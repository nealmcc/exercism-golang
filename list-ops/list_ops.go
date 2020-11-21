// package listops implements some basic list operations for a list of ints
package listops

type IntList []int

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

// Append adds all items in the given list to this list
func (self IntList) Append(other IntList) IntList {
	for _, val := range other {
		self = append(self, val)
	}
	return self
}

// Concat appends all items from all of the given lists to this list
func (self IntList) Concat(lists []IntList) IntList {
	for _, other := range lists {
		self = self.Append(other)
	}
	return self
}

// Filter returns a new list of all items from this list which match the predicate
func (self IntList) Filter(f predFunc) IntList {
	filtered := IntList{}
	for _, val := range self {
		if f(val) {
			filtered = append(filtered, val)
		}
	}
	return filtered
}

// Length returns the number of items in this list.
func (self IntList) Length() int {
	return len(self)
}

// Map appiles the given function to each item in this list, returning the list of results
func (self IntList) Map(f unaryFunc) IntList {
	mapped := make(IntList, len(self))
	for i, val := range self {
		mapped[i] = f(val)
	}
	return mapped
}

// Foldl accumulates the items, using the given function and starting value, starting from the self (index 0)
func (self IntList) Foldl(f binFunc, n int) int {
	for _, val := range self {
		n = f(n, val)
	}
	return n
}

// Foldr accumulates the items, using the given function and starting value, starting from the right (index Length()-1)
func (self IntList) Foldr(f binFunc, n int) int {
	for i := len(self) - 1; i >= 0; i-- {
		n = f(self[i], n)
	}
	return n
}

// Reverse returns a new list of all items in reverse order.
func (self IntList) Reverse() IntList {
	length := len(self)
	reversed := make(IntList, length)
	k := length
	for i := 0; i < length; i++ {
		k--
		reversed[k] = self[i]
	}
	return reversed
}
