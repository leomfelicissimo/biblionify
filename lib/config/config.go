package config

import (
	"github.com/leomfelicissimo/biblionify/lib/types"
	"github.com/leomfelicissimo/biblionify/lib/util"

	"github.com/spf13/viper"
)

// LoadConfiguration returns PORT and SOURCE_PATH from a config.yaml file
func LoadConfiguration() *types.Configuration {
	viper.AddConfigPath(".")
	viper.SetConfigFile("config.yaml")

	viper.SetDefault("port", "3927")
	viper.SetDefault("source_path", "./biblefiles")
	viper.SetDefault("elastic_search_url", "http://localhost:9200")

	err := viper.ReadInConfig()
	util.HandleError(err, "loading config")

	return &types.Configuration{
		Port:             viper.GetInt("port"),
		SourcePath:       viper.GetString("source_path"),
		ElasticSearchURL: viper.GetString("elastic_search_url"),
	}
}
