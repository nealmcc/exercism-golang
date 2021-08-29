// Package game implements a gameboard for Hex.
// 'X' plays left to right and "O" plays top to bottom.
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
func New(input []string) (*Game, error) {
	b, err := parseBoard(input)
	if err != nil {
		return nil, err
	}

	g := Game{board: b}
	g.setLeft(player{name: "X", shape: shapeX})
	g.setTop(player{name: "O", shape: shapeO})

	return &g, nil
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
	none shape = iota
	shapeX
	shapeO
)

// Game is a game of Hex.
type Game struct {
	playerLeft player
	playerTop  player
	board      board
}

func (g *Game) setLeft(p player) {
	(*g).playerLeft = p
	b := g.board
	b.tiles[b.left] = p.shape
	b.tiles[b.right] = p.shape
}

func (g *Game) setTop(p player) {
	(*g).playerTop = p
	b := g.board
	b.tiles[b.top] = p.shape
	b.tiles[b.bottom] = p.shape
}

// Winner returns the name of winning player.
// Returns false iff neither player has won.
func (g Game) Winner() (string, bool) {
	if g.board.canConnect(g.playerLeft.shape, g.board.left, g.board.right) {
		return g.playerLeft.name, true
	}
	if g.board.canConnect(g.playerTop.shape, g.board.top, g.board.bottom) {
		return g.playerTop.name, true
	}
	return "", false
}
