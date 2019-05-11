package middleware

import (
	"net/http"
)

type container struct {
	middlewares []Middleware
	transport   *http.RoundTripper
}

func NewContainer(originalTransport *http.RoundTripper) *container {
	return &container{transport: originalTransport}
}

func (c *container) Push(middleware Middleware) {
	c.middlewares = append(c.middlewares, middleware)
}

func (c *container) RoundTrip(r *http.Request) (res *http.Response, err error) {
	for _, m := range c.middlewares {
		r = m.ModifyRequest(r)
	}
	return (*c.transport).RoundTrip(r)
}
