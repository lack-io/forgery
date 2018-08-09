package forgery

import (
	"testing"
	"reflect"
	"time"
)

func TestDayOfWeek(t *testing.T) {
	day, err := DayOfWeek()
	if day == "" {
		t.Errorf("DayOfWeek() failed. Got day of weekday: %s.", err.Error())
	}
	for _, v := range days {
		if v == day {
			return
		}
	}
	t.Errorf("DayOfWeek() failed. Got error value: %s.", day)
}

func TestDayOfWeekAbbr(t *testing.T) {
	day, err := DayOfWeekAbbr()
	if day == "" {
		t.Errorf("DayOfWeekAbbr() failed. Got abbreviated day of weekday: %s.", err.Error())
	}
	for _, v := range daysAbbr {
		if v == day {
			return
		}
	}
	t.Errorf("DayOfWeekAbbr() failed. Got error value: %s.", day)
}

func TestMonth(t *testing.T) {
	month, err := Month()
	if month == "" {
		t.Errorf("Month() failed. Got month: %s.", err.Error())
	}
	for _, v := range months {
		if v == month {
			return
		}
	}
	t.Errorf("Month() failed. Got error value: %s.", month)
}

func TestMonthAbbr(t *testing.T) {
	month, err := MonthAbbr()
	if month == "" {
		t.Errorf("Month() failed. Got abbreviated month: %s.", err.Error())
	}
	for _, v := range monthsAbbr {
		if v == month {
			return
		}
	}
	t.Errorf("Month() failed. Got error value: %s.", month)
}

func TestMonthNumber(t *testing.T) {
	monthNum, _ := MonthNumber()
	if monthNum < 0 || monthNum > 12 {
		t.Errorf("MonthNumber() failed. Got month number: %d.", monthNum)
	}
}

func TestYear(t *testing.T) {
	year, _ := Year()
	if reflect.TypeOf(year).String() != "int" {
		t.Errorf("Year() failed. Got value format is not int.")
	}
	now := time.Now().Year()
	if year < now - 20 || year > now + 20 {
		t.Errorf("Year() failed. The value out of range: %d.", year)
	}
}

func TestDay(t *testing.T) {
	day, _ := Day()
	if day < 1 || day > 31 {
		t.Errorf("Day() failed. The value out of range: %d.", day)
	}
}

func TestDate(t *testing.T) {
	date, _ := Date()
	if reflect.TypeOf(date).String() != "time.Time" {
		t.Errorf("Date() failed. Got value format is not `time.Time`: %s.", reflect.TypeOf(date).String())
	}
}