package luhn

// Valid determines if the given string is a valid account number.
// To be valid, the last digit of the account must equal the one expected
// according to the Luhn algorithm:
// https://en.wikipedia.org/wiki/Luhn_algorithm
func Valid(s string) bool {
	digits := parseDigits(s)
	length := len(digits)
	if length <= 1 {
		return false
	}
	expected := findCheckDigit(digits[:length-1])
	return expected == digits[length-1]
}

// parseDigits converts a string to a slice of the digits in that string.
// only spaces and digits are permitted.
func parseDigits(s string) []byte {
	digits := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		d := s[i]
		switch {
		case d >= '0' && d <= '9':
			digits = append(digits, d&0x0F)
		case d == ' ':
			continue
		default:
			return nil
		}
	}
	return digits
}

// findCheckDigit calculates the correct final digit to append
func findCheckDigit(digits []byte) byte {
	length := len(digits)
	parity := length & 0x01
	sum := 0
	for i := 0; i < length; i++ {
		d := digits[i]
		if i&0x01 != parity {
			d = d << 1
		}
		if d > 9 {
			d -= 9
		}
		sum = sum + int(d)
	}
	return byte(sum * 9 % 10)
}
