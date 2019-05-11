package models

type FeedItem struct {
	MediaOrAd           *Item          `json:"media_or_ad,omitempty"`
	StoriesNetego       *StoriesNetego `json:"stories_netego,omitempty"`
	Ad4ad               *Ad4ad         `json:"ad_4_ad,omitempty"`
	SuggestedUsers      *Suggested     `json:"suggested_users,omitempty"`
	EndOfFeedDemarcator interface{}    `json:"end_of_feed_demarcator,omitempty"`
	AdLinkType          int            `json:"ad_link_type,omitempty"`
}
