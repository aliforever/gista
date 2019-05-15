package responses

import "github.com/aliforever/gista/models"

type UserInfo struct {
	Response
	Megaphone *interface{} `json:"megaphone,omitempty"`
	User      *models.User `json:"user,omitempty"`
}
