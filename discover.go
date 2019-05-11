package gista

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aliforever/gista/errors"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type discover struct {
	Ig *instagram
}

func NewDiscover(i *instagram) *discover {
	return &discover{Ig: i}
}

func (d *discover) GetExploreFeed(maxId *string, prefetch bool) (res *responses.Explore, err error) {
	res = &responses.Explore{}
	p := "false"
	if prefetch {
		p = "true"
	}
	_, offset := time.Now().Zone()
	j, _ := json.Marshal(constants.SupportedCapabilities)
	request := d.Ig.Client.Request(constants.Explore).
		AddParam("is_prefetch", p).
		AddParam("is_from_promote", "false").
		AddParam("timezone_offset", fmt.Sprintf("%d", offset)).
		AddParam("session_id", d.Ig.sessionId).
		AddParam("supported_capabilities_new", string(j))
	if !prefetch {
		if maxId == nil {
			z := "0"
			maxId = &z
		}
		request.AddParam("max_id", *maxId)
		request.AddParam("module", "explore_popular")
	}
	err = request.GetResponse(res)
	return
}

func (d *discover) GetRecentSearches() (res *responses.RecentSearches, err error) {
	res = &responses.RecentSearches{}
	err = d.Ig.Client.Request(constants.RecentSearches).GetResponse(res)
	return
}

func (d *discover) GetSuggestedSearches(searchType string) (res *responses.SuggestedSearches, err error) {
	found := false
	for _, t := range []string{"blended", "users", "hashtags", "places"} {
		if t == searchType {
			found = true
		}
	}
	if !found {
		err = errors.UnknownSearchType(searchType)
		return
	}
	res = &responses.SuggestedSearches{}
	err = d.Ig.Client.Request(constants.SuggestedSearches).
		AddParam("type", searchType).
		GetResponse(res)
	return
}
