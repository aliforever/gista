package models

type AnimatedMediaImageFixedHeight struct {
	Url      string `json:"url"`
	Width    string `json:"width"`
	Height   string `json:"heigth"`
	Size     string `json:"size"`
	Mp4      string `json:"mp_4"`
	Mp4Size  string `json:"mp_4_size"`
	Webp     string `json:"webp"`
	WebpSize string `json:"webp_size"`
}
