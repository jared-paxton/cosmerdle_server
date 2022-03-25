package game

import "fmt"

func MakeGuess(userWord string, state *GameState) {
	// Default
	guess := Guess{
		Word:          userWord,
		LettersStatus: [NumLetters]LetterStatus{NotPresent, NotPresent, NotPresent, NotPresent, NotPresent},
	}

	correctGuess := isGuessCorrect(&guess, state.CorrectWord)
	state.Guesses = append(state.Guesses, guess)
	updateGameStatus(correctGuess, state)
}

func updateGameStatus(correctGuess bool, state *GameState) {
	if correctGuess {
		state.CurrStatus = Won
	} else if len(state.Guesses) == MaxGuesses {
		state.CurrStatus = Lost
	} else {
		state.CurrStatus = InProgress
	}
}

func isGuessCorrect(guess *Guess, correctWord string) bool {

	correctGuess := true
	for i := range guess.Word {
		fmt.Printf("curr user letter: %s  VS  actual word letter: %s\n", string(guess.Word[i]), string(correctWord[i]))
		if guess.Word[i] == correctWord[i] {
			guess.LettersStatus[i] = Correct
		} else {
			correctGuess = false
		}

	}

	letterRepresented := make(map[int]bool)
	for i := range guess.Word {
		if guess.LettersStatus[i] == Correct {
			continue
		}
		guess.LettersStatus[i] = checkLetterPositions(guess, i, correctWord, letterRepresented)
	}

	return correctGuess
}

func checkLetterPositions(guess *Guess, currPos int, correctWord string, lettersMap map[int]bool) LetterStatus {
	for i := range correctWord {
		if guess.LettersStatus[i] == Correct || lettersMap[i] {
			continue
		}
		if guess.Word[currPos] == correctWord[i] {
			lettersMap[i] = true
			return DiffPosition
		}
	}
	return NotPresent
}

func InitGameState(correctWord string) GameState {
	state := GameState{
		Guesses:     make([]Guess, 0),
		CurrStatus:  InProgress,
		CorrectWord: correctWord,
	}

	return state
}
