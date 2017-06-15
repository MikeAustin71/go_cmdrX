package common

import (
	"fmt"
	"testing"
	"time"
)

func TestGetDateTimeStr(t *testing.T) {
	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	expected := "20170429195430"

	testTime, _ := time.Parse(fmtstr, tstr)

	result := DateTimeUtility{}.GetDateTimeStr(testTime)

	if result != expected {
		t.Error("Expected '20170429195430' got", result)
	}

}

func TestGetDateTimeSecText(t *testing.T) {

	tstr := "04/29/2017 19:54:30 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05 -0700 MST"
	expected := "2017-04-29 19:54:30"
	testTime, _ := time.Parse(fmtstr, tstr)
	result := DateTimeUtility{}.GetDateTimeSecText(testTime)

	if result != expected {
		t.Error("Expected '", expected, "' got", result)
	}

}

func TestGetDateTimeNanoSecText(t *testing.T) {
	tstr := "04/29/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected := "2017-04-29 19:54:30.123456489"
	testTime, _ := time.Parse(fmtstr, tstr)
	dt := DateTimeUtility{}
	result := dt.GetDateTimeNanoSecText(testTime)

	if result != expected {
		t.Error("Expected '", expected, "' got", result)
	}

}

func TestGetDateTimeEverything(t *testing.T) {
	tstr := "04/29/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected := "Saturday April 29, 2017 19:54:30.123456489 -0500 CDT"
	testTime, _ := time.Parse(fmtstr, tstr)
	dt := DateTimeUtility{}
	str := dt.GetDateTimeEverything(testTime)

	if str != expected {
		t.Error(fmt.Sprintf("Expected datetime: '%v', got", expected), str)
	}

}

func TestCustomDateTimeFormat(t *testing.T) {
	dt := DateTimeUtility{}
	tstr := "04/29/2017 19:54:30.123456489 -0500 CDT"
	fmtstr := "01/02/2006 15:04:05.000000000 -0700 MST"
	expected := "Saturday April 29, 2017 19:54:30.123456489 -0500 CDT"
	testTime, _ := time.Parse(fmtstr, tstr)
	result := dt.GetDateTimeCustomFmt(testTime, FmtDateTimeEverything)

	if result != expected {
		t.Error(fmt.Sprintf("Expected: %v, got", expected), result)
	}

}
