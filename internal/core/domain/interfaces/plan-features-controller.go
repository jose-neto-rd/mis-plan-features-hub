package interfaces

import "net/http"

type PlanFeaturesController interface {
	PlanFeatures(w http.ResponseWriter, r *http.Request)
}
