package models

type Story struct {
	Pk        string `json:"pk"`
	Counts    Counts `json:"counts"`
	Args      Args   `json:"args"`
	Type      int    `json:"type"`
	StoryType int    `json:"story_type"`
}
