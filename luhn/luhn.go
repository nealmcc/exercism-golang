package luhn

// Valid determines if the given string is a valid account number
// according to the Luhn algorithm:
// https://en.wikipedia.org/wiki/Luhn_algorithm
func Valid(s string) bool {
	var sum, numDigits int
	for i := len(s) - 1; i >= 0; i-- {
		var d byte = s[i]
		if d == ' ' {
			continue
		}
		d -= '0'
		// bytes are unsigned, so we can use this to rule out invalid digits:
		if d > 9 {
			return false
		}
		if numDigits%2 == 1 {
			d = d << 1
			if d > 9 {
				d -= 9
			}
		}
		sum += int(d)
		numDigits++
	}
	return numDigits > 1 && sum%10 == 0
}
