// Package queenattack solves the exercism problem 'Queen Attack'.
// see: https://exercism.org/tracks/go/exercises/queen-attack
package queenattack

import (
	"errors"
)

// CanQueenAttack accepts the position of two opposing queens on a chessboard,
// and determines if they can attack each other.
// The position of each queen must be given in algebraic notation: a1 to h8,
// and the queens cannot be on the same square.
func CanQueenAttack(w, b string) (bool, error) {
	if w == b {
		return false, errors.New("chess pieces cannot share the same square")
	}

	for _, p := range []string{w, b} {
		if len(p) != 2 {
			return false, errors.New("position must be two characters")
		}
		file, rank := p[0], p[1]
		if file < 'a' || file > 'h' || rank < '1' || rank > '8' {
			return false, errors.New("invalid chessboard location: " + p)
		}
	}

	switch {
	case w[0] == b[0]:
		return true, nil
	case w[1] == b[1]:
		return true, nil
	default:
		dx, dy := int(b[0])-int(w[0]), int(b[1])-int(w[1])
		slope := float64(dy) / float64(dx)
		return slope == 1 || slope == -1, nil
	}
}
