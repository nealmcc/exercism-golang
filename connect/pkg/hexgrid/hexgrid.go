// Package hexgrid implements a hexagonal tile system in two dimensions.
// Tiles can be arranged along three axes:
// E  - W
// NE - SW
// NW - SE
package hexgrid

import "math"

// Vector is a floating-point 2d coordinate.
type Vector struct {
	X, Y float64
}

// Size returns the length of the receiver.
func (a Vector) Size() float64 {
	x2 := math.Pow(a.X, 2)
	y2 := math.Pow(a.Y, 2)
	return math.Sqrt(x2 + y2)
}

// Dist measures the distance between Vectors a and b.
func Dist(a, b Vector) float64 {
	x2 := math.Pow(a.X-b.X, 2)
	y2 := math.Pow(a.Y-b.Y, 2)
	return math.Sqrt(x2 + y2)
}

// Scale multiplies a vector by the given scalar value, returning a new Vector
func Scale(a Vector, sc float64) Vector {
	return Vector{a.X * sc, a.Y * sc}
}

// ToKey rounds the receiver to its nearest discrete key.
func (v Vector) ToKey() Vkey {
	x := int(math.Round(v.X / cos60))
	y := int(math.Round(v.Y / sin60))
	return Vkey{x, y}
}

var (
	sin60 float64 = math.Sqrt(3) / 2
	cos60         = 0.5
)

// Vkey is the logical identifier for a tile on a hexagonal grid where
// the distance between the center points of adjacent tiles is 1 unit.
//
// Vkeys are the preferred way to maintain the state of a grid, because
// they keep full precision through addition and multiplecation, and are
// perfectly comparable.
//
// Note that it is not valid to have x even and y odd or vice versa.
type Vkey struct {
	X int // X multiplied by cos60 gives the geometric x coordinate.
	Y int // Y multiplied by sin60 gives the geometric y coordinate.
}

// East, West, NE, NW, SE, and SW are the directions from a tile to its neighbours.
var (
	East Vkey = Vkey{2, 0}
	West Vkey = Vkey{-2, 0}
	NE   Vkey = Vkey{1, 1}
	NW   Vkey = Vkey{-1, 1}
	SE   Vkey = Vkey{1, -1}
	SW   Vkey = Vkey{-1, -1}
)

// Sum returns the sum of all the given keys.
func Sum(keys ...Vkey) Vkey {
	sum := Vkey{}
	for _, v := range keys {
		sum.X += v.X
		sum.Y += v.Y
	}
	return sum
}

// Times returns k * n.
func (k Vkey) Times(n int) Vkey {
	return Vkey{
		X: k.X * n,
		Y: k.Y * n,
	}
}

// Neighbours returns a slice of tiles adjacent to the receiver.
func (k Vkey) Neighbours() []Vkey {
	adj := make([]Vkey, 6)
	adj[0] = Sum(k, East)
	adj[1] = Sum(k, West)
	adj[2] = Sum(k, NE)
	adj[3] = Sum(k, NW)
	adj[4] = Sum(k, SE)
	adj[5] = Sum(k, SW)
	return adj
}

// ToVector converts the receiver to its floating point representation.
func (k Vkey) ToVector() Vector {
	x := float64(k.X) * cos60
	y := float64(k.Y) * sin60
	return Vector{x, y}
}
