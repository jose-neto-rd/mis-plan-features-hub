package controllers_mock

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockPlanFeaturesController struct {
	mock.Mock
}

func (m *MockPlanFeaturesController) PlanFeatures(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}
