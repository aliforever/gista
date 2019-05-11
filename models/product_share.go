package models

type ProductShare struct {
	Media   Item    `json:"media"`
	Text    string  `json:"text"`
	Product Product `json:"product"`
}
