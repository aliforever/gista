package gista

import (
	"encoding/json"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type story struct {
	Ig *instagram
}

func NewStory(i *instagram) *story {
	return &story{Ig: i}
}

func (s *story) GetReelsTrayFeed() (res *responses.ReelsTrayFeed, err error) {
	res = &responses.ReelsTrayFeed{}
	cps, _ := json.Marshal(constants.SupportedCapabilities)
	err = s.Ig.Client.Request(constants.ReelsTrayFeed).
		SetSignedPost(false).
		AddPost("supported_capabilities_new", string(cps)).
		AddPost("reason", "pull_to_refresh").
		AddCSRFPost().
		AddUuIdPost().
		GetResponse(res)
	return
}
