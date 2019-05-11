package models

type Param struct {
	Name  *string `json:"name,omitempty"`
	Value string  `json:"value,omitempty"`
}
