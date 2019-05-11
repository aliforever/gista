package models

type ActionLog struct {
	Bold        []Bold      `json:"bold"`
	Description interface{} `json:"description"`
}
