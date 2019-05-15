package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type hashtag struct {
	ig *instagram
}

func newHashtag(i *instagram) *hashtag {
	return &hashtag{ig: i}
}

func (h *hashtag) GetBlockedMedia() (res *responses.BlockedMedia, err error) {
	res = &responses.BlockedMedia{}
	err = h.ig.client.Request(constants.BlockedMedia).GetResponse(res)
	return
}
