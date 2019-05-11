package models

import "fmt"

type Item struct {
	TakenAt         int64  `json:"taken_at"`
	Pk              int64  `json:"pk"`
	Id              string `json:"id"`
	DeviceTimestamp int64  `json:"device_timestamp"`
	/*
	 * A number describing what type of media this is. Should be compared
	 * against the `Item::PHOTO` `Item::VIDEO` and `Item::CAROUSEL` constants!
	 */
	MediaType                    int            `json:"media_type"`
	DynamicItemId                string         `json:"dynamic_item_id"`
	Code                         string         `json:"code"`
	ClientCacheKey               string         `json:"client_cache_key"`
	FilterType                   int            `json:"filter_type"`
	ProductType                  string         `json:"product_type"`
	NearlyCompleteCopyrightMatch bool           `json:"nearly_complete_copyright_match"`
	ImageVersions2               ImageVersions2 `json:"image_versions_2"`
	OriginalWidth                int            `json:"original_width"`
	OriginalHeight               int            `json:"original_height"`
	CaptionPosition              float64        `json:"caption_position"`
	IsReelMedia                  bool           `json:"is_reel_media"`
	VideoVersions                []VideoVersion `json:"video_versions"`
	HasAudio                     bool           `json:"has_audio"`
	VideoDuration                float64        `json:"video_duration"`
	User                         User           `json:"user"`
	CanSeeInsightsAsBrand        bool           `json:"can_see_insights_as_brand"`
	Caption                      Caption        `json:"caption"`
	Title                        string         `json:"title"`
	CaptionIsEdited              bool           `json:"caption_is_edited"`
	PhotoOfYou                   bool           `json:"photo_of_you"`
	FbUserTags                   Usertag        `json:"fb_user_tags"`
	CanViewerSave                bool           `json:"can_viewer_save"`
	HasViewerSaved               bool           `json:"has_viewer_saved"`
	OrganicTrackingToken         string         `json:"organic_tracking_token"`
	FollowHashtagInfo            Hashtag        `json:"follow_hashtag_info"`
	ExpiringAt                   int64          `json:"expiring_at"`
	IsDashEligible               int            `json:"is_dash_eligible"`
	VideoDashManifest            string         `json:"video_dash_manifest"`
	NumberOfQualities            int            `json:"number_of_qualities"`
	VideoCodec                   string         `json:"video_codec"`
	Thumbnails                   Thumbnail      `json:"thumbnails"`
	CanReshare                   bool           `json:"can_reshare"`
	CanReply                     bool           `json:"can_reply"`
	CanViewerReshare             bool           `json:"can_viewer_reshare"`
	Visibility                   interface{}    `json:"visibility"`
	Attribution                  Attribution    `json:"attribution"`
	/*
	 * This is actually a float64 in the reply but is always `.0` so we cast
	 * it to an int instead to make the number easier to manage.
	 */
	ViewCount                    int  `json:"view_count"`
	ViewerCount                  int  `json:"viewer_count"`
	CommentCount                 int  `json:"comment_count"`
	CanViewMorePreviewComments   bool `json:"can_view_more_preview_comments"`
	HasMoreComments              bool `json:"has_more_comments"`
	MaxNumVisiblePreviewComments int  `json:"max_num_visible_preview_comments"`
	/*
	 * Preview of comments via feed replies.
	 *
	 * If "has_more_comments" is FALSE then this has ALL of the comments.
	 * Otherwise youll need to get all comments by querying the media.
	 */
	PreviewComments []Comment `json:"preview_comments"`
	/*
	 * Comments for the item.
	 *
	 * TODO: As of mid-2017 this field seems to no longer be used for
	 * timeline feed items? They now use "preview_comments" instead. But we
	 * wont delete it since some other feed MAY use this property for ITS
	 * Item object.
	 */
	Comments                            []Comment                     `json:"comments"`
	CommentsDisabled                    interface{}                   `json:"comments_disabled"`
	ReelMentions                        []ReelMention                 `json:"reel_mentions"`
	StoryCta                            []StoryCta                    `json:"story_cta"`
	NextMaxId                           string                        `json:"next_max_id"`
	CarouselMedia                       []CarouselMedia               `json:"carousel_media"`
	CarouselMediaType                   interface{}                   `json:"carousel_media_type"`
	CarouselMediaCount                  int                           `json:"carousel_media_count"`
	Likers                              []User                        `json:"likers"`
	LikeCount                           int                           `json:"like_count"`
	Preview                             string                        `json:"preview"`
	HasLiked                            bool                          `json:"has_liked"`
	ExploreContext                      string                        `json:"explore_context"`
	ExploreSourceToken                  string                        `json:"explore_source_token"`
	ExploreHideComments                 bool                          `json:"explore_hide_comments"`
	Explore                             Explore                       `json:"explore"`
	ImpressionToken                     string                        `json:"impression_token"`
	Usertags                            Usertag                       `json:"usertags"`
	Media                               Media                         `json:"media"`
	Stories                             Stories                       `json:"stories"`
	TopLikers                           []string                      `json:"top_likers"`
	DirectReplyToAuthorEnabled          bool                          `json:"direct_reply_to_author_enabled"`
	SuggestedUsers                      Suggested                     `json:"suggested_users"`
	IsNewSuggestion                     bool                          `json:"is_new_suggestion"`
	CommentLikesEnabled                 bool                          `json:"comment_likes_enabled"`
	Location                            Location                      `json:"location"`
	Lat                                 float64                       `json:"lat"`
	Lng                                 float64                       `json:"lng"`
	Channel                             Channel                       `json:"channel"`
	Gating                              Gating                        `json:"gating"`
	Injected                            Injected                      `json:"injected"`
	Placeholder                         Placeholder                   `json:"placeholder"`
	Algorithm                           string                        `json:"algorithm"`
	ConnectionId                        string                        `json:"connection_id"`
	SocialContext                       string                        `json:"social_context"`
	Icon                                interface{}                   `json:"icon"`
	MediaIds                            []string                      `json:"media_ids"`
	MediaId                             string                        `json:"media_id"`
	ThumbnailUrls                       interface{}                   `json:"thumbnail_urls"`
	LargeUrls                           interface{}                   `json:"large_urls"`
	MediaInfos                          interface{}                   `json:"media_infos"`
	Value                               float64                       `json:"value"`
	CollapseComments                    bool                          `json:"collapse_comments"`
	Link                                string                        `json:"link"`
	LinkText                            string                        `json:"link_text"`
	LinkHintText                        string                        `json:"link_hint_text"`
	ITunesItem                          interface{}                   `json:"i_tunes_item"`
	AdHeaderStyle                       int                           `json:"ad_header_style"`
	AdMetadata                          []AdMetadata                  `json:"ad_metadata"`
	AdAction                            string                        `json:"ad_action"`
	AdLinkType                          int                           `json:"ad_link_type"`
	DrAdType                            *int                          `json:"dr_ad_type,omitempty"`
	AndroidLinks                        []AndroidLink                 `json:"android_links"`
	IosLinks                            []IOSLink                     `json:"ios_links"`
	ForceOverlay                        bool                          `json:"force_overlay"`
	HideNuxText                         bool                          `json:"hide_nux_text"`
	OverlayText                         string                        `json:"overlay_text"`
	OverlayTitle                        string                        `json:"overlay_title"`
	OverlaySubtitle                     string                        `json:"overlay_subtitle"`
	FbPageUrl                           string                        `json:"fb_page_url"`
	PlaybackDurationSecs                interface{}                   `json:"playback_duration_secs"`
	UrlExpireAtSecs                     interface{}                   `json:"url_expire_at_secs"`
	IsSidecarChild                      interface{}                   `json:"is_sidecar_child"`
	CommentThreadingEnabled             bool                          `json:"comment_threading_enabled"`
	CoverMedia                          CoverMedia                    `json:"cover_media"`
	SavedCollectionIds                  []string                      `json:"saved_collection_ids"`
	BoostedStatus                       interface{}                   `json:"boosted_status"`
	BoostUnavailableReason              interface{}                   `json:"boost_unavailable_reason"`
	Viewers                             []User                        `json:"viewers"`
	ViewerCursor                        interface{}                   `json:"viewer_cursor"`
	TotalViewerCount                    int                           `json:"total_viewer_count"`
	MultiAuthorReelNames                interface{}                   `json:"multi_author_reel_names"`
	ScreenshotterUserIds                interface{}                   `json:"screenshotter_user_ids"`
	ReelShare                           ReelShare                     `json:"reel_share"`
	OrganicPostId                       string                        `json:"organic_post_id"`
	SponsorTags                         []User                        `json:"sponsor_tags"`
	StoryPollVoterInfos                 interface{}                   `json:"story_poll_voter_infos"`
	ImportedTakenAt                     interface{}                   `json:"imported_taken_at"`
	LeadGenFormId                       string                        `json:"lead_gen_form_id"`
	AdId                                string                        `json:"ad_id"`
	ActorFbid                           string                        `json:"actor_fbid"`
	IsAd4ad                             interface{}                   `json:"is_ad_4_ad"`
	CommentingDisabledForViewer         interface{}                   `json:"commenting_disabled_for_viewer"`
	IsSeen                              interface{}                   `json:"is_seen"`
	StoryEvents                         interface{}                   `json:"story_events"`
	StoryHashtags                       []StoryHashtag                `json:"story_hashtags"`
	StoryPolls                          interface{}                   `json:"story_polls"`
	StoryFeedMedia                      interface{}                   `json:"story_feed_media"`
	StorySoundOn                        interface{}                   `json:"story_sound_on"`
	CreativeConfig                      interface{}                   `json:"creative_config"`
	StoryLocations                      []StoryLocation               `json:"story_locations"`
	StorySliders                        interface{}                   `json:"story_sliders"`
	StoryFriendLists                    interface{}                   `json:"story_friend_lists"`
	StoryProductItems                   interface{}                   `json:"story_product_items"`
	StoryQuestions                      []StoryQuestion               `json:"story_questions"`
	StoryQuestionResponderInfos         []StoryQuestionResponderInfos `json:"story_question_responder_infos"`
	StoryCountdowns                     []StoryCountdown              `json:"story_countdowns"`
	StoryMusicStickers                  interface{}                   `json:"story_music_stickers"`
	SupportsReelReactions               bool                          `json:"supports_reel_reactions"`
	ShowOneTapFbShareTooltip            bool                          `json:"show_one_tap_fb_share_tooltip"`
	HasSharedToFb                       bool                          `json:"has_shared_to_fb"`
	MainFeedCarouselStartingMediaId     string                        `json:"main_feed_carousel_starting_media_id"`
	MainFeedCarouselHasUnseenCoverMedia bool                          `json:"main_feed_carousel_has_unseen_cover_media"`
	InventorySource                     string                        `json:"inventory_source"`
	IsEof                               bool                          `json:"is_eof"`
	TopFollowers                        []string                      `json:"top_followers"`
	TopFollowersCount                   int                           `json:"top_followers_count"`
	StoryIsSavedToArchive               bool                          `json:"story_is_saved_to_archive"`
	TimezoneOffset                      int                           `json:"timezone_offset"`
	ProductTags                         ProductTags                   `json:"product_tags"`
	InlineComposerDisplayCondition      string                        `json:"inline_composer_display_condition"`
	InlineComposerImpTriggerTime        int                           `json:"inline_composer_imp_trigger_time"`
	HighlightReelIds                    []string                      `json:"highlight_reel_ids"`
	TotalScreenshotCount                int                           `json:"total_screenshot_count"`
	/*
	 * HTML color string such as "#812A2A".
	 */
	DominantColor string `json:"dominant_color"`
}

func (i *Item) GetItemUrl() string {
	return fmt.Sprintf("https://www.instagram.com/p/%s/", i.Code)
}

func (i *Item) IsAd() bool {
	if i.DrAdType != nil {
		return true
	} else {
		return false
	}
}
