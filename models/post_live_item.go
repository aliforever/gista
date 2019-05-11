package models

type PostLiveItem struct {
	Pk                  string      `json:"pk"`
	User                User        `json:"user"`
	Broadcasts          []Broadcast `json:"broadcasts"`
	PeakViewerCount     int         `json:"peak_viewer_count"`
	LastSeenBroadcastTs interface{} `json:"last_seen_broadcast_ts"`
	CanReply            interface{} `json:"can_reply"`
	RankedPosition      interface{} `json:"ranked_position"`
	SeenRankedPosition  interface{} `json:"seen_ranked_position"`
	Muted               interface{} `json:"muted"`
	CanReshare          interface{} `json:"can_reshare"`
}
