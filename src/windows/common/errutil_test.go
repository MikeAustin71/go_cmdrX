package common

import (
	"errors"
	"fmt"
	"testing"
	"strings"
)

func TestSpecErr_New_001(t *testing.T) {

	ex1 := "errutil_test.go"
	ex1ParentObj := "ErrorObj"
	ex2 := "TestErrorUtility"
	ex3 := int64(10000)


	ex5 := int64(334)
	ex6 := "Test Error #1"
	ex7 := ex5 + ex3

	err := errors.New(ex6)
	a := make([]ErrBaseInfo, 0, 10)
	bi := ErrBaseInfo{}.New(ex1, ex1ParentObj, ex2, ex3)
	se := SpecErr{}.InitializeBaseInfo(a, bi)
	x := se.New("", err, SpecErrTypeFATAL, ex5)


	if x.ErrMsg != ex6 {
		t.Error(fmt.Sprintf("Expected '%v' got", ex6), x.ErrMsg)
	}

	if x.BaseInfo.SourceFileName != ex1 {
		t.Error(fmt.Sprintf("Expected '%v' got", ex1), x.BaseInfo.SourceFileName)
	}

	if x.BaseInfo.ParentObjectName != ex1ParentObj {
		t.Errorf("Expected BaseInfo.ParentObjectName = '%v'. Instead, got ParentObjectName= '%v'", ex1ParentObj, x.BaseInfo.ParentObjectName)
	}

	if x.BaseInfo.FuncName != ex2 {
		t.Error(fmt.Sprintf("Expected '%v' got", ex2), x.BaseInfo.FuncName)
	}

	if x.errNo != ex7 {
		t.Error(fmt.Sprintf("Expected '%v' got", ex7), x.errNo)
	}

}

func TestSpecErr_New_002(t *testing.T) {
	ex1 := "errutil_test.go"
	ex1ParentObj := "ErrorObj"
	ex2 := "TestErrorUtility"
	ex3 := int64(10000)


	ex5 := int64(334)
	ex6 := "Test Error #1"
	ex6Msg := "Func Xray Overloaded!\n"
	ex7 := ex5 + ex3
	expectedMsg := ex6Msg + ex6
	err := errors.New(ex6)
	a := make([]ErrBaseInfo, 0, 10)
	bi := ErrBaseInfo{}.New(ex1, ex1ParentObj, ex2, ex3)
	se := SpecErr{}.InitializeBaseInfo(a, bi)
	x := se.New(ex6Msg, err, SpecErrTypeFATAL, ex5)


	if expectedMsg != x.ErrMsg  {
		t.Errorf("Expected '%v' got %v", expectedMsg, x.ErrMsg)
	}

	if x.BaseInfo.SourceFileName != ex1 {
		t.Error(fmt.Sprintf("Expected '%v' got", ex1), x.BaseInfo.SourceFileName)
	}

	if x.BaseInfo.ParentObjectName != ex1ParentObj {
		t.Errorf("Expected BaseInfo.ParentObjectName = '%v'. Instead, got ParentObjectName= '%v'", ex1ParentObj, x.BaseInfo.ParentObjectName)
	}

	if x.BaseInfo.FuncName != ex2 {
		t.Error(fmt.Sprintf("Expected '%v' got", ex2), x.BaseInfo.FuncName)
	}

	if x.errNo != ex7 {
		t.Error(fmt.Sprintf("Expected '%v' got", ex7), x.errNo)
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

	if se.BaseInfo.BaseErrorId != 0 {
		t.Error("Int64 BaseErrorId was uninitialized. Was expecting value of zero, got", se.BaseInfo.BaseErrorId)
	}

}

func TestInitializeParentInfo(t *testing.T) {

	bi := ErrBaseInfo{}

	x := bi.New("TestSourceFileName", "TestObjectName", "TestFuncName", 9000)
	y := bi.New("TestSrcFileName2", "TestObject2", "TestFuncName2", 14000)
	z := bi.New("TestSrcFileName3", "TestObject3", "TestFuncName3", 15000)

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

	if se.ParentInfo[1].ParentObjectName != "TestObject2" {
		t.Errorf("Expected 2nd Element ParentObjectName = 'TestObject2'. Instead, ParentObjectName= '%v'", se.ParentInfo[1].ParentObjectName)
	}

}

func TestAddSlicesParentInfo(t *testing.T) {
	var bi ErrBaseInfo
	x := bi.New("TestSourceFileName", "TestParentObj", "TestFuncName", 9000)
	y := bi.New("TestSrcFileName2", "TestParentObj2", "TestFuncName2", 14000)
	z := bi.New("TestSrcFileName3", "TestParentObj3", "TestFuncName3", 15000)

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

	if se.ParentInfo[1].ParentObjectName != "TestParentObj2" {
		t.Errorf("Expected 2nd Element 'TestParentObj2', Instead, got: '%v'", se.ParentInfo[1].ParentObjectName)
	}

}

func TestSetParentInfo(t *testing.T) {
	var bi ErrBaseInfo
	x := bi.New("TestSourceFileName", "TestParentObj","TestFuncName", 9000)
	y := bi.New("TestSrcFileName2", "TestParentObj2", "TestFuncName2", 14000)
	z := bi.New("TestSrcFileName3", "TestParentObj3", "TestFuncName3", 15000)

	a := make([]ErrBaseInfo, 0, 30)

	a = append(a, x, y, z)

	var se SpecErr

	se.ParentInfo = se.DeepCopyParentInfo(a)

	l := len(se.ParentInfo)

	if l != 3 {
		t.Error("Expected ParentInfo length of 3, go length of ", l)
	}

	if se.ParentInfo[1].FuncName != "TestFuncName2" {
		t.Errorf("Expected 2nd Element se.ParentInfo[1].FuncName = 'TestFuncName2'. Instead, got '%v'", se.ParentInfo[1].FuncName)
	}

	if se.ParentInfo[1].ParentObjectName != "TestParentObj2" {
		t.Errorf("Expected 2nd Element se.ParentInfo[1].ParentObjectName = 'TestParentObj2'. Instead, got '%v'", se.ParentInfo[1].ParentObjectName)
	}
}

func TestSetErrDetail(t *testing.T) {

	ex1 := "errutil_test.go"
	ex1ParentObj := "TestErrObj"
	ex2 := "TestErrorUtility"
	ex3 := int64(10000)


	ex5 := int64(338)
	ex6 := "Test Error #21"
	ex7 := ex5 + ex3

	err := errors.New(ex6)
	a := make([]ErrBaseInfo, 0, 10)
	bi := ErrBaseInfo{}.New(ex1, ex1ParentObj, ex2, ex3)

	x := SpecErr{}.InitializeBaseInfo(a, bi).New("",err, SpecErrTypeFATAL, ex5)

	if x.errNo != ex7 {
		t.Error(fmt.Sprintf("Expected Err No '%v', got", ex7), x.errNo)
	}

	if x.BaseInfo.SourceFileName != ex1 {
		t.Error(fmt.Sprintf("Expected Source File: '%v',got", ex1), x.BaseInfo.SourceFileName)
	}

	if x.BaseInfo.FuncName != ex2 {
		t.Error(fmt.Sprintf("Expected FuncName: '%v', got", ex2), x.BaseInfo.FuncName)
	}

	if x.BaseInfo.ParentObjectName != ex1ParentObj {
		t.Errorf("Expected x.BaseInfo.ParentObjectName: '%v'. Instead, got '%v'", ex1ParentObj, x.BaseInfo.ParentObjectName)
	}

}

func TestSpecErr_SignalNoErrors_01(t *testing.T) {

	s := SpecErr{}.SignalNoErrors()

	isErr := s.CheckIsSpecErr()

	if isErr {
		t.Errorf("Expected CheckIsSpecErr() to return false. Instead it retrned %v", isErr)
	}

	if s.ErrorMsgType != SpecErrTypeNOERRORSALLCLEAR {
		t.Errorf("Expected s.ErrorMsgType=='SpecErrTypeNOERRORSALLCLEAR'. Instead, got: '%v'", s.ErrorMsgType.String())
	}
}

func TestSpecErr_SignalSuccessfulCompletion_01(t *testing.T) {

	s := SpecErr{}.SignalSuccessfulCompletion()

	isErr := s.CheckIsSpecErr()

	if isErr {
		t.Errorf("Expected CheckIsSpecErr() to return false. Instead it retrned %v", isErr)
	}

	if s.ErrorMsgType != SpecErrTypeSUCCESSFULCOMPLETION {
		t.Errorf("Expected s.ErrorMsgType=='SpecErrTypeSUCCESSFULCOMPLETION'. Instead, got: '%v'", s.ErrorMsgType.String())
	}

}

func TestIsSpecErrYes(t *testing.T) {

	ex4 := int64(334)
	ex99 := "Test Error #1"

	err := errors.New(ex99)

	x := SpecErr{}.New("", err, SpecErrTypeFATAL, ex4)

	isErr := x.CheckIsSpecErr()

	if !isErr {
		t.Errorf("Expected x.CheckIsSpecErr() to return 'true'. Instead it returned %v", isErr)
	}

}

func TestSetNoErr(t *testing.T) {
	x := SpecErr{}.SignalNoErrors()

	if x.IsErr {
		t.Errorf("Expected IsErr= 'false'. Instead IsErr= '%v'", x.IsErr)
	}
}

func TestQuickInitialize(t *testing.T) {

	ex2 := "Error Msg X"
	err := errors.New(ex2)
	ex4 := int64(499)

	x := SpecErr{}.New("",err, SpecErrTypeERROR, ex4)

	if x.errNo != ex4 {
		t.Error(fmt.Sprintf("Expected errNo: '%v', got", ex4), x.errNo)
	}

	if x.ErrMsg != ex2 {
		t.Error(fmt.Sprintf("Expected Error Msg '%v', got,", ex2), x.ErrMsg)
	}

	if x.IsPanic == true {
		t.Errorf("Expected IsPanic == '%v'. Instead IsPanic== '%v' ", false, x.IsPanic)
	}
}

func TestFullInitialize(t *testing.T) {

	bi := ErrBaseInfo{}

	f := bi.New("TestSourceFileName", "TestParentObj", "TestFuncName", 9000)
	g := bi.New("TestSrcFileName2", "TestParentObj2", "TestFuncName2", 14000)
	h := bi.New("TestSrcFileName3", "TestParentObj3", "TestFuncName3", 15000)

	ex1 := make([]ErrBaseInfo, 0, 10)
	ex1 = append(ex1, f, g, h)

	ex21 := "TestSrcFileName99"
	ex21ParentObj := "TestOjb99"
	ex22 := "TestFuncName99"
	ex23 := int64(16000)
	ex2 := bi.New(ex21, ex21ParentObj, ex22, ex23)

	ex4 := "Error Msg 99"
	err := errors.New(ex4)
	ex5 := false
	ex6 := int64(22)
	ex7 := int64(16022)

	x := SpecErr{}.Initialize(ex1, ex2, "", err, SpecErrTypeERROR, ex6)

	pl := len(x.ParentInfo)

	if pl != 3 {
		t.Error("Expected ParentInfo length == 3, got", pl)
	}

	if x.ErrMsg != ex4 {
		t.Error(fmt.Sprintf("Expected ErrMsg '%v', got", ex4), x.ErrMsg)
	}

	if x.IsErr == false {
		t.Error(fmt.Sprintf("Expected IsErr == '%v', got", true), x.IsErr)
	}

	if x.IsPanic != ex5 {
		t.Errorf("Expected IsPanic == '%v'. Instead IsPannic=='%v'", ex5, x.IsPanic)
	}

	if x.errNo != ex7 {
		t.Error(fmt.Sprintf("Expected errNo '%v', got", ex7), x.errNo)
	}

	if x.BaseInfo.SourceFileName != ex21 {
		t.Error(fmt.Sprintf("Expected SourceFileName '%v', got", ex21), x.BaseInfo.SourceFileName)
	}

	if x.BaseInfo.ParentObjectName != ex21ParentObj {
		t.Errorf("Expected x.BaseInfo.ParentObjectName = '%v'. Instead, got '%v'", ex21ParentObj, x.BaseInfo.ParentObjectName)
	}

	if x.BaseInfo.FuncName != ex22 {
		t.Error(fmt.Sprintf("Expected Function Name '%v', got", ex22), x.BaseInfo.FuncName)
	}

	if x.BaseInfo.BaseErrorId != ex23 {
		t.Error(fmt.Sprintf("Expected Base Error ID '%v', got", ex23), x.BaseInfo.BaseErrorId)
	}

}

func TestBlankInitialize(t *testing.T) {

	ex2 := "Error Msg 99"
	err := errors.New(ex2)
	ex3 := false
	ex4 := int64(22)

	x := SpecErr{}.Initialize(blankParentInfo, blankErrBaseInfo, "", err, SpecErrTypeERROR, ex4)

	if x.ErrMsg != ex2 {
		t.Error(fmt.Sprintf("Expected ErrMsg '%v', got", ex2), x.ErrMsg)
	}

	if x.IsErr == false {
		t.Error(fmt.Sprintf("Expected IsErr == '%v', got", true), x.IsErr)
	}

	if x.IsPanic != ex3 {
		t.Error(fmt.Sprintf("Expected IsPanic == '%v', got", ex3), x.IsPanic)
	}

	if x.errNo != ex4 {
		t.Error(fmt.Sprintf("Expected errNo '%v', got", ex4), x.errNo)
	}

	if x.BaseInfo.SourceFileName != "" {
		t.Error("Expected BaseInfo Source File Name == '', got", x.BaseInfo.SourceFileName)
	}

	if x.BaseInfo.FuncName != "" {
		t.Error("Expected BaseInfo FuncName == '', got", x.BaseInfo.FuncName)
	}

	if x.BaseInfo.BaseErrorId != int64(0) {
		t.Error("Expected BaseInfo BaseErrorId == 'Zero', got", x.BaseInfo.BaseErrorId)
	}

}

func TestAddBaseInfoToParent(t *testing.T) {

	bi := ErrBaseInfo{}

	f := bi.New("TestSourceFileName", "TestObject", "TestFuncName", 9000)
	g := bi.New("TestSrcFileName2", "TestObject2", "TestFuncName2", 14000)
	h := bi.New("TestSrcFileName3", "TestObject3", "TestFuncName3", 15000)

	ex1 := make([]ErrBaseInfo, 0, 10)
	ex1 = append(ex1, f, g, h)

	ex21 := "TestSrcFileName99"
	ex21ParentObj := "TestObject99"
	ex22 := "TestFuncName99"
	ex23 := int64(16000)
	ex2 := bi.New(ex21, ex21ParentObj, ex22, ex23)

	ex4 := "Error Msg 99"
	err := errors.New(ex4)
	ex5 := false
	ex6 := int64(22)

	x := SpecErr{}.Initialize(ex1, ex2, "", err, SpecErrTypeERROR, ex6)

	pl := len(x.ParentInfo)

	if pl != 3 {
		t.Error("Expected Parent Info to contain 3-elements. Actual number of elements was ", pl)
	}

	p2 := x.AddBaseToParentInfo()

	pl2 := len(p2)

	if pl2 != 4 {
		t.Error("Expected New Parent Info to contain 4-elements. Actual number of elements was ", pl2)
	}

	if x.IsPanic != ex5 {
		t.Errorf("Expected IsPanic = '%v'. Instead, IsPanic = '%v'", ex5, x.IsPanic)
	}

}

func TestSpecErr_CheckIsSpecErr(t *testing.T) {

	bi := ErrBaseInfo{}

	f := bi.New("TestSourceFileName", "TestObject", "TestFuncName", 9000)
	g := bi.New("TestSrcFileName2", "TestObject2", "TestFuncName2", 14000)
	h := bi.New("TestSrcFileName3", "TestObject3", "TestFuncName3", 15000)

	ex1 := make([]ErrBaseInfo, 0, 10)
	ex1 = append(ex1, f, g, h)

	ex21 := "TestSrcFileName99"
	ex21ParentObj := "TestObject99"
	ex22 := "TestFuncName99"
	ex23 := int64(16000)
	ex2 := bi.New(ex21, ex21ParentObj, ex22, ex23)

	ex4 := "Error Msg 99"
	err := errors.New(ex4)
	ex6 := int64(22)

	x := SpecErr{}.Initialize(ex1, ex2, "", err, SpecErrTypeERROR, ex6)

	result := x.CheckIsSpecErr()

	if result == false {
		t.Errorf("Expected x.CheckIsSpecErr()== 'true'. Instead, x.CheckIsSpecErr()== '%v'", result)
	}

}

func TestSpecErr_CheckIsSpecErrPanic(t *testing.T) {
	bi := ErrBaseInfo{}

	f := bi.New("TestSourceFileName", "TestObject", "TestFuncName", 9000)
	g := bi.New("TestSrcFileName2", "TestObject2", "TestFuncName2", 14000)
	h := bi.New("TestSrcFileName3","TestObject3", "TestFuncName3", 15000)

	ex1 := make([]ErrBaseInfo, 0, 10)
	ex1 = append(ex1, f, g, h)

	ex21 := "TestSrcFileName99"
	ex21ParentObj := "TestObject99"
	ex22 := "TestFuncName99"
	ex23 := int64(16000)
	ex2 := bi.New(ex21, ex21ParentObj, ex22, ex23)

	ex4 := "Error Msg 99"
	err := errors.New(ex4)
	ex6 := int64(22)

	x := SpecErr{}.Initialize(ex1, ex2, "", err, SpecErrTypeERROR, ex6)

	x.IsPanic = true

	result := x.CheckIsSpecErrPanic()

	if result != true {
		t.Errorf("Expected x.CheckIsSpecErrPanic()== 'true' .  Instead, x.CheckIsSpecErrPanic()== '%v'", result)
	}

}

func TestSpecErr_NewErrorMsgString_01(t *testing.T) {

	errMsg := "This is the Error Msg!"
	isPanic := false
	errNo := int64(22)

	s :=  SpecErr{}.NewErrorMsgString(errMsg, SpecErrTypeERROR, errNo)

	if s.ErrMsg != errMsg {
		t.Errorf("Expected Error Message= '%v'.  Instead, Error Message = '%v'", errMsg, s.ErrMsg)
	}

	if isPanic != s.IsPanic {
		t.Errorf("Expected IsPanic = '%v'.  Instead, IsPanic = '%v'", isPanic, s.IsPanic)
	}

}

func TestSpecErr_NewErrorMsgString_02(t *testing.T) {
	errMsg := "This is the Error Msg!"
	isPanic := false
	errNo := int64(0)

	s :=  SpecErr{}.NewErrorMsgString(errMsg, SpecErrTypeERROR, errNo)

	if s.ErrMsg != errMsg {
		t.Errorf("Expected Error Message= '%v'.  Instead, Error Message = '%v'", errMsg, s.ErrMsg)
	}

	errResult := s.Error()

	hasErrNo := strings.Contains(errResult, "errNo:")

	if true == hasErrNo {
		t.Error("Due to Error Number=Zero, expected error message WITHOUT Error Number. Instead, Error Number was included in Error Message")
	}

	if isPanic != s.IsPanic {
		t.Errorf("Expected IsPanic = '%v'. Instead, IsPanic = '%v'.", isPanic, s.IsPanic)
	}

}


func TestSpecErr_ConfigureParentInfoFromParentSpecErr01(t *testing.T) {

	bi := ErrBaseInfo{}

	baseInfo := bi.New("TestSrcFileName6", "TestObject6", "TestFuncName6", 6000)

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()

	e := errors.New("This is the error message")

	se := SpecErr{}.Initialize(parentInfo, baseInfo, "", e, SpecErrTypeFATAL, 553 )

	baseInfo2 := bi.New("TestSrcFileName7", "TestObject7", "TestFuncName7", 7000)

	se2 := SpecErr{}

	se2.ConfigureParentInfoFromParentSpecErr(se)

	se2.ConfigureBaseInfo(baseInfo2)

	e2 := errors.New("This is Error Message # 2")

	se2.SetError(e2, SpecErrTypeERROR, 902)

	if len(se2.ParentInfo) != 6 {

		t.Errorf("Expected length of se2.ParentInfo array == 6. Instead, array length == '%v'", len(se2.ParentInfo))

	}

	if se2.ParentInfo[5].ParentObjectName != "TestObject6" {
		t.Errorf("Expected se2.ParentInfo[5].ParentObjectName != 'TestObject6'. Instead ObjectName='%v'", se2.ParentInfo[5].ParentObjectName)
	}

	if se2.errNo != 7902 {
		t.Errorf("Expected se2.errNo== 7902.  Instead, se2.errNo== '%v'", se2.errNo)
	}

	msg := se2.Error()

	if !strings.Contains(msg,"This is Error Message # 2"){
		t.Error("Expected final error message to contain 'This is Error Message # 2'.  Instead it did NOT!")
	}

}

func TestSpecErr_InitializeBaseInfoWithSpecErr01(t *testing.T) {

	bi := ErrBaseInfo{}

	baseInfo := bi.New("TestSrcFileName6", "TestObject6", "TestFuncName6", 6000)

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()

	e := errors.New("This is the error message")

	se := SpecErr{}.Initialize(parentInfo, baseInfo,"", e, SpecErrTypeFATAL, 553 )

	baseInfo2 := bi.New("TestSrcFileName7", "TestObject7", "TestFuncName7", 7000)

	se2 := SpecErr{}.InitializeBaseInfoWithSpecErr(se, baseInfo2)

	e2 := errors.New("This is Error Message # 2")

	se2.SetError(e2, SpecErrTypeERROR, 902)

	if len(se2.ParentInfo) != 6 {

		t.Errorf("Expected length of se2.ParentInfo array == 6. Instead, array length == '%v'", len(se2.ParentInfo))

	}

	if se2.ParentInfo[5].ParentObjectName != "TestObject6" {
		t.Errorf("Expected se2.ParentInfo[5].ParentObjectName != 'TestObject6'. Instead ObjectName='%v'", se2.ParentInfo[5].ParentObjectName)
	}

	if se2.errNo != 7902 {
		t.Errorf("Expected se2.errNo== 7902.  Instead, se2.errNo== '%v'", se2.errNo)
	}

	msg := se2.Error()

	if !strings.Contains(msg,"This is Error Message # 2"){
		t.Error("Expected final error message to contain 'This is Error Message # 2'.  Instead it did NOT!")
	}

	if se2.ErrorMsgType != SpecErrTypeERROR {
		t.Errorf("Expected 'SpecErrTypeERROR'. Instead, got '%v'", se2.ErrorMsgType)
	}

}

func TestSpecErr_NewInfoMsgString01(t *testing.T) {

	bi := ErrBaseInfo{}

	baseInfo := bi.New("TestSrcFileName6", "TestObject6", "TestFuncName6", 6000)

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()

	e := errors.New("This is an error message #1")

	se := SpecErr{}.Initialize(parentInfo, baseInfo, "", e, SpecErrTypeFATAL, 553 )

	baseInfo2 := bi.New("TestSrcFileName7", "TestObject7", "TestFuncName7", 7000)

	se2 := SpecErr{}.InitializeBaseInfoWithSpecErr(se, baseInfo2)

	iMsg := "This is Information Message # 2"

	se2.SetErrorWithMessage(iMsg, SpecErrTypeINFO, 902)

	if len(se2.ParentInfo) != 6 {

		t.Errorf("Expected length of se2.ParentInfo array == 6. Instead, array length == '%v'", len(se2.ParentInfo))

	}

	if se2.ParentInfo[5].ParentObjectName != "TestObject6" {
		t.Errorf("Expected se2.ParentInfo[5].ParentObjectName != 'TestObject6'. Instead ObjectName='%v'", se2.ParentInfo[5].ParentObjectName)
	}

	if se2.errNo != 7902 {
		t.Errorf("Expected se2.errNo== 7902.  Instead, se2.errNo== '%v'", se2.errNo)
	}

	msg := se2.Error()

	if !strings.Contains(msg,iMsg){
		t.Errorf("Expected final error message to contain '%v'.  Instead it did NOT!", iMsg)
	}

	if se2.ErrorMsgType != SpecErrTypeINFO {
		t.Errorf("Expected 'SpecErrTypeINFO'. Instead, got '%v'", se2.ErrorMsgType)
	}

}

func TestSpecErr_NewWarningMsgString01(t *testing.T) {

	bi := ErrBaseInfo{}

	baseInfo := bi.New("TestSrcFileName6", "TestObject6", "TestFuncName6", 6000)

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()

	e := errors.New("This is an error message #1")

	se := SpecErr{}.Initialize(parentInfo, baseInfo, "", e, SpecErrTypeFATAL, 553 )

	baseInfo2 := bi.New("TestSrcFileName7", "TestObject7", "TestFuncName7", 7000)

	se2 := SpecErr{}.InitializeBaseInfoWithSpecErr(se, baseInfo2)

	iMsg := "This is Warning Message # 2"

	se2.SetErrorWithMessage(iMsg, SpecErrTypeWARNING, 902)

	if len(se2.ParentInfo) != 6 {

		t.Errorf("Expected length of se2.ParentInfo array == 6. Instead, array length == '%v'", len(se2.ParentInfo))

	}

	if se2.ParentInfo[5].ParentObjectName != "TestObject6" {
		t.Errorf("Expected se2.ParentInfo[5].ParentObjectName != 'TestObject6'. Instead ObjectName='%v'", se2.ParentInfo[5].ParentObjectName)
	}

	if se2.errNo != 7902 {
		t.Errorf("Expected se2.errNo== 7902.  Instead, se2.errNo== '%v'", se2.errNo)
	}

	msg := se2.Error()

	if !strings.Contains(msg,iMsg){
		t.Errorf("Expected final error message to contain '%v'.  Instead it did NOT!", iMsg)
	}

	if se2.ErrorMsgType != SpecErrTypeWARNING {
		t.Errorf("Expected 'SpecErrTypeINFO'. Instead, got '%v'", se2.ErrorMsgType)
	}

}

func TestSpecErr_SetStdError_01(t *testing.T) {
	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is error msg 1"
	xType := SpecErrTypeERROR

	se := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	se.SetStdError(xMsg, 822)

	if len(se.ParentInfo) != 5 {
		t.Errorf("Expected length of se.ParentInfo array == 5. Instead, array length == '%v'", len(se.ParentInfo))
	}

	if se.ParentInfo[4].ParentObjectName != "TestObject5" {
		t.Errorf("Expected se2.ParentInfo[5].ParentObjectName != 'TestObject5'. Instead ObjectName='%v'", se.ParentInfo[4].ParentObjectName)
	}

	if se.errNo != 6822 {
		t.Errorf("Expected se.errNo== 6822.  Instead, se.errNo== '%v'", se.errNo)
	}

	msg := se.Error()

	if !strings.Contains(msg,xMsg){
		t.Errorf("Expected final error message to contain '%v'.  Instead it did NOT!", xMsg)
	}

	if se.ErrorMsgType != xType {
		t.Errorf("Expected SpeErrType= '%v'. Instead, got '%v'",xType, se.ErrorMsgType)
	}
}

func TestSpecErr_SetFatalError_01(t *testing.T) {
	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is fatal error msg 1"
	xType := SpecErrTypeFATAL

	se := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	se.SetFatalError(xMsg, 822)

	if len(se.ParentInfo) != 5 {
		t.Errorf("Expected length of se.ParentInfo array == 5. Instead, array length == '%v'", len(se.ParentInfo))
	}

	if se.ParentInfo[4].ParentObjectName != "TestObject5" {
		t.Errorf("Expected se2.ParentInfo[5].ParentObjectName != 'TestObject5'. Instead ObjectName='%v'", se.ParentInfo[4].ParentObjectName)
	}

	if se.errNo != 6822 {
		t.Errorf("Expected se.errNo== 6822.  Instead, se.errNo== '%v'", se.errNo)
	}

	msg := se.Error()

	if !strings.Contains(msg,xMsg){
		t.Errorf("Expected final error message to contain '%v'.  Instead it did NOT!", xMsg)
	}

	str := se.String()

	if str != msg {
		t.Errorf("Compared Error() to String(). Results do not match. Error()= '%v' - String() = '%v'", msg, str)
	}

	if se.ErrorMsgType != xType {
		t.Errorf("Expected SpeErrType= '%v'. Instead, got '%v'",xType, se.ErrorMsgType)
	}
}

func TestSpecErr_SetInfoMsg_01(t *testing.T) {
	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is information msg 1"
	xType := SpecErrTypeINFO

	se := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	se.SetInfoMessage(xMsg, 822)

	if len(se.ParentInfo) != 5 {
		t.Errorf("Expected length of se.ParentInfo array == 5. Instead, array length == '%v'", len(se.ParentInfo))
	}

	if se.ParentInfo[4].ParentObjectName != "TestObject5" {
		t.Errorf("Expected se2.ParentInfo[5].ParentObjectName != 'TestObject5'. Instead ObjectName='%v'", se.ParentInfo[4].ParentObjectName)
	}

	if se.errNo != 6822 {
		t.Errorf("Expected se.errNo== 6822.  Instead, se.errNo== '%v'", se.errNo)
	}

	msg := se.Error()

	if !strings.Contains(msg,xMsg){
		t.Errorf("Expected final error message to contain '%v'.  Instead it did NOT!", xMsg)
	}

	str := se.String()

	if str != msg {
		t.Errorf("Compared Error() to String(). Results do not match. Error()= '%v' - String() = '%v'", msg, str)
	}

	if se.ErrorMsgType != xType {
		t.Errorf("Expected SpeErrType= '%v'. Instead, got '%v'",xType, se.ErrorMsgType)
	}
}

func TestSpecErr_SetWarningMsg_01(t *testing.T) {
	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is warning msg 1"
	xType := SpecErrTypeWARNING

	se := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	se.SetWarningMessage(xMsg, 822)

	if len(se.ParentInfo) != 5 {
		t.Errorf("Expected length of se.ParentInfo array == 5. Instead, array length == '%v'", len(se.ParentInfo))
	}

	if se.ParentInfo[4].ParentObjectName != "TestObject5" {
		t.Errorf("Expected se2.ParentInfo[5].ParentObjectName != 'TestObject5'. Instead ObjectName='%v'", se.ParentInfo[4].ParentObjectName)
	}

	if se.errNo != 6822 {
		t.Errorf("Expected se.errNo== 6822.  Instead, se.errNo== '%v'", se.errNo)
	}

	msg := se.Error()

	if !strings.Contains(msg,xMsg){
		t.Errorf("Expected final error message to contain '%v'.  Instead it did NOT!", xMsg)
	}

	str := se.String()

	if str != msg {
		t.Errorf("Compared Error() to String(). Results do not match. Error()= '%v' - String() = '%v'", msg, str)
	}

	if se.ErrorMsgType != xType {
		t.Errorf("Expected SpeErrType= '%v'. Instead, got '%v'",xType, se.ErrorMsgType)
	}
}

func TestSpecErr_CopyOut_01(t *testing.T) {

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is information msg 1"

	s := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	s.SetInfoMessage(xMsg, 822)

	s2 := s.CopyOut()

	if !s.Equal(&s2) {
		t.Error("Expected after Copy Out s==s2. However, it did NOT!")
	}

}

func TestSpecErr_CopyIn_01(t *testing.T) {

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is information msg 1"

	s := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	s.SetInfoMessage(xMsg, 822)

	s2 := SpecErr{}

	s2.CopyIn(s)

	if !s.Equal(&s2) {
		t.Error("Expected after CopyIn() s==s2. However, it did NOT!")
	}

}

func TestSpecErr_SetDebugMsg_001(t *testing.T) {

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is DEBUG msg #1"
	xErrId := int64(822)
	xErrNo := int64(6822)
	s := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	s.SetDebugMessage(xMsg, xErrId)

	actualMsg := s.String()

	if !strings.Contains(actualMsg, xMsg) {
		t.Errorf("Expected Debug message to contain text '%v'. It did NOT!", xMsg)
	}

	actualErrId := s.GetErrorId()

	if xErrId != actualErrId {
		t.Errorf("Expected ErrId= '%v'. Instead, ErrId= '%v'", xErrId, actualErrId)
	}

	actualErrNo := s.GetErrorNumber()

	if xErrNo != actualErrNo {
		t.Errorf("Expected ErrNo= '%v'.  Instead, ErrNo= '%v'.", xErrNo, actualErrNo)
	}

}

func TestSpecErr_Equal_01(t *testing.T) {

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is information msg 1"

	s := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	s.SetInfoMessage(xMsg, 822)

	s2 := s.CopyOut()

	s2.ParentInfo[1].FuncName = "XXXX"

	if s.Equal(&s2) {
		t.Error("Expected after changes to s2, s!=s2. However, s==s2!")
	}

}

func TestSpecErr_Equal_02(t *testing.T) {

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is information msg 1"

	s := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	s.SetInfoMessage(xMsg, 822)

	s2 := s.CopyOut()

	s2.ParentInfo = s2.ParentInfo[0:2]

	if s.Equal(&s2) {
		t.Error("Expected after changes to s2, s!=s2. However, s==s2!")
	}

}

func TestSpecErr_Equal_03(t *testing.T) {

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is information msg 1"

	s := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	s.SetInfoMessage(xMsg, 822)

	s2 := s.CopyOut()

	s2.BaseInfo.BaseErrorId = 90

	if s.Equal(&s2) {
		t.Error("Expected after changes to s2, s!=s2. However, s==s2!")
	}

}

func TestSpecErr_Equal_04(t *testing.T) {

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is information msg 1"

	s := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	s.SetInfoMessage(xMsg, 822)

	s2 := s.CopyOut()

	s2.ErrorLocalTimeZone = "UTC"

	if s.Equal(&s2) {
		t.Error("Expected after changes to s2, s!=s2. However, s==s2!")
	}

}

func TestSpecErr_Equal_05(t *testing.T) {

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is information msg 1"

	s := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	s.SetInfoMessage(xMsg, 822)

	s2 := s.CopyOut()

	s2.errNo = 47

	if s.Equal(&s2) {
		t.Error("Expected after changes to s2, s!=s2. However, s==s2!")
	}

}

func TestSpecErr_Equal_06(t *testing.T) {

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is information msg 1"

	s := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	s.SetInfoMessage(xMsg, 822)

	s2 := s.CopyOut()

	s2.ErrorMsgType = SpecErrTypeFATAL

	if s.Equal(&s2) {
		t.Error("Expected after changes to s2, s!=s2. However, s==s2!")
	}

}

func TestSpecErr_Equal_07(t *testing.T) {

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is information msg 1"

	s := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	s.SetInfoMessage(xMsg, 822)

	s2 := s.CopyOut()

	s2.IsErr = true

	if s.Equal(&s2) {
		t.Error("Expected after changes to s2, s!=s2. However, s==s2!")
	}

}

func TestSpecErr_Equal_08(t *testing.T) {

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is information msg 1"

	s := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	s.SetInfoMessage(xMsg, 822)

	s2 := s.CopyOut()

	s2.IsPanic = true

	if s.Equal(&s2) {
		t.Error("Expected after changes to s2, s!=s2. However, s==s2!")
	}

}

func TestSpecErr_Equal_09(t *testing.T) {

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is information msg 1"

	s := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	s.SetInfoMessage(xMsg, 822)

	s2 := s.CopyOut()

	s2.BaseInfo.ParentObjectName = "X25"

	if s.Equal(&s2) {
		t.Error("Expected after changes to s2, s!=s2. However, s==s2!")
	}

}

func TestSpecErr_Equal_10(t *testing.T) {

	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is information msg 1"

	s := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	s.SetInfoMessage(xMsg, 822)

	s2 := s.CopyOut()

	if !s.Equal(&s2) {
		t.Error("Expected copying s to s2 , s==s2. However, s!=s2!")
	}

}


func TestSpecErr_Empty_01(t *testing.T) {
	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is warning msg 1"

	se := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	se.SetWarningMessage(xMsg, 822)

	se.Empty()

	if len(se.ParentInfo) != 0 {
		t.Errorf("Expected length of se.ParentInfo array == 0. Instead, array length == '%v'", len(se.ParentInfo))
	}

	if se.errNo != 0 {
		t.Errorf("Expected se.errNo== 0.  Instead, se.errNo== '%v'", se.errNo)
	}

	msg := se.Error()

	if msg != ""{
		t.Errorf("Expected error message = 'No Errors - No Messages'.  Instead it did NOT! msg= '%v'", msg)
	}


	str := se.String()

	if str != msg {
		t.Errorf("Compared Error() to String(). Results do not match. Error()= '%v' - String() = '%v'", msg, str)
	}

	if se.ErrorMsgType != SpecErrTypeNOERRORSALLCLEAR {
		t.Errorf("Expected SpeErrType= 'SpecErrTypeNOERRORSALLCLEAR'. Instead, got '%v'", se.ErrorMsgType)
	}
}

func TestSpecErr_EmptyMsgData_01(t *testing.T) {
	parentInfo := testCreateSpecErrParentBaseInfo5Elements()
	currBaseInfo := testCreateSpecErrBaseInfoObject()
	xMsg := "This is warning msg 1"

	se := SpecErr{}.InitializeBaseInfo(parentInfo, currBaseInfo)

	se.SetWarningMessage(xMsg, 822)

	se.EmptyMsgData()

	if len(se.ParentInfo) != 5 {
		t.Errorf("Expected length of se.ParentInfo array == 5. Instead, array length == '%v'", len(se.ParentInfo))
	}

	if se.ParentInfo[4].ParentObjectName != "TestObject5" {
		t.Errorf("Expected se2.ParentInfo[5].ParentObjectName != 'TestObject5'. Instead ObjectName='%v'", se.ParentInfo[4].ParentObjectName)
	}

	if se.BaseInfo.FuncName != "TestFuncName6" {
		t.Errorf("Expected se.BaseInfo.FuncName == 'TestFuncName6'. Instead FuncName= '%v'",se.BaseInfo.FuncName)
	}

	if se.errNo != 0 {
		t.Errorf("Expected se.errNo== 0.  Instead, se.errNo== '%v'", se.errNo)
	}

	msg := se.Error()

	if msg != ""{
		t.Errorf("Expected error message = 'No Errors - No Messages'.  Instead it did NOT! msg= '%v'", msg)
	}


	str := se.String()

	if str != msg {
		t.Errorf("Compared Error() to String(). Results do not match. Error()= '%v' - String() = '%v'", msg, str)
	}

	if se.ErrorMsgType != SpecErrTypeNOERRORSALLCLEAR {
		t.Errorf("Expected SpeErrType= 'SpecErrTypeNOERRORSALLCLEAR'. Instead, got '%v'", se.ErrorMsgType)
	}
}


func testCreateSpecErrParentBaseInfo5Elements() []ErrBaseInfo {
	parentBaseInfo := make([]ErrBaseInfo, 0, 10)
	bi := ErrBaseInfo{}

	a := bi.New("TestSrcFileName1", "TestObject1", "TestFuncName1", 1000)
	b := bi.New("TestSrcFileName2", "TestObject2", "TestFuncName2", 2000)
	c := bi.New("TestSrcFileName3","TestObject3", "TestFuncName3", 3000)
	d := bi.New("TestSrcFileName4","TestObject4", "TestFuncName4", 4000)
	e := bi.New("TestSrcFileName5","TestObject5", "TestFuncName5", 5000)

	parentBaseInfo = append(parentBaseInfo, a, b, c, d, e)


	return parentBaseInfo
}


func testCreateSpecErrBaseInfoObject() ErrBaseInfo {
	bi := ErrBaseInfo{}

	a := bi.New("TestSrcFileName6", "TestObject6", "TestFuncName6", 6000)

	return a
}