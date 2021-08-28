package game

import (
	"errors"
	"fmt"

	hg "connect/pkg/hexgrid"
)

// board is a Hex gameboard, with the top left corner at position (0, 0).
// The board grows incrementally by adding East and SE to the top left.
//
// Example board of size 5:
//  . . . . .
//   . . . X .
//    . . . . .
//     . O . . .
//      . . . . .
//
// The top right corner is at position 4*East = (8, 0).
//
// The X is at position SE + 3*East = (7, 1).
//
// The O is at position 3*SE + East = (5, 3).
//
// The bottom left corner is at position 4*SE = (4, 4).
//
// The bottom right corner is at position 4*SE + 4*East = (12, 4).
type board struct {
	size   int
	grid   map[hg.Vkey]shape
	top    hg.Vkey // top edge
	right  hg.Vkey // right edge
	bottom hg.Vkey // bottom edge
	left   hg.Vkey // left edge
}

func newBoard(size int) board {
	b := board{
		size:   size,
		top:    hg.NE,
		right:  hg.East.Times(size),
		bottom: hg.SE.Times(size),
		left:   hg.SW,
		grid:   make(map[hg.Vkey]shape, 4),
	}
	// place the appropriate shape on each edge of the board:
	b.grid[b.top] = shapeX
	b.grid[b.right] = shapeO
	b.grid[b.bottom] = shapeX
	b.grid[b.left] = shapeO
	return b
}

// parseBoard reads a text representation of the board in standard form.
//
// 'X' and 'O' are used to represent shapeX and shapeO.
// '.' is an empty tile.
//
// parseBoard will return an error if the input is empty, not square,
// or contains invalid characters.
func parseBoard(lines []string) (board, error) {
	size := len(lines)
	if size == 0 {
		return board{}, errors.New("a board must have at least 1 tile")
	}

	if size != len(lines[0]) {
		return board{}, errors.New("a board in standard form must be square")
	}

	b := newBoard(size)

	for y, row := range lines {
		key := hg.Vkey{
			X: y,
			Y: y,
		}
		for _, symbol := range []byte(row) {
			switch symbol {
			case 'X':
				b.grid[key] = shapeX
			case 'O':
				b.grid[key] = shapeO
			case '.':
			default:
				return board{}, fmt.Errorf("invalid symbol: %q", symbol)
			}
			key = hg.Sum(key, hg.East)
		}
	}

	return b, nil
}

// hasConnection determines if there is a path between start and end (inclusive)
// where every tile has the given shape.
func (b board) hasConnection(symbol shape, start, end hg.Vkey) bool {
	if b.grid[start] != symbol || b.grid[end] != symbol {
		return false
	}

	// depth-first search. If we find the end then we're done.
	visited := make(map[hg.Vkey]bool)
	visited[start] = true

	work := stack{}
	work.push(start)

	for len(work) > 0 {
		tile, _ := work.pop()
		for _, next := range tile.Neighbours() {
			if next == end {
				return true
			}

			if b.grid[next] != symbol || visited[next] || !b.areAdjacent(tile, next) {
				continue
			}
			visited[next] = true
			work.push(next)
		}
	}

	return false
}

// areAdjacent tests to see if k1 and k2 are adjacent to one another.
// If neither tile is on the board, they are not adjacent.
// If one tile is not on the board, it must be an edge.
// Edges are not adjacent to each other.
func (b board) areAdjacent(k1, k2 hg.Vkey) bool {
	if !b.contains(k1) {
		k1, k2 = k2, k1
	}

	if !b.contains(k1) {
		return false
	}

	// k1 is on the board.
	for _, n := range k1.Neighbours() {
		if n == k2 {
			return true
		}
	}

	return false
}

// contains returns true iff the given key is within the bounds of the board.
func (b board) contains(k hg.Vkey) bool {
	// left edge:
	if k.X < k.Y {
		return false
	}

	// right edge:
	if k.X >= k.Y+2*b.size {
		return false
	}

	// top and bottom edges:
	if k.Y < 0 || k.Y >= b.size {
		return false
	}

	return true
}
