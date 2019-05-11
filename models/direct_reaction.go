package models

type DirectReaction struct {
	ReactionType   string `json:"reaction_type"`
	Timestamp      string `json:"timestamp"`
	SenderId       string `json:"sender_id"`
	ClientContext  string `json:"client_context"`
	ReactionStatus string `json:"reaction_status"`
	NodeType       string `json:"node_type"`
	ItemId         string `json:"item_id"`
}
