package httpcore

import (
	"net/http"
)

// HttpClient can be satisfied by a *http.Client so it can be mocked
type HttpClient interface {
	Do(request *http.Request) (*http.Response, error)
}
