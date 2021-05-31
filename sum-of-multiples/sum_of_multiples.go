package summultiples

func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	for n := 1; n < limit; n++ {
		if isMultiple(n, divisors) {
			sum += n
		}
	}
	return sum
}

func isMultiple(n int, divisors []int) bool {
	for _, d := range divisors {
		if d == 0 {
			// ignore 0 as a divisor
			continue
		}
		if n%d == 0 {
			return true
		}
	}
	return false
}
