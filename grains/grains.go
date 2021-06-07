package grains

import "errors"

// Square calculates how many grains of square there are on the nth square
// of a chessboard, where the number doubles each time.
// effectively, we are calculating 2^(n-1)
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("n must be in the range 1 <= n <= 64")
	}
	return 1 << (n - 1), nil
}

// Total returns the sum of the grains on all of the squares.
func Total() uint64 {
	return 1<<64 - 1
}
