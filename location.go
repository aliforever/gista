package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type location struct {
	ig *Instagram
}

func newLocation(i *Instagram) *location {
	return &location{ig: i}
}

func (l *location) GetBlockedMedia() (res *responses.BlockedMedia, err error) {
	res = &responses.BlockedMedia{}
	err = l.ig.client.Request(constants.BlockedMedia).GetResponse(res)
	return
}
