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

func StartMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = true
}

func StopMockServer() {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	mockupServer.enabled = false
}

func AddMock(mock Mock) {
	mockupServer.serverMutex.Lock()
	defer mockupServer.serverMutex.Unlock()

	key := mockupServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	mockupServer.mocks[key] = &mock
}

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
