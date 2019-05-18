package storage_file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/aliforever/gista/errs"
)

func (f *File) createFolder(folderPath string) error {
	return os.MkdirAll(folderPath, os.ModePerm)
}

func (f *File) removeFolder(folderPath string) error {
	return os.Remove(folderPath)
}

func (f *File) moveFile(oldPath, newPath string) error {
	return os.Rename(oldPath, newPath)
}

func (f *File) generateUserPaths(username string) map[string]string {
	userFolder := f.baseFolder + "/" + username
	settingsFile := userFolder + "/" + fmt.Sprintf(settingsFileName, username)
	cookiesFile := userFolder + "/" + fmt.Sprintf(cookiesFileName, username)
	return map[string]string{
		"userFolder":   userFolder,
		"settingsFile": settingsFile,
		"cookiesFile":  cookiesFile,
	}
}

func (f *File) fileOrFolderExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func (f *File) fileGetContents(path string) (content string, err error) {
	data, err := ioutil.ReadFile(path)
	return string(data), err
}

func (f *File) decodeStorage(dataVersion int, rawData string) (userSettings map[string]string, err error) {
	userSettings = map[string]string{}
	switch dataVersion {
	case 1:
		/**
		 * This is the old format from v1.x of InstagramInterface-API.
		 * Terrible format. Basic "key=value\r\n" and very fragile.
		 */

		// Split by system-independent newlines. Tries \r\n (Win), then \r
		// (pre-2000s Mac), then \n\r, then \n (Mac OS X, UNIX, Linux).
		r := regexp.MustCompile(`(\r\n?|\n\r?)`)
		lines := r.Split(rawData, -1)
		if len(lines) > 0 {
			for _, line := range lines {
				r := regexp.MustCompile("^([^=]+)=(.*)$")
				matches := r.FindAllString(line, -1)
				key := matches[1]
				value := strings.TrimRight(matches[2], "\r\n ")
				userSettings[key] = value
			}
		}
		break
	case 2:
		/**
		 * Version 2 uses JSON encoding and perfectly stores any value.
		 * And storage-file corruption can't happen, thanks to the atomic writer.
		 */
		err = json.Unmarshal([]byte(rawData), &userSettings)
		if err != nil {
			err = errs.CannotMarshalJSON(rawData, err.Error())
		}
	default:
		err = errs.InvalidStorageVersion(dataVersion)
	}
	return
}
