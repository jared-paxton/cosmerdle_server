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

	for i := 0; i < NumLetters; i++ {
		fmt.Print("     |")
	}
	fmt.Println()
}

// printWordRow prints each letter of the word in a "table row" like format, a letter per "cell".
func printWordRow(word string, rowNum int) error {
	if len(word) != NumLetters {
		return errors.New("wrong number of letters: cannot print row")
	}

	fmt.Printf("    %d |", rowNum)
	for i := 0; i < len(word); i++ {
		if i == len(word)-1 {
			fmt.Printf("  %c  |", word[i])
		} else {
			fmt.Printf("  %c   ", word[i])
		}
	}
	fmt.Println()

	return nil
}

// printLetterStatusRow prints the statuses of each letter in a "table row" like format, with one letter status per "cell".
// A :( means the letter is not in the word. A :/ means the letter is in the word but in a wrong position. A :) means
// the letter is in the right position.
func printLetterStatusRow(lettersStatus [NumLetters]LetterStatus) {
	row := "      |"
	endChar := " "
	for i, status := range lettersStatus {
		if i == len(lettersStatus)-1 {
			endChar = "|"
		}

		switch status {
		case Correct:
			row += "  :) "
		case DiffPosition:
			row += "  :/ "
		case NotPresent:
			row += "  :( "
		}

		row += endChar
	}
	fmt.Println(row)
}

// PrintBoard prints the game board to the terminal in an understandable format.
func printBoard(guesses []Guess) error {
	var err error
	for i, guess := range guesses {
		err = printWordRow(guess.Word, i+1)
		if err != nil {
			return err
		}
		printLetterStatusRow(guess.LettersStatus)
		fmt.Println()
	}

	printCurrMarker := true
	printedRows := len(guesses)
	for printedRows < MaxGuesses {
		printEmptyRow(printedRows+1, printCurrMarker)
		printCurrMarker = false
		printedRows++
	}
	return nil
}

func printGame(gameState *GameState, correctWord string) error {
	fmt.Println("\n-------------------------------------------")
	err := printBoard(gameState.Guesses)
	if err != nil {
		return err
	}

	fmt.Println("-------------------------------------------")
	switch gameState.CurrStatus {
	case InProgress:
		fmt.Println("           IN PROGRESS")
	case Won:
		fmt.Println("              WON!!!")
	case Lost:
		fmt.Println("             LOST :'(")
		fmt.Println("         Word is:", correctWord)
	}
	fmt.Println("-------------------------------------------")
	fmt.Println()

	return nil
}
