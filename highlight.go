package gista

import (
	"encoding/json"
	"fmt"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type highlight struct {
	ig *Instagram
}

func newHighlight(i *Instagram) *highlight {
	return &highlight{ig: i}
}

func (h *highlight) GetUserFeed(userId int64) (res *responses.HighlightFeed, err error) {
	res = &responses.HighlightFeed{}
	j, _ := json.Marshal(constants.SupportedCapabilities)
	err = h.ig.client.Request(fmt.Sprintf(constants.GetHighlightsUserFeed, userId)).
		AddParam("supported_capabilities_new", string(j)).
		AddPhoneIdParam().
		AddParam("battery_level", "100").
		AddParam("is_charging", "1").
		AddParam("will_sound_on", "1").
		GetResponse(res)
	return
}

func (h *highlight) GetBlockedMedia() (res *responses.BlockedMedia, err error) {
	res = &responses.BlockedMedia{}
	err = h.ig.client.Request(constants.BlockedMedia).GetResponse(res)
	return
}
