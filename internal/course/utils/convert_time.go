package utils

import (
	"time"
)

type covertTimeErr string

func (e covertTimeErr) Error() string {
	return string(e)
}

const (
	allowedStartTime               = "08:00"
	allowedEndTime                 = "22:00"
	timeFormatErr    covertTimeErr = "invalid time format"
	timeRangeErr     covertTimeErr = "invalid time range, should be between 08:00 and 22:00"
)

var timeMap = map[string]int{
	"08:00": 97,
	"08:30": 103,
	"09:00": 109,
	"09:30": 115,
	"10:00": 121,
	"10:30": 127,
	"11:00": 133,
	"11:30": 139,
	"12:00": 145,
	"12:30": 151,
	"13:00": 157,
	"13:30": 163,
	"14:00": 169,
	"14:30": 175,
	"15:00": 181,
	"15:30": 187,
	"16:00": 193,
	"16:30": 199,
	"17:00": 205,
	"17:30": 211,
	"18:00": 217,
	"18:30": 223,
	"19:00": 229,
	"19:30": 235,
	"20:00": 241,
	"20:30": 247,
	"21:00": 253,
	"21:30": 259,
	"22:00": 265,
}

func parseAllowedTime(timeStr string) time.Time {
	t, _ := time.Parse("15:04", timeStr)
	return t
}

func ConvertTimeToInteger(timeStr string) (int, error) {
	parsedTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		return 0, timeFormatErr
	}

	if parsedTime.Before(parseAllowedTime(allowedStartTime)) || parsedTime.After(parseAllowedTime(allowedEndTime)) {
		return 0, timeRangeErr
	}

	minute := parsedTime.Minute()
	if minute < 30 {
		parsedTime = parsedTime.Add(-time.Duration(minute) * time.Minute)
	} else {
		parsedTime = parsedTime.Add((30 - time.Duration(minute)) * time.Minute)
	}

	return timeMap[parsedTime.Format("15:04")], nil
}
