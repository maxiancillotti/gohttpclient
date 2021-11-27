package gohttpclient

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {

	// Initialization
	c := &client{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "client-MaxiAncillotti")

	c.builder = &clientBuilder{}
	c.builder.SetHeaders(commonHeaders)

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	finalHeaders := c.getRequestHeaders(requestHeaders)

	// Validation
	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("Invalid value for Content-Type header")
	}

	if finalHeaders.Get("User-Agent") != "client-MaxiAncillotti" {
		t.Error("Invalid value for User-Agent header")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("Invalid value for X-Request-Id header")
	}

}

func TestAddDefaultRequestHeaders(t *testing.T) {

	// Initialization
	c := &client{}
	headers := make(http.Header)
	headers.Set("Content-Type", "application/xml")
	headers.Set("User-Agent", "client-MaxiAncillotti")

	c.builder = &clientBuilder{}
	c.builder.SetHeaders(headers)

	// Execution
	c.addDefaultRequestHeaders(&headers)

	// Validation
	if headers.Get("Content-Type") != "application/xml" {
		t.Error("Invalid value for Content-Type header")
	}

	if headers.Get("User-Agent") != "client-MaxiAncillotti" {
		t.Error("Invalid value for User-Agent header")
	}

	if headers.Get("Accept") != "application/json" {
		t.Error("Invalid value for Accept header")
	}
}

func TestGetRequestBodyEmpty(t *testing.T) {

	// Initialization
	c := &client{}
	c.builder = &clientBuilder{}
	var bodyStr string
	var bodyBytes []byte

	// Execution
	marshaledBodyFromStr, err := c.getRequestBody(bodyStr, c.builder.headers.Get("Content-Type"))

	// Validation
	if err != nil {
		t.Errorf("Cannot marshal body. %v", err)
	}
	if marshaledBodyFromStr != nil {
		t.Errorf("Marshaled body is not nil: %s", string(marshaledBodyFromStr))
	}

	// Execution
	marshaledBodyFromBytes, err := c.getRequestBody(bodyBytes, c.builder.headers.Get("Content-Type"))

	// Validation
	if err != nil {
		t.Errorf("Cannot marshal body. %v", err)
	}
	if marshaledBodyFromBytes != nil {
		t.Errorf("Marshaled body is not nil: %s", string(marshaledBodyFromBytes))
	}
}

func TestGetRequestBodyDefaultContentType(t *testing.T) {

	// Initialization
	c := &client{}
	c.builder = &clientBuilder{}
	var body struct {
		BodyField1 string `json:"body_field_1"`
		BodyField2 string `json:"body_field_2"`
	}

	body.BodyField1 = "field_1_value"
	body.BodyField2 = "field_2_value"

	// Execution
	marshaledBody, err := c.getRequestBody(body, c.builder.headers.Get("Content-Type"))

	// Validation
	if err != nil {
		t.Errorf("Cannot marshal body. %v", err)
	}

	t.Log(string(marshaledBody))
}

func TestGetRequestBodyContentTypeJSON(t *testing.T) {

	// Initialization
	c := &client{}
	c.builder = &clientBuilder{}
	var body struct {
		BodyField1 string `json:"body_field_1"`
		BodyField2 string `json:"body_field_2"`
	}

	body.BodyField1 = "field_1_value"
	body.BodyField2 = "field_2_value"

	headers := make(http.Header)
	headers.Set("Content-Type", "application/json")
	c.builder.SetHeaders(headers)

	// Execution
	marshaledBody, err := c.getRequestBody(body, c.builder.headers.Get("Content-Type"))

	// Validation
	if err != nil {
		t.Errorf("Cannot marshal body. %v", err)
	}

	t.Log(string(marshaledBody))
}

func TestGetRequestBodyContentTypeXML(t *testing.T) {

	// Initialization
	c := &client{}
	c.builder = &clientBuilder{}
	type XMLStruct struct {
		BodyField1 string `xml:"body_field_1"`
		BodyField2 string `xml:"body_field_2"`
	}

	var body XMLStruct

	body.BodyField1 = "field_1_Value"
	body.BodyField2 = "field_2_Value"

	headers := make(http.Header)
	headers.Set("Content-Type", "application/xml")
	c.builder.SetHeaders(headers)

	// Execution
	marshaledBody, err := c.getRequestBody(body, c.builder.headers.Get("Content-Type"))

	// Validation
	if err != nil {
		t.Errorf("Cannot marshal body. %v", err)
	}

	t.Log(string(marshaledBody))
}
