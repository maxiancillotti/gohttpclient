package httpClient

import (
	"net/http"
	"testing"
)

func TestGetRequestHeaders(t *testing.T) {

	// Initialization
	hc := &httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "httpClient-MaxiAncillotti")
	hc.SetHeaders(commonHeaders)

	// Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")

	finalHeaders := hc.getRequestHeaders(requestHeaders)

	// Validation
	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("Invalid value for Content-Type header")
	}

	if finalHeaders.Get("User-Agent") != "httpClient-MaxiAncillotti" {
		t.Error("Invalid value for User-Agent header")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("Invalid value for X-Request-Id header")
	}

}

func TestAddDefaultRequestHeaders(t *testing.T) {

	// Initialization
	hc := &httpClient{}
	headers := make(http.Header)
	headers.Set("Content-Type", "application/xml")
	headers.Set("User-Agent", "httpClient-MaxiAncillotti")
	hc.SetHeaders(headers)

	// Execution
	hc.addDefaultRequestHeaders(&headers)

	// Validation
	if headers.Get("Content-Type") != "application/xml" {
		t.Error("Invalid value for Content-Type header")
	}

	if headers.Get("User-Agent") != "httpClient-MaxiAncillotti" {
		t.Error("Invalid value for User-Agent header")
	}

	if headers.Get("Accept") != "application/json" {
		t.Error("Invalid value for Accept header")
	}
}

func TestGetRequestBodyNilBody(t *testing.T) {

	// Initialization
	hc := &httpClient{}
	var body string

	// Execution
	marshaledBody, err := hc.getRequestBody(body, hc.headers.Get("Content-Type"))

	// Validation
	if err != nil {
		t.Errorf("Cannot marshal body. %v", err)
	}
	if marshaledBody != nil {
		t.Errorf("Marshaled body is not nil")
	}
}

func TestGetRequestBodyDefaultContentType(t *testing.T) {

	// Initialization
	hc := &httpClient{}
	var body struct {
		BodyField1 string `json:"body_field_1"`
		BodyField2 string `json:"body_field_2"`
	}

	body.BodyField1 = "field_1_value"
	body.BodyField2 = "field_2_value"

	// Execution
	marshaledBody, err := hc.getRequestBody(body, hc.headers.Get("Content-Type"))

	// Validation
	if err != nil {
		t.Errorf("Cannot marshal body. %v", err)
	}

	t.Log(string(marshaledBody))
}

func TestGetRequestBodyContentTypeJSON(t *testing.T) {

	// Initialization
	hc := &httpClient{}
	var body struct {
		BodyField1 string `json:"body_field_1"`
		BodyField2 string `json:"body_field_2"`
	}

	body.BodyField1 = "field_1_value"
	body.BodyField2 = "field_2_value"

	headers := make(http.Header)
	headers.Set("Content-Type", "application/json")
	hc.SetHeaders(headers)

	// Execution
	marshaledBody, err := hc.getRequestBody(body, hc.headers.Get("Content-Type"))

	// Validation
	if err != nil {
		t.Errorf("Cannot marshal body. %v", err)
	}

	t.Log(string(marshaledBody))
}

func TestGetRequestBodyContentTypeXML(t *testing.T) {

	// Initialization
	hc := &httpClient{}

	type XMLStruct struct {
		BodyField1 string `xml:"body_field_1"`
		BodyField2 string `xml:"body_field_2"`
	}

	var body XMLStruct

	body.BodyField1 = "field_1_Value"
	body.BodyField2 = "field_2_Value"

	headers := make(http.Header)
	headers.Set("Content-Type", "application/xml")
	hc.SetHeaders(headers)

	// Execution
	marshaledBody, err := hc.getRequestBody(body, hc.headers.Get("Content-Type"))

	// Validation
	if err != nil {
		t.Errorf("Cannot marshal body. %v", err)
	}

	t.Log(string(marshaledBody))
}
