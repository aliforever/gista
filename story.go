package gista

import (
	"encoding/json"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type story struct {
	ig *Instagram
}

func newStory(i *Instagram) *story {
	return &story{ig: i}
}

func (s *story) GetReelsTrayFeed() (res *responses.ReelsTrayFeed, err error) {
	res = &responses.ReelsTrayFeed{}
	cps, _ := json.Marshal(constants.SupportedCapabilities)
	err = s.ig.client.Request(constants.ReelsTrayFeed).
		SetSignedPost(false).
		AddPost("supported_capabilities_new", string(cps)).
		AddPost("reason", "pull_to_refresh").
		AddCSRFPost().
		AddUuIdPost().
		GetResponse(res)
	return
}
