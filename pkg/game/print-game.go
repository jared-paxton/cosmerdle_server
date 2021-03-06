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
	if len(g.Word) != numLetters {
		return errors.New("wrong number of letters: cannot print row")
	}

	fmt.Printf("    %d |", rowNum)
	for i := 0; i < len(g.Word); i++ {
		if i == len(g.Word)-1 {
			fmt.Printf("  %c  |", g.Word[i])
		} else {
			fmt.Printf("  %c   ", g.Word[i])
		}
	}
	fmt.Println()

	return nil
}

// printStatusesRow prints the statuses of each letter in a "table row" like format, with one letter status per "cell".
// A :( means the letter is not in the word. A :/ means the letter is in the word but in a wrong position. A :) means
// the letter is in the right position.
func printStatusesRow(lettersStatus []letterStatus) {
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
func (gs *gameState) printBoard() error {
	var err error
	for i := 0; i < len(gs.Guesses); i++ {
		err = gs.Guesses[i].printRow(i + 1)
		if err != nil {
			return err
		}
		printStatusesRow(gs.Guesses[i].Statuses)
		fmt.Println()
	}

	printCurrMarker := true
	printedRows := len(gs.Guesses)
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
	err := gs.printBoard()
	if err != nil {
		return err
	}

	fmt.Println("-------------------------------------------")
	switch gs.CurrStatus {
	case InProgress:
		fmt.Println("           IN PROGRESS")
	case Won:
		fmt.Println("              WON!!!")
	case Lost:
		fmt.Println("             LOST :'(")
		fmt.Println("         Word is:", CorrectWord())
	}
	fmt.Println("-------------------------------------------")
	fmt.Println()

	return nil
}
