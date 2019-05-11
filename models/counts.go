package models

type Counts struct {
	Relationships        interface{} `json:"relationships"`
	Requests             interface{} `json:"requests"`
	PhotosOfYou          interface{} `json:"photos_of_you"`
	Usertags             interface{} `json:"usertags"`
	Comments             interface{} `json:"comments"`
	Likes                interface{} `json:"likes"`
	CommentLikes         interface{} `json:"comment_likes"`
	CampaignNotification interface{} `json:"campaign_notification"`
}
