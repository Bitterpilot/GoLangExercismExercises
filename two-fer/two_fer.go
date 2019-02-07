// Package twofer has one function ShareWith
package twofer

// ShareWith takes a string and returns
// One for x, one for me.
// where
// x = a string that is parsed into the function
// empty string = "you,"
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
