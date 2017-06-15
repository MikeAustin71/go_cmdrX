package common

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeZoneUtilConvertTz(t *testing.T) {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	ianaCentralTz := "America/Chicago"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu := TimeZoneUtility{}
	tzu.ConvertTz(tIn, ianaCentralTz)
	tIn = tzu.TimeOut
	tzu.Empty()
	tzu.ConvertTz(tIn, ianaPacificTz)

	exTIn := "2017-04-29 19:54:30 -0500 CDT"
	actTIn := fmt.Sprintf("%v", tzu.TimeIn)
	if actTIn != exTIn {
		t.Error(fmt.Sprintf("Expected tzu.TimeIn %v, got ", exTIn), actTIn)
	}

	exTInLoc := ianaCentralTz
	actTInLoc := fmt.Sprintf("%v", tzu.TimeInLoc)
	if actTInLoc != exTInLoc {
		t.Error(fmt.Sprintf("Expected tzu.TimeInLoc %v, got", exTInLoc), actTInLoc)
	}

	exTOut := "2017-04-29 17:54:30 -0700 PDT"
	actTOut := fmt.Sprintf("%v", tzu.TimeOut)
	if actTOut != exTOut {
		t.Error(fmt.Sprintf("Expected tzu.TimeOut %v, got", exTOut), actTOut)
	}

	exTOutLoc := "America/Los_Angeles"
	actTOutLoc := fmt.Sprintf("%v", tzu.TimeOutLoc)

	if actTOutLoc != exTOutLoc {
		t.Error(fmt.Sprintf("Expected tzu.TimeOutLoc %v, got", exTOutLoc), actTOutLoc)
	}

	exUTC := "2017-04-30 00:54:30 +0000 UTC"
	actUTC := fmt.Sprintf("%v", tzu.TimeUTC)

	if exUTC != actUTC {
		t.Error(fmt.Sprintf("Expected tzu.TimeUTC %v, got", exUTC), actUTC)
	}

}

func TestInvalidTargetTzInConversion(t *testing.T) {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	// Invalid Target Iana Time Zone
	invalidTz := "XUZ Time Zone"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu := TimeZoneUtility{}
	err := tzu.ConvertTz(tIn, invalidTz)

	if err == nil {
		t.Error("ConverTz() failed to detect INVALID Tartet Time Zone. Got: ", "err==nil")
	}

}

func TestTimeZoneUtility_IsValidTimeZone(t *testing.T) {
	tIn := time.Now()

	tzu := TimeZoneUtility{}

	isValidTz, isValidIanaTz, isValidLocalTz := tzu.IsValidTimeZone(tIn.Location().String())

	if isValidTz == false {
		t.Error("Expected Now() Location to yield 'Local' Time Zone isValidTz == VALID ('true'), instead got: ", isValidTz)
	}

	if isValidIanaTz == true {
		t.Error("Passed Time Zone was 'Local' Time Zone. Expected isValidIanaTz == false, got: ", isValidIanaTz)
	}

	if isValidLocalTz == false {
		t.Error("Passed Time Zone was 'Local' Time Zone. Expected isValidLocalTz == true, got: ", isValidIanaTz)
	}

}

func TestCDTIsValidIanaTimeZone(t *testing.T) {

	tzu := TimeZoneUtility{}

	isValidTz, isValidIanaTz, isValidLocalTz := tzu.IsValidTimeZone("America/Chicago")

	if isValidTz == false {
		t.Error("Expected 'America/Chicago' to yield isValidTz = 'true', instead got", isValidTz)
	}

	if isValidIanaTz == false {
		t.Error("Expected 'America/Chicago' to yield isValidIanaTz = 'true', instead got", isValidIanaTz)
	}

	if isValidLocalTz == true {
		t.Error("Expected 'America/Chicago' to yield isValidLocalTz = 'false', instead got", isValidLocalTz)
	}

}

func TestTimeZoneUtility_ReclassifyTimeWithTzLocal(t *testing.T) {
	/*
			Example Method: ReclassifyTimeWithNewTz()
		Input Time :  2017-04-29 17:54:30 -0700 PDT
		Output Time:  2017-04-29 17:54:30 -0500 CDT
	*/

	tPacific := "2017-04-29 17:54:30 -0700 PDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	tz := TimeZoneUtility{}
	tIn, err := time.Parse(fmtstr, tPacific)
	if err != nil {
		fmt.Printf("Error returned from time.Parse: %v", err.Error())
		return
	}

	tOut, err := tz.ReclassifyTimeWithNewTz(tIn, "Local")

	tOutLoc := tOut.Location()

	if tOutLoc.String() != "Local" {
		t.Error("Expected tOutLocation == 'Local', instead go Location: ", tOutLoc.String())
	}

}

func TestTimeZoneUtility_ReclassifyTimeWithNewTz(t *testing.T) {

	tPacific := "2017-04-29 17:54:30 -0700 PDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	tz := TimeZoneUtility{}
	tIn, err := time.Parse(fmtstr, tPacific)
	if err != nil {
		fmt.Printf("Error returned from time.Parse: %v", err.Error())
		return
	}

	tOut, err := tz.ReclassifyTimeWithNewTz(tIn, TzUsHawaii)

	tOutLoc := tOut.Location()

	if tOutLoc.String() != TzUsHawaii {
		t.Error(fmt.Sprintf("Expected tOutLocation == '%v', instead go Location: ", TzUsHawaii), tOutLoc.String())
	}

}

func TestReclassifyTimeAsMountain(t *testing.T) {

	tPacific := "2017-04-29 17:54:30 -0700 PDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	tz := TimeZoneUtility{}
	tIn, err := time.Parse(fmtstr, tPacific)
	if err != nil {
		fmt.Printf("Error returned from time.Parse: %v", err.Error())
		return
	}

	tOut, err := tz.ReclassifyTimeWithNewTz(tIn, TzUsMountain)

	tOutLoc := tOut.Location()

	if tOutLoc.String() != TzUsMountain {
		t.Error(fmt.Sprintf("Expected tOutLocation == '%v', instead go Location: ", TzUsHawaii), tOutLoc.String())
	}

}
