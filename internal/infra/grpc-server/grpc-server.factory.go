package grpcserver

import (
	"mis-plan-features-hub/internal/controllers"
	"mis-plan-features-hub/internal/core/domain/interfaces"
	"mis-plan-features-hub/internal/infra/logger"
	"os"
)

func NewGrpcServer(logger *logger.Logger, planFeaturesService interfaces.PlanFeaturesService) interfaces.GrpcServer {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50050"
	}
	controller := &controllers.PlanFeatureGrpcController{Service: planFeaturesService}
	return &GrpcServer{Port: port, Logger: logger, Controller: controller}
}
