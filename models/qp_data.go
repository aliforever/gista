package models

type QpData struct {
	Surface int          `json:"surface"`
	Data    QpViewerData `json:"data"`
}
