package models

type MediaShare struct {
	Media Item   `json:"media"`
	Text  string `json:"text"`
}
