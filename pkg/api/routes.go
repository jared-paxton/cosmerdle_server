package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jared-paxton/cosmerdle_server/pkg/game"
)

func (srv *application) routes() http.Handler {
	router := chi.NewRouter()

	router.Get("/start", startGame)

	return router
}

func startGame(w http.ResponseWriter, r *http.Request) {
	gameState := game.StartGame()

	response := map[string]interface{}{
		"status": "success",
		"data":   gameState,
	}

	js, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		fmt.Println("error")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
