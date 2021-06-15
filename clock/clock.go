package clock

import (
	"fmt"
)

// Clock holds the current wall time, with no day or timezone information,
// and has minute precision.
type Clock interface {
	// String returns the current time in 24-hour format, with leading 0s
	String() string
	// Add the given number of minutes to this clock, and return a copy
	Add(mins int) Clock
	// Subtract the given number of minutes from this clock, and return a copy
	Subtract(mins int) Clock
}

// we will store the clock as the number of minutes past midnight
type clock struct {
	mins int
}

const (
	minute = 1
	hour   = 60 * minute
	day    = 24 * hour
)

// compile-time interface check
var _ Clock = clock{}

// New creates a new clock set to the given time.  Large or negative values
// for the hours and minutes are allowed.
func New(h, m int) Clock {
	c := clock{hour*h + minute*m}
	return c.normal()
}

func (c clock) String() string {
	hh := c.mins / hour
	mm := c.mins % hour
	return fmt.Sprintf("%02d:%02d", hh, mm)
}

func (c clock) Add(mins int) Clock {
	c.mins += mins
	return c.normal()
}

func (c clock) Subtract(mins int) Clock {
	c.mins -= mins
	return c.normal()
}

// normal adjusts the value of the clock to be in the range
// 0 <= c.mins < day
func (c clock) normal() clock {
	c.mins = c.mins % day
	if c.mins < 0 {
		c.mins += day
	}
	return c
}
