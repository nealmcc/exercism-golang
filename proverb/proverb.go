// Package proverb solves the Proverb side exercise from Exercism
package proverb

// Proverb constructs a saying to solve the programming challenge
func Proverb(rhyme []string) []string {
	length := len(rhyme)
	proverb := make([]string, length)
	if length == 0 {
		return proverb
	}
	for i := 0; i < length-1; i++ {
		proverb[i] = "For want of a " + rhyme[i] +
			" the " + rhyme[i+1] + " was lost."
	}
	proverb[length-1] = "And all for the want of a " + rhyme[0] + "."
	return proverb
}
