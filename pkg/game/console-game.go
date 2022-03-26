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
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

	userGuess = "SSSSS"
	MakeGuess(userGuess, &state)
	state.printGame()
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

	userGuess = "LLLLL"
	MakeGuess(userGuess, &state)
	state.printGame()
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

	userGuess = "SLSLS"
	MakeGuess(userGuess, &state)
	state.printGame()
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

	userGuess = "ADSEL"
	MakeGuess(userGuess, &state)
	state.printGame()
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

	// userGuess = "LADSS"
	// MakeGuess(userGuess, &state)
	// printGame(&state)
	// fmt.Println(state.CorrectWord)
	// fmt.Println(state.CurrStatus)
	// fmt.Println(state.CurrGuess)
	// fmt.Println(state.Guesses)

	userGuess = "ladse"
	MakeGuess(userGuess, &state)
	state.printGame()
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

}
