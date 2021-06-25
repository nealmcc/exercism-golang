package isbn

// IsValidISBN determines if the given string is a valid ISBN
func IsValidISBN(in string) bool {
	mod, digit := 0, 10
	for _, ch := range []byte(in) {
		var n int
		switch ch {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			n = int(ch - '0')
		case 'X', 'x':
			if digit != 1 {
				return false
			}
			n = 10
		case '-':
			continue
		default:
			return false
		}
		mod = (mod + digit*n) % 11
		digit--
	}
	return mod == 0 && digit == 0
}
