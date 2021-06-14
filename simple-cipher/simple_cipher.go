package cipher

import (
	sc "cipher/pkg/shift_cipher"
)

// NewCaesar creates a cipher which shifts its input letters by 3
func NewCaesar() Cipher {
	return sc.New(sc.WithShift(3))
}

// NewShift creates a cipher which shifts its input letter by the given amount
func NewShift(n int) Cipher {
	if n < -25 || n == 0 || n > 25 {
		return nil
	}

	return sc.New(sc.WithShift(n))
}

// NewVigenere creates a cipher which shifts each letter a variable amount
// as determined by the input pattern.
func NewVigenere(p string) Cipher {
	var ok bool

	for _, b := range p {
		if b != 'a' {
			ok = true
		}
		if b < 'a' || b > 'z' {
			return nil
		}
	}

	if !ok {
		return nil
	}

	return sc.New(sc.WithPattern([]byte(p)))
}
