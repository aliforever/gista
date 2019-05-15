package responses

import "github.com/aliforever/gista/models"

type SwitchPersonalProfile struct {
	Response
	Users []models.User `json:"users,omitempty"`
}
