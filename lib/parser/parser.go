package parser

import (
	"encoding/json"
	"io/ioutil"

	"github.com/leomfelicissimo/biblionify/lib/types"
	"github.com/leomfelicissimo/biblionify/lib/util"
)

// ParseJSONFile gets a json file from a certain path parsing it
func parseJSONFile(sourcePath string) types.JSONBiblical {
	jsonContent, err := ioutil.ReadFile(sourcePath)
	util.HandleError(err, "Loading json file")

	var bibleTexts types.JSONBiblical
	errMarshal := json.Unmarshal(jsonContent, &bibleTexts)
	util.HandleError(errMarshal, "Loading json and parsing it")

	return bibleTexts
}

func LoadBiblionTextsFromFile(sourcePath string) types.BiblionTexts {
	jsonContent, err := ioutil.ReadFile(sourcePath)
	util.HandleError(err, "Loading json file")

	var biblionTexts types.BiblionTexts
	errMarshal := json.Unmarshal(jsonContent, &biblionTexts)
	util.HandleError(errMarshal, "Loading json and parsing it")

	return biblionTexts
}

// TransformJSON Load a Biblical JSON file and transforms it to a Biblion JSON file
func transformJSON(filePath string) []byte {
	content := parseJSONFile(filePath)
	var texts []types.BiblionText

	for _, b := range content.JSONBiblicalTexts {
		for j, c := range b.Chapters {
			for k, v := range c {
				texts = append(texts, types.BiblionText{
					Book:    b.Abbrev,
					Chapter: j + 1,
					Verse:   k + 1,
					Text:    v,
				})
			}
		}
	}

	b, err := json.Marshal(texts)
	util.HandleError(err, "Transforming JSON")
	return b
}

// TransformAndSave function gets a JSON Biblical file, transforms it and save to an output
func TransformAndSave(filePath string, outputPath string) {
	b := transformJSON(filePath)
	err := ioutil.WriteFile(outputPath, b, 0644)
	util.HandleError(err, "Writing Biblion JSON")
}
