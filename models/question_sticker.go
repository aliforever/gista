package models

type QuestionSticker struct {
	QuestionId string `json:"question_id"`
	Question   string `json:"question"`
	/*
	 * HTML color string such as "#812A2A".
	 */
	TextColor string `json:"text_color"`
	/*
	 * HTML color string such as "#812A2A".
	 */
	BackgroundColor   string `json:"background_color"`
	ViewerCanInteract bool   `json:"viewer_can_interact"`
	ProfilePicUrl     string `json:"profile_pic_url"`
	QuestionType      string `json:"question_type"`
}
