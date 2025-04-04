package services_test

import (
	"mis-plan-features-hub/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlanFeaturesService(t *testing.T) {
	t.Run("it should check feature is in a plan correctly", func(t *testing.T) {
		values := map[string][]string{
			"test": {"feature-one", "feature-two"},
		}
		service := &services.PlanFeaturesService{Values: values}

		result := service.Allowed("test", "feature-one")
		expected := true

		if result != expected {
			t.Errorf("Expecteded %v, but result was %v", expected, result)
		}
	})

	t.Run("it should check feature is not in a plan correctly", func(t *testing.T) {
		values := map[string][]string{
			"test": {"feature-one", "feature-two"},
		}
		service := &services.PlanFeaturesService{Values: values}

		result := service.Allowed("test", "feature")
		expected := false

		if result != expected {
			t.Errorf("Expecteded %v, but result was %v", expected, result)
		}
	})

	t.Run("it should return features from a plan correctly", func(t *testing.T) {
		values := map[string][]string{
			"test": {"feature-one", "feature-two"},
		}
		service := &services.PlanFeaturesService{Values: values}

		result := service.Features("test")
		expected := []string{"feature-one", "feature-two"}

		assert.Equal(t, expected, result)
	})

	t.Run("it should return features from a plan correctly even if incorrect casetext", func(t *testing.T) {
		values := map[string][]string{
			"test": {"feature-one", "feature-two"},
		}
		service := &services.PlanFeaturesService{Values: values}

		result := service.Features("teSt")
		expected := []string{"feature-one", "feature-two"}

		assert.Equal(t, expected, result)
	})

	t.Run("it should return no features from a plan who doesnt exist", func(t *testing.T) {
		values := map[string][]string{
			"test": {"feature-one", "feature-two"},
		}
		service := &services.PlanFeaturesService{Values: values}

		result := service.Features("test-unknow")
		expected := []string{}

		assert.Equal(t, expected, result)
	})
}
