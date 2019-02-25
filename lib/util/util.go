package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"biblionify/lib/types"

	"github.com/spf13/viper"
)

// HandleError get a error object if it not is nil
// and throws a panic showing a description declared on event
func HandleError(e error, event string) {
	if e != nil {
		panic(fmt.Errorf("One error has ocurred on %s: %s", event, e))
	}
}

// LoadConfiguration returns PORT and SOURCE_PATH from a config.yaml file
func LoadConfiguration() (string, string) {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("PORT", "3927")
	viper.SetDefault("SOURCE_PATH", "./biblefiles")

	err := viper.ReadInConfig()
	HandleError(err, "loading config")

	return viper.GetString("PORT"), viper.GetString("SOURCE_PATH")
}

// ParseJSONFile gets a json file from a certain path parsing it
func ParseJSONFile(sourcePath string) types.BibleTexts {
	jsonContent, err := ioutil.ReadFile(sourcePath)
	HandleError(err, "Loading json file")

	var bibleTexts types.BibleTexts
	errMarshal := json.Unmarshal(jsonContent, &bibleTexts)
	HandleError(errMarshal, "Loading json and parsing it")

	return bibleTexts
}

// GenericFunction is a type for use functions like arguments with
// generic parameters
type GenericFunction func(interface{}) interface{}

// Mapper get a arr and apply a delegate on each item resulting on a new value
func Mapper(arr []interface{}, delegate GenericFunction) []interface{} {
	result := make([]interface{}, len(arr))
	for i, v := range arr {
		result[i] = delegate(v)
	}

	return result
}
