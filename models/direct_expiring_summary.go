package models

type DirectExpiringSummary struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	Count     int    `json:"count"`
}
