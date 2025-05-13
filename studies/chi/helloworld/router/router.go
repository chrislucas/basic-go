package router

import (
	"net/http"

	"main.go/api"
	"main.go/pkg/adapter"
)

func AddRoute(mux http.Handler, routes []api.Route) {
	chiMux := adapter.AdapterChiMux(mux)
	for _, route := range routes {
		chiMux.MethodFunc(route.Method, route.Path, route.HandlerFunc)
	}
}

func AddMiddleware(mux http.Handler, middlewares []func(http.Handler) http.Handler) {
	chiMux := adapter.AdapterChiMux(mux)
	for _, m := range middlewares {
		chiMux.Use(m)
	}
}
