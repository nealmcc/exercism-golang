package game

import (
	"testing"

	"connect/pkg/hex"
)

func TestPush(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name  string
		start stack
		add   hex.Vkey
		want  int
	}{
		{
			name:  "empty stack, push once",
			start: stack{},
			add:   hex.NE,
			want:  1,
		},
		{
			name:  "non-empty stack, push once",
			start: stack{stackitem{tile: hex.East}},
			add:   hex.West,
			want:  2,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			s := tc.start
			s.push(tc.add, nil)

			if len(s) != tc.want {
				t.Logf("got length %d ; want %d", len(s), tc.want)
				t.Fail()
			}
		})
	}
}

func TestPop(t *testing.T) {
	tt := []struct {
		name    string
		start   stack
		count   int
		want    hex.Vkey
		wantOK  bool
		wantLen int
	}{
		{
			name:   "empty stack, pop once => not ok",
			count:  1,
			wantOK: false,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			s := tc.start
			var (
				got hex.Vkey
				ok  bool
			)

			for n := 0; n < tc.count; n++ {
				got, _, ok = s.pop()
			}

			if ok != tc.wantOK {
				t.Fatalf("s.pop() x%d -> ok = %v ; want %v", tc.count, ok, tc.wantOK)
			}

			if got != tc.want {
				t.Fatalf("s.pop() x%d = %v ; want %v", tc.count, got, tc.want)
			}

			if len(s) != tc.wantLen {
				t.Fatalf("pop() got len %d ; want %d", len(s), tc.wantLen)
			}
		})
	}
}
