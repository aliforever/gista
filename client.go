package gista

import (
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/aliforever/gista/middleware"

	"github.com/aliforever/gista/errors"

	"github.com/aliforever/gista/constants"
)

const cookieAutoSaveInterval int64 = 45

type client struct {
	userAgent          string
	client             *http.Client
	instagram          *Instagram
	cookieJar          *cookiejar.Jar
	cookieJarLastSaved int64
	zeroRating         *middleware.ZeroRating
}

func newClient(i *Instagram) (c *client) {
	c = &client{instagram: i}
	c.client = &http.Client{}
	c.zeroRating = middleware.NewZeroRating()
	middleWareContainer := middleware.NewContainer(&http.DefaultTransport)
	middleWareContainer.Push(c.zeroRating)
	c.client.Transport = middleWareContainer
	return
}

func (c *client) Request(address string) (r *request) {
	return newRequest(address, c)
}

/*func (c *client) getStreamForFile(fileMap map[string]*string) (result io.Reader, err error) {
	//https://medium.com/@owlwalks/sending-big-file-with-minimal-memory-in-golang-8f3fc280d2c
	if fileMap["contents"] != nil {
		result = strings.NewReader(*fileMap["contents"])
	} else if fileMap["filepath"] != nil {
		file, fErr := os.Open(*fileMap["filepath"])
		if fErr != nil {
			err = errors.CannotOpenFile(*fileMap["filepath"], fErr.Error())
			return
		}
		defer file.Close()
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile(fileMap, filepath.Base(path))
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(part, file)
	} else {
		err = errors.NoDataForStreamCreation
	}
	return
}
*/

func (c *client) api(request *http.Request) (resp *http.Response, err error) {
	request.Header.Set("User-Agent", c.userAgent)
	request.Header.Set("Connection", "Keep-Alive")
	request.Header.Set("X-FB-HTTP-Engine", constants.XFbHttpEngine)
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Accept-Encoding", constants.AcceptEncoding)
	request.Header.Set("Accept-Language", constants.AcceptLanguage)

	if c.cookieJar != nil {
		c.client.Jar = c.cookieJar
	}
	resp, err = c.client.Do(request)
	if err != nil {
		return
	}
	/*defer resp.Body.Close()*/
	statusCode := resp.StatusCode
	switch statusCode {
	case 429:
		err = errors.ThrottledResponse
		return
	case 431:
		err = errors.RequestHeaderTooLargeResponse
		return
	}
	if time.Now().Unix()-c.cookieJarLastSaved > cookieAutoSaveInterval {
		err = c.SaveCookieJar()
	}
	return
}

func (c *client) UpdateFromCurrentSettings(resetCookieJar bool) {
	c.userAgent = c.instagram.device.GetUserAgent()
	c.LoadCookieJar(resetCookieJar)
	if c.GetToken() == nil {
		c.instagram.isMaybeLoggedIn = false
	}
}

func (c *client) LoadCookieJar(resetCookieJar bool) {
	c.cookieJar = nil
	if resetCookieJar {
		data := ""
		c.instagram.settings.SetCookies(&data)
	}
	var cookies *string
	cookies, _ = c.instagram.settings.GetCookies()

	var restoredCookies []map[string]interface{}
	if cookies != nil && len(*cookies) > 0 {
		json.Unmarshal([]byte(*cookies), &restoredCookies)
	}
	c.cookieJar = c.mapSliceToCookieJar(restoredCookies)
	c.cookieJarLastSaved = time.Now().Unix()
}

func (c *client) cookieJarToMapSlice(cj *cookiejar.Jar) (mapSlice []map[string]interface{}) {
	address, _ := url.Parse(constants.CookieUrl)
	cookies := cj.Cookies(address)
	for _, c := range cookies {
		mapSlice = append(mapSlice, map[string]interface{}{
			"Name":     c.Name,
			"Value":    c.Value,
			"Domain":   c.Domain,
			"Path":     c.Path,
			"MaxAge":   c.MaxAge,
			"Expires":  c.Expires,
			"Secure":   c.Secure,
			"HttpOnly": c.HttpOnly,
		})
	}
	return
}

func (c *client) mapSliceToCookieJar(mapSlice []map[string]interface{}) (cj *cookiejar.Jar) {
	cj, _ = cookiejar.New(nil)
	var cookies []*http.Cookie
	for _, m := range mapSlice {
		cookie := http.Cookie{}
		cookie.Name = m["Name"].(string)
		cookie.Value = m["Value"].(string)
		cookie.Domain = m["Domain"].(string)
		cookie.Path = m["Path"].(string)
		cookie.MaxAge = int(m["MaxAge"].(float64))
		layout := "2006-01-02T15:04:05Z"
		t, _ := time.Parse(layout, m["Expires"].(string))
		cookie.Expires = t
		cookie.Secure = m["Secure"].(bool)
		cookie.HttpOnly = m["HttpOnly"].(bool)
		cookies = append(cookies, &cookie)
	}
	address, _ := url.Parse(constants.CookieUrl)
	cj.SetCookies(address, cookies)
	return
}

func (c *client) GetCookie(name string, domain *string, path *string) (cookie *http.Cookie) {
	if c.cookieJar != nil {
		address, _ := url.Parse(*domain)
		for _, C := range (*c.cookieJar).Cookies(address) {
			if C.Name == name && (C.Expires.IsZero() || !C.Expires.Before(time.Now())) && (path == nil || C.Path == *path) {
				cookie = C
			}
		}
	}
	return
}

func (c *client) SaveCookieJar() (err error) {
	newCookies, err := c.GetCookieJarAsJSON()
	err = c.instagram.settings.SetCookies(newCookies)
	if err != nil {
		return
	}
	c.cookieJarLastSaved = time.Now().Unix()
	return
}

func (c *client) GetCookieJarAsJSON() (jar *string, err error) {
	if c.cookieJar == nil {
		empty := ""
		return &empty, nil
	}
	mapSlice := c.cookieJarToMapSlice(c.cookieJar)
	var jsonByte []byte
	jsonByte, err = json.Marshal(mapSlice)
	if err != nil {
		err = errors.CannotMarshalJSON(mapSlice, err.Error())
		return
	}
	str := string(jsonByte)
	jar = &str
	return
}

func (c *client) GetToken() (t *string) {
	cookie := c.GetCookie("csrftoken", &constants.CookieUrl, nil)
	if cookie == nil || cookie.Value == "" {
		return nil
	}
	t = &cookie.Value
	return
}

func (c *client) ZeroRating() *middleware.ZeroRating {
	return c.zeroRating
}
