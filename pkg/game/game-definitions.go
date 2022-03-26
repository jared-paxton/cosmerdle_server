package game

const MaxGuesses = 6
const NumLetters = 5

type LetterStatus int

const (
	Default      LetterStatus = 0
	Correct      LetterStatus = 1
	DiffPosition LetterStatus = 2
	NotPresent   LetterStatus = 3
)

type Guess struct {
	Word     string
	Statuses [NumLetters]LetterStatus
}

type Status int

const (
	InProgress Status = 0
	Won        Status = 1
	Lost       Status = 2
)

var correctWord string

type GameState struct {
	Guesses    []Guess
	CurrStatus Status
	CurrGuess  int
}
