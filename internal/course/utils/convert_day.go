package utils

import (
	"strings"
)

type covertDayErr string

func (e covertDayErr) Error() string {
	return string(e)
}

const (
	dayErr covertDayErr = "invalid day"
)

type Day int

const (
	Sunday    Day = 1
	Monday    Day = 2
	Tuesday   Day = 3
	Wednesday Day = 4
	Thursday  Day = 5
	Friday    Day = 6
	Saturday  Day = 7
)

var dayMap = map[string]Day{
	"sunday":    Sunday,
	"monday":    Monday,
	"tuesday":   Tuesday,
	"wednesday": Wednesday,
	"thursday":  Thursday,
	"friday":    Friday,
	"saturday":  Saturday,
}

func ConvertDayToInteger(day string) (int, error) {
	dayValue, found := dayMap[strings.ToLower(day)]
	if !found {
		return 0, dayErr
	}
	return int(dayValue), nil
}
