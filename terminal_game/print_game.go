package terminal_game

import (
	"errors"
	"fmt"

	g "github.com/jared-paxton/cosmerdle_server/game"
)

// printEmptyRow prints an empty row for the bame board to indicate the guesses the user has left.
// The next guess will be marked with an arrow.
func printEmptyRow(rowNum int, printMarker bool) {
	if printMarker {
		fmt.Print("--> ")
	} else {
		fmt.Print("    ")
	}
	fmt.Println(rowNum, "|     |     |     |     |     |")
}

// printWordRow prints each letter of the word in a "table row" like format, a letter per "cell".
func printWordRow(word string, rowNum int) error {
	if len(word) != g.NumLetters {
		return errors.New("wrong number of letters: cannot print row")
	}

	fmt.Printf("    %d |  %c  |  %c  |  %c  |  %c  |  %c  | \n", rowNum, word[0], word[1], word[2], word[3], word[4])
	return nil
}

// printLetterStatusRow prints the statuses of each letter in a "table row" like format, with one letter status per "cell".
// A :( means the letter is not in the word. A :/ means the letter is in the word but in a wrong position. A :) means
// the letter is in the right position.
func printLetterStatusRow(lettersStatus [g.NumLetters]g.LetterStatus) {
	row := "      |"
	for _, status := range lettersStatus {
		switch status {
		case g.SamePosition:
			row += "  :) |"
		case g.DiffPosition:
			row += "  :/ |"
		case g.NotPresent:
			row += "  :( |"
		}
	}
	fmt.Println(row)
}

// PrintBoard prints the game board to the terminal in an understandable format.
func printBoard(guesses []g.Guess) error {
	var err error
	for i, guess := range guesses {
		err = printWordRow(guess.Word, i+1)
		if err != nil {
			return err
		}
		printLetterStatusRow(guess.LettersStatus)
	}

	printCurrMarker := true
	printedRows := len(guesses)
	for printedRows < g.MaxGuesses {
		printEmptyRow(printedRows+1, printCurrMarker)
		printCurrMarker = false
		printedRows++
	}
	return nil
}

func PrintGame(gameState g.GameState, correctWord string) error {
	fmt.Println("\n-------------------------------------------")
	err := printBoard(gameState.Guesses)
    if err != nil {
        return err
    }

	fmt.Println("-------------------------------------------")
	switch gameState.CurrStatus {
	case g.InProgress:
		fmt.Println("           IN PROGRESS")
	case g.Won:
		fmt.Println("              WON!!!")
	case g.Lost:
		fmt.Println("             LOST :'(")
		fmt.Println("         Word is:", correctWord)
	}
	fmt.Println("-------------------------------------------")
	fmt.Println()

    return nil
}
