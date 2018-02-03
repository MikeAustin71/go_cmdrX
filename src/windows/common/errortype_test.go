package common

import (
	"testing"
)


func TestErrMsgType_String001(t *testing.T) {

	var r SpecErrMsgType

	r = SpecErrTypeNOERRORSALLCLEAR

	var s string

	s = r.String()

	if s != "NOERRORSALLCLEAR" {
		t.Errorf("Expected string 'NOERRORSALLCLEAR'. Instead got %v", s)
	}

}

func TestErrMsgType_String002(t *testing.T) {

	var r SpecErrMsgType

	r = SpecErrTypeFATAL

	var s string

	s = r.String()

	if s != "FATAL" {
		t.Errorf("Expected string 'FATAL'. Instead got %v", s)
	}

}

func TestErrMsgType_String003(t *testing.T) {

	var r SpecErrMsgType

	r = SpecErrTypeERROR

	var s string

	s = r.String()

	if s != "ERROR" {
		t.Errorf("Expected string 'ERROR'. Instead got %v", s)
	}

}

func TestErrMsgType_String004(t *testing.T) {

	var r SpecErrMsgType

	r = SpecErrTypeWARNING

	var s string

	s = r.String()

	if s != "WARNING" {
		t.Errorf("Expected string 'WARNING'. Instead got %v", s)
	}

}

func TestErrMsgType_String005(t *testing.T) {

	var r SpecErrMsgType

	r = SpecErrTypeINFO

	var s string

	s = r.String()

	if s != "INFO" {
		t.Errorf("Expected string 'INFO'. Instead got %v", s)
	}

}

func TestErrMsgType_String006(t *testing.T) {

	var r SpecErrMsgType

	r = SpecErrTypeDEBUG

	var s string

	s = r.String()

	if s != "DEBUG" {
		t.Errorf("Expected string 'DEBUG'. Instead got %v", s)
	}

}


func TestErrMsgType_String007(t *testing.T) {

	var r SpecErrMsgType

	r = SpecErrTypeSUCCESSFULCOMPLETION

	var s string

	s = r.String()

	if s != "SUCCESS" {
		t.Errorf("Expected string 'SUCCESS'. Instead got %v", s)
	}

}




func TestErrMsgType_Value001(t *testing.T) {

	var r SpecErrMsgType

	var i int

	r = SpecErrTypeNOERRORSALLCLEAR

	i = int(r)

	if r != 0 {
		t.Errorf("Expected SpecErrTypeNOERRORSALLCLEAR value = ZERO (0). Instead got %v", i)
	}

}

func TestErrMsgType_Value002(t *testing.T) {

	var r SpecErrMsgType

	var i int

	r = SpecErrTypeFATAL

	i = int(r)

	if r != 1 {
		t.Errorf("Expected SpecErrTypeFATAL value = 1. Instead got %v", i)
	}

}

func TestErrMsgType_Value003(t *testing.T) {

	var r SpecErrMsgType

	var i int

	r = SpecErrTypeERROR

	i = int(r)

	if r != 2 {
		t.Errorf("Expected SpecErrTypeERROR value = 2. Instead got %v", i)
	}

}

func TestErrMsgType_Value004(t *testing.T) {

	var r SpecErrMsgType

	var i int

	r = SpecErrTypeWARNING

	i = int(r)

	if r != 3 {
		t.Errorf("Expected SpecErrTypeWARNING value = 3. Instead got %v", i)
	}

}

func TestErrMsgType_Value005(t *testing.T) {

	var r SpecErrMsgType

	var i int

	r = SpecErrTypeINFO

	i = int(r)

	if r != 4 {
		t.Errorf("Expected SpecErrTypeINFO value = 4. Instead got %v", i)
	}

}

func TestErrMsgType_Value006(t *testing.T) {

	var r SpecErrMsgType

	var i int

	r = SpecErrTypeDEBUG

	i = int(r)

	if r != 5 {
		t.Errorf("Expected SpecErrTypeDEBUG value = 5. Instead got %v", i)
	}

}


func TestErrMsgType_Value007(t *testing.T) {

	var r SpecErrMsgType

	var i int

	r = SpecErrTypeSUCCESSFULCOMPLETION

	i = int(r)

	if r != 6 {
		t.Errorf("Expected SpecErrTypeSUCCESSFULCOMPLETION value = 6. Instead got %v", i)
	}

}

func TestErrMsgType_Initialization001(t *testing.T) {
	se := SpecErr{}

	if se.ErrorMsgType != SpecErrTypeNOERRORSALLCLEAR {
		t.Error("Exppected uninitialized SpecErr object to set se.ErrMsgType=='SpecErrTypeNOERRORSALLCLEAR'. It did NOT! ")
	}

}

