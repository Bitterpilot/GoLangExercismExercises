// Package raindrops converts a number to a string, the contents of which
// depends on the number's factors.
package raindrops

import "strconv"

// Convert a number to a string, the contents of which depend on the number's
// factors.
// If the number has 3 as a factor, output 'Pling'.
// If the number has 5 as a factor, output 'Plang'.
// If the number has 7 as a factor, output 'Plong'.
// If the number does not have 3, 5, or 7 as a factor, just pass the number's
//   digits straight through.
func Convert(n int) string {
	var output string

	if n%3 == 0 {
		output += "Pling"
	}
	if n%5 == 0 {
		output += "Plang"
	}
	if n%7 == 0 {
		output += "Plong"
	}
	if n%3 != 0 && n%5 != 0 && n%7 != 0 {
		output += strconv.Itoa(n)
	}

	return output
}
