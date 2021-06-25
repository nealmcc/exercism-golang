package series

// All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	var ss []string
	for i := 0; i+n <= len(s); i++ {
		ss = append(ss, s[i:i+n])
	}
	return ss
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}

// First returns the first substring of s with length n, or false if it cannot
func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
