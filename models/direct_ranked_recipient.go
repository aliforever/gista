package models

type DirectRankedRecipient struct {
	Thread DirectThread `json:"thread"`
	User   User         `json:"user"`
}
