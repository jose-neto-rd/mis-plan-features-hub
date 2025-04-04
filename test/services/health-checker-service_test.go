package services_test

import (
	"mis-plan-features-hub/internal/services"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func MakeHealthSut() *services.HealthCheckerService {
	return &services.HealthCheckerService{InstanceId: "test-instance"}
}

func TestHealthCheckerService(t *testing.T) {
	t.Run("it should return true for CheckLiveness", func(t *testing.T) {
		service := MakeHealthSut()
		expected := true
		result := service.CheckLiveness()

		if result != expected {
			t.Errorf("result '%t', expected '%t'", result, expected)
		}
	})

	t.Run("it should return correct structure for CheckReadiness", func(t *testing.T) {
		service := MakeHealthSut()
		os.Setenv("INSTANCE_ID", "test-instance")
		defer os.Unsetenv("INSTANCE_ID")

		result := service.CheckReadiness()
		expected := map[string]interface{}{
			"status":      "ok",
			"instance_id": "test-instance",
			"dependencies": map[string]string{
				"pubsub": "ok",
			},
		}

		assert.Equal(t, expected, result, "CheckReadiness() should return expected structure")
	})

	t.Run("it should return 'local' when INSTANCE_ID is not set", func(t *testing.T) {
		service := &services.HealthCheckerService{}
		os.Unsetenv("INSTANCE_ID")

		result := service.CheckReadiness()
		expected := "local"

		if result["instance_id"] != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})
}
