package responses

import "github.com/aliforever/gista/models"

type UserStoryFeed struct {
	Response
	Broadcast    *models.Broadcast    `json:"broadcast,omitempty"`
	Reel         *models.Reel         `json:"reel,omitempty"`
	PostLiveItem *models.PostLiveItem `json:"post_live_item,omitempty"`
}
