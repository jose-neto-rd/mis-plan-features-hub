package services_mock

import (
	"github.com/stretchr/testify/mock"
)

type MockHealthCheckerService struct {
	mock.Mock
}

func (m *MockHealthCheckerService) CheckLiveness() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockHealthCheckerService) CheckReadiness() map[string]interface{} {
	args := m.Called()
	return args.Get(0).(map[string]interface{})
}
