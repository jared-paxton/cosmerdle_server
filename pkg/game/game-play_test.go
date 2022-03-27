package game

import (
	"strings"
	"testing"
)

func TestIsValidWord(t *testing.T) {
	testGuesses := [3]guess{
		{
			// Word too short
			word:     strings.ToUpper("test"),
			statuses: [numLetters]letterStatus{notPresent, diffPosition, diffPosition, notPresent, notPresent},
		},
		{
			// Word too long
			word:     strings.ToUpper("testing"),
			statuses: [numLetters]letterStatus{diffPosition, diffPosition, diffPosition, diffPosition, diffPosition},
		},
		{
			// Word not in bank
			word:     strings.ToUpper("about"),
			statuses: [numLetters]letterStatus{notPresent, notPresent, notPresent, notPresent, notPresent},
		},
	}

	for _, guess := range testGuesses {
		err := guess.isValid(1)
		if err == nil {
			t.Errorf("\"%s\"should be an invalid word (guess) because of : %s", guess.word, err)
		}
	}

	// Word in bank, but test if the current guess is greater than allowed guesses
	guess := guess{
		word:     strings.ToUpper("storm"),
		statuses: [numLetters]letterStatus{notPresent, notPresent, notPresent, notPresent, notPresent},
	}
	err := guess.isValid(maxGuesses + 1)
	if err == nil {
		t.Error("should be an invalid word (guess) because of word isn't in bank", guess.word)
	}
}

// TestCheckGuess tests what a guess (specifically the statuses field) should look like
// after making a guess
func TestCheckGuess(t *testing.T) {
	word := "storm"
	userWords := [4]string{"honor", "metal", "value", "storm"}
	wantGuesses := [4]guess{
		{
			word:     strings.ToUpper(userWords[0]),
			statuses: [numLetters]letterStatus{notPresent, diffPosition, notPresent, notPresent, diffPosition},
		},
		{
			word:     strings.ToUpper(userWords[1]),
			statuses: [numLetters]letterStatus{diffPosition, notPresent, diffPosition, notPresent, notPresent},
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
	word := "seons"

	userWords := [3]string{"light", "moash", word}
	wantGuesses := [3]guess{
		{
			word:     strings.ToUpper(userWords[0]),
			statuses: [numLetters]letterStatus{notPresent, notPresent, notPresent, notPresent, notPresent},
		},
		{
			word:     strings.ToUpper(userWords[1]),
			statuses: [numLetters]letterStatus{notPresent, diffPosition, notPresent, diffPosition, notPresent},
		},
		{
			word:     strings.ToUpper(userWords[2]),
			statuses: [numLetters]letterStatus{correct, correct, correct, correct, correct},
		},
	}
	wantGameStates := [3]gameState{
		{
			guesses:    []guess{wantGuesses[0]},
			currStatus: InProgress,
			currGuess:  2,
		},
		{
			guesses:    []guess{wantGuesses[0], wantGuesses[1]},
			currStatus: InProgress,
			currGuess:  3,
		},
		{
			guesses:    []guess{wantGuesses[0], wantGuesses[1], wantGuesses[2]},
			currStatus: Won,
			currGuess:  4,
		},
	}

	gs := InitGameState(word)

	for i := range wantGameStates {
		err := MakeGuess(userWords[i], &gs)
		if err != nil {
			t.Fatal(err.Error())
		}
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
	if gs.currGuess != want.currGuess {
		t.Errorf("got %d for CurrGuess, want %d", gs.currGuess, want.currGuess)
		isEqual = false
	}

	return isEqual
}
