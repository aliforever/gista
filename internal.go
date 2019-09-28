package gista

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path"
	"time"

	media2 "github.com/aliforever/gista/media"

	"github.com/aliforever/gista/models/item"

	"github.com/aliforever/gista/utils"

	"github.com/aliforever/gista/metadata"

	"github.com/go-errors/errors"

	"github.com/aliforever/gista/signatures"

	"github.com/aliforever/gista/models"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

const (
	MaxChunkRetries     = 5
	MaxResumableRetries = 15
	MaxConfigureRetries = 5
	MinChunkSize        = 204800
	MaxChunkSize        = 5242880
)

type internal struct {
	ig *Instagram
}

func newInternal(ig *Instagram) (i *internal) {
	i = &internal{ig: ig}
	return
}

func (i *internal) saveExperiments(syncResp *responses.Sync) (response *responses.TokenResult, err error) {
	experiments := map[string]map[string]string{}
	for _, experiment := range syncResp.Experiments {
		group := experiment.Name
		params := experiment.Params
		if group == nil || params == nil {
			continue
		}
		if _, ok := experiments[*group]; !ok {
			experiments[*group] = map[string]string{}
		}
		for _, p := range *params {
			paramName := p.Name
			if paramName == nil {
				continue
			}
			experiments[*group][*paramName] = p.Value
		}
	}

	i.ig.experiments, err = i.ig.settings.SetExperiments(experiments)
	i.ig.settings.Set("last_experiments", fmt.Sprintf("%d", time.Now().Unix()))
	return
}

func (i *internal) saveZeroRatingToken(tokenModel *models.Token) (response *responses.TokenResult, err error) {
	if tokenModel == nil {
		return
	}
	rules := map[string]string{}
	for _, r := range tokenModel.RewriteRules {
		rules[r.Matcher] = r.Replacer
	}
	i.ig.client.ZeroRating().Update(rules)
	err = i.ig.settings.SetRewriteRules(rules)
	if err != nil {
		return
	}
	err = i.ig.settings.Set("zr_token", tokenModel.TokenHash)
	if err != nil {
		return
	}
	err = i.ig.settings.Set("zr_expires", fmt.Sprintf("%d", tokenModel.ExpiresAt()))
	if err != nil {
		return
	}
	return
}

func (i *internal) GetFacebookOTA() (response *responses.FacebookOta, err error) {
	response = &responses.FacebookOta{}
	err = i.ig.client.Request(constants.FacebookOTA).
		AddParam("fields", constants.FacebookOtaFields).
		AddParam("custom_user_id", *i.ig.AccountId).
		AddParam("signed_body", signatures.GenerateSignature("")+".").
		AddParam("ig_sig_key_version", constants.SigKeyVersion).
		AddParam("version_code", constants.VersionCode).
		AddParam("version_name", constants.IgVersion).
		AddParam("custom_app_id", constants.FacebookOrcaApplicationId).
		AddParam("custom_device_id", i.ig.uuid).
		GetResponse(response)
	return
}

func (i *internal) GetQPFetch() (response *responses.FetchQPData, err error) {
	response = &responses.FetchQPData{}
	query := `viewer() {eligible_promotions.surface_nux_id(<surface>).external_gating_permitted_qps(<external_gating_permitted_qps>).supports_client_filters(true) {edges {priority,time_range {start,end},node {id,promotion_id,max_impressions,triggers,contextual_filters {clause_type,filters {filter_type,unknown_action,value {name,required,bool_value,int_value, string_value},extra_datas {name,required,bool_value,int_value, string_value}},clauses {clause_type,filters {filter_type,unknown_action,value {name,required,bool_value,int_value, string_value},extra_datas {name,required,bool_value,int_value, string_value}},clauses {clause_type,filters {filter_type,unknown_action,value {name,required,bool_value,int_value, string_value},extra_datas {name,required,bool_value,int_value, string_value}},clauses {clause_type,filters {filter_type,unknown_action,value {name,required,bool_value,int_value, string_value},extra_datas {name,required,bool_value,int_value, string_value}}}}}},template {name,parameters {name,required,bool_value,string_value,color_value,}},creatives {title {text},content {text},footer {text},social_context {text},primary_action{title {text},url,limit,dismiss_promotion},secondary_action{title {text},url,limit,dismiss_promotion},dismiss_action{title {text},url,limit,dismiss_promotion},image.scale(<scale>) {uri,width,height}}}}}}`
	j, _ := json.Marshal(map[string]string{
		constants.SurfaceParam[0]: query,
		constants.SurfaceParam[1]: query,
	})
	err = i.ig.client.Request(constants.FetchQPData).
		AddPost("vc_policy", "default").
		AddCSRFPost().
		AddUIdPost().
		AddUuIdPost().
		AddPost("surfaces_to_queries", string(j)).
		AddPost("version", "1").
		AddPost("scale", "2").
		GetResponse(response)
	return
}

func (i *internal) GetProfileNotice() (response *responses.ProfileNotice, err error) {
	response = &responses.ProfileNotice{}
	err = i.ig.client.Request(constants.ProfileNotice).GetResponse(response)
	return
}

func (i *internal) GetLoomFetchConfig() (response *responses.LoomFetchConfig, err error) {
	response = &responses.LoomFetchConfig{}
	err = i.ig.client.Request(constants.LoomFetchConfig).GetResponse(response)
	return
}

func (i *internal) FetchZeroRatingToken(reason string) (response *responses.TokenResult, err error) {
	if reason == "" {
		reason = "token_expired"
	}
	response = &responses.TokenResult{}
	zrToken, _ := i.ig.settings.Get("zr_token")
	request := i.ig.client.Request(constants.ZeroRatingToken).
		SetNeedsAuth(false).
		AddParam("custom_device_id", i.ig.uuid).
		AddParam("device_id", i.ig.deviceId).
		AddParam("fetch_reason", reason).
		AddParam("token_hash", zrToken)
	err = request.GetResponse(response)
	if err == nil {
		// Save Zero Rating Token
	}
	//pretty.Println(response.GetHTTPResponse().request.Form)

	return
}

func (i *internal) LogAttribution() (response *responses.Generic, err error) {
	response = &responses.Generic{}
	err = i.ig.client.Request(constants.LogAttribution).
		SetNeedsAuth(false).
		AddPost("adid", i.ig.advertisingId).
		GetResponse(response)
	return
}

func (i *internal) SendLauncherSync(preLogin bool) (response *responses.LauncherSync, err error) {
	csrfToken := i.ig.client.GetToken()
	token := ""
	if csrfToken != nil {
		token = *csrfToken
	}
	request := i.ig.client.Request(constants.LauncherSync).
		AddPost("_csrftoken", token).
		AddPost("configs", constants.LauncherConfigs)
	if preLogin {
		request.SetNeedsAuth(false).
			AddPost("id", i.ig.uuid)
	} else {
		request.AddPost("id", *i.ig.AccountId).
			AddPost("_uuid", i.ig.uuid).
			AddPost("_uid", *i.ig.AccountId).
			AddPost("_csrftoken", token)
	}
	response = &responses.LauncherSync{}
	err = request.GetResponse(response)
	return
}

func (i *internal) SyncDeviceFeatures(preLogin bool) (response *responses.Sync, err error) {
	request := i.ig.client.Request(constants.Sync).
		AddHeader("X-DEVICE-ID", i.ig.uuid).
		AddPost("id", i.ig.uuid).
		AddPost("experiments", constants.LoginExperiments)
	if preLogin {
		request.SetNeedsAuth(false)
	} else {
		csrfToken := i.ig.client.GetToken()
		token := ""
		if csrfToken != nil {
			token = *csrfToken
		}
		request.AddPost("_uuid", i.ig.uuid).
			AddPost("_uid", *i.ig.AccountId).
			AddPost("_csrftoken", token)
	}
	response = &responses.Sync{}
	err = request.GetResponse(response)
	return
}

func (i *internal) SyncUserFeatures() (response *responses.Sync, err error) {
	response = &responses.Sync{}
	err = i.ig.client.Request(constants.Sync).
		AddHeader("X-DEVICE-ID", i.ig.uuid).
		AddPost("id", i.ig.uuid).
		AddPost("experiments", constants.LoginExperiments).
		AddUuIdPost().
		AddUIdPost().
		AddCSRFParam().
		AddIdPost().
		AddPost("experiments", constants.Experiments).
		GetResponse(response)
	i.saveExperiments(response)
	return
}

func (i *internal) ReadMsisdnHeader(usage string, subNoKey *string) (response *responses.MSISDNHeader, err error) {
	request := i.ig.client.Request(constants.MSISDN).
		SetNeedsAuth(false).
		AddHeader("X-DEVICE-ID", i.ig.uuid).
		AddPost("device_id", i.ig.uuid).
		AddPost("mobile_subno_usage", usage)
	if subNoKey != nil {
		request.AddPost("subno_key", *subNoKey)
	}
	response = &responses.MSISDNHeader{}
	err = request.GetResponse(response)
	return
}

func (i *internal) UploadSinglePhoto(targetFeed string, photoFileName string, internalMetaData *metadata.Internal, externalMetaData interface{}) (response *responses.MSISDNHeader, err error) {
	if targetFeed != constants.FeedTimeline && targetFeed != constants.FeedStory && targetFeed != constants.FeedDirectStory {
		err = errors.New(fmt.Sprintf("Bad target feed %s", targetFeed))
		return
	}
	if internalMetaData == nil {
		uploadId := utils.GenerateUploadId(true)
		internalMetaData = metadata.NewInternalMetaData(&uploadId)
	}
	if internalMetaData.GetPhotoDetails() == nil {
		_, err = internalMetaData.SetPhotoDetails(targetFeed, photoFileName)
		if err != nil {
			err = errors.New(fmt.Sprintf("Failed to get photo details: %s", err.Error()))
			return
		}
	}
	i.uploadPhotoData(targetFeed, *internalMetaData)

	i.configureSinglePhoto(targetFeed, *internalMetaData, externalMetaData)
	return
}

func (i *internal) uploadPhotoData(targetFeed string, metaData metadata.Internal) (res interface{}, err error) {
	endPoint := ""
	switch targetFeed {
	case constants.FeedTimeline:
		endPoint = "media/configure"
	case constants.FeedDirect, constants.FeedDirectStory:
		endPoint = "configure_to_story"
	default:
		err = errors.New(fmt.Sprintf("bad target feed %s", targetFeed) + endPoint)
		return
	}
	return
}

func (i *internal) configureSinglePhoto(targetFeed string, metaData metadata.Internal, externalMetaData interface{}) (res interface{}, err error) {
	if targetFeed == constants.FeedDirect {
		err = errors.New(fmt.Sprintf(`Bad target feed "%s".`, targetFeed))
		return
	}
	if metaData.GetPhotoDetails() == nil {
		err = errors.New(`Photo details are missing from the internal metadata.`)
		return
	}
	if i.useResumablePhotoUploader(targetFeed, metaData) {
		res, err = i.uploadResumablePhoto(targetFeed, metaData)
	} else {
		res, err = i.uploadPhotoInOnePiece(targetFeed, metaData)
	}
	return
}

func (i *internal) useResumablePhotoUploader(targetFeed string, metaData metadata.Internal) (result bool) {
	switch targetFeed {
	case constants.FeedTimelineAlbum:
		result = i.ig.isExperimentEnabled("ig_android_sidecar_photo_fbupload_universe", "is_enabled_fbupload_sidecar_photo", false)
		return
	default:
		result = i.ig.isExperimentEnabled("ig_android_photo_fbupload_universe", "is_enabled_fbupload_photo", false)
		return
	}
}

func (i *internal) getPhotoUploadParams(targetFeed string, metaData metadata.Internal) (result map[string]string) {
	retryContext := i.getRetryContext()
	ji, _ := json.Marshal(retryContext)
	mediaType := item.Photo
	if metaData.GetVideoDetails() != nil {
		mediaType = item.Video
	}
	result = map[string]string{
		"upload_id":         metaData.GetUploadId(),
		"retry_context":     string(ji),
		"image_compression": `{"lib_name":"moz","lib_version":"3.1.m","quality":"87"}`,
		"xsharing_user_ids": `[]`,
		"media_type":        fmt.Sprintf("%d", mediaType),
	}
	switch targetFeed {
	case constants.FeedTimelineAlbum:
		result["is_sidecar"] = "1"
		break
	}
	return
}

func (i internal) getRetryContext() map[string]int {
	return map[string]int{
		"num_step_auto_retry":   0,
		"num_reupload":          0,
		"num_step_manual_retry": 0,
	}
}

func (i *internal) uploadResumablePhoto(targetFeed string, metaData metadata.Internal) (res *responses.ResumableUpload, err error) {
	pd := metaData.GetPhotoDetails()
	endPoint := fmt.Sprintf("https://i.instagram.com/rupload_igphoto/%s_%d_%d", metaData.GetUploadId(), 0, utils.HashCode(pd.GetFileName()))
	uploadParams := i.getPhotoUploadParams(targetFeed, metaData)
	j, _ := json.Marshal(uploadParams)
	offsetTemplate := i.ig.client.Request(endPoint)
	offsetTemplate.
		SetAddDefaultHeaders(false).
		AddHeader("X_FB_PHOTO_WATERFALL_ID", signatures.GenerateUUID(true)).
		AddHeader("X-Instagram-Rupload-Params", string(j))
	u, _ := url.Parse(endPoint)
	uploadTemplate := offsetTemplate
	uploadTemplate.
		AddHeader("X-Entity-Type", "image/jpeg").
		AddHeader("X-Entity-Name", path.Base(u.Path)).
		AddHeader("X-Entity-Length", fmt.Sprintf("%d", pd.GetFileSize()))
	res, err = i.uploadResumableMedia(pd, offsetTemplate, uploadTemplate, i.ig.isExperimentEnabled("ig_android_skip_get_fbupload_photo_universe", "photo_skip_get", false))
	return
}

func (i *internal) uploadResumableMedia(md media2.Details, offsetTemplate, uploadTemplate *request, skipGet bool) (res *responses.ResumableUpload, err error) {
	var f *os.File
	f, err = os.Open(md.GetFileName())
	if err != nil {
		err = errors.New(fmt.Sprintf(`Failed to open media file for reading. %s`, err.Error()))
		return
	}
	defer f.Close()
	//fileSize := md.GetFileSize()
	attempt := 0
	var offset int
	for true {
		attempt += 1
		if attempt > MaxResumableRetries {
			err = errors.New("All retries have failed.")
			return
		}
		if attempt == 1 && skipGet {
			offset = 0
		} else {
			offsetRequest := offsetTemplate
			res := &responses.ResumableOffset{}
			err = offsetRequest.GetResponse(res)
			if err != nil {
				return
			}
			oss := res.Offset
			offset = *oss
		}
		res = &responses.ResumableUpload{}
		uploadRequest := uploadTemplate
		err = uploadRequest.AddHeader("Offset", fmt.Sprintf("%d", offset)).
			SetBody(f).GetResponse(res)
		if err != nil {
			return
		}
	}
	err = errors.New("Something went wrong during media upload.")
	return
}

func (i *internal) uploadPhotoInOnePiece(targetFeed string, metaData metadata.Internal) (res *responses.UploadPhoto, err error) {
	res = &responses.UploadPhoto{}
	fileName := "pending_media_" + utils.GenerateUploadId(false) + ".jpg"
	var req *request
	req, err = i.ig.client.Request(constants.UploadPhoto).SetSignedPost(false).AddUuIdPost().AddCSRFPost().AddFile("photo", metaData.GetPhotoDetails().GetFileName(), &fileName, nil)
	if err != nil {
		return
	}
	for k, v := range i.getPhotoUploadParams(targetFeed, metaData) {
		req.AddPost(k, v)
	}
	err = req.GetResponse(res)
	return
}
