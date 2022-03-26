package game

import (
	"testing"
)

func TestPrintGame(t *testing.T) {
	todaysWord := "STORM"

	guess1 := Guess{
		Word:          "GUESS",
		LettersStatus: [NumLetters]LetterStatus{NotPresent, NotPresent, NotPresent, DiffPosition, NotPresent},
	}
	guess2 := Guess{
		Word:          "STATE",
		LettersStatus: [NumLetters]LetterStatus{Correct, Correct, NotPresent, NotPresent, NotPresent},
	}
	// guess3 := Guess{
	// 	Word:          "STONE",
	// 	LettersStatus: [NumLetters]Status{SamePosition, SamePosition, SamePosition, NotPresent, NotPresent},
	// }
	guess4 := Guess{
		Word:          "STOMP",
		LettersStatus: [NumLetters]LetterStatus{Correct, Correct, Correct, DiffPosition, NotPresent},
	}
	guess5 := Guess{
		Word:          "STORM",
		LettersStatus: [NumLetters]LetterStatus{Correct, Correct, Correct, Correct, Correct},
	}

	gameState := GameState{
		Guesses:    []Guess{guess1, guess2, guess4, guess5},
		CurrStatus: Won,
		CurrGuess:  1,
	}

	err := printGame(&gameState, todaysWord)
	if err != nil {
		t.Fatal("printing the game failed for some reason")
	}
}
