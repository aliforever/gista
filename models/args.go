package models

type Args struct {
	MediaDestination        string       `json:"media_destination"`
	Text                    string       `json:"text"`
	IconUrl                 string       `json:"icon_url"`
	Links                   []Link       `json:"links"`
	RichText                string       `json:"rich_text"`
	ProfileId               string       `json:"profile_id"`
	ProfileImage            string       `json:"profile_image"`
	Media                   []Media      `json:"media"`
	CommentNotifType        string       `json:"comment_notif_type"`
	Timestamp               string       `json:"timestamp"`
	Tuuid                   string       `json:"tuuid"`
	Clicked                 bool         `json:"clicked"`
	ProfileName             string       `json:"profile_name"`
	ActionUrl               string       `json:"action_url"`
	Destination             string       `json:"destination"`
	Actions                 []string     `json:"actions"`
	LatestReelMedia         string       `json:"latest_reel_media"`
	CommentId               string       `json:"comment_id"`
	RequestCount            interface{}  `json:"request_count"`
	InlineFollow            InlineFollow `json:"inline_follow"`
	CommentIds              []string     `json:"comment_ids"`
	SecondProfileId         string       `json:"second_profile_id"`
	SecondProfileImage      interface{}  `json:"second_profile_image"`
	ProfileImageDestination interface{}  `json:"profile_image_destination"`
}
