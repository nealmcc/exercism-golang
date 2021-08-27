package hexgrid

import (
	"math"
	"testing"
)

func TestSize(t *testing.T) {
	tt := []struct {
		v    Vector
		want float64
	}{
		{E.ToVector(), 1},
		{W.ToVector(), 1},
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
		{Zero.ToVector(), E.ToVector(), 1},
		{Zero.ToVector(), W.ToVector(), 1},
		{Zero.ToVector(), NE.ToVector(), 1},
		{Zero.ToVector(), SE.ToVector(), 1},
		{Zero.ToVector(), NW.ToVector(), 1},
		{Zero.ToVector(), SW.ToVector(), 1},
		{W.ToVector(), E.ToVector(), 2},
		{NE.ToVector(), SW.ToVector(), 2},
		{NW.ToVector(), SE.ToVector(), 2},
	}
	for _, tc := range tt {
		got := tc.a.Dist(tc.b)
		if math.Abs(tc.want-got) > 1e-15 {
			t.Fatalf("Dist(%v, %v) = %.1f ; want %.1f",
				tc.a, tc.b, got, tc.want)
		}
	}
}

func TestPath(t *testing.T) {
	tt := []struct {
		desc string
		path []Vkey
		want Vector
	}{
		{
			"e w => origin",
			[]Vkey{E, W},
			Zero.ToVector(),
		},
		{
			"ne sw => origin",
			[]Vkey{NE, SW},
			Zero.ToVector(),
		},
		{
			"nw se => origin",
			[]Vkey{NW, SE},
			Zero.ToVector(),
		},
		{
			"e se w => se",
			[]Vkey{E, SE, W},
			SE.ToVector(),
		},
		{
			"nw sw => west",
			[]Vkey{NW, SW},
			W.ToVector(),
		},
	}

	for _, tc := range tt {
		got := Zero.Sum(tc.path...).ToVector()
		if !got.IsClose(tc.want) {
			t.Fatalf("Sum(%s) = %v ; want %v", tc.desc, got, tc.want)
		}
	}
}
