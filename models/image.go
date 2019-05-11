package models

type Image struct {
	Uri    string `json:"uri"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
