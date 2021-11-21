package httpClient

import (
	"net/http"
	"time"
)

type HttpClient interface {
	SetHeaders(headers http.Header)

	SetMaxIdleConnections(maxIdleConnections int)
	SetResponseTimeOut(requestTimeOut time.Duration)
	SetConnectionTimeout(connectionTimeout time.Duration)

	GET(url string, headers http.Header) (*http.Response, error)
	POST(url string, headers http.Header, body interface{}) (*http.Response, error)
	PUT(url string, headers http.Header, body interface{}) (*http.Response, error)
	PATCH(url string, headers http.Header, body interface{}) (*http.Response, error)
	DELETE(url string, headers http.Header) (*http.Response, error)
}

type httpClient struct {
	client *http.Client // Only one http client is created and can be reused on every call

	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeOut    time.Duration

	updateTransportSettings bool

	headers http.Header
}

func New() HttpClient {
	return &httpClient{
		updateTransportSettings: true,
		maxIdleConnections:      defaultMaxIdleConnections,
		connectionTimeout:       defaultConnectionTimeout,
		responseTimeOut:         defaultResponseTimeOut,
	}
}

// Set common headers to use during all client life
func (hc *httpClient) SetHeaders(headers http.Header) {
	hc.headers = headers
}

// Set common headers to use during all client life
func (hc *httpClient) SetMaxIdleConnections(maxIdleConnections int) {
	hc.maxIdleConnections = maxIdleConnections
	hc.updateTransportSettings = true
}

// Set common headers to use during all client life
func (hc *httpClient) SetConnectionTimeout(connectionTimeout time.Duration) {
	hc.connectionTimeout = connectionTimeout
	hc.updateTransportSettings = true
}

// Set common headers to use during all client life
func (hc *httpClient) SetResponseTimeOut(responseTimeOut time.Duration) {
	hc.responseTimeOut = responseTimeOut
	hc.updateTransportSettings = true
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
