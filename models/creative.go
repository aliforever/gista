package models

type Creative struct {
	Title           Text        `json:"title"`
	Footer          Text        `json:"footer"`
	SocialContext   Text        `json:"social_context"`
	Content         Text        `json:"content"`
	PrimaryAction   Action      `json:"primary_action"`
	SecondaryAction Action      `json:"secondary_action"`
	DismissAction   interface{} `json:"dismiss_action"`
	Image           Image       `json:"image"`
}
