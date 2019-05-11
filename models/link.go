package models

type Link struct {
	Start       int         `json:"start"`
	End         int         `json:"end"`
	Id          string      `json:"id"`
	Type        string      `json:"type"`
	Text        string      `json:"text"`
	LinkContext LinkContext `json:"link_context"`
}
