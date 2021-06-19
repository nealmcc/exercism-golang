package primefactors

import "math"

// Source: exercism/problem-specifications
// Commit: d928874 prime-factors: apply "input" policy
// Problem Specifications Version: 1.1.0

var tests = []struct {
	description string
	input       int64
	expected    []int64
}{

	{
		"no factors",
		1,
		[]int64{},
	},
	{
		"prime number",
		2,
		[]int64{2},
	},
	{
		"square of a prime",
		9,
		[]int64{3, 3},
	},
	{
		"cube of a prime",
		8,
		[]int64{2, 2, 2},
	},
	{
		"product of primes and non-primes",
		12,
		[]int64{2, 2, 3},
	},
	{
		"product of primes",
		901255,
		[]int64{5, 17, 23, 461},
	},
	{
		"factors include a large prime",
		93819012551,
		[]int64{11, 9539, 894119},
	},
	{
		"largest int64 - checking for memory / overflow issues",
		math.MaxInt64,
		[]int64{7, 7, 73, 127, 337, 92737, 649657},
	},
}
