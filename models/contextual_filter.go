package models

type ContextualFilter struct {
	ClauseType string      `json:"clause_type"`
	Filters    interface{} `json:"filters"`
	Clauses    interface{} `json:"clauses"`
}
