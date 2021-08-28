// Package game implements a gameboard for Hex.
package game

// New initialises a game using the given board input.
// The input should be normalised to standard form with all spaces removed.
// This will leave a square matrix of characters, all of which are either
// 'X', 'O', or '.'.
//
// For example, this is an empty 3x3 Hex board:
//  ...        . . .
//  ...    ->   . . .
//  ...          . . .
//
// This is a 3x3 Hex board where 'O' has won:
//
//  XOO        X O O
//  OX.    ->   O X .
//  .X.          . X .
//
// 'X' plays top to bottom and "O" plays left to right.
func New(input []string) (*Game, error) {
	b, err := parseBoard(input)
	if err != nil {
		return nil, err
	}

	return &Game{
		p1:    player{name: "x", shape: shapeX},
		p2:    player{name: "o", shape: shapeO},
		board: b,
	}, nil
}

// player is one of the two players.  Each player will be either X's or O's.
type player struct {
	// name is the player's name.
	name string
	// shape is which shape the player uses.
	shape shape
}

// shape is the type of piece that occupies a space on the board.
// The zero-value is an empty space.
type shape int

const (
	none   shape = iota // none is an empty space
	shapeX              // shapeX is an "X"
	shapeO              // shapeO is an "O"
)

// Game is a game of Hex.
type Game struct {
	p1    player
	p2    player
	board board
}

// Winner returns the name of winning player.
// Returns false iff neither player has won.
func (g Game) Winner() (string, bool) {
	if g.board.hasConnection(shapeX, g.board.top, g.board.bottom) {
		return g.p1.name, true
	}
	if g.board.hasConnection(shapeO, g.board.left, g.board.right) {
		return g.p2.name, true
	}
	return "", false
}
