package game

import (
	"testing"
)

func TestPrintGame(t *testing.T) {
	todaysWord := "STORM"

	guess1 := guess{
		Word:     "GUESS",
		Statuses: [numLetters]letterStatus{notPresent, notPresent, notPresent, diffPosition, notPresent},
	}
	guess2 := guess{
		Word:     "STATE",
		Statuses: [numLetters]letterStatus{correct, correct, notPresent, notPresent, notPresent},
	}
	// guess3 := Guess{
	// 	Word:          "STONE",
	// 	LettersStatus: [NumLetters]Status{SamePosition, SamePosition, SamePosition, NotPresent, NotPresent},
	// }
	guess4 := guess{
		Word:     "STOMP",
		Statuses: [numLetters]letterStatus{correct, correct, correct, diffPosition, notPresent},
	}
	guess5 := guess{
		Word:     "STORM",
		Statuses: [numLetters]letterStatus{correct, correct, correct, correct, correct},
	}

	gameState := gameState{
		Guesses:    []guess{guess1, guess2, guess4, guess5},
		CurrStatus: Won,
		CurrGuess:  5,
	}

	// In this case, called just to set the correct word for the test
	initGameState(todaysWord)

	err := gameState.printGame()
	if err != nil {
		t.Fatal("printing the game failed for some reason")
	}
}
