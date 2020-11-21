package internal

// IScale is a sequence of notes
type IScale interface {
	GetNotes() []string
}

type scale struct {
	useFlatNames bool
	notes        []*note
}

// NewScale generates a scale beginning with tonic, using the given interval
func NewScale(tonic string, interval string) IScale {
	s := &scale{
		useFlatNames: isTonicFlat(tonic),
		notes:        make([]*note, len(interval)),
	}

	firstNote := notesByName[tonic]
	id := firstNote.id
	for i, delta := range interval {
		s.notes[i] = notes[id]
		id = (id + intervals[delta]) % _count
	}

	return s
}

// each interval determines how many notes to go up by
var intervals = map[rune]int{
	'm': 1,
	'M': 2,
	'A': 3,
}

// GetNotes gets the names of the notes in this scale
func (s *scale) GetNotes() []string {
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
