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
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	// BuildHash is the git hash that was used to create the binary
	BuildHash string
	// BuildTime is the date and time that the binary was built
	BuildTime string
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The version of the binary",
	Long: `The build date and build hash associated with the build to allow for
	better identification of when the binary was made and what features it
	offers`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Printf("Version:\t%s", runtime.Version())
		logrus.Printf("Git Hash:\t%s", BuildHash)
		logrus.Printf("Build Time:\t%s", BuildTime)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
