package models

type MediaData struct {
	ImageVersions2 ImageVersions2 `json:"image_versions2"`
	OriginalWidth  int            `json:"original_width"`
	OriginalHeight int            `json:"original_height"`
	/*
	 * A number describing what type of media this is.
	 */
	MediaType     int            `json:"media_type"`
	VideoVersions []VideoVersion `json:"video_versions"`
}
