// Package scrabble contains functions to calculate the score on a scrable board
package scrabble

import "strings"

// Score takes a word and returns it's score
func Score(word string) int {
	var score int
	word = strings.ToUpper(word)

	for _, r := range word {
		// r is a rune.
		// Ranging over a string with a for loop returns runes
		// A rune is defined in Go as an int32 the
		// below reflect statement will return int32.
		// fmt.Println(reflect.TypeOf(r))
		score = lookupValue(r) + score
	}

	return score
}

// lookupValue takes a letter and returns it's value
func lookupValue(r rune) int {
	// A, E, I, O, U, L, N, R, S, T       1,
	// D, G                               2,
	// B, C, M, P                         3,
	// F, H, V, W, Y                      4,
	// K                                  5,
	// J, X                               8,
	// Q, Z                               10,
	valueList := map[rune]int{
		'A': 1, 'E': 1, 'I': 1, 'O': 1, 'U': 1, 'L': 1, 'N': 1, 'R': 1, 'S': 1, 'T': 1,
		'D': 2, 'G': 2,
		'B': 3, 'C': 3, 'M': 3, 'P': 3,
		'F': 4, 'H': 4, 'V': 4, 'W': 4, 'Y': 4,
		'K': 5,
		'J': 8, 'X': 8,
		'Q': 10, 'Z': 10,
	}
	value := valueList[r]
	return value
}
