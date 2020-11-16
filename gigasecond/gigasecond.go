package gigasecond

import "time"

const giga = 1e9

// AddGigasecond adds 1,000,000,000 seconds to the input
func AddGigasecond(t time.Time) time.Time {
	return t.Add(giga * time.Second)
}
