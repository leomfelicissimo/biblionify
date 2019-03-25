package types

// GenericFunction is a type for use functions like arguments with
// generic parameters
type GenericFunction func(interface{}) interface{}

// BibleTexts represents a bibleTexts array containing
// many book objects with its chapters and verses
type JSONBiblical struct {
	JSONBiblicalTexts []JSONBiblicalText `json:"bibleTexts"`
}

// BibleText represents a bibleText json schema containing
// a book abbreviation and a array of chapters with its verses
type JSONBiblicalText struct {
	Abbrev   string     `json:"abbrev"`
	Chapters [][]string `json:"chapters"`
}

type BiblionTexts struct {
	BiblionText []BiblionText `json:"bibleTexts"`
}

type BiblionText struct {
	Book    string `json:"book"`
	Chapter int    `json:"chapter"`
	Verse   int    `json:"verse"`
	Text    string `json:"text"`
}

type Configuration struct {
	Port             int
	SourcePath       string
	ElasticSearchURL string
}
