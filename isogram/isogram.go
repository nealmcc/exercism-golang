package isogram

import (
	"bufio"
	"io"
	"strings"
)

// IsIsogram determines if the given string contains at most one of each letter
// in the english alphabet. Any other symbol or character is ignored.
// This function just exists to facilitate the tests from the exercise.
func IsIsogram(s string) bool {
	isIso, err := IsIso(strings.NewReader(s))
	if err != nil {
		return false
	}
	return isIso
}

// IsIso determines if any character in the english alphabet appears more than
// once in the stream. Characters other than a-z and A-Z are ignored.
// I've written this as a stream handler, to simulate the situation where we're
// providing some public API that needs to be resistant to denial of service
// attacks in the form of very large input.
// This function has an upper limit on the amount of memory it can consume,
// regardless of the amount of data on the stream.
func IsIso(stream io.Reader) (bool, error) {
	r := bufio.NewReader(stream)
	// 0=a, 25=z
	found := [26]bool{}
	for {
		rune, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				return true, nil
			}
			return false, err
		}
		if rune < 'A' ||
			rune > 'Z' && rune < 'a' ||
			rune > 'z' {
			// skip any non-english letters
			continue
		}
		var i int
		if rune <= 'Z' {
			i = int(rune - 'A')
		} else {
			i = int(rune - 'a')
		}
		if found[i] {
			return false, nil
		}
		found[i] = true
	}
}
