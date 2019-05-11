package responses

import "github.com/aliforever/gista/models"

type RecentSearches struct {
	Response
	Recent []models.Suggested `json:"recent"`
}
