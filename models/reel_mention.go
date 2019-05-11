package models

type ReelMention struct {
	User     User `json:"user"`
	IsHidden int  `json:"is_hidden"`
}
