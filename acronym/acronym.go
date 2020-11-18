// Package holds a solution to the Exercism side-exercise of the same name
package acronym

import (
	"regexp"
	"strings"
)

var re = regexp.MustCompile(`[a-zA-Z']+`)

// Abbreviate creates an acronym for the given string
// using the first letter of each word
func Abbreviate(s string) string {
	words := re.FindAllString(s, -1)
	acronym := make([]byte, len(words))
	for i, word := range words {
		acronym[i] = word[0]
	}
	return strings.ToUpper(string(acronym))
}
