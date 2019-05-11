package models

type ReelShare struct {
	Tray              []map[string]interface{} `json:"tray"` /*Recursive error on type Item*/
	StoryRankingToken string                   `json:"story_ranking_token"`
	Broadcasts        interface{}              `json:"broadcasts"`
	StickerVersion    int                      `json:"sticker_version"`
	Text              string                   `json:"text"`
	Type              string                   `json:"type"`
	IsReelPersisted   bool                     `json:"is_reel_persisted"`
	ReelOwnerId       string                   `json:"reel_owner_id"`
	ReelType          string                   `json:"reel_type"`
	Media             []map[string]interface{} `json:"media"` /*Recursive error on type Item*/
	MentionedUserId   string                   `json:"mentioned_user_id"`
}
