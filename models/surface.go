package models

type Surface struct {
	Scores    interface{} `json:"scores"`
	RankToken string      `json:"rank_token"`
	TtlSecs   int         `json:"ttl_secs"`
	Name      string      `json:"name"`
}
