package httpClient

import (
	"net/http"
	"time"
)

type HttpClientBuilder interface {
	SetHeaders(headers http.Header) HttpClientBuilder // Returning interface so we can concatenate methods

	SetMaxIdleConnections(maxIdleConnections int) HttpClientBuilder
	SetResponseTimeOut(requestTimeOut time.Duration) HttpClientBuilder
	SetConnectionTimeout(connectionTimeout time.Duration) HttpClientBuilder

	Build() HttpClient
}

type builder struct {
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeOut    time.Duration

	headers http.Header
}

func NewBuilder() HttpClientBuilder {
	return &builder{
		maxIdleConnections: defaultMaxIdleConnections,
		connectionTimeout:  defaultConnectionTimeout,
		responseTimeOut:    defaultResponseTimeOut,
	}
}

func (b *builder) Build() HttpClient {
	return &client{
		maxIdleConnections: b.maxIdleConnections,
		connectionTimeout:  b.connectionTimeout,
		responseTimeOut:    b.responseTimeOut,

		headers: b.headers,
	}
}

// Set common headers to use during all client life
func (b *builder) SetHeaders(headers http.Header) HttpClientBuilder {
	b.headers = headers
	return b
}

func (b *builder) SetMaxIdleConnections(maxIdleConnections int) HttpClientBuilder {
	b.maxIdleConnections = maxIdleConnections
	return b
}

func (b *builder) SetConnectionTimeout(connectionTimeout time.Duration) HttpClientBuilder {
	b.connectionTimeout = connectionTimeout
	return b
}

func (b *builder) SetResponseTimeOut(responseTimeOut time.Duration) HttpClientBuilder {
	b.responseTimeOut = responseTimeOut
	return b
}
