package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	curTime := time.Now()
	fmt.Println(curTime)

	// Name confort
	jsFormat := "YYYY-MM-dddd HH:mm:ss A z"
	goLayout := ConvertDayjsToGoFormat(jsFormat)
	fmt.Println("JS Format:", jsFormat)
	fmt.Println("Go Layout:", goLayout)
	fmt.Println("Current T:", curTime.Format(goLayout))

	// Create Date
	createDate := time.Date(2004, time.April, 25, 3, 40, 20, 33, time.UTC)
	fmt.Println(createDate.Format(goLayout))

}

func ConvertDayjsToGoFormat(jsFormat string) string {
	replacer := strings.NewReplacer(
		"YYYY", "2006",
		"YY", "06",
		"MMMM", "January",
		"MMM", "Jan",
		"MM", "01",
		"M", "1",
		"DD", "02",
		"D", "2",
		"dddd", "Monday",
		"ddd", "Mon",
		"HH", "15", // 24-hour
		"H", "15",
		"hh", "03", // 12-hour
		"h", "3",
		"mm", "04",
		"m", "4",
		"ss", "05",
		"s", "5",
		"A", "PM",
		"a", "pm",
		"ZZ", "-0700",
		"Z", "-07:00",
		"z", "MST",
	)
	return replacer.Replace(jsFormat)
}

// Formating ---------------

// 2006	Year
// 01	Month (zero-padded)
// 02	Day (zero-padded)
// 15	Hour (24-hour)
// 03	Hour (12-hour)
// 04	Minute
// 05	Second
// PM	AM/PM marker
// Mon	Day of the week
// Jan	Month name
// MST	Time zone

// Built-in Formats --------

// time.RFC3339				2006-01-02T15:04:05Z07:00
// time.RFC1123				Mon, 02 Jan 2006 15:04:05 MST
// time.Kitchen				3:04PM
// time.Stamp				Jan _2 15:04:05
// time.DateTime (Go 1.20+)	2006-01-02 15:04:05
