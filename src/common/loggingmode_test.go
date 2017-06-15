package common

import "testing"

func TestLoggingModeText(t *testing.T) {
	var r LoggingMode

	var s string

	r = LogVERBOSE

	s = r.String()

	if s != "VERBOSE" {
		t.Error("Expected string 'VERBOSE' got ", s)
	}
}

func TestLoggingModeValue(t *testing.T) {
	var r LoggingMode

	var i int

	r = LogVERBOSE

	i = int(r)

	if r != 2 {
		t.Error("Expected 'VERBOSE' value 2, got", i)
	}

}
