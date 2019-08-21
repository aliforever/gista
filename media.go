package gista

import (
	"fmt"
	"strconv"

	"github.com/aliforever/gista/errs"

	"github.com/aliforever/gista/signatures"
	"github.com/aliforever/gista/utils"

	"strings"

	"errors"
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

func (m *media) LikeComment(commentId int64) (res *responses.CommentLike, err error) {
	res = &responses.CommentLike{}
	err = m.ig.client.Request(fmt.Sprintf(constants.CommentLike, commentId)).
		AddUuIdPost().
		AddUIdPost().
		AddCSRFPost().
		AddDeviceIdPost().GetResponse(res)
	return
}

func (m *media) Comment(mediaId interface{}, commentText string, replyCommentId *int, module *string) (res *responses.Comment, err error) {
	res = &responses.Comment{}
	mediaIdInt := int64(0)
	switch mediaId.(type) {
	case int64:
		mediaIdInt = mediaId.(int64)
	case string:
		idTemp, _ := strconv.Atoi(mediaId.(string)[:strings.Index(mediaId.(string), "_")])
		mediaIdInt = int64(idTemp)
	}
	moduleText := ""
	if module == nil {
		moduleText = "comments_feed_timeline"
	}

	request := m.ig.client.Request(fmt.Sprintf(constants.CommentMedia, mediaIdInt)).
		AddPost("user_breadcrumb", utils.GenerateUserBreadCrumb(len([]rune(commentText)))).
		AddPost("idempotence_token", signatures.GenerateUUID(true)).
		AddUuIdPost().
		AddUIdPost().
		AddCSRFPost().
		AddPost("comment_text", commentText).
		AddPost("containermodule", moduleText).
		AddPost("radio_type", "wifi-none").
		AddDeviceIdPost()
		/*        if ($replyCommentId !== null) {
		          if ($commentText[0] !== '@') {
		              throw new \InvalidArgumentException('When replying to a comment, your text must begin with an @-mention to their username.');
		          }
		          $request->addPost('replied_to_comment_id', $replyCommentId);
		      }*/
	if replyCommentId != nil {
		if commentText[0] != '@' {
			err = errs.MissingMentionInReply(commentText)
			return
		}
		request.AddPost("replied_to_comment_id", fmt.Sprintf("%d", replyCommentId))
	}
	err = request.GetResponse(res)
	return
}

func (m *media) GetComments(mediaId interface{}, options map[string]string) (res *responses.MediaComment, err error) {
	res = &responses.MediaComment{}
	mediaIdInt := int64(0)
	switch mediaId.(type) {
	case int64:
		mediaIdInt = mediaId.(int64)
	case string:
		idTemp, _ := strconv.Atoi(mediaId.(string)[:strings.Index(mediaId.(string), "_")])
		mediaIdInt = int64(idTemp)
	}
	req := m.ig.client.Request(fmt.Sprintf(constants.GetComments, mediaIdInt)).
		AddParam("can_support_threading", "true")
	if options != nil {
		minId, minIdOk := options["min_id"]
		maxId, maxIdOk := options["max_id"]
		if minIdOk && maxIdOk {
			err = errors.New("you can use either 'min_id' or 'max_id', but not both at the same time")
			return
		}
		if minIdOk {
			req.AddParam("min_id", minId)
		}
		if maxIdOk {
			req.AddParam("max_id", maxId)
		}
		if val, ok := options["target_comment_id"]; ok {
			if minIdOk || maxIdOk {
				err = errors.New("you cannot use the 'target_comment_id' parameter together with the 'min_id' or 'max_id' parameters")
				return
			}
			req.AddParam("target_comment_id", val)
		}
	}
	err = req.GetResponse(res)
	return
}
