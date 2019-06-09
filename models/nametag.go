package models

type Nametag struct {
	Mode          int         `json:"mode"`
	Gradient      interface{} `json:"gradient"`
	Emoji         interface{} `json:"emoji"`
	SelfieSticker interface{} `json:"selfie_sticker"`
}
