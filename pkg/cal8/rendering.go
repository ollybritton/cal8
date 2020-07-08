package cal8

import (
	"fmt"
	"strings"
)

// StringMonth returns a string displaying the calendar for a month.
// Example: c.StringMonth(2020, 7)
//        July 2020
//  Mo Tu We Th Fr Sa Su
//         1  2  3  4  5
//   6  7  8  9 10 11 12
//  13 14 15 16 17 18 19
//  20 21 22 23 24 25 26
//  27 28 29 30 31
// This might just be the messiest code I've ever written.
func (c *Calendar) StringMonth(year, month int) string {
	shortDays := []string{}

	for _, day := range c.Days {
		shortDays = append(shortDays, day[:2])
	}

	dayHeader := strings.Join(shortDays, " ")
	monthHeader := Center(Months[month-1]+" "+fmt.Sprint(year), " ", len(dayHeader))

	startDay := c.Query(year, month, 1)
	startIndex := 0

	for i, day := range c.Days {
		if day == startDay {
			startIndex = i
			break
		}
	}

	dayString := ""

	if startIndex != 0 {
		dayString += strings.Repeat("   ", startIndex)
	}

	var day int
	var monthLength = MonthLens[month-1]

	if month == 2 && (year%400 == 0 || (year%4 == 0 && year%100 != 0)) {
		monthLength++
	}

	for day = 1; day <= len(c.Days)-startIndex; day++ {
		dayString += fmt.Sprintf("%2d ", day)
	}

	dayString += "\n"

	for {
		curr := day

		for ; day < curr+len(c.Days); day++ {
			if day > monthLength {
				break
			}
			dayString += fmt.Sprintf("%2d ", day)
		}

		if day > monthLength {
			break
		}

		dayString += "\n"
	}

	return monthHeader + "\n" + dayHeader + "\n" + dayString
}
