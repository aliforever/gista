package gista

import (
	"encoding/json"
	"fmt"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type story struct {
	ig *Instagram
}

func newStory(i *Instagram) *story {
	return &story{ig: i}
}

func (s *story) GetUserReelMediaFeed(userId int64) (res *responses.UserReelMediaFeed, err error) {
	res = &responses.UserReelMediaFeed{}
	err = s.ig.client.Request(fmt.Sprintf(constants.GetUserReelMediaFeed, userId)).
		GetResponse(res)
	return
}

func (s *story) GetUserStoryFeed(userId int64) (res *responses.UserStoryFeed, err error) {
	res = &responses.UserStoryFeed{}
	c, _ := json.Marshal(constants.SupportedCapabilities)
	err = s.ig.client.Request(fmt.Sprintf(constants.GetUserStoryFeed, userId)).
		AddParam("supported_capabilities_new", string(c)).
		GetResponse(res)
	return
}

func (s *story) GetReelsTrayFeed() (res *responses.ReelsTrayFeed, err error) {
	res = &responses.ReelsTrayFeed{}
	cps, _ := json.Marshal(constants.SupportedCapabilities)
	err = s.ig.client.Request(constants.ReelsTrayFeed).
		SetSignedPost(false).
		AddPost("supported_capabilities_new", string(cps)).
		AddPost("reason", "pull_to_refresh").
		AddCSRFPost().
		AddUuIdPost().
		GetResponse(res)
	return
}

func (s *story) GetReelsMediaFeed(feedList []string, source *string) (res *responses.ReelsMedia, err error) {
	res = &responses.ReelsMedia{}
	src := "feed_timeline"
	if source != nil {
		src = *source
	}
	cps, _ := json.Marshal(constants.SupportedCapabilities)
	err = s.ig.client.Request(constants.GetReelsMediaFeed).
		AddPost("supported_capabilities_new", string(cps)).
		AddPost("user_ids", feedList).
		AddPost("source", src).
		AddCSRFPost().
		AddUIdPost().
		AddUuIdPost().
		GetResponse(res)
	return
}
