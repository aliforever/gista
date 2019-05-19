package models

type Hashtag struct {
	Id               int64       `json:"id"`
	Name             string      `json:"name"`
	MediaCount       int         `json:"media_count"`
	ProfilePicUrl    string      `json:"profile_pic_url"`
	FollowStatus     int         `json:"follow_status"`
	Following        int         `json:"following"`
	AllowFollowing   int         `json:"allow_following"`
	AllowMutingStory bool        `json:"allow_muting_story"`
	RelatedTags      interface{} `json:"related_tags"`
	DebugInfo        interface{} `json:"debug_info"`
}
