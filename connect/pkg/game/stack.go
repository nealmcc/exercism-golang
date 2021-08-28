package game

import "connect/pkg/hexgrid"

// stack is a stack of hexgrid.Vkey.
type stack []hexgrid.Vkey

// push an element on to the top of the stack.
func (s *stack) push(k hexgrid.Vkey) {
	*s = append(*s, k)
}

// pop an element off the top of the stack if possible.
func (s *stack) pop() (hexgrid.Vkey, bool) {
	last := len(*s) - 1
	if last == -1 {
		return hexgrid.Vkey{}, false
	}

	k := (*s)[last]
	*s = (*s)[:last]

	return k, true
}
