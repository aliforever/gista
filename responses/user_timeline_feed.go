package responses

import "github.com/aliforever/gista/models"

type UserTimelineFeed struct {
	Response
	Items               *[]models.Item `json:"items,omitempty"`
	NumResults          int            `json:"num_results"`
	MoreAvailable       bool           `json:"more_available"`
	NextMaxId           string         `json:"next_max_id"`
	MaxId               string         `json:"max_id"`
	AutoLoadMoreEnabled bool           `json:"auto_load_more_enabled"`
}
