package api

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"Projeto/internal/repository"

	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type getShortenedURLResponse struct {
	FullURL string `json:"full_url"`
}

func handleGetShortenedURL(repository repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		fullURL, err := repository.GetFullURL(r.Context(), code)
		fmt.Println(fullURL)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				sendJSON(w, apiResponse{Error: "code not found"}, http.StatusNotFound)
				return
			}
			slog.Error("failed to get code", "error", err)
			sendJSON(w, apiResponse{Error: "something went wrong"}, http.StatusInternalServerError)
			return
		}

		sendJSON(w, apiResponse{Data: getShortenedURLResponse{FullURL: fullURL}}, http.StatusOK)
	}
}
