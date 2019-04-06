package main

import (
	"github.com/leomfelicissimo/biblionify/lib/config"
	"github.com/leomfelicissimo/biblionify/lib/mongoindexer"
	"github.com/leomfelicissimo/biblionify/lib/parser"
)

func main() {
	appConfig := config.LoadConfiguration()
	jsonData := parser.LoadBiblionTextsFromFile(appConfig.SourcePath)
	// firebaseindexer.IndexBiblionData("pt-br", "nvi", jsonData.BiblionTexts)
	// fmt.Println("The file was indexed")
	mongoindexer.IndexBiblionData(jsonData.BiblionTexts, appConfig)
}
