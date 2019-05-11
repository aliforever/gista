package responses

import "github.com/aliforever/gista/models"

type TimelineFeed struct {
	Response
	NumResults                  int               `json:"num_results"`
	ClientGapEnforcerMatrix     interface{}       `json:"client_gap_enforcer_matrix"`
	IsDirectV2Enabled           bool              `json:"is_direct_v_2_enabled"`
	AutoLoadMoreEnabled         bool              `json:"auto_load_more_enabled"`
	MoreAvailable               bool              `json:"more_available"`
	NextMaxId                   string            `json:"next_max_id"`
	PaginationInfo              interface{}       `json:"pagination_info"`
	FeedItems                   []models.FeedItem `json:"feed_items"`
	Megaphone                   models.FeedAysf   `json:"megaphone"`
	ClientFeedChangelistApplied bool              `json:"client_feed_changelist_applied"`
	ViewStateVersion            string            `json:"view_state_version"`
	FeedPillText                string            `json:"feed_pill_text"`
	ClientSessionId             string            `json:"client_session_id"`
}
