package atbash

import (
	"atbash/pkg/cipher"
	"fmt"
)

// Atbash encodes the given plain text using the atbash cipher
func Atbash(plain string) string {
	return cipher.New().Encode(plain)
}

func Demo(plain string) {
	// default options
	atbash := cipher.New().Encode(plain)

	// choose a different transpose function
	rotate := cipher.New(
		cipher.WithTranspose(rot),
	).Encode(plain)

	// override all 3 options
	passthrough := cipher.New(
		cipher.WithKeep(yes),
		cipher.WithTranspose(nop),
		cipher.WithBlockSize(0),
	).Encode(plain)

	// a composite cipher (order doesn't matter in this case)
	both1 := cipher.New(
		cipher.WithTranspose(rotbash),
	).Encode(plain)

	both2 := cipher.New(
		cipher.WithTranspose(bashrot),
	).Encode(plain)

	fmt.Printf(`
	passthrough: %s
	atbash:      %s
	rotate:      %s
	both1:       %s
	both2:       %s
`, passthrough, atbash, rotate, both1, both2)
}

func rot(ch byte) byte {
	switch {
	case 'a' <= ch && ch <= 'z':
		return ((ch-'a')+13)%26 + 'a'
	case 'A' <= ch && ch <= 'Z':
		return ((ch-'A')+13)%26 + 'a'
	case '0' <= ch && ch <= '9':
		return ((ch-'0')+5)%10 + '0'
	default:
		return ch
	}
}

func rotbash(ch byte) byte {
	return rot(cipher.AtbashTranspose(ch))
}

func bashrot(ch byte) byte {
	return cipher.AtbashTranspose(rot(ch))
}

func yes(byte) bool    { return true }
func nop(ch byte) byte { return ch }
