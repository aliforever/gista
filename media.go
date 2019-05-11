package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type media struct {
	ig *instagram
}

func newMedia(i *instagram) *media {
	return &media{ig: i}
}

func (m *media) GetBlockedMedia() (res *responses.BlockedMedia, err error) {
	res = &responses.BlockedMedia{}
	err = m.ig.client.Request(constants.BlockedMedia).GetResponse(res)
	return
}
