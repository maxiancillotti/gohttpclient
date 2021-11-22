package httpClient

import (
	"net/http"
	"time"
)

// Methods returning interface can concatenate method calls
type HttpClientBuilder interface {

	// SetHeaders: Set common headers to use during all client life
	SetHeaders(headers http.Header) HttpClientBuilder
	SetBaseUrl(baseUrl string) HttpClientBuilder

	SetMaxIdleConnections(maxIdleConnections int) HttpClientBuilder
	SetResponseTimeOut(requestTimeOut time.Duration) HttpClientBuilder
	SetConnectionTimeout(connectionTimeout time.Duration) HttpClientBuilder

	Build() HttpClient
}

type clientBuilder struct {
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeOut    time.Duration

	headers http.Header
	baseUrl string
}

func NewBuilder() HttpClientBuilder {
	return &clientBuilder{
		maxIdleConnections: defaultMaxIdleConnections,
		connectionTimeout:  defaultConnectionTimeout,
		responseTimeOut:    defaultResponseTimeOut,
	}
}

func (b *clientBuilder) Build() HttpClient {
	return &client{
		builder: b,
	}
}

func (b *clientBuilder) SetHeaders(headers http.Header) HttpClientBuilder {
	b.headers = headers
	return b
}

func (b *clientBuilder) SetBaseUrl(baseUrl string) HttpClientBuilder {
	b.baseUrl = baseUrl
	return b
}

// Requests per minute is a good metric to set this value
func (b *clientBuilder) SetMaxIdleConnections(maxIdleConnections int) HttpClientBuilder {
	b.maxIdleConnections = maxIdleConnections
	return b
}

// Request connection timeout
func (b *clientBuilder) SetConnectionTimeout(connectionTimeout time.Duration) HttpClientBuilder {
	b.connectionTimeout = connectionTimeout
	return b
}

// Response timeout after we have sent the Request
func (b *clientBuilder) SetResponseTimeOut(responseTimeOut time.Duration) HttpClientBuilder {
	b.responseTimeOut = responseTimeOut
	return b
}
