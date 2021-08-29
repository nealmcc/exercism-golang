package game

import (
	"errors"
	"fmt"

	"connect/pkg/hex"
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
	width  int                // boards should have equal width and height for
	height int                // fair play but this is not enforced.
	top    hex.Vkey           // nominal top edge
	right  hex.Vkey           // nominal right edge
	bottom hex.Vkey           // nominal bottom edge
	left   hex.Vkey           // nominal left edge
	tiles  map[hex.Vkey]shape // all tiles which have a shape. Includes edges.
}

// newBoard initializes a new, empty board of the given size.
func newBoard(width, height int) board {
	b := board{
		width:  width,
		height: height,
		top:    hex.NE,
		right:  hex.East.Times(width),
		bottom: hex.SE.Times(height),
		left:   hex.SW,
		tiles:  make(map[hex.Vkey]shape, 4),
	}
	// place the appropriate shape on each edge of the board:
	b.tiles[b.top] = shapeO
	b.tiles[b.right] = shapeX
	b.tiles[b.bottom] = shapeO
	b.tiles[b.left] = shapeX
	return b
}

// parseBoard reads a text representation of the board.
//
// 'X' and 'O' are the shapes used by playerLeft and playerTop, respectively.
// '.' is an empty tile.
//
// parseBoard will return an error if the input is empty, not a rhombus,
// or contains invalid characters.
func parseBoard(lines []string) (board, error) {
	height := len(lines)
	if height == 0 {
		return board{}, errors.New("a board must have at least 1 tile")
	}

	width := len(lines[0])

	b := newBoard(width, height)

	for y, row := range lines {
		if len(row) != width {
			return board{}, errors.New("a board must be a rhombus")
		}
		key := hex.SE.Times(y)
		for _, symbol := range []byte(row) {
			switch symbol {
			case 'X':
				b.tiles[key] = shapeX
			case 'O':
				b.tiles[key] = shapeO
			case '.':
			default:
				return board{}, fmt.Errorf("invalid symbol: %q", symbol)
			}
			key = key.Plus(hex.East)
		}
	}

	return b, nil
}

// canConnect determines if there is a path between start and end (inclusive)
// where every tile has the given shape.
func (b board) canConnect(symbol shape, start, end hex.Vkey) bool {
	if symbol == none || b.tiles[start] != symbol || b.tiles[end] != symbol {
		return false
	}

	type visits = map[hex.Vkey]bool

	// depth-first search.
	work := stack{{start, visits{start: true}}}

	for len(work) > 0 {
		tile, visited, _ := work.pop()
		for _, next := range b.children(tile) {
			if next == end {
				return true
			}

			if visited[next] {
				continue
			}

			visited[next] = true
			work.push(next, visited)
		}
	}

	return false
}

// children returns a slice of tiles that have the same shape as k,
// and are adjacent to it.
func (b board) children(k hex.Vkey) []hex.Vkey {
	match := b.tiles[k]
	if match == none {
		return nil
	}

	children := make([]hex.Vkey, 0, 3)
	for k2, shape := range b.tiles {
		if shape == match && b.areAdjacent(k, k2) {
			children = append(children, k2)
		}
	}
	return children
}

// areAdjacent tests to see if k1 and k2 are adjacent to one another.
// If neither tile is on the board, they are not adjacent.
// If one tile is not on the board, it must be an edge.
// Edges are not adjacent to each other.
func (b board) areAdjacent(k1, k2 hex.Vkey) bool {
	if !b.allValid(k1, k2) {
		return false
	}

	// both k1 and k2 are either internal or an edge

	if b.isEdge(k1) {
		k1, k2 = k2, k1
	}

	if b.isEdge(k1) {
		return false
	}

	// k1 is internal. k2 is either an edge or internal.

	switch k2 {
	case b.top:
		for n := 0; n < b.width; n++ {
			if k1 == hex.East.Times(n) {
				return true
			}
		}

	case b.right:
		for n := 0; n < b.height; n++ {
			if k1 == hex.East.Times(b.width-1).Plus(hex.SE.Times(n)) {
				return true
			}
		}

	case b.bottom:
		for n := 0; n < b.width; n++ {
			if k1 == hex.SE.Times(b.height-1).Plus(hex.East.Times(n)) {
				return true
			}
		}

	case b.left:
		for n := 0; n < b.height; n++ {
			if k1 == hex.SE.Times(n) {
				return true
			}
		}

	default:
		// k1 and k2 are both internal.
		for _, n := range k1.Neighbours() {
			if n == k2 {
				return true
			}
		}
	}
	return false
}

// allValid returns true if all of the given tiles are either on the receiver's
// board, or are one of its four edges.
func (b board) allValid(keys ...hex.Vkey) bool {
	for _, k := range keys {
		if !b.isInternal(k) && !b.isEdge(k) {
			return false
		}
	}
	return true
}

// isInternal returns true iff the given key is within the bounds of the
// receiver's board.
func (b board) isInternal(k hex.Vkey) bool {
	// left edge or beyond:
	if k.X < k.Y {
		return false
	}

	// right edge or beyond:
	if k.X >= k.Y+2*b.width {
		return false
	}

	// top and bottom edges or beyond:
	if k.Y < 0 || k.Y >= b.height {
		return false
	}

	return true
}

// isEdge returns true iff k is one of the four edges of the receiver.
func (b board) isEdge(k hex.Vkey) bool {
	switch k {
	case b.top, b.right, b.bottom, b.left:
		return true
	default:
		return false
	}
}
