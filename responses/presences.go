package responses

type Presences struct {
	Response
	UserPresence interface{} `json:"user_presence"`
}
