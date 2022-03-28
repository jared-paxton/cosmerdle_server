package game

import (
	"testing"
)

func TestPrintGame(t *testing.T) {
	todaysWord := "STORM"

	guess1 := guess{
		word:     "GUESS",
		statuses: [numLetters]letterStatus{notPresent, notPresent, notPresent, diffPosition, notPresent},
	}
	guess2 := guess{
		word:     "STATE",
		statuses: [numLetters]letterStatus{correct, correct, notPresent, notPresent, notPresent},
	}
	// guess3 := Guess{
	// 	Word:          "STONE",
	// 	LettersStatus: [NumLetters]Status{SamePosition, SamePosition, SamePosition, NotPresent, NotPresent},
	// }
	guess4 := guess{
		word:     "STOMP",
		statuses: [numLetters]letterStatus{correct, correct, correct, diffPosition, notPresent},
	}
	guess5 := guess{
		word:     "STORM",
		statuses: [numLetters]letterStatus{correct, correct, correct, correct, correct},
	}

	gameState := gameState{
		guesses:    []guess{guess1, guess2, guess4, guess5},
		currStatus: Won,
		currGuess:  1,
	}

	// In this case, called just to set the correct word for the test
	initGameState(todaysWord)

	err := gameState.printGame()
	if err != nil {
		t.Fatal("printing the game failed for some reason")
	}
}
