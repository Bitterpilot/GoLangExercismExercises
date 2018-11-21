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
	ul := []string{"One for", "", "one for me."}
	switch name {
	case "":
		ul[1] = "you,"
		return strings.Join(ul, " ")
	default:
		ul[1] = name + ","
		return strings.Join(ul, " ")
	}
}
