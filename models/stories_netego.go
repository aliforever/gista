package models

type StoriesNetego struct {
	TrackingToken  string `json:"tracking_token"`
	HideUnitIfSeen string `json:"hide_unit_if_seen"`
	Id             int    `json:"id"`
}
