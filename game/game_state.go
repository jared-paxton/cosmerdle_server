package game

const MaxGuesses = 6
const NumLetters = 5

type LetterStatus int

const (
	SamePosition LetterStatus = 0
	DiffPosition LetterStatus = 1
	NotPresent   LetterStatus = 2
)

type Guess struct {
	Word          string
	LettersStatus [NumLetters]LetterStatus
}

type Status int

const (
	InProgress Status = 0
	Won        Status = 1
	Lost       Status = 2
)

type GameState struct {
	Guesses    []Guess
	CurrStatus Status
}
