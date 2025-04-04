package routes_test

import (
	"mis-plan-features-hub/internal/routes"
	controllers_mock "mis-plan-features-hub/mocks/controllers"
	http_mock "mis-plan-features-hub/mocks/infra/http"
	"testing"

	"github.com/stretchr/testify/mock"
)

func MakeHealthRouterSut() (*routes.HealthCheckerRouter, *http_mock.MockHttpRouter, *controllers_mock.MockHealthCheckerController) {
	router := new(http_mock.MockHttpRouter)
	controller := new(controllers_mock.MockHealthCheckerController)

	healthRouter := &routes.HealthCheckerRouter{Router: router, Controller: controller}

	router.On("HandleFunc", "/healthz/liveness", mock.AnythingOfType("func(http.ResponseWriter, *http.Request)")).Return()
	router.On("HandleFunc", "/healthz/readiness", mock.AnythingOfType("func(http.ResponseWriter, *http.Request)")).Return()

	return healthRouter, router, controller
}

func TestHealthRouter(t *testing.T) {
	t.Run("it should return call router correctly", func(t *testing.T) {
		healthRouter, router, _ := MakeHealthRouterSut()

		healthRouter.Init()

		router.AssertExpectations(t)
	})
}
