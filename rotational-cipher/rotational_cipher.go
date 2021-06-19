package rotationalcipher

func RotationalCipher(plain string, n int) string {
	shift := func(ch byte) byte {
		return (ch + byte(n)) % 26
	}
	s := []byte(plain)
	for i := 0; i < len(s); i++ {
		ch := s[i]
		switch {
		case 'A' <= ch && ch <= 'Z':
			s[i] = shift(ch-'A') + 'A'
		case 'a' <= ch && ch <= 'z':
			s[i] = shift(ch-'a') + 'a'
		}
	}
	return string(s)
}
