package app

import (
	"net/http"
	"time"

	"github.com/jared-paxton/cosmerdle_server/pkg/game"
)

func (app *application) accessGame(w http.ResponseWriter, r *http.Request) {
	//TODO: implement logic if user has already started game and is just resuming
	gameState := game.StartGame()

	addCookie(w, "userID", "123456789", 30*24*time.Hour)
	err := app.writeJSON(w, 200, gameState, "game_state")

	if err != nil {
		app.logger.Println("could not send accessGame response due to:", err)
	}
}

func (app *application) newUser(w http.ResponseWriter, r *http.Request) {
	user, err := app.appServices.userService.New()
	if err != nil {
		app.logger.Println("could not create user due to:", err)
		// TODO: add json errror writing
	}

	err = app.writeJSON(w, 200, user, "user")

	if err != nil {
		app.logger.Println("could not send newUser response due to:", err)
	}
}
