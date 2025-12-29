package utils

import "time"

func ConvertStringToTime(dateStr string) time.Time {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, dateStr)
	return t
}

func ConvertTimeToString(date time.Time) string {
	layout := "2006-01-02"
	return date.Format(layout)
}
