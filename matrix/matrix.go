// Package matrix solves the exercism problem 'matrix'
package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix is the interface for a matrix of integers.
type Matrix interface {

	// Rows returns a slice of all the rows in the matrix.
	Rows() [][]int

	// Cols returns a slice of all the columns in the matrix.
	Cols() [][]int

	// Set the value of a cell in the matrix.
	Set(r, c, val int) bool
}

// Grid is an implementation of a Matrix using a single slice for the cells.
type Grid struct {
	numRows int
	numCols int
	cells   []int
}

var _ Matrix = new(Grid)

// New creates a new Grid.
func New(in string) (*Grid, error) {
	g := Grid{
		cells: make([]int, 0, 16),
	}

	rows := strings.Split(in, "\n")
	g.numRows = len(rows)
	for r, row := range rows {
		if len(row) == 0 && g.numCols == 0 {
			continue
		}

		row = strings.TrimSpace(row)
		cols := strings.Split(row, " ")

		if r == 0 {
			g.numCols = len(cols)
		} else if len(cols) != g.numCols {
			return nil, errors.New("irregular matrix")
		}

		for _, alpha := range cols {
			n, err := strconv.Atoi(alpha)
			if err != nil {
				return nil, err
			}
			g.cells = append(g.cells, n)
		}
	}

	return &g, nil
}

// Rows implements Matrix.Rows.
func (g *Grid) Rows() [][]int {
	rows := make([][]int, g.numRows)

	for i := 0; i < g.numRows; i++ {
		rows[i] = make([]int, g.numCols)
		start, end := i*g.numCols, (i+1)*g.numCols
		copy(rows[i], g.cells[start:end])
	}

	return rows
}

// Cols implements Matrix.Cols.
func (g *Grid) Cols() [][]int {
	cols := make([][]int, g.numCols)

	for i := 0; i < g.numCols; i++ {
		cols[i] = make([]int, g.numRows)
		for r := 0; r < g.numRows; r++ {
			cols[i][r] = g.cells[r*g.numCols+i]
		}
	}

	return cols
}

// Set implements Matrix.Set.
func (g *Grid) Set(r, c, val int) bool {
	if r < 0 || r >= g.numRows {
		return false
	}

	if c < 0 || c >= g.numCols {
		return false
	}

	g.cells[r*g.numCols+c] = val
	return true
}
