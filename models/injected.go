package models

type Injected struct {
	Label                      string       `json:"label"`
	ShowIcon                   bool         `json:"show_icon"`
	HideLabel                  string       `json:"hide_label"`
	Invalidation               interface{}  `json:"invalidation"` // Only encountered as NULL.
	IsDemo                     bool         `json:"is_demo"`
	ViewTags                   interface{}  `json:"view_tags"` // Only seen as [].
	IsHoldout                  bool         `json:"is_holdout"`
	TrackingToken              string       `json:"tracking_token"`
	ShowAdChoices              bool         `json:"show_ad_choices"`
	AdTitle                    string       `json:"ad_title"`
	AboutAdParams              string       `json:"about_ad_params"`
	DirectShare                bool         `json:"direct_share"`
	AdId                       string       `json:"ad_id"`
	DisplayViewabilityEligible bool         `json:"display_viewability_eligible"`
	FbPageUrl                  string       `json:"fb_page_url"`
	HideReasonsV2              []HideReason `json:"hide_reasons_v_2"`
	HideFlowType               int          `json:"hide_flow_type"`
	Cookies                    []string     `json:"cookies"`
	LeadGenFormId              string       `json:"lead_gen_form_id"`
}
