package responses

import "github.com/aliforever/gista/models"

type BootstrapUsers struct {
	Response
	Surfaces []models.Surface `json:"surfaces"`
	Users    []models.User    `json:"users"`
}
