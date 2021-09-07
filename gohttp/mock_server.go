package gohttp

import "sync"

var (
	mockupServer = mockServer{
		mocks: make(map[string]*Mock),
	}
)
type mockServer struct {
	enabled bool
	serverMutex sync.Mutex

	mocks map[string]*Mock
}

func StartMockServer() {
	// concurrence - one server, multiple routines
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

	key := mockupServer.getMockKey(mock.Method, mock.Url,  mock.RequestBody)
	mockupServer.mocks[key] = &mock
}

func (m *mockServer) getMockKey(method, url, body string) string {
	if !m.enabled{
		return nil
	}
	return method + url + body
}

func (m *mockServer) getMock(method, url, body string) *Mock {
	return m.mocks[m.getMockKey(method,url,body)]
}
