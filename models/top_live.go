package models

type TopLive struct {
	BroadcastOwners []User      `json:"broadcast_owners"`
	RankedPosition  interface{} `json:"ranked_position"`
}
