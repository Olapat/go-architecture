package timeUtils

import (
	"time"
)

var layoutDate = "2006-01-02"
var layoutDateTime = "2006-01-02 15:04:05"
var layoutDateTimeJson = "2006-01-02T15:04:05.000Z"
var layoutTime = "15:04:05"
var layoutTimeHHMM = "15:04"

func TimeParse(value string, layout string) (time.Time, error) {
	var la = layoutDate
	if layout == "dateTime" {
		la = layoutDateTime
	} else if layout == "dateTimeJson" {
		la = layoutDateTimeJson
	} else if layout == "time" {
		la = layoutTime
	}
	t, err := time.Parse(la, value)
	return t, err
}

func TimeFormat(value time.Time, layout string, withLocal bool) string {
	var la = layoutDate
	if layout == "dateTime" {
		la = layoutDateTime
	} else if layout == "dateTimeJson" {
		la = layoutDateTimeJson
	} else if layout == "time" {
		la = layoutTime
	} else if layout == "timeHHMM" {
		la = layoutTimeHHMM
	}

	f := value.Format(la)
	if withLocal {
		f = value.Local().Format(la)
	}
	return f
}
