package config

import (
	"biblionify/lib/util"

	"github.com/spf13/viper"
)

// LoadConfiguration returns PORT and SOURCE_PATH from a config.yaml file
func LoadConfiguration() (string, string) {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("PORT", "3927")
	viper.SetDefault("SOURCE_PATH", "./biblefiles")

	err := viper.ReadInConfig()
	util.HandleError(err, "loading config")

	return viper.GetString("PORT"), viper.GetString("SOURCE_PATH")
}
