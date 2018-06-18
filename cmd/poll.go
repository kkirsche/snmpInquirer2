// Copyright © 2017 Kevin Kirsche <kev.kirsche[at]gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"strings"

	"github.com/kkirsche/snmpInquirer2/libinquirer"
	"github.com/sirupsen/logrus"
	"github.com/soniah/gosnmp"
	"github.com/spf13/cobra"
)

// pollCmd represents the minute command
var pollCmd = &cobra.Command{
	Use:   "poll",
	Short: "SNMP polling for use via cron",
	Long: `Minute is used for per minute polling, most commonly via cron or
another automated service. This does not automate the timing, a tool like cron
must be used to loop this every minute.`,
	Run: func(cmd *cobra.Command, args []string) {
		var constLookup = map[gosnmp.Asn1BER]string{
			gosnmp.UnknownType:       "UnknownType",
			gosnmp.Boolean:           "Boolean",
			gosnmp.Integer:           "Integer",
			gosnmp.BitString:         "BitString",
			gosnmp.OctetString:       "OctetString",
			gosnmp.Null:              "Null",
			gosnmp.ObjectIdentifier:  "ObjectIdentifier",
			gosnmp.ObjectDescription: "ObjectDescription",
			gosnmp.IPAddress:         "IPAddress",
			gosnmp.Counter32:         "Counter32",
			gosnmp.Gauge32:           "Gauge32",
			gosnmp.TimeTicks:         "TimeTicks",
			gosnmp.Opaque:            "Opaque",
			gosnmp.NsapAddress:       "NsapAddress",
			gosnmp.Counter64:         "Counter64",
			gosnmp.Uinteger32:        "Uinteger32",
			gosnmp.OpaqueFloat:       "OpaqueFloat",
			gosnmp.OpaqueDouble:      "OpaqueDouble",
			gosnmp.NoSuchObject:      "NoSuchObject",
			gosnmp.NoSuchInstance:    "NoSuchInstance",
			gosnmp.EndOfMibView:      "EndOfMibView",
		}

		conf, err := libinquirer.ParseConfigFile(cfgFile)
		if err != nil {
			logrus.WithError(err).Errorln("Failed to parse configuration file")
		}

		logrus.WithField("requested_poll_qty", len(conf.Poll)).Infof("%s poll configurations provided", cfgFile)
		for i, cfg := range conf.Poll {
			logrus.WithField("iteration", i).Debugln("Beginning poll process")
			logrus.WithField("version", cfg.Version).Debugln("SNMP version requested --- Validating")
			sv := libinquirer.NewVersion(cfg.Version)

			var auth *libinquirer.SNMPAuth
			if sv.Get() == "v3" {
				auth, err = libinquirer.NewAuth(cfg.Username, cfg.SecurityLevel, cfg.AuthPassword, cfg.AuthProtocol, cfg.PrivPassword, cfg.PrivProtocol)
				if err != nil {
					logrus.WithError(err).Errorln("Failed to create SNMP V3 authentication object")
					return
				}
			}

			logrus.WithField("version", cfg.Version).Debugln("SNMP version accepted")

			logrus.WithFields(logrus.Fields{
				"host":             cfg.Host,
				"community_string": cfg.Community,
				"retries":          cfg.Retries,
				"version":          cfg.Version,
			}).Debugln("Creating SNMP client")

			client, err := libinquirer.CreateClient(cfg.Host, cfg.Community, cfg.Retries, sv, auth)
			if err != nil {
				logrus.WithError(err).Errorln("Failed to create SNMP client")
				return
			}
			logrus.Debugln("SNMP client created")

			logrus.Debugln("Generating OID object for querying process")
			stroids := []string{}
			for oid, mib := range cfg.OIDs {
				logrus.WithFields(logrus.Fields{
					"oid": oid,
					"mib": mib,
				}).Debugln("Adding object for querying")
				stroids = append(stroids, oid)
			}

			logrus.WithFields(logrus.Fields{
				"host": cfg.Host,
			}).Debugln("Creating client connection to host")
			err = client.Connect()
			if err != nil {
				logrus.WithError(err).Errorln("Failed to open SNMP connection")
				continue
			}
			defer client.Conn.Close()
			logrus.Debugln("Client connection created successfully")

			logrus.WithFields(logrus.Fields{
				"nonrepeaters":   0,
				"maxrepetitions": 10,
			}).Debugln("Beginning bulk walk")
			for _, oid := range stroids {
				pdus, err := client.BulkWalkAll(oid)
				if err != nil {
					logrus.WithError(err).Errorln("Failed to execute bulk walk request")
					continue
				}
				logrus.Debugln("PDU's retrieved, checking for PDU error(s)")
				for _, pdu := range pdus {
					logrus.Debugln("Bulk walk completed successfully")
					logrus.Debugln("Outputting result values")
					splitOID := strings.Split(pdu.Name, ".")
					intIndex := strings.Join(splitOID[len(splitOID)-1:len(splitOID)], ".")
					switch pdu.Type {
					case gosnmp.OctetString:
						logrus.WithFields(logrus.Fields{
							"full_oid":        pdu.Name,
							"host_queried":    cfg.Host,
							"oid":             oid,
							"oid_name":        cfg.OIDs[oid],
							"interface_index": intIndex,
							"pdu_type":        fmt.Sprintf("0x%x", pdu.Type),
							"pdu_type_name":   constLookup[pdu.Type],
							"value":           string(pdu.Value.([]byte)),
						}).Infoln("OID successfully retrieved")
					default:
						logrus.WithFields(logrus.Fields{
							"full_oid":        pdu.Name,
							"host_queried":    cfg.Host,
							"oid":             oid,
							"oid_name":        cfg.OIDs[oid],
							"interface_index": intIndex,
							"pdu_type":        fmt.Sprintf("0x%x", pdu.Type),
							"pdu_type_name":   constLookup[pdu.Type],
							"value":           gosnmp.ToBigInt(pdu.Value),
						}).Infoln("OID successfully retrieved")
					}

				}
			}
			logrus.Debugln("Host output complete")
		}
	},
}

func init() {
	RootCmd.AddCommand(pollCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pollCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pollCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
