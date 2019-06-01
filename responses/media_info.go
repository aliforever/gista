package responses

import "github.com/aliforever/gista/models"

type MediaInfo struct {
	Response
	AutoLoadMoreEnabled bool          `json:"auto_load_more_enabled"`
	NumResults          int           `json:"num_results"`
	MoreAvailable       bool          `json:"more_available"`
	Items               []models.Item `json:"items"`
}
