package httpClient

import (
	"net/http"
)

type HttpClient interface {
	SetHeaders(headers http.Header)
	GET(url string, headers http.Header) (*http.Response, error)
	POST(url string, headers http.Header, body interface{}) (*http.Response, error)
	PUT(url string, headers http.Header, body interface{}) (*http.Response, error)
	PATCH(url string, headers http.Header, body interface{}) (*http.Response, error)
	DELETE(url string, headers http.Header) (*http.Response, error)
}

type httpClient struct {
	headers http.Header
}

func New() HttpClient {
	return &httpClient{}
}

// Set common headers to use during all client life
func (hc *httpClient) SetHeaders(headers http.Header) {
	hc.headers = headers
}

func (hc *httpClient) GET(url string, headers http.Header) (*http.Response, error) {
	return hc.do(http.MethodGet, url, headers, nil)
}

func (hc *httpClient) POST(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return hc.do(http.MethodGet, url, headers, body)
}

func (hc *httpClient) PUT(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return hc.do(http.MethodGet, url, headers, body)
}

func (hc *httpClient) PATCH(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return hc.do(http.MethodGet, url, headers, body)
}

func (hc *httpClient) DELETE(url string, headers http.Header) (*http.Response, error) {
	return hc.do(http.MethodGet, url, headers, nil)
}
