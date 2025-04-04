package controllers

import (
	"mis-plan-features-hub/internal/core/domain/interfaces"
)

func NewHealthCheckerController(service interfaces.HealthChecker) interfaces.HealthCheckerController {
	return &HealthCheckerController{service}
}
