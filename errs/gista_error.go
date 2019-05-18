package errs

import "net/http"

type GistaError struct {
	Type         *string
	Message      *string
	HTTPResponse *http.Response
}

func (g GistaError) Error() string {
	m := "unknown"
	if g.Message != nil {
		m = *g.Message
	}
	return m
}
