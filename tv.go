package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type tv struct {
	ig *Instagram
}

func newTv(i *Instagram) *tv {
	return &tv{ig: i}
}

func (t *tv) GetBlockedMedia() (res *responses.BlockedMedia, err error) {
	res = &responses.BlockedMedia{}
	err = t.ig.client.Request(constants.BlockedMedia).GetResponse(res)
	return
}
