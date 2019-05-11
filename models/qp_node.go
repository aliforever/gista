package models

type QpNode struct {
	Id                string           `json:"id"`
	PromotionId       string           `json:"promotion_id"`
	MaxImpressions    int              `json:"max_impressions"`
	Triggers          []string         `json:"triggers"`
	ContextualFilters ContextualFilter `json:"contextual_filters"`
	Template          Template         `json:"template"`
	Creatives         []Creative       `json:"creatives"`
}
