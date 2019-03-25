package main

import (
	"fmt"

	"github.com/leomfelicissimo/biblionify/lib/parser"
)

func main() {
	filePath := "./biblefiles/nvi.json"
	outputPath := "/Users/leonardo/go/src/biblionify/biblefiles/biblion.nvi.json"
	parser.TransformAndSave(filePath, outputPath)
	fmt.Println("The file was transformed")
}
