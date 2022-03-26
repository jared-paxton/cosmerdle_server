package game

import (
	"errors"
	"strings"

	"github.com/jared-paxton/cosmerdle_server/pkg/db"
)

func GetCorrectWord() string {
	return correctWord
}

func InitGameState(word string) GameState {
	state := GameState{
		Guesses:    make([]Guess, MaxGuesses),
		CurrStatus: InProgress,
		CurrGuess:  1,
	}
	correctWord = strings.ToUpper(word)

	return state
}

func MakeGuess(userWord string, state *GameState) error {
	userWord = strings.ToUpper(userWord)

	// Default each letter to NotPresent
	guess := Guess{
		Word:     userWord,
		Statuses: [NumLetters]LetterStatus{NotPresent, NotPresent, NotPresent, NotPresent, NotPresent},
	}

	err := guess.isValid(state.CurrGuess)
	if err != nil {
		return err
	}

	correctGuess := guess.isCorrect(correctWord)
	state.addGuess(guess)
	state.updateGameStatus(correctGuess)

	return nil
}

func (gs *GameState) addGuess(guess Guess) {
	gs.Guesses[gs.CurrGuess-1] = guess
}

func (gs *GameState) updateGameStatus(correctGuess bool) {
	if correctGuess {
		gs.CurrStatus = Won
	} else if gs.CurrGuess == MaxGuesses {
		gs.CurrStatus = Lost
	} else {
		gs.CurrStatus = InProgress
	}

	gs.CurrGuess++
}

func (g *Guess) isValid(currGuess int) error {
	if len(g.Word) > NumLetters {
		return errors.New("number of letters exceeds max")
	} else if len(g.Word) < NumLetters {
		return errors.New("not enough letters in word")
	} else if !db.IsWordInBank(g.Word) {
		return errors.New("word is not a part of the cosmere")
	} else if currGuess > MaxGuesses {
		return errors.New("exceeded max number of guesses")
	}

	return nil
}

func (g *Guess) isCorrect(correctWord string) bool {
	correctGuess := true
	for i := range g.Word {
		if g.Word[i] == correctWord[i] {
			g.Statuses[i] = Correct
		} else {
			correctGuess = false
		}
	}

	letterRepresented := make(map[int]bool)
	for i := range g.Word {
		if g.Statuses[i] != Correct {
			g.checkLetterPositions(i, correctWord, letterRepresented)
		}
	}

	return correctGuess
}

func (g *Guess) checkLetterPositions(currPos int, correctWord string, lettersMap map[int]bool) {
	for i := range correctWord {
		if g.Statuses[i] == Correct || lettersMap[i] {
			continue
		}
		if g.Word[currPos] == correctWord[i] {
			lettersMap[i] = true
			g.Statuses[currPos] = DiffPosition
			return
		}
	}
	g.Statuses[currPos] = NotPresent
}
