package utils

import "time"

func ConvertStringToTime(dateStr string) *time.Time {
	layout := "2006-01-02T15:04:05Z07:00"
	t, _ := time.Parse(layout, dateStr)
	return &t
}

func ConvertTimeToString(date time.Time) string {
	layout := "2006-01-02T15:04:05Z07:00"
	return date.Format(layout)
}
