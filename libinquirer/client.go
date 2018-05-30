package libinquirer

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/soniah/gosnmp"
)

// CreateClient is used to generate a SNMP client to query one or more hosts
// for host metrics
func CreateClient(a, c string, r int, vers *SNMPVersion, auth *SNMPAuth) (*gosnmp.GoSNMP, error) {
	if err := vers.Validate(); err != nil {
		return nil, err
	}

	var v gosnmp.SnmpVersion
	vs := vers.Get()
	switch vs {
	case v1:
		v = gosnmp.Version1
		logrus.WithField("version", v1).Debugln("SNMP Version 1 enabled")
	case v2:
		v = gosnmp.Version2c
		logrus.WithField("version", v2).Debugln("SNMP Version 2c enabled")
	case v3:
		v = gosnmp.Version3
		logrus.WithField("version", v3).Debugln("SNMP Version 3 enabled")
	}

	if vs == v1 || vs == v2 {
		params := &gosnmp.GoSNMP{
			Version:   v,
			Target:    a,
			Port:      161,
			Timeout:   time.Duration(30) * time.Second,
			Retries:   r,
			Community: c,
		}
		return params, nil
	}

	params := &gosnmp.GoSNMP{
		Target:        a,
		Port:          161,
		Version:       v,
		Timeout:       time.Duration(30) * time.Second,
		SecurityModel: gosnmp.UserSecurityModel,
		Community:     c,
		Retries:       r,
		MsgFlags:      auth.SecurityLevel,
		SecurityParameters: &gosnmp.UsmSecurityParameters{
			UserName:                 auth.Username,
			AuthenticationProtocol:   auth.AuthProtocol,
			AuthenticationPassphrase: auth.AuthPassword,
			PrivacyProtocol:          auth.PrivProtocol,
			PrivacyPassphrase:        auth.PrivPassword,
		},
	}

	return params, nil
}
