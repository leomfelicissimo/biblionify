package mongoindexer

import (
	"context"
	"strconv"
	"time"

	"github.com/leomfelicissimo/biblionify/lib/config"
	"github.com/leomfelicissimo/biblionify/lib/types"
	"github.com/leomfelicissimo/biblionify/lib/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getMongoClient(uri string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	util.HandleError(err, "Error connecting at mongo URI")
	return client
}

func parseStructToBson(biblionTexts []types.BiblionText) []interface{} {
	var mongoData []interface{}
	for _, text := range biblionTexts {
		id := text.Book + strconv.Itoa(text.Chapter) + strconv.Itoa(text.Verse)
		var data = map[string]interface{}{
			"reference": id,
			"book":      text.Book,
			"chapter":   text.Chapter,
			"verse":     text.Verse,
			"text":      text.Text,
		}
		mongoData = append(mongoData, data)
	}

	return mongoData
}

func readConfigurationData(config *config.Configuration) (string, string, string) {
	return config.MongoURL, config.MongoDatabase, config.MongoCollection
}

// IndexBiblionData get a array of biblion texts struct, and import basead a certain configuration
func IndexBiblionData(biblionTexts []types.BiblionText, config *config.Configuration) bool {
	url, databaseName, collectionName := readConfigurationData(config)
	client := getMongoClient(url)
	collection := client.Database(databaseName).Collection(collectionName)

	data := parseStructToBson(biblionTexts)
	ctx := context.Background()
	res, err := collection.InsertMany(ctx, data)

	util.HandleError(err, "Error on indexing data at mongo")

	return len(res.InsertedIDs) != 0
}
