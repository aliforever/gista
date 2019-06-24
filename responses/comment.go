package responses

import "github.com/aliforever/gista/models"

type Comment struct {
	Response
	Comment *models.Comment `json:"comment"`
}
