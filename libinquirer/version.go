package libinquirer

import "github.com/pkg/errors"

const (
	v1 = "v1"
	v2 = "v2c"
	v3 = "v3"
)

// SNMPVersion allows us to contain booleans for the different versions and
// perform simpler validation
type SNMPVersion struct {
	V1 bool
	V2 bool
	V3 bool
}

// NewVersion creates a new SNMP version object
func NewVersion(v string) *SNMPVersion {
	sv := SNMPVersion{V1: false, V2: false, V3: false}
	if v == v1 {
		sv.V1 = true
	}

	if v == v2 {
		sv.V2 = true
	}

	if v == v3 {
		sv.V3 = true
	}

	return &sv
}

// Validate is used to validate that only one version of SNMP
// has been enabled for use
func (s *SNMPVersion) Validate() error {
	if s.V1 && s.V2 && s.V3 {
		return errors.Errorf("Please select only one of the three SNMP versions")
	}

	if !s.V1 && !s.V2 && !s.V3 {
		return errors.Errorf("Please select at one version of SNMP to use")
	}

	if s.V1 && s.V2 {
		return errors.Errorf("SNMP v1 and v2c may not both be selected")
	}

	if s.V2 && s.V3 {
		return errors.Errorf("SNMP v2c and v3 may not both be selected")
	}

	if s.V1 && s.V3 {
		return errors.Errorf("SNMP v1 and v3 may not both be selected")
	}

	return nil
}

// Get is used to retrieve a string representation of the current SNMP version
func (s *SNMPVersion) Get() string {
	if s.V1 {
		return v1
	}

	if s.V2 {
		return v2
	}

	if s.V3 {
		return v3
	}

	return ""
}
