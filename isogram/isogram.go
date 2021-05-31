package isogram

import (
	"unicode"
)

func IsIsogram(s string) bool {
	found := make(map[rune]bool, len(s))
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToLower(r)
		if found[r] {
			return false
		}
		found[r] = true
	}
	return true
}
