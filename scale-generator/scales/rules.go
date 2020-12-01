package scales

import (
	"fmt"
	"strings"
)

// IScaleRules encapsulates the following domain logic:
// - the superset of notes and pitches
// - the default set of intervals to use
// - decision making about when to use sharps or flats to name a pitch
// -
type IScaleRules interface {
	DefaultInterval() string
	IntervalSizes() IntervalMap
	UseFlats(string) bool
	AllPitches() PitchSuperset
}

// PitchSuperset is the set of all pitches.
type PitchSuperset = []*Pitch

// Interval is a single step between notes
type Interval = rune

// An IntervalMap determines how large each interval is
type IntervalMap = map[Interval]int

// findPitch returns the pitch that has a note with the given name (case insensitive)
func findPitch(all *PitchSuperset, name string) (*Pitch, error) {
	for _, p := range *all {
		if strings.EqualFold(p.Sharp.Name, name) ||
			strings.EqualFold(p.Flat.Name, name) {
			return p, nil
		}
	}
	return nil, fmt.Errorf("no pitch found for %s", name)
}

// addInterval returns the pitch that is n above this one,
// continuing into the next octav as necessary
func addInterval(p *Pitch, delta Interval, r IScaleRules) *Pitch {
	sizes := r.IntervalSizes()
	pitches := r.AllPitches()
	count := len(pitches)
	next := (p.Seq + sizes[delta]) % count
	return pitches[next]
}
