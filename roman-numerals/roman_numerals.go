// Package romannumerals converts Arabic numerals to Roman numerals
package romannumerals

import (
	"bytes"
	"fmt"
)

// ToRomanNumeral takes a decimal number and converts it to a roman numeral.
func ToRomanNumeral(a int) (string, error) {
	if a > 3000 || a < 1 {
		return "", fmt.Errorf("zero, negative and above 3000 numbers aren't allowed. Got: %d", a)
	}
	var buffer bytes.Buffer

	for a > 0 {
		switch {
		case a >= 1000:
			a -= 1000
			buffer.WriteString("M")

		case a >= 900:
			a -= 900
			buffer.WriteString("CM")

		case a >= 500:
			a -= 500
			buffer.WriteString("D")

		case a >= 400:
			a -= 400
			buffer.WriteString("CD")

		case a >= 100:
			a -= 100
			buffer.WriteString("C")

		case a >= 90:
			a -= 90
			buffer.WriteString("XC")

		case a >= 50:
			a -= 50
			buffer.WriteString("L")

		case a >= 40:
			a -= 40
			buffer.WriteString("XL")

		case a >= 10:
			a -= 10
			buffer.WriteString("X")

		case a >= 9:
			a -= 9
			buffer.WriteString("IX")

		case a >= 5:
			a -= 5
			buffer.WriteString("V")

		case a >= 4:
			a -= 4
			buffer.WriteString("IV")

		case a >= 1:
			a--
			buffer.WriteString("I")
		}
	}

	return buffer.String(), nil
}
