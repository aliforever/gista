package responses

import (
	"fmt"
	"go/types"
	"net/http"
	"strings"
)

type Response struct {
	httpResponse *http.Response
	rawResponse  string
	Status       string `json:"status,omitempty"`
	isOk         bool
	Message      interface{} `json:"message,omitempty"`
	ErrorType    *string     `json:"error_type,omitempty"`
	ErrorTitle   interface{} `json:"error_title,omitempty"`
}

func (r *Response) GetHTTPResponse() *http.Response {
	return r.httpResponse
}

func (r *Response) SetHTTPResponse(response *http.Response) {
	r.httpResponse = response
}

func (r *Response) GetRawResponse() string {
	return r.rawResponse
}

func (r *Response) SetRawResponse(response string) {
	r.rawResponse = response
}

func (r *Response) IsOk() bool {
	return r.isOk
}

func (r *Response) SetIsOk() {
	if r.Status == "ok" {
		r.isOk = true
	} else {
		r.isOk = false
	}
}

func (r *Response) GetErrorType() (str *string) {
	return r.ErrorType
}

func (r *Response) GetMessage() (str string) {
	switch r.Message.(type) {
	case string, types.Nil:
		str = r.Message.(string)
		return
	case map[string][]string:
		mp := r.Message.(map[string][]string)
		_, ok := mp["errors"]
		if len(mp) == 1 && ok && len(mp["errors"]) > 0 {
			str = ""
			if len(mp["errors"]) > 1 {
				str = "Multiple Errors: "
			}
			str = strings.Join(mp["errors"], " AND ")
			return
		}
	default:
		str = fmt.Sprintf("%+v", r.Message)
		return
	}
	return
}
