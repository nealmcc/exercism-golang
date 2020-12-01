// Package acronym holds a solution to the Exercism side-exercise of the same name
package acronym

import "strings"

// Abbreviate creates an acronym for the given string
// using the first letter of each word
func Abbreviate(s string) string {
	phrase := TrimUntilMatch([]byte(s), IsWordPart)
	var acronym []byte
	for len(phrase) > 0 {
		acronym = append(acronym, phrase[0])
		phrase = TrimUntilMatch(phrase, not(IsWordPart))
		phrase = TrimUntilMatch(phrase, IsWordPart)
	}
	return strings.ToUpper(string(acronym))
}

func isWordPart(b byte) bool {
	return 'a' <= b && b <= 'z' ||
		'A' <= b && b <= 'Z' ||
		'\'' == b
}

func not(m matcher) matcher {
	return func(b byte) bool {
		return !m(b)
	}
}

type matcher func(byte) bool

func trimUntilMatch(s []byte, isMatch matcher) []byte {
	i, length := 0, len(s)
	found := false
	for i < length && !found {
		found = isMatch(s[i])
		i++
	}
	if found {
		return s[i-1:]
	}
	return []byte{}
}
