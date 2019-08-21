package gista

import (
	"encoding/json"
	"fmt"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type people struct {
	ig *Instagram
}

func newPeople(i *Instagram) *people {
	return &people{ig: i}
}

func (p *people) GetInfoById(userId int64, module *string) (res *responses.UserInfo, err error) {
	/*modules:
		 *							  "comment_likers",
	     *                            "comment_owner",
	     *                            "followers",
	     *                            "following",
	     *                            "likers_likers_media_view_profile",
	     *                            "likers_likers_photo_view_profile",
	     *                            "likers_likers_video_view_profile",
	     *                            "newsfeed",
	     *                            "self_followers",
	     *                            "self_following",
	     *                            "self_likers_self_likers_media_view_profile",
	     *                            "self_likers_self_likers_photo_view_profile",
	     *                            "self_likers_self_likers_video_view_profile".
	*/
	res = &responses.UserInfo{}
	req := p.ig.client.Request(fmt.Sprintf(constants.GetInfoById, userId))
	if module != nil {
		req.AddParam("from_module", *module)
	}
	err = req.GetResponse(res)
	return
}

func (p *people) GetInfoByName(username string, module *string) (res *responses.UserInfo, err error) {
	/*modules:
		 *							  "comment_likers",
	     *                            "comment_owner",
	     *                            "followers",
	     *                            "following",
	     *                            "likers_likers_media_view_profile",
	     *                            "likers_likers_photo_view_profile",
	     *                            "likers_likers_video_view_profile",
	     *                            "newsfeed",
	     *                            "self_followers",
	     *                            "self_following",
	     *                            "self_likers_self_likers_media_view_profile",
	     *                            "self_likers_self_likers_photo_view_profile",
	     *                            "self_likers_self_likers_video_view_profile".
	*/
	res = &responses.UserInfo{}
	req := p.ig.client.Request(fmt.Sprintf(constants.GetInfoByUsername, username))
	if module != nil {
		req.AddParam("from_module", *module)
	}
	err = req.GetResponse(res)
	return
}

func (p *people) GetBootstrapUsers() (res *responses.BootstrapUsers, err error) {
	res = &responses.BootstrapUsers{}
	surfaces := []string{
		"coefficient_direct_closed_friends_ranking",
		"coefficient_direct_recipients_ranking_variant_2",
		"coefficient_rank_recipient_user_suggestion",
		"coefficient_ios_section_test_bootstrap_ranking",
		"autocomplete_user_list",
	}
	j, _ := json.Marshal(surfaces)
	err = p.ig.client.Request(constants.BootstrapUsers).AddParam("surfaces", string(j)).GetResponse(res)
	return
}

func (p *people) GetRecentActivityInbox() (res *responses.ActivityNews, err error) {
	res = &responses.ActivityNews{}
	err = p.ig.client.Request(constants.ActivityNews).GetResponse(res)
	return
}

func (p *people) GetFriendship(userId int64) (res *responses.FriendshipsShow, err error) {
	res = &responses.FriendshipsShow{}
	err = p.ig.client.Request(fmt.Sprintf(constants.GetFriendship, userId)).GetResponse(res)
	return
}

func (p *people) Follow(userId int64) (res *responses.Friendship, err error) {
	res = &responses.Friendship{}
	err = p.ig.client.Request(fmt.Sprintf(constants.FollowUser, userId)).
		AddUuIdPost().
		AddUIdPost().
		AddCSRFPost().
		AddPost("user_id", userId).
		AddPost("radio_type", "wifi-none").
		AddDeviceIdPost().GetResponse(res)
	return
}
