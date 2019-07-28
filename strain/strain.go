// Package strain filters keeps or discards items from a list depending on
// a filtering value; ie. keep odd numbers
//
// The following is taken from the testing package of the exercise:
// Collections, hm?  For this exercise in Go you'll work with slices as
// collections.  Define the following in your solution:
//
//    type Ints []int
//    type Lists [][]int
//    type Strings []string
//
// Then complete the exercise by implementing these methods:
//
//    (Ints) Keep(func(int) bool) Ints
//    (Ints) Discard(func(int) bool) Ints
//    (Lists) Keep(func([]int) bool) Lists
//    (Strings) Keep(func(string) bool) Strings
package strain

// Ints a collection of intergers.
type Ints []int

// Lists is a list of intergers.
type Lists [][]int

// Strings is a collection of strings.
type Strings []string

// Keep returns the type you want to keep.
func (i Ints) Keep(f func(int) bool) Ints {
	var n Ints
	for _, v := range i {
		if f(v) {
			n = append(n, v)
		}
	}
	return n
}

// Discard removes the type you want and returns the remainder.
func (i Ints) Discard(f func(int) bool) Ints {
	var n Ints
	for _, v := range i {
		if !f(v) {
			n = append(n, v)
		}
	}
	return n
}

// Keep returns the type you want to keep.
func (l Lists) Keep(f func([]int) bool) Lists {
	var n Lists
	for _, v := range l {
		if f(v) {
			n = append(n, v)
		}
	}
	return n
}

// Keep returns the type you want to keep.
func (s Strings) Keep(f func(string) bool) Strings {
	var n Strings
	for _, v := range s {
		if f(v) {
			n = append(n, v)
		}
	}
	return n
}
