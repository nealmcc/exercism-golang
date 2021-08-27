// Package game implements a gameboard for Hex.
package game

import (
	"errors"

	"connect/pkg/hexgrid"
)

// New initialises a game using the given starting board.
// The board may have pieces on it already.
// Player 1 plays from top to bottom.
// Player 2 plays from left to right.
func New(lines []string) (*Game, error) {
	b, err := parseBoard(lines)
	if err != nil {
		return nil, err
	}

	return &Game{
		p1:    player{name: "x", shape: ex},
		p2:    player{name: "o", shape: oh},
		board: b,
	}, nil
}

// player is one of the two players.  Each player will be either EX's or OH's.
type player struct {
	// name is the player's name.
	name string
	// shape is which shape the player uses.
	shape shape
}

// Game is a game of Hex.
type Game struct {
	p1    player
	p2    player
	board board
}

// Winner returns the name of winning player if there is one and false if not.
func (g Game) Winner() (string, bool) {
	return "", false
}

// board is a Hex gameboard.
// The top left corner of the board is position 0, 0.
// Any tile with x = 0 or x = 1 connects to the left edge of the board.
// Any tile with x = width or x = width-1 connext to the right edge.
// Any tile with y = 0 connects to the top edge of the board.
// Any tile with y = height connects to the bottom edge of the board.
type board struct {
	width  int
	height int
	grid   map[hexgrid.Vkey]shape
}

// shape is the type of piece that occupies a space on the board.
// The zero-value is an empty space.
type shape int

const (
	// none is an empty space
	none shape = iota
	// ex is an "X"
	ex
	// oh is an "O"
	oh
)

func parseBoard(lines []string) (board, error) {
	return board{}, errors.New("not implemented")
}
