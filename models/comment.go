package models

type Comment struct {
	Status string `json:"status"`
	UserId int64  `json:"user_id"`
	/*
	 * Unix timestamp (UTC) of when the comment was posted.
	 * Yes this is the UTC timestamp even though its not named "utc"!
	 */
	CreatedAt string `json:"created_at"`
	/*
	 * WARNING: DO NOT USE THIS VALUE! It is NOT a real UTC timestamp.
	 * Instagram has messed up their values of "created_at" vs "created_at_utc".
	 * In `getComments()` both have identical values. In `getCommentReplies()`
	 * both are identical too. But in the `getComments()` "reply previews"
	 * their "created_at_utc" values are completely wrong (always +8 hours into
	 * the future beyond the real UTC time). So just ignore this bad value!
	 * The real app only reads "created_at" for showing comment timestamps!
	 */
	CreatedAtUtc string `json:"created_at_utc"`
	BitFlags     int    `json:"bit_flags"`
	User         User   `json:"user"`
	Pk           int64  `json:"pk"`
	MediaId      string `json:"media_id"`
	Text         string `json:"text"`
	ContentType  string `json:"content_type"`
	/*
	 * A number describing what type of comment this is. Should be compared
	 * against the `Comment::PARENT` and `Comment::CHILD` constants. All
	 * replies are of type `CHILD` and all parents are of type `PARENT`.
	 */
	Type             int  `json:"type"`
	CommentLikeCount int  `json:"comment_like_count"`
	HasLikedComment  bool `json:"has_liked_comment"`
	HasTranslation   bool `json:"has_translation"`
	DidReportAsSpam  bool `json:"did_report_as_spam"`
	/*	`json:""
	 * If this is a child in a thread this is the ID of its parent thread.
	 */
	ParentCommentId string `json:"parent_comment_id"`
	/*
	 * Number of child comments in this comment thread.
	 */
	ChildCommentCount int `json:"child_comment_count"`
	/*
	 * Previews of some of the child comments. Compare it to the child
	 * comment count. If there are more you must request the comment thread.
	 */
	PreviewChildComments []Comment `json:"preview_child_comments"`
	/*
	 * Previews of users in very long comment threads.
	 */
	OtherPreviewUsers              []User `json:"other_preview_users"`
	InlineComposerDisplayCondition string `json:"inline_composer_display_condition"`
	/*
	 * When "has_more_tail_child_comments" is true you can use the value
	 * in "next_max_child_cursor" as "max_id" parameter to load up to
	 * "num_tail_child_comments" older child-comments.
	 */
	HasMoreTailChildComments bool   `json:"has_more_tail_child_comments"`
	NextMaxChildCursor       string `json:"next_max_child_cursor"`
	NumTailChildComments     int    `json:"num_tail_child_comments"`
	/*
	 * When "has_more_head_child_comments" is true you can use the value
	 * in "next_min_child_cursor" as "min_id" parameter to load up to
	 * "num_head_child_comments" newer child-comments.
	 */
	HasMoreHeadChildComments bool   `json:"has_more_head_child_comments"`
	NextMinChildCursor       string `json:"next_min_child_cursor"`
	NumHeadChildComments     int    `json:"num_head_child_comments"`
}
