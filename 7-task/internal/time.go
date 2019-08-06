package internal

import "time"

// SameDay check if the given times are in the same day or not
func SameDay(t1 time.Time, t2 time.Time) bool {
	utc1 := t1.UTC()
	utc2 := t2.UTC()
	return utc1.Year() == utc2.Year() && utc1.Month() == utc2.Month() && utc1.Day() == utc2.Day()
}
