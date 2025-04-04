package datadog

import (
	"mis-plan-features-hub/internal/core/domain/interfaces"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type DatadogService struct {
	DD_ENV      string
	IsDDEnabled string
	Logger      interfaces.Logger
}

func (d DatadogService) InitDatadog() {
	isDDEnabled := d.Contains([]string{"homologacao", "projeto", "producao"}, d.DD_ENV) && d.IsDDEnabled == "true"
	if isDDEnabled {
		tracer.Start()
		defer tracer.Stop()
		d.Logger.Info("Monitoring is running")
	} else {
		d.Logger.Info("Monitoring is not running")
	}
}

func (d DatadogService) Contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
