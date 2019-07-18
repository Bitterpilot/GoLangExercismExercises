// Package scrabble contains functions to calculate the score on a scrable board
package scrabble

// valueList contains the letter score value
var valueList = []int{
 // A, B, C, D, E, F, G, H, I, J, K, L, M, N, O, P,  Q, R, S, T, U, V, W, X, Y,  Z, [, \, ], ^, _, `, 
	1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10, 0, 0, 0, 0, 0, 0,
 // a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p,  q, r, s, t, u, v, w, x, y,  z
	1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 10, 1, 1, 1, 1, 4, 4, 8, 4, 10,
}

// offset is the number of characters before the alphabet starts. i.e;
// Take the rune int32 value and offset so rune 'A' matchs with the 'A' score in value list
const offset = 65

// Score takes a word and returns it's score
func Score(word string) int {
	var score int

	for _, r := range word {
		score = valueList[ r - offset ] + score
	}

	return score
}
