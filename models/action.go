package models

type Action struct {
	Title            Text   `json:"title"`
	Url              string `json:"url"`
	Limit            int    `json:"limit"`
	DismissPromotion bool   `json:"dismiss_promotion"`
}
