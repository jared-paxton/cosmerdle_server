package app

import (
	"net/http"

	"github.com/jared-paxton/cosmerdle_server/pkg/game"
)

func (app *application) accessGame(w http.ResponseWriter, r *http.Request) {
	//TODO: implement logic if user has already started game and is just resuming
	gameState := game.StartGame()

	err := app.writeJSON(w, 200, gameState, "game_state")

	if err != nil {
		app.logger.Println("could not send accessGame response due to:", err)
	}
}
