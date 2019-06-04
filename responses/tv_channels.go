package responses

import "github.com/aliforever/gista/models"

type TVChannels struct {
	Response
	Type          string         `json:"type"`
	Title         string         `json:"title"`
	Id            string         `json:"id"`
	Items         *[]models.Item `json:"items,omitempty"`
	MoreAvailable bool           `json:"more_available"`
	MaxId         string         `json:"max_id"`
	SeenState     interface{}    `json:"seen_state"`
	UserDict      models.User    `json:"user_dict"`
}
