package internal

// IScale is a sequence of notes
type IScale interface {
	GetNotes() []string
}

type scale []*note

// NewScale generates a scale beginning with tonic, using the given interval
func NewScale(tonic string, interval string) IScale {
	useFlatNames := isScaleFlat(tonic)
	scale := make(scale, len(interval))
	note := notes[tonic]
	pitch := note.pitch
	for i, delta := range interval {
		scale[i] = note
		pitch = note.addInterval(delta)
		note = pitch.getNote(useFlatNames)
	}
	return scale
}

func isScaleFlat(tonic string) bool {
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		return true
	default:
		return false
	}
}

// GetNotes gets the names of the notes in this scale
func (s scale) GetNotes() []string {
	names := make([]string, len(s))
	for i, n := range s {
		names[i] = n.name
	}
	return names
}
