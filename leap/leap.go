// Package leap calculates if leap year is true or false
package leap

// IsLeapYear check if provided int is a leap year.
func IsLeapYear(year int) bool {
	switch {
	case (year % 400) == 0:
		return true
	case (year % 100) == 0:
		return false
	case (year % 4) == 0:
		return true
	default:
		return false
	}
}
