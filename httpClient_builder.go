package httpClient

import (
	"net/http"
	"time"
)

type HttpClientBuilder interface {
	SetHeaders(headers http.Header)

	SetMaxIdleConnections(maxIdleConnections int)
	SetResponseTimeOut(requestTimeOut time.Duration)
	SetConnectionTimeout(connectionTimeout time.Duration)
}

type httpClientBuilder struct{}

func New() HttpClientBuilder {
	return &httpClientBuilder{}
}

// Set common headers to use during all client life
func (hc *httpClientBuilder) SetHeaders(headers http.Header) {
	hc.headers = headers
}

// Set common headers to use during all client life
func (hc *httpClientBuilder) SetMaxIdleConnections(maxIdleConnections int) {
	hc.maxIdleConnections = maxIdleConnections
	hc.updateTransportSettings = true
}

// Set common headers to use during all client life
func (hc *httpClientBuilder) SetConnectionTimeout(connectionTimeout time.Duration) {
	hc.connectionTimeout = connectionTimeout
	hc.updateTransportSettings = true
}

// Set common headers to use during all client life
func (hc *httpClientBuilder) SetResponseTimeOut(responseTimeOut time.Duration) {
	hc.responseTimeOut = responseTimeOut
	hc.updateTransportSettings = true
}
