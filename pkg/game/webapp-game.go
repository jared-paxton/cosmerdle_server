package game

func StartGame() *gameState {
	todaysWord := GetWordOfTheDay()

	state := initGameState(todaysWord)

	// // Get Guesses from user
	// userGuess := "STORM"
	// state.makeGuess(userGuess)

	return state
}
