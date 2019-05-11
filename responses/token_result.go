package responses

import "github.com/aliforever/gista/models"

type TokenResult struct {
	Response
	Token models.Token
}
