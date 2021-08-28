package game

import (
	"connect/pkg/hexgrid"
	"testing"
)

func TestPush(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name  string
		start stack
		add   hexgrid.Vkey
		want  stack
	}{
		{
			name:  "empty stack, push once",
			start: stack{},
			add:   hexgrid.NE,
			want:  stack{hexgrid.NE},
		},
		{
			name:  "non-empty stack, push once",
			start: stack{hexgrid.E},
			add:   hexgrid.W,
			want:  stack{hexgrid.E, hexgrid.W},
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			s := tc.start
			s.push(tc.add)

			if len(s) != len(tc.want) {
				t.Fatalf("push() got len %d ; want %d", len(s), len(tc.want))
			}

			for i, val := range s {
				if tc.want[i] != val {
					t.Logf("got stack[%d] = %v ; want %v", i, val, tc.want[i])
					t.Fail()
				}
			}
		})
	}
}

func TestPop(t *testing.T) {
	tt := []struct {
		name   string
		start  stack
		n      int
		want   hexgrid.Vkey
		wantOK bool
		after  stack
	}{
		{
			name:   "empty stack, pop => not ok",
			n:      1,
			wantOK: false,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			s := tc.start
			var (
				got hexgrid.Vkey
				ok  bool
			)
			for n := 0; n < tc.n; n++ {
				got, ok = s.pop()
			}
			if ok != tc.wantOK {
				t.Fatalf("s.pop() x%d -> ok = %v ; want %v", tc.n, ok, tc.wantOK)
			}
			if got != tc.want {
				t.Fatalf("s.pop() x%d = %v ; want %v", tc.n, got, tc.want)
			}

			if len(s) != len(tc.after) {
				t.Fatalf("pop() got len %d ; want %d", len(s), len(tc.after))
			}

			for i, val := range s {
				if tc.after[i] != val {
					t.Logf("got stack[%d] = %v ; want %v", i, val, tc.after[i])
					t.Fail()
				}
			}
		})
	}
}
