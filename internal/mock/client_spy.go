package mock

import (
	"io"
	"net/http"
)

type IClientSpy struct {
	StatusCode int
	Body       io.ReadCloser
	Err        error
}

func (c *IClientSpy) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{StatusCode: c.StatusCode, Body: c.Body}, c.Err
}
