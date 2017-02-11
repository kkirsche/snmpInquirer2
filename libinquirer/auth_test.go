package libinquirer

import (
	"testing"

	"github.com/Sirupsen/logrus"
)

func TestValidNoAuthNoPrivSecurityLevel(t *testing.T) {
	_, err := retrieveSecurityLevel(noauthnopriv)
	if err != nil {
		logrus.WithError(err).Errorln("Failed to correctly set NoAuthNoPriv security level")
		t.Fail()
	}
}

func TestValidAuthNoPrivSecurityLevel(t *testing.T) {
	_, err := retrieveSecurityLevel(authnopriv)
	if err != nil {
		logrus.WithError(err).Errorln("Failed to correctly set AuthNoPriv security level")
		t.Fail()
	}
}

func TestValidAuthPrivSecurityLevel(t *testing.T) {
	_, err := retrieveSecurityLevel(authpriv)
	if err != nil {
		logrus.WithError(err).Errorln("Failed to correctly set AuthPriv security level")
		t.Fail()
	}
}

func TestInvalidSecurityLevel(t *testing.T) {
	_, err := retrieveSecurityLevel(invalid)
	if err == nil {
		logrus.WithError(err).Errorln("Incorrectly set an invalid security level")
		t.Fail()
	}
}

func TestValidMD5RetrieveAuthProto(t *testing.T) {
	_, err := retrieveAuthProto(md5)
	if err != nil {
		logrus.WithError(err).Errorln("Failed to correctly set MD5 authentication protocol")
		t.Fail()
	}
}

func TestValidSHARetrieveAuthProto(t *testing.T) {
	_, err := retrieveAuthProto(sha)
	if err != nil {
		logrus.WithError(err).Errorln("Failed to correctly set SHA authentication protocol")
		t.Fail()
	}
}

func TestInvalidAuthProto(t *testing.T) {
	_, err := retrieveAuthProto(invalid)
	if err == nil {
		logrus.WithError(err).Errorln("Incorrectly set an invalid authentication protocl")
		t.Fail()
	}
}

func TestValidDESRetrievePrivProto(t *testing.T) {
	_, err := retrievePrivProto(des)
	if err != nil {
		logrus.WithError(err).Errorln("Failed to correctly set DES private communication protocol")
		t.Fail()
	}
}

func TestValidAESRetrievePrivProto(t *testing.T) {
	_, err := retrievePrivProto(aes)
	if err != nil {
		logrus.WithError(err).Errorln("Failed to correctly set AES private communication protocol")
		t.Fail()
	}
}

func TestInvalidPrivProto(t *testing.T) {
	_, err := retrievePrivProto(invalid)
	if err == nil {
		logrus.WithError(err).Errorln("Incorrectly set an invalid authentication protocl")
		t.Fail()
	}
}

func TestValidNewAuthStructure(t *testing.T) {
	_, err := NewAuth("test_user", authpriv, "test_auth_pass", sha, "test_priv_pass", aes)
	if err != nil {
		logrus.WithError(err).Errorln("Failed to correctly set DES private communication protocol")
		t.Fail()
	}
}

func TestInvalidNewAuthStructureSecLevel(t *testing.T) {
	_, err := NewAuth("test_user", invalid, "test_auth_pass", sha, "test_priv_pass", aes)
	if err == nil {
		logrus.WithError(err).Errorln("Failed to correctly set DES private communication protocol")
		t.Fail()
	}
}

func TestInvalidNewAuthStructureAuthProto(t *testing.T) {
	_, err := NewAuth("test_user", authpriv, "test_auth_pass", invalid, "test_priv_pass", aes)
	if err == nil {
		logrus.WithError(err).Errorln("Failed to correctly set DES private communication protocol")
		t.Fail()
	}
}

func TestInvalidNewAuthStructurePrivProto(t *testing.T) {
	_, err := NewAuth("test_user", authpriv, "test_auth_pass", sha, "test_priv_pass", invalid)
	if err == nil {
		logrus.WithError(err).Errorln("Failed to correctly set DES private communication protocol")
		t.Fail()
	}
}
