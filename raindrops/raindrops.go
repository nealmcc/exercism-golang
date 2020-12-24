// Package raindrops solves the Exercism exercise
package raindrops

import (
	"strconv"
	"strings"
)

// Convert returns raindrop sounds based on the input number
func Convert(n int) string {
	d := []string{}

	if n%3 == 0 {
		d = append(d, "Pling")
	}

	if n%5 == 0 {
		d = append(d, "Plang")
	}

	if n%7 == 0 {
		d = append(d, "Plong")
	}

	if len(d) == 0 {
		return strconv.Itoa(n)
	}

	return strings.Join(d, "")
}
