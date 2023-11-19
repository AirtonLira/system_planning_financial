package time

import "time"

const layout = "2006-01-02T15:04:05"

// ConvertStringToTime - Function for the convert string to time and return this new time
func ConvertStringToTime(value string) time.Time {

	convertedToTime, _ := time.Parse(layout, value)

	return convertedToTime
}
