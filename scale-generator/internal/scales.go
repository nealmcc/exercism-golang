package internal

// IScale is a sequence of notes
type IScale interface {
	GetNames() []string
}

type scale struct {
	useFlatNames bool
	notes        []*note
}

// NewScale generates a scale beginning with tonic, using the given interval
func NewScale(tonic string, interval string) *scale {
	s := &scale{
		useFlatNames: isTonicFlat(tonic),
		notes:        make([]*note, len(interval)),
	}

	firstNote := notesByName[tonic]
	id := firstNote.id
	for i, delta := range interval {
		s.notes[i] = notes[id]
		id = (id + intervals[delta]) % count
	}

	return s
}

// GetNames gets the names of the notes in this scale
func (s *scale) GetNames() []string {
	names := make([]string, len(s.notes))
	for i, n := range s.notes {
		if s.useFlatNames {
			names[i] = n.flatName
		} else {
			names[i] = n.name
		}
	}
	return names
}
