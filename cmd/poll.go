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

import "github.com/spf13/cobra"

// pollCmd represents the poll command
var pollCmd = &cobra.Command{
	Use:   "poll",
	Short: "Poll is used for time based polling of SNMP OIDs",
	Long: `Poll is used when doing either cron or custom time based SNMP polling of
an external IP addresses for data.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(pollCmd)
}
