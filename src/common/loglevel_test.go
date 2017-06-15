package common

import "testing"

func TestLogLevelString(t *testing.T) {
	var r LogLevel

	r = LogOPERROR

	var s string

	s = r.String()

	if s != "OPERROR" {
		t.Error("Expected string 'OPERROR' got ", s)
	}

}

func TestLogLevelValue(t *testing.T) {
	var r LogLevel

	var i int

	r = LogINFO

	i = int(r)

	if r != 3 {
		t.Error("Expected INFO value 3, got", i)
	}
}
