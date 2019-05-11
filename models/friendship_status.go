package models

type FriendshipStatus struct {
	Following       bool `json:"following"`
	FollowedBy      bool `json:"followed_by"`
	IncomingRequest bool `json:"incoming_request"`
	OutgoingRequest bool `json:"outgoing_request"`
	IsPrivate       bool `json:"is_private"`
	IsBlockingReel  bool `json:"is_blocking_reel"`
	IsMutingReel    bool `json:"is_muting_reel"`
	Blocking        bool `json:"blocking"`
	Muting          bool `json:"muting"`
	IsBestie        bool `json:"is_bestie"`
}
