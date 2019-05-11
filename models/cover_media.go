package models

type CoverMedia struct {
	Id      string `json:"id"`
	MediaId string `json:"media_id"`
	/*
	 * A number describing what type of media this is.
	 */
	MediaType           int            `json:"media_type"`
	ImageVersions2      ImageVersions2 `json:"image_versions_2"`
	OriginalWidth       int            `json:"original_width"`
	OriginalHeight      int            `json:"original_height"`
	CroppedImageVersion ImageCandidate `json:"cropped_image_version"`
	CropRect            []int          `json:"crop_rect"`
	FullImageVersion    ImageCandidate `json:"full_image_version"`
}
