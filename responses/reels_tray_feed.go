package responses

import (
	"github.com/aliforever/gista/models"
)

type ReelsTrayFeed struct {
	Response
	StoryRankingToken                string                   `json:"story_ranking_token"`
	Broadcasts                       []models.Broadcast       `json:"broadcasts"`
	Tray                             []models.StoryTray       `json:"tray"`
	PostLive                         models.PostLive          `json:"post_live"`
	StickerVersion                   int                      `json:"sticker_version"`
	FaceFilterNuxVersion             int                      `json:"face_filter_nux_version"`
	StoriesViewerGesturesNuxEligible bool                     `json:"stories_viewer_gestures_nux_eligible"`
	HasNewNuxStory                   bool                     `json:"has_new_nux_story"`
	Suggestions                      []models.TraySuggestions `json:"suggestions"`
}
