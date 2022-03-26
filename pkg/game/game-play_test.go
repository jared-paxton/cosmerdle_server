package game

import (
	"testing"
)

func TestCheckGuess(t *testing.T) {
	// Correct guess
	guess1 := guess{
		word:     "BLAST",
		statuses: [numLetters]letterStatus{notPresent, notPresent, notPresent, diffPosition, notPresent},
	}
	if !guess1.isCorrect("BLAST") {
		t.Error("guess is not correct")
	}
}

func (g *guess) equals(want *guess, t *testing.T) bool {
	isEqual := true
	if g.word != want.word {
		t.Errorf("got %s for guess.Word, want %s", g.word, want.word)
		isEqual = false
	}
	for i := range want.statuses {
		if g.statuses[i] != want.statuses[i] {
			t.Errorf("got %d for g.LettersStatus[%d], want %d", g.statuses[i], i, want.statuses[i])
			isEqual = false
		}
	}
	return isEqual
}

func (gs *gameState) equals(want *gameState, t *testing.T) bool {
	isEqual := true
	for i := range want.guesses {
		isEqual = gs.guesses[i].equals(&want.guesses[i], t)
	}
	if gs.currStatus != want.currStatus {
		t.Errorf("got %d for CurrStatus, want %d", gs.currStatus, want.currStatus)
		isEqual = false
	}
	if gs.currGuess != want.currGuess {
		t.Errorf("got %d for CurrGuess, want %d", gs.currGuess, want.currGuess)
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
	want := gameState{
		guesses: []guess{
			{
				word:     userWord,
				statuses: [numLetters]letterStatus{correct, correct, correct, correct, correct},
			},
		},
		currStatus: Won,
		currGuess:  2,
	}
	gs.equals(&want, t)
}
