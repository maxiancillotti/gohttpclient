package gohttpclient

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"
)

type mockServer struct {
	enabled     bool
	serverMutex sync.Mutex

	mocks map[string]*Mock
}

var (
	mockupServer = mockServer{
		mocks: make(map[string]*Mock),
	}
)

// StartMockServer enables the usage of mock data to execute the HTTP calls
// not doing then the actual call to the target url
func StartMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = true
}

// StopMockServer disables the usage of mock data to execute the HTTP calls
// doing then the actual call to the target url
func StopMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = false
}

// AddMock adds mock data which can be used to execute tests without needing
// to do an actual HTTP call
func AddMock(mock Mock) {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	key := mockupServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	mockupServer.mocks[key] = &mock
}

// RemoveAllMocks removes all mocks that could have been added by AddMock
// so you can clean mock data added before that you now don't want to use anymore
func RemoveAllMocks() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.mocks = make(map[string]*Mock)
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

func (m *mockServer) getMock(method, url, body string) *Mock {
	if !m.enabled {
		return nil
	}
	if mock := m.mocks[m.getMockKey(method, url, body)]; mock != nil {
		return mock
	}
	return &Mock{
		Error: fmt.Errorf("no mock matching key for method %s, url %s, and body %s", method, url, body),
	}
}
