// Package collatzconjecture solves the side-exercise from Exercism
package collatzconjecture

import "errors"

// CollatzConjecture counts the number of steps to solve the '3x+1 problem'
func CollatzConjecture(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("input must be positive")
	}
	i := 0
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		i++
	}
	return i, nil
}
