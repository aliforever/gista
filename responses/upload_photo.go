package responses

type UploadPhoto struct {
	Response
	UploadId string `json:"upload_id"`
	MediaId  string `json:"media_id"`
}
