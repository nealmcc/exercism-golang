package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode returns a cipher text for the given UTF-8 plain text.
func Encode(plain string) string {
	sq := chop(normalize(plain))
	return sq.Cipher()
}

// normalize returns a new string, all lowercase string,
// with non-letter and non-digit runes removed.
func normalize(plain string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, plain)
}

type square struct {
	rows, cols int
	text       [][]rune
}

// chop splits the incoming string into a rectangle of
// r slices of c runes, where c >= r and c-r <= 1
func chop(s string) *square {
	runes := []rune(s)
	length := len(runes)
	rows := int(math.Sqrt(float64(length)))
	cols := rows
	if cols*rows < length {
		cols += 1
	}
	if cols*rows < length {
		rows += 1
	}

	sq := newSquare(rows, cols)

	i := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if i < length {
				sq.text[r][c] = runes[i]
			} else {
				sq.text[r][c] = ' '
			}
			i++
		}
	}
	return sq
}

// newSquare allocates memory for a square of size rows x cols
func newSquare(rows, cols int) *square {
	sq := square{
		rows: rows,
		cols: cols,
		text: make([][]rune, rows),
	}
	for r := 0; r < rows; r++ {
		sq.text[r] = make([]rune, cols)
	}
	return &sq
}

// Cipher returns the cipher text for this square
func (sq square) Cipher() string {
	var b strings.Builder
	b.Grow(sq.rows * sq.cols)
	for c := 0; c < sq.cols; c++ {
		for r := 0; r < sq.rows; r++ {
			b.WriteRune(sq.text[r][c])
		}
		if c != sq.cols-1 {
			b.WriteRune(' ')
		}
	}
	return b.String()
}

// String returns the plain text for this square
func (sq square) String() string {
	var b strings.Builder
	b.Grow(sq.rows * sq.cols)
	for r := 0; r < sq.rows; r++ {
		for c := 0; c < sq.cols; c++ {
			b.WriteRune(sq.text[r][c])
		}
	}
	return b.String()
}
