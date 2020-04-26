package handler

import (
	"github.com/go-chi/chi"
	"net/http"
)

func NewHealthHandler(router *chi.Mux) {
	router.Get("/health", health)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("hello everybody"))
}
