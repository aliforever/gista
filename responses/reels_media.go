package responses

import (
	"github.com/aliforever/gista/models"
)

type ReelsMedia struct {
	Response
	ReelsMedia []models.Reel `json:"reels_media"`
	Reels      interface{}   `json:"reels"`
}
