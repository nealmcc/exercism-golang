// Package triangle solves the Exercism 'Triangle' problem.
package triangle

import "math"

// Kind refers to a type of triangle (see the constants Equ, Iso, Sca, NaT).
type Kind int

const (
	// Equ is a triangle with all three sides having the same length.
	Equ Kind = 3

	// Iso is a triagle with exactly two sides having the same length.
	Iso Kind = 2

	// Sca is a triangle with every side having different lengths.
	Sca Kind = 1

	// NaT is not a triangle.
	NaT Kind = -1
)

// KindFromSides takes the lengths of the sides of a triangle and
// determines what Kind it is.
func KindFromSides(a, b, c float64) Kind {
	if !isTriangle(a, b, c) {
		return NaT
	}

	if isEquilateral(a, b, c) {
		return Equ
	}

	if isScalene(a, b, c) {
		return Sca
	}

	return Iso
}

func isTriangle(a, b, c float64) bool {
	sidesPositive := a > 0 && b > 0 && c > 0
	inf := math.Inf(1)
	sidesFinite := a < inf && b < inf && c < inf
	passesInequality := (b+c >= a) && (a+c >= b) && (a+b >= c)
	return sidesPositive && sidesFinite && passesInequality
}

func isEquilateral(a, b, c float64) bool {
	return a == b && a == c && b == c
}

func isScalene(a, b, c float64) bool {
	return (a != b) && (a != c) && (b != c)
}
