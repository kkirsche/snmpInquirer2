package libinquirer

import (
	"testing"

	"github.com/juju/errors"
	"github.com/sirupsen/logrus"
)

// Valid Configurations
func TestValidNewSNMPV1(t *testing.T) {
	s := NewVersion(v1)
	if s.V1 != true {
		logrus.Errorln("Failed to create New SNMPVersion for v1")
		t.Fail()
	}

	if s.V2 == true || s.V3 == true {
		logrus.Errorln("Incorrectly set versions with v1 target")
		t.Fail()
	}
}

func TestValidNewSNMPV2(t *testing.T) {
	s := NewVersion(v2)
	if s.V2 != true {
		logrus.Errorln("Failed to create New SNMPVersion for v1")
		t.Fail()
	}

	if s.V1 == true || s.V3 == true {
		logrus.Errorln("Incorrectly set versions with v1 target")
		t.Fail()
	}
}

func TestValidNewSNMPV3(t *testing.T) {
	s := NewVersion(v3)
	if s.V3 != true {
		logrus.Errorln("Failed to create New SNMPVersion for v1")
		t.Fail()
	}

	if s.V1 == true || s.V2 == true {
		logrus.Errorln("Incorrectly set versions with v1 target")
		t.Fail()
	}
}

func TestInvalidNewSNMP(t *testing.T) {
	s := NewVersion(invalid)

	if s.V1 == true || s.V2 == true || s.V3 == true {
		logrus.Errorln("Incorrectly set versions with invalid target")
		t.Fail()
	}
}

func TestValidSNMPValidationV1Only(t *testing.T) {
	s := SNMPVersion{V1: true, V2: false, V3: false}
	err := s.Validate()

	if err != nil {
		logrus.WithError(err).Errorln(errors.ErrorStack(err))
		t.Fail()
	}
}

func TestV1SNMPGet(t *testing.T) {
	s := SNMPVersion{V1: true, V2: false, V3: false}
	v := s.Get()

	if v != v1 {
		logrus.WithField("version", v).Errorln("Incorrect SNMP version returned")
		t.Fail()
	}
}

func TestValidSNMPValidationV2cOnly(t *testing.T) {
	s := SNMPVersion{V1: false, V2: true, V3: false}
	err := s.Validate()

	if err != nil {
		logrus.WithError(err).Errorln(errors.ErrorStack(err))
		t.Fail()
	}
}

func TestV2cSNMPGet(t *testing.T) {
	s := SNMPVersion{V1: false, V2: true, V3: false}
	v := s.Get()

	if v != v2 {
		logrus.WithField("version", v).Errorln("Incorrect SNMP version returned")
		t.Fail()
	}
}

func TestValidSNMPValidationV3Only(t *testing.T) {
	s := SNMPVersion{V1: false, V2: false, V3: true}
	err := s.Validate()

	if err != nil {
		logrus.WithError(err).Errorln(errors.ErrorStack(err))
		t.Fail()
	}
}

func TestV3SNMPGet(t *testing.T) {
	s := SNMPVersion{V1: false, V2: false, V3: true}
	v := s.Get()

	if v != v3 {
		logrus.WithField("version", v).Errorln("Incorrect SNMP version returned")
		t.Fail()
	}
}

// Invalid Configurations
func TestInvalidSNMPValidationNoSelection(t *testing.T) {
	s := SNMPVersion{V1: false, V2: false, V3: false}
	err := s.Validate()

	if err == nil {
		logrus.WithError(err).Errorln(errors.ErrorStack(err))
		t.Fail()
	}
}

func TestNoVersionGet(t *testing.T) {
	s := SNMPVersion{V1: false, V2: false, V3: false}
	v := s.Get()

	if v != "" {
		logrus.Errorln("No SNMP version returned")
		t.Fail()
	}
}

func TestInvalidSNMPValidationAllVersions(t *testing.T) {
	s := SNMPVersion{V1: true, V2: true, V3: true}
	err := s.Validate()

	if err == nil {
		logrus.WithError(err).Errorln(errors.ErrorStack(err))
		t.Fail()
	}
}

func TestInvalidSNMPValidationV1andV2Only(t *testing.T) {
	s := SNMPVersion{V1: true, V2: true, V3: false}
	err := s.Validate()

	if err == nil {
		logrus.WithError(err).Errorln(errors.ErrorStack(err))
		t.Fail()
	}
}

func TestInvalidSNMPValidationV2andV3Only(t *testing.T) {
	s := SNMPVersion{V1: false, V2: true, V3: true}
	err := s.Validate()

	if err == nil {
		logrus.WithError(err).Errorln(errors.ErrorStack(err))
		t.Fail()
	}
}

func TestInvalidSNMPValidationV1andV3Only(t *testing.T) {
	s := SNMPVersion{V1: true, V2: false, V3: true}
	err := s.Validate()

	if err == nil {
		logrus.WithError(err).Errorln(errors.ErrorStack(err))
		t.Fail()
	}
}
