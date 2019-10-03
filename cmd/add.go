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
	"github.com/Adron/cobra-cli-samples/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var theKey string
		theKey, _ = cmd.Flags().GetString("key")
		var theValue string
		theValue, _ = cmd.Flags().GetString("value")

		if len(theKey) == 0 || len(theValue) == 0 {
			fmt.Println("The key and value must both contain contents to write to the configuration file.")
			return
		}

		// Determine if an existing key, if so notify.
		if findExistingKey(theKey) {
			fmt.Println("This key already exists. Create a key value pair with a different key, or if this is an update use the update command.")
			return
		}

		writeKeyValuePair(theKey, theValue)
	},
}

func writeKeyValuePair(key string, value string) {
	viper.Set(key, value)
	err := viper.WriteConfig()
	helper.Check(err)
	fmt.Printf("Wrote the %s pair.\n", key)
}

func findExistingKey(theKey string) bool {
	existingKey := false
	for i := 0; i < len(viper.AllKeys()); i++ {
		if viper.AllKeys()[i] == theKey {
			existingKey = true
		}
	}
	return existingKey
}

func init() {
	configCmd.AddCommand(addCmd)
}
