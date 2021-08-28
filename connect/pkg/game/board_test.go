package game

import (
	"testing"

	hg "connect/pkg/hexgrid"
)

func TestAreAdjacent(t *testing.T) {
	var (
		tiny   board = newBoard(1)
		little board = newBoard(2)
	)

	tt := []struct {
		name string
		b    board
		keys [2]hg.Vkey
		want bool
	}{
		// tiny
		{
			name: "tiny(1): center, top",
			b:    tiny,
			keys: [2]hg.Vkey{{}, tiny.top},
			want: true,
		},
		{
			name: "tiny(1): center, right",
			b:    tiny,
			keys: [2]hg.Vkey{{}, tiny.right},
			want: true,
		},
		{
			name: "tiny(1): center, bottom",
			b:    tiny,
			keys: [2]hg.Vkey{{}, tiny.bottom},
			want: true,
		},
		{
			name: "tiny(1): center, left",
			b:    tiny,
			keys: [2]hg.Vkey{{}, tiny.left},
			want: true,
		},
		{
			name: "tiny(1): left right",
			b:    tiny,
			keys: [2]hg.Vkey{tiny.left, tiny.right},
			want: false,
		},
		{
			name: "tiny(1): top bottom",
			b:    tiny,
			keys: [2]hg.Vkey{tiny.top, tiny.bottom},
			want: false,
		},
		// little
		{
			name: "little(2): top 1",
			b:    little,
			keys: [2]hg.Vkey{{}, little.top},
			want: true,
		},
		{
			name: "little(2): top 2",
			b:    little,
			keys: [2]hg.Vkey{{X: 2, Y: 0}, little.top},
			want: true,
		},
		{
			name: "little(2): right 1",
			b:    little,
			keys: [2]hg.Vkey{{X: 2, Y: 0}, little.right},
			want: true,
		},
		{
			name: "little(2): right 2",
			b:    little,
			keys: [2]hg.Vkey{{X: 3, Y: 1}, little.right},
			want: true,
		},
		{
			name: "little(2): bottom 1",
			b:    little,
			keys: [2]hg.Vkey{{X: 1, Y: 1}, little.bottom},
			want: true,
		},
		{
			name: "little(2): bottom 2",
			b:    little,
			keys: [2]hg.Vkey{{X: 3, Y: 1}, little.bottom},
			want: true,
		},
		{
			name: "little(2): left 1",
			b:    little,
			keys: [2]hg.Vkey{{}, little.left},
			want: true,
		},
		{
			name: "little(2): left 2",
			b:    little,
			keys: [2]hg.Vkey{{X: 1, Y: 1}, little.left},
			want: true,
		},
		{
			name: "little(2): 0,0 2,0",
			b:    little,
			keys: [2]hg.Vkey{{}, {X: 2, Y: 0}},
			want: true,
		},
		{
			name: "little(2): 2,0 3,1",
			b:    little,
			keys: [2]hg.Vkey{{X: 2, Y: 0}, {X: 3, Y: 1}},
			want: true,
		},
		{
			name: "little(2): 3,1 1,1",
			b:    little,
			keys: [2]hg.Vkey{{X: 3, Y: 1}, {X: 1, Y: 1}},
			want: true,
		},
		{
			name: "little(2): 1,1 0,0",
			b:    little,
			keys: [2]hg.Vkey{{X: 1, Y: 1}, {}},
			want: true,
		},
		{
			name: "little(2): 1,1 2,0",
			b:    little,
			keys: [2]hg.Vkey{{X: 1, Y: 1}, {X: 2, Y: 0}},
			want: true,
		},
		{
			name: "little(2): 0,0 3,1",
			b:    little,
			keys: [2]hg.Vkey{{X: 0, Y: 0}, {X: 3, Y: 1}},
			want: false,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			t.Logf("grid: %+v", tc.b)
			assertAdjacent(t, tc.b, tc.keys[0], tc.keys[1], tc.want)
			assertAdjacent(t, tc.b, tc.keys[1], tc.keys[0], tc.want)
		})
	}
}

func assertAdjacent(t *testing.T, b board, k1, k2 hg.Vkey, want bool) {
	t.Helper()
	if got := b.areAdjacent(k1, k2); got != want {
		t.Logf("areAdjacent(%+v, %+v) = %v ; want %v", k1, k2, got, want)
		t.Fail()
	}
}

func TestHasConnection(t *testing.T) {
	/*
		X O O
		 O X O
		  O X X
	*/
	b, err := parseBoard([]string{
		"XOO",
		"OXO",
		"OXX",
	})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	tt := []struct {
		name  string
		shape shape
		start hg.Vkey
		end   hg.Vkey
		want  bool
	}{
		{
			name:  "adjacent, no shape -> false",
			shape: shapeX,
			start: hg.Vkey{X: 1, Y: 1},
			end:   hg.Vkey{X: 2, Y: 0},
			want:  false,
		},
		{
			name:  "disconnected, same shape -> false",
			shape: shapeX,
			start: hg.Vkey{},
			end:   hg.Vkey{X: 3, Y: 1},
			want:  false,
		},
		{
			name:  "same shape, adjacent -> true",
			shape: shapeX,
			start: hg.Vkey{X: 3, Y: 1},
			end:   hg.Vkey{X: 4, Y: 2},
			want:  true,
		},
		{
			name:  "same shape, distance 2 -> true",
			shape: shapeX,
			start: hg.Vkey{X: 3, Y: 1},
			end:   hg.Vkey{X: 6, Y: 2},
			want:  true,
		},
		{
			name:  "same shape, distance 2 -> true",
			shape: shapeX,
			start: hg.Vkey{X: 3, Y: 1},
			end:   hg.Vkey{X: 6, Y: 2},
			want:  true,
		}, {
			name:  "same shape, long path -> true",
			shape: shapeO,
			start: hg.Vkey{X: 2, Y: 2},
			end:   hg.Vkey{X: 5, Y: 1},
			want:  true,
		}, {
			name:  "left to right, multiple paths -> true",
			shape: shapeO,
			start: b.left,
			end:   b.right,
			want:  true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := b.hasConnection(tc.shape, tc.start, tc.end)
			if got != tc.want {
				t.Logf("got = %v ; want %v", got, tc.want)
				t.Fail()
			}
			// test in reverse
			got = b.hasConnection(tc.shape, tc.end, tc.start)
			if got != tc.want {
				t.Logf("got = %v ; want %v", got, tc.want)
				t.Fail()
			}
		})
	}
}
