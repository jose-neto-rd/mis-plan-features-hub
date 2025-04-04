package httpserver_mock

import (
	"github.com/stretchr/testify/mock"
)

type MockHttpServer struct {
	mock.Mock
}

func (m *MockHttpServer) ListenAndServe(address string) error {
	m.Called(address)
	return nil
}

func (m *MockHttpServer) InitHttpServer() {
	m.Called()
}
