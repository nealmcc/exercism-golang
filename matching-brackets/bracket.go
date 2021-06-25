package brackets

type bracket byte

const (
	openSquare, closeSquare bracket = '[', ']'
	openRound, closeRound   bracket = '(', ')'
	openCurly, closeCurly   bracket = '{', '}'
)

func isBracket(b bracket) bool {
	switch b {
	case openSquare, openRound, openCurly, closeSquare, closeRound, closeCurly:
		return true
	default:
		return false
	}
}

func (b bracket) isOpen() bool {
	switch b {
	case openSquare, openRound, openCurly:
		return true
	default:
		return false
	}
}

func (b bracket) closes(o bracket) bool {
	switch b {
	case closeSquare:
		return o == openSquare
	case closeRound:
		return o == openRound
	case closeCurly:
		return o == openCurly
	default:
		return false
	}
}
