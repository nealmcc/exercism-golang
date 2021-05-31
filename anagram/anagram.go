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
		upperCand := strings.ToUpper(c)
		if upperCand == upperSubj || len(c) != len(subj) {
			continue
		}

		isAnagram := true
		cand := sortable([]byte(upperCand))
		sort.Sort(cand)
		for i, b := range cand {
			if subj[i] != b {
				isAnagram = false
				break
			}
		}

		if isAnagram {
			anagrams = append(anagrams, c)
		}
	}

	return anagrams
}

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
