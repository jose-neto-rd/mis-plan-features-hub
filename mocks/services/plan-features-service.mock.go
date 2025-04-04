package services_mock

import (
	"github.com/stretchr/testify/mock"
)

type MockPlanFeaturesService struct {
	mock.Mock
}

func (m *MockPlanFeaturesService) Allowed(plan, feature string) bool {
	args := m.Called(plan, feature)
	return args.Get(0).(bool)
}

func (m *MockPlanFeaturesService) Features(plan string) []string {
	args := m.Called(plan)
	return args.Get(0).([]string)
}
