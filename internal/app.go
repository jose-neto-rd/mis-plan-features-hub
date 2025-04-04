package app

import (
	"mis-plan-features-hub/internal/core/domain/interfaces"
)

type App struct {
	Logger     interfaces.Logger
	Datadog    interfaces.Datadog
	HttpServer interfaces.HttpServer
	GrpcServer interfaces.GrpcServer
}

func (a App) InitApp() {
	a.Datadog.InitDatadog()
	go a.GrpcServer.InitGrpcServer()
	a.HttpServer.InitHttpServer()
}
