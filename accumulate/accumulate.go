// Package accumulate solves the Exercism Accumulate side exercise
package accumulate

type operation func(string) string

// Accumulate transforms each element of the input using the given operation
func Accumulate(in []string, op operation) []string {
	out := make([]string, len(in))
	for i, s := range in {
		out[i] = op(s)
	}
	return out
}
