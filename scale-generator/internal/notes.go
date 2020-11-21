// package internal hides the 'enum-like' collection of notes and only exposes the scale
package internal

import "strings"

// a note is a label attached to a pitch.
type note struct {
	name  string
	pitch *pitch
}

// a pitch can have one or two notes labelling it. The scale that
// the pitch belongs to will determine which way to describe the pitch
// (either sharp/default or flat)
type pitch struct {
	id    int
	sharp *note
	flat  *note
}

func (p *pitch) getNote(useFlatName bool) *note {
	if useFlatName {
		return p.flat
	}
	return p.sharp
}

func (n *note) addInterval(r rune) *pitch {
	id := n.pitch.id
	next := (id + intervals[r]) % count
	return pitches[next]
}

// each interval determines how many notes to go up by
var intervals = map[rune]int{
	'm': 1,
	'M': 2,
	'A': 3,
}

// count is how many pitches have been defined
var count int

// pitches is our enumeration of pitches
var pitches = make(map[int]*pitch)

// notes is our enumeration of notes
var notes = make(map[string]*note)

// init defines our pitches and notes.
// In this exercise, we use the chromatic scale.
func init() {
	newPitch := func() *pitch {
		p := &pitch{count, nil, nil}
		pitches[count] = p
		count++
		return p
	}

	saveNote := func(n *note) {
		notes[n.name] = n
		notes[strings.ToLower(n.name)] = n
	}

	newNote := func(name string) *pitch {
		p := newPitch()
		p.sharp = &note{name, p}
		p.flat = p.sharp
		saveNote(p.sharp)
		return p
	}

	newHalfNote := func(sharp, flat string) *pitch {
		p := newPitch()
		p.sharp = &note{sharp, p}
		p.flat = &note{flat, p}
		saveNote(p.sharp)
		saveNote(p.flat)
		return p
	}

	newNote("A")
	newHalfNote("A#", "Bb")
	newNote("B")
	newNote("C")
	newHalfNote("C#", "Db")
	newNote("D")
	newHalfNote("D#", "Eb")
	newNote("E")
	newNote("F")
	newHalfNote("F#", "Gb")
	newNote("G")
	newHalfNote("G#", "Ab")
}
