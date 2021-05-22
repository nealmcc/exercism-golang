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
	ones = map[int]string{
		0: "",
		1: "I",
		2: "II",
		3: "III",
		4: "IV",
		5: "V",
		6: "VI",
		7: "VII",
		8: "VIII",
		9: "IX",
	}

	tens = map[int]string{
		0: "",
		1: "X",
		2: "XX",
		3: "XXX",
		4: "XL",
		5: "L",
		6: "LX",
		7: "LXX",
		8: "LXX",
		9: "XC",
	}

	hundreds = map[int]string{
		0: "",
		1: "C",
		2: "CC",
		3: "CCC",
		4: "CD",
		5: "D",
		6: "DC",
		7: "DCC",
		8: "DCCC",
		9: "CM",
	}

	thousands = map[int]string{
		0: "",
		1: "M",
		2: "MM",
		3: "MMM",
	}

	ErrOutOfBounds = errors.New("decimal must be between 1 and 3000")
)
