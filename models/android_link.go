package models

type AndroidLink struct {
	LinkType          int    `json:"link_type"`
	WebUri            string `json:"web_uri"`
	AndroidClass      string `json:"android_class"`
	Package           string `json:"package"`
	DeeplinkUri       string `json:"deeplink_uri"`
	CallToActionTitle string `json:"call_to_action_title"`
	RedirectUri       string `json:"redirect_uri"`
	IgUserId          string `json:"ig_user_id"`
	LeadGenFormId     string `json:"lead_gen_form_id"`
	CanvasDocId       string `json:"canvas_doc_id"`
}
