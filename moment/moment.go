package moment

import "time"

func BeginOfDay(t ...time.Time) time.Time {
	var finalT time.Time
	if len(t) > 0 {
		finalT = t[0]
	} else {
		finalT = time.Now()
	}
	return time.Date(finalT.Year(), finalT.Month(), finalT.Day(), 0, 0, 0, 0, finalT.Location())
}

func EndOfDay(t ...time.Time) time.Time {
	var finalT time.Time
	if len(t) > 0 {
		finalT = t[0]
	} else {
		finalT = time.Now()
	}
	return time.Date(finalT.Year(), finalT.Month(), finalT.Day(), 23, 59, 59, 999, finalT.Location())
}

func BeginOfMonth(t ...time.Time) time.Time {
	var finalT time.Time
	if len(t) > 0 {
		finalT = t[0]
	} else {
		finalT = time.Now()
	}
	return time.Date(finalT.Year(), finalT.Month(), 1, 0, 0, 0, 0, finalT.Location())
}

func EndOfMonth(t ...time.Time) time.Time {
	var finalT time.Time
	if len(t) > 0 {
		finalT = t[0]
	} else {
		finalT = time.Now()
	}
	year, month, _ := finalT.Date()
	return time.Date(year, month, DayOfMonth(year, int(month)), 23, 59, 59, 999, finalT.Location())
}

func DayOfYear(year int) int {
	if IsLeapYear(year) {
		return 366
	}
	return 365
}

func DayOfMonth(year int, month int) int {
	if month < 1 || month > 12 {
		return 0
	}
	switch month {
	case 2:
		if IsLeapYear(year) {
			return 29
		}
		return 28
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	default:
		return 30
	}
}

func IsLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}
	return false
}
