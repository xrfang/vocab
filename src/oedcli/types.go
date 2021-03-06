package oed

type note struct {
	Text string
	Type string
}

type textEntry struct {
	Text string
}

type ThesaurusLink struct {
	EntryID string `json:"entry_id"`
	SenseID string `json:"sense_id"`
}

type Sense struct {
	ID             string
	Definitions    []string
	Domains        []string
	Examples       []textEntry
	Notes          []note
	Registers      []string
	SubSenses      []Sense
	ThesaurusLinks []ThesaurusLink
	Synonyms       []textEntry
	Antonyms       []textEntry
}

type Entry struct {
	GrammaticalFeatures []note
	Senses              []Sense
}

type Pronunciation struct {
	AudioFile        string
	PhoneticSpelling string
}

type LexicalEntry struct {
	Entries         []Entry
	LexicalCategory string
	Pronunciations  []Pronunciation
	Text            string
	ID              string
}

type QueryResult struct {
	ID             string
	LexicalEntries []LexicalEntry
}

type QueryReply struct {
	Error   string        `json:",omitifempty"`
	Results []QueryResult `json:",omitifempty"`
}
