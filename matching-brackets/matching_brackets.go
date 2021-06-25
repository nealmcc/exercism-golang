package brackets

// Bracket is used to check if all brackets in the incoming string match
// and are correctly nested
func Bracket(in string) bool {
	open := &stack{}

	for _, ch := range []byte(in) {
		b := bracket(ch)
		if !isBracket(b) {
			continue
		}

		if b.isOpen() {
			open.push(b)
			continue
		}

		top, ok := open.pop()
		if !ok || !b.closes(top) {
			return false
		}
	}

	return open.len() == 0
}
