package interfaces

import "net/http"

type HttpServer interface {
	InitHttpServer()
	Start(address string) error
}

type HttpServerAdapter interface {
	ListenAndServe(address string, handler HttpHandler) error
}

type HttpHandler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type HttpRouterInit interface {
	Init() error
}

type HttpRouter interface {
	HandleFunc(pattern string, handlerFunc func(http.ResponseWriter, *http.Request))
}
