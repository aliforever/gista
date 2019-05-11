package models

type StoryQuestion struct {
	X               float64         `json:"x"`
	Y               float64         `json:"y"`
	Z               float64         `json:"z"`
	Width           float64         `json:"width"`
	Height          float64         `json:"height"`
	Rotation        float64         `json:"rotation"`
	IsPinned        int             `json:"is_pinned"`
	IsHidden        int             `json:"is_hidden"`
	QuestionSticker QuestionSticker `json:"question_sticker"`
}
