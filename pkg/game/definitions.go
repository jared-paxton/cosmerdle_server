package game

const maxGuesses = 6
const numLetters = 5

type letterStatus int

const (
	initial      letterStatus = 0
	correct      letterStatus = 1
	diffPosition letterStatus = 2
	notPresent   letterStatus = 3
)

type guess struct {
	Word     string
	Statuses []letterStatus
}

type gameStatus int

const (
	InProgress gameStatus = 0
	Won        gameStatus = 1
	Lost       gameStatus = 2
)

// correctWord stores the word of the day
var correctWord string

type gameState struct {
	Guesses    []guess    `json:"guesses"`
	CurrStatus gameStatus `json:"curr_status"`
	CurrGuess  int        `json:"curr_guess"`
}
