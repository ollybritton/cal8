package cal8

// Days represents the new names of the week under the cal8 system.
var Days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday", "Valday"}

// NormalDays represents the typical names of the week.
var NormalDays = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

// GetDay returns the day of the week for a given date under the cal8 system.
var GetDay = NewCalendar(Days, 2000, 1, 3)

// GetNormalDay returns the normal day of the week for a given date, using the 7-day week.
var GetNormalDay = NewCalendar(NormalDays, 2000, 1, 3)

// Months is the list of months.
var Months = []string{"January", "Feburary", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

// MonthLens is a list of the length of the months.
// Feburary is assumed to have 28 days. In the code that generates the month view strings, this is accounted for.
var MonthLens = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
