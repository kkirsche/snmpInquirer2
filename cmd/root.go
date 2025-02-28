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
	"log/syslog"
	"os"

	"github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	"github.com/spf13/cobra"
)

var (
	cfgFile        string
	verboseEnabled bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "inquirer2",
	Short: "Multi-host SNMP collection tool",
	Long: `Inquirer is a multi-host SNMP collection tool designed to support
SNMP versions 1, 2c, and 3. It is designed to leverage best practices in
logging formats to be easier to parse and quicker to use than other solutions`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "/etc/shield/snmp/inquirer_2.json", "config file (default is /etc/shield/snmp/inquirer_2.json)")
	RootCmd.PersistentFlags().BoolVarP(&verboseEnabled, "verbose", "v", false, "Enable verbose logging")

	hook, err := logrus_syslog.NewSyslogHook("", "", syslog.LOG_INFO, "Inquirer2")
	if err != nil {
		logrus.WithError(err).Errorln("Unable to create syslog hook")
		return
	}

	if verboseEnabled {
		logrus.SetLevel(logrus.DebugLevel)
	}

	logrus.AddHook(hook)
}
