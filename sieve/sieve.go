package sieve

// Sieve finds prime numbers up to and including the limit,
// using the sieve of Eristosthenes
func Sieve(limit int) []int {
	primes := make([]int, 0, 12)

	isComposite := make([]bool, limit+1)
	p, done := nextPrime(isComposite, 1)
	for !done {
		primes = append(primes, p)
		for k := p * 2; k <= limit; k += p {
			isComposite[k] = true
		}
		p, done = nextPrime(isComposite, p)
	}

	return primes
}

func nextPrime(composites []bool, after int) (p int, done bool) {
	for p = after + 1; p < len(composites); p++ {
		if !composites[p] {
			return p, false
		}
	}
	return 0, true
}
