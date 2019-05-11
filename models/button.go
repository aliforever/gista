package models

type Button struct {
	Text            string      `json:"text"`
	Url             string      `json:"url"`
	Action          interface{} `json:"action"`
	BackgroundColor interface{} `json:"background_color"`
	BorderColor     interface{} `json:"border_color"`
	TextColor       interface{} `json:"text_color"`
	ActionInfo      interface{} `json:"action_info"`
}
