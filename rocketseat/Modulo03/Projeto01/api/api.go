package api

import (
	"encoding/json"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal json data", "error", err)
		sendJSON(
			w,
			Response{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}
}

func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/shorten", handlePost(db))
	r.Get("/{code}", handleGet(db))

	return r
}

type PostBory struct {
	URL string `json:"url"`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func handlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var bory PostBory

		if err := json.NewDecoder(r.Body).Decode(&bory); err != nil {
			sendJSON(w, Response{Error: "invalid bory"}, http.StatusUnprocessableEntity)
			return
		}

		if _, err := url.Parse(bory.URL); err != nil {
			sendJSON(w, Response{Error: "invalid URL passed"}, http.StatusBadRequest)
		}
		code := genCode()
		db[code] = bory.URL
		sendJSON(w, Response{Data: code}, http.StatusCreated)
	}
}

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"

func genCode() string {
	const n = 8
	byts := make([]byte, n)

	for i := range 8 {
		byts[i] = characters[rand.IntN(len(characters))]
	}
	return string(byts)
}

func handleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		data, ok := db[code]
		if !ok {
			sendJSON(w, Response{Error: "URL nao encontrada"}, http.StatusNotFound)
			return
		}
		http.Redirect(w, r, data, http.StatusPermanentRedirect)
	}
}
