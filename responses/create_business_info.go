package responses

import "github.com/aliforever/gista/models"

type CreateBusinessInfo struct {
	Response
	Users []models.User `json:"users,omitempty"`
}
