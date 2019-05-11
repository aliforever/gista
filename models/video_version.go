package models

type VideoVersion struct {
	Type   int    `json:"type"` // Some kinda internal type ID, such as int(102).
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Url    string `json:"url"`
	Id     string `json:"id"`
}
