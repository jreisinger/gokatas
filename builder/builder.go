// Package builder builds *http.Request objects using the builder design
// pattern. The pattern is useful for creating complex objects step by step so
// you don't end up with a single function with dozens of parameters. Based on
// "GoF design patterns that still make sense in Go" (bit.ly/37BaClv).
package builder

import (
	"context"
	"io"
	"net/http"
)

// NewBuilder creates a HTTPBuilder given a URL. We don't leak the actual
// builder to achieve information hiding. Also we set some sane defaults and
// don't have to worry about null/empty values.
func NewBuilder(url string) HTTPBuilder {
	return &builder{
		url:     url,
		ctx:     context.Background(),
		method:  http.MethodGet,
		headers: map[string][]string{},
		body:    nil,
		close:   false,
	}
}

// HTTPBuilder defines the fields we want to set on this builder, you could
// add/remove fields here.
type HTTPBuilder interface {
	Method(method string) HTTPBuilder
	AddHeader(name, value string) HTTPBuilder
	Body(r io.Reader) HTTPBuilder
	Close(close bool) HTTPBuilder
	Build() (*http.Request, error)
}

type builder struct {
	ctx     context.Context
	url     string
	method  string
	headers map[string][]string
	body    io.Reader
	close   bool
}

func (b *builder) Method(method string) HTTPBuilder {
	b.method = method
	return b
}
func (b *builder) AddHeader(name, value string) HTTPBuilder {
	values, found := b.headers[name]
	if !found {
		values = make([]string, 0, 10)
	}
	b.headers[name] = append(values, value)
	return b
}
func (b *builder) Body(r io.Reader) HTTPBuilder {
	b.body = r
	return b
}
func (b *builder) Close(close bool) HTTPBuilder {
	b.close = close
	return b
}

func (b *builder) Build() (*http.Request, error) {
	r, err := http.NewRequestWithContext(b.ctx, b.method, b.url, b.body)
	if err != nil {
		return nil, err
	}
	for key, values := range b.headers {
		for _, value := range values {
			r.Header.Add(key, value)
		}
	}
	r.Close = b.close
	return r, nil
}
