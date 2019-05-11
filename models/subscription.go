package models

type Subscription struct {
	Topic    interface{} `json:"topic"`
	Url      string      `json:"url"`
	Sequence interface{} `json:"sequence"`
	Auth     interface{} `json:"auth"`
}
