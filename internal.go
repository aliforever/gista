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
	Ig *instagram
}

func newInternal(ig *instagram) (i *internal) {
	i = &internal{Ig: ig}
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

	i.Ig.experiments, err = i.Ig.Settings.SetExperiments(experiments)
	i.Ig.Settings.Set("last_experiments", fmt.Sprintf("%d", time.Now().Unix()))
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
	i.Ig.Client.ZeroRating().Update(rules)
	err = i.Ig.Settings.SetRewriteRules(rules)
	if err != nil {
		return
	}
	err = i.Ig.Settings.Set("zr_token", tokenModel.TokenHash)
	if err != nil {
		return
	}
	err = i.Ig.Settings.Set("zr_expires", fmt.Sprintf("%d", tokenModel.ExpiresAt()))
	if err != nil {
		return
	}
	return
}

func (i *internal) GetFacebookOTA() (response *responses.FacebookOta, err error) {
	response = &responses.FacebookOta{}
	err = i.Ig.Client.Request(constants.FacebookOTA).
		AddParam("fields", constants.FacebookOtaFields).
		AddParam("custom_user_id", *i.Ig.AccountId).
		AddParam("signed_body", signatures.GenerateSignature("")+".").
		AddParam("ig_sig_key_version", constants.SigKeyVersion).
		AddParam("version_code", constants.VersionCode).
		AddParam("version_name", constants.IgVersion).
		AddParam("custom_app_id", constants.FacebookOrcaApplicationId).
		AddParam("custom_device_id", i.Ig.Uuid).
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
	err = i.Ig.Client.Request(constants.FetchQPData).
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
	err = i.Ig.Client.Request(constants.ProfileNotice).GetResponse(response)
	return
}

func (i *internal) GetLoomFetchConfig() (response *responses.LoomFetchConfig, err error) {
	response = &responses.LoomFetchConfig{}
	err = i.Ig.Client.Request(constants.LoomFetchConfig).GetResponse(response)
	return
}

func (i *internal) FetchZeroRatingToken(reason string) (response *responses.TokenResult, err error) {
	if reason == "" {
		reason = "token_expired"
	}
	response = &responses.TokenResult{}
	zrToken, _ := i.Ig.Settings.Get("zr_token")
	request := i.Ig.Client.Request(constants.ZeroRatingToken).
		SetNeedsAuth(false).
		AddParam("custom_device_id", i.Ig.Uuid).
		AddParam("device_id", i.Ig.DeviceId).
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
	err = i.Ig.Client.Request(constants.LogAttribution).
		SetNeedsAuth(false).
		AddPost("adid", i.Ig.advertisingId).
		GetResponse(response)
	return
}

func (i *internal) SendLauncherSync(preLogin bool) (response *responses.LauncherSync, err error) {
	csrfToken := i.Ig.Client.GetToken()
	token := ""
	if csrfToken != nil {
		token = *csrfToken
	}
	request := i.Ig.Client.Request(constants.LauncherSync).
		AddPost("_csrftoken", token).
		AddPost("configs", "ig_android_felix_release_players,ig_user_mismatch_soft_error,ig_android_os_version_blocking_config,ig_android_carrier_signals_killswitch,fizz_ig_android,ig_mi_block_expired_events,ig_android_killswitch_perm_direct_ssim,ig_fbns_blocked")
	if preLogin {
		request.SetNeedsAuth(false).
			AddPost("id", i.Ig.Uuid)
	} else {
		request.AddPost("id", *i.Ig.AccountId).
			AddPost("_uuid", i.Ig.Uuid).
			AddPost("_uid", *i.Ig.AccountId).
			AddPost("_csrftoken", token)
	}
	response = &responses.LauncherSync{}
	err = request.GetResponse(response)
	return
}

func (i *internal) SyncDeviceFeatures(preLogin bool) (response *responses.Sync, err error) {
	request := i.Ig.Client.Request(constants.Sync).
		AddHeader("X-DEVICE-ID", i.Ig.Uuid).
		AddPost("id", i.Ig.Uuid).
		AddPost("experiments", constants.LoginExperiments)
	if preLogin {
		request.SetNeedsAuth(false)
	} else {
		csrfToken := i.Ig.Client.GetToken()
		token := ""
		if csrfToken != nil {
			token = *csrfToken
		}
		request.AddPost("_uuid", i.Ig.Uuid).
			AddPost("_uid", *i.Ig.AccountId).
			AddPost("_csrftoken", token)
	}
	response = &responses.Sync{}
	err = request.GetResponse(response)
	return
}

func (i *internal) SyncUserFeatures() (response *responses.Sync, err error) {
	response = &responses.Sync{}
	err = i.Ig.Client.Request(constants.Sync).
		AddHeader("X-DEVICE-ID", i.Ig.Uuid).
		AddPost("id", i.Ig.Uuid).
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
	request := i.Ig.Client.Request(constants.MSISDN).
		SetNeedsAuth(false).
		AddHeader("X-DEVICE-ID", i.Ig.Uuid).
		AddPost("device_id", i.Ig.Uuid).
		AddPost("mobile_subno_usage", usage)
	if subNoKey != nil {
		request.AddPost("subno_key", *subNoKey)
	}
	response = &responses.MSISDNHeader{}
	err = request.GetResponse(response)
	return
}
