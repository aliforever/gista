package models

type CarouselMedia struct {
	Pk                string         `json:"pk"`
	Id                string         `json:"id"`
	CarouselParentId  string         `json:"carousel_parent_id"`
	FbUserTags        Usertag        `json:"fb_user_tags"`
	NumberOfQualities int            `json:"number_of_qualities"`
	IsDashEligible    int            `json:"is_dash_eligible"`
	VideoDashManifest string         `json:"video_dash_manifest"`
	ImageVersions2    ImageVersions2 `json:"image_versions_2"`
	VideoVersions     []VideoVersion `json:"video_versions"`
	HasAudio          bool           `json:"has_audio"`
	VideoDuration     float64        `json:"video_duration"`
	VideoSubtitlesUri string         `json:"video_subtitles_uri"`
	OriginalHeight    int            `json:"original_height"`
	OriginalWidth     int            `json:"original_width"`
	/*
	 * A number describing what type of media this is. Should be compared
	 * against the `CarouselMedia::PHOTO` and `CarouselMedia::VIDEO`
	 * constants!
	 */
	MediaType       int           `json:"media_type"`
	DynamicItemId   string        `json:"dynamic_item_id"`
	Usertags        Usertag       `json:"usertags"`
	Preview         string        `json:"preview"`
	Headline        Headline      `json:"headline"`
	Link            string        `json:"link"`
	LinkText        string        `json:"link_text"`
	LinkHintText    string        `json:"link_hint_text"`
	AndroidLinks    []AndroidLink `json:"android_links"`
	AdMetadata      []AdMetadata  `json:"ad_metadata"`
	AdAction        string        `json:"ad_action"`
	AdLinkType      int           `json:"ad_link_type"`
	ForceOverlay    bool          `json:"force_overlay"`
	HideNuxText     bool          `json:"hide_nux_text"`
	OverlayText     string        `json:"overlay_text"`
	OverlayTitle    string        `json:"overlay_title"`
	OverlaySubtitle string        `json:"overlay_subtitle"`
	/*
	 * HTML color string such as "#812A2A".
	 */
	DominantColor string `json:"dominant_color"`
}
