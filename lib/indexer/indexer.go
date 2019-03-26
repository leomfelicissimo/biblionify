package indexer

import (
	"context"
	"fmt"
	"strconv"

	"github.com/leomfelicissimo/biblionify/lib/types"
	"github.com/leomfelicissimo/biblionify/lib/util"
	"github.com/olivere/elastic"
)

const mapping = `
{
    "settings": {
        "number_of_shards": 1,
        "number_of_replicas": 0
    },
    "mappings": {
        "bibleText": {
            "properties": {
                "book": { "type": "keyword" },
                "chapter": { "type": "long" },
                "verse": { "type": "long" },
                "text": { "type": "text" }
            }
        }
    }
}`

func createElasticSearchClient(elasticURL string) *elastic.Client {
	client, err := elastic.NewClient(
		elastic.SetURL(elasticURL),
		elastic.SetSniff(false),
	)

	util.HandleError(err, "Getting ElasticSearch client")

	return client
}

func getOrCreateIndex(client *elastic.Client, name string) *elastic.IndexService {
	ctx := context.Background()

	indexExists, err := client.IndexExists(name).Do(ctx)
	util.HandleError(err, "Verifiying if an index exists")

	if !indexExists {
		createIndex, err := client.CreateIndex(name).BodyString(mapping).Do(ctx)
		util.HandleError(err, "Creating a new index")

		if !createIndex.Acknowledged {
			panic("Cannot create a new bibleText index. Creation not acknowledged.")
		}
	}

	return client.Index()
}

func HealthSearchEngine(elasticURL string) {
	ctx := context.Background()

	client := createElasticSearchClient(elasticURL)

	info, code, err := client.Ping(elasticURL).Do(ctx)
	util.HandleError(err, "Pinging ElasticSearch")
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	esversion, err := client.ElasticsearchVersion(elasticURL)
	util.HandleError(err, "Getting ElasticSearch version")
	fmt.Printf("Elasticsearch version %s\n", esversion)

}

func IndexBiblionData(elasticSearchURL string, biblionTexts []types.BiblionText) {
	ctx := context.Background()
	// TODO: Load from configuration env
	// conf := config.LoadConfiguration()
	client := createElasticSearchClient(elasticSearchURL)
	index := getOrCreateIndex(client, "nvi")

	for _, text := range biblionTexts {
		id := text.Book + strconv.Itoa(text.Chapter) + strconv.Itoa(text.Verse)
		_, err := index.Index("nvi").
			Type("bibleText").
			Id(id).
			BodyJson(text).
			Do(ctx)

		util.HandleError(err, "Putting data on index")
	}
}
