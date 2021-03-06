package scales

// IScale can be described as a sequence of notes
type IScale interface {
	Describe() []string
}

// Scale is a sequence of notes, and can be generated by a
// tonic and sequence of intervals
type Scale struct {
	Tonic     string
	Intervals string
	Notes     []*Note
}

// Note is a label attached to a pitch.
type Note struct {
	Name  string
	Pitch *Pitch
}

// NewNote creates a new instance of a Note, which labels the given Pitch
func NewNote(n string, p *Pitch) *Note {
	return &Note{n, p}
}

// Pitch is a sound frequency, and can be labelled multiple notes.
type Pitch struct {
	Seq   int
	Sharp *Note
	Flat  *Note
}

// NewPitch creates a new pitch. The sequence should be unique.
func NewPitch(seq int) *Pitch {
	return &Pitch{seq, nil, nil}
}

// NewScale generates a new scale, starting with the given tonic,
// incrementing according to the interval pattern, and
// taken from the given superset of all pitches
func NewScale(tonic string, intervals string, rules IScaleRules) (IScale, error) {
	if intervals == "" {
		intervals = rules.DefaultInterval()
	}
	scale := Scale{
		tonic,
		intervals,
		make([]*Note, len(intervals)),
	}

	superset := rules.AllPitches()
	useFlats := rules.UseFlats(tonic)

	p, err := findPitch(&superset, tonic)
	if err != nil {
		return nil, err
	}

	for i, delta := range intervals {
		if useFlats {
			scale.Notes[i] = p.Flat
		} else {
			scale.Notes[i] = p.Sharp
		}
		p = addInterval(p, delta, rules)
	}
	return &scale, nil
}

// Describe returns the notes in this scale using either
// sharps or flats, depending on the scale's tonic
func (s *Scale) Describe() []string {
	names := make([]string, len(s.Notes))
	for i, n := range s.Notes {
		names[i] = n.Name
	}
	return names
}
