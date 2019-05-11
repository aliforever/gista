package models

type Headline struct {
	ContentType  interface{} `json:"content_type"`
	User         User        `json:"user"`
	UserId       string      `json:"user_id"`
	Pk           string      `json:"pk"`
	Text         string      `json:"text"`
	Type         int         `json:"type"`
	CreatedAt    string      `json:"created_at"`
	CreatedAtUtc string      `json:"created_at_utc"`
	MediaId      string      `json:"media_id"`
	BitFlags     int         `json:"bit_flags"`
	Status       interface{} `json:"status"`
}
