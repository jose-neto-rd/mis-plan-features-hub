package routes_test

import (
	"mis-plan-features-hub/internal/routes"
	controllers_mock "mis-plan-features-hub/mocks/controllers"
	http_mock "mis-plan-features-hub/mocks/infra/http"
	"testing"

	"github.com/stretchr/testify/mock"
)

func MakePlanFeaturesRouterSut() (*routes.PlanFeaturesRouter, *http_mock.MockHttpRouter, *controllers_mock.MockPlanFeaturesController) {
	router := new(http_mock.MockHttpRouter)
	controller := new(controllers_mock.MockPlanFeaturesController)

	healthRouter := &routes.PlanFeaturesRouter{Router: router, Controller: controller}

	router.On("HandleFunc", "/features/", mock.AnythingOfType("func(http.ResponseWriter, *http.Request)")).Return()

	return healthRouter, router, controller
}

func TestPlanFeaturesRouter(t *testing.T) {
	t.Run("it should return call router correctly", func(t *testing.T) {
		healthRouter, router, _ := MakePlanFeaturesRouterSut()

		healthRouter.Init()

		router.AssertExpectations(t)
	})
}
