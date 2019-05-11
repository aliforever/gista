package models

type ActionBadge struct {
	ActionType      interface{} `json:"action_type"`
	ActionCount     interface{} `json:"action_count"`
	ActionTimestamp interface{} `json:"action_timestamp"`
}
