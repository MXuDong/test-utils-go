package base

import (
	"fmt"
	"github.com/MXuDong/test-utils-go/patch"
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
	setDefaultValue()
}

func setDefaultValue() {
	setTimeToGlobalValue(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	SetDuration(0)
}

// timeValue is a global value.
var timeValue map[string]int = map[string]int{}
var timeLocation *time.Location
var timeDuration time.Duration
var isFreeze bool

//-------------------------------------------------- base function

// GetFreezeTimePoint will return the freeze time point.
func GetFreezeTimePoint() time.Time {
	return blockTimeFunc()
}

//-------------------------------------------------- duration controller

// SetDuration set the time offset value directly.
func SetDuration(d time.Duration) {
	timeDuration = d
}

func GetDuration() time.Duration {
	return timeDuration
}

// CleanDuration set the time offset value to 0 directly
func CleanDuration() {
	SetDuration(0)
}

// AddDuration aAdd an offset to an existing offset
func AddDuration(d time.Duration) {
	SetDuration(time.Duration(timeDuration.Nanoseconds() + d.Nanoseconds()))
}

//-------------------------------------------------- time controller

// FreezeWithTime set time to specify value, and freeze time flow
func FreezeWithTime(year, month, day, hour, minute, second, nano int, location *time.Location) error {
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

func FreezeWithTimeStruct(t time.Time) error {
	if isFreeze {
		return fmt.Errorf("time is freeze now")
	}

	setTime(t)
	patch.Cover(time.Now, blockTimeFunc)
	return nil
}

// FreezeTime locks the time to the current point in time
func FreezeTime() error {
	if isFreeze {
		return fmt.Errorf("time is freeze now")
	}

	setTime(time.Now())
	patch.Cover(time.Now, blockTimeFunc)
	return nil
}

// UnFreezeTime unfreeze time to now, and reset all value to default.
func UnFreezeTime() {
	setDefaultValue()
	patch.Restore(time.Now)
}

// changeLocation change timeLocation for time freeze. Default is time.UTC
func changeLocation(l *time.Location) {
	timeLocation = l
}

// ---------------------------------------------------------------------------------------------------------------------
// define the common method

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
