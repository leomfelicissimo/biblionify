package indexer

import (
	"testing"

	"github.com/leomfelicissimo/biblionify/lib/types"
)

func TestHealthSearchEngine(t *testing.T) {
	HealthSearchEngine("http://localhost:9200")
}

func TestIndexBiblionData(t *testing.T) {
	biblionData := []types.BiblionText{
		{Book: "gn", Chapter: 1, Verse: 1, Text: "No princípio criou Deus os céus e a terra"},
		{Book: "jo", Chapter: 1, Verse: 1, Text: "No princípio era o Verbo"},
	}

	IndexBiblionData("http://localhost:9200", biblionData)
}
