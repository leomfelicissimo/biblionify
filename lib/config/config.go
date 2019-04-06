package config

import (
	"github.com/leomfelicissimo/biblionify/lib/util"

	"github.com/spf13/viper"
)

// Configuration represents the data loaded from a config.yaml file
type Configuration struct {
	Port             int
	SourcePath       string
	ElasticSearchURL string
	MongoURL         string
	MongoDatabase    string
	MongoCollection  string
}

// LoadConfiguration returns a Configuration struct
func LoadConfiguration() *Configuration {
	// viper.AddConfigPath(".")
	// viper.SetConfigName("config.yml")
	// viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yml")

	viper.SetDefault("port", "3927")
	viper.SetDefault("source_path", "../../biblefiles")
	viper.SetDefault("elastic_search_url", "http://localhost:9200")
	viper.SetDefault("mongodb_url", "http://localhost:27017")

	err := viper.ReadInConfig()
	util.HandleError(err, "Loading Configuration")

	return &Configuration{
		Port:             viper.GetInt("port"),
		SourcePath:       viper.GetString("source_path"),
		ElasticSearchURL: viper.GetString("elastic_search_url"),
		MongoURL:         viper.GetString("mongodb_url"),
		MongoDatabase:    viper.GetString("mongodb_database"),
		MongoCollection:  viper.GetString("mongodb_collection"),
	}
}
