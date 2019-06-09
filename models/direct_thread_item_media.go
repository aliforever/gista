package models

type DirectThreadItemMedia struct {
	MediaType      int             `json:"media_type"`
	ImageVersions2 *ImageVersions2 `json:"image_versions2"`
	VideoVersions  *[]VideoVersion `json:"video_versions"`
	OriginalWidth  int             `json:"original_width"`
	OriginalHeight int             `json:"original_height"`
}
