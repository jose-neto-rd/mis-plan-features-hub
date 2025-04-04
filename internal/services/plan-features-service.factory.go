package services

import (
	"mis-plan-features-hub/internal/core/domain/constants"
	"mis-plan-features-hub/internal/core/domain/interfaces"
)

func NewPlanFeaturesService() interfaces.PlanFeaturesService {
	values := make(map[string][]string)
	values[constants.BasicPlan] = constants.BasicFeatures
	values[constants.ProPlan] = constants.ProFeatures
	values[constants.AdvancedPlan] = constants.AdvancedFeatures

	return &PlanFeaturesService{Values: values}
}
