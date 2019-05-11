package models

type BiographyEntities struct {
	Entities interface{} `json:"entities"`
	RawText  string      `json:"raw_text"`
	NuxType  string      `json:"nux_type"`
}
