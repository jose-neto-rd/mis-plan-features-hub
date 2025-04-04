package controllers

import (
	"encoding/json"
	"mis-plan-features-hub/internal/core/domain/interfaces"
	"net/http"
	"strings"
)

type PlanFeaturesController struct {
	Service interfaces.PlanFeaturesService
}

func (p *PlanFeaturesController) PlanFeatures(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	path := r.URL.Path
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		http.Error(w, `{"error": "Path parameter 'plan' is missing"}`, http.StatusBadRequest)
		return
	}
	plan := parts[2]
	features := p.Service.Features(plan)

	w.WriteHeader(http.StatusOK)
	jsonData, err := json.Marshal(features)
	if err != nil {
		json.NewEncoder(w).Encode("[]")
		return
	}

	w.Write(jsonData)
}
