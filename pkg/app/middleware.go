package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func useCORS(router *chi.Mux) {
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedHeaders: []string{"Content-Type"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
}
