// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"sort"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind

const (
    // Pick values for the following identifiers used by the test program.
    NaT // not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var k Kind
	return k
}

// tests for a inequality triangle
func inequality(a, b, c float64) bool {
	// collect and sort side lengths in ascending order
	lengths := []float64{a, b, c}
	sort.Float64s(lengths)

	// select the largest value for z
	z := lengths[2]
	// use the other values for x and y
	x, y := lengths[0], lengths[1]

	// test for basic inequality
	if (x + y) >= z {
		return false
	}
	return true
}
