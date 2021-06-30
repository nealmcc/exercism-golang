package binarysearch

// SearchInts returns the index of n within ints, or -1 if n is not found.
// ints is assumed to be sorted in ascending order.
func SearchInts(ints []int, n int) int {
	var (
		length = len(ints)
		mid    = length / 2
	)
	switch {
	case length == 0:
		return -1
	case n == ints[mid]:
		return mid
	case n < ints[mid]:
		return SearchInts(ints[:mid], n)
	default:
		ix := SearchInts(ints[mid+1:], n)
		if ix < 0 {
			return -1
		}
		return mid + ix + 1
	}
}
