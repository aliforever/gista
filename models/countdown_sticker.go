package models

type CountdownSticker struct {
	CountdownId string `json:"countdown_id"`
	EndTs       string `json:"end_ts"`
	Text        string `json:"text"`
	/*
	 * HTML color string such as "#812A2A".
	 */
	TextColor string `json:"text_color"`
	/*
	 * HTML color string such as "#812A2A".
	 */
	StartBackgroundColor string `json:"start_background_color"`
	/*
	 * HTML color string such as "#812A2A".
	 */
	EndBackgroundColor string `json:"end_background_color"`
	/*
	 * HTML color string such as "#812A2A".
	 */
	DigitColor string `json:"digit_color"`
	/*
	 * HTML color string such as "#812A2A".
	 */
	DigitCardColor    string      `json:"digit_card_color"`
	FollowingEnabled  bool        `json:"following_enabled"`
	IsOwner           bool        `json:"is_owner"`
	Attribution       interface{} `json:"attribution"`
	ViewerIsFollowing bool        `json:"viewer_is_following"`
}
