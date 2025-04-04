package routes

import (
	"mis-plan-features-hub/internal/core/domain/interfaces"
)

type PlanFeaturesRouter struct {
	Router     interfaces.HttpRouter
	Controller interfaces.PlanFeaturesController
}

func (p PlanFeaturesRouter) Init() error {
	p.Router.HandleFunc("/features/", p.Controller.PlanFeatures)
	return nil
}
