package routes

import (
	"mis-plan-features-hub/internal/controllers"
	"mis-plan-features-hub/internal/core/domain/interfaces"
)

func NewPlanFeaturesRouterInit(router interfaces.HttpRouter, service interfaces.PlanFeaturesService) interfaces.HttpRouterInit {
	controller := controllers.NewPlanFeaturesController(service)
	return &PlanFeaturesRouter{Router: router, Controller: controller}
}
