package http_mock

import (
	"mis-plan-features-hub/internal/core/domain/interfaces"

	"github.com/stretchr/testify/mock"
)

type MockHttpServerAdapter struct {
	mock.Mock
}

func (m *MockHttpServerAdapter) ListenAndServe(addr string, handler interfaces.HttpHandler) error {
	m.Called(addr, handler)
	return nil
}
