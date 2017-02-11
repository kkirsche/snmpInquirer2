package libinquirer

import (
	"testing"

	"github.com/Sirupsen/logrus"
)

// Valid Client Configurations
func TestCreateV1SNMPClient(t *testing.T) {
	v := &SNMPVersion{V1: true, V2: false, V3: false}
	client, err := CreateClient(localhost, testCommunity, 1, v, nil)
	if err != nil {
		logrus.WithError(err).Errorln("Could not create SNMP client")
		t.Fail()
	}
	client.Close()
}

func TestCreateV2cSNMPClient(t *testing.T) {
	v := &SNMPVersion{V1: false, V2: true, V3: false}
	client, err := CreateClient(localhost, testCommunity, 1, v, nil)
	if err != nil {
		logrus.WithError(err).Errorln("Could not create SNMP client")
		t.Fail()
	}
	client.Close()
}

func TestCreateV3SNMPClient(t *testing.T) {
	v := &SNMPVersion{V1: false, V2: false, V3: true}
	a, _ := NewAuth("user", authnopriv, "auth_pass", sha, "priv_pass", aes)
	client, err := CreateClient(localhost, testCommunity, 1, v, a)
	if err != nil {
		logrus.WithError(err).Errorln("Could not create SNMP client")
		t.Fail()
	}
	client.Close()
}

// Invalid Client Configurations
func TestCreateInvalidSNMPClient(t *testing.T) {
	v := &SNMPVersion{V1: false, V2: false, V3: false}
	_, err := CreateClient(localhost, testCommunity, 1, v, nil)
	if err == nil {
		logrus.WithError(err).Errorln("SNMP client erroneously created")
		t.Fail()
	}
}
