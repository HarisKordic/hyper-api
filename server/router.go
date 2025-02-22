package server

import (
	"hyper-api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter creates a new router
func NewRouter() *chi.Mux {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Register routes
	router.Route("/api", func(r chi.Router) {
		r.Mount("/dashboard", handlers.DashboardRoutes())
	})

	return router
}
