// Package twelve solves the exercism problem 'Twelve Days of Christmas'
package twelve

import (
	"fmt"
	"strings"
)

var verses = []struct {
	day  string
	gift string
}{
	{"first", "a Partridge in a Pear Tree."},
	{"second", "two Turtle Doves, "},
	{"third", "three French Hens, "},
	{"fourth", "four Calling Birds, "},
	{"fifth", "five Gold Rings, "},
	{"sixth", "six Geese-a-Laying, "},
	{"seventh", "seven Swans-a-Swimming, "},
	{"eighth", "eight Maids-a-Milking, "},
	{"ninth", "nine Ladies Dancing, "},
	{"tenth", "ten Lords-a-Leaping, "},
	{"eleventh", "eleven Pipers Piping, "},
	{"twelfth", "twelve Drummers Drumming, "},
}

// Song return the full song.
func Song() string {
	var b strings.Builder
	for i := 1; i < 12; i++ {
		b.WriteString(Verse(i))
		b.WriteByte('\n')
	}
	b.WriteString(Verse(12))
	return b.String()
}

// Verse returns the requested verse, starting at 1.
func Verse(n int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s",
		verses[n-1].day, gifts(n))
}

func gifts(n int) string {
	isMultiple := n > 1
	var b strings.Builder
	for n > 1 {
		n--
		b.WriteString(verses[n].gift)
	}
	if isMultiple {
		b.WriteString("and ")
	}
	b.WriteString(verses[0].gift)
	return b.String()
}
