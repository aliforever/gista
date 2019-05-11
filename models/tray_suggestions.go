package models

type TraySuggestions struct {
	Tray           []StoryTray `json:"tray"`
	TrayTitle      string      `json:"tray_title"`
	BannerTitle    string      `json:"banner_title"`
	BannerSubtitle string      `json:"banner_subtitle"`
	SuggestionType string      `json:"suggestion_type"`
}
