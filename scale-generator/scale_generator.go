// Package scale solves the Exercism side-exercise 'scale generator'
package scale

import (
	"scale/chromatic"
	"scale/scales"
)

// Scale creates a scale, and returns the names of the notes in it
func Scale(tonic string, interval string) []string {
	rules := chromatic.Rules

	scale, err := scales.NewScale(tonic, interval, rules)
	if err != nil {
		panic(err)
	}

	return scale.Describe()
}
