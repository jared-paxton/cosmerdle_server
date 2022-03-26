package game

import (
	"fmt"

	"github.com/jared-paxton/cosmerdle_server/pkg/db"
)

func StartConsoleGame() {
	todaysWord := db.GetWordOfTheDay()

	state := InitGameState(todaysWord)

	// Get Guesses from user
	userGuess := "SALES"
	MakeGuess(userGuess, &state)
	state.printGame()
	fmt.Println(todaysWord)
	fmt.Println(state.currStatus)
	fmt.Println(state.guesses)

	userGuess = "SSSSS"
	MakeGuess(userGuess, &state)
	state.printGame()
	fmt.Println(todaysWord)
	fmt.Println(state.currStatus)
	fmt.Println(state.guesses)

	userGuess = "LLLLL"
	MakeGuess(userGuess, &state)
	state.printGame()
	fmt.Println(todaysWord)
	fmt.Println(state.currStatus)
	fmt.Println(state.guesses)

	userGuess = "SLSLS"
	MakeGuess(userGuess, &state)
	state.printGame()
	fmt.Println(todaysWord)
	fmt.Println(state.currStatus)
	fmt.Println(state.guesses)

	userGuess = "ADSEL"
	MakeGuess(userGuess, &state)
	state.printGame()
	fmt.Println(todaysWord)
	fmt.Println(state.currStatus)
	fmt.Println(state.guesses)

	userGuess = "LADSS"
	MakeGuess(userGuess, &state)
	state.printGame()
	fmt.Println(todaysWord)
	fmt.Println(state.currStatus)
	fmt.Println(state.guesses)

	userGuess = "ladse"
	MakeGuess(userGuess, &state)
	state.printGame()
	fmt.Println(todaysWord)
	fmt.Println(state.currStatus)
	fmt.Println(state.guesses)

}
