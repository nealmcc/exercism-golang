package twofer

import "fmt"

// ShareWith solves the 'two-fer' problem in Exercism.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
