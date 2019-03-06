// Package triangle takes values from a triangle and identifies its type.
package triangle

import (
	"math"
	"sort"
)

// Kind holds the kind of triangle as an int.
type Kind int

// Declare diffrent types of triangle
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch {
	case infinite(a, b, c) || negative(a, b, c) || inequality(a, b, c):
		k = NaT
	case a == b && b == c:
		k = Equ
	case a == b || b == c || a == c:
		k = Iso
	default:
		k = Sca
	}
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

// tests for negative or zero sides
func negative(a, b, c float64) bool {
	switch {
	case a <= 0 || b <= 0 || c <= 0:
		return true
	default:
		return false
	}
}

// tests for infinite sides, the negative infinity case would be handled by the
// func negative()
func infinite(a, b, c float64) bool {
	switch {
	case a == math.Inf(1) || b == math.Inf(1) || c == math.Inf(1):
		return true
	default:
		return false
	}
}
