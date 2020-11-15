// Package triangle solves the Exercism 'Triangle' problem.
package triangle

import "math"

// Kind refers to a type of triangle (see Equ, Iso, Sca, NaT).
type Kind int

const (
	// Equ is an equilateral triangle.
	Equ Kind = 3
	// Iso is an isosceles triangle.
	Iso Kind = 2
	// Sca is a scalene triangle.
	Sca Kind = 1
	// NaT is not a triangle.
	NaT Kind = -1
)

// KindFromSides finds the Kind of a potential triangle with sides of length a,b,c
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
	inf := math.Inf(1)
	return a > 0 && b > 0 && c > 0 &&
		a < inf && b < inf && c < inf &&
		(b+c >= a) && (a+c >= b) && (a+b >= c)
}

func isEquilateral(a, b, c float64) bool {
	return a == b && b == c
}

func isScalene(a, b, c float64) bool {
	return (a != b) && (a != c) && (b != c)
}
