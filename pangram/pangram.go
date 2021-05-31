package pangram

// IsPangram returns true iff s contains every letter in the English alphabet
func IsPangram(s string) bool {
	found := [26]bool{}

	// using a for.. i loop iterates by byte, not rune
	for i := 0; i < len(s); i++ {
		l := getAlphaIndex(s[i])
		if l <= 25 {
			found[l] = true
		}
	}

	for _, x := range found {
		if !x {
			return false
		}
	}
	return true
}

// getAlphaIndex returns a number between 0..25, indicating which letter
// of the english alphabet is represented by b. if the result is >25 then
// the byte is not an english letter
func getAlphaIndex(b byte) byte {
	// the only difference between lowercase and uppercase letters is that the
	// '32s bit' is set for lowercase, and unset for uppercase.
	// also, bytes are inherently unsigned, so a negative value will become > 25
	return b | 1<<5 - 'a'
}
