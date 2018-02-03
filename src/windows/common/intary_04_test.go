package common

import (
	"math/big"
	"testing"
)

func TestIntAry_Equals_01(t *testing.T) {
	nStr1 := "000549721.32178000"

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()
	ia2.CopyIn(&ia, false)

	if !ia.Equals(&ia2) {
		t.Error("Error: ia NOT EQUAL to ia2!")
	}

}

func TestIntAry_Equals_02(t *testing.T) {
	nStr1 := "-000549721.32178000"

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.CopyToBackUp()

	ia2 := IntAry{}.New()
	ia2.CopyIn(&ia, true)

	if !ia.Equals(&ia2) {
		t.Error("Error: ia NOT EQUAL to ia2!")
	}

	if !ia.BackUp.Equals(&ia.BackUp) {
		t.Error("Error: ia.Backup != ia2.Backup!")
	}

}

func TestIntAry_Equals_03(t *testing.T) {
	nStr1 := "-000549721.32178000"

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.CopyToBackUp()

	ia2 := IntAry{}.New()
	ia2.CopyIn(&ia, true)

	ia2.SetSign(1)

	if ia.Equals(&ia2) {
		t.Error("Error: ia EQUALS ia2!")
	}

}

func TestIntAry_Equals_04(t *testing.T) {
	nStr1 := "-000549721.32178000"

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.CopyToBackUp()

	ia2 := IntAry{}.New()
	ia2.CopyIn(&ia, true)

	ia2.BackUp.SetSignValue(1)

	if !ia.Equals(&ia2) {
		t.Error("Error: ia NOT EQUAL ia2!")
	}

	if ia.BackUp.Equals(&ia2.BackUp) {
		t.Error("Error: ia.BackUp SHOULD NOT EQUAL ia2.BackUp")
	}

}

func TestIntAry_Floor_01(t *testing.T) {
	nStr1 := "99.925"
	expected := "99.000"
	precision := 3

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_02(t *testing.T) {
	nStr1 := "-99.925"
	expected := "-100.000"
	precision := 3

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_03(t *testing.T) {
	nStr1 := "0.925"
	expected := "0.000"
	precision := 3

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_04(t *testing.T) {
	nStr1 := "2.0"
	expected := "2.0"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_05(t *testing.T) {
	nStr1 := "-2.7"
	expected := "-3.0"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_06(t *testing.T) {
	nStr1 := "-2"
	expected := "-2"
	precision := 0

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_Floor_07(t *testing.T) {
	nStr1 := "2.9"
	expected := "2.0"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.Floor()

	if err != nil {
		t.Errorf("Received Error from ia.Floor(). Error:= %v", err)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

}

func TestIntAry_GetAbsoluteValue_01(t *testing.T) {
	numStr := "-927.351"
	eNumStr := "927.351"

	ia1, _ := IntAry{}.NewNumStr(numStr)

	if numStr != ia1.GetNumStr() {
		t.Errorf("Error. Expected initialized NumStr= %v .  Instead, NumStr= %v", numStr, ia1.GetNumStr())
	}

	ia2 := ia1.GetAbsoluteValue()

	if eNumStr != ia2.GetNumStr() {
		t.Errorf("Error. Expected absolute value NumStr= %v .  Instead, NumStr= %v", eNumStr, ia2.GetAbsoluteValue())
	}

}

func TestIntAry_GetCurrencySymbol_01(t *testing.T) {

	ia := IntAry{}.New()
	curSymbol := ia.GetCurrencySymbol()
	if '$' != curSymbol {
		t.Errorf("Error: Expected Currency Symbol= '$'. Instead, received Currency Symbol= '%v'", curSymbol)
	}

}

func TestIntAry_GetCurrencySymbol_03(t *testing.T) {

	ia, _ := IntAry{}.NewNumStr("50.37")

	curSymbol := ia.GetCurrencySymbol()

	if '$' != curSymbol {
		t.Errorf("Error: Expected Currency Symbol= '$'. Instead, received Currency Symbol= '%v'", curSymbol)
	}

}

func TestIntAry_GetCurrencySymbol_04(t *testing.T) {
	ia, _ := IntAry{}.NewInt64(int64(5034), 2)

	curSymbol := ia.GetCurrencySymbol()

	if '$' != curSymbol {
		t.Errorf("Error: Expected Currency Symbol= '$'. Instead, received Currency Symbol= '%v'", curSymbol)
	}

}

func TestIntAry_GetCurrencySymbol_05(t *testing.T) {

	ia, _ := IntAry{}.NewNumStr("50.37")

	var poundSym  rune

	poundSym = '\U000000a3'

	ia.SetCurrencySymbol(poundSym)

	curSymbol := ia.GetCurrencySymbol()

	if poundSym != curSymbol {
		t.Errorf("Error: Expected Currency Symbol= '%v'. Instead, received Currency Symbol= '%v'", poundSym, curSymbol)
	}

}

func TestIntAry_GetDecimalSeparator_01(t *testing.T) {
	ia := IntAry{}.New()

	decimalSeparator := ia.GetDecimalSeparator()

	if '.' != decimalSeparator {
		t.Errorf("Error: Expected Currency Symbol= '.'. Instead, received Currency Symbol= '%v'", decimalSeparator)
	}

}


func TestIntAry_GetDecimalSeparator_02(t *testing.T) {
	ia, _ := IntAry{}.NewNumStr("50.37")

	decimalSeparator := ia.GetDecimalSeparator()

	if '.' != decimalSeparator {
		t.Errorf("Error: Expected Currency Symbol= '.'. Instead, received Currency Symbol= '%v'", decimalSeparator)
	}

}

func TestIntAry_GetDecimalSeparator_03(t *testing.T) {
	ia, _ := IntAry{}.NewInt(5064, 2)

	decimalSeparator := ia.GetDecimalSeparator()

	if '.' != decimalSeparator {
		t.Errorf("Error: Expected Currency Symbol= '.'. Instead, received Currency Symbol= '%v'", decimalSeparator)
	}

}

func TestIntAry_GetDecimalSeparator_04(t *testing.T) {
	ia := IntAry{}.New()

	var frenchDecSeparator rune

	frenchDecSeparator = ','

	var frenchThousandsSeparator rune

	frenchThousandsSeparator = ' '

	ia.SetDecimalSeparator(frenchDecSeparator)
	ia.SetThousandsSeparator(frenchThousandsSeparator)

	ia.SetIntAryWithNumStr("450 123 647,1234")

	decimalSeparator := ia.GetDecimalSeparator()

	if frenchDecSeparator != decimalSeparator {
		t.Errorf("Error: Expected Currency Symbol= '%v'. Instead, received Currency Symbol= '%v'", frenchDecSeparator, decimalSeparator)
	}

	numStr := ia.GetNumStr()

	expectedNumStr := "450123647,1234"

	if expectedNumStr != numStr {
		t.Errorf("Error: Expected French Decimal separated NumStr= '%v'. Instead received NumStr= '%v'", expectedNumStr, numStr)
	}

}


func TestIntAry_GetInt_01(t *testing.T) {

	nStr1 := "50"
	expected := int(50)
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	result, err := ia.GetInt()

	if err != nil {
		t.Errorf("Error returned from ia.GetInt(). nStr= '%v' .  Error= %v", nStr1, err)
	}

	if expected != result {
		t.Errorf("Error Expected ia.GetInt()= %v .  Instead ia.GetInt()= %v .", expected, result)
	}

}

func TestIntAry_GetInt_02(t *testing.T) {

	nStr1 := "-50"
	expected := int(-50)
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	result, err := ia.GetInt()

	if err != nil {
		t.Errorf("Error returned from ia.GetInt(). nStr= '%v' .  Error= %v", nStr1, err)
	}

	if expected != result {
		t.Errorf("Error Expected ia.GetInt()= %v .  Instead ia.GetInt()= %v .", expected, result)
	}

}

func TestIntAry_GetInt_03(t *testing.T) {

	nStr1 := "2147483647"
	expected := int(2147483647)
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	result, err := ia.GetInt()

	if err != nil {
		t.Errorf("Error returned from ia.GetInt(). nStr= '%v' .  Error= %v", nStr1, err)
	}

	if expected != result {
		t.Errorf("Error Expected ia.GetInt()= %v .  Instead ia.GetInt()= %v .", expected, result)
	}

}

func TestIntAry_GetInt_04(t *testing.T) {

	nStr1 := "2147483648"
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	_, err := ia.GetInt()

	if err == nil {
		t.Error("Input Value Exceeded maximum allowable value for a 32-bit Int. This should have thrown an error. However, no error was thrown! ")
	}

}

func TestIntAry_GetInt_05(t *testing.T) {

	nStr1 := "-2147483649"
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	_, err := ia.GetInt()

	if err == nil {
		t.Error("Input Value was less than the minimum allowable value for a 32-big Int. This should have thrown an error. However, no error was thrown! ")
	}

}

func TestIntAry_GetIntAry_01(t *testing.T) {

	expectedAry := []uint8{1,2,3,4,5,6,9,4,8,2,9}

	ia, _ := IntAry{}.NewNumStr("12345694829")

	iAry, iAryLen := ia.GetIntAry()

	expectedLen := 11

	if expectedLen != iAryLen {
		t.Errorf("Error: Expected Retrned Int Ary Length= '%v'. Actual Int Ary Length= '%v'", expectedLen, iAryLen)
	}

	for i:= 0 ; i < iAryLen; i++ {
		if expectedAry[i] != iAry[i] {
			t.Errorf("Error: Returned Int Ary Element NOT EQUAL to expected Array Element. i= '%v' Expected Value='%v'. Actual Value='%v'", i, expectedAry[i], iAry[i])
		}
	}

}

func TestIntAry_GetIntAry_02(t *testing.T) {


	ia, _ := IntAry{}.NewNumStr("12345694829")

	iAry, iAryLen := ia.GetIntAry()

	iAry[4] = 9

	expectedLen := 11

	if expectedLen != iAryLen {
		t.Errorf("Error: Expected Retrned Int Ary Length= '%v'. Actual Int Ary Length= '%v'", expectedLen, iAryLen)
	}


	iAry2, _ := ia.GetIntAry()

	// GetIntAry() returns a reference to the internal array because we are using 'slices'. Be careful!!!
	if iAry[4] != iAry2[4] {
		t.Errorf("Error: Changed Reference to iAry[4]. However the internal Int Ary Value remained unchanged. Old Value='%v'  New Value=%v", iAry[4], iAry2[4])
	}

}

func TestIntAry_GetIntAryDeepCopy(t *testing.T) {

	expectedAry := []uint8{1,2,3,4,5,6,9,4,8,2,9}

	ia, _ := IntAry{}.NewNumStr("12345694829")

	iAry, iAryLen := ia.GetIntAryDeepCopy()

	expectedLen := 11

	if expectedLen != iAryLen {
		t.Errorf("Error: Expected Retrned Int Ary Length= '%v'. Actual Int Ary Length= '%v'", expectedLen, iAryLen)
	}

	for i:= 0 ; i < iAryLen; i++ {
		if expectedAry[i] != iAry[i] {
			t.Errorf("Error: Returned Int Ary Element NOT EQUAL to expected Array Element. i= '%v' Expected Value='%v'. Actual Value='%v'", i, expectedAry[i], iAry[i])
		}
	}

}

func TestIntAry_GetIntAryDeepCopy_02(t *testing.T) {


	ia, _ := IntAry{}.NewNumStr("12345694829")

	iAry, iAryLen := ia.GetIntAryDeepCopy()

	iAry[4] = 9

	expectedLen := 11

	if expectedLen != iAryLen {
		t.Errorf("Error: Expected Retrned Int Ary Length= '%v'. Actual Int Ary Length= '%v'", expectedLen, iAryLen)
	}


	iAry2, _ := ia.GetIntAryDeepCopy()

	// GetIntAry() returns a reference to the internal array because we are using 'slices'. Be careful!!!
	if iAry[4] == iAry2[4] {
		t.Errorf("Error: Changed Reference to iAry[4]. However the internal Int Ary Value was also changed. Old Value='%v'  New Value=%v", iAry[4], iAry2[4])
	}

}

func TestIntAry_GetInt64_01(t *testing.T) {

	nStr1 := "50"
	expected := int64(50)
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	result, err := ia.GetInt64()

	if err != nil {
		t.Errorf("Error returned from ia.GetInt64(). nStr= '%v' .  Error= %v", nStr1, err)
	}

	if expected != result {
		t.Errorf("Error Expected ia.GetInt64()= %v .  Instead ia.GetInt64()= %v .", expected, result)
	}

}

func TestIntAry_GetInt64_02(t *testing.T) {

	nStr1 := "-50"
	expected := int64(-50)
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	result, err := ia.GetInt64()

	if err != nil {
		t.Errorf("Error returned from ia.GetInt64(). nStr= '%v' .  Error= %v", nStr1, err)
	}

	if expected != result {
		t.Errorf("Error Expected ia.GetInt64()= %v .  Instead ia.GetInt64()= %v .", expected, result)
	}

}

func TestIntAry_GetInt64_03(t *testing.T) {

	nStr1 := "9223372036854775807"
	expected := int64(9223372036854775807)
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	result, err := ia.GetInt64()

	if err != nil {
		t.Errorf("Error returned from ia.GetInt64(). nStr= '%v' .  Error= %v", nStr1, err)
	}

	if expected != result {
		t.Errorf("Error Expected ia.GetInt64()= %v .  Instead ia.GetInt64()= %v .", expected, result)
	}

}

func TestIntAry_GetInt64_04(t *testing.T) {

	nStr1 := "9223372036854775808"
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	_, err := ia.GetInt64()

	if err == nil {
		t.Error("Input Value Exceeded maximum allowable value for an Int64. This should have thrown an error. However, no error was thrown! ")
	}

}

func TestIntAry_GetInt64_05(t *testing.T) {

	nStr1 := "-9223372036854775809"
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	_, err := ia.GetInt64()

	if err == nil {
		t.Error("Input Value was less than the minimum allowable value for an Int64. This should have thrown an error. However, no error was thrown! ")
	}
}

func TestIntAry_GetInt64_06(t *testing.T) {

	nStr1 := "-9223372036854775808"
	expected := int64(-9223372036854775808)
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	result, err := ia.GetInt64()

	if err != nil {
		t.Errorf("Error returned from ia.GetInt64(). nStr= '%v' .  Error= %v", nStr1, err)
	}

	if expected != result {
		t.Errorf("Error Expected ia.GetInt64()= %v .  Instead ia.GetInt64()= %v .", expected, result)
	}

}

func TestIntAry_GetFractionalDigits_01(t *testing.T) {

	nStr1 := "2.7894"
	expected := "0.7894"
	precision := 4
	signVal := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetFractionalDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	iAry2Stats := iAry2.GetIntAryStats()

	if false != iAry2Stats.IsZeroValue {
		t.Errorf("Error. Expected iAry2 IsZeroValue= '%v'. Instead, got isZeroValue='%v'\n", false, iAry2Stats.IsZeroValue)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

	if signVal != iAry2.GetSign() {
		t.Errorf("Error. Expected signVal= '%v'. Instead, got signVal='%v'\n", signVal, iAry2.GetSign())
	}

}

func TestIntAry_GetFractionalDigits_02(t *testing.T) {

	nStr1 := "2"
	expected := "0"
	precision := 0
	signVal := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetFractionalDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	iAry2Stats := iAry2.GetIntAryStats()

	if true != iAry2Stats.IsZeroValue {
		t.Errorf("Error. Expected iAry2 IsZeroValue= '%v'. Instead, got isZeroValue='%v'\n", true, iAry2Stats.IsZeroValue)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

	if signVal != iAry2.GetSign() {
		t.Errorf("Error. Expected signVal= '%v'. Instead, got signVal='%v'\n", signVal, iAry2.GetSign())
	}

}

func TestIntAry_GetFractionalDigits_03(t *testing.T) {

	nStr1 := "2.00"
	expected := "0.00"
	precision := 2
	signVal := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetFractionalDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	iAry2Stats := iAry2.GetIntAryStats()

	if true != iAry2Stats.IsZeroValue {
		t.Errorf("Error. Expected iAry2 IsZeroValue= '%v'. Instead, got isZeroValue='%v'\n", true, iAry2Stats.IsZeroValue)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

	if signVal != iAry2.GetSign() {
		t.Errorf("Error. Expected signVal= '%v'. Instead, got signVal='%v'\n", signVal, iAry2.GetSign())
	}

}

func TestIntAry_GetFractionalDigits_04(t *testing.T) {

	nStr1 := "-2.978562154907"
	expected := "0.978562154907"
	precision := 12
	signVal := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetFractionalDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	iAry2Stats := iAry2.GetIntAryStats()

	if false != iAry2Stats.IsZeroValue {
		t.Errorf("Error. Expected iAry2 IsZeroValue= '%v'. Instead, got isZeroValue='%v'\n", false, iAry2Stats.IsZeroValue)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

	if signVal != iAry2.GetSign() {
		t.Errorf("Error. Expected signVal= '%v'. Instead, got signVal='%v'\n", signVal, iAry2.GetSign())
	}

}

func TestIntAry_GetIntegerDigits_01(t *testing.T) {

	nStr1 := "997562.4692"
	expected := "997562"
	precision := 0
	signVal := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetIntegerDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	iAry2Stats := iAry2.GetIntAryStats()

	if false != iAry2Stats.IsZeroValue {
		t.Errorf("Error. Expected iAry2 IsZeroValue= '%v'. Instead, got isZeroValue='%v'\n", false, iAry2Stats.IsZeroValue)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

	if signVal != iAry2.GetSign() {
		t.Errorf("Error. Expected signVal= '%v'. Instead, got signVal='%v'\n", signVal, iAry2.GetSign())
	}

}

func TestIntAry_GetIntegerDigits_02(t *testing.T) {

	nStr1 := "0.4692"
	expected := "0"
	precision := 0
	signVal := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetIntegerDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	iAry2Stats := iAry2.GetIntAryStats()

	if true != iAry2Stats.IsZeroValue {
		t.Errorf("Error. Expected iAry2 IsZeroValue= '%v'. Instead, got isZeroValue='%v'\n", true, iAry2Stats.IsZeroValue)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

	if signVal != iAry2.GetSign() {
		t.Errorf("Error. Expected signVal= '%v'. Instead, got signVal='%v'\n", signVal, iAry2.GetSign())
	}

}

func TestIntAry_GetIntegerDigits_03(t *testing.T) {

	nStr1 := "-987.4692"
	expected := "-987"
	precision := 0
	signVal := -1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetIntegerDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	iAry2Stats := iAry2.GetIntAryStats()

	if false != iAry2Stats.IsZeroValue {
		t.Errorf("Error. Expected iAry2 IsZeroValue= '%v'. Instead, got isZeroValue='%v'\n", false, iAry2Stats.IsZeroValue)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, iAry2.GetPrecision())
	}

	if signVal != iAry2.GetSign() {
		t.Errorf("Error. Expected signVal= '%v'. Instead, got signVal='%v'\n", signVal, iAry2.GetSign())
	}

}

func TestIntAry_GetIntegerDigits_04(t *testing.T) {

	nStr1 := "-0.4692"
	expected := "0"
	precision := 0
	signVal := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	iAry2, err := ia.GetIntegerDigits()

	if err != nil {
		t.Errorf("Received Error from ia.Ceiling(). Error:= %v", err)
	}

	iAry2Stats := iAry2.GetIntAryStats()

	if true != iAry2Stats.IsZeroValue {
		t.Errorf("Error. Expected iAry2 IsZeroValue= '%v'. Instead, got IsZeroValue='%v'", true, iAry2Stats.IsZeroValue)
	}

	s := iAry2.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'", expected, s)
	}

	if iAry2.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'", precision, iAry2.GetPrecision())
	}

	if signVal != iAry2.GetSign() {
		t.Errorf("Error. Expected signVal= '%v'. Instead, got signVal='%v'", signVal, iAry2.GetSign())
	}

}

func TestIntAry_GetNthRootOfThis_01(t *testing.T) {
	numStr1 := "125"
	nthRoot := uint(5)
	maxPrecision := uint(14)
	expected := "2.62652780440377"
	ia, _ := IntAry{}.NewNumStr(numStr1)
	iaResult, err := ia.GetNthRootOfThis(nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia.GetNthRootOfThis(nthRoot, maxPrecision) - Error= %v", err)
	}

	if expected != iaResult.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead iaResult.GetNumStr()= %v .", expected, iaResult.GetNumStr())
	}

}

func TestIntAry_GetNthRootOfThis_02(t *testing.T) {

	numStr1 := "5604423"
	nthRoot := uint(6)
	maxPrecision := uint(13)
	expected := "13.3276982415963"
	ia, _ := IntAry{}.NewNumStr(numStr1)
	iaResult, err := ia.GetNthRootOfThis(nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != iaResult.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead iaResult.GetNumStr()= %v .", expected, iaResult.GetNumStr())
	}

}

func TestIntAry_GetNthRootOfThis_03(t *testing.T) {

	numStr1 := "5604423.924"
	nthRoot := uint(6)
	maxPrecision := uint(13)
	expected := "13.3276986078187"
	ia, _ := IntAry{}.NewNumStr(numStr1)

	iaResult, err := ia.GetNthRootOfThis(nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from iaResult.GetNthRootIntAry() - %v", err)
	}

	if expected != iaResult.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead iaResult.GetNumStr()= %v .", expected, iaResult.GetNumStr())
	}

}

func TestIntAry_GetNthRootOfThis_04(t *testing.T) {

	numStr1 := "-27"
	nthRoot := uint(3)
	maxPrecision := uint(2)
	expected := "-3.00"
	ia, _ := IntAry{}.NewNumStr(numStr1)

	iaResult, err := ia.GetNthRootOfThis(nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from iaResult.GetNthRootIntAry() - %v", err)
	}

	if expected != iaResult.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead iaResult.GetNumStr()= %v .", expected, iaResult.GetNumStr())
	}

}

func TestIntAry_GetNthRootOfThis_05(t *testing.T) {
	ia := IntAry{}.New()
	numStr1 := "-27"
	nthRoot := uint(4)
	maxPrecision := uint(2)
	ia.SetIntAryWithNumStr(numStr1)
	_, err := ia.GetNthRootOfThis(nthRoot, maxPrecision)

	if err == nil {
		t.Error("Expected Error from iaResult.GetNthRootIntAry() for negative number with even nthRoot. No Error triggered")
	}

}

func TestIntAry_GetNthRootOfThis_06(t *testing.T) {
	numStr1 := "-5604423.924"
	nthRoot := uint(5)
	maxPrecision := uint(13)
	expected := "-22.3720713464898"
	ia, _ := IntAry{}.NewNumStr(numStr1)

	iaResult, err := ia.GetNthRootOfThis(nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from nRt.GetNthRootIntAry() - %v", err)
	}

	if expected != iaResult.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead iaResult.GetNumStr()= %v .", expected, iaResult.GetNumStr())
	}

}

func TestIntAry_GetNthRootOfThis_07(t *testing.T) {

	numStr1 := "5604423.924"
	nthRoot := uint(0)
	maxPrecision := uint(1)
	expected := "1.0"
	ia, _ := IntAry{}.NewNumStr(numStr1)

	iaResult, err := ia.GetNthRootOfThis(nthRoot, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia.GetNthRootOfThis(nthRoot, maxPrecision) - %v", err)
	}

	if expected != iaResult.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead iaResult.GetNumStr()= %v .", expected, iaResult.GetNumStr())
	}

}

func TestIntAry_GetNthRootOfThis_08(t *testing.T) {
	numStr1 := "27"
	ia, _ := IntAry{}.NewNumStr(numStr1)
	nthRoot := uint(1)
	maxPrecision := uint(2)
	_, err := ia.GetNthRootOfThis(nthRoot, maxPrecision)
	if err == nil {
		t.Error("Expected Error from ia.GetNthRootOfThis(nthRoot, maxPrecision) for nthRoot == 1. No Error triggered")
	}

}

func TestIntAry_GetScaleFactor_01(t *testing.T) {
	numSr1 := "2686.12345"
	exScale := big.NewInt(100000)

	ia, _ := IntAry{}.NewNumStr(numSr1)

	iaScaleFac, err := ia.GetScaleFactor()

	if err != nil {
		t.Errorf("Error returned from ia.GetScaleFactor() - Error: %v", err)
	}

	if iaScaleFac.Cmp(exScale) != 0 {
		t.Errorf("Expected scale factor of %v  .   Instead, actual scale factor = %v", exScale, iaScaleFac.String())
	}

}

func TestIntAry_GetSquareRootInt_01(t *testing.T) {

	numSr1 := "2686.5"
	maxPrecision := uint(30)
	expected := "51.831457629512986714934518985668"
	ai, _ := IntAry{}.NewNumStr(numSr1)

	aiResult, err := ai.GetSquareRootOfThis(maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ai.GetSquareRootOfThis(30) - Error: %v", err)
	}

	if expected != aiResult.GetNumStr() {
		t.Errorf("Expected result= %v .  Instead ai.GetNumStr()= %v .", expected, aiResult.GetNumStr())
	}

	if int(maxPrecision) != aiResult.GetPrecision() {
		t.Errorf("Expected precision= %v .  Instead precision= %v .", maxPrecision, aiResult.GetPrecision())
	}

}

func TestIntAry_GetThousandsSeparator_01(t *testing.T) {
	ia := IntAry{}.New()

	var expectedDecimalSeparator rune

	expectedDecimalSeparator = '.'

	decimalSeparator := ia.GetDecimalSeparator()

	if expectedDecimalSeparator != decimalSeparator {
		t.Errorf("Error: Expected Currency Symbol= '%v'. Instead, received Currency Symbol= '%v'", expectedDecimalSeparator, decimalSeparator)
	}

}

func TestIntAry_GetThousandsSeparator_02(t *testing.T) {
	ia, _ := IntAry{}.NewNumStr("50.47")

	var expectedDecimalSeparator rune

	expectedDecimalSeparator = '.'

	decimalSeparator := ia.GetDecimalSeparator()

	if expectedDecimalSeparator != decimalSeparator {
		t.Errorf("Error: Expected Currency Symbol= '%v'. Instead, received Currency Symbol= '%v'", expectedDecimalSeparator, decimalSeparator)
	}

}

func TestIntAry_GetThousandsSeparator_03(t *testing.T) {
	ia, _ := IntAry{}.NewFloat64(float64(50.47), 2)

	var expectedDecimalSeparator rune

	expectedDecimalSeparator = '.'

	decimalSeparator := ia.GetDecimalSeparator()

	if expectedDecimalSeparator != decimalSeparator {
		t.Errorf("Error: Expected Currency Symbol= '%v'. Instead, received Currency Symbol= '%v'", expectedDecimalSeparator, decimalSeparator)
	}

}

func TestIntAry_GetThousandsSeparator_04(t *testing.T) {
	ia := IntAry{}.New()

	var frenchDecSeparator rune

	frenchDecSeparator = ','

	var frenchThousandsSeparator rune

	frenchThousandsSeparator = ' '

	ia.SetDecimalSeparator(frenchDecSeparator)
	ia.SetThousandsSeparator(frenchThousandsSeparator)

	ia.SetIntAryWithNumStr("450 123 647,1234")

	decimalSeparator := ia.GetDecimalSeparator()

	if frenchDecSeparator != decimalSeparator {
		t.Errorf("Error: Expected Currency Symbol= '%v'. Instead, received Currency Symbol= '%v'", frenchDecSeparator, decimalSeparator)
	}

	numStr := ia.GetNumStr()

	expectedNumStr := "450123647,1234"

	if expectedNumStr != numStr {
		t.Errorf("Error: Expected French Decimal separated NumStr= '%v'. Instead received NumStr= '%v'", expectedNumStr, numStr)
	}

}


func TestIntAry_IncrementIntegerOne_01(t *testing.T) {
	expected := "100.123"
	nStr1 := "-100.123"
	cycles := 200
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	for i := 0; i < cycles; i++ {
		ia.IncrementIntegerOne()
	}

	if expected != ia.GetNumStr() {
		t.Errorf("Error - Expected numStr= '%v'. Instead, numStr= '%v'", expected, ia.GetNumStr())
	}

}

func TestIntAry_IncrementIntegerOne_02(t *testing.T) {
	nStr1 := "-2000"
	expected := "2000"
	cycles := 4000
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	for i := 0; i < cycles; i++ {
		ia.IncrementIntegerOne()
	}

	if expected != ia.GetNumStr() {
		t.Errorf("Error - Expected numStr= '%v'. Instead, numStr= '%v'", expected, ia.GetNumStr())
	}

}

func TestIntAry_IncrementIntegerOne_03(t *testing.T) {
	nStr1 := "-2000.123"
	expected := "2000.123"
	cycles := 4000
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	for i := 0; i < cycles; i++ {
		ia.IncrementIntegerOne()
	}

	if expected != ia.GetNumStr() {
		t.Errorf("Error - Expected numStr= '%v'. Instead, numStr= '%v'", expected, ia.GetNumStr())
	}

}

func TestIntAry_IncrementIntegerOne_04(t *testing.T) {
	nStr1 := "0"
	expected := "40"
	cycles := 40
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	for i := 0; i < cycles; i++ {
		ia.IncrementIntegerOne()
	}

	if expected != ia.GetNumStr() {
		t.Errorf("Error - Expected numStr= '%v'. Instead, numStr= '%v'", expected, ia.GetNumStr())
	}

}

func TestIntAry_Inverse_01(t *testing.T) {
	nStr := "25"
	eNumStr := "0.04"
	iaBase, _ := IntAry{}.NewNumStr(nStr)
	maxPrecision := 2
	signVal := 1

	iaInverse, err := iaBase.Inverse(maxPrecision)

	if err != nil {
		t.Errorf("Error received from iaBase.Inverse(3) - Error= %v", err)
	}

	if eNumStr != iaInverse.GetNumStr() {
		t.Errorf("Expected NumStr= %v .  Instead, NumStr= %v .", eNumStr, iaInverse.GetNumStr() )
	}

	if maxPrecision != iaInverse.GetPrecision() {
		t.Errorf("Expected precision= %v .  Instead, precision= %v .", maxPrecision, iaInverse.GetPrecision())
	}

	if signVal != iaInverse.GetSign() {
		t.Errorf("Expected precision= %v .  Instead, precision= %v .", maxPrecision, iaInverse.GetPrecision())
	}

}

func TestIntAry_Inverse_02(t *testing.T) {
	nStr := "30517578125"
	eNumStr := "0.000000000032768"
	iaBase, _ := IntAry{}.NewNumStr(nStr)
	maxPrecision := 15
	signVal := 1

	iaInverse, err := iaBase.Inverse(maxPrecision)

	if err != nil {
		t.Errorf("Error received from iaBase.Inverse(3) - Error= %v", err)
	}

	if eNumStr != iaInverse.GetNumStr() {
		t.Errorf("Expected NumStr= %v .  Instead, NumStr= %v .", eNumStr, iaInverse.GetNumStr() )
	}

	if maxPrecision != iaInverse.GetPrecision() {
		t.Errorf("Expected precision= %v .  Instead, precision= %v .", maxPrecision, iaInverse.GetPrecision())
	}

	if signVal != iaInverse.GetSign() {
		t.Errorf("Expected precision= %v .  Instead, precision= %v .", maxPrecision, iaInverse.GetPrecision())
	}

}

func TestIntAry_Inverse_03(t *testing.T) {
	nStr := "25"
	eNumStr := "0.04"
	iaBase, _ := IntAry{}.NewNumStr(nStr)
	maxPrecision := 2
	signVal := 1

	iaInverse, err := iaBase.Inverse(maxPrecision)

	if err != nil {
		t.Errorf("Error received from iaBase.Inverse(3) - Error= %v", err)
	}

	if eNumStr != iaInverse.GetNumStr() {
		t.Errorf("Expected NumStr= %v .  Instead, NumStr= %v .", eNumStr, iaInverse.GetNumStr() )
	}

	if maxPrecision != iaInverse.GetPrecision() {
		t.Errorf("Expected precision= %v .  Instead, precision= %v .", maxPrecision, iaInverse.GetPrecision())
	}

	if signVal != iaInverse.GetSign() {
		t.Errorf("Expected precision= %v .  Instead, precision= %v .", maxPrecision, iaInverse.GetPrecision())
	}
}


