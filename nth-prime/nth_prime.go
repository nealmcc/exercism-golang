package prime

import (
	"errors"
	"log"
	"sort"
)

// Nth returns the nth prime, assuming n is > 1.
// if n is <1, returns 0 and false
func Nth(n int) (int, bool) {
	if n < 1 {
		return 0, false
	}

	var primes = []int{2, 3, 5}

	for len(primes) < n {
		maxP := primes[len(primes)-1]
		limit := maxP * maxP
		primes = Sieve(primes, limit)
	}

	return primes[n-1], true
}

// Sieve accepts a list of primes that have been discovered so far, and an
// upper limit to search for more. It returns a newly allocated slice of
// primes. This function is only thread safe if the underlying contents of
// primes is not altered.
func Sieve(primes []int, limit int) []int {
	if len(primes) < 1 {
		log.Fatal(ErrPrimesEmpty)
	}

	maxPrime := primes[len(primes)-1]
	if maxPrime*maxPrime < limit {
		log.Fatal(ErrOutOfBounds)
	}

	// We don't care about the value, just the key.
	// an empty struct{} takes 0 space to store.
	newPrimes := make(map[int]struct{}, (limit-maxPrime)/2)
	for i := maxPrime + 2; i <= limit; i += 2 {
		newPrimes[i] = struct{}{}
	}

	// remove all multiples of the primes we already know about.
	// skip the first prime (2)
	for i := 1; i < len(primes); i++ {
		p := primes[i]
		for k := p; k <= limit; k += p {
			delete(newPrimes, k)
		}
	}

	// all the remaining elements in the set are prime
	result := make([]int, len(primes)+len(newPrimes))
	copy(result, primes)
	i := len(primes)
	for k := range newPrimes {
		result[i] = k
		i++
	}

	// we have to sort the results, because ranging over the keys in a map
	// doesn't necessarily happen in any particular order.
	sort.Ints(result)
	return result
}

var (
	ErrPrimesEmpty = errors.New("the slice of found primes must not be empty")
	ErrOutOfBounds = errors.New("n must be at most the largest prime squared")
)
