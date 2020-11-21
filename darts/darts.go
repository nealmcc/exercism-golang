// packate darts solves the Exercism side-exercise
package darts

import "math"

// radii of each zone on the dart board
const (
	inner  = 1
	middle = 5
	outer  = 10
)

// Score assigns points for a thrown dart
func Score(x, y float64) int {
	switch dist := math.Sqrt(x*x + y*y); {
	case dist <= inner:
		return 10
	case dist <= middle:
		return 5
	case dist <= outer:
		return 1
	default:
		return 0
	}
}
