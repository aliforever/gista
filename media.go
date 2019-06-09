package gista

import (
	"fmt"
	"strconv"

	"strings"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type media struct {
	ig *Instagram
}

func newMedia(i *Instagram) *media {
	return &media{ig: i}
}

func (m *media) GetInfo(mediaId interface{}) (res *responses.MediaInfo, err error) {
	mediaIdInt := int64(0)
	switch mediaId.(type) {
	case int64:
		mediaIdInt = mediaId.(int64)
	case string:
		idTemp, _ := strconv.Atoi(mediaId.(string)[:strings.Index(mediaId.(string), "_")])
		mediaIdInt = int64(idTemp)
	}

	res = &responses.MediaInfo{}
	err = m.ig.client.Request(fmt.Sprintf(constants.GetMediaInfo, mediaIdInt)).GetResponse(res)
	return
}

func (m *media) GetBlockedMedia() (res *responses.BlockedMedia, err error) {
	res = &responses.BlockedMedia{}
	err = m.ig.client.Request(constants.BlockedMedia).GetResponse(res)
	return
}
