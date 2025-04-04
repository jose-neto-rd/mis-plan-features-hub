package app

import (
	"mis-plan-features-hub/internal/infra/datadog"
	grpcserver "mis-plan-features-hub/internal/infra/grpc-server"
	httpserver "mis-plan-features-hub/internal/infra/http-server"
	"mis-plan-features-hub/internal/infra/logger"
	"mis-plan-features-hub/internal/services"
)

func NewApp() *App {
	logger := logger.NewLogger()
	planFeaturesService := services.NewPlanFeaturesService()
	datadog := datadog.NewDatadogService(logger)
	httpServer := httpserver.NewHttpServer(logger, planFeaturesService)
	grpcServer := grpcserver.NewGrpcServer(logger, planFeaturesService)

	return &App{
		Logger:     logger,
		Datadog:    datadog,
		HttpServer: httpServer,
		GrpcServer: grpcServer,
	}
}
