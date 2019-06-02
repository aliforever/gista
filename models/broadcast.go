package models

type Broadcast struct {
	BroadcastOwner User `json:"broadcast_owner"`
	/*
	 * A string such as "active" or "post_live".
	 */
	BroadcastStatus      string      `json:"broadcast_status"`
	CoverFrameUrl        string      `json:"cover_frame_url"`
	PublishedTime        int64       `json:"published_time"`
	BroadcastMessage     string      `json:"broadcast_message"`
	Muted                interface{} `json:"muted"`
	MediaId              string      `json:"media_id"`
	Id                   int64       `json:"id"`
	RtmpPlaybackUrl      string      `json:"rtmp_playback_url"`
	DashAbrPlaybackUrl   string      `json:"dash_abr_playback_url"`
	DashPlaybackUrl      string      `json:"dash_playback_url"`
	RankedPosition       interface{} `json:"ranked_position"`
	OrganicTrackingToken string      `json:"organic_tracking_token"`
	SeenRankedPosition   int         `json:"seen_ranked_position"`
	ViewerCount          float64     `json:"viewer_count"`
	DashManifest         string      `json:"dash_manifest"`
	/*
	 * Unix timestamp of when the "post_live" will expire.
	 */
	ExpireAt               int64  `json:"expire_at"`
	EncodingTag            string `json:"encoding_tag"`
	TotalUniqueViewerCount int    `json:"total_unique_viewer_count"`
	InternalOnly           bool   `json:"internal_only"`
	NumberOfQualities      int    `json:"number_of_qualities"`
}
