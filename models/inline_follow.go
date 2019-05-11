package models

type InlineFollow struct {
	UserInfo        User `json:"user_info"`
	Following       bool `json:"following"`
	OutgoingRequest bool `json:"outgoing_request"`
}
