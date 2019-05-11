package models

type Stories struct {
	IsPortrait interface{} `json:"is_portrait"`
	Tray       []StoryTray `json:"tray"`
	Id         string      `json:"id"`
	TopLive    TopLive     `json:"top_live"`
}
