package libinquirer

import (
	"github.com/Sirupsen/logrus"
	"github.com/k-sone/snmpgo"
	"github.com/pkg/errors"
)

const (
	noauthnopriv = "NoAuthNoPriv"
	authnopriv   = "AuthNoPriv"
	authpriv     = "AuthPriv"
	md5          = "MD5"
	sha          = "SHA"
	des          = "DES"
	aes          = "AES"

	securityLevel          = "security_level"
	authenticationProtocol = "auth_proto"
	privateProtocol        = "priv_proto"
)

// SNMPAuth is used to collect the necessary data to authenticate over SNMP v3
// connections
type SNMPAuth struct {
	Username      string
	SecurityLevel snmpgo.SecurityLevel
	AuthPassword  string
	AuthProtocol  snmpgo.AuthProtocol
	PrivPassword  string
	PrivProtocol  snmpgo.PrivProtocol
}

// auth is used when parsing the user's configuration file
type auth struct {
	Username      string `json:"username"`
	SecurityLevel string `json:"security_level"`
	AuthPassword  string `json:"auth_password"`
	AuthProtocol  string `json:"auth_protocol"`
	PrivPassword  string `json:"priv_password"`
	PrivProtocol  string `json:"priv_protocol"`
}

func retrieveSecurityLevel(s string) (snmpgo.SecurityLevel, error) {
	switch s {
	case noauthnopriv:
		logrus.WithField(securityLevel, noauthnopriv).Debugln("Security level set")
		return snmpgo.NoAuthNoPriv, nil
	case authnopriv:
		logrus.WithField(securityLevel, authnopriv).Debugln("Security level set")
		return snmpgo.AuthNoPriv, nil
	case authpriv:
		logrus.WithField(securityLevel, authpriv).Debugln("Security level set")
		return snmpgo.AuthPriv, nil
	default:
		logrus.WithField(securityLevel, s).Debugln("Invalid security level detected")
		return snmpgo.AuthPriv, errors.Errorf("Invalid security level. Please select NoAuthNoPriv, AuthNoPriv, or AuthPriv")
	}
}

func retrieveAuthProto(a string) (snmpgo.AuthProtocol, error) {
	switch a {
	case md5:
		logrus.WithField(authenticationProtocol, md5).Debugln("Authentication protocol set")
		return snmpgo.Md5, nil
	case sha:
		logrus.WithField(authenticationProtocol, sha).Debugln("Authentication protocol set")
		return snmpgo.Sha, nil
	default:
		logrus.WithField(authenticationProtocol, a).Debugln("Invalid authentication protocol detected")
		return snmpgo.Sha, errors.Errorf("Invalid auth protocol. Please select MD5 or SHA")
	}
}

func retrievePrivProto(p string) (snmpgo.PrivProtocol, error) {
	switch p {
	case des:
		logrus.WithField(privateProtocol, des).Debugln("Private communication protocol set")
		return snmpgo.Des, nil
	case aes:
		logrus.WithField(privateProtocol, aes).Debugln("Private communication protocol set")
		return snmpgo.Aes, nil
	default:
		logrus.WithField(privateProtocol, p).Debugln("Invalid authentication protocol detected")
		return snmpgo.Aes, errors.Errorf("Invalid private communication protocol. Please select DES or AES")
	}
}

// NewAuth is used to create the proper SNMP authentication object for use when
// using SNMP v3
func NewAuth(u, s, apass, a, ppass, p string) (*SNMPAuth, error) {
	sl, err := retrieveSecurityLevel(s)
	if err != nil {
		return nil, err
	}

	aproto, err := retrieveAuthProto(a)
	if err != nil {
		return nil, err
	}

	pproto, err := retrievePrivProto(p)
	if err != nil {
		return nil, err
	}

	return &SNMPAuth{
		Username:      u,
		SecurityLevel: sl,
		AuthPassword:  apass,
		AuthProtocol:  aproto,
		PrivPassword:  ppass,
		PrivProtocol:  pproto,
	}, nil
}
