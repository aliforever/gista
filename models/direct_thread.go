package models

type DirectThread struct {
	ThreadId                  string             `json:"thread_id"`
	ThreadV2Id                string             `json:"thread_v_2_id"`
	Users                     []User             `json:"users"`
	LeftUsers                 []User             `json:"left_users"`
	Items                     []DirectThreadItem `json:"items"`
	LastActivityAt            int64              `json:"last_activity_at"`
	Muted                     bool               `json:"muted"`
	IsPin                     bool               `json:"is_pin"`
	Named                     bool               `json:"named"`
	Canonical                 bool               `json:"canonical"`
	Pending                   bool               `json:"pending"`
	ValuedRequest             bool               `json:"valued_request"`
	ThreadType                string             `json:"thread_type"`
	ViewerId                  int64              `json:"viewer_id"`
	ThreadTitle               string             `json:"thread_title"`
	PendingScore              int64              `json:"pending_score"`
	VcMuted                   bool               `json:"vc_muted"`
	IsGroup                   bool               `json:"is_group"`
	ReshareSendCount          int                `json:"reshare_send_count"`
	ReshareReceiveCount       int                `json:"reshare_receive_count"`
	ExpiringMediaSendCount    int                `json:"expiring_media_send_count"`
	ExpiringMediaReceiveCount int                `json:"expiring_media_receive_count"`
	Inviter                   User               `json:"inviter"`
	HasOlder                  bool               `json:"has_older"`
	HasNewer                  bool               `json:"has_newer"`
	LastSeenAt                interface{}        `json:"last_seen_at"`
	NewestCursor              string             `json:"newest_cursor"`
	OldestCursor              string             `json:"oldest_cursor"`
	IsSpam                    bool               `json:"is_spam"`
	LastPermanentItem         PermanentItem      `json:"last_permanent_item"`
	UnseenCount               int                `json:"unseen_count"`
	ActionBadge               ActionBadge        `json:"action_badge"`
	LastActivityAtSecs        interface{}        `json:"last_activity_at_secs"`
}
