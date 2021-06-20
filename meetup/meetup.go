package meetup

import "time"

type WeekSchedule uint

const (
	First WeekSchedule = iota + 1
	Second
	Third
	Fourth
	Teenth
	Last
)

// Day finds the date for the nth weekday of the given month and year
func Day(nth WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	var day int
	switch nth {
	case First, Last:
		day = 1
	case Second:
		day = 8
	case Third:
		day = 15
	case Fourth:
		day = 22
	case Teenth:
		day = 13
	}
	minDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
	if nth == Last {
		minDate = minDate.AddDate(0, 1, -7)
	}
	return after(weekday, minDate).Day()
}

func after(weekday time.Weekday, minDate time.Time) time.Time {
	delta := weekday - minDate.Weekday()
	if delta < 0 {
		delta += 7
	}
	return minDate.AddDate(0, 0, int(delta))
}
