package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type live struct {
	ig *instagram
}

func newLive(i *instagram) *live {
	return &live{ig: i}
}

func (l *live) GetBlockedMedia() (res *responses.BlockedMedia, err error) {
	res = &responses.BlockedMedia{}
	err = l.ig.client.Request(constants.BlockedMedia).GetResponse(res)
	return
}
