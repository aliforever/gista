package models

type LinkContext struct {
	LinkUrl      string `json:"link_url"`
	LinkTitle    string `json:"link_title"`
	LinkSummary  string `json:"link_summary"`
	LinkImageUrl string `json:"link_image_url"`
}
