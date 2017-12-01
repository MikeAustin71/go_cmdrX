package common

import (
	"errors"
	"fmt"
	"math"
	"time"
)

/*
	The principal component of this library is the DurationUtility. This
	type plus associated methods is used to manage and control time
	duration calculations.

	Usage requires two operations:
	1. You must first initialize the DurationUtility type using one of the
		 four 'Set' methods shown below:
		 	a. SetStartTimeDuration()
		 	b. SetStartEndTimes()
		 	c. SetStartTimePlusTime()
		 	d. SetSTartTimeMinusTime()

	2. After the DurationUtility is initialized in step one above, you are free
		 to call any of the following 'Get' methods in order to return
		 formatted time durations. A call to any of these methods will
		 return a DurationDto which contains a record of the duration
		 calculation broken down by years, months, weeks, days, hours,
		 minutes, seconds, milliseconds, microseconds and nanoseconds.
		 In addition, the DurationDto contains a field named, 'DisplayStr'
		 which contains the formatted text version of the duration output.

			a. GetYearsMthDays()
			b. GetYearsMthsWeeksTime()
			c. GetWeeksDaysTime()
			d. GetDaysTime()
			e. GetHoursTime()
			f. GetYrMthsWkDayHourSecNanosecDuration()
			g. GetNanosecondsDuration()
			h. GetDefaultDuration()
			i. GetGregorianYearDuration()


*/

const (
	// MicroSecondNanoseconds - Number of Nanoseconds in a Microsecond
	MicroSecondNanoseconds = int64(time.Microsecond)
	// MilliSecondNanoseconds - Number of Nanoseconds in a MilliSecond
	MilliSecondNanoseconds = int64(time.Millisecond)
	// SecondNanoseconds - Number of Nanoseconds in a Second
	SecondNanoseconds = int64(time.Second)
	// MinuteNanoSeconds - Number of Nanoseconds in a minute
	MinuteNanoSeconds = int64(time.Minute)
	// HourNanoSeconds - Number of Nanoseconds in an hour
	HourNanoSeconds = int64(time.Hour)
	// DayNanoSeconds - Number of Nanoseconds in a 24-hour day
	DayNanoSeconds = int64(24) * HourNanoSeconds

	WeekNanoSeconds = int64(7) * DayNanoSeconds

	/*
		For the Gregorian calendar the average length of the calendar year
		(the mean year) across the complete leap cycle of 400 Years is 365.2425 days.
		The Gregorian Average Year is therefore equivalent to 365 days, 5 hours,
		49 minutes and 12 seconds.
		Sources:
		https://en.wikipedia.org/wiki/Year
		Source: https://en.wikipedia.org/wiki/Gregorian_calendar
	*/

	GregorianYearNanoSeconds = int64(31556952000000000)
)

// TimeDto - used for transmitting
// time elements.
type TimeDto struct {
	Years        int64
	Months       int64
	Weeks        int64
	Days         int64
	Hours        int64
	Minutes      int64
	Seconds      int64
	Milliseconds int64
	Microseconds int64
	Nanoseconds  int64
}

func (tDto *TimeDto) ConvertToAbsoluteValues() {
	tDto.Years = int64(math.Abs(float64(tDto.Years)))
	tDto.Months = int64(math.Abs(float64(tDto.Months)))
	tDto.Weeks = int64(math.Abs(float64(tDto.Weeks)))
	tDto.Days = int64(math.Abs(float64(tDto.Days)))
	tDto.Hours = int64(math.Abs(float64(tDto.Hours)))
	tDto.Minutes = int64(math.Abs(float64(tDto.Minutes)))
	tDto.Seconds = int64(math.Abs(float64(tDto.Seconds)))
	tDto.Milliseconds = int64(math.Abs(float64(tDto.Milliseconds)))
	tDto.Microseconds = int64(math.Abs(float64(tDto.Microseconds)))
	tDto.Nanoseconds = int64(math.Abs(float64(tDto.Nanoseconds)))
}

func (tDto *TimeDto) ConvertToNegativeValues() {
	tDto.ConvertToAbsoluteValues()
	tDto.Years = tDto.Years * -1
	tDto.Months = tDto.Months * -1
	tDto.Weeks = tDto.Weeks * -1
	tDto.Days = tDto.Days * -1
	tDto.Hours = tDto.Hours * -1
	tDto.Minutes = tDto.Minutes * -1
	tDto.Seconds = tDto.Seconds * -1
	tDto.Milliseconds = tDto.Milliseconds * -1
	tDto.Microseconds = tDto.Microseconds * -1
	tDto.Nanoseconds = tDto.Nanoseconds * -1
}

type DurationDto struct {
	StartDateTime        time.Time
	EndDateTime          time.Time
	TimeDuration         time.Duration
	Years                int64
	YearsNanosecs        int64
	Months               int64
	MonthsNanosecs       int64
	Weeks                int64
	WeeksNanosecs        int64
	Days                 int64
	DaysNanosecs         int64
	Hours                int64
	HoursNanosecs        int64
	Minutes              int64
	MinutesNanosecs      int64
	Seconds              int64
	SecondsNanosecs      int64
	Milliseconds         int64
	MillisecondsNanosecs int64
	Microseconds         int64
	MicrosecondsNanosecs int64
	Nanoseconds          int64
	NanosecondsNanosecs  int64
	DisplayStr           string
}

func (dDto *DurationDto) CalcTotalNanoSecs() int64 {
	ns := dDto.YearsNanosecs
	ns += dDto.MonthsNanosecs
	ns += dDto.WeeksNanosecs
	ns += dDto.DaysNanosecs
	ns += dDto.HoursNanosecs
	ns += dDto.MinutesNanosecs
	ns += dDto.SecondsNanosecs
	ns += dDto.MillisecondsNanosecs
	ns += dDto.MicrosecondsNanosecs
	ns += dDto.NanosecondsNanosecs

	return ns
}

func (dDto *DurationDto) InitializeTime(tDto TimeDto) {
	dDto.EmptyTimeValues()
	dDto.Years = tDto.Years
	dDto.Months = tDto.Months
	dDto.Weeks = tDto.Weeks
	dDto.Days = tDto.Days
	dDto.Hours = tDto.Hours
	dDto.Minutes = tDto.Minutes
	dDto.Seconds = tDto.Seconds
	dDto.Milliseconds = tDto.Milliseconds
	dDto.Microseconds = tDto.Microseconds
	dDto.Nanoseconds = tDto.Nanoseconds
}

func (dDto *DurationDto) Copy() DurationDto {
	d := DurationDto{}

	d.StartDateTime = dDto.StartDateTime
	d.EndDateTime = dDto.EndDateTime
	d.TimeDuration = dDto.TimeDuration
	d.Years = dDto.Years
	d.YearsNanosecs = dDto.YearsNanosecs
	d.Months = dDto.Months
	d.MonthsNanosecs = dDto.MonthsNanosecs
	d.Weeks = dDto.Weeks
	d.WeeksNanosecs = dDto.WeeksNanosecs
	d.Days = dDto.Days
	d.DaysNanosecs = dDto.DaysNanosecs
	d.Hours = dDto.Hours
	d.HoursNanosecs = dDto.HoursNanosecs
	d.Minutes = dDto.Minutes
	d.MinutesNanosecs = dDto.MinutesNanosecs
	d.Seconds = dDto.Seconds
	d.SecondsNanosecs = dDto.SecondsNanosecs
	d.Milliseconds = dDto.Milliseconds
	d.MillisecondsNanosecs = dDto.MillisecondsNanosecs
	d.Microseconds = dDto.Microseconds
	d.MicrosecondsNanosecs = dDto.MicrosecondsNanosecs
	d.Nanoseconds = dDto.Nanoseconds
	d.NanosecondsNanosecs = dDto.MicrosecondsNanosecs
	d.DisplayStr = dDto.DisplayStr

	return d
}

func (dDto *DurationDto) Empty() {
	dDto.StartDateTime = time.Time{}
	dDto.EndDateTime = time.Time{}
	dDto.TimeDuration = time.Duration(0)
	dDto.Years = 0
	dDto.YearsNanosecs = 0
	dDto.Months = 0
	dDto.MonthsNanosecs = 0
	dDto.Weeks = 0
	dDto.WeeksNanosecs = 0
	dDto.Days = 0
	dDto.DaysNanosecs = 0
	dDto.Hours = 0
	dDto.HoursNanosecs = 0
	dDto.Minutes = 0
	dDto.MinutesNanosecs = 0
	dDto.Seconds = 0
	dDto.SecondsNanosecs = 0
	dDto.Milliseconds = 0
	dDto.MillisecondsNanosecs = 0
	dDto.Microseconds = 0
	dDto.MicrosecondsNanosecs = 0
	dDto.Nanoseconds = 0
	dDto.NanosecondsNanosecs = 0
	dDto.DisplayStr = ""
}

func (dDto *DurationDto) EmptyTimeValues() {

	dDto.Years = 0
	dDto.YearsNanosecs = 0
	dDto.Months = 0
	dDto.MonthsNanosecs = 0
	dDto.Weeks = 0
	dDto.WeeksNanosecs = 0
	dDto.Days = 0
	dDto.DaysNanosecs = 0
	dDto.Hours = 0
	dDto.HoursNanosecs = 0
	dDto.Minutes = 0
	dDto.MinutesNanosecs = 0
	dDto.Seconds = 0
	dDto.SecondsNanosecs = 0
	dDto.Milliseconds = 0
	dDto.MillisecondsNanosecs = 0
	dDto.Microseconds = 0
	dDto.MicrosecondsNanosecs = 0
	dDto.Nanoseconds = 0
	dDto.NanosecondsNanosecs = 0
	dDto.DisplayStr = ""
}

func (dDto *DurationDto) EmptyNanosecs() {
	dDto.YearsNanosecs = 0
	dDto.MonthsNanosecs = 0
	dDto.WeeksNanosecs = 0
	dDto.DaysNanosecs = 0
	dDto.HoursNanosecs = 0
	dDto.MinutesNanosecs = 0
	dDto.SecondsNanosecs = 0
	dDto.MillisecondsNanosecs = 0
	dDto.MicrosecondsNanosecs = 0
	dDto.NanosecondsNanosecs = 0
	dDto.DisplayStr = ""
}

func (dDto *DurationDto) Equal(dto2 DurationDto) bool {

	if dDto.StartDateTime != dto2.StartDateTime ||
		dDto.EndDateTime != dto2.EndDateTime ||
		dDto.TimeDuration != dto2.TimeDuration ||
		dDto.Years != dto2.Years ||
		dDto.YearsNanosecs != dto2.YearsNanosecs ||
		dDto.Months != dto2.Months ||
		dDto.MonthsNanosecs != dto2.MonthsNanosecs ||
		dDto.Weeks != dto2.Weeks ||
		dDto.WeeksNanosecs != dto2.WeeksNanosecs ||
		dDto.Days != dto2.Days ||
		dDto.DaysNanosecs != dto2.DaysNanosecs ||
		dDto.Hours != dto2.Hours ||
		dDto.HoursNanosecs != dto2.HoursNanosecs ||
		dDto.Minutes != dto2.Minutes ||
		dDto.MinutesNanosecs != dto2.MinutesNanosecs ||
		dDto.Seconds != dto2.Seconds ||
		dDto.SecondsNanosecs != dto2.SecondsNanosecs ||
		dDto.Milliseconds != dto2.Milliseconds ||
		dDto.MillisecondsNanosecs != dto2.MillisecondsNanosecs ||
		dDto.Microseconds != dto2.Microseconds ||
		dDto.MicrosecondsNanosecs != dto2.MicrosecondsNanosecs ||
		dDto.Nanoseconds != dto2.Nanoseconds ||
		dDto.NanosecondsNanosecs != dto2.NanosecondsNanosecs ||
		dDto.DisplayStr != dto2.DisplayStr {

		return false
	}

	return true
}

// DurationUtility - holds elements of
// time duration
type DurationUtility struct {
	StartDateTime time.Time
	EndDateTime   time.Time
	TimeDuration  time.Duration
}

// CopyToThis - Receives and incoming DurationUtility data
// structure and copies the values to the current DurationUtility
// data structure.
func (du *DurationUtility) CopyToThis(duIn DurationUtility) {
	du.Empty()
	du.TimeDuration = duIn.TimeDuration
	du.StartDateTime = duIn.StartDateTime
	du.EndDateTime = duIn.EndDateTime

	return
}

// Copy - Returns a deep copy of the current
// DurationUtility data fields.
func (du *DurationUtility) Copy() DurationUtility {
	duOut := DurationUtility{}
	duOut.TimeDuration = du.TimeDuration
	duOut.StartDateTime = du.StartDateTime
	duOut.EndDateTime = du.EndDateTime

	return duOut
}

// Equal - This method may be used to determine if two
// DurationUtility data structures are equivalent.
func (du *DurationUtility) Equal(duIn DurationUtility) bool {

	if du.TimeDuration != duIn.TimeDuration ||
		du.StartDateTime != duIn.StartDateTime ||
		du.EndDateTime != duIn.EndDateTime {

		return false
	}

	return true

}

// Empty - This method initializes
// all of the fields in this
// DurationUtility structure to their
// zero values.
func (du *DurationUtility) Empty() {
	du.TimeDuration = time.Duration(0)
	du.StartDateTime = time.Time{}
	du.EndDateTime = time.Time{}
}

// GetYearMthDaysTimeAbbrv - Abbreviated formatting of Years, Months,
// Days, Hours, Minutes, Seconds, Milliseconds, Microseconds and
// Nanoseconds. At a minimum only Hours, Minutes, Seconds, Milliseconds,
// Microseconds and Nanoseconds.
// Abbreviated Years Mths Days Time Duration - Example Return:
// 0-Hours 0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (du *DurationUtility) GetYearMthDaysTimeAbbrv() (DurationDto, error) {

	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	du.calcYearsFromDuration(&dDto)

	du.calcMonthsFromDuration(&dDto)

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	yearsElement := ""

	monthsElement := ""

	daysElement := ""

	if dDto.Years > 0 {
		yearsElement = fmt.Sprintf("%v-Years ", dDto.Years)
	}

	if dDto.Months > 0 {
		monthsElement = fmt.Sprintf("%v-Months ", dDto.Months)
	}

	if dDto.Days > 0 {
		daysElement = fmt.Sprintf("%v-Days ", dDto.Days)
	}

	str := fmt.Sprintf("%v-Hours ", dDto.Hours)

	str += fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	str += fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = yearsElement + monthsElement + daysElement + str

	return dDto, nil

}

// GetYearMthDaysTime - Calculates Duration and breakdowns
// time elements by Years, Months, days, hours, minutes,
// seconds, milliseconds, microseconds and nanoseconds.
// Example DisplayStr:
// Years Mths Days Time Duration - Example Return:
// 12-Years 3-Months 2-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (du *DurationUtility) GetYearMthDaysTime() (DurationDto, error) {

	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	du.calcYearsFromDuration(&dDto)

	du.calcMonthsFromDuration(&dDto)

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	str := fmt.Sprintf("%v-Years ", dDto.Years)

	str += fmt.Sprintf("%v-Months ", dDto.Months)

	str += fmt.Sprintf("%v-Days ", dDto.Days)

	str += fmt.Sprintf("%v-Hours ", dDto.Hours)

	str += fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	str += fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = str

	return dDto, nil
}

// GetYearsMthsWeeksTimeAbbrv - Abbreviated formatting of Years, Months,
// Weeks, Days, Hours, Minutes, Seconds, Milliseconds, Microseconds,
// Nanoseconds. At a minimum only Hours, Minutes, Seconds, Milliseconds,
// Microseconds, Nanoseconds are displayed. Example return when Years,
// Months, Weeks and Days are zero:
// 0-Hours 0-Minutes 0-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (du *DurationUtility) GetYearsMthsWeeksTimeAbbrv() (DurationDto, error) {
	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	du.calcYearsFromDuration(&dDto)

	du.calcMonthsFromDuration(&dDto)

	du.calcWeeksFromDuration(&dDto)

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	yearsElement := ""

	monthsElement := ""

	weeksElement := ""

	daysElement := ""

	if dDto.Years > 0 {
		yearsElement = fmt.Sprintf("%v-Years ", dDto.Years)
	}

	if dDto.Months > 0 {
		monthsElement = fmt.Sprintf("%v-Months ", dDto.Months)
	}

	if dDto.Weeks > 0 {
		weeksElement = fmt.Sprintf("%v-Weeks ", dDto.Weeks)
	}

	if dDto.Days > 0 {
		daysElement = fmt.Sprintf("%v-Days ", dDto.Days)
	}

	hoursElement := fmt.Sprintf("%v-Hours ", dDto.Hours)

	minutesElement := fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	secondsElement := fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	millisecondsElement := fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	microsecondsElement := fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	nanosecondsElement := fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = yearsElement + monthsElement +
		weeksElement + daysElement +
		hoursElement + minutesElement + secondsElement +
		millisecondsElement + microsecondsElement +
		nanosecondsElement

	return dDto, nil

}

// GetYearsMthsWeeksTime - Example Return:
// 12-Years 3-Months 2-Weeks 1-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (du *DurationUtility) GetYearsMthsWeeksTime() (DurationDto, error) {

	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	du.calcYearsFromDuration(&dDto)

	du.calcMonthsFromDuration(&dDto)

	du.calcWeeksFromDuration(&dDto)

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	yearsElement := fmt.Sprintf("%v-Years ", dDto.Years)

	monthsElement := fmt.Sprintf("%v-Months ", dDto.Months)

	weeksElement := fmt.Sprintf("%v-Weeks ", dDto.Weeks)

	daysElement := fmt.Sprintf("%v-Days ", dDto.Days)

	hoursElement := fmt.Sprintf("%v-Hours ", dDto.Hours)

	minutesElement := fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	secondsElement := fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	millisecondsElement := fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	microsecondsElement := fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	nanosecondsElement := fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = yearsElement + monthsElement +
		weeksElement + daysElement +
		hoursElement + minutesElement + secondsElement +
		millisecondsElement + microsecondsElement +
		nanosecondsElement

	return dDto, nil

}

// GetWeeksDaysTime - Example DisplayStr
// 126-Weeks 1-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (du *DurationUtility) GetWeeksDaysTime() (DurationDto, error) {
	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	du.calcWeeksFromDuration(&dDto)

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	str := ""

	str += fmt.Sprintf("%v-Weeks ", dDto.Weeks)

	str += fmt.Sprintf("%v-Days ", dDto.Days)

	str += fmt.Sprintf("%v-Hours ", dDto.Hours)

	str += fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	str += fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = str

	return dDto, nil

}

// GetDaysTime - Returns duration formatted as
// days, hours, minutes, seconds, milliseconds, microseconds,
// and nanoseconds.
// Example: 97-Days 13-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (du *DurationUtility) GetDaysTime() (DurationDto, error) {
	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	str := ""

	str += fmt.Sprintf("%v-Days ", dDto.Days)

	str += fmt.Sprintf("%v-Hours ", dDto.Hours)

	str += fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	str += fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = str

	return dDto, nil
}

// GetHoursTime - Returns duration formatted as hours,
// minutes, seconds, milliseconds, microseconds, nanoseconds.
// Example: 152-Hours 26-Minutes 46-Seconds 864-Milliseconds 197-Microseconds 832-Nanoseconds
func (du *DurationUtility) GetHoursTime() (DurationDto, error) {
	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	str := ""

	str += fmt.Sprintf("%v-Hours ", dDto.Hours)

	str += fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	str += fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = str

	return dDto, nil

}

// GetYrMthWkDayHrMinSecNanosecs - Returns duration formatted
// as Year, Month, Day, Hour, Second and Nanoseconds.
// Example: 3-Years 2-Months 3-Weeks 2-Days 13-Hours 26-Minutes 46-Seconds 864197832-Nanoseconds
func (du *DurationUtility) GetYrMthWkDayHrMinSecNanosecs() (DurationDto, error) {
	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	du.calcYearsFromDuration(&dDto)

	du.calcMonthsFromDuration(&dDto)

	du.calcWeeksFromDuration(&dDto)

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcNanoseconds(&dDto)

	str := ""

	str += fmt.Sprintf("%v-Years ", dDto.Years)

	str += fmt.Sprintf("%v-Months ", dDto.Months)

	str += fmt.Sprintf("%v-Weeks ", dDto.Weeks)

	str += fmt.Sprintf("%v-Days ", dDto.Days)

	str += fmt.Sprintf("%v-Hours ", dDto.Hours)

	str += fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	str += fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	str += fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = str

	return dDto, nil
}

// GetNanosecondsDuration - Returns duration formatted as
// Nonseconds. DisplayStr shows Nanoseconds expressed as a
// 64-bit integer value.
func (du *DurationUtility) GetNanosecondsDuration() (DurationDto, error) {
	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	du.calcNanoseconds(&dDto)

	str := fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = str

	return dDto, nil

}

// GetDefaultDuration returns duration formatted
// as nanoseconds. The DisplayStr shows the default
// string value for duration.
// Example: 61h26m46.864197832s
func (du *DurationUtility) GetDefaultDuration() (DurationDto, error) {

	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	dDto.Nanoseconds = rd

	dur := time.Duration(rd)

	dDto.DisplayStr = fmt.Sprintf("%v", dur)

	return dDto, nil
}

// GetGregorianYearDuration - Returns a string showing the
// breakdown of duration by Gregorian Years, Days, Hours, Minutes,
// Seconds, Milliseconds, Microseconds and Nanoseconds. Unlike
// other calculations which use a Standard 365-day year consisting
// of 24-hour days, a Gregorian Year consists of 365 days, 5-hours,
// 59-minutes and 12 Seconds. For the Gregorian calendar the
// average length of the calendar year (the mean year) across
// the complete leap cycle of 400 Years is 365.2425 days.
// Sources:
// https://en.wikipedia.org/wiki/Year
// Source: https://en.wikipedia.org/wiki/Gregorian_calendar

func (du *DurationUtility) GetGregorianYearDuration() (DurationDto, error) {

	rd := int64(du.TimeDuration)

	dDto := DurationDto{}

	if rd == 0 {
		dDto.DisplayStr = "0-Nanoseconds"
		return dDto, nil
	}

	du.calcBaseData(&dDto)

	if rd > GregorianYearNanoSeconds {
		dDto.Years = rd / GregorianYearNanoSeconds
		dDto.YearsNanosecs = dDto.Years * GregorianYearNanoSeconds
	}

	du.calcDaysFromDuration(&dDto)

	du.calcHoursMinSecs(&dDto)

	du.calcMilliSeconds(&dDto)

	du.calcMicroSeconds(&dDto)

	du.calcNanoseconds(&dDto)

	str := fmt.Sprintf("%v-Gregorian Years ", dDto.Years)

	str += fmt.Sprintf("%v-Days ", dDto.Days)

	str += fmt.Sprintf("%v-Hours ", dDto.Hours)

	str += fmt.Sprintf("%v-Minutes ", dDto.Minutes)

	str += fmt.Sprintf("%v-Seconds ", dDto.Seconds)

	str += fmt.Sprintf("%v-Milliseconds ", dDto.Milliseconds)

	str += fmt.Sprintf("%v-Microseconds ", dDto.Microseconds)

	str += fmt.Sprintf("%v-Nanoseconds", dDto.Nanoseconds)

	dDto.DisplayStr = str

	return dDto, nil

}

// GetDurationFromStartEndTimes - Computes the duration
// by subtracting Starting Date Time from the Ending Date
// time. No changes are made to or stored in the
// existing DurationUtility data structures.
func (du *DurationUtility) GetDurationFromStartEndTimes(startDateTime time.Time, endDateTime time.Time) (time.Duration, error) {

	if startDateTime.IsZero() {
		return time.Duration(0), errors.New("DurationUtility.GetDurationFromStartEndTimes() ERROR - startDateTime is ZERO!")
	}

	if endDateTime.IsZero() {
		return time.Duration(0), errors.New("DurationUtility.GetDurationFromStartEndTimes() ERROR - endDateTime is ZERO!")
	}

	if endDateTime == startDateTime {

		return time.Duration(0), nil

	} else if endDateTime.Before(startDateTime) {

		return startDateTime.Sub(endDateTime), nil

	} else {

		return endDateTime.Sub(startDateTime), nil

	}

}

// GetDurationFromSeconds - returns a time Duration value
// based on the number of seconds passed to this method.
// No changes are made to or stored in the existing
// DurationUtility data structures.
func (du *DurationUtility) GetDurationFromSeconds(seconds int64) time.Duration {

	return time.Duration(seconds) * time.Second

}

// GetDurationFromMinutes - returns a time Duration value
// based on the number of minutes passed to this method.
// No changes are made to or stored in the existing
// DurationUtility data structures.
func (du DurationUtility) GetDurationFromMinutes(minutes int64) time.Duration {

	return time.Duration(minutes) * time.Minute

}

// SetStartTimeDuration - Receives a starting date time and
// a time duration. The method then calculates the ending
// date time, duration and populates the DurationUtility
// data fields.
//
// The Method will except negative time durations. This means
// that the duration will be subtracted from the starting
// date time.
func (du *DurationUtility) SetStartTimeDuration(startDateTime time.Time, duration time.Duration) error {

	if startDateTime.IsZero() {
		return fmt.Errorf("DurationUtility.SetStartTimeDuration() Error - Start Time is Zero!")
	}

	x := int64(duration)

	du.Empty()

	if x < 0 {
		eTime := startDateTime
		du.StartDateTime = eTime.Add(duration)
		du.EndDateTime = eTime
		du.TimeDuration = time.Duration(x * -1)
	} else if x == 0 {

		return fmt.Errorf("DurationUtility.SetStartTimeDuration() Error - Duration is Zero!")

	} else {
		du.StartDateTime = startDateTime
		du.EndDateTime = startDateTime.Add(duration)
		du.TimeDuration = duration
	}

	err := du.ValidateDurationBaseData()

	if err != nil {
		return fmt.Errorf("DurationUtility.SetStartEndTimes() ERROR - %v", err.Error())
	}

	return nil
}

// SetStartEndTimes - Calculate duration values and save the results in the DurationUtility
// data fields. Calculations are based on a starting date time and an ending date time passed
// to the method.
func (du *DurationUtility) SetStartEndTimes(startDateTime time.Time, endDateTime time.Time) error {

	if startDateTime.IsZero() &&
		endDateTime.IsZero() {
		return errors.New("Both startDateTime and endDateTime have ZERO Values!")
	}

	du.Empty()

	sTime := startDateTime
	eTime := endDateTime

	if eTime.Before(sTime) {
		s2 := sTime
		sTime = eTime
		eTime = s2
	}

	du.StartDateTime = sTime

	du.EndDateTime = eTime

	du.TimeDuration = du.EndDateTime.Sub(du.StartDateTime)

	err := du.ValidateDurationBaseData()

	if err != nil {
		return fmt.Errorf("DurationUtility.SetStartEndTimes() ERROR - %v", err.Error())
	}

	return nil
}

// SetStartTimePlusTime - Calculate duration values based on a Starting Date Time and
// time values (Years, Months, weeks, days, hours, minutes etc.) passed to the method
// in the 'times' parameter. The 'timeDto' parameter is added to
// 'StartDateTime'.
//
// Values in the 'timeDto' parameter are automatically converted to positive
// numeric values before being added to 'StartDateTime'.
//
// True values for StartDateTime, EndDateTime and TimeDuration are
// then stored in the DurationUtility data structure.
func (du *DurationUtility) SetStartTimePlusTime(startDateTime time.Time, timeDto TimeDto) error {
	du.Empty()
	timeDto.ConvertToAbsoluteValues()
	dur := DurationDto{StartDateTime: startDateTime}
	dur.InitializeTime(timeDto)
	du.calcFromYears(&dur)
	du.calcFromMonths(&dur)
	du.calcFromWeeks(&dur)
	du.calcFromDays(&dur)
	du.calcFromHoursMinSecs(&dur)
	du.calcFromMilliSecs(&dur)
	du.calcFromMicroSecs(&dur)
	du.calcFromNanoSecs(&dur)
	du.StartDateTime = startDateTime
	du.TimeDuration = time.Duration(dur.CalcTotalNanoSecs())
	du.EndDateTime = du.StartDateTime.Add(du.TimeDuration)

	err := du.ValidateDurationBaseData()

	if err != nil {
		return fmt.Errorf("DurationUtility.SetStartEndTimes() ERROR - %v", err.Error())
	}

	return nil
}

// SetStartTimeMinusTime - Calculate duration values based on a Starting Date Time and
// time values (Years, Months, weeks, days, hours, minutes etc.) passed to the method
// in the 'timeDto' parameter. The time values in the 'timeDto' parameter are subtracted
// from 'StartDateTime'.
//
// Time values in the 'timeDto' parameter are first converted to negative
// numeric values. Then these values are added to the 'startDateTime' value
// which is effective treated as an End Date Time.
//
// As a result. true values for StartDateTime, EndDateTime and
// TimeDuration are stored in the DurationUtility data structure.

func (du *DurationUtility) SetStartTimeMinusTime(startDateTime time.Time, timeDto TimeDto) error {
	du.Empty()
	timeDto.ConvertToNegativeValues()
	dur := DurationDto{StartDateTime: startDateTime}
	dur.InitializeTime(timeDto)
	du.calcFromYears(&dur)
	du.calcFromMonths(&dur)
	du.calcFromWeeks(&dur)
	du.calcFromDays(&dur)
	du.calcFromHoursMinSecs(&dur)
	du.calcFromMilliSecs(&dur)
	du.calcFromMicroSecs(&dur)
	du.calcFromNanoSecs(&dur)

	ns := dur.CalcTotalNanoSecs()
	tDur := time.Duration(ns)

	sTime := startDateTime.Add(tDur)

	du.StartDateTime = sTime
	du.TimeDuration = time.Duration(ns * -1)
	du.EndDateTime = startDateTime

	err := du.ValidateDurationBaseData()

	if err != nil {
		return fmt.Errorf("DurationUtility.SetStartEndTimes() ERROR - %v", err.Error())
	}

	return nil
}

// ValidateDurationBaseData - Validates DurationUtility.TimeDuration
// DurationUtility.StartDateTime and DurationUtility.EndDateTime.
// Note: if DurationUtility.StartDateTime and DurationUtility.EndDateTime
// have zero values, DurationUtility.StartDateTime will be defaulted to
// time.Now().UTC()
func (du *DurationUtility) ValidateDurationBaseData() error {

	rd := int64(du.TimeDuration)

	if du.StartDateTime.IsZero() && du.EndDateTime.IsZero() &&
		rd == 0 {

		return fmt.Errorf("DurationUtility.ValidateDurationBaseData() - ERROR: Time Duration, Start Time and End Time are ZERO!")
	}

	if du.StartDateTime.IsZero() && du.EndDateTime.IsZero() {
		du.StartDateTime = time.Now().UTC()
	}

	if rd == 0 && du.StartDateTime == du.EndDateTime {
		return fmt.Errorf("DurationUtility.ValidateDurationBaseData() - ERROR: Time Duration is Zero, Start Time and End Time are EQUAL!")
	}

	if rd < 0 && du.StartDateTime.IsZero() && !du.EndDateTime.IsZero() {
		du.StartDateTime = du.EndDateTime.Add(du.TimeDuration)
		rd = rd * -1
		du.TimeDuration = time.Duration(rd)
	}

	if rd > 0 && !du.StartDateTime.IsZero() {
		du.EndDateTime = du.StartDateTime.Add(du.TimeDuration)
	}

	if du.EndDateTime.IsZero() && !du.StartDateTime.IsZero() &&
		rd < 0 {

		rd = rd * -1
		du.TimeDuration = time.Duration(rd)
		du.EndDateTime = du.StartDateTime.Add(du.TimeDuration)
	}

	return nil
}

// calcBaseData - Validates Time Duration
func (du *DurationUtility) calcBaseData(dDto *DurationDto) error {
	dDto.TimeDuration = du.TimeDuration
	dDto.StartDateTime = du.StartDateTime
	dDto.EndDateTime = du.EndDateTime
	return nil
}

func (du *DurationUtility) calcYearsFromDuration(dDto *DurationDto) {

	yearDateTime := dDto.StartDateTime
	rd := int64(dDto.TimeDuration)

	i := 0
	var prevDateTime time.Time

	for yearDateTime.Before(dDto.EndDateTime) {
		prevDateTime = yearDateTime
		i++
		yearDateTime = dDto.StartDateTime.AddDate(i, 0, 0)
	}

	i -= 1

	if i > 0 {
		yearDateTime = prevDateTime
		dDto.Years = int64(i)
		dDto.YearsNanosecs = int64(yearDateTime.Sub(dDto.StartDateTime))
		rd -= dDto.YearsNanosecs
	}

	return
}

func (du *DurationUtility) calcMonthsFromDuration(dDto *DurationDto) {

	rd := int64(dDto.TimeDuration)

	rd -= dDto.YearsNanosecs

	i := 0
	yearDateTime := dDto.StartDateTime.Add(time.Duration(dDto.YearsNanosecs))
	mthDateTime := yearDateTime
	prevDateTime := mthDateTime

	for mthDateTime.Before(dDto.EndDateTime) {
		prevDateTime = mthDateTime
		i++
		mthDateTime = yearDateTime.AddDate(0, i, 0)
	}

	i -= 1

	if i > 0 {
		mthDateTime = prevDateTime
		dDto.Months = int64(i)
		dDto.MonthsNanosecs = int64(mthDateTime.Sub(yearDateTime))

	}

	return

}

func (du *DurationUtility) calcWeeksFromDuration(dDto *DurationDto) {
	rd := int64(dDto.TimeDuration)

	rd -= dDto.YearsNanosecs + dDto.MonthsNanosecs

	if rd >= WeekNanoSeconds {
		dDto.Weeks = rd / WeekNanoSeconds
		dDto.WeeksNanosecs = dDto.Weeks * WeekNanoSeconds
	}

	return
}

func (du *DurationUtility) calcDaysFromDuration(dDto *DurationDto) {

	rd := int64(dDto.TimeDuration)

	rd -= dDto.YearsNanosecs + dDto.MonthsNanosecs +
		dDto.WeeksNanosecs

	if rd >= DayNanoSeconds {
		dDto.Days = rd / DayNanoSeconds
		dDto.DaysNanosecs = DayNanoSeconds * dDto.Days
	}
}

func (du *DurationUtility) calcHoursMinSecs(dDto *DurationDto) {

	rd := int64(dDto.TimeDuration)

	rd -= dDto.YearsNanosecs + dDto.MonthsNanosecs +
		dDto.WeeksNanosecs + dDto.DaysNanosecs

	if rd >= HourNanoSeconds {
		dDto.Hours = rd / HourNanoSeconds
		dDto.HoursNanosecs = HourNanoSeconds * dDto.Hours
		rd -= dDto.HoursNanosecs
	}

	if rd >= MinuteNanoSeconds {
		dDto.Minutes = rd / MinuteNanoSeconds
		dDto.MinutesNanosecs = MinuteNanoSeconds * dDto.Minutes
		rd -= dDto.MinutesNanosecs
	}

	if rd >= SecondNanoseconds {
		dDto.Seconds = rd / SecondNanoseconds
		dDto.SecondsNanosecs = SecondNanoseconds * dDto.Seconds
		rd -= dDto.SecondsNanosecs
	}

}

func (du *DurationUtility) calcMilliSeconds(dDto *DurationDto) {
	rd := int64(dDto.TimeDuration)

	rd -= dDto.YearsNanosecs + dDto.MonthsNanosecs +
		dDto.WeeksNanosecs + dDto.DaysNanosecs + dDto.HoursNanosecs +
		dDto.MinutesNanosecs + dDto.SecondsNanosecs

	if rd >= MilliSecondNanoseconds {
		dDto.Milliseconds = rd / MilliSecondNanoseconds
		dDto.MillisecondsNanosecs = MilliSecondNanoseconds * dDto.Milliseconds
	}

}

func (du *DurationUtility) calcMicroSeconds(dDto *DurationDto) {
	rd := int64(dDto.TimeDuration)

	rd -= dDto.YearsNanosecs + dDto.MonthsNanosecs +
		dDto.WeeksNanosecs + dDto.DaysNanosecs + dDto.HoursNanosecs +
		dDto.MinutesNanosecs + dDto.SecondsNanosecs +
		dDto.MillisecondsNanosecs

	if rd >= MicroSecondNanoseconds {
		dDto.Microseconds = rd / MicroSecondNanoseconds
		dDto.MicrosecondsNanosecs = MicroSecondNanoseconds * dDto.Microseconds
	}

}

func (du *DurationUtility) calcNanoseconds(dDto *DurationDto) {

	rd := int64(dDto.TimeDuration)

	rd -= dDto.YearsNanosecs + dDto.MonthsNanosecs +
		dDto.WeeksNanosecs + dDto.DaysNanosecs + dDto.HoursNanosecs +
		dDto.MinutesNanosecs + dDto.SecondsNanosecs +
		dDto.MillisecondsNanosecs + dDto.MicrosecondsNanosecs

	dDto.Nanoseconds = rd

}

func (du *DurationUtility) calcFromYears(dDto *DurationDto) {

	yearDateTime := dDto.StartDateTime.AddDate(int(dDto.Years), 0, 0)

	dDto.YearsNanosecs = int64(yearDateTime.Sub(dDto.StartDateTime))

	return

}

func (du *DurationUtility) calcFromMonths(dDto *DurationDto) {
	mthStartDateTime := dDto.StartDateTime.Add(time.Duration(dDto.YearsNanosecs))
	mthEndDateTime := mthStartDateTime.AddDate(0, int(dDto.Months), 0)

	dDto.MonthsNanosecs = int64(mthEndDateTime.Sub(mthStartDateTime))

	return
}

func (du *DurationUtility) calcFromWeeks(dDto *DurationDto) {

	dDto.WeeksNanosecs = dDto.Weeks * WeekNanoSeconds
}

func (du *DurationUtility) calcFromDays(dDto *DurationDto) {
	dDto.DaysNanosecs = dDto.Days * DayNanoSeconds
}

func (du *DurationUtility) calcFromHoursMinSecs(dDto *DurationDto) {
	dDto.HoursNanosecs = dDto.Hours * HourNanoSeconds
	dDto.MinutesNanosecs = dDto.Minutes * MinuteNanoSeconds
	dDto.SecondsNanosecs = dDto.Seconds * SecondNanoseconds
}

func (du *DurationUtility) calcFromMilliSecs(dDto *DurationDto) {
	dDto.MillisecondsNanosecs = dDto.Milliseconds * MilliSecondNanoseconds
}

func (du *DurationUtility) calcFromMicroSecs(dDto *DurationDto) {
	dDto.MicrosecondsNanosecs = dDto.Microseconds * MicroSecondNanoseconds
}

func (du *DurationUtility) calcFromNanoSecs(dDto *DurationDto) {
	dDto.NanosecondsNanosecs = dDto.Nanoseconds
}
