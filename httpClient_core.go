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

func (hc *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {

	fullHeaders := hc.getRequestHeaders(headers)
	hc.addDefaultRequestHeaders(&fullHeaders)

	marshaledBody, err := hc.getRequestBody(body, fullHeaders.Get("Content-Type"))
	if err != nil {
		return nil, fmt.Errorf("unable to marshal body. %v", err)
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(marshaledBody))
	if err != nil {
		return nil, fmt.Errorf("unable to create new request")
	}
	request.Header = fullHeaders

	hc.setHttpClient()

	return hc.client.Do(request)
}

func (hc *httpClient) setHttpClient() {

	if hc.updateTransportSettings || hc.client == nil {
		hc.client = &http.Client{
			Timeout: hc.connectionTimeout + hc.responseTimeOut,
			Transport: &http.Transport{
				//MaxIdleConnsPerHost:   hc.getMaxIdleConnections(), // Requests per minute is a good metric to set this value
				MaxIdleConnsPerHost: hc.maxIdleConnections, // Requests per minute is a good metric to set this value
				//ResponseHeaderTimeout: hc.getResponseTimeOut(),    // Response timeout after we have sent the Request
				ResponseHeaderTimeout: hc.responseTimeOut, // Response timeout after we have sent the Request
				DialContext: (&net.Dialer{
					//Timeout:   hc.getConnectionTimeout(), // Request connection timeout
					Timeout:   hc.connectionTimeout, // Request connection timeout
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

	hc.updateTransportSettings = false
}

/*
func (hc *httpClient) getMaxIdleConnections() int {
	if hc.maxIdleConnections > 0 {
		return hc.maxIdleConnections
	}
	return defaultMaxIdleConnections
}

func (hc *httpClient) getResponseTimeOut() time.Duration {
	if hc.responseTimeOut > 0 {
		return hc.responseTimeOut
	}
	return defaultResponseTimeOut
}

func (hc *httpClient) getConnectionTimeout() time.Duration {
	if hc.connectionTimeout > 0 {
		return hc.connectionTimeout
	}
	return defaultConnectionTimeout
}
*/
func (hc *httpClient) getRequestBody(body interface{}, contentType string) ([]byte, error) {

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

func (hc *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {

	result := make(http.Header)

	// Adding common headers
	for headerKey, headerVal := range hc.headers {
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

func (hc *httpClient) addDefaultRequestHeaders(requestHeaders *http.Header) {

	if requestHeaders.Get("Content-Type") == "" {
		requestHeaders.Set("Content-Type", "application/json")
	}
	if requestHeaders.Get("Accept") == "" {
		requestHeaders.Set("Accept", "application/json")
	}
}
