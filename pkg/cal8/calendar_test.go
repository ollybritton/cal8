package cal8

import (
	"fmt"
	"testing"
)

// TestCase represents a calendar test case. It contains a date and then the expected result for that date.
type TestCase struct {
	year, month, day int
	expected         string
}

// TestNormal tests that the package can function as a normal calendar.
func TestNormal(t *testing.T) {
	cal := NewCalendar(NormalDays, 2000, 1, 3)

	tests := []TestCase{
		{
			2000, 1, 1,
			"Saturday",
		},
		{
			2000, 2, 1,
			"Tuesday",
		},
		{
			9999, 12, 31,
			"Friday",
		},
		{
			2020, 7, 7,
			"Tuesday",
		},
		{
			2020, 2, 29,
			"Saturday",
		},
		{
			1753, 5, 20,
			"Sunday",
		},
	}

	for _, ts := range tests {
		t.Run(fmt.Sprintf("testing normal %d-%d-%d", ts.year, ts.month, ts.day), func(t *testing.T) {
			got := cal.Query(ts.year, ts.month, ts.day)
			if got != ts.expected {
				t.Errorf("fail got %s want %s", got, ts.expected)
			}
		})
	}
}

// TestNew tests that the package can function as a calendar for the modified cal8 system.
// I can't think of a good test suite for this one, since there's not really a reference I can use.
func TestNew(t *testing.T) {
	cal := NewCalendar(Days, 2000, 1, 1)

	tests := []TestCase{
		{
			2000, 1, 1,
			"Monday",
		},
		{
			2000, 1, 2,
			"Tuesday",
		},
		{
			2000, 1, 3,
			"Wednesday",
		},
		{
			2000, 1, 4,
			"Thursday",
		},
		{
			2000, 1, 5,
			"Friday",
		},
		{
			2000, 1, 6,
			"Saturday",
		},
		{
			2000, 1, 7,
			"Sunday",
		},
		{
			2000, 1, 8,
			"Valday",
		},
	}

	for _, ts := range tests {
		t.Run(fmt.Sprintf("testing new %d-%d-%d", ts.year, ts.month, ts.day), func(t *testing.T) {
			got := cal.Query(ts.year, ts.month, ts.day)
			if got != ts.expected {
				t.Errorf("fail got %s want %s", got, ts.expected)
			}
		})
	}
}
