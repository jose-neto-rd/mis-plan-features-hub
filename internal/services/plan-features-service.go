package services

import (
	"slices"
	"strings"
)

type PlanFeaturesService struct {
	Values map[string][]string
}

func (p *PlanFeaturesService) Allowed(plan, feature string) bool {
	values, exists := p.Values[strings.ToLower(plan)]
	if !exists {
		return false
	}

	return slices.Contains(values, feature)
}

func (p *PlanFeaturesService) Features(plan string) []string {
	values, exists := p.Values[strings.ToLower(plan)]
	if !exists {
		return []string{}
	}

	return values
}
