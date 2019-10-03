/*
Copyright © 2019 Adron Hall <adron@thrashingcode.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/Adron/cobra-cli-samples/configMgmt"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "The 'update' subcommand will update a passed in key value pair for an existing set of data to the application configuration file.",
	Long: `The 'update' subcommand updates a key value pair, if the key value pair already exists it is updated, if it does
not exist then the passed in values are added to the application configuration file. For example:

'<cmd> config add --key theKey --value "the new value which will be updated for this particular key value pair."'.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		value, _ := cmd.Flags().GetString("value")
		configMgmt.ConfigKeyValuePairUpdate(key, value)
	},
}

func init() {
	configCmd.AddCommand(updateCmd)
}
