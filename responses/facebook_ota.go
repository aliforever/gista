package responses

type FacebookOta struct {
	Response
	Bundles   interface{} `json:"bundles"`
	RequestId string      `json:"request_id"`
}
