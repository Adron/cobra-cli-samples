package configMgmt

import (
	"fmt"
	"github.com/Adron/cobra-cli-samples/helper"
	"github.com/spf13/viper"
	"log"
)

func ConfigKeyValuePairDelete(key string) {

	settings := viper.AllSettings()
	fmt.Println(settings)

	delete(settings, key)

	fmt.Println(settings)

	// How to write back the things minus what I just deleted?!?!?!?!

	// No idea how this is supposed to work!?!?! Ugh!

	// TODO: Please submit a PR with this or fix viper or something?!

	helper.HandleError(viper.WriteConfig())
}

func ConfigKeyValuePairUpdate(key string, value string) {
	writeKeyValuePair(key, value)
}

func ConfigKeyValuePairAdd(key string, value string) {
	if validateKeyValuePair(key, value) {
		log.Printf("Validation not met for %s.", key)
	} else {
		writeKeyValuePair(key, value)
	}
}

func validateKeyValuePair(key string, value string) bool {
	if len(key) == 0 || len(value) == 0 {
		fmt.Println("The key and value must both contain contents to write to the configuration file.")
		return true
	}
	// Determine if an existing key, if so notify.
	if findExistingKey(key) {
		fmt.Println("This key already exists. Create a key value pair with a different key, or if this is an update use the update command.")
		return true
	}
	return false
}

func writeKeyValuePair(key string, value string) {
	viper.Set(key, value)
	err := viper.WriteConfig()
	helper.HandleError(err)
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
