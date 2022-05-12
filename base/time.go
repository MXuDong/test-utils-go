package base

import (
	"fmt"
	"sync"
	"test-utils/patch"
	"time"
)

const (
	_year   = "year"
	_month  = "month"
	_day    = "day"
	_hour   = "hour"
	_minute = "minute"
	_second = "second"
	_nano   = "nano"
)

func init() {
	setTimeToGlobalValue(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	operatorLock = &sync.Mutex{}
}

// timeValue is a global value.
var timeValue map[string]int = map[string]int{}
var timeLocation *time.Location
var timeDuration time.Duration
var operatorLock *sync.Mutex
var isFreeze bool

func setTimeToGlobalValue(y, mo, d, h, mi, s, n int, l *time.Location) {
	timeValue[_year],
		timeValue[_month],
		timeValue[_day],
		timeValue[_hour],
		timeValue[_minute],
		timeValue[_second],
		timeValue[_nano] = y, mo, d, h, mi, s, n
	changeLocation(l)
}

func setTime(t time.Time) {
	setTimeToGlobalValue(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
}

func blockTimeFunc() time.Time {

	return time.Date(getYearValue(), time.Month(getMonthValue()), getDayValue(),
		getHourValue(), getMinuteValue(), getSecondValue(), getNanoValue(), getLocation()).Add(timeDuration)
}

func FreezeWithTime(year, month, day, hour, minute, second, nano int, location *time.Location) error {
	operatorLock.Lock()
	defer operatorLock.Unlock()

	if isFreeze {
		return fmt.Errorf("time is freeze now")
	}

	if location == nil {
		location = time.UTC
	}
	setTimeToGlobalValue(year, month, day, hour, minute, second, nano, location)
	patch.Cover(time.Now, blockTimeFunc)
	return nil
}

func FreezeTime() error {
	operatorLock.Lock()
	defer operatorLock.Unlock()

	if isFreeze {
		return fmt.Errorf("time is freeze now")
	}

	setTime(time.Now())
	patch.Cover(time.Now, blockTimeFunc)
	return nil
}

func UnFreezeTime() {
	operatorLock.Lock()
	defer operatorLock.Unlock()
	patch.Restore(time.Now())
}

// changeLocation change timeLocation for time freeze. Default is time.UTC
func changeLocation(l *time.Location) {
	timeLocation = l
}

// ---------------------------------------------------------------------------------------------------------------------
// define the common method

func getYearValue() int {
	if value, ok := timeValue[_year]; ok {
		return value
	}
	return 0
}
func getMonthValue() int {
	if value, ok := timeValue[_month]; ok {
		return value
	}
	return 0
}
func getDayValue() int {
	if value, ok := timeValue[_day]; ok {
		return value
	}
	return 0
}
func getHourValue() int {
	if value, ok := timeValue[_hour]; ok {
		return value
	}
	return 0
}
func getSecondValue() int {
	if value, ok := timeValue[_second]; ok {
		return value
	}
	return 0
}
func getMinuteValue() int {
	if value, ok := timeValue[_minute]; ok {
		return value
	}
	return 0
}
func getNanoValue() int {
	if value, ok := timeValue[_nano]; ok {
		return value
	}
	return 0
}
func getLocation() *time.Location {
	return timeLocation
}
