package datadog_mock

import (
	"github.com/stretchr/testify/mock"
)

type MockDatadog struct {
	mock.Mock
}

func (m *MockDatadog) InitDatadog() {
	m.Called()
}
