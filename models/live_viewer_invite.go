package models

type LiveViewerInvite struct {
	Text      string    `json:"text"`
	Broadcast Broadcast `json:"broadcast"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
}
