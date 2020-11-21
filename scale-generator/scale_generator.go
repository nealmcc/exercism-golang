// package scale solves the Exercism side-exercise 'scale generator'
package scale

import "scale/chromatic"

// Scale creates a scale, and returns the names of the notes in it
func Scale(tonic string, interval string) []string {
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}
	scale := chromatic.NewScale(tonic, interval)
	return scale.Describe()
}
