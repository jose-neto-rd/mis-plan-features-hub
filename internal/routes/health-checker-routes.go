package routes

import (
	"mis-plan-features-hub/internal/core/domain/interfaces"
)

type HealthCheckerRouter struct {
	Router     interfaces.HttpRouter
	Controller interfaces.HealthCheckerController
}

func (h HealthCheckerRouter) Init() error {
	h.Router.HandleFunc("/healthz/liveness", h.Controller.CheckLiveness)
	h.Router.HandleFunc("/healthz/readiness", h.Controller.CheckReadiness)
	return nil
}
