package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type direct struct {
	ig *Instagram
}

func newDirect(i *Instagram) *direct {
	return &direct{ig: i}
}

func (p *direct) GetPresences() (res *responses.Presences, err error) {
	res = &responses.Presences{}
	err = p.ig.client.Request(constants.Presence).GetResponse(res)
	return
}

func (p *direct) GetInbox(cursorId *string) (res *responses.DirectInbox, err error) {
	res = &responses.DirectInbox{}
	req := p.ig.client.Request(constants.DirectInbox).AddParam("persistentBadging", "true").AddParam("use_unified_inbox", "true")
	if cursorId != nil {
		req.AddParam("cursor", *cursorId)
	}
	err = req.GetResponse(res)
	return
}

func (p *direct) GetRankedRecipients(mode string, showThreads bool, query *string) (res *responses.DirectRankedRecipients, err error) {
	res = &responses.DirectRankedRecipients{}
	showThreadsStr := "false"
	if showThreads {
		showThreadsStr = "true"
	}
	req := p.ig.client.Request(constants.DirectRankedRecipients).AddParam("mode", mode).AddParam("show_threads", showThreadsStr).AddParam("use_unified_inbox", "true")
	if query != nil {
		req.AddParam("query", *query)
	}
	err = req.GetResponse(res)
	return
}
