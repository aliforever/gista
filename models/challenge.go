package models

type Challenge struct {
	Url               string `json:"url"`
	ApiPath           string `json:"api_path"`
	HideWebviewHeader bool   `json:"hide_webview_header"`
	Lock              bool   `json:"lock"`
	Logout            bool   `json:"logout"`
	NativeFlow        bool   `json:"native_flow"`
}
