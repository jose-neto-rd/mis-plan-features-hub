package controllers

import (
	"encoding/json"
	"mis-plan-features-hub/internal/core/domain/interfaces"
	"net/http"
)

type HealthCheckerController struct {
	Service interfaces.HealthChecker
}

func (c *HealthCheckerController) CheckLiveness(w http.ResponseWriter, r *http.Request) {
	if c.Service.CheckLiveness() {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"status":"error"}`))
	}
}

func (c *HealthCheckerController) CheckReadiness(w http.ResponseWriter, r *http.Request) {
	readiness := c.Service.CheckReadiness()
	response, _ := json.Marshal(readiness)

	if readiness["status"] == "error" {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Write(response)
}
