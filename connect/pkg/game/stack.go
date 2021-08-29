package game

import "connect/pkg/hex"

type stack []stackitem

type stackitem struct {
	tile    hex.Vkey
	visited map[hex.Vkey]bool
}

func (s *stack) push(tile hex.Vkey, visited map[hex.Vkey]bool) {
	*s = append(*s, stackitem{
		tile:    tile,
		visited: clone(visited),
	})
}

func clone(m map[hex.Vkey]bool) map[hex.Vkey]bool {
	target := make(map[hex.Vkey]bool, len(m))
	for key, val := range m {
		target[key] = val
	}
	return target
}

func (s *stack) pop() (tile hex.Vkey, visited map[hex.Vkey]bool, ok bool) {
	last := len(*s) - 1
	if last == -1 {
		ok = false
		return
	}

	item := (*s)[last]
	*s = (*s)[:last]

	return item.tile, item.visited, true
}
