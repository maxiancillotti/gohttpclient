package httpClient

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

func (hc *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {

	client := http.Client{}

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

	return client.Do(request)
}

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
