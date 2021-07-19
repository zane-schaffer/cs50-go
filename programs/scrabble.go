package programs

import (
	"fmt"
	"strings"

	"github.com/zane-schaffer/cs50/utils"
)

var points = []int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10}

func Scrabble() {

	// Get input words from both players
	word1 := utils.StringInput("Player 1: ")
	word2 := utils.StringInput("Player 2: ")

	// Score both words
	score1 := compute_score(word1)
	score2 := compute_score(word2)

	// Print the winner
	switch {
	case score1 > score2:
		fmt.Println("Player 1 is the winner")
	case score2 > score1:
		fmt.Println("Player 2 is the winner")
	case score1 == score2:
		fmt.Println("There is a tie")
	default:
		fmt.Println("Something went wrong!")
	}
}

func compute_score(word string) int {
	// Compute and return score for string
	letterPoints := 0
	word = strings.ToUpper(word)
	for _, char := range word {
		letterPoints += points[char-65]
	}
	return letterPoints
}
