package tree

import (
	"errors"
	"sort"
)

// Record is a row in a database representing a message in a forum
type Record struct {
	ID, Parent int
}

// Node is an element of a tree, representing a message in a forum
type Node struct {
	ID       int
	Children []*Node
}

// Build assembles the given records into a tree, and returns the root node
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	// sorting the records makes it easy to find the root node
	// and makes it easy to validate the data
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	r := records[0]
	if r.ID != 0 || r.Parent != 0 {
		return nil, errors.New("bad root node")
	}
	// the nodes are mapped to an array to quickly look them up by id.
	nodes := make([]Node, len(records))
	for id := 1; id < len(records); id++ {
		r := records[id]
		if r.ID != id || r.Parent >= id {
			return nil, errors.New("bad data")
		}
		c, p := &nodes[r.ID], &nodes[r.Parent]
		c.ID = id
		p.Children = append(p.Children, c)
	}
	return &nodes[0], nil
}
