package errs

import (
	"github.com/aliforever/gista/responses"
)

type GistaError struct {
	Type         *string
	Message      *string
	HTTPResponse responses.ResponseInterface
}

func (g GistaError) Error() string {
	m := "unknown"
	if g.Message != nil {
		m = *g.Message
	}
	return m
}
