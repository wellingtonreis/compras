package date_custom

import "time"

func ConvertDateStringToISO8601(dateStr string) time.Time {
	parsedDate, _ := time.Parse("2006-01-02", dateStr)
	date := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, time.FixedZone("", -3*60*60))
	return date
}
