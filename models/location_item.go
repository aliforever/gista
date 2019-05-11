package models

type LocationItem struct {
	MediaBundles interface{} `json:"media_bundles"`
	Subtitle     interface{} `json:"subtitle"`
	Location     Location    `json:"location"`
	Title        interface{} `json:"title"`
}
