package game

func StartGame() *gameState {
	todaysWord := GetWordOfTheDay()

	state := initGameState(todaysWord)

	// // Get Guesses from user
	// userGuess := "FASIL"
	// state.makeGuess(userGuess)

	return state
}
