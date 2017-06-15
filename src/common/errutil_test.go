package common

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorUtility(t *testing.T) {

	ex1 := "errutil_test.go"
	ex2 := "TestErrorUtility"
	ex3 := int64(10000)

	ex4 := "errprefix"
	ex5 := int64(334)
	ex6 := "Test Error #1"
	ex7 := ex5 + ex3

	err := errors.New(ex6)
	a := make([]ErrBaseInfo, 0, 10)
	bi := ErrBaseInfo{}.New(ex1, ex2, ex3)
	se := SpecErr{}.InitializeBaseInfo(a, bi)
	x := se.New(ex4, err, true, ex5)

	if x.PrefixMsg != ex4 {
		t.Error(fmt.Sprintf("Expected '%v' got", ex4), x.PrefixMsg)
	}

	if x.ErrMsg != ex6 {
		t.Error(fmt.Sprintf("Expected '%v' got", ex6), x.ErrMsg)
	}

	if x.BaseInfo.SourceFileName != ex1 {
		t.Error(fmt.Sprintf("Expected '%v' got", ex1), x.BaseInfo.SourceFileName)
	}

	if x.BaseInfo.FuncName != ex2 {
		t.Error(fmt.Sprintf("Expected '%v' got", ex2), x.BaseInfo.FuncName)
	}

	if x.ErrNo != ex7 {
		t.Error(fmt.Sprintf("Expected '%v' got", ex7), x.ErrNo)
	}

}

func TestUninitializedBaseInfo(t *testing.T) {
	var se SpecErr

	if se.BaseInfo.SourceFileName != "" {
		t.Error("String SourceFileName was uninitialized. Was expecting empty string, got", se.BaseInfo.SourceFileName)
	}

	if se.BaseInfo.FuncName != "" {
		t.Error("String FuncName was uninitialized. Was expecting empty string, got", se.BaseInfo.FuncName)
	}

	if se.BaseInfo.BaseErrorID != 0 {
		t.Error("Int64 BaseErrorID was uninitialized. Was expecting value of zero, got", se.BaseInfo.BaseErrorID)
	}

}

func TestInitializeParentInfo(t *testing.T) {

	bi := ErrBaseInfo{}

	x := bi.New("TestSourceFileName", "TestFuncName", 9000)
	y := bi.New("TestSrcFileName2", "TestFuncName2", 14000)
	z := bi.New("TestSrcFileName3", "TestFuncName3", 15000)

	se := SpecErr{}
	se.ParentInfo = append(se.ParentInfo, x)
	se.ParentInfo = append(se.ParentInfo, y)
	se.ParentInfo = append(se.ParentInfo, z)

	l := len(se.ParentInfo)

	if l != 3 {
		t.Error("Expected ParentInfo Length of 3, got", l)
	}

	if se.ParentInfo[1].FuncName != "TestFuncName2" {
		t.Error("Expected 2nd Element 'TestFuncName2', got", se.ParentInfo[1].FuncName)
	}

}

func TestAddSlicesParentInfo(t *testing.T) {
	var bi ErrBaseInfo
	x := bi.New("TestSourceFileName", "TestFuncName", 9000)
	y := bi.New("TestSrcFileName2", "TestFuncName2", 14000)
	z := bi.New("TestSrcFileName3", "TestFuncName3", 15000)

	a := make([]ErrBaseInfo, 0, 30)

	a = append(a, x, y, z)

	var se SpecErr

	se.ParentInfo = a

	l := len(se.ParentInfo)

	if l != 3 {
		t.Error("Expected ParentInfo Length of 3, got", l)
	}

	if se.ParentInfo[1].FuncName != "TestFuncName2" {
		t.Error("Expected 2nd Element 'TestFuncName2', got", se.ParentInfo[1].FuncName)
	}

}

func TestSetParentInfo(t *testing.T) {
	var bi ErrBaseInfo
	x := bi.New("TestSourceFileName", "TestFuncName", 9000)
	y := bi.New("TestSrcFileName2", "TestFuncName2", 14000)
	z := bi.New("TestSrcFileName3", "TestFuncName3", 15000)

	a := make([]ErrBaseInfo, 0, 30)

	a = append(a, x, y, z)

	var se SpecErr

	se.ParentInfo = se.DeepCopyParentInfo(a)

	l := len(se.ParentInfo)

	if l != 3 {
		t.Error("Expected ParentInfo length of 3, go length of ", l)
	}

	if se.ParentInfo[1].FuncName != "TestFuncName2" {
		t.Error("Expected 2nd Element 'TestFuncName2', got", se.ParentInfo[1].FuncName)
	}
}

func TestSetErrDetail(t *testing.T) {

	ex1 := "errutil_test.go"
	ex2 := "TestErrorUtility"
	ex3 := int64(10000)

	ex4 := "errprefix"
	ex5 := int64(338)
	ex6 := "Test Error #21"
	ex7 := ex5 + ex3

	err := errors.New(ex6)
	a := make([]ErrBaseInfo, 0, 10)
	bi := ErrBaseInfo{}.New(ex1, ex2, ex3)

	x := SpecErr{}.InitializeBaseInfo(a, bi).New(ex4, err, true, ex5)

	if x.ErrNo != ex7 {
		t.Error(fmt.Sprintf("Expected Err No '%v', got", ex7), x.ErrNo)
	}

	if x.BaseInfo.SourceFileName != ex1 {
		t.Error(fmt.Sprintf("Expected Source File: '%v',got", ex1), x.BaseInfo.SourceFileName)
	}

	if x.BaseInfo.FuncName != ex2 {
		t.Error(fmt.Sprintf("Expected FuncName: '%v', got", ex2), x.BaseInfo.FuncName)
	}

}

func TestIsSpecErrNo(t *testing.T) {

	s := SpecErr{}.SignalNoErrors()

	isErr := CheckIsSpecErr(s)

	if isErr {
		t.Error("Expected CheckIsSpecErr() to return false, go", isErr)
	}

}

func TestIsSpecErrYes(t *testing.T) {
	ex1 := "errprefix"
	ex4 := int64(334)
	ex99 := "Test Error #1"

	err := errors.New(ex99)

	x := SpecErr{}.New(ex1, err, true, ex4)

	isErr := CheckIsSpecErr(x)

	if !isErr {
		t.Error("Expected CheckIsSpecErr() to return true, go", isErr)
	}

}

func TestSetNoErr(t *testing.T) {
	x := SpecErr{}.SignalNoErrors()

	if x.IsErr {
		t.Error("Expected IsErr = 'false', got", x.IsErr)
	}
}

func TestQuickInitialize(t *testing.T) {

	ex1 := "prefixMsg"
	ex2 := "Error Msg X"
	err := errors.New(ex2)
	ex3 := false
	ex4 := int64(499)

	x := SpecErr{}.New(ex1, err, ex3, ex4)

	if x.ErrNo != ex4 {
		t.Error(fmt.Sprintf("Expected ErrNo: '%v', got", ex4), x.ErrNo)
	}

	if x.PrefixMsg != ex1 {
		t.Error(fmt.Sprintf("Expected PrefixMsg == '%v', got", ex1), x.PrefixMsg)
	}

	if x.ErrMsg != ex2 {
		t.Error(fmt.Sprintf("Expected Error Msg '%v', got,", ex2), x.ErrMsg)
	}

	if x.IsPanic == true {
		t.Error(fmt.Sprintf("Expected IsPanic == '%v', got", ex3), x.IsPanic)
	}
}

func TestFullInitialize(t *testing.T) {

	bi := ErrBaseInfo{}

	f := bi.New("TestSourceFileName", "TestFuncName", 9000)
	g := bi.New("TestSrcFileName2", "TestFuncName2", 14000)
	h := bi.New("TestSrcFileName3", "TestFuncName3", 15000)

	ex1 := make([]ErrBaseInfo, 0, 10)
	ex1 = append(ex1, f, g, h)

	ex2_1 := "TestSrcFileName99"
	ex2_2 := "TestFuncName99"
	ex2_3 := int64(16000)
	ex2 := bi.New(ex2_1, ex2_2, ex2_3)

	ex3 := "prefixString"
	ex4 := "Error Msg 99"
	err := errors.New(ex4)
	ex5 := false
	ex6 := int64(22)
	ex7 := int64(16022)

	x := SpecErr{}.Initialize(ex1, ex2, ex3, err, ex5, ex6)

	pl := len(x.ParentInfo)

	if pl != 3 {
		t.Error("Expected ParentInfo length == 3, got", pl)
	}

	if x.PrefixMsg != ex3 {
		t.Error(fmt.Sprintf("Expected PrefixMsg '%v', got", ex3), x.PrefixMsg)
	}

	if x.ErrMsg != ex4 {
		t.Error(fmt.Sprintf("Expected ErrMsg '%v', got", ex4), x.ErrMsg)
	}

	if x.IsErr == false {
		t.Error(fmt.Sprintf("Expected IsErr == '%v', got", true), x.IsErr)
	}

	if x.IsPanic != ex5 {
		t.Error(fmt.Sprintf("Expected IsPanic == '%v', got", ex5), x.IsPanic)
	}

	if x.ErrNo != ex7 {
		t.Error(fmt.Sprintf("Expected ErrNo '%v', got", ex7), x.ErrNo)
	}

	if x.BaseInfo.SourceFileName != ex2_1 {
		t.Error(fmt.Sprintf("Expected SourceFileName '%v', got", ex2_1), x.BaseInfo.SourceFileName)
	}

	if x.BaseInfo.FuncName != ex2_2 {
		t.Error(fmt.Sprintf("Expected Function Name '%v', got", ex2_2), x.BaseInfo.FuncName)
	}

	if x.BaseInfo.BaseErrorID != ex2_3 {
		t.Error(fmt.Sprintf("Expected Base Error ID '%v', got", ex2_3), x.BaseInfo.BaseErrorID)
	}

}

func TestBlankInitialize(t *testing.T) {

	ex1 := "prefixString"
	ex2 := "Error Msg 99"
	err := errors.New(ex2)
	ex3 := false
	ex4 := int64(22)

	x := SpecErr{}.Initialize(blankParentInfo, blankErrBaseInfo, ex1, err, ex3, ex4)

	if x.PrefixMsg != ex1 {
		t.Error(fmt.Sprintf("Expected PrefixMsg '%v', got", ex1), x.PrefixMsg)
	}

	if x.ErrMsg != ex2 {
		t.Error(fmt.Sprintf("Expected ErrMsg '%v', got", ex2), x.ErrMsg)
	}

	if x.IsErr == false {
		t.Error(fmt.Sprintf("Expected IsErr == '%v', got", true), x.IsErr)
	}

	if x.IsPanic != ex3 {
		t.Error(fmt.Sprintf("Expected IsPanic == '%v', got", ex3), x.IsPanic)
	}

	if x.ErrNo != ex4 {
		t.Error(fmt.Sprintf("Expected ErrNo '%v', got", ex4), x.ErrNo)
	}

	if x.BaseInfo.SourceFileName != "" {
		t.Error("Expected BaseInfo Source File Name == '', got", x.BaseInfo.SourceFileName)
	}

	if x.BaseInfo.FuncName != "" {
		t.Error("Expected BaseInfo FuncName == '', got", x.BaseInfo.FuncName)
	}

	if x.BaseInfo.BaseErrorID != int64(0) {
		t.Error("Expected BaseInfo BaseErrorID == 'Zero', got", x.BaseInfo.BaseErrorID)
	}

}

func TestAddBaseInfoToParent(t *testing.T) {

	bi := ErrBaseInfo{}

	f := bi.New("TestSourceFileName", "TestFuncName", 9000)
	g := bi.New("TestSrcFileName2", "TestFuncName2", 14000)
	h := bi.New("TestSrcFileName3", "TestFuncName3", 15000)

	ex1 := make([]ErrBaseInfo, 0, 10)
	ex1 = append(ex1, f, g, h)

	ex2_1 := "TestSrcFileName99"
	ex2_2 := "TestFuncName99"
	ex2_3 := int64(16000)
	ex2 := bi.New(ex2_1, ex2_2, ex2_3)

	ex3 := "prefixString"
	ex4 := "Error Msg 99"
	err := errors.New(ex4)
	ex5 := false
	ex6 := int64(22)

	x := SpecErr{}.Initialize(ex1, ex2, ex3, err, ex5, ex6)

	pl := len(x.ParentInfo)

	if pl != 3 {
		t.Error("Expected Parent Info to contain 3-elements. Actual number of elements was ", pl)
	}

	p2 := x.AddBaseToParentInfo()

	pl2 := len(p2)

	if pl2 != 4 {
		t.Error("Expected New Parent Info to contain 4-elements. Actual number of elements was ", pl2)
	}

}
