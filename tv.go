package gista

import (
	"regexp"

	"github.com/aliforever/gista/constants"

	"github.com/aliforever/gista/errs"

	"github.com/aliforever/gista/responses"
)

type tv struct {
	ig *Instagram
}

func newTv(i *Instagram) *tv {
	return &tv{ig: i}
}

/**
 * Get channel.
 *
 * You can filter the channel with different IDs: "for_you", "chrono_following", "popular", "continue_watching"
 * and using a user ID in the following format: "user_1234567891".
 *
 */
func (tv *tv) GetChannel(id string, maxId *string) (res *responses.TVChannels, err error) {
	res = &responses.TVChannels{}
	found := false
	allowed := []string{"for_you", "chrono_following", "popular", "continue_watching"}
	for _, item := range allowed {
		if item == id {
			found = true
			break
		}
	}
	if !found {
		r := regexp.MustCompile(`^user_[1-9]\d*$`)
		if r.MatchString(id) {
			found = true
		}
	}
	if !found {
		err = errs.InvalidIdForIgtv(id)
		return
	}
	req := tv.ig.client.Request(constants.GetIGTVChannel).
		AddPost("id", id).
		AddUIdPost().
		AddUuIdPost().
		AddCSRFPost()
	if maxId != nil {
		req.AddPost("max_id", *maxId)
	}
	err = req.GetResponse(res)
	return
}
