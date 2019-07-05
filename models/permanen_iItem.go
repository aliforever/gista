package models

type PermanentItem struct {
	ItemId         string         `json:"item_id"`
	UserId         int64          `json:"user_id"`
	Timestamp      int64          `json:"timestamp"`
	ItemType       string         `json:"item_type"`
	Profile        User           `json:"profile"`
	Text           string         `json:"text"`
	Location       Location       `json:"location"`
	Like           interface{}    `json:"like"`
	Media          MediaData      `json:"media"`
	Link           Link           `json:"link"`
	MediaShare     Item           `json:"media_share"`
	ReelShare      ReelShare      `json:"reel_share"`
	ClientContext  string         `json:"client_context"`
	LiveVideoShare LiveVideoShare `json:"live_video_share"`
	HideInThread   int            `json:"hide_in_thread"`
}
