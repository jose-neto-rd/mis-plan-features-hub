package interfaces

type HealthChecker interface {
	CheckLiveness() bool
	CheckReadiness() map[string]interface{}
}
