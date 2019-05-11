package models

type User struct {
	Username                       string `json:"username"`
	HasAnonymousProfilePicture     bool   `json:"has_anonymous_profile_picture"`
	HasHighlightReels              bool   `json:"has_highlight_reels"`
	IsFavorite                     bool   `json:"is_favorite"`
	IsFavoriteForStories           bool   `json:"is_favorite_for_stories"`
	IsFavoriteForHighlights        bool   `json:"is_favorite_for_highlights"`
	IsInterestAccount              bool   `json:"is_interest_account"`
	canBeReportedAsFraud           bool
	ProfilePicUrl                  string               `json:"profile_pic_url"`
	ProfilePicId                   string               `json:"profile_pic_id"`
	Permission                     bool                 `json:"permission"`
	FullName                       string               `json:"full_name"`
	UserId                         string               `json:"user_id"`
	Pk                             int64                `json:"pk"`
	Id                             *int64               `json:"id,omitempty"`
	IsVerified                     bool                 `json:"is_verified"`
	IsPrivate                      bool                 `json:"is_private"`
	CoeffWeight                    interface{}          `json:"coeff_weight"`
	FriendshipStatus               FriendshipStatus     `json:"friendship_status"`
	HdProfilePicVersions           []ImageCandidate     `json:"hd_profile_pic_versions"`
	Byline                         interface{}          `json:"byline"`
	SearchSocialContext            interface{}          `json:"search_social_context"`
	UnseenCount                    interface{}          `json:"unseen_count"`
	MutualFollowersCount           int                  `json:"mutual_followers_count"`
	FollowerCount                  int                  `json:"follower_count"`
	SearchSubtitle                 string               `json:"search_subtitle"`
	SocialContext                  interface{}          `json:"social_context"`
	MediaCount                     int                  `json:"media_count"`
	FollowingCount                 int                  `json:"following_count"`
	FollowingTagCount              int                  `json:"following_tag_count"`
	IsBusiness                     bool                 `json:"is_business"`
	UserTagsCount                  int                  `json:"usertags_count"`
	ProfileContext                 interface{}          `json:"profile_context"`
	Biography                      string               `json:"biography"`
	GeoMediaCount                  int                  `json:"geo_media_count"`
	IsUnpublished                  bool                 `json:"is_unpublished"`
	AllowContactsSync              interface{}          `json:"allow_contacts_sync"`
	ShowFeedBizConversionIcon      interface{}          `json:"show_feed_biz_conversion_icon"`
	AutoExpandChaining             interface{}          `json:"auto_expand_chaining"`
	CanBoostPost                   interface{}          `json:"can_boost_post"`
	IsProfileActionNeeded          bool                 `json:"is_profile_action_needed"`
	HasChaining                    bool                 `json:"has_chaining"`
	HasRecommendAccounts           bool                 `json:"has_recommend_accounts"`
	ChainingSuggestions            []ChainingSuggestion `json:"chaining_suggestions"`
	IncludeDirectBlacklistStatus   interface{}          `json:"include_direct_blacklist_status"`
	CanSeeOrganicInsights          bool                 `json:"can_see_organic_insights"`
	HasPlacedOrders                bool                 `json:"has_placed_orders"`
	CanConvertToBusiness           bool                 `json:"can_convert_to_business"`
	ConvertFromPages               interface{}          `json:"convert_from_pages"`
	ShowBusinessConversionIcon     bool                 `json:"show_business_conversion_icon"`
	ShowConversionEditEntry        bool                 `json:"show_conversion_edit_entry"`
	ShowInsightsTerms              bool                 `json:"show_insights_terms"`
	CanCreateSponsorTags           interface{}          `json:"can_create_sponsor_tags"`
	HdProfilePicUrlInfo            ImageCandidate       `json:"hd_profile_pic_url_info"`
	UserTagReviewEnabled           interface{}          `json:"usertag_review_enabled"`
	ProfileContextMutualFollowIds  []string             `json:"profile_context_mutual_follow_ids"`
	ProfileContextLinksWithUserIds []Link               `json:"profile_context_links_with_user_ids"`
	HasBiographyTranslation        bool                 `json:"has_biography_translation"`
	TotalIGTVVideos                int                  `json:"total_igtv_videos"`
	TotalArEffects                 int                  `json:"total_ar_effects"`
	CanLinkEntitiesInBio           bool                 `json:"can_link_entities_in_bio"`
	BiographyWithEntities          BiographyEntities    `json:"biography_with_entities"`
	MaxNumLinkedEntitiesInBio      int                  `json:"max_num_linked_entities_in_bio"`
	BusinessContactMethod          string               `json:"business_contact_method"`
	HighlightReShareDisabled       bool                 `json:"highlight_reshare_disabled"`
	/*
	 * Business category.
	 */
	Category                       string      `json:"category"`
	DirectMessaging                string      `json:"direct_messaging"`
	PageName                       interface{} `json:"page_name"`
	FbPageCallToActionId           string      `json:"fb_page_call_to_action_id"`
	IsCallToActionEnabled          bool        `json:"is_call_to_action_enabled"`
	AccountType                    int         `json:"account_type"`
	PublicPhoneCountryCode         string      `json:"public_phone_country_code"`
	PublicPhoneNumber              string      `json:"public_phone_number"`
	ContactPhoneNumber             string      `json:"contact_phone_number"`
	Latitude                       float64     `json:"latitude"`
	Longitude                      float64     `json:"longitude"`
	AddressStreet                  string      `json:"address_street"`
	Zip                            string      `json:"zip"`
	CityId                         string      `json:"city_id"` // 64-bit number.
	CityName                       string      `json:"city_name"`
	PublicEmail                    string      `json:"public_email"`
	IsNeedy                        bool        `json:"is_needy"`
	ExternalUrl                    string      `json:"external_url"`
	ExternalLynxUrl                string      `json:"external_lynx_url"`
	Email                          string      `json:"email"`
	CountryCode                    int         `json:"country_code"`
	Birthday                       interface{} `json:"birthday"`
	NationalNumber                 string      `json:"national_number"` // Really int but may be >32bit.
	Gender                         int         `json:"gender"`
	PhoneNumber                    string      `json:"phone_number"`
	NeedsEmailConfirm              interface{} `json:"needs_email_confirm"`
	IsActive                       bool        `json:"is_active"`
	BlockAt                        interface{} `json:"block_at"`
	AggregatePromoteEngagement     interface{} `json:"aggregate_promote_engagement"`
	FBUID                          interface{} `json:"fbuid"`
	PageId                         string      `json:"page_id"`
	CanClaimPage                   bool        `json:"can_claim_page"`
	FbPageCallToActionIxAppId      int         `json:"fb_page_call_to_action_ix_app_id"`
	FbPageCallToActionIxUrl        string      `json:"fb_page_call_to_action_ix_url"`
	CanCrossPostWithoutFbToken     bool        `json:"can_crosspost_without_fb_token"`
	FbPageCallToActionIxPartner    string      `json:"fb_page_call_to_action_ix_partner"`
	ShoppablePostsCount            int         `json:"shoppable_posts_count"`
	ShowShoppableFeed              bool        `json:"show_shoppable_feed"`
	ShowAccountTransparencyDetails bool        `json:"show_account_transparency_details"`
	/*
	 * Unix "taken_at" timestamp of the newest item in their story reel.
	 */
	LatestReelMedia        string      `json:"latest_reel_media"`
	HasUnseenBestiesMedia  bool        `json:"has_unseen_besties_media"`
	AllowedCommenterType   string      `json:"allowed_commenter_type"`
	ReelAutoArchive        string      `json:"reel_auto_archive"`
	IsDirectAppInstalled   bool        `json:"is_directapp_installed"`
	BestiesCount           int         `json:"besties_count"`
	CanBeTaggedAsSponsor   bool        `json:"can_be_tagged_as_sponsor"`
	CanFollowHashtag       bool        `json:"can_follow_hashtag"`
	HasProfileVideoFeed    bool        `json:"has_profile_video_feed"`
	IsVideoCreator         bool        `json:"is_video_creator"`
	ShowBestiesBadge       bool        `json:"show_besties_badge"`
	ScreenShotted          bool        `json:"screenshotted"`
	Nametag                Nametag     `json:"nametag"`
	School                 interface{} `json:"school"`
	IsBestie               bool        `json:"is_bestie"`
	LiveSubscriptionStatus string      `json:"live_subscription_status"`
}
