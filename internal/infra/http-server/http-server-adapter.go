package httpserver

import (
	"mis-plan-features-hub/internal/core/domain/interfaces"
	"net/http"
)

type HttpServerAdapter struct{}

func (h *HttpServerAdapter) ListenAndServe(addr string, hander interfaces.HttpHandler) error {
	return http.ListenAndServe(addr, hander)
}
