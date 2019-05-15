package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type collection struct {
	ig *instagram
}

func newCollection(i *instagram) *collection {
	return &collection{ig: i}
}

func (c *collection) GetBlockedMedia() (res *responses.BlockedMedia, err error) {
	res = &responses.BlockedMedia{}
	err = c.ig.client.Request(constants.BlockedMedia).GetResponse(res)
	return
}
