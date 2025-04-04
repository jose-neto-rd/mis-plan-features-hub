package httpserver

import "mis-plan-features-hub/internal/core/domain/interfaces"

func NewHttpServerAdapter() interfaces.HttpServerAdapter {
	return &HttpServerAdapter{}
}
