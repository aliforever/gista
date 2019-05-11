package models

type Challenge struct {
	Url               string      `json:"url"`
	ApiPath           interface{} `json:"api_path"`
	HideWebviewHeader interface{} `json:"hide_webview_header"`
	Lock              interface{} `json:"lock"`
	Logout            interface{} `json:"logout"`
	NativeFlow        interface{} `json:"native_flow"`
}
