package meetup

import "time"

type WeekSchedule int

const (
	First  WeekSchedule = 1
	Second WeekSchedule = 8
	Third  WeekSchedule = 15
	Fourth WeekSchedule = 22
	Teenth WeekSchedule = 13
	Last   WeekSchedule = -6
)

// Day finds the date for the nth weekday of the given month and year
func Day(nth WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	if nth == Last {
		month++
	}
	minDate := time.Date(year, month, int(nth), 0, 0, 0, 0, time.UTC)
	return after(weekday, minDate).Day()
}

// after finds the first occurrence of the given weekday on or after the date
func after(weekday time.Weekday, minDate time.Time) time.Time {
	delta := weekday - minDate.Weekday()
	if delta < 0 {
		delta += 7
	}
	return minDate.AddDate(0, 0, int(delta))
}
