package responses

import "github.com/aliforever/gista/models"

type SuggestedSearches struct {
	Response
	Suggested []models.Suggested `json:"suggested"`
	RankToken string             `json:"rank_token"`
}
