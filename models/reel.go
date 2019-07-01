package models

type Reel struct {
	Id string `json:"id,string"`
	/*
	 * Unix "taken_at" timestamp of the newest item in their story reel.
	 */
	LatestReelMedia int64 `json:"latest_reel_media"`
	/*
	 * The "taken_at" timestamp of the last story media you have seen for
	 * that user (the current reels user). Defaults to `0` (not seen).
	 */
	Seen               int64       `json:"seen"`
	CanReply           bool        `json:"can_reply"`
	CanReshare         bool        `json:"can_reshare"`
	ReelType           string      `json:"reel_type"`
	CoverMedia         *CoverMedia `json:"cover_media,omitempty"`
	User               User        `json:"user"`
	Items              *[]Item     `json:"items,omitempty"`
	RankedPosition     string      `json:"ranked_position"`
	Title              string      `json:"title"`
	SeenRankedPosition string      `json:"seen_ranked_position"`
	ExpiringAt         int64       `json:"expiring_at"`
	HasBestiesMedia    bool        `json:"has_besties_media"` // Uses int(0) for false and 1 for true.
	Location           *Location   `json:"location,omitempty"`
	PrefetchCount      int         `json:"prefetch_count"`
	Broadcast          *Broadcast  `json:"broadcast,omitempty"`
}
