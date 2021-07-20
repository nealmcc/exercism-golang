// Package house solves the exercism problem of the same name
package house

import (
	"fmt"
	"strings"
)

// Song returns all verses of the song.
func Song() string {
	s := make([]string, 12)
	for i := 1; i <= 12; i++ {
		s[i-1] = verses[i].String()
	}
	return strings.Join(s, "\n\n")
}

// Verse returns the nth verse of the song, starting at 1.
func Verse(n int) string {
	return verses[n].String()
}

var (
	verses = []*verse{
		nil, house, malt, rat, cat, dog, cow,
		maiden, man, priest, rooster, farmer, horse,
	}

	horse = &verse{
		noun: "horse and the hound and the horn",
		verb: "belonged to",
		next: farmer,
	}

	farmer = &verse{
		noun: "farmer",
		adj:  "sowing his corn",
		verb: "kept",
		next: rooster,
	}

	rooster = &verse{
		noun: "rooster",
		adj:  "that crowed in the morn",
		verb: "woke",
		next: priest,
	}

	priest = &verse{
		noun: "priest",
		adj:  "all shaven and shorn",
		verb: "married",
		next: man,
	}

	man = &verse{
		noun: "man",
		adj:  "all tattered and torn",
		verb: "kissed",
		next: maiden,
	}

	maiden = &verse{
		noun: "maiden",
		adj:  "all forlorn",
		verb: "milked",
		next: cow,
	}

	cow = &verse{
		noun: "cow",
		adj:  "with the crumpled horn",
		verb: "tossed",
		next: dog,
	}

	dog = &verse{
		noun: "dog",
		verb: "worried",
		next: cat,
	}

	cat = &verse{
		noun: "cat",
		verb: "killed",
		next: rat,
	}

	rat = &verse{
		noun: "rat",
		verb: "ate",
		next: malt,
	}

	malt = &verse{
		noun: "malt",
		verb: "lay in",
		next: house,
	}

	house = &verse{
		noun: "house",
		adj:  "that Jack built.",
	}
)

type verse struct {
	noun string
	adj  string
	verb string
	next *verse
}

// compile-time interface check
var _ fmt.Stringer = verse{}

// String implements fmt.Stringer
func (v verse) String() string {
	var b strings.Builder
	b.WriteString("This is the ")
	v.writeNounPhrase(&b)
	v.writeActions(&b)
	return b.String()
}

func (v verse) writeNounPhrase(b *strings.Builder) {
	b.WriteString(v.noun)
	if len(v.adj) > 0 {
		b.WriteByte(' ')
		b.WriteString(v.adj)
	}
}

func (v verse) writeActions(b *strings.Builder) {
	if v.next == nil {
		return
	}
	b.WriteByte('\n')
	b.WriteString("that ")
	b.WriteString(v.verb)
	b.WriteString(" the ")
	v.next.writeNounPhrase(b)
	v.next.writeActions(b)
}
