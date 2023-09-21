package utils

import "time"

func GetDate(day int, month time.Month, year int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
