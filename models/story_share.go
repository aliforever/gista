package models

type StoryShare struct {
	Media    Item   `json:"media"`
	Text     string `json:"text"`
	Title    string `json:"title"`
	Message  string `json:"message"`
	IsLinked bool   `json:"is_linked"`
}
