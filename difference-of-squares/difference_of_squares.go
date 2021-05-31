package diffsquares

func SquareOfSum(n int) int {
	// 1 + 2 + 3 + 4
	// 4 + 3 + 2 + 1
	// -------------
	// 5 + 5 + 5 + 5
	// == (n)(n+1)
	// == 2x(sum of 1..4)
	sum := n * (n + 1) / 2
	sq := sum * sum
	return sq
}

func SumOfSquares(n int) int {
	// https://proofwiki.org/wiki/Sum_of_Sequence_of_Squares
	sum := n * (n + 1) * (2*n + 1) / 6
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
