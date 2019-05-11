package models

type DirectLink struct {
	Text        string      `json:"text"`
	LinkContext LinkContext `json:"link_context"`
}
