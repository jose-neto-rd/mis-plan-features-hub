package controllers

import (
	"mis-plan-features-hub/internal/core/domain/interfaces"
)

func NewPlanFeaturesController(service interfaces.PlanFeaturesService) interfaces.PlanFeaturesController {
	return &PlanFeaturesController{service}
}
