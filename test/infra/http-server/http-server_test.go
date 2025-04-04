package httpserver_test

import (
	httpserver "mis-plan-features-hub/internal/infra/http-server"
	http_mock "mis-plan-features-hub/mocks/infra/http"
	logger_mock "mis-plan-features-hub/mocks/infra/logger"
	"testing"
)

func MakeHttpServerSut() (*httpserver.HttpServer, *logger_mock.MockLogger, *http_mock.MockHttpRouterInit, *http_mock.MockHttpServerAdapter) {
	port := "3000"
	logger := &logger_mock.MockLogger{}
	handler := &http_mock.MockHttpHandler{}
	routerInit := &http_mock.MockHttpRouterInit{}
	serverAdapter := &http_mock.MockHttpServerAdapter{}

	routerInit.On("Init").Return()
	logger.On("Info", "HTTP Server running on port "+port).Return()
	serverAdapter.On("ListenAndServe", ":"+port, handler).Return()

	httpServer := &httpserver.HttpServer{
		Port:        port,
		Logger:      logger,
		RouterInit:  routerInit,
		Handler:     handler,
		HttpAdapter: serverAdapter,
	}

	return httpServer, logger, routerInit, serverAdapter
}

func TestHttpServer(t *testing.T) {
	t.Run("it should initialize the HTTP server correctly", func(t *testing.T) {
		httpServer, logger, routerInit, serverAdapter := MakeHttpServerSut()

		httpServer.InitHttpServer()

		logger.AssertExpectations(t)
		routerInit.AssertExpectations(t)
		serverAdapter.AssertExpectations(t)
	})

	t.Run("it should listen and serve on the specified port", func(t *testing.T) {
		httpServer, _, _, _ := MakeHttpServerSut()

		err := httpServer.Start(":3000")
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
	})
}
