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

// matrix is the implementation of a Matrix.
type matrix struct {
	numRows int
	numCols int
	cells   []int
}

var _ Matrix = new(matrix)

// New creates a new matrix.
func New(in string) (*matrix, error) {
	m := matrix{
		cells: make([]int, 0, 16),
	}

	rows := strings.Split(in, "\n")
	m.numRows = len(rows)
	for r, row := range rows {
		if len(row) == 0 && m.numCols == 0 {
			continue
		}

		row = strings.TrimSpace(row)
		cols := strings.Split(row, " ")

		if r == 0 {
			m.numCols = len(cols)
		} else if len(cols) != m.numCols {
			return nil, errors.New("irregular matrix")
		}

		for _, alpha := range cols {
			n, err := strconv.Atoi(alpha)
			if err != nil {
				return nil, err
			}
			m.cells = append(m.cells, n)
		}
	}

	return &m, nil
}

// Rows implements Matrix.Rows.
func (m *matrix) Rows() [][]int {
	rows := make([][]int, m.numRows)

	for i := 0; i < m.numRows; i++ {
		rows[i] = make([]int, m.numCols)
		start, end := i*m.numCols, (i+1)*m.numCols
		copy(rows[i], m.cells[start:end])
	}

	return rows
}

// Cols implements Matrix.Cols.
func (m *matrix) Cols() [][]int {
	cols := make([][]int, m.numCols)

	for i := 0; i < m.numCols; i++ {
		cols[i] = make([]int, m.numRows)
		for r := 0; r < m.numRows; r++ {
			cols[i][r] = m.cells[r*m.numCols+i]
		}
	}

	return cols
}

// Set implements Matrix.Set.
func (m *matrix) Set(r, c, val int) bool {
	if r < 0 || r >= m.numRows {
		return false
	}

	if c < 0 || c >= m.numCols {
		return false
	}

	m.cells[r*m.numCols+c] = val
	return true
}
