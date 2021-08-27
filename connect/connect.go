// Package connect solves the exercism problem 'connect'.
package connect

import (
	"fmt"

	"connect/pkg/game"
)

// ResultOf evaluates a gameboard and determines if player "x"
// or player "o" has won the game.
func ResultOf(lines []string) (string, error) {
	g, err := game.New(lines)
	if err != nil {
		return "", fmt.Errorf("failed to initialise board: %w", err)
	}

	winner, ok := g.Winner()
	if !ok {
		return "", nil
	}

	return winner, nil
}
