// Package builder shows the builder pattern. It's good when creating a complex
// object so that you don't end up with a single function call receiving dozens
// of parameters. Based on "GoF design patterns that still make sense in Go"
// (bit.ly/37BaClv).
package builder

import (
	"context"
	"io"
	"net/http"
)

func NewHTTPRequest(url string) HTTPRequest {
	return &request{
		ctx:     context.Background(),
		url:     url,
		method:  http.MethodGet,
		headers: map[string][]string{},
		body:    nil,
		close:   false,
	}
}

type HTTPRequest interface {
	Method(method string) HTTPRequest
	AddHeader(name, value string) HTTPRequest
	Body(r io.Reader) HTTPRequest
	Close(close bool) HTTPRequest
	Build() (*http.Request, error)
}

type request struct {
	ctx     context.Context
	url     string
	method  string
	headers map[string][]string
	body    io.Reader
	close   bool
}

func (req *request) Method(method string) HTTPRequest {
	req.method = method
	return req
}
func (req *request) AddHeader(name, value string) HTTPRequest {
	values, found := req.headers[name]
	if !found {
		values = make([]string, 0, 10)
	}
	req.headers[name] = append(values, value)
	return req
}
func (req *request) Body(r io.Reader) HTTPRequest {
	req.body = r
	return req
}
func (req *request) Close(close bool) HTTPRequest {
	req.close = close
	return req
}

func (req *request) Build() (*http.Request, error) {
	r, err := http.NewRequestWithContext(req.ctx, req.method, req.url, req.body)
	if err != nil {
		return nil, err
	}
	for key, values := range req.headers {
		for _, value := range values {
			r.Header.Add(key, value)
		}
	}
	r.Close = req.close
	return r, nil
}
