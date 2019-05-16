package gista

import (
	"fmt"
	"strings"
	"time"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
	"github.com/aliforever/gista/utils"
)

type timeline struct {
	ig *Instagram
}

func newTimeline(i *Instagram) *timeline {
	return &timeline{ig: i}
}

func (t *timeline) GetTimelineFeed(maxId *string, options map[string]interface{}) (res *responses.TimelineFeed, err error) {
	res = &responses.TimelineFeed{}
	_, offset := time.Now().Zone()
	asyncAds := t.ig.isExperimentEnabled("ig_android_ad_async_ads_universe", "is_enabled", false)
	asyncAds2 := t.ig.isExperimentEnabled("ig_android_ad_async_ads_universe", "is_async_ads_in_headload_enabled", false)
	asyncAds3 := t.ig.isExperimentEnabled("ig_android_ad_async_ads_universe", "is_double_request_enabled", false)
	asyncAds4 := t.ig.isExperimentEnabled("ig_android_ad_async_ads_universe", "is_rti_enabled", false)
	asyncAds5 := t.ig.isExperimentEnabled("ig_android_ad_async_ads_universe", "rti_delivery_backend", false)
	asyncAdsStr := "0"
	if asyncAds && asyncAds2 {
		asyncAdsStr = "1"
	}
	asyncAdsStr2 := "0"
	if asyncAds && asyncAds3 {
		asyncAdsStr2 = "1"
	}
	asyncAdsStr3 := "0"
	if asyncAds && asyncAds4 {
		asyncAdsStr3 = "1"
	}
	asyncAdsStr4 := "0"
	if asyncAds && asyncAds5 {
		asyncAdsStr4 = "1"
	}
	request := t.ig.client.Request(constants.TimelineFeed).
		SetSignedPost(false).
		SetIsBodyCompressed(true).
		AddHeader("X-Ads-Opt-Out", "0").
		AddHeader("X-Google-AD-ID", t.ig.advertisingId).
		AddHeader("X-DEVICE-ID", t.ig.uuid).
		AddCSRFPost().
		AddUuIdPost().
		AddPost("is_prefetch", "0").
		AddPhoneIdPost().
		AddDeviceIdPost().
		AddPost("client_session_id", t.ig.sessionId).
		AddPost("battery_level", fmt.Sprintf("%d", utils.MtRand(25, 100))).
		AddPost("is_charging", "0").
		AddPost("will_sound_on", "1").
		AddPost("is_on_screen", "true").
		AddPost("timezone_offset", fmt.Sprintf("%d", offset)).
		AddPost("is_async_ads_in_headload_enabled", asyncAdsStr).
		AddPost("is_async_ads_double_request", asyncAdsStr2).
		AddPost("is_async_ads_rti", asyncAdsStr3).
		AddPost("rti_delivery_backend", asyncAdsStr4)
	optionsHaveItem := func(item string) bool {
		if options != nil {
			if _, ok := options[item]; ok {
				return true
			}
		}
		return false
	}
	if optionsHaveItem("latest_story_pk") {
		request.AddPost("latest_story_pk", options["latest_story_pk"].(string))
	}
	if maxId != nil {
		request.AddPost("reason", "pagination")
		request.AddPost("max_id", *maxId)
		request.AddPost("is_pull_to_refresh", "0")
	} else if optionsHaveItem("is_pull_to_refresh") && options["is_pull_to_refresh"] != "" {
		request.AddPost("reason", "pull_to_refresh")
		request.AddPost("is_pull_to_refresh", "1")
	} else if optionsHaveItem("is_pull_to_refresh") {
		request.AddPost("reason", "warm_start_fetch")
		request.AddPost("is_pull_to_refresh", "0")
	} else {
		request.AddPost("reason", "cold_start_fetch")
		request.AddPost("is_pull_to_refresh", "0")
	}
	if optionsHaveItem("seen_posts") {
		switch options["seen_posts"].(type) {
		case []string:
			request.AddPost("seen_posts", strings.Join(options["seen_posts"].([]string), ","))
			break
		case string:
			request.AddPost("seen_posts", options["seen_posts"].(string))
			break
		}
	} else if maxId == nil {
		request.AddPost("seen_posts", "")
	}
	if optionsHaveItem("unseen_posts") {
		switch options["unseen_posts"].(type) {
		case []string:
			request.AddPost("unseen_posts", strings.Join(options["unseen_posts"].([]string), ","))
			break
		case string:
			request.AddPost("unseen_posts", options["unseen_posts"].(string))
			break
		}
	} else if maxId == nil {
		request.AddPost("unseen_posts", "")
	}
	if optionsHaveItem("feed_view_info") {
		switch options["feed_view_info"].(type) {
		case []string:
			request.AddPost("feed_view_info", strings.Join(options["feed_view_info"].([]string), ","))
			break
		case string:
			request.AddPost("feed_view_info", options["feed_view_info"].(string))
			break
		}
	} else if maxId == nil {
		request.AddPost("feed_view_info", "")
	}
	if optionsHaveItem("push_disabled") && options["push_disabled"] != "" {
		request.AddPost("push_disabled", "true")
	}
	if optionsHaveItem("recovered_from_crash") && options["recovered_from_crash"] != "" {
		request.AddPost("recovered_from_crash", "true")
	}
	err = request.GetResponse(res)
	return
}
