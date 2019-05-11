package responses

import "net/http"

type ResponseInterface interface {
	SetHTTPResponse(resp *http.Response)
	SetRawResponse(raw string)
	SetIsOk()
	IsOk() bool
}
