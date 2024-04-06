package dto

import (
	"strconv"
	"strings"
	"time"
)

const (
	MinYear  = 2000
	MaxYear  = 2100
	MinMonth = 1
	MaxMonth = 12
	MinDay   = 1
	MaxDay   = 31
)

type DDMMYYYY string

func (d DDMMYYYY) Validate() bool {
	separatedString := strings.Split(string(d), "-")
	if len(separatedString) != 3 {
		return false
	}

	day, err := strconv.Atoi(separatedString[0])
	if err != nil || day < MinDay || day > MaxDay {
		return false
	}

	month, err := strconv.Atoi(separatedString[1])
	if err != nil || month < MinMonth || month > MaxMonth {
		return false
	}

	year, err := strconv.Atoi(separatedString[2])
	if err != nil || year < MinYear || year > MaxYear {
		return false
	}

	switch month {
	case 2:
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			return day >= MinDay && day <= 29
		} else {
			return day >= MinDay && day <= 28
		}
	case 4, 6, 9, 11:
		return day >= MinDay && day <= 30
	default:
		return day >= MinDay && day <= 31
	}
}

func (date DDMMYYYY) ToStdDate() time.Time {
	day, _ := strconv.Atoi(string(date[0:2]))
	month, _ := strconv.Atoi(string(date[3:5]))
	year, _ := strconv.Atoi(string(date[6:]))

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
