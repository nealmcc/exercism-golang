package armstrong

import "math"

// IsNumber determines if the given number is an Armstrong Number
// (the sum of its own digits, each raised to the power of the number of digits)
func IsNumber(n int) bool {
	var (
		digits = getDigits(n)
		exp    = float64(len(digits))
		sum    = 0
	)
	for _, d := range digits {
		sum += int(math.Pow(float64(d), exp))
	}
	return n == sum
}

func getDigits(n int) []int {
	digits := make([]int, 0, 8)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	return digits
}
