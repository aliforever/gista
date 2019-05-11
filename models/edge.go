package models

type Edge struct {
	Priority  int       `json:"priority"`
	TimeRange TimeRange `json:"time_range"`
	Node      QpNode    `json:"node"`
}
