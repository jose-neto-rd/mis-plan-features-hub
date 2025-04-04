package controllers_test

import (
	"encoding/json"
	"mis-plan-features-hub/internal/controllers"
	services_mock "mis-plan-features-hub/mocks/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MakeHealthControllerSut() (*controllers.HealthCheckerController, *services_mock.MockHealthCheckerService) {
	service := new(services_mock.MockHealthCheckerService)

	controller := &controllers.HealthCheckerController{Service: service}

	return controller, service
}

func TestHealthController(t *testing.T) {
	t.Run("it should return 200 OK when CheckLiveness returns true", func(t *testing.T) {
		controller, service := MakeHealthControllerSut()

		service.On("CheckLiveness").Return(true)

		req := httptest.NewRequest("GET", "/healthz/liveness", nil)
		w := httptest.NewRecorder()

		controller.CheckLiveness(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "ok", response["status"])

		service.AssertExpectations(t)
	})

	t.Run("it should return 500 Internal Server Error when CheckLiveness returns false", func(t *testing.T) {
		controller, service := MakeHealthControllerSut()

		service.On("CheckLiveness").Return(false)

		req := httptest.NewRequest("GET", "/healthz/liveness", nil)
		w := httptest.NewRecorder()

		controller.CheckLiveness(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])

		service.AssertExpectations(t)
	})

	t.Run("it should return 200 OK and readiness status when CheckReadiness returns a valid map", func(t *testing.T) {
		controller, service := MakeHealthControllerSut()

		service.On("CheckReadiness").Return(map[string]interface{}{"status": "ok"})

		req := httptest.NewRequest("GET", "/healthz/readiness", nil)
		w := httptest.NewRecorder()

		controller.CheckReadiness(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "ok", response["status"])

		service.AssertExpectations(t)
	})

	t.Run("it should return 500 Internal Server Error when CheckReadiness returns error status", func(t *testing.T) {
		controller, service := MakeHealthControllerSut()

		service.On("CheckReadiness").Return(map[string]interface{}{"status": "error"})

		req := httptest.NewRequest("GET", "/healthz/readiness", nil)
		w := httptest.NewRecorder()

		controller.CheckReadiness(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "error", response["status"])

		service.AssertExpectations(t)
	})
}
