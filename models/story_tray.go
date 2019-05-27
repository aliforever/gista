package models

type StoryTray struct {
	Id                 string      `json:"id"`
	Items              []Item      `json:"items"`
	User               User        `json:"user"`
	CanReply           interface{} `json:"can_reply"`
	ExpiringAt         int64       `json:"expiring_at"`
	SeenRankedPosition int         `json:"seen_ranked_position"`
	/*
	 * The "taken_at" timestamp of the last story media you have seen for
	 * that user (the current trays user). Defaults to `0` (not seen).
	 */
	Seen int64 `json:"seen"`
	/*
	 * Unix "taken_at" timestamp of the newest item in their story reel.
	 */
	LatestReelMedia     int64       `json:"latest_reel_media"`
	RankedPosition      int         `json:"ranked_position"`
	IsNux               interface{} `json:"is_nux"`
	ShowNuxTooltip      interface{} `json:"show_nux_tooltip"`
	Muted               interface{} `json:"muted"`
	PrefetchCount       int         `json:"prefetch_count"`
	Location            Location    `json:"location"`
	SourceToken         interface{} `json:"source_token"`
	Owner               Owner       `json:"owner"`
	NuxId               string      `json:"nux_id"`
	DismissCard         DismissCard `json:"dismiss_card"`
	CanReshare          interface{} `json:"can_reshare"`
	HasBestiesMedia     bool        `json:"has_besties_media"`
	ReelType            string      `json:"reel_type"`
	UniqueIntegerReelId string      `json:"unique_integer_reel_id"`
	CoverMedia          CoverMedia  `json:"cover_media"`
	Title               string      `json:"title"`
	MediaCount          int         `json:"media_count"`
}
