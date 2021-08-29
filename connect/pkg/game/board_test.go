package game

import (
	"testing"

	"connect/pkg/hex"
)

func TestAreAdjacent(t *testing.T) {
	var (
		tiny   board = newBoard(1)
		little board = newBoard(2)
	)

	tt := []struct {
		name string
		b    board
		keys [2]hex.Vkey
		want bool
	}{
		// tiny
		{
			name: "tiny(1): center, top",
			b:    tiny,
			keys: [2]hex.Vkey{{}, tiny.top},
			want: true,
		},
		{
			name: "tiny(1): center, right",
			b:    tiny,
			keys: [2]hex.Vkey{{}, tiny.right},
			want: true,
		},
		{
			name: "tiny(1): center, bottom",
			b:    tiny,
			keys: [2]hex.Vkey{{}, tiny.bottom},
			want: true,
		},
		{
			name: "tiny(1): center, left",
			b:    tiny,
			keys: [2]hex.Vkey{{}, tiny.left},
			want: true,
		},
		{
			name: "tiny(1): left right",
			b:    tiny,
			keys: [2]hex.Vkey{tiny.left, tiny.right},
			want: false,
		},
		{
			name: "tiny(1): top bottom",
			b:    tiny,
			keys: [2]hex.Vkey{tiny.top, tiny.bottom},
			want: false,
		},
		// little
		{
			name: "little(2): top 1",
			b:    little,
			keys: [2]hex.Vkey{{}, little.top},
			want: true,
		},
		{
			name: "little(2): top 2",
			b:    little,
			keys: [2]hex.Vkey{{X: 2, Y: 0}, little.top},
			want: true,
		},
		{
			name: "little(2): right 1",
			b:    little,
			keys: [2]hex.Vkey{{X: 2, Y: 0}, little.right},
			want: true,
		},
		{
			name: "little(2): right 2",
			b:    little,
			keys: [2]hex.Vkey{{X: 3, Y: 1}, little.right},
			want: true,
		},
		{
			name: "little(2): bottom 1",
			b:    little,
			keys: [2]hex.Vkey{{X: 1, Y: 1}, little.bottom},
			want: true,
		},
		{
			name: "little(2): bottom 2",
			b:    little,
			keys: [2]hex.Vkey{{X: 3, Y: 1}, little.bottom},
			want: true,
		},
		{
			name: "little(2): left 1",
			b:    little,
			keys: [2]hex.Vkey{{}, little.left},
			want: true,
		},
		{
			name: "little(2): left 2",
			b:    little,
			keys: [2]hex.Vkey{{X: 1, Y: 1}, little.left},
			want: true,
		},
		{
			name: "little(2): 0,0 2,0",
			b:    little,
			keys: [2]hex.Vkey{{}, {X: 2, Y: 0}},
			want: true,
		},
		{
			name: "little(2): 2,0 3,1",
			b:    little,
			keys: [2]hex.Vkey{{X: 2, Y: 0}, {X: 3, Y: 1}},
			want: true,
		},
		{
			name: "little(2): 3,1 1,1",
			b:    little,
			keys: [2]hex.Vkey{{X: 3, Y: 1}, {X: 1, Y: 1}},
			want: true,
		},
		{
			name: "little(2): 1,1 0,0",
			b:    little,
			keys: [2]hex.Vkey{{X: 1, Y: 1}, {}},
			want: true,
		},
		{
			name: "little(2): 1,1 2,0",
			b:    little,
			keys: [2]hex.Vkey{{X: 1, Y: 1}, {X: 2, Y: 0}},
			want: true,
		},
		{
			name: "little(2): 0,0 3,1",
			b:    little,
			keys: [2]hex.Vkey{{X: 0, Y: 0}, {X: 3, Y: 1}},
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

func assertAdjacent(t *testing.T, b board, k1, k2 hex.Vkey, want bool) {
	t.Helper()
	if got := b.areAdjacent(k1, k2); got != want {
		t.Logf("areAdjacent(%+v, %+v) = %v ; want %v", k1, k2, got, want)
		t.Fail()
	}
}

func TestCanConnect(t *testing.T) {
	b, err := parseBoard([]string{
		"OXX.", // O X X .
		"OOXX", //  O O X X
		".XO.", //   . X O .
		"XO..", //    X O . .
	})
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	b.tiles[b.left] = shapeX
	b.tiles[b.right] = shapeX
	b.tiles[b.top] = shapeO
	b.tiles[b.top] = shapeO

	tt := []struct {
		name  string
		shape shape
		start hex.Vkey
		end   hex.Vkey
		want  bool
	}{
		{
			name:  "no shape, adjacent -> false",
			shape: none,
			start: hex.Vkey{X: 7, Y: 3},
			end:   hex.Vkey{X: 9, Y: 3},
			want:  false,
		},
		{
			name:  "same shape, disconnected -> false",
			shape: shapeO,
			start: hex.Vkey{},
			end:   hex.Vkey{X: 6, Y: 2},
			want:  false,
		},
		{
			name:  "same shape, adjacent -> true",
			shape: shapeO,
			start: hex.Vkey{X: 1, Y: 1},
			end:   hex.Vkey{X: 3, Y: 1},
			want:  true,
		},
		{
			name:  "same shape, distance 2 -> true",
			shape: shapeO,
			start: hex.Vkey{},
			end:   hex.Vkey{X: 3, Y: 1},
			want:  true,
		},
		{
			name:  "same shape, long path -> true",
			shape: shapeX,
			start: hex.Vkey{X: 2, Y: 0},
			end:   hex.Vkey{X: 3, Y: 3},
			want:  true,
		},
		{
			name:  "left edge onto 'distant' tile -> true",
			shape: shapeX,
			start: b.left,
			end:   hex.Vkey{X: 3, Y: 3},
			want:  true,
		},
		{
			name:  "left to right edge -> true",
			shape: shapeX,
			start: b.left,
			end:   b.right,
			want:  true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := b.canConnect(tc.shape, tc.start, tc.end)
			if got != tc.want {
				t.Logf("got = %v ; want %v", got, tc.want)
				t.Fail()
			}
			// test in reverse
			got = b.canConnect(tc.shape, tc.end, tc.start)
			if got != tc.want {
				t.Logf("got = %v ; want %v", got, tc.want)
				t.Fail()
			}
		})
	}
}
