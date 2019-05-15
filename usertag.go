package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type usertag struct {
	ig *instagram
}

func newUsertag(i *instagram) *usertag {
	return &usertag{ig: i}
}

func (ut *usertag) GetBlockedMedia() (res *responses.BlockedMedia, err error) {
	res = &responses.BlockedMedia{}
	err = ut.ig.client.Request(constants.BlockedMedia).GetResponse(res)
	return
}
