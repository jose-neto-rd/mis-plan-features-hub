package routes

import (
	"mis-plan-features-hub/internal/controllers"
	"mis-plan-features-hub/internal/core/domain/interfaces"
	"mis-plan-features-hub/internal/services"
)

func NewHealthCheckerRouterInit(router interfaces.HttpRouter) interfaces.HttpRouterInit {
	service := services.NewHealthCheckerService()
	controller := controllers.NewHealthCheckerController(service)
	return &HealthCheckerRouter{Router: router, Controller: controller}
}
