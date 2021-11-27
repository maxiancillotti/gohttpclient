package gohttpclient

import (
	"net/http"
	"sync"
)

type Client interface {
	GET(url string, headers http.Header) (*http.Response, error)
	POST(url string, headers http.Header, body interface{}) (*http.Response, error)
	PUT(url string, headers http.Header, body interface{}) (*http.Response, error)
	PATCH(url string, headers http.Header, body interface{}) (*http.Response, error)
	DELETE(url string, headers http.Header) (*http.Response, error)
}

type client struct {
	httpClient *http.Client // Only one http client is created and can be reused on every call
	builder    *clientBuilder
	clientOnce sync.Once
}

func (c *client) GET(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *client) POST(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, body)
}

func (c *client) PUT(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, body)
}

func (c *client) PATCH(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, body)
}

func (c *client) DELETE(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}
