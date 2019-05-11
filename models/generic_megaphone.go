package models

type GenericMegaphone struct {
	Type             interface{} `json:"type"`
	Title            interface{} `json:"title"`
	Message          interface{} `json:"message"`
	Dismissible      interface{} `json:"dismissible"`
	Icon             interface{} `json:"icon"`
	Buttons          []Button    `json:"buttons"`
	MegaphoneVersion interface{} `json:"megaphone_version"`
	ButtonLayout     interface{} `json:"button_layout"`
	ActionInfo       interface{} `json:"action_info"`
	ButtonLocation   interface{} `json:"button_location"`
	BackgroundColor  interface{} `json:"background_color"`
	TitleColor       interface{} `json:"title_color"`
	MessageColor     interface{} `json:"message_color"`
	Uuid             string      `json:"uuid"`
}
