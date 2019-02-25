package types

// BibleTexts represents a bibleTexts array containing
// many book objects with its chapters and verses
type BibleTexts struct {
	BibleTexts []BibleText `json:"bibleTexts"`
}

// BibleText represents a bibleText json schema containing
// a book abbreviation and a array of chapters with its verses
type BibleText struct {
	Abbrev   string     `json:"abbrev"`
	Chapters [][]string `json:"chapters"`
}
