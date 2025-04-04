package httpserver

import (
	"mis-plan-features-hub/internal/core/domain/interfaces"
)

type HttpServer struct {
	Port        string
	Logger      interfaces.Logger
	Handler     interfaces.HttpHandler
	RouterInit  interfaces.HttpRouterInit
	HttpAdapter interfaces.HttpServerAdapter
}

func (h *HttpServer) Start(addr string) error {
	return h.HttpAdapter.ListenAndServe(addr, h.Handler)
}

func (h *HttpServer) InitHttpServer() {
	h.RouterInit.Init()
	h.Logger.Info("HTTP Server running on port " + h.Port)

	if err := h.Start(":" + h.Port); err != nil {
		h.Logger.Error(err.Error())
	}
}
