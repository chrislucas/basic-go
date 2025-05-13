package adapter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func AdapterChiMux(handler http.Handler) *chi.Mux {
	return handler.(*chi.Mux)
}
