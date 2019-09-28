package responses

type ResumableUpload struct {
	Response
	XsharingNonces interface{} `json:"xsharing_nonces"`
	UploadId       int         `json:"upload_id"`
}
