package constants

// Internals
const MSISDN string = "accounts/read_msisdn_header/"
const Sync string = "qe/sync/"
const LauncherSync string = "launcher/sync/"
const LogAttribution string = "attribution/log_attribution/"
const ZeroRatingToken string = "zr/token/result/"
const LoomFetchConfig string = "loom/fetch_config/"
const ProfileNotice string = "users/profile_notice/"
const FetchQPData string = "qp/batch_fetch/"
const FacebookOTA string = "facebook_ota/"
const UploadPhoto string = "upload/photo/"

// Account
const ContactPointPreFill string = "accounts/contact_point_prefill/"
const Login string = "accounts/login/"
const TwoFactorLogin string = "accounts/two_factor_login/"
const CurrentUser string = "accounts/current_user/"
const SetBiography string = "accounts/set_biography/"
const EditProfile string = "accounts/set_biography/"
const ChangeProfilePicture string = "accounts/change_profile_picture/"
const RemoveProfilePicture string = "accounts/remove_profile_picture/"
const SetPublic string = "accounts/set_public/"
const SetPrivate string = "accounts/set_private/"
const SwitchToBusinessProfile string = "business_conversion/get_business_convert_social_context/"
const SwitchToPersonalProfile string = "business_conversion/convert_to_personal/"
const CreateBusinessInfo string = "business_conversion/create_business_info/"
const CheckUsername string = "users/check_username/"
const AgreeConsentFirstStep string = "consent/existing_user_flow/"
const GetCommentFilter string = "accounts/get_comment_filter/"
const SetCommentFilter string = "accounts/set_comment_filter/"
const GetCommentCategoryFilterDisabled string = "accounts/get_comment_category_filter_disabled/"
const GetCommentFilterKeywords string = "accounts/get_comment_filter_keywords/"
const SetCommentFilterKeywords string = "accounts/set_comment_filter_keywords/"
const ChangePassword string = "accounts/change_password/"
const GetSecurityInfo string = "accounts/account_security_info/"
const SendTwoFactorEnableSms string = "accounts/send_two_factor_enable_sms/"
const EnableTwoFactorSms string = "accounts/enable_sms_two_factor/"
const DisableTwoFactorSms string = "accounts/disable_sms_two_factor/"
const GetPresenceStatus string = "accounts/get_presence_disabled/"
const EnablePresence string = "accounts/set_presence_disabled/"
const DisablePresence string = "accounts/set_presence_disabled/"
const SendConfirmEmail string = "accounts/send_confirm_email/"
const SendSmsCode string = "accounts/send_sms_code/"
const VerifySmsCode string = "accounts/verify_sms_code/"
const SetContactPointPrefil string = "accounts/contact_point_prefill/"
const GetBadgeNotifications string = "notifications/badge/"
const GetProcessContactPointSignals string = "accounts/process_contact_point_signals/"

// Business
const GetInsights string = "insights/account_organic_insights/"
const GetMediaInsights string = "insights/media_organic_insights/%s/"
const GetStatistics string = "ads/graphql/"

// Collection
const GetCollectionList string = "collections/list/"
const GetCollectionFeed string = "feed/collection/%s/"
const CreateCollection string = "collections/create/"
const DeleteCollection string = "collections/%s/delete/"
const EditCollection string = "collections/%s/edit/"
const RemoveMediaFromCollections string = "media/{$mediaId}/save/"

// Timeline
const TimelineFeed string = "feed/timeline/"
const GetUserTimelineFeed string = "feed/user/%d/"

// Story
const ReelsTrayFeed string = "feed/reels_tray/"
const GetUserReelMediaFeed string = "feed/user/%d/reel_media/"
const GetUserStoryFeed string = "feed/user/%d/story/"
const GetReelsMediaFeed string = "feed/reels_media/"

// Discover
const SuggestedSearches string = "fbsearch/suggested_searches/"
const RecentSearches string = "fbsearch/recent_searches/"
const Explore string = "discover/explore/"

// Push Notification
const PushRegister string = "push/register/"

// Direct
const DirectRankedRecipients string = "direct_v2/ranked_recipients/"
const DirectInbox string = "direct_v2/inbox/"
const Presence string = "direct_v2/get_presence/"

// People
const ActivityNews string = "news/inbox/"
const BootstrapUsers string = "scores/bootstrap/users/"
const GetInfoById string = "users/%d/info/"
const GetInfoByUsername string = "users/%s/usernameinfo/"
const GetFriendship string = "friendships/show/%d/"
const FollowUser string = "friendships/create/%d/"
const Followers string = "friendships/%d/followers/"

// Media
const BlockedMedia string = "media/blocked/"
const GetMediaInfo string = "media/%s/info/"
const CommentMedia string = "media/%s/comment/"
const CommentLike string = "media/%d/comment_like/"
const GetComments string = "media/%d/comments/"

// Highlight
const GetHighlightsUserFeed string = "highlights/%d/highlights_tray/"

// Tv
const GetIGTVChannel string = "igtv/channel/"
