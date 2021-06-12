package sieve

import (
	"log"
	"sort"
)

// Sieve finds prime numbers up to and including the limit by using
// the sieve of Eristosthenes
func Sieve(limit int) []int {
	if limit < 2 {
		return nil
	}

	if limit == 2 {
		return []int{2}
	}

	primes := []int{2, 3}
	for {
		p := primes[len(primes)-1]
		lim := min(p*p, limit)
		primes = sieve(primes, lim)
		if lim == limit {
			return primes
		}
	}
}

var yes = struct{}{}

// sieve performs a bulk iteration of the sieve of Eristosthenes, expanding the
// list of discovered primes up to the given limit.
// Each time this function is called, the given limit should be at most the
// square of the highest prime found so far, so that:
//  a) we reduce the number of candidates in memory at once, and
//  b) we don't need to sort the hashmap keys (to find the next lowest prime)
//     after every individual iteration.
func sieve(primes []int, lim int) []int {
	if len(primes) < 2 {
		primes = []int{2, 3}
	}

	maxp := primes[len(primes)-1]
	if maxp*maxp < lim {
		log.Fatal("each iteration, lim must be <= the highest prime squared")
	}

	// fill the top of the 'sieve'
	newPrimes := make(map[int]struct{}, (lim-maxp)/2)
	for i := maxp + 2; i <= lim; i += 2 {
		newPrimes[i] = yes
	}

	// filter out the composite numbers:
	for i := 1; i < len(primes); i++ {
		p := primes[i]
		for k := p + p; k <= lim; k += p {
			delete(newPrimes, k)
		}
	}

	return append(primes, sortedKeys(newPrimes)...)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func sortedKeys(p map[int]struct{}) []int {
	result := make([]int, len(p))
	i := 0
	for k := range p {
		result[i] = k
		i++
	}
	sort.Ints(result)
	return result
}
