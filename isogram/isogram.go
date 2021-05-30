package isogram

import "strings"

func IsIsogram(s string) bool {
	found := make(map[rune]bool, len(s))
	s = strings.ToLower(s)
	for _, r := range s {
		if found[r] && r != ' ' && r != '-' {
			return false
		}
		found[r] = true
	}
	return true
}
