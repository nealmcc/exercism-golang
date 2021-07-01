package palindrome

import (
	"errors"
	"strconv"
)

// Product is a palindrome number and its factor pairs
type Product struct {
	Product        int
	Factorizations [][2]int
}

// Products is used to find the smallest and largest palindrome number
// with factor pairs in the range [fmin, fmax]
// returns an error if fmin > fmax, or there are no palindromes
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = errors.New("fmin > fmax")
		return
	}

	min, max := fmax*fmax+1, 0
	factors := make(map[int][][2]int, 8)
	for a := fmin; a <= fmax; a++ {
		for b := a; b <= fmax; b++ {
			p := a * b
			if min < p && p < max {
				continue
			}
			if !isPalindrome(p) {
				continue
			}
			factors[p] = append(factors[p], [2]int{a, b})
			if p < min {
				min = p
			}
			if p > max {
				max = p
			}
		}
	}

	if len(factors) == 0 {
		err = errors.New("no palindromes")
		return
	}

	pmin = Product{min, factors[min]}
	pmax = Product{max, factors[max]}
	return
}

// isPalindrome returns true if the base-10 representation of n is a palindrome
func isPalindrome(n int) bool {
	digits := strconv.Itoa(n)
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}
