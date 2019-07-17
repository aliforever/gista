package storage_file

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/aliforever/gista/errs"

	"github.com/natefinch/atomic"

	storage_handler "github.com/aliforever/gista/settings/storage-handler"

	"strings"
)

type File struct {
	baseFolder   string
	userFolder   string
	userSettings *map[string]string
	settingsFile string
	cookiesFile  string
	username     string
	storage      storage_handler.StorageInterface
}

func (f *File) SetUsername(username string) error {
	panic("implement me")
}

func (f *File) SetSettings(key, value string) error {
	panic("implement me")
}

func (f *File) OpenLocation(config map[string]string) error {
	baseFolder := ""
	if val, ok := config["baseFolder"]; ok {
		if val != "" {
			baseFolder = val
		}
	}
	if baseFolder == "" {
		baseFolder = "sessions"
	}
	err := f.createFolder(baseFolder)
	if err != nil {
		return errs.CreateFolder(err.Error())
	}
	f.baseFolder = baseFolder
	return nil
}

func (f *File) HasUser(username string) bool {
	hasUser := f.generateUserPaths(username)
	if _, err := os.Stat(hasUser["settingsFile"]); os.IsNotExist(err) {
		return false
	}
	return true
}

func (f *File) MoveUser(oldUsername, newUsername string) error {
	oldUser := f.generateUserPaths(oldUsername)
	newUser := f.generateUserPaths(oldUsername)
	if !f.fileOrFolderExists(oldUser["userFolder"]) {
		return errs.PathNotExist(oldUser["userFolder"])
	}
	if f.fileOrFolderExists(newUser["userFolder"]) {
		return errs.PathAlreadyExists(newUser["userFolder"])
	}
	if err := f.createFolder(newUser["userFolder"]); err != nil {
		return errs.CreateFolder(err.Error())
	}
	if err := f.moveFile(oldUser["settingsFile"], newUser["settingsFile"]); err != nil {
		return errs.MoveFile(oldUser["settingsFile"], newUser["settingsFile"], err.Error())
	}
	if err := f.moveFile(oldUser["cookiesFile"], newUser["cookiesFile"]); err != nil {
		return errs.MoveFile(oldUser["cookiesFile"], newUser["cookiesFile"], err.Error())
	}
	_ = f.removeFolder(oldUser["userFolder"])
	return nil
}

func (f *File) DeleteUser(username string) error {
	userPaths := f.generateUserPaths(username)
	return f.removeFolder(userPaths["userFolder"])
}

func (f *File) OpenUser(username string) error {
	f.username = username
	userPaths := f.generateUserPaths(username)
	f.userFolder = userPaths["userFolder"]
	f.settingsFile = userPaths["settingsFile"]
	f.cookiesFile = userPaths["cookiesFile"]
	return f.createFolder(f.userFolder)
}

func (f *File) LoadUserSettings() (userSettings map[string]string, err error) {
	userSettings = map[string]string{}
	if !f.fileOrFolderExists(f.settingsFile) {
		return
	}
	rawData, err := f.fileGetContents(f.settingsFile)
	if err != nil {
		err = errs.ReadFile(f.settingsFile, err.Error())
		return
	}
	dataVersion := 1
	if idx := strings.Index(rawData, "FILESTORAGEv"); idx != -1 {
		v := rawData[idx+len("FILESTORAGEv")]
		if vInt, err := strconv.Atoi(string(v)); err == nil {
			dataVersion = vInt
			rawData = rawData[strings.Index(rawData, ";")+1:]
		}
	}
	userSettings, err = f.decodeStorage(dataVersion, rawData)
	return
}

func (f *File) SaveUserSettings(userSettings *map[string]string, key string) error {
	versionHeader := fmt.Sprintf("FILESTORAGEv%d;", storageVersion)
	settingsJson, err := json.Marshal(userSettings)
	if err != nil {
		return err
	}
	encodedData := versionHeader + string(settingsJson)
	return atomic.WriteFile(f.settingsFile, strings.NewReader(encodedData))
}

func (f *File) HasUserCookies() bool {
	fi, err := os.Stat(f.cookiesFile)
	if err == nil {
		if f.fileOrFolderExists(f.cookiesFile) && fi.Size() > 0 {
			return true
		}
	}
	return false
}

func (f *File) GetUserCookiesFilePath() string {
	return f.cookiesFile
}

func (f *File) LoadUserCookies() (cookies *string, err error) {
	// never called for
	return
}

func (f *File) SaveUserCookies(rawData string) error {
	// never called for
	return nil
}

func (f *File) CloseUser() {
	f.userFolder = ""
	f.settingsFile = ""
	f.cookiesFile = ""
	f.username = ""
}

func (f *File) CloseLocation() error {
	panic("implement me")
}
