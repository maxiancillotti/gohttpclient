package gohttpclient

import (
	"net/http"
	"time"
)

// Methods returning interface can concatenate method calls
type ClientBuilder interface {

	// SetHeaders: Set common headers to use during all client life
	SetHeaders(headers http.Header) ClientBuilder

	SetMaxIdleConnections(maxIdleConnections int) ClientBuilder
	SetResponseTimeOut(requestTimeOut time.Duration) ClientBuilder
	SetConnectionTimeout(connectionTimeout time.Duration) ClientBuilder

	Build() Client
}

type clientBuilder struct {
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeOut    time.Duration

	headers http.Header
}

func NewBuilder() ClientBuilder {
	return &clientBuilder{
		maxIdleConnections: defaultMaxIdleConnections,
		connectionTimeout:  defaultConnectionTimeout,
		responseTimeOut:    defaultResponseTimeOut,
	}
}

func (b *clientBuilder) Build() Client {
	return &client{
		builder: b,
	}
}

func (b *clientBuilder) SetHeaders(headers http.Header) ClientBuilder {
	b.headers = headers
	return b
}

// Requests per minute is a good metric to set this value
func (b *clientBuilder) SetMaxIdleConnections(maxIdleConnections int) ClientBuilder {
	b.maxIdleConnections = maxIdleConnections
	return b
}

// Request connection timeout
func (b *clientBuilder) SetConnectionTimeout(connectionTimeout time.Duration) ClientBuilder {
	b.connectionTimeout = connectionTimeout
	return b
}

// Response timeout after we have sent the Request
func (b *clientBuilder) SetResponseTimeOut(responseTimeOut time.Duration) ClientBuilder {
	b.responseTimeOut = responseTimeOut
	return b
}
