// Package cal8 provides support for a 8-day week calendar, though it's functionality can be used to model a calendar of any period.
//
// The algorithm used is based of a modified version of zeller's formula.
package cal8

import (
	"math"
	"strings"
	"time"
)

// Calendar represents... a calendar. It's modeled as a function which, when given a year, month and day will return the day name.
type Calendar struct {
	Query func(year, month, day int) string
	Days  []string
}

// NewCalendar returns a new calendar from the given options.
// days specifies the names of the days, i.e. []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
// year, month, day specifies a day when the calendar starts. The calendar can still be queried for dates before this, it's just that the date given
// is equal to the first day name specified.
func NewCalendar(days []string, year, month, day int) Calendar {
	length := len(days)
	gradientApprox := float64(3*(31%length)+2*(30%length)) / 5.0

	f := func(year, month, day int) int {
		if month < 3 {
			month += 12
			year--
		}

		monthOffset := int(math.Ceil(
			(float64(month) - 2.0) * gradientApprox,
		))

		yearOffset := year*(365%length) + year/4 - year/100 + year/400

		r := (day + monthOffset + yearOffset) % length

		return r
	}

	// The following code applies the start date. It calculates the offset it needs to add in order to
	// have the dates match up with what the user wanted.
	// There's probably a more mathematic way of doing this but that's a little beyond what I can do I think.
	start := f(year, month, day)

	return Calendar{
		Query: func(year, month, day int) string {
			return days[(f(year, month, day)+(length-start))%length]
		},
		Days: days,
	}
}

// Format parses a Go time format string and returns a result suited to the modified calendar.
// It is identical to time.Time.Format
func (c *Calendar) Format(t time.Time, str string) string {
	res := t.Format(str)
	day := c.Query(t.Year(), int(t.Month()), t.Day())

	for _, dayName := range NormalDays {
		res = strings.ReplaceAll(res, dayName, day)
		res = strings.ReplaceAll(res, dayName[:3], day[:3])
	}

	res = strings.ReplaceAll(res, "%a", t.Weekday().String()[:3])
	res = strings.ReplaceAll(res, "%A", t.Weekday().String())

	return res
}
