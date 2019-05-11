package models

type Gating struct {
	GatingType  interface{} `json:"gating_type"`
	Description interface{} `json:"description"`
	Buttons     interface{} `json:"buttons"`
	Title       interface{} `json:"title"`
}
