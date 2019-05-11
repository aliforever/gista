package middleware

import (
	"net/http"
	"net/url"
	"regexp"

	"github.com/huandu/xstrings"

	zero_rating "github.com/aliforever/gista/middleware/zero-rating"
)

type ZeroRating struct {
	rules map[string]string
}

func NewZeroRating() *ZeroRating {
	zr := &ZeroRating{}
	zr.Reset()
	return zr
}

func (zr *ZeroRating) Reset() {
	zr.Update(zero_rating.DefaultRewrite)
}

func (zr *ZeroRating) ReWrite(uri string) (result string) {
	for from, to := range zr.rules {
		r := regexp.MustCompile(from)
		result = r.ReplaceAllString(uri, to)
		if result != uri && result != "" {
			return
		}
	}
	result = uri
	return
}

func (zr *ZeroRating) ModifyRequest(req *http.Request) (modifiedRequest *http.Request) {
	modifiedRequest = req
	if zr.rules == nil || len(zr.rules) == 0 {
		return
	}
	oldUri := req.URL.String()
	uri := zr.ReWrite(oldUri)
	if uri != oldUri {
		u, _ := url.Parse(uri)
		modifiedRequest.URL = u
	}
	return
}

func (zr *ZeroRating) Update(rules map[string]string) {
	zr.rules = map[string]string{}
	for from, to := range rules {
		regx := "#" + from + "#"
		r := regexp.MustCompile(regx)
		if r.Match([]byte("")) {
			continue
		}
		zr.rules[regx] = xstrings.Translate(to, `\.`, ".")
	}
}
