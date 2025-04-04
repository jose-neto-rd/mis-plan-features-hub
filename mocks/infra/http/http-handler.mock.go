package http_mock

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockHttpHandler struct {
	mock.Mock
}

func (m *MockHttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Called(w, r)
}
