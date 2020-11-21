// package listops implements some basic list operations for a list of ints
package listops

type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Append adds all items in the given list to this list
func (left IntList) Append(right IntList) IntList {
	return left
}

// Concat appends all items from all of the given lists to this list
func (left IntList) Concat(right []IntList) IntList {
	return left
}

// Filter returns a new list of all items from this list which match the predicate
func (left IntList) Filter(f predFunc) IntList {
	return left
}

// Length returns the number of items in this list.
func (left IntList) Length() int {
	return len(left)
}

// Map appiles the given function to each item in this list, returning the list of results
func (left IntList) Map(f unaryFunc) IntList {
	return left
}

// Foldl accumulates the items, using the given function and starting value, starting from the left (index 0)
func (left IntList) Foldl(f binFunc, start int) int {
	return start
}

// Foldr accumulates the items, using the given function and starting value, starting from the right (index Length()-1)
func (left IntList) Foldr(f binFunc, start int) int {
	return start
}

// Reverse returns a new list of all items in reverse order.
func (left IntList) Reverse() IntList {
	return left
}
