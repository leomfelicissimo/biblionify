package main

import (
	"biblionify/lib/parser"
	"fmt"
)

func main() {
	filePath := "./biblefiles/mini.json"
	outputPath := "/Users/leonardo/go/src/biblionify/biblefiles/biblion.mini.json"
	parser.TransformAndSave(filePath, outputPath)
	fmt.Println("The file was transformed")
}
