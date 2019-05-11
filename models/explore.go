package models

type Explore struct {
	Explanation interface{} `json:"explanation"`
	ActorId     string      `json:"actor_id"`
	SourceToken interface{} `json:"source_token"`
}
