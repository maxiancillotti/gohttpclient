package httpClient

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
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

	return c.client.Do(request)
}

func (c *client) setupHttpClient() {

	if c.client == nil {
		c.client = &http.Client{
			Timeout: c.connectionTimeout + c.responseTimeOut,
			Transport: &http.Transport{
				//MaxIdleConnsPerHost:   c.getMaxIdleConnections(), // Requests per minute is a good metric to set this value
				MaxIdleConnsPerHost: c.maxIdleConnections, // Requests per minute is a good metric to set this value
				//ResponseHeaderTimeout: c.getResponseTimeOut(),    // Response timeout after we have sent the Request
				ResponseHeaderTimeout: c.responseTimeOut, // Response timeout after we have sent the Request
				DialContext: (&net.Dialer{
					//Timeout:   c.getConnectionTimeout(), // Request connection timeout
					Timeout:   c.connectionTimeout, // Request connection timeout
					KeepAlive: 30 * time.Second,
				}).DialContext,
				// Acá abajo es igual a DefaulTransport
				ForceAttemptHTTP2:     true,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
		}
	}
}

/*
func (c *client) getMaxIdleConnections() int {
	if c.maxIdleConnections > 0 {
		return c.maxIdleConnections
	}
	return defaultMaxIdleConnections
}

func (c *client) getResponseTimeOut() time.Duration {
	if c.responseTimeOut > 0 {
		return c.responseTimeOut
	}
	return defaultResponseTimeOut
}

func (c *client) getConnectionTimeout() time.Duration {
	if c.connectionTimeout > 0 {
		return c.connectionTimeout
	}
	return defaultConnectionTimeout
}
*/
func (c *client) getRequestBody(body interface{}, contentType string) ([]byte, error) {

	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {

	case "application/json":
		return json.Marshal(body)

	case "application/xml":
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}

}

func (c *client) getRequestHeaders(requestHeaders http.Header) http.Header {

	result := make(http.Header)

	// Adding common headers
	for headerKey, headerVal := range c.headers {
		if len(headerVal) > 0 {
			result.Set(headerKey, headerVal[0])
		}
	}

	// Adding custom headers
	for headerKey, headerVal := range requestHeaders {
		if len(headerVal) > 0 {
			result.Set(headerKey, headerVal[0])
		}
	}

	return result
}

func (c *client) addDefaultRequestHeaders(requestHeaders *http.Header) {

	if requestHeaders.Get("Content-Type") == "" {
		requestHeaders.Set("Content-Type", "application/json")
	}
	if requestHeaders.Get("Accept") == "" {
		requestHeaders.Set("Accept", "application/json")
	}
}
