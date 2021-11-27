package gohttpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Mock struct {

	// These form the Mock Key
	Method      string
	Url         string
	RequestBody string

	Error              error
	ResponseBody       []byte
	ResponseStatusCode int
}

func (m *Mock) GetResponse() (*http.Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}

	responseBodyReadCloser := ioutil.NopCloser(bytes.NewReader(m.ResponseBody))

	return &http.Response{
		StatusCode: m.ResponseStatusCode,
		Body:       responseBodyReadCloser,
	}, nil
}
