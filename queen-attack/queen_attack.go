// Package queenattack solves the exercism problem 'Queen Attack'.
// see: https://exercism.org/tracks/go/exercises/queen-attack
package queenattack

import (
	"errors"
	"math"
)

// CanQueenAttack accepts the position of two opposing queens on a chessboard,
// and determines if they can attack each other.  The position of the queens
// are given in standard chess notation: a1 to h8.  If either piece is off
// the board, or the chess notation is not used, CanQueenAttack will return
// an error.
func CanQueenAttack(white, black string) (attack bool, err error) {
	var v1, v2 vector

	if v1, err = getCoord(white); err != nil {
		return false, err
	}

	if v2, err = getCoord(black); err != nil {
		return false, err
	}

	switch slope := v1.minus(v2).slope(); true {
	case slope == 0:
		// same rank
		return true, nil
	case slope == 1, slope == -1:
		// same file
		return true, nil
	case math.IsInf(slope, 1) || math.IsInf(slope, -1):
		// a diagonal
		return true, nil
	case math.IsNaN(slope):
		// same square
		return false, errors.New("chess pieces cannot share the same square")
	default:
		return false, nil
	}
}

// getCoord interprets the location on a chessboard as an integer coordinate on
// the x,y plane where a1 is 0,0 and h8 is 7,7.
// getCoord returns an error if the input string does not follow chess notation,
// or is off the board.
func getCoord(pos string) (vector, error) {
	if len(pos) != 2 {
		return vector{}, errors.New("position must be two characters")
	}

	file, rank := pos[0], pos[1]

	if file < 'a' || file > 'h' || rank < '1' || rank > '8' {
		return vector{}, errors.New("invalid chessboard location: " + pos)
	}

	return vector{
		x: int(file - 'a'),
		y: int(rank - '1'),
	}, nil
}

// vector is a 2D integer coordinate on the x,y plane.
// vector is also the direction from one coordinate to another.
type vector struct {
	x, y int
}

// minus returns the difference v - v2.
func (v vector) minus(v2 vector) vector {
	return vector{
		x: v.x - v2.x,
		y: v.y - v2.y,
	}
}

// slope determines the slope of a line that intersects 0,0 and (v.x, v.y).
func (v vector) slope() float64 {
	switch {
	case v.x == 0 && v.y == 0:
		return math.NaN()
	case v.x == 0:
		return math.Inf(v.y)
	default:
		return float64(v.y) / float64(v.x)
	}
}
