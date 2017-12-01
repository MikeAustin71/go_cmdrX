package common

import (
	"math/big"
	"testing"
)

/*
	These tests are designed to test library methods found in
	source file nthroot.go .

	This test file is located in source code repository:

       https://github.com/MikeAustin71/mathhlpr.git

*/

func TestNthRootOp_GetNthRootFloat32_01(t *testing.T) {
	num := float32(125.0)
	nthRoot := uint(5)
	maxPrecision := uint(14)
	expected := "2.62652780440377"

	nRt := NthRootOp{}
	ai, err := nRt.GetNthRootFloat32(num, 0, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootFloat32() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootFloat64_01(t *testing.T) {
	num := float64(125.0)
	nthRoot := uint(5)
	maxPrecision := uint(14)
	expected := "2.62652780440377"

	nRt := NthRootOp{}
	ai, err := nRt.GetNthRootFloat64(num, 0, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootFloat64() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootBigFloat_01(t *testing.T) {
	num := big.NewFloat(float64(125.0))
	nthRoot := uint(5)
	maxPrecision := uint(14)
	expected := "2.62652780440377"

	nRt := NthRootOp{}
	ai, err := nRt.GetNthRootBigFloat(num, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootBigFloat() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootInt32_01(t *testing.T) {

	num := 125
	nthRoot := uint(5)
	maxPrecision := uint(14)
	expected := "2.62652780440377"

	nRt := NthRootOp{}
	ai, err := nRt.GetNthRootInt(num, uint(0), nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootInt() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootInt64_01(t *testing.T) {

	num := int64(125)
	nthRoot := uint(5)
	maxPrecision := uint(14)
	expected := "2.62652780440377"

	nRt := NthRootOp{}
	ai, err := nRt.GetNthRootInt64(num, uint(0), nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootBigInt_01(t *testing.T) {

	num := big.NewInt(125)
	nthRoot := uint(5)
	maxPrecision := uint(14)
	expected := "2.62652780440377"

	nRt := NthRootOp{}
	ai, err := nRt.GetNthRootBigInt(num, uint(0), nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_01(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "125"
	nthRoot := uint(5)
	maxPrecision := uint(14)
	expected := "2.62652780440377"
	originalNum.SetIntAryWithNumStr(numStr1)
	ai, err := nRt.GetNthRootIntAry(&originalNum, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_02(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "5604423"
	nthRoot := uint(6)
	maxPrecision := uint(13)
	expected := "13.3276982415963"
	originalNum.SetIntAryWithNumStr(numStr1)
	ai, err := nRt.GetNthRootIntAry(&originalNum, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_03(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "5604423.924"
	nthRoot := uint(6)
	maxPrecision := uint(13)
	expected := "13.3276986078187"
	originalNum.SetIntAryWithNumStr(numStr1)
	ai, err := nRt.GetNthRootIntAry(&originalNum, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_04(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "-27"
	nthRoot := uint(3)
	maxPrecision := uint(2)
	expected := "-3.00"
	originalNum.SetIntAryWithNumStr(numStr1)
	ai, err := nRt.GetNthRootIntAry(&originalNum, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_05(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "-27"
	nthRoot := uint(4)
	maxPrecision := uint(2)
	originalNum.SetIntAryWithNumStr(numStr1)
	_, err := nRt.GetNthRootIntAry(&originalNum, nthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected Error from nRt.GetNthRootIntAry() for negative number with even nthRoot. No Error triggered")
	}

}

func TestNthRootOp_GetNthRootIntAry_06(t *testing.T) {

	nRt := NthRootOp{}
	numStr1 := "-5604423.924"
	nthRoot := uint(5)
	maxPrecision := uint(13)
	expected := "-22.3720713464898"
	originalNum, _ := IntAry{}.NewNumStr(numStr1)

	ai, err := nRt.GetNthRootIntAry(&originalNum, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_07(t *testing.T) {

	nRt := NthRootOp{}
	numStr1 := "5604423.924"
	nthRoot := uint(0)
	maxPrecision := uint(1)
	expected := "1.0"
	originalNum, _ := IntAry{}.NewNumStr(numStr1)

	ai, err := nRt.GetNthRootIntAry(&originalNum, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetNthRootIntAry_08(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "27"
	nthRoot := uint(1)
	maxPrecision := uint(2)
	originalNum.SetIntAryWithNumStr(numStr1)
	_, err := nRt.GetNthRootIntAry(&originalNum, nthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected Error from nRt.GetNthRootIntAry() for nthRoot == 1. No Error triggered")
	}

}

func TestNthRootOp_GetSquareRootFloat32_01(t *testing.T) {

	nRt := NthRootOp{}
	num := float32(2686.5)

	maxPrecision := uint(30)
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootFloat32(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootFloat32() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootFloat64_01(t *testing.T) {

	nRt := NthRootOp{}
	num := float64(2686.5)

	maxPrecision := uint(30)
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootFloat64(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootFloat64() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootBigFloat_01(t *testing.T) {

	nRt := NthRootOp{}
	num, ok := big.NewFloat(0).SetString("2686.5")

	if !ok {
		t.Error("Conversion failed big.NewFloat(0).SetString(\"2686.5\")")
	}

	maxPrecision := uint(30)
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootBigFloat(num, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootBigFloat() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootInt_01(t *testing.T) {

	nRt := NthRootOp{}
	num := int(26865)
	maxPrecision := uint(30)
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootInt(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootInt32() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootInt32_01(t *testing.T) {

	nRt := NthRootOp{}
	num := int32(26865)
	maxPrecision := uint(30)
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootInt32(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootInt32() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootInt64_01(t *testing.T) {

	nRt := NthRootOp{}
	num := int64(26865)
	maxPrecision := uint(30)
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootInt64(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootInt64() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootBigInt_01(t *testing.T) {

	nRt := NthRootOp{}
	num := big.NewInt(int64(26865))
	maxPrecision := uint(30)
	expected := "51.831457629512986714934518985668"
	ai, err := nRt.GetSquareRootBigInt(num, 1, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootBigInt() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_GetSquareRootIntAry_01(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "2686.5"
	maxPrecision := uint(30)
	expected := "51.831457629512986714934518985668"
	originalNum.SetIntAryWithNumStr(numStr1)
	ai, err := nRt.GetSquareRootIntAry(&originalNum, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetSquareRootIntAry() - %v", err)
	}

	if expected != ai.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, ai.GetNumStr())
	}

}

func TestNthRootOp_SetNthRootIntAry_01(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "125"
	nthRoot := uint(5)
	maxPrecision := uint(14)
	expected := "2.62652780440377"
	originalNum.SetIntAryWithNumStr(numStr1)
	err := nRt.SetNthRootIntAry(&originalNum, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != nRt.ResultAry.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, nRt.ResultAry.GetNumStr())
	}

}

func TestNthRootOp_SetNthRootIntAry_02(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "5604423"
	nthRoot := uint(6)
	maxPrecision := uint(13)
	expected := "13.3276982415963"
	originalNum.SetIntAryWithNumStr(numStr1)
	err := nRt.SetNthRootIntAry(&originalNum, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != nRt.ResultAry.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, nRt.ResultAry.GetNumStr())
	}

}

func TestNthRootOp_SetNthRootIntAry_03(t *testing.T) {

	nRt := NthRootOp{}
	originalNum := IntAry{}.New()
	numStr1 := "5604423.924"
	nthRoot := uint(6)
	maxPrecision := uint(13)
	expected := "13.3276986078187"
	originalNum.SetIntAryWithNumStr(numStr1)
	err := nRt.SetNthRootIntAry(&originalNum, nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != nRt.ResultAry.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, nRt.ResultAry.GetNumStr())
	}

}
