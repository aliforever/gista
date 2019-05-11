package models

type DirectInbox struct {
	HasOlder            bool           `json:"has_older"`
	UnseenCount         int            `json:"unseen_count"`
	UnseenCountTs       int64          `json:"unseen_count_ts"` // Is a timestamp.
	BlendedInboxEnabled bool           `json:"blended_inbox_enabled"`
	OldestCursor        Cursor         `json:"oldest_cursor"`
	PrevCursor          Cursor         `json:"prev_cursor"`
	NextCursor          Cursor         `json:"next_cursor"`
	Threads             []DirectThread `json:"threads"`
}
