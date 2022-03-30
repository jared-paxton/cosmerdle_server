package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jared-paxton/cosmerdle_server/pkg/game"
)

func (app *application) startGame(w http.ResponseWriter, r *http.Request) {
	gameState := game.StartGame()

	response := map[string]interface{}{
		"status": "success",
		"data":   gameState,
	}

	var js []byte
	var err error
	if app.config.Env == "development" {
		js, err = json.MarshalIndent(response, "", "    ")
	} else {
		js, err = json.Marshal(response)
	}

	if err != nil {
		fmt.Println("error")
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
