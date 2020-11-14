package twofer

// ShareWith solves the 'two-fer' problem in Exercism.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
