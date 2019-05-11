package middleware

import "net/http"

type Middleware interface {
	ModifyRequest(request *http.Request) (modifiedRequest *http.Request)
}
