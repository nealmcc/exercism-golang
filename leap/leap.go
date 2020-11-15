// Package leap holds the solution for the 'leap year' exercise
package leap

// IsLeapYear determines if the given year is a leap year
func IsLeapYear(y int) bool {
	if y%400 == 0 {
		return true
	}

	if y%100 == 0 {
		return false
	}

	return y%4 == 0
}
