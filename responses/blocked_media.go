package responses

type BlockedMedia struct {
	Response
	MediaIds interface{} `json:"media_ids"`
}
