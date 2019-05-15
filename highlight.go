package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type highlight struct {
	ig *instagram
}

func newHighlight(i *instagram) *highlight {
	return &highlight{ig: i}
}

func (h *highlight) GetBlockedMedia() (res *responses.BlockedMedia, err error) {
	res = &responses.BlockedMedia{}
	err = h.ig.client.Request(constants.BlockedMedia).GetResponse(res)
	return
}
