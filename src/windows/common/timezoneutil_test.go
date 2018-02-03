package common

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeZoneUtilConvertTz(t *testing.T) {
	utcTime := "2017-04-30 00:54:30 +0000 UTC"
	pacificTime := "2017-04-29 17:54:30 -0700 PDT"
	mountainTime := "2017-04-29 18:54:30 -0600 MDT"
	centralTime := "2017-04-29 19:54:30 -0500 CDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"
	ianaCentralTz := "America/Chicago"
	ianaMountainTz := "America/Denver"
	tPacificIn, err := time.Parse(fmtstr, pacificTime)

	if err != nil {
		t.Errorf("Received error from time parse tPacificIn: %v", err.Error())
	}

	tzu := TimeZoneUtility{}
	tzuCentral, err := tzu.ConvertTz(tPacificIn, ianaCentralTz)

	if err != nil {
		t.Errorf("Error from TimeZoneUtility.ConvertTz(). Error: %v", err.Error())
	}

	centralTOut := tzuCentral.TimeOut.Format(fmtstr)

	if centralTime != centralTOut {
		t.Errorf("Expected tzuCentral.TimeOut %v, got %v", centralTime, centralTOut)
	}

	tzuMountain, err := tzu.ConvertTz(tzuCentral.TimeOut, ianaMountainTz)

	if err != nil {
		t.Errorf("Error from  tzuMountain TimeZoneUtility.ConvertTz(). Error: %v", err.Error())
	}

	mountainTOut := tzuMountain.TimeOut.Format(fmtstr)

	if mountainTime != mountainTOut {
		t.Errorf("Expected tzuMountain.TimeOut %v, got %v", mountainTime, mountainTOut)
	}

	tzuPacific, err := tzu.ConvertTz(tzuMountain.TimeOut, ianaPacificTz)

	if err != nil {
		t.Errorf("Error from  tzuMountain TimeZoneUtility.ConvertTz(). Error: %v", err.Error())
	}

	pacificTOut := tzuPacific.TimeOut.Format(fmtstr)

	if pacificTime != pacificTOut {

		t.Errorf("Expected tzuPacific.TimeOut %v, got %v", pacificTime, pacificTOut)
	}

	exTOutLoc := "America/Los_Angeles"

	if exTOutLoc != tzuPacific.TimeOutLoc.String() {
		t.Errorf("Expected tzu.TimeOutLoc %v, got %v", exTOutLoc, tzuPacific.TimeOutLoc.String())
	}

	pacificUtcOut := tzuPacific.TimeUTC.Format(fmtstr)

	if utcTime != pacificUtcOut {
		t.Errorf("Expected tzuPacific.TimeUTC %v, got %v", utcTime, pacificUtcOut)
	}

	centralUtcOut := tzuCentral.TimeUTC.Format(fmtstr)

	if utcTime != centralUtcOut {
		t.Errorf("Expected tzuCentral.TimeUTC %v, got %v", utcTime, pacificUtcOut)
	}

	mountainUtcOut := tzuMountain.TimeUTC.Format(fmtstr)

	if utcTime != mountainUtcOut {
		t.Errorf("Expected tzuMountain.TimeUTC %v, got %v", utcTime, pacificUtcOut)
	}

}

func TestInvalidTargetTzInConversion(t *testing.T) {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	// Invalid Target Iana Time Zone
	invalidTz := "XUZ Time Zone"
	tIn, _ := time.Parse(fmtstr, tstr)
	tzu := TimeZoneUtility{}
	_, err := tzu.ConvertTz(tIn, invalidTz)

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

func TestTimeZoneUtility_MakeDateTz(t *testing.T) {
	tPacific := "2017-04-29 17:54:30 -0700 PDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"

	dtTzDto := DateTzDto{Year: 2017, Month: 4, Day: 29, Hour: 17, Minute: 54, Second: 30, IANATimeZone: "America/Los_Angeles"}

	tzu := TimeZoneUtility{}

	tOut, err := tzu.MakeDateTz(dtTzDto)

	if err != nil {
		t.Errorf("Error returned from TimeZoneUtility.MakeDateTz(). Error: %v", err.Error())
	}

	tOutStr := tOut.Format(fmtstr)

	if tPacific != tOutStr {
		t.Errorf("Error - Expected output time string: %v. Instead, got %v.", tPacific, tOutStr)
	}

}

func TestTimeZoneUtility_ConvertTz(t *testing.T) {

	pacificTime := "2017-04-29 17:54:30 -0700 PDT"
	centralTime := "2017-04-29 19:54:30 -0500 CDT"
	ianaCentralTz := "America/Chicago"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"

	tPacific, err := time.Parse(fmtstr, pacificTime)

	if err != nil {
		t.Errorf("Error from time.Parse. pacificTime = %v. Error= %v", pacificTime, err.Error())
	}

	tzuCentral, err := TimeZoneUtility{}.ConvertTz(tPacific, ianaCentralTz)

	tOutStr := tzuCentral.TimeOut.Format(fmtstr)

	if centralTime != tOutStr {
		t.Errorf("Error. Central Time zone conversion failed! Expected %v. Instead, got %v.", centralTime, tOutStr)
	}

}

func TestTimeZoneUtility_ConvertTz_02(t *testing.T) {
	moscowTz := "Europe/Moscow"
	beijingTz := "Asia/Shanghai"
	centralTime := "2017-04-29 19:54:30 -0500 CDT"
	moscowTime := "2017-04-30 03:54:30 +0300 MSK"
	utcTime := "2017-04-30 00:54:30 +0000 UTC"
	beijingTime := "2017-04-30 08:54:30 +0800 CST"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"

	tCentral, err := time.Parse(fmtstr, centralTime)

	if err != nil {
		t.Errorf("Error from time.Parse. centralTime = %v. Error= %v", centralTime, err.Error())
	}

	tzuMoscow, err := TimeZoneUtility{}.ConvertTz(tCentral, moscowTz)

	if err != nil {
		t.Errorf("Error from TimeZoneUtility{}.ConvertTz. Central Time = %v. Error= %v", centralTime, err.Error())
	}

	moscowTOut := tzuMoscow.TimeOut.Format(fmtstr)

	if moscowTime != moscowTOut {
		t.Errorf("Error. Moscow Time zone conversion failed! Expected %v. Instead, got %v.", moscowTime, moscowTOut)
	}

	tzuBeijing, err := tzuMoscow.ConvertTz(tCentral, beijingTz)

	if err != nil {
		t.Errorf("Error from tzuMoscow.ConvertTz. Central Time = %v. Error= %v", centralTime, err.Error())
	}

	beijingTOut := tzuBeijing.TimeOut.Format(fmtstr)

	if beijingTime != beijingTOut {
		t.Errorf("Error. Beijing Time zone conversion failed! Expected %v. Instead, got %v.", beijingTime, beijingTOut)
	}

	utcTOut := tzuBeijing.TimeUTC.Format(fmtstr)

	if utcTime != utcTOut {
		t.Errorf("Error. UTC Time from tzuBeijing.TimeUTC failed! Expected %v. Instead, got %v.", utcTime, utcTOut)

	}

}

func TestTimeZoneUtility_ConvertTz_03(t *testing.T) {
	t1UTCStr := "2017-07-02 22:00:18.423111300 +0000 UTC"
	fmtstr := "2006-01-02 15:04:05.000000000 -0700 MST"
	t2LocalStr := "2017-07-02 17:00:18.423111300 -0500 CDT"
	localTzStr := "America/Chicago"

	t1, _ := time.Parse(fmtstr, t1UTCStr)

	tz := TimeZoneUtility{}

	tzLocal, _ := tz.ConvertTz(t1, localTzStr)
	t1OutStr := tzLocal.TimeIn.Format(fmtstr)
	t2OutStr := tzLocal.TimeOut.Format(fmtstr)

	if t1UTCStr != t1OutStr {
		t.Errorf("Expected Input Time: %v. Error - Instead, got %v", t1UTCStr, t1OutStr)
	}

	if t2LocalStr != t2OutStr {
		t.Errorf("Expected Output Local Time: %v. Error - Instead, got %v", t2LocalStr, t2OutStr)
	}

}

func TestTimeZoneUtility_Location_01(t *testing.T) {
	utcTime := "2017-04-30 00:54:30 +0000 UTC"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	ianaPacificTz := "America/Los_Angeles"

	tUtc, _ := time.Parse(fmtstr, utcTime)

	tzu := TimeZoneUtility{}

	tzuPacific, err := tzu.ConvertTz(tUtc, ianaPacificTz)

	if err != nil {
		t.Errorf("Error from TimeZoneUtility{}.ConvertTz. Utc Time = %v. Error= %v", utcTime, err.Error())
	}

	tzOutPacific := tzuPacific.TimeOutLoc.String()

	if tzOutPacific != ianaPacificTz {
		t.Errorf("Error: Expected tzOutPacific %v. Instead, got %v", ianaPacificTz, tzOutPacific)
	}

}
func TestTimeZoneUtility_Location_02(t *testing.T) {

	pacificTime := "2017-04-29 17:54:30 -0700 PDT"
	fmtstr := "2006-01-02 15:04:05 -0700 MST"
	tPacific, _ := time.Parse(fmtstr, pacificTime)

	tzu := TimeZoneUtility{}

	tzuLocal, err := tzu.ConvertTz(tPacific, "Local")

	if err != nil {
		t.Errorf("Error from TimeZoneUtility{}.ConvertTz. Pacific Time = %v. Error= %v", pacificTime, err.Error())
	}

	tzOutLocal := tzuLocal.TimeOutLoc.String()

	if "Local" != tzOutLocal {
		t.Errorf("Error: Expected tzOutLocal 'Local'. Instead, got %v", tzOutLocal)
	}

}
