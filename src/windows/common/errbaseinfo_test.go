package common

import (
	"fmt"
	"testing"
)

func TestUninitializedErrBaseInfo(t *testing.T) {
	var se ErrBaseInfo

	if se.SourceFileName != "" {
		t.Error("String SourceFileName was uninitialized. Was expecting empty string, got", se.SourceFileName)
	}

	if se.FuncName != "" {
		t.Error("String FuncName was uninitialized. Was expecting empty string, got", se.FuncName)
	}

	if se.BaseErrorId != 0 {
		t.Error("Int64 BaseErrorId was uninitialized. Was expecting value of zero, got", se.BaseErrorId)
	}

}

// Creates new BaseInfo Structure
func TestNewBaseInfo(t *testing.T) {
	ex1 := "TestSourceFileName"
	ex1ParentObj := "TestObject"
	ex2 := "TestFuncName"
	ex3 := int64(9000)

	x := ErrBaseInfo{}.New(ex1, ex1ParentObj, ex2, ex3)

	if x.SourceFileName != ex1 {
		t.Error(fmt.Sprintf("Expected %v got,", ex1), x.SourceFileName)
	}

	if x.ParentObjectName != ex1ParentObj {
		t.Errorf("Expected x.ParentObjectName= %v. Instead, got: '%v'", ex1ParentObj, x.ParentObjectName)
	}

	if x.FuncName != ex2 {
		t.Error(fmt.Sprintf("Expected %v got,", ex2), x.FuncName)
	}

	if x.BaseErrorId != ex3 {
		t.Error(fmt.Sprintf("Expected %v got,", ex3), x.BaseErrorId)
	}

}

// Changes Function Name Only
func TestNewFunc(t *testing.T) {
	ex1 := "TestSourceFileName"
	ex1ParentObj := "TestObject"
	ex2 := "TestFuncName"
	ex3 := int64(9000)
	ex4 := "NewFuncName"

	x := ErrBaseInfo{}.New(ex1, ex1ParentObj, ex2, ex3)

	if x.FuncName != ex2 {
		t.Error(fmt.Sprintf("Expected %v got,", ex2), x.FuncName)
	}

	y := x.NewFunc(ex4)

	if y.SourceFileName != ex1 {
		t.Error(fmt.Sprintf("Expected %v got,", ex1), y.SourceFileName)
	}

	if y.ParentObjectName != ex1ParentObj {
		t.Errorf("Expected y.ParentObjectName= %v. Instead, got: %v", ex1ParentObj, y.ParentObjectName)
	}

	if y.FuncName != ex4 {
		t.Error(fmt.Sprintf("Expected %v got,", ex4), y.FuncName)
	}

	if y.BaseErrorId != ex3 {
		t.Error(fmt.Sprintf("Expected %v got,", ex3), y.BaseErrorId)
	}
}

func TestBaseInfoDeepCopy(t *testing.T) {
	ex1 := "TestSourceFileName"
	ex1ParentObj:= "TestObject"
	ex2 := "TestFuncName"
	ex3 := int64(9000)

	x := ErrBaseInfo{}.New(ex1, ex1ParentObj, ex2, ex3)

	y := x.NewBaseInfo()

	if y.SourceFileName != ex1 {
		t.Error(fmt.Sprintf("Expected %v got,", ex1), y.SourceFileName)
	}

	if y.ParentObjectName != ex1ParentObj {
		t.Errorf("Expected y.ParentObjectName='%v'. Instead, got: '%v'", ex1, y.ParentObjectName)
	}

	if y.FuncName != ex2 {
		t.Error(fmt.Sprintf("Expected %v got,", ex2), y.FuncName)
	}

	if y.BaseErrorId != ex3 {
		t.Error(fmt.Sprintf("Expected %v got,", ex3), y.BaseErrorId)
	}

}

func TestGetSpecErrFromBaseInfo(t *testing.T) {
	ex1 := "TestSourceFileName"
	ex1ParentObj:= "TestObject"
	ex2 := "TestFuncName"
	ex3 := int64(9000)

	x := ErrBaseInfo{}.New(ex1, ex1ParentObj, ex2, ex3)

	se := x.GetBaseSpecErr()

	if se.BaseInfo.SourceFileName != "TestSourceFileName" {
		t.Error("Expected BaseInfo.SourceFileName 'TestSourceFileName', got ", se.BaseInfo.SourceFileName)
	}

	if se.BaseInfo.ParentObjectName != ex1ParentObj {
		t.Errorf("Expected BaseInfo.ParentObjectName= '%v'.  Instead, got: '%v' ",ex1ParentObj, se.BaseInfo.ParentObjectName)
	}

	if se.BaseInfo.FuncName != "TestFuncName" {
		t.Error("Expected BaseInfo.FuncName 'TestFuncName', got ", se.BaseInfo.FuncName)
	}

	if se.BaseInfo.BaseErrorId != 9000 {
		t.Error("Expected BaseInfo.BaseErrorId '9000', got ", se.BaseInfo.BaseErrorId)
	}

}

func TestGetNewParentInfo(t *testing.T) {
	ex1 := "test11.go"
	ex1ParentObj:="TestObject"
	ex2 := "test11"
	ex3 := int64(1000)

	parent := ErrBaseInfo{}.GetNewParentInfo(ex1, ex1ParentObj, ex2, ex3)

	l := len(parent)

	if l != 1 {
		t.Error("Expected 1 element. Slice length is:", l)
	}

	if parent[0].SourceFileName != ex1 {
		t.Error(fmt.Sprintf("Expected %v, got", ex1), parent[0].SourceFileName)
	}

	if parent[0].ParentObjectName != ex1ParentObj {
		t.Errorf("Expected parent[0].ParentObjectName='%v'. Instead, got: '%v'", ex1ParentObj, parent[0].ParentObjectName)
	}

	if parent[0].FuncName != ex2 {
		t.Error(fmt.Sprintf("Expected %v, got", ex2), parent[0].FuncName)
	}

	if parent[0].BaseErrorId != ex3 {
		t.Error(fmt.Sprintf("Expected %v, got", ex3), parent[0].BaseErrorId)
	}
}

func TestErrBaseInfo_Equal_01(t *testing.T) {

	e1 := testCreateErrBaseInfoObject()

	e2 := e1.DeepCopyBaseInfo()


	if !e1.Equal(&e2) {
		t.Error("After copying e1 to e2, e1 IS NOT EQUAL to e2!")
	}

}

func TestErrBaseInfo_Equal_02(t *testing.T) {

	e1 := testCreateErrBaseInfoObject()

	e2 := e1.DeepCopyBaseInfo()

	e2.FuncName = "xx"

	if e1.Equal(&e2) {
		t.Error("After copying e1 to e2 and changing a field value, e1 IS STILL EQUAL to e2!")
	}

}

func TestErrBaseInfo_Equal_03(t *testing.T) {

	e1 := testCreateErrBaseInfoObject()

	e2 := e1.DeepCopyBaseInfo()

	e2.SourceFileName = "xx"

	if e1.Equal(&e2) {
		t.Error("After copying e1 to e2 and changing a field value, e1 IS STILL EQUAL to e2!")
	}

}

func TestErrBaseInfo_Equal_04(t *testing.T) {

	e1 := testCreateErrBaseInfoObject()

	e2 := e1.DeepCopyBaseInfo()

	e2.ParentObjectName = "xx"

	if e1.Equal(&e2) {
		t.Error("After copying e1 to e2 and changing a field value, e1 IS STILL EQUAL to e2!")
	}

}

func TestErrBaseInfo_Equal_05(t *testing.T) {

	e1 := testCreateErrBaseInfoObject()

	e2 := e1.DeepCopyBaseInfo()

	e2.BaseErrorId = 25

	if e1.Equal(&e2) {
		t.Error("After copying e1 to e2 and changing a field value, e1 IS STILL EQUAL to e2!")
	}

}

func testCreateErrBaseInfoObject() ErrBaseInfo {
	ex1 := "test11.go"
	ex2 :="TestObj11"
	ex3 := "test11"
	ex4 := int64(1000)

	return ErrBaseInfo{SourceFileName: ex1, ParentObjectName: ex2, FuncName:ex3, BaseErrorId: ex4}

}
