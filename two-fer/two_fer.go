// Package twofer has one function ShareWith
package twofer

import (
	"strings"
)

// ShareWith takes a string and returns
// One for x, one for me.
// where
// x = a string that is parsed into the function
// empty string = "you,"
func ShareWith(name string) string {
	ul := []string{"One for", name + ",", "one for me."}
	if name == "" {
		ul[1] = "you,"
	}
	return strings.Join(ul, " ")
}
