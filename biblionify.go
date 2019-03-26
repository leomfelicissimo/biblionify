package main

import (
	"fmt"

	"github.com/leomfelicissimo/biblionify/lib/indexer"
	"github.com/leomfelicissimo/biblionify/lib/parser"
)

func main() {
	// filePath := "./biblefiles/nvi.json"
	outputPath := "/Users/leonardo/go/src/github.com/leomfelicissimo/biblionify/biblefiles/biblion.nvi.json"
	// parser.TransformAndSave(filePath, outputPath)
	// fmt.Println("The file was transformed")

	jsonData := parser.LoadBiblionTextsFromFile(outputPath)
	indexer.IndexBiblionData("http://localhost:9200", jsonData.BiblionTexts)
	fmt.Println("The file was indexed")
}
