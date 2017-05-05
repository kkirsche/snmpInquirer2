package libinquirer

import (
	"encoding/json"
	"os"

	"github.com/Sirupsen/logrus"
)

// Configuration object for the inquirer tool
type Configuration struct {
	Poll []PollConfiguration `json:"poll"`
}

// PollConfiguration represents the configuration on a host by host basis for
// the inquirer tool
type PollConfiguration struct {
	Community string            `json:"community"`
	Host      string            `json:"host"`
	Version   string            `json:"version"`
	OIDs      map[string]string `json:"oids"`
	Retries   int               `json:"retries"`
	auth
}

// ParseConfigFile is used to retrieve an SNMP configuration object from
func ParseConfigFile(c string) (*Configuration, error) {
	configFile, err := os.Open(c)
	if err != nil {
		logrus.WithError(err).Debugln("Could not open configuration file")
		return nil, err
	}
	defer configFile.Close()

	var conf Configuration
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&conf); err != nil {
		logrus.WithError(err).Debugln("Could not parse configuration file")
		return nil, err
	}

	return &conf, nil
}
