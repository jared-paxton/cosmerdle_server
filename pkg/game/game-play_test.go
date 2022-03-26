package game

import (
	"testing"
)

func TestCheckGuess(t *testing.T) {
	// Correct guess
	guess1 := Guess{
		Word:     "BLAST",
		Statuses: [NumLetters]LetterStatus{NotPresent, NotPresent, NotPresent, DiffPosition, NotPresent},
	}
	if !guess1.isCorrect("BLAST") {
		t.Error("guess is not correct")
	}
}

func (g *Guess) equals(want *Guess, t *testing.T) bool {
	isEqual := true
	if g.Word != want.Word {
		t.Errorf("got %s for guess.Word, want %s", g.Word, want.Word)
		isEqual = false
	}
	for i := range want.Statuses {
		if g.Statuses[i] != want.Statuses[i] {
			t.Errorf("got %d for g.LettersStatus[%d], want %d", g.Statuses[i], i, want.Statuses[i])
			isEqual = false
		}
	}
	return isEqual
}

func (gs *GameState) equals(want *GameState, t *testing.T) bool {
	isEqual := true
	for i := range want.Guesses {
		isEqual = gs.Guesses[i].equals(&want.Guesses[i], t)
	}
	if gs.CurrStatus != want.CurrStatus {
		t.Errorf("got %d for CurrStatus, want %d", gs.CurrStatus, want.CurrStatus)
		isEqual = false
	}
	if gs.CurrGuess != want.CurrGuess {
		t.Errorf("got %d for CurrGuess, want %d", gs.CurrGuess, want.CurrGuess)
		isEqual = false
	}

	return isEqual
}

/*
type GameState struct {
	Guesses    []Guess
	CurrStatus Status
	CurrGuess  int
}

type Guess struct {
	Word          string
	LettersStatus [NumLetters]LetterStatus
}

*/

func TestMakeGuess(t *testing.T) {

	gs := InitGameState("blast")

	// var tests = []struct {
	//     a, b int
	//     want int
	// }{
	//     {0, 1, 0},
	//     {1, 0, 0},
	//     {2, -2, -2},
	//     {0, -1, -1},
	//     {-1, 0, -1},
	// }

	// Correct guess
	userWord := "BLAST"
	MakeGuess(userWord, &gs)
	want := GameState{
		Guesses: []Guess{
			{
				Word:     userWord,
				Statuses: [NumLetters]LetterStatus{Correct, Correct, Correct, Correct, Correct},
			},
		},
		CurrStatus: Won,
		CurrGuess:  2,
	}
	gs.equals(&want, t)
}
