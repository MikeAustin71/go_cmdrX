package common

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	// MicroSecondNanoseconds - Number of Nanoseconds in a Microsecond
	MicroSecondNanoseconds = int64(1000)
	// MilliSecondNanoseconds - Number of Nanoseconds in a MilliSecond
	MilliSecondNanoseconds = int64(1000 * 1000)
	// SecondNanoseconds - Number of Nanoseconds in a Second
	SecondNanoseconds = int64(1000 * 1000 * 1000)
	// MinuteNanoSeconds - Number of Nanoseconds in a minute
	MinuteNanoSeconds = int64(1000 * 1000 * 1000 * 60)
	// HourNanoSeconds - Number of Nanoseconds in an hour
	HourNanoSeconds = int64(1000 * 1000 * 1000 * 60 * 60)
	// DayNanoSeconds - Number of Nanoseconds in a 24-hour day
	DayNanoSeconds = int64(1000 * 1000 * 1000 * 60 * 60 * 24)
	// YearNanoSeconds - Number of Nanoseconds in a 365-day year

	/*
	For the Gregorian calendar the average length of the calendar year
	(the mean year) across the complete leap cycle of 400 years is 365.2425 days.
	The Gregorian Average Year is therefore equivalent to 365 days, 5 hours,
	49 minutes and 12 seconds.
	Sources:
	https://en.wikipedia.org/wiki/Year
	Source: https://en.wikipedia.org/wiki/Gregorian_calendar
*/
	YearNanoSeconds = int64(31556952000000000)

)

// DurationUtility - holds elements of
// time duration
type DurationUtility struct {
	TimeDuration time.Duration
	Years        int64
	Days         int64
	Hours        int64
	Minutes      int64
	Seconds      int64
	MilliSeconds int64
	MicroSeconds int64
	NanoSeconds  int64
	// NanosecStr - Example: 2-Days 13-Hours 26-Minutes 46-Seconds 864197832-Nanoseconds
	NanosecStr string
	// DurationStr - Example: 2-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
	DurationStr string
	// DefaultStr - Example: 61h26m46.864197832s - format provided by 'go' library
	DefaultStr string
}


func (du *DurationUtility) AddDurationToThis(duration time.Duration) {

	durPlus := du.TimeDuration + duration

	elapsedDuration := du.GetDurationBreakDown(durPlus)

	du.CopyToThis(elapsedDuration)
}

func (du *DurationUtility) AddToThis(duIn DurationUtility) {

	durPlus := du.TimeDuration + duIn.TimeDuration

	elapsedDuration := du.GetDurationBreakDown(durPlus)

	du.CopyToThis(elapsedDuration)

}

// Empty - This method initializes
// all of the fields in this
// DurationUtility structure to thier
// zero values.
func (du *DurationUtility) Empty() {
	du.TimeDuration = time.Duration(0)
	du.Years = 0
	du.Days = 0
	du.Hours = 0
	du.Minutes = 0
	du.Seconds = 0
	du.MilliSeconds = 0
	du.MicroSeconds = 0
	du.NanoSeconds = 0
	du.NanosecStr = ""
	du.DurationStr = ""
	du.DefaultStr = ""

}

func (du *DurationUtility) CopyToThis(duIn DurationUtility) {
	du.Empty()
	du.TimeDuration = duIn.TimeDuration
	du.Years = duIn.Years
	du.Days = duIn.Days
	du.Hours = duIn.Hours
	du.Minutes = duIn.Minutes
	du.Seconds = duIn.Seconds
	du.MilliSeconds = duIn.MilliSeconds
	du.MicroSeconds = duIn.MicroSeconds
	du.NanoSeconds = duIn.NanoSeconds
	du.NanosecStr = duIn.NanosecStr
	du.DurationStr = duIn.DurationStr
	du.DefaultStr = duIn.DefaultStr

	return
}

func (du *DurationUtility) Equal(duIn DurationUtility) bool {

	if du.TimeDuration != duIn.TimeDuration ||
		du.Years != duIn.Years ||
		du.Days != duIn.Days ||
		du.Hours != duIn.Hours ||
		du.Minutes != duIn.Minutes ||
		du.Seconds != duIn.Seconds ||
		du.MilliSeconds != duIn.MilliSeconds ||
		du.MicroSeconds != duIn.MicroSeconds ||
		du.NanoSeconds != duIn.NanoSeconds ||
		du.NanosecStr != duIn.NanosecStr ||
		du.DurationStr != duIn.DurationStr ||
		du.DefaultStr != duIn.DefaultStr {
		return false
	}

	return true

}


// GetDuration - Returns a time.Duration structure defining the duration between
// input parameters startTime and endTime
func (du DurationUtility) GetDuration(startTime time.Time, endTime time.Time) (time.Duration, error) {

	def := time.Duration(0)

	if startTime.Equal(endTime) {
		return def, nil
	}

	if endTime.Before(startTime) {
		return def, errors.New("DateTimeUtility.GetDuration() Error: endTime less than startTime")
	}

	return endTime.Sub(startTime), nil
}

func (du DurationUtility) GenerateDuration(duIn DurationUtility) ( time.Duration, error){
	return  du.GetDurationFromElapsedTime(duIn)
}

// GetDurationBreakDown - Receives a Duration type
// and returns a breakdown of duration by years,
// days, hours, minutes, seconds, milliseconds,
// microseconds and nanoseconds.
// NOTE: Years is arbitrarily set to the number of
// nanoseconds in a standard 365-day year. The Years
// calculation does NOT take Leap Years into account.
func (du DurationUtility) GetDurationBreakDown(d time.Duration) DurationUtility {
	str := ""
	durationUtility := DurationUtility{TimeDuration: d}
	firstEle := false
	rd := int64(d)
	durationUtility.DefaultStr = fmt.Sprintf("%v", d)
	if rd == 0 {
		durationUtility.DurationStr = "0-Nanoseconds"
		durationUtility.NanosecStr =  "0-Nanoseconds"
		return durationUtility
	}

	if rd >= YearNanoSeconds {
		durationUtility.Years = rd / YearNanoSeconds
		rd -= YearNanoSeconds * durationUtility.Years
	}

	if durationUtility.Years > 0 {
		str = fmt.Sprintf("%v-Years ", durationUtility.Years)
		firstEle = true
	}

	if rd >= DayNanoSeconds {
		durationUtility.Days = rd / DayNanoSeconds
		rd -= DayNanoSeconds * durationUtility.Days
	}

	if durationUtility.Days > 0 || firstEle {
		str += fmt.Sprintf("%v-Days ", durationUtility.Days)
		firstEle = true
	}

	if rd >= HourNanoSeconds {
		durationUtility.Hours = rd / HourNanoSeconds
		rd -= HourNanoSeconds * durationUtility.Hours
	}

	if durationUtility.Hours > 0 || firstEle {
		str += fmt.Sprintf("%v-Hours ", durationUtility.Hours)
		firstEle = true
	}

	if rd >= MinuteNanoSeconds {
		durationUtility.Minutes = rd / MinuteNanoSeconds
		rd -= MinuteNanoSeconds * durationUtility.Minutes
	}

	if durationUtility.Minutes > 0 || firstEle {
		str += fmt.Sprintf("%v-Minutes ", durationUtility.Minutes)
		firstEle = true
	}

	if rd >= SecondNanoseconds {
		durationUtility.Seconds = rd / SecondNanoseconds
		rd -= SecondNanoseconds * durationUtility.Seconds
	}

	if durationUtility.Seconds > 0 || firstEle {
		str += fmt.Sprintf("%v-Seconds ", durationUtility.Seconds)
		firstEle = true
	}

	durationUtility.NanosecStr = str + fmt.Sprintf("%v-Nanoseconds", rd)

	if rd >= MilliSecondNanoseconds {
		durationUtility.MilliSeconds = rd / MilliSecondNanoseconds
		rd -= MilliSecondNanoseconds * durationUtility.MilliSeconds
	}

	if durationUtility.MilliSeconds > 0 || firstEle {
		str += fmt.Sprintf("%v-Milliseconds ", durationUtility.MilliSeconds)
		firstEle = true
	}

	if rd >= MicroSecondNanoseconds {
		durationUtility.MicroSeconds = rd / MicroSecondNanoseconds
		rd -= MicroSecondNanoseconds * durationUtility.MicroSeconds
	}

	if durationUtility.MicroSeconds > 0 || firstEle {
		str += fmt.Sprintf("%v-Microseconds ", durationUtility.MicroSeconds)
		firstEle = true
	}

	durationUtility.NanoSeconds = rd

	if durationUtility.NanoSeconds > 0 || firstEle {
		str += fmt.Sprintf("%v-Nanoseconds", durationUtility.NanoSeconds)
		firstEle = true
	}

	durationUtility.DurationStr = strings.TrimRight(str, " ")

	return durationUtility

}

func (du DurationUtility) GetDurationFromElapsedTime(elapsedTime DurationUtility) (time.Duration, error) {
	var dns int64

	dns = elapsedTime.Years * YearNanoSeconds
	dns += elapsedTime.Days * DayNanoSeconds
	dns += elapsedTime.Hours * HourNanoSeconds
	dns += elapsedTime.Minutes * MinuteNanoSeconds
	dns += elapsedTime.Seconds * SecondNanoseconds
	dns += elapsedTime.MilliSeconds * MilliSecondNanoseconds
	dns += elapsedTime.MicroSeconds * MicroSecondNanoseconds
	dns += elapsedTime.NanoSeconds

	s := fmt.Sprintf("%vns", dns)

	dur, err := time.ParseDuration(s)

	if err != nil {
		e := errors.New("DurationUtility:GetDurationFromElapsedTime() Error Parsing Duration: " + err.Error())
		return dur, e
	}

	return dur, nil

}


// GetElapsedTime - calculates the elapsed time
// between input parameters startTime and endTime.
// The result is returned in an DurationUtility
// structure.
func (du DurationUtility) GetElapsedTime(startTime time.Time, endTime time.Time) (DurationUtility, error) {

	dur, err := du.GetDuration(startTime, endTime)

	if err != nil {
		s := "DateTimeUtility-GetElapsedTime Error: " + err.Error()

		return DurationUtility{}, errors.New(s)
	}

	return du.GetDurationBreakDown(dur), nil

}

// GetTimePlusDuration - Returns time plus input duration as a time.Time type.
func (du DurationUtility) GetTimePlusDuration(tStartTime time.Time, duration time.Duration) time.Time {

	return tStartTime.Add(duration)
}

// GetTimeMinusDuration - Returns time minus input duration as a time.Type type.
func (du DurationUtility) GetTimeMinusDuration(tStartTime time.Time, duration time.Duration) time.Time {

	return tStartTime.Add(-duration)
}

