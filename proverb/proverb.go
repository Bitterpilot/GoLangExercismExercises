// Package proverb writes proverbs
package proverb

import "fmt"

// Proverb takes a slice of strings containing nouns for a proverb.
// When parssing in the nouns "nail", "shoe" and "horse" the proverb is
// returned as follows:
//
// "For want of a nail the shoe was lost.",
// "For want of a shoe the horse was lost.",
// "And all for the want of a nail."
func Proverb(rhyme []string) []string {
	var proverb []string
	if len(rhyme) == 0 {
		return proverb
	}
	firstPiece := rhyme[0]
	for i, piece := range rhyme {
		if i != len(rhyme)-1 {
			line := fmt.Sprintf("For want of a %s the %s was lost.", piece, rhyme[i+1])
			proverb = append(proverb, line)
		} else {
			line := fmt.Sprintf("And all for the want of a %s.", firstPiece)
			proverb = append(proverb, line)
		}
	}
	return proverb
}
