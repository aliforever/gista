package models

type DirectReactions struct {
	LikesCount int              `json:"likes_count"`
	Likes      []DirectReaction `json:"likes"`
}
