package common

import (
	"time"
)

const (
	// FmtDateTimeSecondStr - Date Time format used
	// for file names and directory names
	FmtDateTimeSecondStr = "20060102150405"
	// FmtDateTimeNanoSecondStr - Custom Date Time Format
	FmtDateTimeNanoSecondStr = "2006-01-02 15:04:05.000000000"
	// FmtDateTimeSecText - Custom Date Time Format
	FmtDateTimeSecText = "2006-01-02 15:04:05"

	FmtDateTimeTzNano = "01/02/2006 15:04:05.000000000 -0700 MST"

	FmtDateTimeTzSec = "01/02/2006 15:04:05 -0700 MST"

	// FmtDateTimeEverything - Custom Date Time Format showing virtually
	// all elements of a date time string.
	FmtDateTimeEverything = "Monday January 2, 2006 15:04:05.000000000 -0700 MST"
)

// DateTimeUtility - struct used to export
// Date Time Management methods.
type DateTimeUtility struct {
	TimeIn     time.Time
	TimeOut    time.Time
	TimeStart  time.Time
	TimeEnd    time.Time
	Duration   time.Duration
	TimeInStr  string
	TimeOutStr string
	TimeFmtStr string
}

// GetDateTimeStrNowLocal - Gets current
// local time and formats it as a date time
// string
func (dt DateTimeUtility) GetDateTimeStrNowLocal() string {

	return dt.GetDateTimeStr(time.Now().Local())

}

// GetDateTimeStr - Returns a date time string
// in the format 20170427211307
func (dt DateTimeUtility) GetDateTimeStr(t time.Time) string {

	// Time Format down to the second
	return t.Format(FmtDateTimeSecondStr)

}

// GetDateTimeSecText - Returns formatted
// date time with seconds for display,
// 2006-01-02 15:04:05.
func (dt DateTimeUtility) GetDateTimeSecText(t time.Time) string {
	// Time Display Format with seconds
	return t.Format(FmtDateTimeSecText)
}

// GetDateTimeNanoSecText - Returns formated
// date time string with nanoseconds
// 2006-01-02 15:04:05.000000000.
func (dt DateTimeUtility) GetDateTimeNanoSecText(t time.Time) string {
	// Time Format down to the nanosecond
	return t.Format(FmtDateTimeNanoSecondStr)
}

// GetDateTimeEverything - Receives a time value and formats as
// a date time string in the format:
//  Saturday April 29, 2017 19:54:30.123456489 -0500 CDT
func (dt DateTimeUtility) GetDateTimeEverything(t time.Time) string {
	return t.Format(FmtDateTimeEverything)
}

// GetDateTimeCustomFmt - Returns time string
// formatted according to passed in format
// string. Example format string:
// 'Monday January 2, 2006 15:04:05.000000000 -0700 MST'
func (dt DateTimeUtility) GetDateTimeCustomFmt(t time.Time, fmt string) string {
	return t.Format(fmt)
}
