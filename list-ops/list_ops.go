// Package listops implements some basic list operations for a list of ints
package listops

type intList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Append adds all items in the given list to this list
func (list intList) Append(other intList) intList {
	for _, val := range other {
		list = append(list, val)
	}
	return list
}

// Concat appends all items from all of the given lists to this list
func (list intList) Concat(lists []intList) intList {
	for _, other := range lists {
		list = list.Append(other)
	}
	return list
}

// Filter returns a new list of all items from this list which match the predicate
func (list intList) Filter(f predFunc) intList {
	filtered := intList{}
	for _, val := range list {
		if f(val) {
			filtered = append(filtered, val)
		}
	}
	return filtered
}

// Length returns the number of items in this list.
func (list intList) Length() int {
	return len(list)
}

// Map appiles the given function to each item in this list, returning the list of results
func (list intList) Map(f unaryFunc) intList {
	mapped := make(intList, len(list))
	for i, val := range list {
		mapped[i] = f(val)
	}
	return mapped
}

// Foldl accumulates the items, using the given function and starting value, starting from the list (index 0)
func (list intList) Foldl(f binFunc, n int) int {
	for _, val := range list {
		n = f(n, val)
	}
	return n
}

// Foldr accumulates the items, using the given function and starting value, starting from the right (index Length()-1)
func (list intList) Foldr(f binFunc, n int) int {
	for i := len(list) - 1; i >= 0; i-- {
		n = f(list[i], n)
	}
	return n
}

// Reverse returns a new list of all items in reverse order.
func (list intList) Reverse() intList {
	length := len(list)
	reversed := make(intList, length)
	k := length
	for i := 0; i < length; i++ {
		k--
		reversed[k] = list[i]
	}
	return reversed
}
