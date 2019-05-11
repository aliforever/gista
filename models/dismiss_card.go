package models

type DismissCard struct {
	CardId       interface{} `json:"card_id"`
	ImageUrl     string      `json:"image_url"`
	Title        interface{} `json:"title"`
	Message      interface{} `json:"message"`
	ButtonText   interface{} `json:"button_text"`
	CameraTarget interface{} `json:"camera_target"`
	FaceFilterId interface{} `json:"face_filter_id"`
}
