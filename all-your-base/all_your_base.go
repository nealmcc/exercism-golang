package allyourbase

import (
	"errors"
)

var (
	errBaseIn   = errors.New("input base must be >= 2")
	errBaseOut  = errors.New("output base must be >= 2")
	errBadDigit = errors.New("all digits must satisfy 0 <= d < input base")
)

// ConvertToBase converts the given digits from b1 to b2
func ConvertToBase(b1 int, digits []int, b2 int) ([]int, error) {
	if b1 < 2 {
		return nil, errBaseIn
	}
	if b2 < 2 {
		return nil, errBaseOut
	}
	total, err := sum(digits, b1)
	if err != nil {
		return nil, err
	}
	d2 := splitDigits(total, b2)
	return d2, nil
}

func sum(digits []int, base int) (int, error) {
	total, exp := 0, len(digits)
	for _, d := range digits {
		if d < 0 || d >= base {
			return 0, errBadDigit
		}
		total += d * pow(base, exp-1)
		exp--
	}
	return total, nil
}

// pow returns x ^ y
func pow(x, y int) int {
	pow := 1
	for i := 0; i < y; i++ {
		pow *= x
	}
	return pow
}

func splitDigits(total, base int) []int {
	if total == 0 {
		return []int{0}
	}
	digits := make([]int, 0, 3)
	for total > 0 {
		digits = append(digits, total%base)
		total = total / base
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return digits
}
