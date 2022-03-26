package game

import (
	"fmt"

	"github.com/jared-paxton/cosmerdle_server/pkg/db"
)

func StartGame() {
	todaysWord := db.GetWordOfTheDay()

	state := InitGameState(todaysWord)

	// Get Guesses from user
	userGuess := "SALES"
	MakeGuess(userGuess, &state)
	printGame(&state, todaysWord)
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

	userGuess = "SSSSS"
	MakeGuess(userGuess, &state)
	printGame(&state, todaysWord)
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

	userGuess = "LLLLL"
	MakeGuess(userGuess, &state)
	printGame(&state, todaysWord)
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

	userGuess = "SLSLS"
	MakeGuess(userGuess, &state)
	printGame(&state, todaysWord)
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

	userGuess = "ADSEL"
	MakeGuess(userGuess, &state)
	printGame(&state, todaysWord)
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
	printGame(&state, todaysWord)
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

}
