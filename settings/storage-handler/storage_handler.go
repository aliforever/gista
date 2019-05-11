package storage_handler

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/natefinch/atomic"

	"github.com/aliforever/gista/utils"

	"github.com/aliforever/gista/errors"
)

type StorageHandler struct {
	storage         StorageInterface
	username        string
	userSettings    *map[string]string
	cookiesFilePath *string
}

func NewStorageHandler(storageInstance StorageInterface, locationConfig map[string]string) (sh *StorageHandler, err error) {
	sh = &StorageHandler{}
	sh.storage = storageInstance
	err = sh.storage.OpenLocation(locationConfig)
	return
}

func (sh *StorageHandler) PersistentKeys() []string {
	return []string{
		"account_id",        // The numerical UserPK ID of the account.
		"devicestring",      // Which Android device they"re identifying as.
		"device_id",         // Hardware identifier.
		"phone_id",          // Hardware identifier.
		"uuid",              // Universally unique identifier.
		"advertising_id",    // Google Play advertising ID.
		"session_id",        // The user"s current application session ID.
		"experiments",       // Interesting experiment variables for this account.
		"fbns_auth",         // Serialized auth credentials for FBNS.
		"fbns_token",        // Serialized FBNS token.
		"last_fbns_token",   // Tracks time elapsed since our last FBNS token refresh.
		"last_login",        // Tracks time elapsed since our last login state refresh.
		"last_experiments",  // Tracks time elapsed since our last experiments refresh.
		"datacenter",        // Preferred data center (region-based).
		"presence_disabled", // Whether the presence feature has been disabled by user.
		"zr_token",          // Zero rating token.
		"zr_expires",        // Zero rating token expiration timestamp.
		"zr_rules",          // Zero rating rewrite rules.
	}
}

func (sh *StorageHandler) hasPersistentKey(key string) (has bool) {
	for _, v := range sh.PersistentKeys() {
		if v == key {
			return true
		}
	}
	return
}

func (sh *StorageHandler) hasKeepKeysWhenErasingDevice(key string) (has bool) {
	for _, v := range sh.KeepKeysWhenErasingDevice() {
		if v == key {
			return true
		}
	}
	return
}

func (sh *StorageHandler) KeepKeysWhenErasingDevice() []string {
	return []string{
		"account_id",
	}
}

func (sh *StorageHandler) ExperimentKeys() []string {
	return []string{
		"ig_android_2fac",
		"ig_android_realtime_iris",
		"ig_android_skywalker_live_event_start_end",
		"ig_android_gqls_typing_indicator",
		"ig_android_upload_reliability_universe",
		"ig_android_photo_fbupload_universe",
		"ig_android_video_segmented_upload_universe",
		"ig_android_direct_video_segmented_upload_universe",
		"ig_android_reel_raven_video_segmented_upload_universe",
		"ig_android_ad_async_ads_universe",
		"ig_android_direct_inbox_presence",
		"ig_android_direct_thread_presence",
		"ig_android_rtc_reshare",
		"ig_android_sidecar_photo_fbupload_universe",
		"ig_android_fbupload_sidecar_video_universe",
		"ig_android_skip_get_fbupload_photo_universe",
		"ig_android_skip_get_fbupload_universe",
		"ig_android_loom_universe",
	}
}

func (sh *StorageHandler) SetActiveUser(username string) (err error) {
	if username == "" {
		return errors.EmptyParameter("username")
	}
	if username == sh.username {
		return
	}
	if sh.username != "" {
		_ = sh.storage.CloseUser()
	}
	sh.username = username
	sh.userSettings = &map[string]string{}
	err = sh.storage.OpenUser(username)
	if err != nil {
		return
	}
	loadedSettings, err := sh.storage.LoadUserSettings()
	if err != nil {
		return
	}

	for k, v := range loadedSettings {
		var key string
		key = k
		if k == "username_id" {
			key = "account_id"
		} else if k == "adid" {
			key = "advertising_id"
		}
		if sh.hasPersistentKey(key) {
			(*sh.userSettings)[key] = v
		}
	}
	cookiesFilePath := sh.storage.GetUserCookiesFilePath()
	sh.cookiesFilePath = &cookiesFilePath
	return
}

func (sh *StorageHandler) Get(key string) (val string, err error) {
	if sh.username == "" {
		err = errors.StorageNoUsernameIsSet
		return
	}
	if !sh.hasPersistentKey(key) {
		err = errors.NotValidPersistentKey(key)
		return
	}

	if value, ok := (*sh.userSettings)[key]; ok {
		return value, nil
	} else {
		err = errors.SettingsKeyNotFound(key)
		return
	}
}

func (sh *StorageHandler) Set(key, val string) (err error) {
	if sh.username == "" {
		err = errors.StorageNoUsernameIsSet
		return
	}
	if !sh.hasPersistentKey(key) {
		err = errors.NotValidPersistentKey(key)
		return
	}
	if _, ok := (*sh.userSettings)[key]; !ok || (*sh.userSettings)[key] != val {
		(*sh.userSettings)[key] = val
		err = sh.storage.SaveUserSettings(sh.userSettings, key)
	} /*else {
		if val == "" {
			fmt.Println("empty val")
		}
		fmt.Println("not saving", val, key, sh.userSettings)
		fmt.Println()
	}*/
	return
}

func (sh *StorageHandler) EraseDeviceSettings() error {
	for _, k := range sh.PersistentKeys() {
		if !sh.hasKeepKeysWhenErasingDevice(k) {
			err := sh.Set(k, "")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (sh *StorageHandler) IsMaybeLoggedIn() (bl bool) {
	accountId, _ := sh.Get("account_id")
	bl = sh.storage.HasUserCookies() && accountId != "" && sh.username != ""
	return
}

func (sh *StorageHandler) SetCookies(rawData *string) (err error) {
	if sh.username == "" {
		err = errors.StorageNoUsernameIsSet
		return
	}

	if rawData == nil {
		err = errors.ParameterMustBeString("cookie_raw_data")
		return
	}

	if sh.cookiesFilePath == nil {
		sh.storage.SaveUserCookies(*rawData)
	} else { // Cookie file on disk
		if len(*rawData) > 0 {
			sh.createCookiesFileDirectory()
			timeOut := 5
			t := time.Now().Unix()
			written := atomic.WriteFile(*sh.cookiesFilePath, strings.NewReader(*rawData))
			for written != nil {
				time.Sleep(time.Microsecond * time.Duration(utils.MtRand(400000, 600000)))
				if time.Now().Unix()-t > int64(timeOut) {
					break
				}
				written = atomic.WriteFile(*sh.cookiesFilePath, strings.NewReader(*rawData))
			}
			if written != nil {
				err = errors.CookiesFileNotWritable(*sh.cookiesFilePath, written.Error())
				return
			}
		} else {
			if utils.FileOrFolderExists(*sh.cookiesFilePath) {
				err = os.Remove(*sh.cookiesFilePath)
				if err != nil {
					err = errors.CannotDeleteCookiesFile(*sh.cookiesFilePath, err.Error())
					return
				}
			}
		}
	}
	return
}

func (sh *StorageHandler) GetCookies() (cookies *string, err error) {
	if sh.username == "" {
		err = errors.StorageNoUsernameIsSet
		return
	}
	cookies = nil

	if sh.cookiesFilePath == nil {
		cookies, err = sh.storage.LoadUserCookies()
	} else { // Cookie file on disk
		if *sh.cookiesFilePath == "" {
			err = errors.EmptyCookiesFilePath
			return
		}
		sh.createCookiesFileDirectory()
		if utils.FileOrFolderExists(*sh.cookiesFilePath) {
			data, dataErr := utils.FileGetContents(*sh.cookiesFilePath)
			if dataErr != nil {
				err = dataErr
				return
			}
			cookiesStr := string(data)
			cookies = &cookiesStr
		}
	}
	return
}

func (sh *StorageHandler) createCookiesFileDirectory() (err error) {
	if sh.cookiesFilePath == nil {
		return
	}
	cookieDir := filepath.Dir(*sh.cookiesFilePath)
	err = utils.CreateFolder(cookieDir)
	if err != nil {
		err = errors.CreateFolder(cookieDir)
	}
	return
}

func (sh *StorageHandler) packJSON(obj interface{}) (jsonString string, err error) {
	var jsonBytes []byte
	jsonBytes, err = json.Marshal(obj)
	jsonString = string(jsonBytes)
	return
}

func (sh *StorageHandler) unpackJSON(jsonString string, obj interface{}) (err error) {
	err = json.Unmarshal([]byte(jsonString), obj)
	return
}

func (sh *StorageHandler) SetExperiments(experiments map[string]map[string]string) (filtered map[string]map[string]string, err error) {
	filtered = map[string]map[string]string{}
	for _, key := range sh.ExperimentKeys() {
		if _, ok := experiments[key]; !ok {
			continue
		}
		filtered[key] = experiments[key]
	}
	jsonMap, _ := sh.packJSON(filtered)
	sh.Set("experiments", jsonMap)
	return
}

func (sh *StorageHandler) SetRewriteRules(rules map[string]string) (err error) {
	var jp string
	jp, err = sh.packJSON(rules)
	if err != nil {
		return
	}
	err = sh.Set("zr_rules", jp)
	return
}

func (sh *StorageHandler) GetExperiments() (experiments map[string]map[string]string, err error) {
	experiments = map[string]map[string]string{}
	var exps string
	exps, err = sh.Get("experiments")
	if err != nil {
		return
	}
	err = sh.unpackJSON(exps, &experiments)
	return
}
