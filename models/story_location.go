package models

type StoryLocation struct {
	// Sticker Missing
	Location    Location `json:"location"`
	Attribution string   `json:"attribution"`
	IsHidden    int      `json:"is_hidden"`
}
