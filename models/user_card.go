package models

type UserCard struct {
	User            User        `json:"user"`
	Algorithm       string      `json:"algorithm"`
	SocialContext   string      `json:"social_context"`
	Caption         interface{} `json:"caption"`
	Icon            interface{} `json:"icon"`
	MediaIds        interface{} `json:"media_ids"`
	ThumbnailUrls   interface{} `json:"thumbnail_urls"`
	LargeUrls       interface{} `json:"large_urls"`
	MediaInfos      interface{} `json:"media_infos"`
	Value           float64     `json:"value"`
	IsNewSuggestion bool        `json:"is_new_suggestion"`
	Uuid            string      `json:"uuid"`
}
