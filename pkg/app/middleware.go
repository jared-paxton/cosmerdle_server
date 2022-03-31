package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func useCORS(router *chi.Mux) {
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedHeaders: []string{"Content-Type"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		//AllowCredentials: true,
	}))
}
