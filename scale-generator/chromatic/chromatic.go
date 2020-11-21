// package chromatic hides internal implementation of scales,
// so that if we need to rework them due to improved understanding,
// we can change the implemntation without affecting external consumers
package chromatic

import "strings"

var count int
var pitches = make(map[int]*pitch)
var notes = make(map[string]*note)

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

	newNote := func(name string) {
		p := newPitch()
		p.sharp = &note{name, p}
		p.flat = p.sharp
		saveNote(p.sharp)
	}

	newHalfNote := func(sharp, flat string) {
		p := newPitch()
		p.sharp = &note{sharp, p}
		p.flat = &note{flat, p}
		saveNote(p.sharp)
		saveNote(p.flat)
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
