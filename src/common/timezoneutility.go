package common

import (
	"errors"
	"fmt"
	"time"
)

// NOTE: See https://golang.org/pkg/time/#LoadLocation
// and https://www.iana.org/time-zones to ensure that
// the IANA Time Zone Database is properly configured
// on your system.
const (
	// TzUsEast - USA Eastern Time Zone
	// IANA database identifier
	TzUsEast = "America/New_York"

	// TzUsCentral - USA Central Time Zone
	// IANA database identifier
	TzUsCentral = "America/Chicago"

	// TzUsMountain - USA Mountain Time Zone
	// IANA database identifier
	TzUsMountain = "America/Denver"

	// TzUsPacific - USA Pacific Time Zone
	// IANA database identifier
	TzUsPacific = "America/Los_Angeles"

	// TzUsHawaii - USA Hawaiian Time Zone
	// IANA database identifier
	TzUsHawaii = "Pacific/Honolulu"

	// tzUTC - UTC Time Zone IANA database
	// identifier
	TzUTC = "Zulu"

	NeutralDateFmt = "2006-01-02 15:04:05.000000000"
)

// TimeZoneUtility - Time Zone Data and Methods
type TimeZoneUtility struct {
	Description string
	TimeIn      time.Time
	TimeInLoc   *time.Location
	TimeOut     time.Time
	TimeOutLoc  *time.Location
	TimeUTC     time.Time
}

// ConvertTz - Convert Time from existing time zone to targetTZone.
// The results are stored in the TimeZoneUtility data structure.
// The input time and output time are equivalent times adjusted
// for different time zones.
//
// Input Parameters:
// tIn - time.Time initial time
// targetTz - The IANA Time Zone or Time Zone 'Local' to which
// input parameter 'tIn' will be converted.

// Output Values are returned in the tzu (TimeZoneUtility)
// data fields.
func (tzu *TimeZoneUtility) ConvertTz(tIn time.Time, targetTz string) error {

	isValidTz, _, _ := tzu.IsValidTimeZone(targetTz)

	if !isValidTz {
		return errors.New(fmt.Sprintf("TimeZoneUtility:ConvertTz() Error: targetTz is INVALID!! Input Time Zone == %v", targetTz))
	}

	if tIn.IsZero() {
		return errors.New("TimeZoneUtility:ConvertTz() Error: Input parameter time, 'tIn' is zero and INVALID")
	}

	tzOut, err := time.LoadLocation(targetTz)

	if err != nil {
		return fmt.Errorf("TimeZoneUtility:ConvertTz() - Error Loading Target IANA Time Zone 'targetTz', %v. Errors: %v ", targetTz, err.Error())
	}

	tzu.Empty()

	tzu.SetTimeIn(tIn)

	tzu.SetTimeOut(tIn.In(tzOut))

	tzu.SetUTCTime(tIn)

	return nil
}

// Empty - Clears or returns this
// TimeZoneUtility to an uninitialized
// state.
func (tzu *TimeZoneUtility) Empty() {
	tzu.Description = ""
	tzu.TimeIn = time.Time{}
	tzu.TimeInLoc = nil
	tzu.TimeOut = time.Time{}
	tzu.TimeOutLoc = nil
	tzu.TimeUTC = time.Time{}

}

// CopyToThis - Copies another TimeZoneUtility
// to the current TimeZoneUtility data fields.
func (tzu *TimeZoneUtility) CopyToThis(tzu2 TimeZoneUtility) {
	tzu.Empty()

	tzu.Description = tzu2.Description
	tzu.TimeIn = tzu2.TimeIn
	tzu.TimeInLoc = tzu2.TimeInLoc
	tzu.TimeOut = tzu2.TimeOut
	tzu.TimeOutLoc = tzu2.TimeOutLoc
	tzu.TimeUTC = tzu2.TimeUTC
}

// Equal - returns a boolean value indicating
// whether two TimeZoneUtility data structures
// are equivalent.
func (tzu *TimeZoneUtility) Equal(tzu2 TimeZoneUtility) bool {
	if tzu.TimeIn != tzu2.TimeIn ||
		tzu.TimeInLoc != tzu2.TimeInLoc ||
		tzu.TimeOut != tzu2.TimeOut ||
		tzu.TimeOutLoc != tzu2.TimeOutLoc ||
		tzu.TimeUTC != tzu2.TimeUTC {

		return false
	}

	return true
}

// IsValidTimeZone - Tests a Time Zone string and returns three boolean values
// designating whether the passed Time Zone string is:
// (1.) a valid time zone
// (2.) a valid iana time zone
// (3.) a valid Local time zone
func (tzu *TimeZoneUtility) IsValidTimeZone(tZone string) (isValidTz, isValidIanaTz, isValidLocalTz bool) {

	isValidTz = false

	isValidIanaTz = false

	isValidLocalTz = false

	if tZone == "" {
		return
	}

	if tZone == "Local" {
		isValidTz = true
		isValidLocalTz = true
		return
	}

	_, err := time.LoadLocation(tZone)

	if err != nil {
		return
	}

	isValidTz = true

	isValidIanaTz = true

	isValidLocalTz = false

	return

}

func (tzu *TimeZoneUtility) ReclassifyTimeWithNewTz(tIn time.Time, tZone string) (time.Time, error) {
	strTime := tzu.TimeWithoutTimeZone(tIn)

	isValidTz, _, _ := tzu.IsValidTimeZone(tZone)

	if !isValidTz {
		return time.Time{}, fmt.Errorf("TimeZoneUtility:ReclassifyTimeWithNewTz() Error: Input Time Zone is INVALID!")
	}


	tzNew, err := time.LoadLocation(tZone)

	if err != nil {
		return time.Time{}, fmt.Errorf("TimeZoneUtility:ReclassifyTimeWithNewTz() - Error from time.Location('%v') - Error: %v",tZone, err.Error())
	}

	tOut, err := time.ParseInLocation(NeutralDateFmt, strTime, tzNew)

	if err != nil {
		return tOut, fmt.Errorf("TimeZoneUtility:ReclassifyTimeWithNewTz() - Error from time.Parse - Error: %v", err.Error())
	}

	return tOut, nil
}

// SetTimeIn - Assigns value to field 'TimeIn'
func (tzu *TimeZoneUtility) SetTimeIn(tIn time.Time) {
	tzu.TimeIn = tIn
	tzu.TimeInLoc = tIn.Location()
}

// SetTimeOut - Assigns value to field 'TimeOut'
func (tzu *TimeZoneUtility) SetTimeOut(tOut time.Time) {
	tzu.TimeOut = tOut
	tzu.TimeOutLoc = tOut.Location()
}

// SetUTCTime - Assigns UTC Time to field 'TimeUTC'
func (tzu *TimeZoneUtility) SetUTCTime(t time.Time) {

	tzu.TimeUTC = t.UTC()
}

// TimeWithoutTimeZone - Returns a Time String containing
// NO time zone. - When the returned string is converted to
// time.Time - in defaults to a UTC time zone.
func (tzu *TimeZoneUtility) TimeWithoutTimeZone(t time.Time) string {
	return t.Format(NeutralDateFmt)
}