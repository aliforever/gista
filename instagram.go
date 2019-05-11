package gista

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/aliforever/gista/utils"

	"github.com/aliforever/gista/responses"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/devices"
	"github.com/aliforever/gista/errors"
	"github.com/aliforever/gista/settings/factory"
	storage_handler "github.com/aliforever/gista/settings/storage-handler"
	"github.com/aliforever/gista/signatures"
)

const experimentsRefresh int = 7200

type instagram struct {
	device               devices.DeviceInterface
	Username             string
	Password             string
	Settings             *storage_handler.StorageHandler
	client               *client
	Uuid                 string
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
	Internal *internal
	Account  *account
	Timeline *timeline
	Story    *story
	Discover *discover
	Push     *push
	Direct   *direct
	People   *people
	Media    *media
}

func New(storageConfig *map[string]string) (i *instagram, err error) {
	rand.Seed(time.Now().UTC().UnixNano())
	i = &instagram{}
	/*i.device = devices.NewDevice(constants.IgVersion, constants.VersionCode, constants.UserAgentLocale, "")
	if i.device.GetDeviceString() == "" {
		i.device.SetDeviceString(good_devices.GetRandomGoodDevice())
	}*/
	i.Settings, err = factory.CreateHandler(storageConfig)
	if err != nil {
		return
	}
	i.client = newClient(i)
	i.Internal = newInternal(i)
	i.Account = newAccount(i)
	i.Timeline = newTimeline(i)
	i.Story = newStory(i)
	i.Discover = newDiscover(i)
	i.Push = newPush(i)
	i.Direct = newDirect(i)
	i.People = newPeople(i)
	i.Media = newMedia(i)
	return
}

func (i *instagram) isExperimentEnabled(experiment, param string, defaultVal bool /*false*/) bool {
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

func (i *instagram) setUser(username, password string) (err error) {
	if username == "" || password == "" {
		return errors.EmptyUsernameOrPassword
	}
	err = i.Settings.SetActiveUser(username)
	if err != nil {
		return
	}
	savedDeviceString, err := i.Settings.Get("devicestring")

	var dstr *string = nil
	if err == nil && savedDeviceString != "" {
		dstr = &savedDeviceString
	}
	i.device = devices.NewDevice(constants.IgVersion, constants.VersionCode, constants.UserAgentLocale, dstr, true)
	deviceString := i.device.GetDeviceString()
	uuId, _ := i.Settings.Get("uuid")
	phoneId, _ := i.Settings.Get("phone_id")
	deviceId, _ := i.Settings.Get("device_id")
	resetCookieJar := false

	if deviceString != savedDeviceString || uuId == "" || phoneId == "" || deviceId == "" {
		i.Settings.EraseDeviceSettings()
		err = i.Settings.Set("devicestring", deviceString)
		if err != nil {
			return
		}
		err = i.Settings.Set("device_id", signatures.GenerateDeviceId())
		if err != nil {
			return
		}
		err = i.Settings.Set("phone_id", signatures.GenerateUUID(true))
		if err != nil {
			return
		}
		err = i.Settings.Set("uuid", signatures.GenerateUUID(true))
		if err != nil {
			return
		}
		err = i.Settings.Set("account_id", "")
		if err != nil {
			return
		}
		resetCookieJar = true
	}
	advId, _ := i.Settings.Get("advertising_id")
	if advId == "" {
		err = i.Settings.Set("advertising_id", signatures.GenerateUUID(true))
		if err != nil {
			return
		}
	}
	sessionId, _ := i.Settings.Get("session_id")
	if sessionId == "" {
		err = i.Settings.Set("session_id", signatures.GenerateUUID(true))
		if err != nil {
			return
		}
	}
	i.Username = username
	i.Password = password
	i.Uuid, _ = i.Settings.Get("uuid")
	i.advertisingId, _ = i.Settings.Get("advertising_id")
	i.deviceId, _ = i.Settings.Get("device_id")
	i.phoneId, _ = i.Settings.Get("phone_id")
	i.sessionId, _ = i.Settings.Get("session_id")
	/*i.experiments, _ = i.Settings.Get("Uuid")*/
	if !resetCookieJar && i.Settings.IsMaybeLoggedIn() {
		i.isMaybeLoggedIn = true
		accId, _ := i.Settings.Get("account_id")
		i.AccountId = &accId
	} else {
		i.isMaybeLoggedIn = false
		i.AccountId = nil
	}
	i.client.UpdateFromCurrentSettings(resetCookieJar)
	return nil
}

func (i *instagram) login(user, password string, appRefreshInterval int /*1800*/, forceLogin bool) (err error) {
	if user == "" || password == "" {
		return errors.EmptyUsernameOrPassword
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
		loginResponse, err = i.Account.Login(user, password)
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

func (i *instagram) Login(user, password string, forceLogin bool) (err error) {
	if user == "" || password == "" {
		return errors.EmptyUsernameOrPassword
	}
	err = i.login(user, password, 1800, false)
	//constants.ApiUrls
	return nil
}

func (i *instagram) SetAddHTTPResponseToResult(status bool) {
	i.httpResponseInResult = status
}

func (i *instagram) SetAddRawResponseToResult(status bool) {
	i.rawResponseInResult = status
}

func (i *instagram) sendLoginFlow(justLoggedIn bool, appRefreshInterval int /*1800*/) (err error) {
	if appRefreshInterval < 0 {
		err = errors.InvalidAppRefreshInterval(appRefreshInterval)
		return
	}
	if appRefreshInterval > 21600 {
		err = errors.TooHighAppRefreshInterval(appRefreshInterval)
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
		lastLoginTime, _ := i.Settings.Get("last_login")
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
			if err == errors.NotLoggedIn {
				err = i.login(i.Username, i.Password, appRefreshInterval, true)
				return
			}
		}
		if isSessionExpired {
			i.Settings.Set("last_login", fmt.Sprintf("%d", time.Now().Unix()))
			i.sessionId = signatures.GenerateUUID(true)
			i.Settings.Set("session_id", i.sessionId)
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
		lastExperimentsTime, exErr := i.Settings.Get("last_experiments")
		lastExperimentsTimeInt, _ := strconv.Atoi(lastExperimentsTime)
		if exErr != nil || time.Now().Unix()-int64(lastExperimentsTimeInt) > int64(experimentsRefresh) {
			i.Internal.SyncUserFeatures()
			i.Internal.SyncDeviceFeatures(false)
		}
		zrExpires, _ := i.Settings.Get("zr_expires")
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

func (i *instagram) getExperimentParam(experiment, param string, def string) string {
	if _, ok := i.experiments[experiment]; ok {
		if _, ok := i.experiments[experiment][param]; ok {
			return i.experiments[experiment][param]
		}
	}
	return def
}

func (i *instagram) registerPushChannels() (err error) {
	var (
		lastFbnsTokenInt int
		fbnsToken        string
	)
	lastFbnsToken, tErr := i.Settings.Get("last_fbns_token")
	if tErr != nil {
		lastFbnsTokenInt = 0
	} else {
		lastFbnsTokenInt, _ = strconv.Atoi(lastFbnsToken)
	}
	if lastFbnsTokenInt != 0 || int64(lastFbnsTokenInt) < time.Now().Add(-(time.Hour*time.Duration(24))).Unix() {
		err = i.Settings.Set("fbns_token", "")
	}
	// Ignore Storage Errors
	err = nil
	fbnsToken, err = i.Settings.Get("last_fbns_token")
	if err != nil {
		return
	}
	_, err = i.Push.Register("mqtt", fbnsToken)
	if err != nil {
		err = i.Settings.Set("fbns_token", "")
	}
	// Ignore Storage Errors
	err = nil
	return
}

func (i *instagram) sendPreLoginFlow() {
	i.SetAddRawResponseToResult(true)
	i.Internal.ReadMsisdnHeader("ig_select_app", nil)
	i.Internal.SyncDeviceFeatures(true)
	i.Internal.SendLauncherSync(true)
	i.Internal.LogAttribution()
	i.Internal.FetchZeroRatingToken("token_expired")
	i.Account.SetContactPointPreFill("prefill")
}

/*func (i *instagram) request(address string) *request {
	return NewRequest(i, address)
}*/
func (i *instagram) updateLoginState(loginResponse *responses.Login) (err error) {
	if !loginResponse.IsOk() {
		m, _ := loginResponse.GetMessage()
		err = errors.InvalidLoginResponse(m)
		return
	}
	i.isMaybeLoggedIn = true
	stringId := fmt.Sprintf("%d", loginResponse.LoggedInUser.Pk)
	i.AccountId = &stringId
	i.Settings.Set("account_id", stringId)
	i.Settings.Set("last_login", fmt.Sprintf("%d", time.Now().Unix()))
	return
}
