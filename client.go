package gohttpclient

import (
	"net/http"
	"sync"

	"github.com/maxiancillotti/gohttpclient/httpcore"
)

type Client interface {
	GET(url string, headers http.Header) (*http.Response, error)
	POST(url string, headers http.Header, body interface{}) (*http.Response, error)
	PUT(url string, headers http.Header, body interface{}) (*http.Response, error)
	PATCH(url string, headers http.Header, body interface{}) (*http.Response, error)
	DELETE(url string, headers http.Header) (*http.Response, error)

	OPTIONS(url string, headers http.Header) (*http.Response, error)
	HEAD(url string, headers http.Header) (*http.Response, error)
	CONNECT(url string, headers http.Header) (*http.Response, error)
	TRACE(url string, headers http.Header) (*http.Response, error)
}

type client struct {
	httpClient httpcore.HttpClient //*http.Client. Only one http client is created and can be reused on every call
	builder    *clientBuilder
	clientOnce sync.Once
}

func (c *client) GET(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *client) POST(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *client) PUT(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *client) PATCH(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}

func (c *client) DELETE(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, nil)
}

func (c *client) OPTIONS(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodOptions, url, headers, nil)
}

func (c *client) HEAD(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodHead, url, headers, nil)
}

func (c *client) CONNECT(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodConnect, url, headers, nil)
}

func (c *client) TRACE(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodTrace, url, headers, nil)
}
