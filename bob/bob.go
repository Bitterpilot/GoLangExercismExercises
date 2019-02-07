// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

var isLetter = regexp.MustCompile(`[a-zA-Z]+`).MatchString
var isWhiteSpace = regexp.MustCompile(`\s{2}`).MatchString

// Hey takes in a remark and responds appropriately
func Hey(remark string) string {
	switch true {
	case strings.HasSuffix(strings.TrimRight(remark, " "), "?"):
		switch {
		case (remark == strings.ToUpper(remark) && isLetter(remark)):
			// if you yell a question at him.
			return "Calm down, I know what I'm doing!"
		default:
			// if you ask him a question.
			return "Sure."

		}
	case (remark == strings.ToUpper(remark) && isLetter(remark)):
		// if you yell at him.
		return "Whoa, chill out!"
	case remark == "" || (isWhiteSpace(remark) && !isLetter(remark)):
		// if you address him without actually saying anything.
		return "Fine. Be that way!"
	default:
		// to anything else.
		return "Whatever."
	}
}
