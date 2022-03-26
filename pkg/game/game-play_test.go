package game

import (
	"fmt"
	"testing"
)


func TestCheckGuess(t *testing.T) {
	// Correct guess
	guess1 := Guess{
		Word:          "BLAST",
		LettersStatus: [NumLetters]LetterStatus{NotPresent, NotPresent, NotPresent, DiffPosition, NotPresent},
	}

	guess1.isCorrect("blaht")
	fmt.Println("new guess", guess1.Word)
}
