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

	min, max := fmax*fmax+1, -1
	for a := fmin; a <= fmax; a++ {
		for b := a; b <= fmax; b++ {
			p := a * b
			if min < p && p < max || !isPalindrome(p) {
				continue
			}
			if p < min {
				min = p
				pmin = Product{min, nil}
			}
			if p == min {
				pmin.Factorizations = append(pmin.Factorizations, [2]int{a, b})
			}
			if p > max {
				max = p
				pmax = Product{max, nil}
			}
			if p == max {
				pmax.Factorizations = append(pmax.Factorizations, [2]int{a, b})
			}
		}
	}

	if min > max {
		err = errors.New("no palindromes")
	}

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
