package common

import "testing"

/*  'opsmsgtype_test.go' is located in source code
		repository:

		https://github.com/MikeAustin71/ErrHandlerGo.git

*/

func TestOpsMsgTypeText001(t *testing.T) {
	var r OpsMsgType

	r = OpsMsgTypeNOERRORNOMSGS

	var s string

	s = r.String()

	if s != "NOERRORSNOMSGS" {
		t.Errorf("Expected string 'NOERRORSNOMSGS'. Instead, got %v", s)
	}

}

func TestOpsMsgTypeText002(t *testing.T) {
	var r OpsMsgType

	r = OpsMsgTypeERRORMSG

	var s string

	s = r.String()

	if s != "ERROR" {
		t.Errorf("Expected string 'ERROR'. Instead, got %v", s)
	}

}

func TestOpsMsgTypeText003(t *testing.T) {
	var r OpsMsgType

	r = OpsMsgTypeFATALERRORMSG

	var s string

	s = r.String()

	if s != "FATALERROR" {
		t.Errorf("Expected string 'FATALERROR'. Instead, got %v", s)
	}

}


func TestOpsMsgTypeText004(t *testing.T) {
	var r OpsMsgType

	r = OpsMsgTypeINFOMSG

	var s string

	s = r.String()

	if s != "INFO" {
		t.Errorf("Expected string 'INFO'. Instead, got %v", s)
	}

}

func TestOpsMsgTypeText005(t *testing.T) {
	var r OpsMsgType

	r = OpsMsgTypeWARNINGMSG

	var s string

	s = r.String()

	if s != "WARNING" {
		t.Errorf("Expected string 'WARNING'. Instead, got %v", s)
	}

}

func TestOpsMsgTypeText006(t *testing.T) {
	var r OpsMsgType

	r = OpsMsgTypeDEBUGMSG

	var s string

	s = r.String()

	if s != "DEBUG" {
		t.Errorf("Expected string 'DEBUG'. Instead, got %v", s)
	}

}

func TestOpsMsgTypeText007(t *testing.T) {
	var r OpsMsgType

	r = OpsMsgTypeSUCCESSFULCOMPLETION

	var s string

	s = r.String()

	if s != "SUCCESS" {
		t.Errorf("Expected string 'SUCCESS'. Instead, got %v", s)
	}

}


func TestOpsMsgTypeValue001(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsMsgTypeNOERRORNOMSGS

	i = int(r)

	if r != 0 {
		t.Errorf("Expected 'OpsMsgTypeNOERRORNOMSG' value = 0. Instead, got %v", i)
	}

}

func TestOpsMsgTypeValue002(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsMsgTypeERRORMSG

	i = int(r)

	if r != 1 {
		t.Errorf("Expected 'OpsMsgTypeERRORMSG' value = 1. Instead, got %v", i)
	}

}

func TestOpsMsgTypeValue003(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsMsgTypeFATALERRORMSG

	i = int(r)

	if r != 2 {
		t.Errorf("Expected 'OpsMsgTypeINFOMSG' value = 2. Instead, got %v", i)
	}

}

func TestOpsMsgTypeValue004(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsMsgTypeINFOMSG

	i = int(r)

	if r != 3 {
		t.Errorf("Expected 'OpsMsgTypeINFOMSG' value = 3. Instead, got %v", i)
	}

}

func TestOpsMsgTypeValue005(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsMsgTypeWARNINGMSG

	i = int(r)

	if r != 4 {
		t.Errorf("Expected 'OpsMsgTypeWARNINGMSG' value = 4. Instead, got %v", i)
	}

}

func TestOpsMsgTypeValue006(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsMsgTypeDEBUGMSG

	i = int(r)

	if r != 5 {
		t.Errorf("Expected 'OpsMsgTypeDEBUGMSG' value = 5. Instead, got %v", i)
	}

}

func TestOpsMsgTypeValue007(t *testing.T) {
	var r OpsMsgType

	var i int

	r = OpsMsgTypeSUCCESSFULCOMPLETION

	i = int(r)

	if r != 6 {
		t.Errorf("Expected 'OpsMsgTypeSUCCESSFULCOMPLETION' value = 6. Instead, got %v", i)
	}

}

func TestOpsMsgTypeInitialization001(t *testing.T) {

	om:= OpsMsgDto{}

	if om.MsgType != OpsMsgTypeNOERRORNOMSGS {
		t.Errorf("Expected uninitialized OpsMsgDto object to show MsgType= 'OpsMsgTypeNOERRORNOMSG'.  Instead, MsgType= '%v' ", om.MsgType)
	}

}