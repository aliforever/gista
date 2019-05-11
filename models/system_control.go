package models

type SystemControl struct {
	UploadMaxBytes       int `json:"upload_max_bytes"`
	UploadTimePeriodSec  int `json:"upload_time_period_sec"`
	UploadBytesPerUpdate int `json:"upload_bytes_per_update"`
}
