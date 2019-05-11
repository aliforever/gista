package models

type StoryHashtag struct {
	// Sticker Missing
	Hashtag     Hashtag `json:"hashtag"`
	Attribution string  `json:"attribution"`
	CustomTitle string  `json:"custom_title"`
	IsHidden    int     `json:"is_hidden"`
}
