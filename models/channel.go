package models

type Channel struct {
	ChannelId   string                   `json:"channel_id"`
	ChannelType interface{}              `json:"channel_type"`
	Title       interface{}              `json:"title"`
	Header      interface{}              `json:"header"`
	MediaCount  int                      `json:"media_count"`
	Media       []map[string]interface{} `json:"media"` /*Recursive error on type Item*/
	Context     interface{}              `json:"context"`
}
