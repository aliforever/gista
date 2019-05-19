package models

type QpData struct {
	Surface string       `json:"surface"`
	Data    QpViewerData `json:"data"`
}
