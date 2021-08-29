package game

import (
	"testing"
)

func TestNew(t *testing.T) {
	tt := []struct {
		name    string
		in      []string
		want    boardInfo
		wantErr bool
	}{
		{
			name:    "empty input -> error",
			in:      nil,
			wantErr: true,
		},
		{
			name:    "bad input character -> error",
			in:      []string{"q"},
			wantErr: true,
		},
		{
			name:    "non-square input -> error",
			in:      []string{".", "", "."},
			wantErr: true,
		},
		{
			name: "pretty input -> error",
			in: []string{
				". . .  ",
				" . . . ",
				"  . . .",
			},
			wantErr: true,
		},
		{
			name: "size one grid -> ok, the edges have a shape on them",
			in:   []string{"."},
			want: boardInfo{width: 1, height: 1, x: 2, o: 2},
		},
		{
			name: "size 3 board with Xs and Os -> ok",
			in: []string{
				"X.O",
				".XO",
				"OOX",
			},
			want: boardInfo{width: 3, height: 3, x: 5, o: 6},
		},
	}
	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got, err := New(tc.in)
			if err != nil {
				if !tc.wantErr {
					t.Fatalf("got unexpected error: %s", err)
				} else {
					return
				}
			}

			t.Logf("%#v", got.board)

			assertBoardMatches(t, tc.want, got.board)
		})
	}
}

type boardInfo struct {
	width, height int
	x, o          int
}

func assertBoardMatches(t *testing.T, want boardInfo, got board) {
	if got.width != want.width {
		t.Logf("got width %d ; want width %d", got.width, want.width)
		t.Fail()
	}

	if got.height != want.height {
		t.Logf("got height %d ; want height %d", got.height, want.height)
		t.Fail()
	}

	var x, o int
	for _, shape := range got.tiles {
		switch shape {
		case shapeX:
			x++
		case shapeO:
			o++
		default:
			t.Logf("got unexpected shape on grid: %q", shape)
			t.Fail()
		}
	}

	if x != want.x {
		t.Logf("count of 'X' = %d ; want %d", x, want.x)
		t.Fail()
	}

	if o != want.o {
		t.Logf("count of 'O' = %d ; want %d", o, want.o)
		t.Fail()
	}
}
