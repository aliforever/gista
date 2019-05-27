package responses

import "github.com/aliforever/gista/models"

type HighlightFeed struct {
	Response
	AutoLoadMoreEnabled bool                `json:"auto_load_more_enabled"`
	NextMaxId           string              `json:"next_max_id"`
	Stories             *[]models.Story     `json:"stories,omitempty"`
	ShowEmptyState      bool                `json:"show_empty_state"`
	Tray                *[]models.StoryTray `json:"tray,omitempty"`
	TvChannel           interface{}         `json:"tv_channel"`
}
