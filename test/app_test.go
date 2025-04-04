package app_test

import (
	app "mis-plan-features-hub/internal"
	datadog_mock "mis-plan-features-hub/mocks/infra/datadog"
	grpc_mock "mis-plan-features-hub/mocks/infra/grpc"
	http_mock "mis-plan-features-hub/mocks/infra/http"
	logger_mock "mis-plan-features-hub/mocks/infra/logger"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

func MakeSut() (*app.App, *logger_mock.MockLogger, *datadog_mock.MockDatadog, *http_mock.MockHttpServer, *grpc_mock.MockGrpcServer) {
	logger := new(logger_mock.MockLogger)
	datadog := new(datadog_mock.MockDatadog)
	httpServer := new(http_mock.MockHttpServer)
	grpcServer := new(grpc_mock.MockGrpcServer)

	var wg sync.WaitGroup
	wg.Add(1)

	application := &app.App{
		Logger:     logger,
		Datadog:    datadog,
		HttpServer: httpServer,
		GrpcServer: grpcServer,
	}

	datadog.On("InitDatadog").Return()
	httpServer.On("InitHttpServer").Run(func(args mock.Arguments) {
		wg.Done()
	}).Return()

	return application, logger, datadog, httpServer, grpcServer
}

func MakeSutMocked() (*app.App, *logger_mock.MockLogger, *datadog_mock.MockDatadog, *http_mock.MockHttpServer, *grpc_mock.MockGrpcServer) {
	application, logger, datadog, httpServer, grpcServer := MakeSut()
	datadog.On("InitDatadog").Return()
	httpServer.On("InitHttpServer").Return()
	grpcServer.On("InitGrpcServer").Return()
	return application, logger, datadog, httpServer, grpcServer
}

func TestApp(t *testing.T) {
	t.Run("it should call logger correctly", func(t *testing.T) {
		application, logger, _, _, _ := MakeSutMocked()

		application.InitApp()

		logger.AssertExpectations(t)
	})

	t.Run("it should call datadog correctly", func(t *testing.T) {
		application, _, datadog, _, _ := MakeSutMocked()

		application.InitApp()

		datadog.AssertExpectations(t)
	})

	t.Run("it should call httpServer correctly", func(t *testing.T) {
		application, _, _, httpServer, _ := MakeSutMocked()

		application.InitApp()

		time.Sleep(100 * time.Millisecond)

		httpServer.AssertExpectations(t)
	})

	t.Run("it should call grpcServer correctly", func(t *testing.T) {
		application, _, _, _, grpcServer := MakeSutMocked()

		application.InitApp()

		time.Sleep(100 * time.Millisecond)

		grpcServer.AssertExpectations(t)
	})
}
