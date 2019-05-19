package models

type Caption struct {
	Status          interface{} `json:"status"`
	UserId          int64       `json:"user_id"`
	CreatedAtUtc    interface{} `json:"created_at_utc"`
	CreatedAt       interface{} `json:"created_at"`
	BitFlags        int         `json:"bit_flags"`
	User            User        `json:"user"`
	ContentType     interface{} `json:"content_type"`
	Text            string      `json:"text"`
	MediaId         int         `json:"media_id"`
	Pk              int64       `json:"pk"`
	Type            int         `json:"type"`
	HasTranslation  bool        `json:"has_translation"`
	DidReportAsSpam bool        `json:"did_report_as_spam"`
}
