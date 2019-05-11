package models

type ExploreItem struct {
	Media           Item            `json:"media"`
	Stories         Stories         `json:"stories"`
	Channel         Channel         `json:"channel"`
	ExploreItemInfo ExploreItemInfo `json:"explore_item_info"`
}
