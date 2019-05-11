package models

type Cursor struct {
	CursorTimestampSeconds float64 `json:"cursor_timestamp_seconds"`
	CursorThreadV2Id       float64 `json:"cursor_thread_v2_id"`
}
