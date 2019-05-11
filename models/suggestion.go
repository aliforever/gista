package models

type Suggestion struct {
	MediaInfos      interface{} `json:"media_infos"`
	SocialContext   string      `json:"social_context"`
	Algorithm       string      `json:"algorithm"`
	ThumbnailUrls   []string    `json:"thumbnail_urls"`
	Value           float64     `json:"value"`
	Caption         interface{} `json:"caption"`
	User            User        `json:"user"`
	LargeUrls       []string    `json:"large_urls"`
	MediaIds        interface{} `json:"media_ids"`
	Icon            interface{} `json:"icon"`
	IsNewSuggestion bool        `json:"is_new_suggestion"`
	Uuid            string      `json:"uuid"`
}
