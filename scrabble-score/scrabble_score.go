// Package scrabble contains functions to calculate the score on a scrable board
package scrabble

var valueList = []int{
	1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10, 0, 0, 0, 0, 0, 0,
	1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10,
}

// Score takes a word and returns it's score
func Score(word string) int {
	var score int

	// TODO: comments before submitting to exorcism
	for _, r := range word {
		score = valueList[r-65] + score
	}

	return score
}
