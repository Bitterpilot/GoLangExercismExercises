// Package accumulate performs operations on collections
package accumulate

// Accumulate takes in something to be converted and
// the function to do the conversion
func Accumulate(convert []string, with func(string) string) []string {
	var rtn []string
	for _, s := range convert {
		rtn = append(rtn, with(s))
	}
	return rtn
}
