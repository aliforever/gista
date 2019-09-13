package models

type SuggestedUsers struct {
	Id               string           `json:"id"`
	ViewAllText      string           `json:"view_all_text"`
	Title            string           `json:"title"`
	AutoDvance       string           `json:"auto_dvance"`
	Type             int              `json:"type"`
	TrackingToken    string           `json:"tracking_token"`
	LandingSiteType  string           `json:"landing_site_type"`
	LandingSiteTitle string           `json:"landing_site_title"`
	UpsellFbPos      string           `json:"upsell_fb_pos"`
	Suggestions      []Suggestion     `json:"suggestions"`
	SuggestionCards  []SuggestionCard `json:"suggestion_cards"`
	NetegoType       string           `json:"netego_type"`
}
