package models

type Location struct {
	Name               string                 `json:"name"`
	ExternalIdSource   string                 `json:"external_id_source"`
	ExternalSource     string                 `json:"external_source"`
	Address            string                 `json:"address"`
	Lat                float64                `json:"lat"`
	Lng                float64                `json:"lng"`
	ExternalId         string                 `json:"external_id"`
	FacebookPlacesId   string                 `json:"facebook_places_id"`
	City               string                 `json:"city"`
	Pk                 string                 `json:"pk"`
	ShortName          string                 `json:"short_name"`
	FacebookEventsId   string                 `json:"facebook_events_id"`
	StartTime          interface{}            `json:"start_time"`
	EndTime            interface{}            `json:"end_time"`
	LocationDict       map[string]interface{} `json:"location_dict"` /*recursive error on Location type*/
	Type               interface{}            `json:"type"`
	ProfilePicUrl      string                 `json:"profile_pic_url"`
	ProfilePicUsername string                 `json:"profile_pic_username"`
	TimeGranularity    interface{}            `json:"time_granularity"`
	Timezone           interface{}            `json:"timezone"`
	/*
	 * Country number such as int(398) but it has no relation to actual
	 * country codes so the number is useless...
	 */
	Country int `json:"country"`
	/*
	 * Regular unix timestamp of when the location was created.
	 */
	CreatedAt string `json:"created_at"`
	/*
	 * Some kind of internal number to signify what type of event a special
	 * location (such as a festival) is. Weve only seen this with int(0).
	 */
	EventCategory int `json:"event_category"`
	/*
	 * 64-bit integer with the facebook places ID for the location.
	 */
	PlaceFbid string `json:"place_fbid"`
	/*
	 * Human-readable name of the facebook place for the location.
	 */
	PlaceName string `json:"place_name"`
}
