package models

type AymfItem struct {
	Algorithm     string      `json:"algorithm"`
	SocialContext string      `json:"social_context"`
	Caption       string      `json:"caption"`
	Uuid          string      `json:"uuid"`
	MediaIds      []int64     `json:"media_ids"`
	ThumbnailUrls []string    `json:"thumbnail_urls"`
	LargeUrls     []string    `json:"large_urls"`
	MediaInfos    interface{} `json:"media_infos"`
	FollowedBy    bool        `json:"followed_by"`
	Value         float64     `json:"value"`
}
