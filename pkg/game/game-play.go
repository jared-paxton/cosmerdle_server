package game

import (
	"errors"
	"strings"
)

func getCorrectWord() string {
	return correctWord
}

func initGameState(word string) *gameState {
	state := &gameState{
		Guesses:    []guess{},
		CurrStatus: InProgress,
		CurrGuess:  1,
	}
	correctWord = strings.ToUpper(word)

	return state
}

func (gs *gameState) makeGuess(userWord string) error {
	userWord = strings.ToUpper(userWord)

	// Default each letter to NotPresent
	guess := guess{
		Word:     userWord,
		Statuses: [numLetters]letterStatus{notPresent, notPresent, notPresent, notPresent, notPresent},
	}

	err := guess.isValid(gs.CurrGuess)
	if err != nil {
		return err
	}

	correctguess := guess.isCorrect(correctWord)
	gs.addGuess(guess)
	gs.updateGameStatus(correctguess)

	return nil
}

func (gs *gameState) addGuess(guess guess) {
	gs.Guesses = append(gs.Guesses, guess)
}

func (gs *gameState) updateGameStatus(correctGuess bool) {
	if correctGuess {
		gs.CurrStatus = Won
	} else if gs.CurrGuess >= maxGuesses {
		gs.CurrStatus = Lost
	} else {
		gs.CurrStatus = InProgress
	}

	gs.CurrGuess++
}

func (g *guess) isValid(currGuess int) error {
	if len(g.Word) > numLetters {
		return errors.New("exceeded max number of letters")
	} else if len(g.Word) < numLetters {
		return errors.New("not enough letters in word")
	} else if !IsWordInBank(g.Word) {
		return errors.New("word is not a part of the cosmere")
	} else if currGuess > maxGuesses {
		return errors.New("exceeded max number of guesses")
	}

	return nil
}

func (g *guess) isCorrect(correctWord string) bool {
	correctguess := true
	for i := range g.Word {
		if g.Word[i] == correctWord[i] {
			g.Statuses[i] = correct
		} else {
			correctguess = false
		}
	}

	letterRepresented := make(map[int]bool)
	for i := range g.Word {
		if g.Statuses[i] != correct {
			g.checkLetterPositions(i, correctWord, letterRepresented)
		}
	}

	return correctguess
}

func (g *guess) checkLetterPositions(currPos int, correctWord string, lettersMap map[int]bool) {
	for i := range correctWord {
		if g.Statuses[i] == correct || lettersMap[i] {
			continue
		}
		if g.Word[currPos] == correctWord[i] {
			lettersMap[i] = true
			g.Statuses[currPos] = diffPosition
			return
		}
	}
	g.Statuses[currPos] = notPresent
}
