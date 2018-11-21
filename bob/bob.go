// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"regexp"
	"strings"
)

var isLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

// Hey takes in a remark and responds appropriatly
func Hey(remark string) string {
	switch true {
	case (strings.HasSuffix(remark, "?") && remark == strings.ToUpper(remark)):
		// if you yell a question at him.
		return "Calm down, I know what I'm doing!"
	case strings.HasSuffix(remark, "?"):
		// if you ask him a question.
		return "Sure."
	case (remark == strings.ToUpper(remark) && !(isLetter(remark))):
		// if you yell at him.
		return "Whoa, chill out!"
	// case gibberish:
	// if you address him without actually saying anything.
	// return "Fine. Be that way!"
	default:
		// to anything else.
		return "Whatever."
	}
}
