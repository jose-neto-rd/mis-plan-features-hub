package controllers_test

import (
	"context"
	"mis-plan-features-hub/internal/controllers"
	services_mock "mis-plan-features-hub/mocks/services"
	plan_feature "mis-plan-features-hub/proto/plan-feature"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MakePlanFeatureGrpcControllerSut() (*controllers.PlanFeatureGrpcController, *services_mock.MockPlanFeaturesService) {
	service := new(services_mock.MockPlanFeaturesService)
	controller := &controllers.PlanFeatureGrpcController{Service: service}
	return controller, service
}

func TestPlanFeatureGrpcController(t *testing.T) {
	t.Run("it should return true when feature is allowed", func(t *testing.T) {
		controller, service := MakePlanFeatureGrpcControllerSut()

		service.On("Allowed", "plan", "feature").Return(true)

		req := &plan_feature.AllowedRequest{
			Plan:    "plan",
			Feature: "feature",
		}

		response, err := controller.Allowed(context.Background(), req)

		assert.NoError(t, err)
		assert.True(t, response.Allowed)
		service.AssertExpectations(t)
	})

	t.Run("it should return false when feature is not allowed", func(t *testing.T) {
		controller, service := MakePlanFeatureGrpcControllerSut()

		service.On("Allowed", "plan", "feature").Return(false)

		req := &plan_feature.AllowedRequest{
			Plan:    "plan",
			Feature: "feature",
		}

		response, err := controller.Allowed(context.Background(), req)

		assert.NoError(t, err)
		assert.False(t, response.Allowed)
		service.AssertExpectations(t)
	})
}
