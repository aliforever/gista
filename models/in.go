package models

type In struct {
	Position              []float64   `json:"position"`
	User                  User        `json:"user"`
	TimeInVideo           interface{} `json:"time_in_video"`
	StartTimeInVideoInSec interface{} `json:"start_time_in_video_in_sec"`
	DurationInVideoInSec  interface{} `json:"duration_in_video_in_sec"`
	Product               Product     `json:"product"`
}
