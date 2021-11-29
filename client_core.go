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
	defaultMaxIdleConnections int           = 5
	defaultResponseTimeOut    time.Duration = 10 * time.Second
	defaultConnectionTimeout  time.Duration = 30 * time.Second
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

		c.httpClient = &http.Client{
			Timeout: c.builder.connectionTimeout + c.builder.responseTimeOut,
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.builder.maxIdleConnections,
				ResponseHeaderTimeout: c.builder.responseTimeOut,
				DialContext: (&net.Dialer{
					Timeout:   c.builder.connectionTimeout,
					KeepAlive: 30 * time.Second,
				}).DialContext,

				// The following properties are equal to DefaulTransport
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		}
	})
}
