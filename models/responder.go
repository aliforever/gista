package models

type Responder struct {
	Response          string `json:"response"`
	HasSharedResponse bool   `json:"has_shared_response"`
	Id                string `json:"id"`
	User              User   `json:"user"`
	Ts                int    `json:"ts"`
}
