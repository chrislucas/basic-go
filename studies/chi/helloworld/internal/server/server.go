package server

import (
	"fmt"
	"net/http"

	"main.go/api"
	"main.go/router"
)

type Server struct {
	httpServer  *http.Server
	routes      []api.Route
	middlewares []func(http.Handler) http.Handler
}

type ApplyOption func(*Server)

func WithHandler(handler http.Handler) ApplyOption {
	return func(s *Server) {
		s.httpServer.Handler = handler
	}
}

func WithRoutes(routes []api.Route) ApplyOption {
	return func(s *Server) {
		s.routes = routes
	}
}

func WithMiddleware(middlewares []func(http.Handler) http.Handler) ApplyOption {
	return func(s *Server) {
		s.middlewares = middlewares
	}
}

func (s *Server) Start() error {
	fmt.Println("Server is running on port", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func New(addr string, applyOptions ...ApplyOption) *Server {
	server := &Server{
		httpServer: &http.Server{
			Addr: addr,
		},
	}

	for _, applyOption := range applyOptions {
		applyOption(server)
	}

	if server.routes != nil {
		router.AddRoute(server.httpServer.Handler, server.routes)
	}

	if server.middlewares != nil {
		router.AddMiddleware(server.httpServer.Handler, server.middlewares)
	}

	return server
}
