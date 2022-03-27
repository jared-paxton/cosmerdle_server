package game

import (
	"errors"
	"strings"

	"github.com/jared-paxton/cosmerdle_server/pkg/db"
)

func GetCorrectWord() string {
	return correctWord
}

func InitGameState(word string) gameState {
	state := gameState{
		guesses:    make([]guess, maxGuesses),
		currStatus: InProgress,
		currGuess:  1,
	}
	correctWord = strings.ToUpper(word)

	return state
}

func MakeGuess(userWord string, state *gameState) error {
	userWord = strings.ToUpper(userWord)

	// Default each letter to NotPresent
	guess := guess{
		word:     userWord,
		statuses: [numLetters]letterStatus{notPresent, notPresent, notPresent, notPresent, notPresent},
	}

	err := guess.isValid(state.currGuess)
	if err != nil {
		return err
	}

	correctguess := guess.isCorrect(correctWord)
	state.addGuess(guess)
	state.updateGameStatus(correctguess)

	return nil
}

func (gs *gameState) addGuess(guess guess) {
	gs.guesses[gs.currGuess-1] = guess
}

func (gs *gameState) updateGameStatus(correctGuess bool) {
	if correctGuess {
		gs.currStatus = Won
	} else if gs.currGuess >= maxGuesses {
		gs.currStatus = Lost
	} else {
		gs.currStatus = InProgress
	}

	gs.currGuess++
}

func (g *guess) isValid(currGuess int) error {
	if len(g.word) > numLetters {
		return errors.New("exceeded max number of letters")
	} else if len(g.word) < numLetters {
		return errors.New("not enough letters in word")
	} else if !db.IsWordInBank(g.word) {
		return errors.New("word is not a part of the cosmere")
	} else if currGuess > maxGuesses {
		return errors.New("exceeded max number of guesses")
	}

	return nil
}

func (g *guess) isCorrect(correctWord string) bool {
	correctguess := true
	for i := range g.word {
		if g.word[i] == correctWord[i] {
			g.statuses[i] = correct
		} else {
			correctguess = false
		}
	}

	letterRepresented := make(map[int]bool)
	for i := range g.word {
		if g.statuses[i] != correct {
			g.checkLetterPositions(i, correctWord, letterRepresented)
		}
	}

	return correctguess
}

func (g *guess) checkLetterPositions(currPos int, correctWord string, lettersMap map[int]bool) {
	for i := range correctWord {
		if g.statuses[i] == correct || lettersMap[i] {
			continue
		}
		if g.word[currPos] == correctWord[i] {
			lettersMap[i] = true
			g.statuses[currPos] = diffPosition
			return
		}
	}
	g.statuses[currPos] = notPresent
}
