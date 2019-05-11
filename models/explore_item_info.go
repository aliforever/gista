package models

type ExploreItemInfo struct {
	NumColumns      int    `json:"num_columns"`
	TotalNumColumns int    `json:"total_num_columns"`
	AspectRatio     int    `json:"aspect_ratio"`
	Autoplay        bool   `json:"autoplay"`
	DestinationView string `json:"destination_view"`
}
