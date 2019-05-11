package models

type AnimatedMedia struct {
	Id     string             `json:"id"`
	Images AnimatedMediaImage `json:"images"`
}
