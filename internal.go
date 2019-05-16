package gista

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aliforever/gista/signatures"

	"github.com/aliforever/gista/models"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
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
		AddPost("configs", "ig_android_felix_release_players,ig_user_mismatch_soft_error,ig_android_os_version_blocking_config,ig_android_carrier_signals_killswitch,fizz_ig_android,ig_mi_block_expired_events,ig_android_killswitch_perm_direct_ssim,ig_fbns_blocked")
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
