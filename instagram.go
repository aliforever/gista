package gista

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/aliforever/gista/errs"

	"github.com/aliforever/gista/settings/factory"

	"github.com/aliforever/gista/utils"

	"github.com/aliforever/gista/responses"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/devices"
	storage_handler "github.com/aliforever/gista/settings/storage-handler"
	"github.com/aliforever/gista/signatures"
)

const experimentsRefresh int = 7200

type Instagram struct {
	device               devices.DeviceInterface
	Username             string
	Password             string
	settings             *storage_handler.StorageHandler
	client               *client
	uuid                 string
	advertisingId        string
	deviceId             string
	AccountId            *string
	phoneId              string
	sessionId            string
	experiments          map[string]map[string]string
	isMaybeLoggedIn      bool
	httpResponseInResult bool
	rawResponseInResult  bool
	//Properties
	Account    *account
	Business   *business
	Collection *collection
	Creative   *creative
	Direct     *direct
	Hashtag    *hashtag
	Highlight  *highlight
	Tv         *tv
	Internal   *internal
	Live       *live
	Location   *location
	Media      *media
	People     *people
	Push       *push
	Shopping   *shopping
	Story      *story
	Timeline   *timeline
	Usertag    *usertag
	Discover   *discover
}

func New(storageConfig *map[string]string) (i *Instagram, err error) {
	rand.Seed(time.Now().UTC().UnixNano())
	i = &Instagram{}
	i.Account = newAccount(i)
	i.Business = newBusiness(i)
	i.Collection = newCollection(i)
	i.Creative = newCreative(i)
	i.Direct = newDirect(i)
	i.Discover = newDiscover(i)
	i.Hashtag = newHashtag(i)
	i.Highlight = newHighlight(i)
	i.Tv = newTv(i)
	i.Internal = newInternal(i)
	i.Live = newLive(i)
	i.Location = newLocation(i)
	i.Media = newMedia(i)
	i.People = newPeople(i)
	i.Push = newPush(i)
	i.Shopping = newShopping(i)
	i.Story = newStory(i)
	i.Timeline = newTimeline(i)
	i.Usertag = newUsertag(i)
	i.client = newClient(i)
	i.settings, err = factory.CreateHandler(storageConfig)
	if err != nil {
		return
	}
	i.experiments = map[string]map[string]string{}
	return
}

func (i *Instagram) isExperimentEnabled(experiment, param string, defaultVal bool /*false*/) bool {
	if i.experiments == nil {
		return false
	}
	if _, ok := i.experiments[experiment]; ok {
		if _, ok := i.experiments[experiment][param]; ok {
			for _, item := range []string{"enabled", "true", "1"} {
				if i.experiments[experiment][param] == item {
					return true
				}
			}
		}
	}
	return defaultVal
}

func (i *Instagram) setUser(username, password string) (err error) {
	if username == "" || password == "" {
		return errs.EmptyUsernameOrPassword
	}
	err = i.settings.SetActiveUser(username)
	if err != nil {
		return
	}
	savedDeviceString, err := i.settings.Get("devicestring")

	var dstr *string = nil
	if err == nil && savedDeviceString != "" {
		dstr = &savedDeviceString
	}
	i.device = devices.NewDevice(constants.IgVersion, constants.VersionCode, constants.UserAgentLocale, dstr, true)
	deviceString := i.device.GetDeviceString()
	uuId, _ := i.settings.Get("uuid")
	phoneId, _ := i.settings.Get("phone_id")
	deviceId, _ := i.settings.Get("device_id")
	resetCookieJar := false

	if deviceString != savedDeviceString || uuId == "" || phoneId == "" || deviceId == "" {
		i.settings.EraseDeviceSettings()
		err = i.settings.Set("devicestring", deviceString)
		if err != nil {
			return
		}
		err = i.settings.Set("device_id", signatures.GenerateDeviceId())
		if err != nil {
			return
		}
		err = i.settings.Set("phone_id", signatures.GenerateUUID(true))
		if err != nil {
			return
		}
		err = i.settings.Set("uuid", signatures.GenerateUUID(true))
		if err != nil {
			return
		}
		err = i.settings.Set("account_id", "")
		if err != nil {
			return
		}
		resetCookieJar = true
	}
	advId, _ := i.settings.Get("advertising_id")
	if advId == "" {
		err = i.settings.Set("advertising_id", signatures.GenerateUUID(true))
		if err != nil {
			return
		}
	}
	sessionId, _ := i.settings.Get("session_id")
	if sessionId == "" {
		err = i.settings.Set("session_id", signatures.GenerateUUID(true))
		if err != nil {
			return
		}
	}
	i.Username = username
	i.Password = password
	i.uuid, _ = i.settings.Get("uuid")
	i.advertisingId, _ = i.settings.Get("advertising_id")
	i.deviceId, _ = i.settings.Get("device_id")
	i.phoneId, _ = i.settings.Get("phone_id")
	i.sessionId, _ = i.settings.Get("session_id")
	/*i.experiments, _ = i.settings.Get("uuid")*/
	if !resetCookieJar && i.settings.IsMaybeLoggedIn() {
		i.isMaybeLoggedIn = true
		accId, _ := i.settings.Get("account_id")
		i.AccountId = &accId
	} else {
		i.isMaybeLoggedIn = false
		i.AccountId = nil
	}
	i.client.UpdateFromCurrentSettings(resetCookieJar)
	return nil
}

func (i *Instagram) FinishTwoFactorLogin(username, password, twoFactorIdentifier, verificationCode string) (err error) {
	_, err = i.Account.finishTwoFactorLogin(username, password, twoFactorIdentifier, verificationCode, "1", 1800, nil)
	return
}

func (i *Instagram) login(user, password string, appRefreshInterval int /*1800*/, forceLogin bool) (err error) {
	if user == "" || password == "" {
		return errs.EmptyUsernameOrPassword
	}
	if i.Username != user || i.Password != password {
		err = i.setUser(user, password)
		if err != nil {
			return
		}
	}
	if !i.isMaybeLoggedIn || forceLogin {
		i.sendPreLoginFlow()
		var loginResponse *responses.Login
		//i.SetAddHTTPResponseToResult(true)
		loginResponse, err = i.Account.login(user, password)
		if err != nil {
			return
		}
		//pretty.Println("here bro", loginResponse.GetHTTPResponse().Cookies())
		err = i.updateLoginState(loginResponse)
		if err != nil {
			return
		}
		i.sendLoginFlow(true, appRefreshInterval)
	}
	return
}

func (i *Instagram) Login(user, password string, forceLogin bool) (err error) {
	if user == "" || password == "" {
		return errs.EmptyUsernameOrPassword
	}
	err = i.login(user, password, 1800, false)
	//constants.ApiUrls
	return err
}

func (i *Instagram) SetAddHTTPResponseToResult(status bool) {
	i.httpResponseInResult = status
}

func (i *Instagram) SetAddRawResponseToResult(status bool) {
	i.rawResponseInResult = status
}

func (i *Instagram) sendLoginFlow(justLoggedIn bool, appRefreshInterval int /*1800*/) (err error) {
	if appRefreshInterval < 0 {
		err = errs.InvalidAppRefreshInterval(appRefreshInterval)
		return
	}
	if appRefreshInterval > 21600 {
		err = errs.TooHighAppRefreshInterval(appRefreshInterval)
		return
	}
	if justLoggedIn {
		i.Internal.SendLauncherSync(false)
		i.Internal.SyncUserFeatures()
		i.Timeline.GetTimelineFeed(nil, map[string]interface{}{"recovered_from_crash": true})
		i.Story.GetReelsTrayFeed()
		i.Discover.GetSuggestedSearches("users")
		i.Discover.GetRecentSearches()
		i.Discover.GetSuggestedSearches("blended")
		i.Internal.FetchZeroRatingToken("token_expired")
		i.registerPushChannels()
		i.Direct.GetRankedRecipients("reshare", true, nil)
		i.Direct.GetRankedRecipients("raven", true, nil)
		i.Direct.GetInbox(nil)
		i.Direct.GetPresences()
		i.People.GetRecentActivityInbox()
		experimentParam := i.getExperimentParam("ig_android_loom_universe", "cpu_sampling_rate_ms", "0")
		experimentParamInt, _ := strconv.Atoi(experimentParam)
		if experimentParamInt > 0 {
			i.Internal.GetLoomFetchConfig()
		}
		i.Internal.GetProfileNotice()
		i.Media.GetBlockedMedia()
		i.People.GetBootstrapUsers()
		i.Discover.GetExploreFeed(nil, true)
		i.Internal.GetQPFetch()
		i.Internal.GetFacebookOTA()
	} else {
		lastLoginTime, _ := i.settings.Get("last_login")
		lastLoginTimeInt, _ := strconv.Atoi(lastLoginTime)
		isSessionExpired := lastLoginTime == "" || (time.Now().Unix()-int64(lastLoginTimeInt)) > int64(appRefreshInterval)
		options := map[string]interface{}{
			"is_pull_to_refresh": nil,
		}
		if !isSessionExpired {
			options["is_pull_to_refresh"] = utils.MtRand(1, 3) < 3
		}
		_, err = i.Timeline.GetTimelineFeed(nil, options)
		if err != nil {
			if err == errs.NotLoggedIn {
				err = i.login(i.Username, i.Password, appRefreshInterval, true)
				return
			}
		}
		if isSessionExpired {
			i.settings.Set("last_login", fmt.Sprintf("%d", time.Now().Unix()))
			i.sessionId = signatures.GenerateUUID(true)
			i.settings.Set("session_id", i.sessionId)
			i.People.GetBootstrapUsers()
			i.Story.GetReelsTrayFeed()
			i.Direct.GetRankedRecipients("reshare", true, nil)
			i.Direct.GetRankedRecipients("raven", true, nil)
			i.registerPushChannels()
			i.Direct.GetInbox(nil)
			i.Direct.GetPresences()
			i.People.GetRecentActivityInbox()
			i.Internal.GetProfileNotice()
			i.Discover.GetExploreFeed(nil, false)
		}
		lastExperimentsTime, exErr := i.settings.Get("last_experiments")
		lastExperimentsTimeInt, _ := strconv.Atoi(lastExperimentsTime)
		if exErr != nil || time.Now().Unix()-int64(lastExperimentsTimeInt) > int64(experimentsRefresh) {
			i.Internal.SyncUserFeatures()
			i.Internal.SyncDeviceFeatures(false)
		}
		zrExpires, _ := i.settings.Get("zr_expires")
		zrExpiresInt, _ := strconv.Atoi(zrExpires)
		expired := time.Now().Unix() - int64(zrExpiresInt)
		reason := "token_expired"
		if expired > 7200 {
			reason = "token_stale"
		}
		if expired > 0 {
			i.client.ZeroRating().Reset()
			i.Internal.FetchZeroRatingToken(reason)
		}
	}
	err = i.client.SaveCookieJar()
	// Left Here
	return
}

func (i *Instagram) getExperimentParam(experiment, param string, def string) string {
	if _, ok := i.experiments[experiment]; ok {
		if _, ok := i.experiments[experiment][param]; ok {
			return i.experiments[experiment][param]
		}
	}
	return def
}

func (i *Instagram) registerPushChannels() (err error) {
	var (
		lastFbnsTokenInt int
		fbnsToken        string
	)
	lastFbnsToken, tErr := i.settings.Get("last_fbns_token")
	if tErr != nil {
		lastFbnsTokenInt = 0
	} else {
		lastFbnsTokenInt, _ = strconv.Atoi(lastFbnsToken)
	}
	if lastFbnsTokenInt != 0 || int64(lastFbnsTokenInt) < time.Now().Add(-(time.Hour*time.Duration(24))).Unix() {
		err = i.settings.Set("fbns_token", "")
	}
	// Ignore Storage Errors
	err = nil
	fbnsToken, err = i.settings.Get("last_fbns_token")
	if err != nil {
		return
	}
	_, err = i.Push.Register("mqtt", fbnsToken)
	if err != nil {
		err = i.settings.Set("fbns_token", "")
	}
	// Ignore Storage Errors
	err = nil
	return
}

func (i *Instagram) sendPreLoginFlow() {
	//i.SetAddRawResponseToResult(true)
	i.Internal.ReadMsisdnHeader("ig_select_app", nil)
	i.Internal.SyncDeviceFeatures(true)
	i.Internal.SendLauncherSync(true)
	i.Internal.LogAttribution()
	i.Internal.FetchZeroRatingToken("token_expired")
	i.Account.SetContactPointPreFill("prefill")
}

/*func (i *InstagramInterface) request(address string) *request {
	return NewRequest(i, address)
}*/
func (i *Instagram) updateLoginState(loginResponse *responses.Login) (err error) {
	if !loginResponse.IsOk() {
		err = errs.InvalidLoginResponse(loginResponse.GetMessage())
		return
	}
	i.isMaybeLoggedIn = true
	stringId := fmt.Sprintf("%d", loginResponse.LoggedInUser.Pk)
	i.AccountId = &stringId
	i.settings.Set("account_id", stringId)
	i.settings.Set("last_login", fmt.Sprintf("%d", time.Now().Unix()))
	return
}
