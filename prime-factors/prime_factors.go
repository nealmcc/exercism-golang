package prime

import (
	"sort"
)

// cache the prime numbers as we find them:
var primeCache = []int64{2, 3, 5, 7, 11, 13, 17, 19}

// Factors determines the prime factors of the given natural number
func Factors(n int64) []int64 {
	if n < 2 {
		return []int64{}
	}

	quo, factors := findFactors(n, primeCache)
	maxp := primeCache[len(primeCache)-1]

	for quo != 1 && quo/maxp > maxp {
		morePrimes := nextPrimes(primeCache, maxp*maxp)
		q2, moreFactors := findFactors(quo, morePrimes)
		quo = q2
		primeCache = append(primeCache, morePrimes...)
		factors = append(factors, moreFactors...)
		maxp = primeCache[len(primeCache)-1]
	}

	if quo != 1 {
		factors = append(factors, quo)
	}

	return factors
}

// findFactors finds a subset of the prime factors of n,
// using trial division with the given set of prime numbers.
// returns all of the factors found using the given set of primes, and
// the quotient that is left, such that quo * product(factors) == n
func findFactors(n int64, primes []int64) (quo int64, factors []int64) {
	for _, p := range primes {
		if n < p {
			break
		}
		quo, rem := divmod(n, p)
		for rem == 0 {
			n = quo
			factors = append(factors, p)
			quo, rem = divmod(n, p)
		}
	}

	return n, factors
}

// divmod finds the quotient and remainder of n/d
func divmod(n, d int64) (quo, rem int64) {
	quo = n / d
	rem = n % d
	return
}

// nextPrimes accepts a list of primes that have been discovered so far
// and an upper limit to search for more. It returns a sorted slice of
// primes, not including the ones it was given.
func nextPrimes(primes []int64, limit int64) []int64 {
	maxp := primes[len(primes)-1]

	// We don't care about the value, just the key.
	// an empty struct{} takes 0 space to store.
	newPrimes := make(map[int64]struct{}, (limit-maxp)/2)
	for i := maxp + 2; i <= limit; i += 2 {
		newPrimes[i] = struct{}{}
	}

	// remove all multiples of the primes we already know about.
	// skip the first prime (2)
	for i := 1; i < len(primes); i++ {
		p := primes[i]
		for k := p * p; k <= limit; k += p {
			delete(newPrimes, k)
		}
	}

	// all the remaining elements in the set are prime
	result := make([]int64, len(newPrimes))
	i := 0
	for k := range newPrimes {
		result[i] = k
		i++
	}

	// we have to sort the results, because ranging over the keys in a map
	// doesn't necessarily happen in any particular order.
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	return result
}
