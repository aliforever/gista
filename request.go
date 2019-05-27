package gista

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aliforever/gista/errs"

	"github.com/aliforever/gista/responses"

	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/signatures"
	"github.com/aliforever/gista/utils"
)

type request struct {
	parent           *client
	url              string
	params           map[string]string
	body             *string
	headers          map[string]string
	posts            map[string]string
	defaultHeaders   bool
	files            map[string]map[string]interface{}
	needsAuth        bool
	signedGet        bool
	signedPost       bool
	excludeSigned    *[]string
	handles          map[string]io.Reader
	isBodyCompressed bool
	apiVersion       int
	httpResponse     *http.Response
}

func newRequest(address string, parent *client) (r *request) {
	r = &request{parent: parent}
	r.url = address
	r.apiVersion = 1
	r.headers = map[string]string{}
	r.params = map[string]string{}
	r.posts = map[string]string{}
	r.files = map[string]map[string]interface{}{}
	r.handles = map[string]io.Reader{}
	r.needsAuth = true
	r.signedGet = false
	r.signedPost = true
	r.isBodyCompressed = false
	r.excludeSigned = &[]string{}
	r.defaultHeaders = true
	return
}

func (r *request) SetIsBodyCompressed(val bool) *request {
	r.isBodyCompressed = val
	return r
}

func (r *request) SetSignedPost(val bool) *request {
	r.signedPost = val
	return r
}

func (r *request) AddParam(key, val string) *request {
	if r.params == nil {
		r.params = map[string]string{}
	}
	r.params[key] = val
	return r
}

func (r *request) AddCSRFParam() *request {
	t := r.parent.GetToken()
	token := ""
	if t != nil {
		token = *t
	}
	r.params["_csrftoken"] = token
	return r
}

func (r *request) AddCSRFPost() *request {
	t := r.parent.GetToken()
	token := ""
	if t != nil {
		token = *t
	}
	r.posts["_csrftoken"] = token
	return r
}

func (r *request) AddPhoneIdPost() *request {
	r.posts["phone_id"] = r.parent.instagram.phoneId
	return r
}

func (r *request) AddAdIdPost() *request {
	r.posts["adid"] = r.parent.instagram.advertisingId
	return r
}

func (r *request) AddGuIdPost() *request {
	r.posts["guid"] = r.parent.instagram.uuid
	return r
}

func (r *request) AddUuIdPost() *request {
	r.posts["_uuid"] = r.parent.instagram.uuid
	return r
}

func (r *request) AddUIdPost() *request {
	r.posts["_uid"] = *r.parent.instagram.AccountId
	return r
}

func (r *request) AddIdPost() *request {
	r.posts["id"] = *r.parent.instagram.AccountId
	return r
}

func (r *request) AddDeviceIdPost() *request {
	r.posts["device_id"] = r.parent.instagram.deviceId
	return r
}

func (r *request) AddPost(key, val string) *request {
	if r.posts == nil {
		r.posts = map[string]string{}
	}
	r.posts[key] = val
	return r
}

func (r *request) AddFile(key, filePath string, fileName *string, headers map[string]string) (*request, error) {
	if !utils.FileOrFolderExists(filePath) {
		return r, errs.PathNotExist(filePath)
	}
	if fileName == nil {
		fileName = &filePath
	}
	name := filepath.Base(*fileName)
	fileName = &name
	if headers == nil {
		headers = map[string]string{}
	}
	headers["Content-Type"] = "application/octet-stream"
	headers["Content-Transfer-Encoding"] = "binary"
	r.files[key] = map[string]interface{}{
		"filepath": filePath,
		"filename": *fileName,
		"headers":  headers,
	}
	return r, nil
}

func (r *request) resetHandles() {
	r.handles = map[string]io.Reader{}
}

func (r *request) closeHandles() {
	if len(r.handles) > 0 {
		for k, v := range r.handles {
			switch v.(type) {
			case *os.File:
				v.(*os.File).Close()
				delete(r.handles, k)
			}
		}
	}
	r.resetHandles()
}

func (r *request) AddHeader(key, val string) *request {
	if r.headers == nil {
		r.headers = map[string]string{}
	}
	r.headers[key] = val
	return r
}

func (r *request) addDefaultHeaders() *request {
	if r.defaultHeaders {
		r.headers["X-IG-App-ID"] = constants.FacebookAnalyticsApplicationId
		r.headers["X-IG-Capabilities"] = constants.XIgCapabilities
		r.headers["X-IG-Connection-Type"] = constants.XIgConnectionType
		r.headers["X-IG-Connection-Speed"] = fmt.Sprintf("%dkbps", utils.MtRand(1000, 3700))
		r.headers["X-IG-Bandwidth-Speed-KBPS"] = "-1.000"
		r.headers["X-IG-Bandwidth-TotalBytes-B"] = "0"
		r.headers["X-IG-Bandwidth-TotalTime-MS"] = "0"
	}
	return r
}

func (r *request) SetNeedsAuth(needs bool) *request {
	r.needsAuth = needs
	return r
}

func (r *request) getRequestBody() (body io.Reader, contentType string, err error) {
	if r.body != nil {
		if r.isBodyCompressed {
			//
		}
		body = strings.NewReader(*r.body)
		return
	}
	if len(r.posts) == 0 && len(r.files) == 0 {
		return
	}
	if r.signedPost {
		r.posts = signatures.SignData(r.posts, r.excludeSigned)
	}
	if len(r.files) == 0 {
		body = r.getUrlEncodedBody()
	} else {
		body, contentType, err = r.getMultiPartBody()
	}
	if r.isBodyCompressed {
		//ZLIB Encode Body
	}
	return
}

func (r *request) getMultiPartBody() (body *bytes.Buffer, contentType string, err error) {
	newMap := map[string]interface{}{}
	for k, v := range r.files {
		newMap[k] = v
	}
	for k, v := range r.posts {
		newMap[k] = v
	}
	// ReorderByHashCodeNeededForNewMap
	body = &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, v := range newMap {
		if _, ok := r.files[k]; !ok {
			_ = writer.WriteField(k, v.(string))
		} else {
			fileMap := r.files[k]
			if fileMap["contents"] != nil {
				part, partErr := writer.CreateFormFile(k, filepath.Base((fileMap["filepath"]).(string)))
				if partErr != nil {
					err = errs.CannotCreateFormFieldFromFile((fileMap["filepath"]).(string), partErr.Error())
					return
				}
				_, err = io.Copy(part, bytes.NewReader((fileMap["filepath"]).([]byte)))
				if err != nil {
					return
				}
			} else if fileMap["filepath"] != nil {
				filePath := (fileMap["filepath"]).(string)
				file, fErr := os.Open(filePath)
				if fErr != nil {
					err = errs.CannotOpenFile(filePath, fErr.Error())
					return
				}
				r.handles[k] = file
				/*for k, v := range r.handles {
					switch v.(type) {
					case *os.File:
						v.(*os.File).Close()
					}
				}*/
				part, partErr := writer.CreateFormFile(k, fileMap["filename"].(string))
				if partErr != nil {
					err = errs.CannotCreateFormFieldFromFile(filePath, partErr.Error())
					return
				}
				_, err = io.Copy(part, file)
				if err != nil {
					return
				}
			} else {
				err = errs.NoDataForStreamCreation
				return
			}
		}
	}
	contentType = writer.FormDataContentType()
	// ReorderByHashCodeNeededForPosts
	err = writer.Close()
	return
}

func (r *request) getUrlEncodedBody() io.Reader {
	r.headers["Content-Type"] = constants.ContentType
	// ReorderByHashCodeNeededForPosts
	return strings.NewReader(r.mapToForm(r.posts).Encode())
}

func (r *request) mapToForm(data map[string]string) url.Values {
	values := url.Values{}
	for k, v := range data {
		values.Add(k, v)
	}
	return values
}

func (r *request) buildHttpRequest() (req *http.Request, err error) {
	endPoint := r.url
	if strings.Index(endPoint, "http:") != 0 && strings.Index(endPoint, "https:") != 6 {
		endPoint = constants.ApiUrls[r.apiVersion] + endPoint
	}
	if r.signedGet {
		//r.params = signatures.SignData(r.params, nil)
	}
	if len(r.params) > 0 {
		if strings.Index(endPoint, "?") == -1 {
			endPoint += "?"
		} else {
			endPoint += "&"
		}
		frm := url.Values{}
		for k, v := range r.params {
			frm.Add(k, v)
		}
		endPoint += frm.Encode()
	}
	r.addDefaultHeaders()

	postData, contentType, err := r.getRequestBody()
	r.closeHandles()
	if err != nil {
		err = errs.ErrorBuildingHTTPRequest(err.Error())
		return
	}
	method := "GET"
	if len(r.posts) != 0 {
		method = "POST"
	}
	req, err = http.NewRequest(method, endPoint, postData)
	if err != nil {
		err = errs.ErrorBuildingHTTPRequest(err.Error())
	}
	if len(r.headers) > 0 {
		for k, v := range r.headers {
			req.Header.Set(k, v)
		}
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	return
}

func (r *request) getHTTPResponse() (resp *http.Response, err error) {
	if r.httpResponse == nil {
		if r.needsAuth {
			if !r.parent.instagram.isMaybeLoggedIn {
				err = errs.NotLoggedIn
				return
			}
		}
		r.resetHandles()
		var (
			req *http.Request
			/*respByte []byte*/
		)
		req, err = r.buildHttpRequest()
		if err != nil {
			err = errs.ErrorBuildingHTTPRequest(err.Error())
			return
		}
		dump, dumpErr := httputil.DumpRequest(req, true)
		if dumpErr == nil {
			filePrefix := fmt.Sprintf("%s_%s", time.Now().Format("20060102-150405"), req.Method)
			if err := ioutil.WriteFile(filePrefix+"_request.dump", dump, 0644); err != nil {
				fmt.Println(err)
			}
		}
		r.httpResponse, err = r.parent.api(req)
		if err != nil {
			err = errs.ErrorGettingHTTPResponse(err.Error())
			return
		}
		/*respByte, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			err = errors.ErrorReadingHTTPResponseBody(err.Error())
			return
		}
		strBody := string(respByte)
		r.httpResponse = &strBody*/
	}
	resp = r.httpResponse
	return
}

func (r *request) GetRawResponse() (raw string, err error) {
	httpResponse, respError := r.getHTTPResponse()
	if respError != nil {
		err = respError
		return
	}
	defer httpResponse.Body.Close()
	var reader io.ReadCloser
	switch httpResponse.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(httpResponse.Body)
		defer reader.Close()
	default:
		reader = httpResponse.Body
	}
	body, readErr := ioutil.ReadAll(reader)
	if readErr != nil {
		err = readErr
		return
	}
	raw = string(body)
	return
}

func (r *request) GetResponse(object interface{}) (err error) {
	raw, rawResp := r.GetRawResponse()
	if rawResp != nil {
		err = rawResp
		return
	}
	httpResp, respErr := r.getHTTPResponse()
	if respErr != nil {
		err = respErr
		return
	}
	//{"message": "CSRF token missing or incorrect", "status": "fail"}
	err = r.MapServerResponse(object, raw, httpResp)
	return
}

func (r *request) MapServerResponse(object interface{}, rawResponse string, httpResponse *http.Response) (err error) {
	if r.parent.instagram.httpResponseInResult {
		object.(responses.ResponseInterface).SetHTTPResponse(httpResponse)
	}
	if r.parent.instagram.rawResponseInResult {
		object.(responses.ResponseInterface).SetRawResponse(rawResponse)
	}
	err = json.Unmarshal([]byte(rawResponse), &object)
	object.(responses.ResponseInterface).SetIsOk()
	if err != nil {
		httpStatusCode := httpResponse.StatusCode
		switch httpStatusCode {
		case 400:
			err = errs.InvalidRequestOptions
		case 404:
			err = errs.RequestedResourceNotExist
		default:
			switch err.(type) {
			case *json.UnmarshalTypeError:
				t, m := "json_unmarshal", err.Error()
				err = errs.JsonUnmarshal{Type: &t, Message: &m}
			default:
				err = errs.NoResponseFromServer
			}
		}
		return
	}
	if !object.(responses.ResponseInterface).IsOk() {
		message := object.(responses.ResponseInterface).GetMessage()
		t := fmt.Sprintf("%T", object)
		t = strings.Replace(t, "*responses.", "", -1)
		t += "Response"
		err = errs.GetError(&t, &message, object.(responses.ResponseInterface), httpResponse)
	}
	return
}
