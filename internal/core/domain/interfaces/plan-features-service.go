package interfaces

type PlanFeaturesService interface {
	Allowed(plan, feature string) bool
	Features(plan string) []string
}
