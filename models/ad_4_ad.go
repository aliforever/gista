package models

type Ad4ad struct {
	Type          interface{} `json:"type"`
	Title         interface{} `json:"title"`
	Media         Item        `json:"media"`
	Footer        interface{} `json:"footer"`
	Id            string      `json:"id"`
	TrackingToken string      `json:"tracking_token"`
}
