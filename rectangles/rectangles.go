// Package rectangles solves the exercism problem of the same name
package rectangles

// Count returns the number of rectangles in the input.
func Count(rows []string) int {
	numRows := len(rows)
	if numRows < 2 {
		return 0
	}

	numCols := len(rows[0])
	if numCols < 2 {
		return 0
	}

	var num int
	for r1 := 0; r1 < numRows; r1++ {
		for c1 := 0; c1 < numCols; c1++ {
			for r2 := r1 + 1; r2 < numRows; r2++ {
				for c2 := c1 + 1; c2 < numCols; c2++ {

					segr1 := []byte(rows[r1][c1 : c2+1])
					if !isEdge(segr1, '-') {
						continue
					}

					segr2 := []byte(rows[r2][c1 : c2+1])
					if !isEdge(segr2, '-') {
						continue
					}

					segc1 := vslice(rows[r1:r2+1], c1)
					if !isEdge(segc1, '|') {
						continue
					}

					segc2 := vslice(rows[r1:r2+1], c2)
					if isEdge(segc2, '|') {
						num++
					}
				}
			}
		}
	}
	return num
}

func isEdge(segment []byte, connector byte) bool {
	if len(segment) < 2 {
		return false
	}

	const corner = '+'
	if segment[0] != corner {
		return false
	}

	last := len(segment) - 1
	if segment[last] != corner {
		return false
	}

	for _, v := range segment[1:last] {
		if v != connector && v != corner {
			return false
		}
	}

	return true
}

func vslice(rows []string, col int) []byte {
	s := make([]byte, len(rows))
	for i, row := range rows {
		s[i] = row[col]
	}
	return s
}
