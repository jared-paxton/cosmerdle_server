package console

import (
	"testing"

	g "github.com/jared-paxton/cosmerdle_server/pkg/game"
)

func TestPrintGame(t *testing.T) {
	todaysWord := "STORM"

	guess1 := g.Guess{
		Word:          "GUESS",
		LettersStatus: [g.NumLetters]g.LetterStatus{g.NotPresent, g.NotPresent, g.NotPresent, g.DiffPosition, g.NotPresent},
	}
	guess2 := g.Guess{
		Word:          "STATE",
		LettersStatus: [g.NumLetters]g.LetterStatus{g.Correct, g.Correct, g.NotPresent, g.NotPresent, g.NotPresent},
	}
	// guess3 := g.Guess{
	// 	Word:          "STONE",
	// 	LettersStatus: [g.NumLetters]g.Status{g.SamePosition, g.SamePosition, g.SamePosition, g.NotPresent, g.NotPresent},
	// }
	guess4 := g.Guess{
		Word:          "STOMP",
		LettersStatus: [g.NumLetters]g.LetterStatus{g.Correct, g.Correct, g.Correct, g.DiffPosition, g.NotPresent},
	}
	guess5 := g.Guess{
		Word:          "STORM",
		LettersStatus: [g.NumLetters]g.LetterStatus{g.Correct, g.Correct, g.Correct, g.Correct, g.Correct},
	}

	gameState := g.GameState{
		Guesses:     []g.Guess{guess1, guess2, guess4, guess5},
		CurrStatus:  g.Won,
		CorrectWord: todaysWord,
	}

	err := printGame(&gameState)
	if err != nil {
		t.Fatal("printing the game failed for some reason")
	}
}
