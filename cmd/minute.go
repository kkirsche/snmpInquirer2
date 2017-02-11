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

	"github.com/Sirupsen/logrus"
	"github.com/k-sone/snmpgo"
	"github.com/spf13/cobra"
	"ni.vzbi.com/stash/scm/ncsddos/inquirer2/libinquirer"
)

// minuteCmd represents the minute command
var minuteCmd = &cobra.Command{
	Use:   "minute",
	Short: "Per-minute polling, for use via cron",
	Long: `Minute is used for per minute polling, most commonly via cron or
another automated service. This does not automate the timing, a tool like cron
must be used to loop this every minute.`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := libinquirer.ParseConfigFile(cfgFile)
		if err != nil {
			log.WithError(err).Errorln("Failed to parse configuration file")
		}
		log.WithField("requested_poll_qty", len(conf.Poll)).Infof("%s poll configurations provided", cfgFile)
		for i, cfg := range conf.Poll {
			logrus.WithField("iteration", i).Debugln("Beginning poll process")
			logrus.WithField("version", cfg.Version).Debugln("SNMP version requested")
			sv := libinquirer.NewVersion(cfg.Version)

			var auth *libinquirer.SNMPAuth
			if sv.Get() == "v3" {
				auth, err = libinquirer.NewAuth(cfg.Username, cfg.SecurityLevel, cfg.AuthPassword, cfg.AuthProtocol, cfg.PrivPassword, cfg.PrivProtocol)
				if err != nil {
					log.WithError(err).Errorln("Failed to create SNMP V3 authentication object")
					return
				}
			}

			logrus.WithField("version", cfg.Version).Debugln("SNMP version accepted")

			client, err := libinquirer.CreateClient(cfg.Host, cfg.Community, cfg.Retries, sv, auth)
			if err != nil {
				log.WithError(err).Errorln("Failed to create SNMP client")
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

			oids, err := snmpgo.NewOids(stroids)
			if err != nil {
				log.WithError(err).Errorln("Failed to create OIDs object")
				continue
			}
			logrus.Debugln("OID object successfully created")

			logrus.WithField("host", cfg.Host).Debugln("Creating client connection to host")
			err = client.Open()
			if err != nil {
				log.WithError(err).Errorln("Failed to open SNMP connection")
				continue
			}
			logrus.Debugln("Client connection created successfully")

			logrus.WithFields(logrus.Fields{
				"nonrepeaters":   0,
				"maxrepetitions": 10,
			}).Debugln("Beginning bulk walk")
			pdu, err := client.GetBulkWalk(oids, 0, 10)
			if err != nil {
				log.WithError(err).Errorln("Failed to execute bulk walk request")
				client.Close()
				continue
			}

			if pdu.ErrorStatus() != snmpgo.NoError {
				log.WithFields(logrus.Fields{
					"error_index": pdu.ErrorIndex(),
				}).Errorln(pdu.ErrorStatus())
				client.Close()
				continue
			}

			logrus.Debugln("Bulk walk completed successfully")
			logrus.Debugln("Outputting result values")
			for _, val := range pdu.VarBinds() {
				splitOID := strings.Split(val.Oid.String(), ".")
				oid := strings.Join(splitOID[:len(splitOID)-1], ".")
				intIndex := strings.Join(splitOID[len(splitOID)-1:len(splitOID)], ".")
				log.WithFields(logrus.Fields{
					"full_oid":        val.Oid,
					"host_queried":    cfg.Host,
					"oid":             oid,
					"oid_name":        cfg.OIDs[fmt.Sprintf(".%s", oid)],
					"interface_index": intIndex,
					"type":            val.Variable.Type(),
					"value":           val.Variable,
				}).Infoln("OID successfully retrieved")
			}
			client.Close()
		}
	},
}

func init() {
	pollCmd.AddCommand(minuteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// minuteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// minuteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
