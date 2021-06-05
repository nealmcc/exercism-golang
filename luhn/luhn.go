package luhn

import (
	"strings"
)

// Valid determines if the given string is a valid account number
// according to the Luhn algorithm:
// https://en.wikipedia.org/wiki/Luhn_algorithm
func Valid(s string) bool {
	digits := strings.ReplaceAll(s, " ", "")
	if len(digits) <= 1 {
		return false
	}

	var (
		sum    int
		parity int = len(digits) % 2
	)

	for i, d := range []byte(digits) {
		d -= '0'
		if d > 9 {
			return false
		}
		if i%2 == parity {
			d += d
			if d > 9 {
				d -= 9
			}
		}
		sum += int(d)
	}

	return sum%10 == 0
}
