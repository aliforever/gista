package models

type DirectThreadItem struct {
	ItemId                     string                `json:"item_id"`
	ItemType                   interface{}           `json:"item_type"`
	Text                       string                `json:"text"`
	MediaShare                 Item                  `json:"media_share"`
	PreviewMedias              []Item                `json:"preview_medias"`
	Media                      DirectThreadItemMedia `json:"media"`
	UserId                     string                `json:"user_id"`
	Timestamp                  float64               `json:"timestamp"`
	ClientContext              string                `json:"client_context"`
	HideInThread               int                   `json:"hide_in_thread"`
	ActionLog                  ActionLog             `json:"action_log"`
	Link                       DirectLink            `json:"link"`
	Reactions                  DirectReactions       `json:"reactions"`
	RavenMedia                 Item                  `json:"raven_media"`
	SeenUserIds                []string              `json:"seen_user_ids"`
	ExpiringMediaActionSummary DirectExpiringSummary `json:"expiring_media_action_summary"`
	ReelShare                  ReelShare             `json:"reel_share"`
	Placeholder                Placeholder           `json:"placeholder"`
	Location                   Location              `json:"location"`
	Like                       interface{}           `json:"like"`
	LiveVideoShare             LiveVideoShare        `json:"live_video_share"`
	LiveViewerInvite           LiveViewerInvite      `json:"live_viewer_invite"`
	Profile                    User                  `json:"profile"`
	StoryShare                 StoryShare            `json:"story_share"`
	DirectMediaShare           MediaShare            `json:"direct_media_share"`
	VideoCallEvent             VideoCallEvent        `json:"video_call_event"`
	ProductShare               ProductShare          `json:"product_share"`
	AnimatedMedia              AnimatedMedia         `json:"animated_media"`
	FelixShare                 FelixShare            `json:"felix_share"`
}
