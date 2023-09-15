package helper

import "time"

func ComposeLike(val string) string {
	charLike := "%" + val + "%"
	if val == "" {
		charLike = "%" + "''" + "%"
	}
	return charLike
}

func StringToDate(val string) time.Time {
	// Define the layout for parsing the input date string
	inputLayout := "02/01/2006" // The layout for the input date string

	// Parse the input date string into a time.Time value
	parsedTime, _ := time.Parse(inputLayout, val)
	location, _ := time.LoadLocation("Asia/Bangkok")
	timeInZone := parsedTime.In(location)
	return timeInZone
}
