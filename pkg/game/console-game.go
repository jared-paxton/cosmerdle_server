package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jared-paxton/cosmerdle_server/pkg/db"
)

func printGameIntro() {
	fmt.Println()
	fmt.Println("**********************************************************************")
	fmt.Println()
	fmt.Println("                              COSMERDLE                               ")
	fmt.Println()
	fmt.Println("It's WORDLE, but only with words from Brandon Sanderson's Cosmere!")
	fmt.Println("Proper nouns like charaters and places are valid words.")
	fmt.Println("Remember, to only guess 5 letter words...")
	fmt.Println()
	fmt.Println("**********************************************************************")
	fmt.Println()
}

func StartConsoleGame() {
	printGameIntro()

	todaysWord := db.GetWordOfTheDay()
	gs := InitGameState(todaysWord)

	for gs.currStatus == InProgress {
		fmt.Printf("Enter guess #%d (or enter 9 to quit): ", gs.currGuess)
		var reader = bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Ooops... something went wrong with your input. Try again.")
			continue
		}

		word := strings.TrimSuffix(input, "\n")
		fmt.Println(word)
		fmt.Println(len(word))

		var num int
		num, err = strconv.Atoi(word)
		if err == nil && num == 9 {
			break
		}

		err = MakeGuess(word, &gs)
		if err != nil {
			fmt.Println("error: ", err.Error())
		} else {
			err = gs.printGame()
		}

		if err != nil {
			fmt.Println("error: ", err)
		}
	}
}
