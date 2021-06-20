package meetup

import "time"

// WeekSchedule stores an offset from the beginning of a month
type WeekSchedule struct {
	month time.Month
	day   int
}

var (
	First  = WeekSchedule{0, 1}
	Second = WeekSchedule{0, 8}
	Third  = WeekSchedule{0, 15}
	Fourth = WeekSchedule{0, 22}
	Teenth = WeekSchedule{0, 13}
	Last   = WeekSchedule{1, -6}
)

// Day finds the date for the nth weekday of the given month and year
func Day(nth WeekSchedule, weekday time.Weekday, month time.Month, year int) int {
	month = month + nth.month
	day := nth.day
	minDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
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
