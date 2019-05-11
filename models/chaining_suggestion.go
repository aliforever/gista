package models

type ChainingSuggestion struct {
	ChainingInfo                  ChainingInfo `json:"chaining_info"`
	ProfileChainingSecondaryLabel interface{}  `json:"profile_chaining_secondary_label"`
}
