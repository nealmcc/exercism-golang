package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// we re-use a read buffer to reduce garbage collection during the benchmark
var readBuf = make([]byte, 1024*4)

func Tally(r io.Reader, w io.Writer) error {
	s := bufio.NewScanner(r)
	s.Buffer(readBuf, len(readBuf))

	l := &leaderboard{teams: make([]team, 0, 4)}
	for s.Scan() {
		if err := s.Err(); err != nil {
			return err
		}

		row := s.Bytes()
		if len(row) == 0 || row[0] == '#' {
			continue
		}

		team1, team2, err := parseRow(row)
		if err != nil {
			return err
		}
		l.add(team1)
		l.add(team2)
	}

	if err := s.Err(); err != nil {
		return err
	}

	l.sort()

	_, err := l.WriteTo(w)
	if err != io.EOF {
		return err
	}
	return nil
}

func parseRow(row []byte) (team, team, error) {
	parts := strings.Split(string(row), ";")
	if len(parts) != 3 {
		return team{}, team{}, errors.New("invalid row - need three parts")
	}

	var (
		t1 team = team{name: parts[0], played: 1}
		t2 team = team{name: parts[1], played: 1}
	)
	switch parts[2] {
	case "win":
		t1.won = 1
		t1.points = 3
		t2.lost = 1
	case "loss":
		t2.won = 1
		t2.points = 3
		t1.lost = 1
	case "draw":
		t1.drew = 1
		t1.points = 1
		t2.drew = 1
		t2.points = 1
	default:
		return team{}, team{}, errors.New("invalid match result")
	}

	return t1, t2, nil
}

type leaderboard struct {
	// given the small number of teams in the tournament, a slice
	// is faster and more memory efficient than a map
	teams     []team
	readIndex int
}

type team struct {
	name   string
	played int
	won    int
	drew   int
	lost   int
	points int
}

func (l *leaderboard) add(t team) {
	ix := -1
	for i, val := range l.teams {
		if t.name == val.name {
			ix = i
			break
		}
	}
	if ix == -1 {
		l.teams = append(l.teams, t)
	} else {
		l.teams[ix].played += t.played
		l.teams[ix].won += t.won
		l.teams[ix].drew += t.drew
		l.teams[ix].lost += t.lost
		l.teams[ix].points += t.points
	}
}

func (l *leaderboard) sort() {
	sort.Slice(l.teams, func(i, j int) bool {
		if l.teams[i].points < l.teams[j].points {
			return false
		}
		if l.teams[i].points > l.teams[j].points {
			return true
		}
		return l.teams[i].name < l.teams[j].name
	})
}

// WriteTo implements io.WriterTo
func (l *leaderboard) WriteTo(w io.Writer) (int64, error) {
	const header = "Team                           | MP |  W |  D |  L |  P\n"
	var n64 int64

	n, err := w.Write([]byte(header))
	n64 += int64(n)
	if err != nil {
		return n64, err
	}

	for _, t := range l.teams {
		// this fprintf call consumes the most memory, but optimising it would
		// cost more in readability than I'm comfortable with.
		n, err := fmt.Fprintf(w, "%-30s | %2d | %2d | %2d | %2d | %2d\n",
			t.name, t.played, t.won, t.drew, t.lost, t.points)
		n64 += int64(n)
		if err != nil {
			return n64, err
		}
	}

	return n64, io.EOF
}
