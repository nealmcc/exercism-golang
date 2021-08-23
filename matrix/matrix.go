// Package matrix solves the exercism problem 'matrix'
package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Matrix is a regular two-dimensional array of integers.
type Matrix [][]int

// New creates a new Matrix.
// If the input contains no cells, the Matrix will be an empty (valid) matrix.
func New(in string) (Matrix, error) {
	m := make([][]int, 0, 4)

	rows := strings.Split(in, "\n")
	numCols := 0
	for r, row := range rows {
		if len(row) == 0 && numCols == 0 {
			continue
		}

		row = strings.TrimSpace(row)
		cols := strings.Split(row, " ")

		if r == 0 {
			numCols = len(cols)
		} else if len(cols) != numCols {
			return nil, errors.New("irregular matrix")
		}

		cells := make([]int, 0, 4)

		for _, alpha := range cols {
			n, err := strconv.Atoi(alpha)
			if err != nil {
				return nil, fmt.Errorf("a matrix only supports integers: %w", err)
			}
			cells = append(cells, n)
		}
		m = append(m, cells)
	}

	return m, nil
}

// Rows returns the rows of the matrix
func (m Matrix) Rows() [][]int {
	numRows, numCols := m.Size()
	if numRows == 0 {
		return nil
	}

	rows := make([][]int, numRows)

	for r := 0; r < numRows; r++ {
		rows[r] = make([]int, numCols)
		copy(rows[r], m[r])
	}

	return rows
}

// Cols returns the columns of the matrix.
func (m Matrix) Cols() [][]int {
	numRows, numCols := m.Size()
	if numCols == 0 {
		return nil
	}

	cols := make([][]int, numCols)

	for c := 0; c < numCols; c++ {
		cols[c] = make([]int, numRows)
		for r := 0; r < numRows; r++ {
			cols[c][r] = m[r][c]
		}
	}

	return cols
}

// Set attempts to set the given row and column.  Returns true if successful.
func (m Matrix) Set(r, c, val int) bool {
	numRows, numCols := m.Size()

	if r < 0 || r >= numRows {
		return false
	}

	if c < 0 || c >= numCols {
		return false
	}

	m[r][c] = val
	return true
}

// Size returns the number of rows and columns in the matrix.
func (m Matrix) Size() (rows int, cols int) {
	rows = len(m)
	if rows == 0 {
		return
	}
	cols = len(m[0])
	return
}
