package logger_mock

import "github.com/stretchr/testify/mock"

type MockLogger struct {
	LastMessage string
	mock.Mock
}

func (m *MockLogger) Info(args string) {
	m.LastMessage = args
	m.Called(args)
}

func (m *MockLogger) Error(args string) {
	m.LastMessage = args
	m.Called(args)
}
