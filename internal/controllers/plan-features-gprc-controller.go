package controllers

import (
	"context"
	"mis-plan-features-hub/internal/core/domain/interfaces"
	plan_feature "mis-plan-features-hub/proto/plan-feature"
)

type PlanFeatureGrpcController struct {
	plan_feature.UnimplementedPlanFeatureServer
	Service interfaces.PlanFeaturesService
}

func (s *PlanFeatureGrpcController) Allowed(ctx context.Context, in *plan_feature.AllowedRequest) (*plan_feature.AllowedResponse, error) {
	allowed := s.Service.Allowed(in.Plan, in.Feature)
	return &plan_feature.AllowedResponse{Allowed: allowed}, nil
}
