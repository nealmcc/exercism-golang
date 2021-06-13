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
	d := data{[]byte(plain)}
	return d.normalize().transpose().format(5)
}

type data struct {
	text []byte
}

// normalize removes any invalid characters, and converts letters to lowercase
func (d data) normalize() data {
	var curr, next int
	for ; next < len(d.text); next++ {
		ch := d.text[next]
		switch {
		case 'a' <= ch && ch <= 'z':
			d.text[curr] = ch
			curr++
		case '0' <= ch && ch <= '9':
			d.text[curr] = ch
			curr++
		case 'A' <= ch && ch <= 'Z':
			d.text[curr] = ch + 'a' - 'A'
			curr++
		}
	}
	d.text = d.text[:curr]
	return d
}

// transpose flips the letters of the text
func (d data) transpose() data {
	for i, char := range d.text {
		if 'a' <= char && char <= 'z' {
			d.text[i] = inverse[char-'a']
		}
	}
	return d
}

var inverse = []byte{'z', 'y', 'x', 'w', 'v', 'u', 't', 's', 'r', 'q', 'p', 'o', 'n', 'm', 'l', 'k', 'j', 'i', 'h', 'g', 'f', 'e', 'd', 'c', 'b', 'a'}

// format returns the cipher text as a space-separated string with chunks
// of the given size. If the size <= 0 then the cipher text is not separated.
func (d data) format(size int) string {
	if size <= 0 {
		return string(d.text)
	}

	// allocate a string builder of sufficient size
	b := new(strings.Builder)
	b.Grow(len(d.text) + len(d.text)/size)

	// write each chunk to the string builder
	for start, end := 0, 0; end < len(d.text); start = end {
		end += size

		// the last chunk could be shorter than 5 bytes
		if end > len(d.text) {
			end = len(d.text)
		}

		b.Write(d.text[start:end])

		// add a space if we have at least one more chunk
		if end < len(d.text) {
			b.WriteByte(' ')
		}
	}
	return b.String()
}
