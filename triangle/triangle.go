// Package triangle takes values from a triangle and identifies its type.
package triangle

import (
	"sort"
)

type Kind
// Kind holds the kind of triangle as an int.

// Declare diffrent types of triangle
const (
    NaT // not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
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
