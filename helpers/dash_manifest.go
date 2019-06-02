package helpers

import (
	"regexp"
	"strconv"
	"strings"
)

func GetDashManifestBestMediaUrl(dashManifest string) (address string) {
	dashManifest = strings.Replace(dashManifest, `\"`, `"`, -1)
	dashManifest = strings.Replace(dashManifest, "&amp;", "&", -1)
	r := regexp.MustCompile(`(?U)codecs=".+ width="(\d+)" height="(\d+?)".+<BaseURL>(.+)</BaseURL>`)
	result := r.FindAllStringSubmatch(dashManifest, -1)
	type Info struct {
		Width  int
		Height int
		Url    string
	}
	biggestHeight := 0
	for _, info := range result {
		//width, _ := strconv.Atoi(info[1])
		height, _ := strconv.Atoi(info[2])
		link := info[3]
		if height > biggestHeight {
			biggestHeight = height
			address = link
		}
	}
	return
}
