// package internal hides the 'enum-like' collection of notes and only exposes the scale
package internal

import "strings"

type note struct {
	id       int
	name     string
	flatName string
}

// count is the number of distinct pitches that we have notes for.
var count int

// notes is our collection of notes, accessed by id.
// Both A# and Bb will be represented by the same note.
var notes map[int]*note

// notesByName is our collection of notes, accessed by name.
// Both uppercase and lowercase names work for finding the note.
var notesByName map[string]*note

// init() defines our collection of notes to be the chromatic scale.
func init() {
	count = 0
	notes = make(map[int]*note)
	notesByName = make(map[string]*note)

	makeNote := func(name, flatName string) *note {
		n := note{count, name, flatName}
		count++
		notes[n.id] = &n
		return &n
	}

	addByName := func(name string, n *note) {
		notesByName[name] = n
		notesByName[strings.ToLower(name)] = n
	}

	addNote := func(name string) {
		n := makeNote(name, name)
		addByName(name, n)
	}

	addHalfNote := func(sharp, flat string) {
		n := makeNote(sharp, flat)
		addByName(sharp, n)
		addByName(flat, n)
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

// isTonicFlat checks if a scale with this tonic should use 'flat' names
// for half-notes.  If this returns false, the scale will use 'sharp' names.
func isTonicFlat(tonic string) bool {
	switch tonic {
	case "F":
		return true
	case "Bb":
		return true
	case "Eb":
		return true
	case "Ab":
		return true
	case "Db":
		return true
	case "Gb":
		return true
	case "d":
		return true
	case "g":
		return true
	case "c":
		return true
	case "f":
		return true
	case "bb":
		return true
	case "eb":
		return true

	default:
		return false
	}
}
