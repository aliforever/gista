package models

import "fmt"

type Item struct {
	TakenAt         int64  `json:"taken_at,omitempty"`
	Pk              int64  `json:"pk,omitempty"`
	Id              string `json:"id,omitempty"`
	DeviceTimestamp int64  `json:"device_timestamp,omitempty"`
	/*
	 * A number describing what type of media this is. Should be compared
	 * against the `Item::PHOTO` `Item::VIDEO` and `Item::CAROUSEL` constants!
	 */
	MediaType                    int            `json:"media_type,omitempty"`
	DynamicItemId                string         `json:"dynamic_item_id,omitempty"`
	Code                         string         `json:"code,omitempty"`
	ClientCacheKey               string         `json:"client_cache_key,omitempty"`
	FilterType                   int            `json:"filter_type,omitempty"`
	ProductType                  string         `json:"product_type,omitempty"`
	NearlyCompleteCopyrightMatch bool           `json:"nearly_complete_copyright_match,omitempty"`
	ImageVersions2               ImageVersions2 `json:"image_versions2,omitempty"`
	OriginalWidth                int            `json:"original_width,omitempty"`
	OriginalHeight               int            `json:"original_height,omitempty"`
	CaptionPosition              float64        `json:"caption_position,omitempty"`
	IsReelMedia                  bool           `json:"is_reel_media,omitempty"`
	VideoVersions                []VideoVersion `json:"video_versions,omitempty"`
	HasAudio                     bool           `json:"has_audio,omitempty"`
	VideoDuration                float64        `json:"video_duration,omitempty"`
	User                         User           `json:"user,omitempty"`
	CanSeeInsightsAsBrand        bool           `json:"can_see_insights_as_brand,omitempty"`
	Caption                      Caption        `json:"caption,omitempty"`
	Title                        string         `json:"title,omitempty"`
	CaptionIsEdited              bool           `json:"caption_is_edited,omitempty"`
	PhotoOfYou                   bool           `json:"photo_of_you,omitempty"`
	FbUserTags                   *Usertag       `json:"fb_user_tags,omitempty"`
	CanViewerSave                bool           `json:"can_viewer_save,omitempty"`
	HasViewerSaved               bool           `json:"has_viewer_saved,omitempty"`
	OrganicTrackingToken         string         `json:"organic_tracking_token,omitempty"`
	FollowHashtagInfo            *Hashtag       `json:"follow_hashtag_info,omitempty"`
	ExpiringAt                   int64          `json:"expiring_at,omitempty"`
	IsDashEligible               int            `json:"is_dash_eligible,omitempty"`
	VideoDashManifest            string         `json:"video_dash_manifest,omitempty"`
	NumberOfQualities            int            `json:"number_of_qualities,omitempty"`
	VideoCodec                   string         `json:"video_codec,omitempty"`
	Thumbnails                   *Thumbnail     `json:"thumbnails,omitempty"`
	CanReshare                   bool           `json:"can_reshare,omitempty"`
	CanReply                     bool           `json:"can_reply,omitempty"`
	CanViewerReshare             bool           `json:"can_viewer_reshare,omitempty"`
	Visibility                   interface{}    `json:"visibility,omitempty"`
	Attribution                  *Attribution   `json:"attribution,omitempty"`
	/*
	 * This is actually a float64 in the reply but is always `.0` so we cast
	 * it to an int instead to make the number easier to manage.
	 */
	ViewCount                    int  `json:"view_count,omitempty"`
	ViewerCount                  int  `json:"viewer_count,omitempty"`
	CommentCount                 int  `json:"comment_count,omitempty"`
	CanViewMorePreviewComments   bool `json:"can_view_more_preview_comments,omitempty"`
	HasMoreComments              bool `json:"has_more_comments,omitempty"`
	MaxNumVisiblePreviewComments int  `json:"max_num_visible_preview_comments,omitempty"`
	/*
	 * Preview of comments via feed replies.
	 *
	 * If "has_more_comments" is FALSE then this has ALL of the comments.
	 * Otherwise youll need to get all comments by querying the media.
	 */
	PreviewComments []Comment `json:"preview_comments,omitempty"`
	/*
	 * Comments for the item.
	 *
	 * TODO: As of mid-2017 this field seems to no longer be used for
	 * timeline feed items? They now use "preview_comments" instead. But we
	 * wont delete it since some other feed MAY use this property for ITS
	 * Item object.
	 */
	Comments                            []Comment                     `json:"comments,omitempty"`
	CommentsDisabled                    interface{}                   `json:"comments_disabled,omitempty"`
	ReelMentions                        []ReelMention                 `json:"reel_mentions,omitempty"`
	StoryCta                            []StoryCta                    `json:"story_cta,omitempty"`
	NextMaxId                           interface{}                   `json:"next_max_id,omitempty"`
	CarouselMedia                       []CarouselMedia               `json:"carousel_media,omitempty"`
	CarouselMediaType                   interface{}                   `json:"carousel_media_type,omitempty"`
	CarouselMediaCount                  int                           `json:"carousel_media_count,omitempty"`
	Likers                              []User                        `json:"likers,omitempty"`
	LikeCount                           int                           `json:"like_count,omitempty"`
	Preview                             string                        `json:"preview,omitempty"`
	HasLiked                            bool                          `json:"has_liked,omitempty"`
	ExploreContext                      string                        `json:"explore_context,omitempty"`
	ExploreSourceToken                  string                        `json:"explore_source_token,omitempty"`
	ExploreHideComments                 bool                          `json:"explore_hide_comments,omitempty"`
	Explore                             *Explore                      `json:"explore,omitempty"`
	ImpressionToken                     string                        `json:"impression_token,omitempty"`
	Usertags                            *Usertag                      `json:"usertags,omitempty"`
	Media                               *Media                        `json:"media,omitempty"`
	Stories                             *Stories                      `json:"stories,omitempty"`
	TopLikers                           []string                      `json:"top_likers,omitempty"`
	DirectReplyToAuthorEnabled          bool                          `json:"direct_reply_to_author_enabled,omitempty"`
	SuggestedUsers                      *Suggested                    `json:"suggested_users,omitempty"`
	IsNewSuggestion                     bool                          `json:"is_new_suggestion,omitempty"`
	CommentLikesEnabled                 bool                          `json:"comment_likes_enabled,omitempty"`
	Location                            *Location                     `json:"location,omitempty"`
	Lat                                 float64                       `json:"lat,omitempty"`
	Lng                                 float64                       `json:"lng,omitempty"`
	Channel                             *Channel                      `json:"channel,omitempty"`
	Gating                              *Gating                       `json:"gating,omitempty"`
	Injected                            *Injected                     `json:"injected,omitempty"`
	Placeholder                         *Placeholder                  `json:"placeholder,omitempty"`
	Algorithm                           string                        `json:"algorithm,omitempty"`
	ConnectionId                        string                        `json:"connection_id,omitempty"`
	SocialContext                       string                        `json:"social_context,omitempty"`
	Icon                                interface{}                   `json:"icon,omitempty"`
	MediaIds                            []string                      `json:"media_ids,omitempty"`
	MediaId                             string                        `json:"media_id,omitempty"`
	ThumbnailUrls                       interface{}                   `json:"thumbnail_urls,omitempty"`
	LargeUrls                           interface{}                   `json:"large_urls,omitempty"`
	MediaInfos                          interface{}                   `json:"media_infos,omitempty"`
	Value                               float64                       `json:"value,omitempty"`
	CollapseComments                    bool                          `json:"collapse_comments,omitempty"`
	Link                                string                        `json:"link,omitempty"`
	LinkText                            string                        `json:"link_text,omitempty"`
	LinkHintText                        string                        `json:"link_hint_text,omitempty"`
	ITunesItem                          interface{}                   `json:"i_tunes_item,omitempty"`
	AdHeaderStyle                       int                           `json:"ad_header_style,omitempty"`
	AdMetadata                          []AdMetadata                  `json:"ad_metadata,omitempty"`
	AdAction                            string                        `json:"ad_action,omitempty"`
	AdLinkType                          int                           `json:"ad_link_type,omitempty"`
	DrAdType                            *int                          `json:"dr_ad_type,omitempty"`
	AndroidLinks                        []AndroidLink                 `json:"android_links,omitempty"`
	IosLinks                            []IOSLink                     `json:"ios_links,omitempty"`
	ForceOverlay                        bool                          `json:"force_overlay,omitempty"`
	HideNuxText                         bool                          `json:"hide_nux_text,omitempty"`
	OverlayText                         string                        `json:"overlay_text,omitempty"`
	OverlayTitle                        string                        `json:"overlay_title,omitempty"`
	OverlaySubtitle                     string                        `json:"overlay_subtitle,omitempty"`
	FbPageUrl                           string                        `json:"fb_page_url,omitempty"`
	PlaybackDurationSecs                interface{}                   `json:"playback_duration_secs,omitempty"`
	UrlExpireAtSecs                     interface{}                   `json:"url_expire_at_secs,omitempty"`
	IsSidecarChild                      interface{}                   `json:"is_sidecar_child,omitempty"`
	CommentThreadingEnabled             bool                          `json:"comment_threading_enabled,omitempty"`
	CoverMedia                          *CoverMedia                   `json:"cover_media,omitempty"`
	SavedCollectionIds                  []string                      `json:"saved_collection_ids,omitempty"`
	BoostedStatus                       interface{}                   `json:"boosted_status,omitempty"`
	BoostUnavailableReason              interface{}                   `json:"boost_unavailable_reason,omitempty"`
	Viewers                             []User                        `json:"viewers,omitempty"`
	ViewerCursor                        interface{}                   `json:"viewer_cursor,omitempty"`
	TotalViewerCount                    int                           `json:"total_viewer_count,omitempty"`
	MultiAuthorReelNames                interface{}                   `json:"multi_author_reel_names,omitempty"`
	ScreenshotterUserIds                interface{}                   `json:"screenshotter_user_ids,omitempty"`
	ReelShare                           *ReelShare                    `json:"reel_share,omitempty"`
	OrganicPostId                       string                        `json:"organic_post_id,omitempty"`
	SponsorTags                         []User                        `json:"sponsor_tags,omitempty"`
	StoryPollVoterInfos                 interface{}                   `json:"story_poll_voter_infos,omitempty"`
	ImportedTakenAt                     interface{}                   `json:"imported_taken_at,omitempty"`
	LeadGenFormId                       string                        `json:"lead_gen_form_id,omitempty"`
	AdId                                string                        `json:"ad_id,omitempty"`
	ActorFbid                           string                        `json:"actor_fbid,omitempty"`
	IsAd4ad                             interface{}                   `json:"is_ad_4_ad,omitempty"`
	CommentingDisabledForViewer         interface{}                   `json:"commenting_disabled_for_viewer,omitempty"`
	IsSeen                              interface{}                   `json:"is_seen,omitempty"`
	StoryEvents                         interface{}                   `json:"story_events,omitempty"`
	StoryHashtags                       []StoryHashtag                `json:"story_hashtags,omitempty"`
	StoryPolls                          interface{}                   `json:"story_polls,omitempty"`
	StoryFeedMedia                      interface{}                   `json:"story_feed_media,omitempty"`
	StorySoundOn                        interface{}                   `json:"story_sound_on,omitempty"`
	CreativeConfig                      interface{}                   `json:"creative_config,omitempty"`
	StoryLocations                      []StoryLocation               `json:"story_locations,omitempty"`
	StorySliders                        interface{}                   `json:"story_sliders,omitempty"`
	StoryFriendLists                    interface{}                   `json:"story_friend_lists,omitempty"`
	StoryProductItems                   interface{}                   `json:"story_product_items,omitempty"`
	StoryQuestions                      []StoryQuestion               `json:"story_questions,omitempty"`
	StoryQuestionResponderInfos         []StoryQuestionResponderInfos `json:"story_question_responder_infos,omitempty"`
	StoryCountdowns                     []StoryCountdown              `json:"story_countdowns,omitempty"`
	StoryMusicStickers                  interface{}                   `json:"story_music_stickers,omitempty"`
	SupportsReelReactions               bool                          `json:"supports_reel_reactions,omitempty"`
	ShowOneTapFbShareTooltip            bool                          `json:"show_one_tap_fb_share_tooltip,omitempty"`
	HasSharedToFb                       int                           `json:"has_shared_to_fb,omitempty"`
	MainFeedCarouselStartingMediaId     string                        `json:"main_feed_carousel_starting_media_id,omitempty"`
	MainFeedCarouselHasUnseenCoverMedia bool                          `json:"main_feed_carousel_has_unseen_cover_media,omitempty"`
	InventorySource                     string                        `json:"inventory_source,omitempty"`
	IsEof                               bool                          `json:"is_eof,omitempty"`
	TopFollowers                        []string                      `json:"top_followers,omitempty"`
	TopFollowersCount                   int                           `json:"top_followers_count,omitempty"`
	StoryIsSavedToArchive               bool                          `json:"story_is_saved_to_archive,omitempty"`
	TimezoneOffset                      int                           `json:"timezone_offset,omitempty"`
	ProductTags                         *ProductTags                  `json:"product_tags,omitempty"`
	InlineComposerDisplayCondition      string                        `json:"inline_composer_display_condition,omitempty"`
	InlineComposerImpTriggerTime        int                           `json:"inline_composer_imp_trigger_time,omitempty"`
	HighlightReelIds                    []string                      `json:"highlight_reel_ids,omitempty"`
	TotalScreenshotCount                int                           `json:"total_screenshot_count,omitempty"`
	/*
	 * HTML color string such as "#812A2A".
	 */
	DominantColor string `json:"dominant_color,omitempty"`
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
