package controllers_test

import (
	"encoding/json"
	"mis-plan-features-hub/internal/controllers"
	services_mock "mis-plan-features-hub/mocks/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func MakePlanFeaturesControllerSut() (*controllers.PlanFeaturesController, *services_mock.MockPlanFeaturesService) {
	service := new(services_mock.MockPlanFeaturesService)
	controller := &controllers.PlanFeaturesController{Service: service}
	return controller, service
}

func TestPlanFeaturesController(t *testing.T) {
	t.Run("it should return 400 Bad Request when plan parameter is missing", func(t *testing.T) {
		controller, _ := MakePlanFeaturesControllerSut()

		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()

		controller.PlanFeatures(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Path parameter 'plan' is missing", response["error"])
	})

	t.Run("it should return 200 OK with features when plan exists", func(t *testing.T) {
		controller, service := MakePlanFeaturesControllerSut()

		expectedFeatures := []string{
			"Feature 1",
			"Feature 2",
		}

		service.On("Features", mock.Anything).Return(expectedFeatures)

		req := httptest.NewRequest("GET", "/plans/gold", nil)
		w := httptest.NewRecorder()

		controller.PlanFeatures(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		var response []string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, expectedFeatures, response)

		service.AssertExpectations(t)
	})

	t.Run("it should return empty array when plan has no features", func(t *testing.T) {
		controller, service := MakePlanFeaturesControllerSut()

		service.On("Features", mock.Anything).Return([]string{})

		req := httptest.NewRequest("GET", "/plans/basic", nil)
		w := httptest.NewRecorder()

		controller.PlanFeatures(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Empty(t, response)

		service.AssertExpectations(t)
	})

	t.Run("it should set correct CORS headers", func(t *testing.T) {
		controller, service := MakePlanFeaturesControllerSut()

		service.On("Features", mock.Anything).Return([]string{})

		req := httptest.NewRequest("GET", "/plans/gold", nil)
		w := httptest.NewRecorder()

		controller.PlanFeatures(w, req)

		assert.Equal(t, "GET, OPTIONS", w.Header().Get("Access-Control-Allow-Methods"))
		assert.Equal(t, "Content-Type", w.Header().Get("Access-Control-Allow-Headers"))
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		service.AssertExpectations(t)
	})
}
