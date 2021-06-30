package robotname

import (
	"errors"
	"math/rand"
	"sync"
)

// Robot has a name
type Robot struct {
	name string
}

// Reset the robot's name
func (r *Robot) Reset() {
	r.name = ""
}

// Name gets the robots's name
func (r *Robot) Name() (string, error) {
	if len(r.name) > 0 {
		return r.name, nil
	}

	next, err := f.name()
	if err != nil {
		return "", err
	}
	r.name = next
	return r.name, nil
}

var f = newFactory()

type factory struct {
	mu      *sync.Mutex
	usedIDs map[int]struct{}
}

func newFactory() *factory {
	return &factory{
		mu:      new(sync.Mutex),
		usedIDs: make(map[int]struct{}),
	}
}

func (f *factory) name() (string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	const max = 26 * 26 * 10 * 10 * 10
	if len(f.usedIDs) >= max {
		return "", errors.New("no more names available")
	}

	var (
		id int
		ok bool
	)
	for !ok {
		id = rand.Intn(max)
		_, didUse := f.usedIDs[id]
		ok = !didUse
	}
	f.usedIDs[id] = struct{}{}

	return idToName(id), nil
}

func idToName(id int) string {
	x, id := 'A'+byte(id%26), id/26
	y, id := 'A'+byte(id%26), id/26
	a, id := '0'+byte(id%10), id/10
	b, id := '0'+byte(id%10), id/10
	c := '0' + byte(id)
	return string([]byte{x, y, a, b, c})
}
