package models

type VideoCallEvent struct {
	Action string `json:"action"`
	VcId   string `json:"vc_id"`
}
