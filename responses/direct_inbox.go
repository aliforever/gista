package responses

import "github.com/aliforever/gista/models"

type DirectInbox struct {
	Response
	PendingRequestsTotal interface{}        `json:"pending_requests_total"`
	SeqId                int64              `json:"seq_id"`
	Viewer               models.User        `json:"viewer"`
	PendingRequestsUsers []models.User      `json:"pending_requests_users"`
	Inbox                models.DirectInbox `json:"inbox"`
	Megaphone            models.Megaphone   `json:"megaphone"`
	SnapshotAtMs         string             `json:"snapshot_at_ms"`
}
