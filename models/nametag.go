package models

type Nametag struct {
	Mode          int    `json:"mode"`
	Gradient      string `json:"gradient"`
	Emoji         string `json:"emoji"`
	SelfieSticker int    `json:"selfie_sticker"`
}
