package common

import "testing"

func TestLogMsgTypeText(t *testing.T) {
	var r LogMsgType

	r = LogINFOMSGTYPE

	var s string

	s = r.String()

	if s != "INFO" {
		t.Error("Expected string 'INFO' got ", s)
	}

}

func TestLogMsgTypeValue(t *testing.T) {
	var r LogMsgType

	var i int

	r = LogINFOMSGTYPE

	i = int(r)

	if r != 1 {
		t.Error("Expected 'INFOMSGTYPE' value = 1, got ", i)
	}

}
