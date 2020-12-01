// Package raindrops solves the Exercism exercise
package raindrops

import (
	"strconv"
	"strings"
)

// Convert returns raindrop sounds based on the input number
func Convert(n int) string {
	d := []string{}

	if hasFactor(n, 3) {
		d = append(d, "Pling")
	}

	if hasFactor(n, 5) {
		d = append(d, "Plang")
	}

	if hasFactor(n, 7) {
		d = append(d, "Plong")
	}

	if len(d) == 0 {
		return strconv.Itoa(n)
	}

	return strings.Join(d, "")
}

func hasFactor(n, f int) bool {
	return (n % f) == 0
}
