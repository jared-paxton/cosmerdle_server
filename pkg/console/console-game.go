package console

import (
	"fmt"

	"github.com/jared-paxton/cosmerdle_server/pkg/db"
	"github.com/jared-paxton/cosmerdle_server/pkg/game"
)

func StartGame() {
	todaysWord := db.GetWordOfTheDay()

	state := game.InitGameState(todaysWord)

	// Get Guesses from user
	userGuess := "SALES"
	game.MakeGuess(userGuess, &state)
	printGame(&state, todaysWord)
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

	userGuess = "SSSSS"
	game.MakeGuess(userGuess, &state)
	printGame(&state, todaysWord)
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

	userGuess = "LLLLL"
	game.MakeGuess(userGuess, &state)
	printGame(&state, todaysWord)
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

	userGuess = "SLSLS"
	game.MakeGuess(userGuess, &state)
	printGame(&state, todaysWord)
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

	userGuess = "ADSEL"
	game.MakeGuess(userGuess, &state)
	printGame(&state, todaysWord)
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

	// userGuess = "LADSS"
	// game.MakeGuess(userGuess, &state)
	// printGame(&state)
	// fmt.Println(state.CorrectWord)
	// fmt.Println(state.CurrStatus)
	// fmt.Println(state.CurrGuess)
	// fmt.Println(state.Guesses)

	userGuess = "ladse"
	game.MakeGuess(userGuess, &state)
	printGame(&state, todaysWord)
	fmt.Println(todaysWord)
	fmt.Println(state.CurrStatus)
	fmt.Println(state.Guesses)

}
