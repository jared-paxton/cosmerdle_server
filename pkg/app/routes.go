package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	// Middleware
	useCORS(router)

	router.Get("/", app.accessGame)

	return router
}
