package err

import "net/http"

type Err struct {
	error
	httpResponse *http.Response
}

func (e *Err) setResponse(resp *http.Response) {
	e.httpResponse = resp
}
