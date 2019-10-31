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
	"fmt"
	"github.com/Adron/cobra-cli-samples/configMgmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "The 'delete' subcommand removes a key value pair from the configuration file.",
	Long:  `The 'delete' subcommand removes a key value pair from the configuration file.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		fmt.Printf("\n\n    **** Deleting key: %s ****\n\n", key)
		configMgmt.ConfigKeyValuePairDelete(key)
	},
}

func init() {
	configCmd.AddCommand(deleteCmd)
}
