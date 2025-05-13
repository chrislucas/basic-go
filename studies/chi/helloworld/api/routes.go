package api

import (
	"net/http"

	"main.go/api/handlers"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
}

func newRoute(method, path string, handlerFunc func(w http.ResponseWriter, r *http.Request)) Route {
	return Route{
		Method:      method,
		Path:        path,
		HandlerFunc: handlerFunc,
	}
}

func Routes() []Route {
	return []Route{
		newRoute(http.MethodGet, "/ping", handlers.Ping),
		newRoute(http.MethodGet, "/helloworld", handlers.HelloWorld),
	}
}
