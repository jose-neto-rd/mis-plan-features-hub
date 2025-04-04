package interfaces

import "net/http"

type HealthCheckerController interface {
	CheckLiveness(w http.ResponseWriter, r *http.Request)
	CheckReadiness(w http.ResponseWriter, r *http.Request)
}
