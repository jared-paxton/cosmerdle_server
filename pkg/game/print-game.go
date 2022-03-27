package game

import (
	"errors"
	"fmt"
)

// printEmptyRow prints an empty row for the bame board to indicate the guesses the user has left.
// The next guess will be marked with an arrow.
func printEmptyRow(rowNum int, printMarker bool) {
	if printMarker {
		fmt.Print("--> ")
	} else {
		fmt.Print("    ")
	}
	fmt.Printf("%d |", rowNum)

	for i := 0; i < numLetters; i++ {
		fmt.Print("     |")
	}
	fmt.Println()
}

// printRow prints each letter of the word in a "table row" like format, a letter per "cell".
func (g *guess) printRow(rowNum int) error {
	if len(g.word) != numLetters {
		return errors.New("wrong number of letters: cannot print row")
	}

	fmt.Printf("    %d |", rowNum)
	for i := 0; i < len(g.word); i++ {
		if i == len(g.word)-1 {
			fmt.Printf("  %c  |", g.word[i])
		} else {
			fmt.Printf("  %c   ", g.word[i])
		}
	}
	fmt.Println()

	return nil
}

// printStatusesRow prints the statuses of each letter in a "table row" like format, with one letter status per "cell".
// A :( means the letter is not in the word. A :/ means the letter is in the word but in a wrong position. A :) means
// the letter is in the right position.
func printStatusesRow(lettersStatus [numLetters]letterStatus) {
	row := "      |"
	endChar := " "
	for i, status := range lettersStatus {
		if i == len(lettersStatus)-1 {
			endChar = "|"
		}

		switch status {
		case correct:
			row += "  :) "
		case diffPosition:
			row += "  :/ "
		case notPresent:
			row += "  :( "
		}

		row += endChar
	}
	fmt.Println(row)
}

// PrintBoard prints the game board to the terminal in an understandable format.
func printBoard(guesses []guess) error {
	var err error
	for i, guess := range guesses {
		err = guess.printRow(i + 1)
		if err != nil {
			return err
		}
		printStatusesRow(guess.statuses)
		fmt.Println()
	}

	printCurrMarker := true
	printedRows := len(guesses)
	for printedRows < maxGuesses {
		printEmptyRow(printedRows+1, printCurrMarker)
		printCurrMarker = false
		printedRows++
	}
	return nil
}

// printGame prints the game and game board at the current state
func (gs *gameState) printGame() error {
	fmt.Println("\n-------------------------------------------")
	err := printBoard(gs.guesses)
	if err != nil {
		return err
	}

	fmt.Println("-------------------------------------------")
	switch gs.currStatus {
	case InProgress:
		fmt.Println("           IN PROGRESS")
	case Won:
		fmt.Println("              WON!!!")
	case Lost:
		fmt.Println("             LOST :'(")
		fmt.Println("         Word is:", GetCorrectWord())
	}
	fmt.Println("-------------------------------------------")
	fmt.Println()

	return nil
}
