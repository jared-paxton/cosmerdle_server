package game

import (
	"strings"
	"testing"
)

func TestIsValidWord(t *testing.T) {
	testGuesses := [3]guess{
		{
			// Word too short
			Word:     strings.ToUpper("test"),
			Statuses: [numLetters]letterStatus{notPresent, diffPosition, diffPosition, notPresent, notPresent},
		},
		{
			// Word too long
			Word:     strings.ToUpper("testing"),
			Statuses: [numLetters]letterStatus{diffPosition, diffPosition, diffPosition, diffPosition, diffPosition},
		},
		{
			// Word not in bank
			Word:     strings.ToUpper("about"),
			Statuses: [numLetters]letterStatus{notPresent, notPresent, notPresent, notPresent, notPresent},
		},
	}

	for _, guess := range testGuesses {
		err := guess.isValid(1)
		if err == nil {
			t.Errorf("\"%s\"should be an invalid word (guess) because of : %s", guess.Word, err)
		}
	}

	// Word in bank, but test if the current guess is greater than allowed guesses
	guess := guess{
		Word:     strings.ToUpper("storm"),
		Statuses: [numLetters]letterStatus{notPresent, notPresent, notPresent, notPresent, notPresent},
	}
	err := guess.isValid(maxGuesses + 1)
	if err == nil {
		t.Error("should be an invalid word (guess) because of word isn't in bank", guess.Word)
	}
}

// TestCheckGuess tests what a guess (specifically the statuses field) should look like
// after making a guess
func TestCheckGuess(t *testing.T) {
	word := "storm"
	userWords := [4]string{"honor", "metal", "elend", "storm"}
	wantGuesses := [4]guess{
		{
			Word:     strings.ToUpper(userWords[0]),
			Statuses: [numLetters]letterStatus{notPresent, diffPosition, notPresent, notPresent, diffPosition},
		},
		{
			Word:     strings.ToUpper(userWords[1]),
			Statuses: [numLetters]letterStatus{diffPosition, notPresent, diffPosition, notPresent, notPresent},
		},
		{
			Word:     strings.ToUpper(userWords[2]),
			Statuses: [numLetters]letterStatus{notPresent, notPresent, notPresent, notPresent, notPresent},
		},
		{
			Word:     strings.ToUpper(userWords[3]),
			Statuses: [numLetters]letterStatus{correct, correct, correct, correct, correct},
		},
	}

	gs := initGameState(word)

	for i := range wantGuesses {
		gs.makeGuess(userWords[i])
		gs.Guesses[i].equals(&wantGuesses[i], t)
	}
}

// TestMakeGuess tests what the game state should be after making different guesses
func TestMakeGuess(t *testing.T) {
	word := "seons"

	userWords := [3]string{"light", "moash", word}
	wantGuesses := [3]guess{
		{
			Word:     strings.ToUpper(userWords[0]),
			Statuses: [numLetters]letterStatus{notPresent, notPresent, notPresent, notPresent, notPresent},
		},
		{
			Word:     strings.ToUpper(userWords[1]),
			Statuses: [numLetters]letterStatus{notPresent, diffPosition, notPresent, diffPosition, notPresent},
		},
		{
			Word:     strings.ToUpper(userWords[2]),
			Statuses: [numLetters]letterStatus{correct, correct, correct, correct, correct},
		},
	}
	wantGameStates := [3]gameState{
		{
			Guesses:    []guess{wantGuesses[0]},
			CurrStatus: InProgress,
			CurrGuess:  2,
		},
		{
			Guesses:    []guess{wantGuesses[0], wantGuesses[1]},
			CurrStatus: InProgress,
			CurrGuess:  3,
		},
		{
			Guesses:    []guess{wantGuesses[0], wantGuesses[1], wantGuesses[2]},
			CurrStatus: Won,
			CurrGuess:  4,
		},
	}

	gs := initGameState(word)

	for i := range wantGameStates {
		err := gs.makeGuess(userWords[i])
		if err != nil {
			t.Fatal(err.Error())
		}
		gs.equals(&wantGameStates[i], t)
	}
}

func (g *guess) equals(want *guess, t *testing.T) bool {
	isEqual := true
	if g.Word != want.Word {
		t.Errorf("got %s for guess.Word, want %s", g.Word, want.Word)
		isEqual = false
	}
	for i := range want.Statuses {
		if g.Statuses[i] != want.Statuses[i] {
			t.Errorf("got %d for g.LettersStatus[%d], want %d (guess word: %s)", g.Statuses[i], i, want.Statuses[i], g.Word)
			isEqual = false
		}
	}
	return isEqual
}

func (gs *gameState) equals(want *gameState, t *testing.T) bool {
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
