package strain

// Ints is a slice of ints
type Ints []int

// Keep returns a new slice of all the elements where the predicate is true.
// The order of elements is preserved.
func (i Ints) Keep(pred func(int) bool) Ints {
	var kept Ints
	for _, el := range i {
		if pred(el) {
			kept = append(kept, el)
		}
	}
	return kept
}

// Discard returns a new slice of all the elements where the predicate is false.
// The order of elements is preserved.
func (i Ints) Discard(pred func(int) bool) Ints {
	var trash Ints
	for _, el := range i {
		if !pred(el) {
			trash = append(trash, el)
		}
	}
	return trash
}

// Lists is a slice of slice of ints
type Lists [][]int

// Keep returns a new slice of all the elements where the predicate is true.
// The order of elements is preserved.
func (l Lists) Keep(pred func([]int) bool) Lists {
	var kept Lists
	for _, el := range l {
		if pred(el) {
			kept = append(kept, el)
		}
	}
	return kept
}

// Strings is a slice of strings
type Strings []string

// Keep returns a new slice of all the elements where the predicate is true.
// The order of elements is preserved.
func (s Strings) Keep(pred func(string) bool) Strings {
	var kept Strings
	for _, el := range s {
		if pred(el) {
			kept = append(kept, el)
		}
	}
	return kept
}
