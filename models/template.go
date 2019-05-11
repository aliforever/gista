package models

type Template struct {
	Name       string      `json:"name"`
	Parameters interface{} `json:"parameters"`
}
