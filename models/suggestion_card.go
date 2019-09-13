package models

type SuggestionCard struct {
	UserCard      UserCard    `json:"user_card"`
	UpsellCiCard  interface{} `json:"upsell_ci_card"`
	UpsellFbcCard interface{} `json:"upsell_fbc_card"`
}
