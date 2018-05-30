package libinquirer

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestParseValidConfigFile(t *testing.T) {
	cwd, _ := os.Getwd()
	c, err := ParseConfigFile(fmt.Sprintf("%s/fixtures/inquirer.json", path.Dir(cwd)))
	if err != nil {
		logrus.WithError(err).Errorln("Failed to parse configuration file")
		t.FailNow()
	}

	if c.Poll[0].Community != testCommunity {
		logrus.Errorln("Parsed community string is invalid")
		t.Fail()
	}

	if c.Poll[0].Host != localhost {
		logrus.Errorln("Parsed community string is invalid")
		t.Fail()
	}

	if c.Poll[0].OIDs == nil {
		logrus.Errorln("Parsed OIDs are invalid")
		t.Fail()
	}

	if c.Poll[0].Version != v2 {
		logrus.Errorln("Parsed version string is invalid")
		t.Fail()
	}
}

func TestParseInvalidConfigFile(t *testing.T) {
	cwd, _ := os.Getwd()
	_, err := ParseConfigFile(fmt.Sprintf("%s/fixtures/invalid_inquirer.json", path.Dir(cwd)))
	if err == nil {
		logrus.WithError(err).Errorln("Failed to parse configuration file")
		t.FailNow()
	}
}

func TestParseNonExistantConfigFile(t *testing.T) {
	_, err := ParseConfigFile(testCommunity)
	if err == nil {
		logrus.WithError(err).Errorln("Parsed non-existant configuration file")
		t.FailNow()
	}
}
