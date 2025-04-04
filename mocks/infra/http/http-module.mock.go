package http_mock

import (
	"github.com/stretchr/testify/mock"
)

type MockHttpRouterInit struct {
	mock.Mock
}

func (m *MockHttpRouterInit) Init() error {
	m.Called()
	return nil
}
