package chromatic

// an IScale has a sequence of notes which can be expressed as strings
type IScale interface {
	Describe() []string
}

type scale struct {
	tonic    string
	interval string
	notes    []*note
}

// NewScale generates a scale defined by a tonic and an interval
func NewScale(tonic string, interval string) IScale {
	scale := scale{tonic, interval, make([]*note, len(interval))}
	useFlatNames := useFlatNames(tonic)

	n := notes[tonic]
	for i, delta := range interval {
		scale.notes[i] = n
		n = n.addInterval(delta, useFlatNames)
	}
	return &scale
}

// Describe gets the names of the notes in this scale
func (s *scale) Describe() []string {
	names := make([]string, len(s.notes))
	for i, n := range s.notes {
		names[i] = n.name
	}
	return names
}

func useFlatNames(tonic string) bool {
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		return true
	default:
		return false
	}
}
