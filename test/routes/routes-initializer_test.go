package routes_test

import (
	"mis-plan-features-hub/internal/routes"
	http_mock "mis-plan-features-hub/mocks/infra/http"
	logger_mock "mis-plan-features-hub/mocks/infra/logger"
	"testing"
)

func MakeRoutesInitializerSut() (*routes.RouterInitializer, *logger_mock.MockLogger, *http_mock.MockHttpRouterInit, *http_mock.MockHttpRouterInit) {
	logger := new(logger_mock.MockLogger)
	routerIniterOne := new(http_mock.MockHttpRouterInit)
	routerIniterTwo := new(http_mock.MockHttpRouterInit)

	routerInitializer := &routes.RouterInitializer{Logger: logger}

	return routerInitializer, logger, routerIniterOne, routerIniterTwo
}

func TestRoutesInitializerRouter(t *testing.T) {
	t.Run("it should call register correctly", func(t *testing.T) {
		routerInitializer, _, router, _ := MakeRoutesInitializerSut()

		routerInitializer.Register(router)

		expected := 1
		result := len(routerInitializer.Components)

		if result != expected {
			t.Errorf("Expecteded %v, but result was %v", expected, result)
		}
	})

	t.Run("it should logger call correctly", func(t *testing.T) {
		routerInitializer, logger, router, routerTwo := MakeRoutesInitializerSut()

		logger.On("Info", "Initializing router: MockHttpRouterInit").Return()
		logger.On("Info", "Initializing router: MockHttpRouterInit").Return()

		router.On("Init").Return()
		routerTwo.On("Init").Return()

		routerInitializer.Register(router)
		routerInitializer.Register(routerTwo)

		routerInitializer.Init()

		logger.AssertExpectations(t)
	})

	t.Run("it should all routes call init correctly", func(t *testing.T) {
		routerInitializer, logger, router, routerTwo := MakeRoutesInitializerSut()

		logger.On("Info", "Initializing router: MockHttpRouterInit").Return()
		logger.On("Info", "Initializing router: MockHttpRouterInit").Return()

		router.On("Init").Return()
		routerTwo.On("Init").Return()

		routerInitializer.Register(router)
		routerInitializer.Register(routerTwo)

		routerInitializer.Init()

		router.AssertExpectations(t)
		routerTwo.AssertExpectations(t)
	})

	t.Run("it should all routes call init correctly", func(t *testing.T) {
		routerInitializer, logger, router, routerTwo := MakeRoutesInitializerSut()

		logger.On("Info", "Initializing router: MockHttpRouterInit").Return()
		logger.On("Info", "Initializing router: MockHttpRouterInit").Return()

		router.On("Init").Return()
		routerTwo.On("Init").Return()

		routerInitializer.Register(router)
		routerInitializer.Register(routerTwo)

		routerInitializer.Init()

		router.AssertExpectations(t)
		routerTwo.AssertExpectations(t)
	})

	t.Run("it should not break when receive a nil router", func(t *testing.T) {
		routerInitializer, logger, router, _ := MakeRoutesInitializerSut()

		logger.On("Info", "Initializing router: MockHttpRouterInit").Return()

		router.On("Init").Return()

		routerInitializer.Register(nil)
		routerInitializer.Register(router)
		routerInitializer.Register(nil)

		routerInitializer.Init()

		router.AssertExpectations(t)
	})
}
