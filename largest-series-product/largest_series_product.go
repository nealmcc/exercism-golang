package lsproduct

import (
	"errors"
	"fmt"
)

var (
	ErrSpanNegative  = errors.New("span must be greater than zero")
	ErrSpanTooLong   = errors.New("span must be smaller than string length")
	ErrDigitsInvalid = errors.New("digits input must only contain digits")
)

// LargestSeriesProduct calculates the largest product for
// a contiguous substring of digits from the given string
func LargestSeriesProduct(s string, span int) (int64, error) {
	if span < 0 {
		return 0, fmt.Errorf("%w", ErrSpanNegative)
	}
	if span > len(s) {
		return 0, fmt.Errorf("%w", ErrSpanTooLong)
	}
	digits, err := tryParse(s)
	if err != nil {
		return 0, err
	}
	var max int64
	for i := 0; i <= len(digits)-span; i++ {
		prod := multiply(digits[i : i+span]...)
		if prod > max {
			max = prod
		}
	}
	return max, nil
}

func tryParse(s string) ([]byte, error) {
	digits := make([]byte, 0, 50)
	for i := 0; i < len(s); i++ {
		d := s[i] - '0'
		if d > 9 {
			return nil, fmt.Errorf("%w", ErrDigitsInvalid)
		}
		digits = append(digits, d)
	}
	return digits, nil
}

func multiply(bytes ...byte) int64 {
	prod := int64(1)
	for _, b := range bytes {
		prod *= int64(b)
	}
	return prod
}
