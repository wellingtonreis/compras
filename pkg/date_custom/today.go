package date_custom

import "time"

func GetToday() time.Time {
	now := time.Now().UTC()
	year, month, day := now.Date()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	today := time.Date(year, month, day, hour, minute, second, 0, time.FixedZone("", -3*60*60))

	return today
}
