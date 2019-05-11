package models

type HideReason struct {
	/*
	 * A human string such as "Its not relevant" and "I see it too often".
	 */
	Text string `json:"text"`
	/*
	 * A computer string such as "NOT_RELEVANT" or "KEEP_SEEING_THIS".
	 */
	Reason string `json:"reason"`
}
