package services

import (
	"mis-plan-features-hub/internal/core/domain/interfaces"
	"os"
)

func NewHealthCheckerService() interfaces.HealthChecker {
	instanceID := os.Getenv("INSTANCE_ID")
	return &HealthCheckerService{InstanceId: instanceID}
}
