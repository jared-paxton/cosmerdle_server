package game

import (
	"strings"
	"testing"
)

// TestCheckGuess tests what a guess (specifically the statuses field) should look like
// after making a guess
func TestCheckGuess(t *testing.T) {
	word := "TRAIT"
	userWords := [4]string{"ptttp", "aittr", "blown", "trait"}
	wantGuesses := [4]guess{
		{
			word:     strings.ToUpper(userWords[0]),
			statuses: [numLetters]letterStatus{notPresent, diffPosition, diffPosition, notPresent, notPresent},
		},
		{
			word:     strings.ToUpper(userWords[1]),
			statuses: [numLetters]letterStatus{diffPosition, diffPosition, diffPosition, diffPosition, diffPosition},
		},
		{
			word:     strings.ToUpper(userWords[2]),
			statuses: [numLetters]letterStatus{notPresent, notPresent, notPresent, notPresent, notPresent},
		},
		{
			word:     strings.ToUpper(userWords[3]),
			statuses: [numLetters]letterStatus{correct, correct, correct, correct, correct},
		},
	}

	gs := InitGameState(word)

	for i := range wantGuesses {
		MakeGuess(userWords[i], &gs)
		gs.guesses[i].equals(&wantGuesses[i], t)
	}
}

// TestMakeGuess tests what the game state should be after making different guesses
func TestMakeGuess(t *testing.T) {
	word := "blast"

	userWords := [3]string{"chore", "bares", word}
	wantGuesses := [3]guess{
		{
			word:     strings.ToUpper(userWords[0]),
			statuses: [numLetters]letterStatus{notPresent, notPresent, notPresent, notPresent, notPresent},
		},
		{
			word:     strings.ToUpper(userWords[1]),
			statuses: [numLetters]letterStatus{correct, diffPosition, notPresent, notPresent, diffPosition},
		},
		{
			word:     strings.ToUpper(userWords[2]),
			statuses: [numLetters]letterStatus{correct, correct, correct, correct, correct},
		},
	}
	wantGameStates := [3]gameState{
		{
			guesses:    []guess{},
			currStatus: InProgress,
			nextGuess:  2,
		},
		{
			guesses:    []guess{wantGuesses[0], wantGuesses[1]},
			currStatus: InProgress,
			nextGuess:  3,
		},
		{
			guesses:    []guess{wantGuesses[0], wantGuesses[1], wantGuesses[2]},
			currStatus: Won,
			nextGuess:  4,
		},
	}

	gs := InitGameState(word)

	for i := range wantGameStates {
		MakeGuess(userWords[i], &gs)
		gs.equals(&wantGameStates[i], t)
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
			t.Errorf("got %d for g.LettersStatus[%d], want %d (guess word: %s)", g.statuses[i], i, want.statuses[i], g.word)
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
	if gs.nextGuess != want.nextGuess {
		t.Errorf("got %d for CurrGuess, want %d", gs.nextGuess, want.nextGuess)
		isEqual = false
	}

	return isEqual
}
