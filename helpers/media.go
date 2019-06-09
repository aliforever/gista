package helpers

import (
	"io"
	"net/http"
	"os"
	"path"

	errs2 "github.com/aliforever/gista/errs"

	"github.com/aliforever/gista/utils"

	"github.com/aliforever/gista/constants"

	"github.com/aliforever/gista/models/item"

	"github.com/aliforever/gista/models"
)

func ShortCodeToMediaId(shortCode string) int64 {
	alphabet := map[string]int64{
		"-": 62, "1": 53, "0": 52, "3": 55, "2": 54, "5": 57, "4": 56, "7": 59, "6": 58, "9": 61,
		"8": 60, "A": 0, "C": 2, "B": 1, "E": 4, "D": 3, "G": 6, "F": 5, "I": 8, "H": 7,
		"K": 10, "J": 9, "M": 12, "L": 11, "O": 14, "N": 13, "Q": 16, "P": 15, "S": 18, "R": 17,
		"U": 20, "T": 19, "W": 22, "V": 21, "Y": 24, "X": 23, "Z": 25, "_": 63, "a": 26, "c": 28,
		"b": 27, "e": 30, "d": 29, "g": 32, "f": 31, "i": 34, "h": 33, "k": 36, "j": 35, "m": 38,
		"l": 37, "o": 40, "n": 39, "q": 42, "p": 41, "s": 44, "r": 43, "u": 46, "t": 45, "w": 48,
		"v": 47, "y": 50, "x": 49, "z": 51,
	}
	n := int64(0)
	i := 0

	for i < len(shortCode) {
		c := shortCode[i : i+1]
		n = n*64 + alphabet[c]
		i++
	}
	return n
}

func DownloadItem(i *models.Item, path *string) (errs map[string]error) {
	if i == nil {
		errs = map[string]error{}
		errs["default"] = errs2.InvalidItem(i)
		return
	}
	info, err := GetItemBestQualityUrl(i)
	if err != nil {
		return
	}
	p := func(key string, info map[string]string) (path *string) {
		_path := constants.DownloadPaths[info["type"]] + "/" + key
		if info["type"] == "photo" {
			_path += ".jpg"
		} else {
			_path += ".mp4"
		}
		return &_path
	}
	mediaId := ""
	if path == nil {
		for mId := range info {
			mediaId = mId
			path = p(mediaId, info[mId])
			err = Download(info[mediaId]["url"], *path)
			if err != nil {
				if errs == nil {
					errs = map[string]error{}
				}
				errs[info[mediaId]["url"]] = err
			}
		}
	}

	return
}

func Download(address, filePath string) (err error) {
	err = utils.CreateFolder(path.Dir(filePath))
	if err != nil {
		err = errs2.CreateFolder(err.Error())
		return
	}
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	resp, err := http.Get(address)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = errs2.InvalidHTTPStatus(resp.StatusCode)
		return
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return
}

func GetItemBestQualityUrl(i *models.Item) (info map[string]map[string]string, err error) {
	if i == nil {
		err = errs2.InvalidItem(i)
		return
	}

	info = map[string]map[string]string{}
	var b string
	if i.MediaType == item.Photo {
		info[i.Id] = map[string]string{}
		info[i.Id]["type"] = "photo"
		b, err = GetMediaBestQualityUrl(i.ImageVersions2)
		info[i.Id]["url"] = b
	} else if i.MediaType == item.Video {
		info[i.Id] = map[string]string{}
		info[i.Id]["type"] = "video"
		b, err = GetMediaBestQualityUrl(i.VideoVersions)
		info[i.Id]["url"] = b
	} else if i.MediaType == item.Carousel {
		for _, m := range i.CarouselMedia {
			var t interface{}
			info[m.Id] = map[string]string{}
			if m.MediaType == 1 {
				info[m.Id]["type"] = "photo"
				t = m.ImageVersions2
			} else if m.MediaType == 2 {
				info[m.Id]["type"] = "video"
				t = m.VideoVersions
			}
			info[m.Id]["url"], err = GetMediaBestQualityUrl(t)
		}
	}
	return
}

func GetMediaBestQualityUrl(i interface{}) (url string, err error) {
	switch i.(type) {
	case *models.ImageVersions2:
		biggestHeight := 0
		candidates := i.(*models.ImageVersions2).Candidates
		for _, c := range candidates {
			if c.Height > biggestHeight {
				biggestHeight = c.Height
				url = c.Url
			}
		}
		break
	case *[]models.VideoVersion:
		biggestHeight := 0
		for _, v := range *i.(*[]models.VideoVersion) {
			if v.Height > biggestHeight {
				biggestHeight = v.Height
				url = v.Url
			}
		}
		break
		//sort.Ints(widths)
	default:
		err = errs2.UnknownItem(i)
	}
	return
}
