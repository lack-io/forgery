package forgery

import (
	"math/rand"
	"time"
)

var (
	days []string = []string{
		"Monday", "Tuesday", "Wednesday", "Thursday",
		"Friday", "Saturday", "Sunday",
	}

	daysAbbr []string = []string{
		"Mon", "Tue", "Wed", "Thu", "Fri", "Sat",
	}

	months []string = []string{
		"January", "February", "March", "April",
		"May", "June", "July", "August",
		"September", "October", "November", "December",
	}

	monthsAbbr []string = []string{
		"Jan", "Feb", "Mar", "Apr",
		"May", "Jun", "Jul", "Aug",
		"Sep", "Oct", "Nov", "Dec",
	}
)

// Random day of week name.
func DayOfWeek() (string, error) {
	return random(days), nil
}

// Random abbreviated day of week name.
func DayOfWeekAbbr() (string, error) {
	return random(daysAbbr), nil
}

// Random month name.
func Month() (string, error) {
	return random(months), nil
}

// Random month abbreviated name.
func MonthAbbr() (string, error) {
	return random(monthsAbbr), nil
}

// Random month number
func MonthNumber() (int, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(months)), nil
}

// Random year (this year +/- 20).
func Year() (int, error) {
	rand.Seed(time.Now().UnixNano())
	return time.Now().Year() + rand.Intn(40) - 20, nil
}


// Random day number.
func Day() (int, error) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(30) + 1, nil
}

// Random date (today +/- 20d)
func Date() (time.Time, error) {
	delta := rand.Intn(40) - 20
	return time.Now().AddDate(0, 0, delta), nil
}