package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type media struct {
	Ig *instagram
}

func newMedia(i *instagram) *media {
	return &media{Ig: i}
}

func (m *media) GetBlockedMedia() (res *responses.BlockedMedia, err error) {
	res = &responses.BlockedMedia{}
	err = m.Ig.Client.Request(constants.BlockedMedia).GetResponse(res)
	return
}
