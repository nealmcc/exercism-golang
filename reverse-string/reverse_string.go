package reverse

// Reverse returns the reverse of the given string
func Reverse(s string) string {
	runes := []rune(s)
	length := len(runes)
	rev := make([]rune, length)
	for i, r := range runes {
		rev[length-i-1] = r
	}
	return string(rev)
}
