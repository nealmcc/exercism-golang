package rectangles

import (
	"testing"
)

func TestRectangles(t *testing.T) {
	for _, tc := range testCases {
		if actual := Count(tc.input); actual != tc.expected {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkRectangles(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Count(tc.input)
		}
	}
}

func TestIsHorizontalEdge(t *testing.T) {
	tt := []struct {
		name string
		in   string
		want bool
	}{
		{
			name: "empty string",
			in:   "",
			want: false,
		},
		{
			name: "one corner",
			in:   "+",
			want: false,
		},
		{
			name: "missing left",
			in:   "-+",
			want: false,
		},
		{
			name: "missing right",
			in:   "+-",
			want: false,
		},
		{
			name: "adjacent corners",
			in:   "++",
			want: true,
		},
		{
			name: "one edge node",
			in:   "+-+",
			want: true,
		},
		{
			name: "many edge nodes",
			in:   "+-----+",
			want: true,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := isEdge([]byte(tc.in), '-'); got != tc.want {
				t.Fatalf("isEdge(%s, '-') = %v ; want %v",
					tc.in, got, tc.want)
			}
		})
	}
}

func TestIsVerticalEdge(t *testing.T) {
	tt := []struct {
		name   string
		in     string
		c1, c2 int
		want   bool
	}{
		{
			name: "empty string",
			in:   "",
			want: false,
		},
		{
			name: "one corner",
			in:   "+",
			want: false,
		},
		{
			name: "missing top",
			in:   "|+",
			want: false,
		},
		{
			name: "missing bottom",
			in:   "+|",
			want: false,
		},
		{
			name: "adjacent corners",
			in:   "++",
			want: true,
		},
		{
			name: "one edge node",
			in:   "+|+",
			want: true,
		},
		{
			name: "many edge nodes",
			in:   "+||||||+",
			want: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if got := isEdge([]byte(tc.in), '|'); got != tc.want {
				t.Fatalf("isEdge(%s, '|') = %v ; want %v",
					tc.in, got, tc.want)
			}
		})
	}
}
