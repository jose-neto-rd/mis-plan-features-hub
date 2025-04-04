package httprouter_mock

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

type MockHttpRouter struct {
	mock.Mock
}

func (m *MockHttpRouter) HandleFunc(address string, handlerFunc func(http.ResponseWriter, *http.Request)) {
	m.Called(address, handlerFunc)
}
