package osrs_api

type RarityNote struct {
	Text      string
	Reference *Reference
}

type Reference struct {
	Title  string
	Name   string
	Url    string
	Author string
	Date   string
}

func NewNote(str string) *RarityNote {
	n := RarityNote{}

	return &n
}
