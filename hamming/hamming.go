// Package hamming holds exercise 2 from Exercism
package hamming

import "errors"

// Distance calculates the Hamming Distance between two strings.
func Distance(a, b string) (int, error) {
	length := len(a)
	if length != len(b) {
		return 0, errors.New("Hamming distance is only defined for strings of equal length.")
	}
	dist := 0
	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist, nil
}
