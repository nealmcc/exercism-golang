package pascal

// Triangle computes pascal's triangle up to and including the nth row.
func Triangle(n int) [][]int {
	var rows = make([][]int, n)
	for i := 0; i < n; i++ {
		rows[i] = makeRow(i + 1)
	}
	return rows
}

// makeRow calculates the nth (1-based) row of Pascal's triangle
// The row is:
//
//    n   n(n-1)   n(n-1)(n-2)    n(n-1)(n-2)(n-3)
// 1, - , ------ , ----------- , ----------------- ...
//    1     2          2(3)            2(3)(4)
//
func makeRow(n int) []int {
	if n < 1 {
		return nil
	}
	var (
		row        []int = make([]int, n)
		num, denom int   = 1, 1
	)
	// The row is symmetrical, so we calculate from the outside in:
	row[0], row[n-1] = 1, 1
	for j, k := 1, n-2; j <= k; j, k = j+1, k-1 {
		num, denom = num*(k+1), denom*(j)
		v := num / denom
		row[j], row[k] = v, v
	}

	return row
}
