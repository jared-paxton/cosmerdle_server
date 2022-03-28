package game


func StartGame() {
	todaysWord := GetWordOfTheDay()

	state := initGameState(todaysWord)

	// Get Guesses from user
	userGuess := "FASIL"
	makeGuess(userGuess, &state)
}
