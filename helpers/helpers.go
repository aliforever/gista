package helpers

import (
	"io"
	"net/http"
	"os"
	"path"
	"sort"

	errs2 "github.com/aliforever/gista/errs"

	"github.com/aliforever/gista/utils"

	"github.com/aliforever/gista/constants"

	"github.com/aliforever/gista/models/item"

	"github.com/aliforever/gista/models"
)

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
			info[m.Id]["url"], _ = GetMediaBestQualityUrl(t)
		}
	}
	return
}

func GetMediaBestQualityUrl(i interface{}) (url string, err error) {
	switch i.(type) {
	case models.ImageVersions2:
		var heights []int
		var widths []int
		cMapByH := map[int]string{}
		candidates := i.(models.ImageVersions2).Candidates
		for _, c := range candidates {
			heights = append(heights, c.Height)
			widths = append(widths, c.Width)
			cMapByH[c.Height] = c.Url
		}
		sort.Ints(heights)
		//sort.Ints(widths)
		url = cMapByH[heights[len(heights)-1]]
		break
	case []models.VideoVersion:
		var heights []int
		var widths []int
		cMapByH := map[int]string{}
		for _, v := range i.([]models.VideoVersion) {
			heights = append(heights, v.Height)
			widths = append(widths, v.Width)
			cMapByH[v.Height] = v.Url
		}
		sort.Ints(heights)
		url = cMapByH[heights[len(heights)-1]]
		break
		//sort.Ints(widths)
	default:
		err = errs2.UnknownItem(i)
	}
	return
}
