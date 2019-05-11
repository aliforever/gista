package models

type Media struct {
	Image                   string      `json:"image"`
	Id                      string      `json:"id"`
	User                    User        `json:"user"`
	ExpiringAt              interface{} `json:"expiring_at"`
	CommentThreadingEnabled bool        `json:"comment_threading_enabled"`
}
