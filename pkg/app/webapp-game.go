package app

import (
	"github.com/jared-paxton/cosmerdle_server/pkg/db"
	"github.com/jared-paxton/cosmerdle_server/pkg/game"
)

func StartGame() {
	todaysWord := db.GetWordOfTheDay()

	state := game.InitGameState(todaysWord)

	// Get Guesses from user
	userGuess := "FASIL"
	game.MakeGuess(userGuess, &state)
}
