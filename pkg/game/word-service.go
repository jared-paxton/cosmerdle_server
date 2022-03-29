package game

import (
	"math/rand"
	"strings"
	"time"
)

func getRandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

var wordBank = []string{
	"storm",
	"elend",
	"light",
	"moash",
	"seons",
	"steel",
	"metal",
	"spren",
	"honor",
	"oaths",
}

func GetWordOfTheDay() string {
	randNum := getRandomNumber(0, len(wordBank)-1)
	return wordBank[randNum]
}

// TODO: implement this using database after database is implemented
func IsWordInBank(word string) bool {
	for _, w := range wordBank {
		if strings.ToUpper(w) == word {
			return true
		}
	}

	return false
}
