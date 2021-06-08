package perfect

import (
	"errors"
)

type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("only positive numbers can be classified")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}

	done := make(chan struct{})
	defer close(done)

	factors := genFactors(n, done)
	limit := 2 * n
	sum := <-sumToLimit(factors, limit)

	switch {
	case sum < limit:
		return ClassificationDeficient, nil
	case sum > limit:
		return ClassificationAbundant, nil
	default:
		return ClassificationPerfect, nil
	}
}

// sumToLimit produces the sum of the values it receives on the input.
// It will stop early if the given limit is exceeded.
func sumToLimit(in <-chan int64, limit int64) <-chan int64 {
	out := make(chan int64)
	go func() {
		var sum int64
		for n := range in {
			sum += n
			if sum > limit {
				out <- sum
			}
		}
		out <- sum
	}()
	return out
}

// genFactors produces all the factors of n.
// It may be cancelled early by closing the quit channel.
func genFactors(n int64, quit <-chan struct{}) <-chan int64 {
	factors := make(chan int64, 4)
	go func() {
		defer close(factors)
		a := int64(1)
		for a*a < n {
			b, rem := n/a, n%a
			if rem == 0 {
				select {
				case factors <- a:
				case <-quit:
				}
				if a != b {
					select {
					case factors <- b:
					case <-quit:
					}
				}
			}
			a++
		}
	}()
	return factors
}
