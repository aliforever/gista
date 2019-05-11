package models

type Suggested struct {
	Position   int          `json:"position"`
	Hashtag    Hashtag      `json:"hashtag"`
	User       User         `json:"user"`
	Place      LocationItem `json:"place"`
	ClientTime interface{}  `json:"client_time"`
}
