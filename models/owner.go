package models

type Owner struct {
	Type               interface{} `json:"type"`
	Pk                 string      `json:"pk"`
	Name               string      `json:"name"`
	ProfilePicUrl      string      `json:"profile_pic_url"`
	ProfilePicUsername string      `json:"profile_pic_username"`
	ShortName          string      `json:"short_name"`
	Lat                float64     `json:"lat"`
	Lng                float64     `json:"lng"`
	LocationDict       Location    `json:"location_dict"`
}
