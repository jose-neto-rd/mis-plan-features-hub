package httpserver

import (
	"mis-plan-features-hub/internal/core/domain/interfaces"
	"mis-plan-features-hub/internal/infra/logger"
	"mis-plan-features-hub/internal/routes"
	"net/http"
	"os"
)

func NewHttpServer(logger *logger.Logger, planFeaturesService interfaces.PlanFeaturesService) interfaces.HttpServer {
	handler := http.NewServeMux()
	httpServerAdapter := NewHttpServerAdapter()
	healthCheckerRouter := routes.NewHealthCheckerRouterInit(handler)
	planFeaturesRouter := routes.NewPlanFeaturesRouterInit(handler, planFeaturesService)

	routerInit := &routes.RouterInitializer{Logger: logger}
	routerInit.Register(healthCheckerRouter)
	routerInit.Register(planFeaturesRouter)

	port := os.Getenv("SERVICE_PORT")
	if port == "" {
		port = "3000"
	}
	return &HttpServer{Port: port, Logger: logger, RouterInit: routerInit, Handler: handler, HttpAdapter: httpServerAdapter}
}
