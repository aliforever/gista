package gista

import (
	"encoding/json"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type people struct {
	Ig *instagram
}

func NewPeople(i *instagram) *people {
	return &people{Ig: i}
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
	err = p.Ig.Client.Request(constants.BootstrapUsers).AddParam("surfaces", string(j)).GetResponse(res)
	return
}

func (p *people) GetRecentActivityInbox() (res *responses.ActivityNews, err error) {
	res = &responses.ActivityNews{}
	err = p.Ig.Client.Request(constants.ActivityNews).GetResponse(res)
	return
}
