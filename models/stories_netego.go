package models

type StoriesNetego struct {
	TrackingToken  string `json:"tracking_token"`
	HideUnitIfSeen bool   `json:"hide_unit_if_seen"`
	Id             int    `json:"id"`
}
