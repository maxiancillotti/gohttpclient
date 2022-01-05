package gohttpclient

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/maxiancillotti/gohttpclient/mock"
)

const (
	defaultResponseTimeOut       time.Duration = 10 * time.Second
	defaultConnectionTimeout     time.Duration = 30 * time.Second
	defaultFallbackDelay         time.Duration = 300 * time.Millisecond
	defaultExpectContinueTimeout time.Duration = 1 * time.Second
	defaultTLSHandshakeTimeout   time.Duration = 10 * time.Second

	defaultKeepAliveTime       time.Duration = 30 * time.Second
	defaultIdleConnTimeout     time.Duration = 90 * time.Second
	defaultMaxIdleConnections  int           = 100
	defaultMaxIdleConnsPerHost int           = 20
	defaultMaxConnsPerHost     int           = 512

	defaultForceAttemptHTTP2Enabled bool = true
)

func (c *client) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {

	fullHeaders := c.getRequestHeaders(headers)
	c.addDefaultRequestHeaders(&fullHeaders)

	marshaledBody, err := c.getRequestBody(body, fullHeaders.Get("Content-Type"))
	if err != nil {
		return nil, fmt.Errorf("unable to marshal body. %v", err)
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(marshaledBody))
	if err != nil {
		return nil, fmt.Errorf("unable to create new request")
	}
	request.Header = fullHeaders

	c.setupHttpClient()

	return c.httpClient.Do(request)
}

func (c *client) setupHttpClient() {

	if mock.MockupServer.IsEnabled() {
		c.httpClient = mock.MockupServer.GetClient()
		return
	}
	c.clientOnce.Do(func() {

		// Covers the entire exchange, from Dial (if a connection is not reused) to reading the body
		totalTimeout := c.builder.expectContinueTimeout + c.builder.tlsHandshakeTimeout + c.builder.connectionTimeout + c.builder.responseTimeOut

		customTransport := http.DefaultTransport.(*http.Transport).Clone()

		// Dialer contains options for connecting to an address
		customTransport.DialContext = (&net.Dialer{
			// Dial Timeout limits the time spent establishing a TCP connection (if a new one is needed)
			Timeout:       c.builder.connectionTimeout,
			KeepAlive:     c.builder.keepAliveTime,
			FallbackDelay: c.builder.fallbackDelay,
			LocalAddr:     c.builder.localAddr,
		}).DialContext

		customTransport.ResponseHeaderTimeout = c.builder.responseTimeOut
		customTransport.ExpectContinueTimeout = c.builder.expectContinueTimeout
		customTransport.TLSHandshakeTimeout = c.builder.tlsHandshakeTimeout

		customTransport.IdleConnTimeout = c.builder.idleConnTimeout
		customTransport.MaxIdleConns = c.builder.maxIdleConns
		customTransport.MaxIdleConnsPerHost = c.builder.maxIdleConnsPerHost
		customTransport.MaxConnsPerHost = c.builder.maxConnsPerHost

		customTransport.ForceAttemptHTTP2 = c.builder.forceAttemptHTTP2Enabled

		c.httpClient = &http.Client{
			Transport: customTransport,
			// CheckRedirect: func(req *http.Request, via []*http.Request) error {
			// 	return errors.New("error")
			// },
			Jar:     c.builder.cookieJar,
			Timeout: totalTimeout,
		}
	})
}
