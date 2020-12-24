// Package raindrops solves the Exercism exercise
package raindrops

import (
	"strconv"
)

// Convert an integer n to raindrop sounds
func Convert(n int) (drops string) {
	if n%3 == 0 {
		drops += "Pling"
	}

	if n%5 == 0 {
		drops += "Plang"
	}

	if n%7 == 0 {
		drops += "Plong"
	}

	if len(drops) == 0 {
		return strconv.Itoa(n)
	}

	return drops
}
