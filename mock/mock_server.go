package mock

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"

	"github.com/maxiancillotti/gohttpclient/httpcore"
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex

	httpClient httpcore.HttpClient

	mocks map[string]*Mock
}

var (
	MockupServer = mockServer{
		mocks:      make(map[string]*Mock),
		httpClient: &httpClientMock{},
	}
)

// Start enables the usage of mock data to execute the HTTP calls
// not doing then the actual call to the target url
func (m *mockServer) Start() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.enabled = true
}

// Stop disables the usage of mock data to execute the HTTP calls
// doing then the actual call to the target url
func (m *mockServer) Stop() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.enabled = false
}

// IsEnabled returns true after Start() is called and
// false on default state or after Stop() is called.
func (m *mockServer) IsEnabled() bool {
	return m.enabled
}

// AddMock adds mock data which can be used to execute tests without needing
// to do an actual HTTP call
func (m *mockServer) AddMock(mock Mock) {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	key := m.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	m.mocks[key] = &mock
}

// DeleteMocks deletes all mocks that could have been added by AddMock
// so you can clean mock data added before that you now don't want to use anymore
func (m *mockServer) DeleteMocks() {
	m.serverMutex.Lock()
	defer m.serverMutex.Unlock()

	m.mocks = make(map[string]*Mock)
}

// GetClient returns a mocked HTTP client
func (m *mockServer) GetClient() httpcore.HttpClient {
	return m.httpClient
}

// getMock returns a previously added mock, that matches the indicated parameters.
func (m *mockServer) getMock(method, url, body string) *Mock {

	if !m.enabled {
		//return nil
		return &Mock{
			// This will show up at testing result alerting something is wrong be it an error
			// when you expect a successful response, or an error with a different message.
			Error: fmt.Errorf("mockup server is not enabled"),
		}
	}
	if mock := m.mocks[m.getMockKey(method, url, body)]; mock != nil {
		return mock
	}
	return &Mock{
		// This will show up at testing result alerting something is wrong be it an error
		// when you expect a successful response, or an error with a different message.
		Error: fmt.Errorf("no mock matching key for method %s, url %s, and body %s", method, url, body),
	}

}

func (m *mockServer) getMockKey(method, url, body string) string {
	plainkey := fmt.Sprint(method, url, m.cleanRequestBody(body))
	hasher := md5.New()
	hasher.Write([]byte(plainkey))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (m *mockServer) cleanRequestBody(body string) string {
	body = strings.TrimSpace(body)
	if body == "" {
		return ""
	}
	body = strings.ReplaceAll(body, "\t", "")
	body = strings.ReplaceAll(body, "\n", "")
	body = strings.ReplaceAll(body, "\r", "")
	return body
}
