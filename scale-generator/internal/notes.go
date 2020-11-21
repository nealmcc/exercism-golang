// package internal hides the 'enum-like' collection of notes and only exposes the scale
package internal

import "strings"

// note is our 'enum-like' collection of notes, accessed by id.
// every pitch will have a distinct id, but may be referred
// to by either its 'sharp' name (by default) or else its 'flat' name if
// it has one.
// notes without a 'flat' name will have their 'sharp' name repeated here.
type note struct {
	id       int
	name     string
	flatName string
	pitch    *pitch
}

type pitch struct {
	id    int
	notes *[]note
}

// _count is the number of pitches that we have notes for.
var _count int
var notes = make(map[int]*note)

// notesByName provides a performance shortcut to access notes by their name.
// Both uppercase and lowercase names work for finding the note,
// but the canonical 'names' of the note will still be uppercase.
var notesByName = make(map[string]*note)

// init() defines our collection of notes to be the chromatic scale.
func init() {
	newNote := func(name, flatName string) *note {
		n := note{_count, name, flatName, nil}
		_count++
		notes[n.id] = &n
		return &n
	}

	addByName := func(name string, n *note) {
		notesByName[name] = n
		notesByName[strings.ToLower(name)] = n
	}

	genNote := func(name string) {
		n := newNote(name, name)
		addByName(name, n)
	}

	genHalfNote := func(sharp, flat string) {
		n := newNote(sharp, flat)
		addByName(sharp, n)
		addByName(flat, n)
	}

	genNote("A")
	genHalfNote("A#", "Bb")
	genNote("B")
	genNote("C")
	genHalfNote("C#", "Db")
	genNote("D")
	genHalfNote("D#", "Eb")
	genNote("E")
	genNote("F")
	genHalfNote("F#", "Gb")
	genNote("G")
	genHalfNote("G#", "Ab")
}

// isTonicFlat checks if a scale with this tonic should use 'flat' names
// for half-notes.  If this returns false, the scale will use 'sharp' names.
func isTonicFlat(tonic string) bool {
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		return true
	default:
		return false
	}
}
