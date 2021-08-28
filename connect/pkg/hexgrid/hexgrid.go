// Package hexgrid defines a geometric system based on three axis in
// two dimensions.  The three axis are:
// E  - W
// NE - SW
// NW - SE
// The grid can be either discrete with integer coordinates, or continuous
// using floating point coordinates.
package hexgrid

import "math"

// Vector is the 2d coordinate of the center of a hexagonal tile on a grid.
// The distance between the center points of tiles is defined to be 1 unit.
type Vector struct {
	X, Y float64
}

// Vkey is a vector expressed in a form suitable for the key in a hashtable.
// x is a multiple of cos60
// y is a multiple of sin60
// Note that it is not valid to have x even and y odd or vice versa.
type Vkey struct {
	X, Y int
}

// E, W, NE, NW, SE, and SW are the defined keys to navigate
// from a given tile to its neighbours.
var (
	E    Vkey = Vkey{2, 0}
	W    Vkey = Vkey{-2, 0}
	NE   Vkey = Vkey{1, 1}
	NW   Vkey = Vkey{-1, 1}
	SE   Vkey = Vkey{1, -1}
	SW   Vkey = Vkey{-1, -1}
	Zero Vkey = Vkey{0, 0}
)

// EGeom, WGeom, NEGeom, NWGeom, SEGeom, and SWGeom are the geometric Vectors
// from a given tile to its neighbours.
var (
	EGeom    Vector = E.ToVector()
	WGeom    Vector = W.ToVector()
	NEGeom   Vector = NE.ToVector()
	NWGeom   Vector = NW.ToVector()
	SEGeom   Vector = SE.ToVector()
	SWGeom   Vector = SW.ToVector()
	ZeroGeom Vector = Zero.ToVector()
)

var (
	sin60 float64 = math.Sqrt(3) / 2
	cos60         = 0.5
)

// ToKey rounds this vector to its nearest discrete (Vkey) representation.
func (v Vector) ToKey(dia float64) Vkey {
	x := int(math.Round(v.X / cos60))
	y := int(math.Round(v.Y / sin60))
	return Vkey{x, y}
}

// ToVector converts this Vkey to its floating point (Vector) representation.
func (k Vkey) ToVector() Vector {
	x := float64(k.X) * cos60
	y := float64(k.Y) * sin60
	return Vector{x, y}
}

// Adjacent returns a list of keys for tiles adjacent to this one.
func (k Vkey) Adjacent() []Vkey {
	adj := make([]Vkey, 6)
	adj[0] = k.Sum(E)
	adj[1] = k.Sum(W)
	adj[2] = k.Sum(NE)
	adj[3] = k.Sum(NW)
	adj[4] = k.Sum(SE)
	adj[5] = k.Sum(SW)
	return adj
}

// Sum returns the sum this Vkey and all the given keys.
func (k Vkey) Sum(keys ...Vkey) Vkey {
	sum := k
	for _, v := range keys {
		sum.X += v.X
		sum.Y += v.Y
	}
	return sum
}

// Size measures the length of a Vector.
func (a Vector) Size() float64 {
	x2 := math.Pow(a.X, 2)
	y2 := math.Pow(a.Y, 2)
	return math.Sqrt(x2 + y2)
}

// Dist measures the distance between this Vector and the other one.
func (a Vector) Dist(b Vector) float64 {
	x2 := math.Pow(a.X-b.X, 2)
	y2 := math.Pow(a.Y-b.Y, 2)
	return math.Sqrt(x2 + y2)
}

// Scale multiplies a vector by the given scalar value, returning a new Vector
func (a Vector) Scale(sc float64) Vector {
	return Vector{a.X * sc, a.Y * sc}
}

// IsClose will return true iff both the x and y values are within 1e-15.
func (a Vector) IsClose(b Vector) bool {
	delta := 1e-15
	return math.Abs(a.X-b.X) < delta && math.Abs(a.Y-b.Y) < delta
}
