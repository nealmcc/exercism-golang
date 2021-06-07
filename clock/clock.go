package clock

import (
	"fmt"
	"time"
)

// a Clock holds the current wall time, with no day or timezone information.
// a Clock has minute precision.
type Clock interface {
	// String returns the current time in 24-hour format, with leading 0s
	String() string
	// Add the given number of minutes to this clock,
	// and return a copy of this clock.
	Add(mins int) Clock
	// Subtract the given number of minutes from this clock,
	// and return a copy of this clock.
	Subtract(mins int) Clock
	// Return the time.Time equivalent of this clock (location is set to UTC)
	Time() time.Time
}

// we will store the clock as the number of minutes past midnight
type clock int

// compile-time interface check
var _ Clock = clock(0)

const (
	minute = 1
	hour   = 60 * minute
	day    = 24 * hour
)

// Create a new clock. Negative values are allowed.
func New(h, m int) Clock {
	t := (hour*h + minute*m) % day
	if t < 0 {
		t += day
	}
	return clock(t)
}

func (c clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours(), c.Minutes())
}

func (c clock) Hours() int {
	return int(c) / hour
}

func (c clock) Minutes() int {
	return int(c) % hour
}

func (c clock) Add(mins int) Clock {
	t := int(c) + mins
	t = t % day
	if t < 0 {
		t += day
	}
	c = clock(t)
	return c
}

func (c clock) Subtract(mins int) Clock {
	return c.Add(-1 * mins)
}

func (c clock) Time() time.Time {
	return time.Date(0 /* year */, 0 /* month */, 0, /* day */
		c.Hours(), c.Minutes(), 0 /* seconds */, 0, /* nanoseconds */
		time.UTC)
}
