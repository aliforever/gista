package responses

type MSISDNHeader struct {
	Response
	PhoneNumber         string `json:"phone_number,omitempty"`
	Url                 string `json:"url,omitempty"`
	RemainingTTLSeconds int    `json:"remaining_ttl_seconds,omitempty"`
	Ttl                 int    `json:"ttl,omitempty"`
}
