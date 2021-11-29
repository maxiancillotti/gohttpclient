package gohttpclient

import "net/http"

func (c *client) getRequestHeaders(requestHeaders http.Header) http.Header {

	result := make(http.Header)

	// Adding common headers
	for headerKey, headerVal := range c.builder.headers {
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
