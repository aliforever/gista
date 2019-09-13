package responses

import "github.com/aliforever/gista/models"

type FollowerAndFollowing struct {
	Response
	Users                         []models.User  `json:"users"`
	SuggestedUsers                SuggestedUsers `json:"suggested_users"`
	TruncateFollowRequestsAtIndex int            `json:"truncate_follow_requests_at_index"`
	NextMaxId                     string         `json:"next_max_id"`
	PageSize                      interface{}    `json:"page_size"`
	BigList                       interface{}    `json:"big_list"`
}
