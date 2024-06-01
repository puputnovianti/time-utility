package timeutilitypackage

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// FormatDate formats a time.Time object to a string based on the provided layout.
func FormatDate(t time.Time, layout string) string {
	return t.Format(layout)
}

// ParseDate parses a date string based on the provided layout and returns a time.Time object.
func ParseDate(dateStr, layout string) (time.Time, error) {
	return time.Parse(layout, dateStr)
}

// DaysBetween calculates the number of days between two dates.
func DaysBetween(start, end time.Time) int {
	duration := end.Sub(start)
	return int(duration.Hours() / 24)
}

// IsWeekend checks if the given date falls on a weekend.
func IsWeekend(t time.Time) bool {
	weekday := t.Weekday()
	return weekday == time.Saturday || weekday == time.Sunday
}

// AddDays adds the specified number of days to the given date.
func AddDays(t time.Time, days int) time.Time {
	return t.AddDate(0, 0, days)
}

// NextWeekday returns the next occurrence of the specified weekday.
func NextWeekday(t time.Time, weekday time.Weekday) (time.Time, error) {
	if weekday < time.Sunday || weekday > time.Saturday {
		return time.Time{}, errors.New("invalid weekday")
	}

	for {
		t = t.AddDate(0, 0, 1)
		if t.Weekday() == weekday {
			return t, nil
		}
	}
}

// ParseTimezone parses a timezone string and returns the corresponding *time.Location.
func ParseTimezone(tz string) (*time.Location, error) {
	location, err := time.LoadLocation(tz)
	if err != nil {
		return nil, err
	}
	return location, nil
}

// ConvertTimezone converts a time.Time object to the specified timezone.
func ConvertTimezone(t time.Time, tz string) (time.Time, error) {
	location, err := ParseTimezone(tz)
	if err != nil {
		return time.Time{}, err
	}
	return t.In(location), nil
}

// IsLeapYear checks if a given year is a leap year.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

// AddMonths adds the specified number of months to the given date.
func AddMonths(t time.Time, months int) time.Time {
	return t.AddDate(0, months, 0)
}

// RecurringEvent calculates the next occurrence of a recurring event based on the given interval (days, weeks, months).
func RecurringEvent(start time.Time, interval string) (time.Time, error) {
	interval = strings.ToLower(interval)
	switch {
	case strings.HasSuffix(interval, "day"):
		days, err := extractNumber(interval)
		if err != nil {
			return time.Time{}, err
		}
		return AddDays(start, days), nil
	case strings.HasSuffix(interval, "week"):
		weeks, err := extractNumber(interval)
		if err != nil {
			return time.Time{}, err
		}
		return AddDays(start, weeks*7), nil
	case strings.HasSuffix(interval, "month"):
		months, err := extractNumber(interval)
		if err != nil {
			return time.Time{}, err
		}
		return AddMonths(start, months), nil
	default:
		return time.Time{}, errors.New("invalid interval format")
	}
}

func extractNumber(s string) (int, error) {
	re := regexp.MustCompile(`\d+`)
	numStr := re.FindString(s)
	if numStr == "" {
		return 0, errors.New("no number found in interval")
	}
	var num int
	fmt.Sscanf(numStr, "%d", &num)
	return num, nil
}
