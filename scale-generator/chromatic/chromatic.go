// Package chromatic hides internal implementation of scales,
// so that if we need to rework them due to improved understanding,
// we can change the implemntation without affecting external consumers
package chromatic

import "scale/scales"

// Rules provide the domain logic for the chromatic scale
var Rules chromatic

type chromatic struct{}

func init() {
	newPitch := func() *scales.Pitch {
		p := scales.NewPitch(count)
		count++
		pitches = append(pitches, p)
		return p
	}

	addNote := func(name string) {
		p := newPitch()
		n := scales.NewNote(name, p)
		p.Sharp = n
		p.Flat = n
	}

	addHalfNote := func(sharp, flat string) {
		p := newPitch()
		p.Sharp = scales.NewNote(sharp, p)
		p.Flat = scales.NewNote(flat, p)
	}

	addNote("A")
	addHalfNote("A#", "Bb")
	addNote("B")
	addNote("C")
	addHalfNote("C#", "Db")
	addNote("D")
	addHalfNote("D#", "Eb")
	addNote("E")
	addNote("F")
	addHalfNote("F#", "Gb")
	addNote("G")
	addHalfNote("G#", "Ab")
}

var (
	pitches = []*scales.Pitch{}
	count   = 0
)

func (chromatic) UseFlats(tonic string) bool {
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		return true
	default:
		return false
	}
}

func (chromatic) AllPitches() []*scales.Pitch {
	return pitches
}

func (chromatic) DefaultInterval() string {
	return "mmmmmmmmmmmm"
}

// each interval determines how many notes to go up by
var intervals = map[rune]int{
	'm': 1,
	'M': 2,
	'A': 3,
}

func (chromatic) IntervalSizes() map[rune]int {
	return intervals
}
