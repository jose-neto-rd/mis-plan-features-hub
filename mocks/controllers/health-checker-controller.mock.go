package controllers_mock

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockHealthCheckerController struct {
	mock.Mock
}

func (m *MockHealthCheckerController) CheckLiveness(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}

func (m *MockHealthCheckerController) CheckReadiness(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}
