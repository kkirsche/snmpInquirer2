package libinquirer

import (
	"net"

	"github.com/Sirupsen/logrus"
	"github.com/k-sone/snmpgo"
)

// CreateClient is used to generate a SNMP client to query one or more hosts
// for host metrics
func CreateClient(a, c string, r uint, vers *SNMPVersion, auth *SNMPAuth) (*snmpgo.SNMP, error) {
	if err := vers.Validate(); err != nil {
		return nil, err
	}

	var v snmpgo.SNMPVersion
	vs := vers.Get()
	switch vs {
	case v1:
		v = snmpgo.V1
		logrus.WithField("version", v1).Debugln("SNMP Version 1 enabled")
	case v2:
		v = snmpgo.V2c
		logrus.WithField("version", v2).Debugln("SNMP Version 2c enabled")
	case v3:
		v = snmpgo.V3
		logrus.WithField("version", v3).Debugln("SNMP Version 3 enabled")
	}

	addr := net.JoinHostPort(a, "161")

	if vs == v1 || vs == v2 {
		client, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
			Version:   v,
			Address:   addr,
			Retries:   r,
			Community: c,
		})
		return client, err
	}

	client, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version:       v,
		Address:       a,
		Retries:       r,
		Community:     c,
		UserName:      auth.Username,
		SecurityLevel: auth.SecurityLevel,
		AuthPassword:  auth.AuthPassword,
		AuthProtocol:  auth.AuthProtocol,
		PrivPassword:  auth.PrivPassword,
		PrivProtocol:  auth.PrivProtocol,
	})

	return client, err
}
