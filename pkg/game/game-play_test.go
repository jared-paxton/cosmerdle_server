package game

import (
	"fmt"
	"testing"
)

func TestCheckGuess(t *testing.T) {
	guess1 := Guess{
		Word:          "BLAST",
		LettersStatus: [NumLetters]LetterStatus{NotPresent, NotPresent, NotPresent, DiffPosition, NotPresent},
	}

	isGuessCorrect(&guess1, "blaht")
	fmt.Println("new guess", guess1.Word)
}
