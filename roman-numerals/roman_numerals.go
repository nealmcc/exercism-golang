package romannumerals

import "errors"

// ToRomanNumeral converts the given decimal to its Roman numeral form
func ToRomanNumeral(n int) (string, error) {
	if n <= 0 || n > 3000 {
		return "", ErrOutOfBounds
	}
	sum := thousands[n/1000]
	n = n % 1000
	sum += hundreds[n/100]
	n = n % 100
	sum += tens[n/10]
	n = n % 10
	sum += ones[n]
	return sum, nil
}

var (
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXX", "XC"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousands = []string{"", "M", "MM", "MMM"}
)

var ErrOutOfBounds = errors.New("decimal must be between 1 and 3000")
