package hex

import (
	"math"
	"testing"
)

func TestPath(t *testing.T) {
	tt := []struct {
		desc string
		path []Vkey
		want Vector
	}{
		{
			"e w => origin",
			[]Vkey{East, West},
			Vector{},
		},
		{
			"ne sw => origin",
			[]Vkey{NE, SW},
			Vector{},
		},
		{
			"nw se => origin",
			[]Vkey{NW, SE},
			Vector{},
		},
		{
			"e se w => se",
			[]Vkey{East, SE, West},
			SE.ToVector(),
		},
		{
			"nw sw => west",
			[]Vkey{NW, SW},
			West.ToVector(),
		},
	}

	for _, tc := range tt {
		got := Sum(tc.path...).ToVector()
		if !nearlyEq(t, tc.want, got) {
			t.Fatalf("Sum(%s) = %v ; want %v", tc.desc, got, tc.want)
		}
	}
}

func nearlyEq(t *testing.T, a, b Vector) bool {
	t.Helper()
	delta := 1e-15
	return math.Abs(a.X-b.X) < delta && math.Abs(a.Y-b.Y) < delta
}

func TestSize(t *testing.T) {
	tt := []struct {
		v    Vector
		want float64
	}{
		{East.ToVector(), 1},
		{West.ToVector(), 1},
		{NE.ToVector(), 1},
		{SE.ToVector(), 1},
		{NW.ToVector(), 1},
		{SW.ToVector(), 1},
	}
	for _, tc := range tt {
		got := tc.v.Size()
		if math.Abs(tc.want-got) > 1e-15 {
			t.Fatalf("Len(%v) = %.1f ; want %.1f",
				tc.v, got, tc.want)
		}
	}
}

func TestDist(t *testing.T) {
	tt := []struct {
		a, b Vector
		want float64
	}{
		{Vkey{}.ToVector(), East.ToVector(), 1},
		{Vkey{}.ToVector(), West.ToVector(), 1},
		{Vkey{}.ToVector(), NE.ToVector(), 1},
		{Vkey{}.ToVector(), SE.ToVector(), 1},
		{Vkey{}.ToVector(), NW.ToVector(), 1},
		{Vkey{}.ToVector(), SW.ToVector(), 1},
		{West.ToVector(), East.ToVector(), 2},
		{NE.ToVector(), SW.ToVector(), 2},
		{NW.ToVector(), SE.ToVector(), 2},
	}
	for _, tc := range tt {
		got := Dist(tc.a, tc.b)
		if math.Abs(tc.want-got) > 1e-15 {
			t.Fatalf("Dist(%v, %v) = %.1f ; want %.1f",
				tc.a, tc.b, got, tc.want)
		}
	}
}
