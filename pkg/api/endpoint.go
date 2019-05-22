package api

import (
	"github.com/go-chi/chi"
)

func RegisterEndpoints(router *chi.Mux) {
	router.Get("/health", handleHealthCheck)
}
