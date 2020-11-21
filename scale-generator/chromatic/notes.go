package chromatic

// a note is a label attached to a pitch.
type note struct {
	name  string
	pitch *pitch
}

// addInterval adds an interval on
func (n *note) addInterval(r interval, useFlatNames bool) *note {
	id := n.pitch.id
	next := (id + sizes[r]) % count
	return pitches[next].getNote(useFlatNames)
}
