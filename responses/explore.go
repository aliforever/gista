package responses

import "github.com/aliforever/gista/models"

type Explore struct {
	Response
	NumResults          int                  `json:"num_results"`
	AutoLoadMoreEnabled interface{}          `json:"auto_load_more_enabled"`
	Items               []models.ExploreItem `json:"items"`
	MoreAvailable       interface{}          `json:"more_available"`
	NextMaxId           interface{}          `json:"next_max_id"`
	MaxId               string               `json:"max_id"`
	RankToken           string               `json:"rank_token"`
}
