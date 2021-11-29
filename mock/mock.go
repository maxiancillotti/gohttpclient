package mock

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// The Mock struct provides a clean way to configure HTTP mocks
// based on the combination between request method, URL and request body.

type Mock struct {

	// These form the Mock Key
	Method      string
	Url         string
	RequestBody string

	Error              error
	ResponseBody       []byte
	ResponseStatusCode int
}

// GetResponse returns an *http.Response based on the mock config and http request submited.
func (m *Mock) GetResponse(request *http.Request) (*http.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	responseBodyReadCloser := ioutil.NopCloser(bytes.NewReader(m.ResponseBody))

	return &http.Response{
		StatusCode:    m.ResponseStatusCode,
		Body:          responseBodyReadCloser,
		ContentLength: int64(len(m.ResponseBody)),
		Request:       request,
	}, nil
}
