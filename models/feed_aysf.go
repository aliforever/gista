package models

type FeedAysf struct {
	LandingSiteType  interface{}  `json:"landing_site_type"`
	Uuid             string       `json:"uuid"`
	ViewAllText      interface{}  `json:"view_all_text"`
	FeedPosition     interface{}  `json:"feed_position"`
	LandingSiteTitle interface{}  `json:"landing_site_title"`
	IsDismissable    interface{}  `json:"is_dismissable"`
	Suggestions      []Suggestion `json:"suggestions"`
	ShouldRefill     interface{}  `json:"should_refill"`
	DisplayNewUnit   interface{}  `json:"display_new_unit"`
	FetchUserDetails interface{}  `json:"fetch_user_details"`
	Title            interface{}  `json:"title"`
	Activator        interface{}  `json:"activator"`
}
