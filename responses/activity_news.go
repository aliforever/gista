package responses

import "github.com/aliforever/gista/models"

type ActivityNews struct {
	Response
	NewStories           []models.Story      `json:"new_stories"`
	OldStories           []models.Story      `json:"old_stories"`
	Continuation         interface{}         `json:"continuation"`
	FriendRequestStories []models.Story      `json:"friend_request_stories"`
	Counts               models.Counts       `json:"counts"`
	Subscription         models.Subscription `json:"subscription"`
	Partition            interface{}         `json:"partition"`
	ContinuationToken    interface{}         `json:"continuation_token"`
	AdsManager           interface{}         `json:"ads_manager"`
	Aymf                 models.Aymf         `json:"aymf"`
}
