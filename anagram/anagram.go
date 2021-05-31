package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	upperSubj := strings.ToUpper(subject)
	subj := sortable([]byte(upperSubj))
	sort.Sort(subj)

	anagrams := make([]string, 0, len(candidates))

	for _, c := range candidates {
		if len(c) != len(subject) {
			continue
		}

		upperCand := strings.ToUpper(c)
		if upperCand == upperSubj {
			continue
		}

		cand := sortable([]byte(upperCand))
		sort.Sort(cand)
		if bytesMatch(subj, cand) {
			anagrams = append(anagrams, c)
		}
	}

	return anagrams
}

// bytesMatch compares the two byte arrays to see that they are identical.
// It assumes that both inputs are already sorted, and of the same length.
func bytesMatch(a, b []byte) bool {
	for i, x := range a {
		if b[i] != x {
			return false
		}
	}
	return true
}

// sortable is a byte array that implements sort.Interface
type sortable []byte

// compile time interface check
var _ sort.Interface = sortable{}

// Len is part of sort.Interface
func (s sortable) Len() int {
	return len(s)
}

// Less is part of sort.Interface
func (s sortable) Less(i, j int) bool {
	return s[i] < s[j]
}

// Swap is part of sort.Interface
func (s sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
