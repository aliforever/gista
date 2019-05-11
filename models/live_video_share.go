package models

type LiveVideoShare struct {
	Text        string    `json:"text"`
	Broadcast   Broadcast `json:"broadcast"`
	VideoOffset int       `json:"video_offset"`
}
