package gohttpclient

import (
	"net/http"
	"time"
)

// Methods returning interface can concatenate method calls
type ClientBuilder interface {

	// SetHeaders: set common headers to use during all client life
	SetHeaders(headers http.Header) ClientBuilder

	// SetMaxIdleConnections sets http.Transaport's MaxIdleConnsPerHost property.
	// Requests per minute is a good metric to set this value.
	// Default is 5.
	SetMaxIdleConnections(maxIdleConnections int) ClientBuilder

	// SetConnectionTimeout sets the request connection timeout.
	// Default is 10 seconds.
	SetConnectionTimeout(connectionTimeout time.Duration) ClientBuilder

	// SetResponseTimeOut sets the response timeout after we have sent the Request.
	// Default is 30 seconds.
	SetResponseTimeOut(requestTimeOut time.Duration) ClientBuilder

	// Build sets the previously configured parameters into our HTTP client
	// and returns it to perform the desired HTTP calls.
	Build() Client
}

type clientBuilder struct {
	maxIdleConnections int
	connectionTimeout  time.Duration
	responseTimeOut    time.Duration

	headers http.Header
}

// NewBuiler returns a ClientBuilder that you can configure to build
// finally your HTTP client.
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

func (b *clientBuilder) SetMaxIdleConnections(maxIdleConnections int) ClientBuilder {
	b.maxIdleConnections = maxIdleConnections
	return b
}

func (b *clientBuilder) SetConnectionTimeout(connectionTimeout time.Duration) ClientBuilder {
	b.connectionTimeout = connectionTimeout
	return b
}

func (b *clientBuilder) SetResponseTimeOut(responseTimeOut time.Duration) ClientBuilder {
	b.responseTimeOut = responseTimeOut
	return b
}
