package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"main.go/api"
	"main.go/internal/server"

	"log"
)

func TestRouter() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	http.ListenAndServe(":3000", r)
}

func main() {
	chi := chi.NewRouter()
	routes := api.Routes()
	middlewares := []func(http.Handler) http.Handler{middleware.Logger}
	server := server.New(":3000", server.WithHandler(chi),
		server.WithRoutes(routes),
		server.WithMiddleware(middlewares),
	)
	if error := server.Start(); error != nil {
		log.Fatal(error)
	}
}
