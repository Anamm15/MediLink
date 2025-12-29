package utils

import (
	"fmt"
	"time"
)

func ParseDate(dateStr string) time.Time {
	const layout = "2006-01-02"
	result, err := time.Parse(layout, dateStr)
	if err != nil {
		fmt.Println("Error while parsing date: ", err)
	}
	return result
}

func FormatDate(t time.Time) string {
	const layout = "2006-01-02"
	return t.Format(layout)
}

func ParseTime(timeStr string) time.Time {
	const layout = "15:04"
	result, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error while parsing time: ", err)
	}
	return result
}

func FormatTime(t time.Time) string {
	const layout = "15:04"
	return t.Format(layout)
}

func CombineDateAndTime(dateStr, timeStr string, loc *time.Location) (time.Time, error) {
	if loc == nil {
		loc = time.Local
	}

	date, err := time.ParseInLocation("2006-01-02", dateStr, loc)
	if err != nil {
		return time.Time{}, err
	}

	tm, err := time.Parse("15:04", timeStr)
	if err != nil {
		return time.Time{}, err
	}

	return time.Date(
		date.Year(), date.Month(), date.Day(),
		tm.Hour(), tm.Minute(), 0, 0,
		loc,
	), nil
}
