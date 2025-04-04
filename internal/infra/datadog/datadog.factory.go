package datadog

import (
	"mis-plan-features-hub/internal/core/domain/interfaces"
	"os"
)

func NewDatadogService(logger interfaces.Logger) *DatadogService {
	dd_Env := os.Getenv("DD_ENV")
	isEnabled := os.Getenv("isDDEnabled")
	return &DatadogService{DD_ENV: dd_Env, IsDDEnabled: isEnabled, Logger: logger}
}
