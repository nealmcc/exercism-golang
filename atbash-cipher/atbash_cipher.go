package atbash

import (
	"strings"
)

// Atbash produces the cipher text for the given plaintext
// It ignores anything in the input other than ASCII letters and digits.
// Letters are converted to lowercase, and transposed such that
// 'a' <-> 'z', 'b' <-> 'y' and so on.
// Digits are not transformed.
// the cipher text is divided into 5-byte chunks separated by spaces
func Atbash(plain string) string {
	text := []byte(plain)

	// import the plain text, transposing as we go:
	var curr, next int
	for ; next < len(text); next++ {
		ch := text[next]
		switch {
		case 'a' <= ch && ch <= 'z':
			text[curr] = inverse[ch-'a']
			curr++
		case '0' <= ch && ch <= '9':
			text[curr] = ch
			curr++
		case 'A' <= ch && ch <= 'Z':
			text[curr] = inverse[ch-'A']
			curr++
		}
	}
	text = text[:curr]

	const chunkSize = 5

	// allocate a string builder of sufficient size
	b := new(strings.Builder)
	b.Grow(len(text) + len(text)/chunkSize)

	// write each chunk to the string builder
	for start, end := 0, 0; end < len(text); start = end {
		end += chunkSize

		// the last chunk could be shorter than 5 bytes
		if end > len(text) {
			end = len(text)
		}

		b.Write(text[start:end])

		// add a space if we have at least one more chunk
		if end < len(text) {
			b.WriteByte(' ')
		}
	}
	return b.String()
}

var inverse = []byte{'z', 'y', 'x', 'w', 'v', 'u', 't', 's', 'r', 'q', 'p', 'o', 'n', 'm', 'l', 'k', 'j', 'i', 'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a'}
