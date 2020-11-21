package chromatic

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
