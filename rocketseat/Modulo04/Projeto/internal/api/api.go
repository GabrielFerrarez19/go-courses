package api

import (
	"net/http"

	"Projeto/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type apiResponse struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewHandler(repository repository.Repository) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/url", func(r chi.Router) {
			r.Post("/shorten", handleShortenURL(repository))
			r.Get("/{code}", handleGetShortenedURL(repository))
		})
	})
	return r
}
