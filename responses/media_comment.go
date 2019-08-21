package responses

import "github.com/aliforever/gista/models"

type MediaComment struct {
	Response
	Comments                []models.Comment `json:"comments"`
	CommentCount            int              `json:"comment_count"`
	CommentLikesEnabled     bool             `json:"comment_likes_enabled"`
	NextMaxId               string           `json:"next_max_id"`
	NextMinId               string           `json:"next_min_id"`
	Caption                 models.Caption   `json:"caption"`
	HasMoreComments         bool             `json:"has_more_comments"`
	CaptionIsEdited         bool             `json:"caption_is_edited"`
	PreviewComments         interface{}      `json:"preview_comments"`
	HasMoreHeadloadComments bool             `json:"has_more_headload_comments"`
	MediaHeaderDisplay      string           `json:"media_header_display"`
	ThreadingEnabled        bool             `json:"threading_enabled"`
}
