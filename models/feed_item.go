package models

type FeedItem struct {
	MediaOrAd           Item          `json:"media_or_ad"`
	StoriesNetego       StoriesNetego `json:"stories_netego"`
	Ad4ad               Ad4ad         `json:"ad_4_ad"`
	SuggestedUsers      Suggested     `json:"suggested_users"`
	EndOfFeedDemarcator interface{}   `json:"end_of_feed_demarcator"`
	AdLinkType          int           `json:"ad_link_type"`
}
