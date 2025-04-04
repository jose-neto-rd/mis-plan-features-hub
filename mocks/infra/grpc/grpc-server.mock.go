package grpc_mock

import "github.com/stretchr/testify/mock"

type MockGrpcServer struct {
	mock.Mock
}

func (m *MockGrpcServer) InitGrpcServer() {
	m.Called()
}
