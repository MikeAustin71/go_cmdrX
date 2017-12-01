package common

import (
	"math/big"
	"testing"
)

func TestNumStrDto_AddNumStrs_01(t *testing.T) {
	nStr1 := "-9589.21"
	nStr2 := "9211.40"
	nStr3 := "-377.81"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_01() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_02(t *testing.T) {
	nStr1 := "9589.21"
	nStr2 := "-9211.40"
	nStr3 := "377.81"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_02() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_03(t *testing.T) {
	nStr1 := "9589.21"
	nStr2 := "9211.40"
	nStr3 := "18800.61"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_03() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_04(t *testing.T) {
	nStr1 := "-9589.21"
	nStr2 := "-9211.40"
	nStr3 := "-18800.61"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_04() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_05(t *testing.T) {
	nStr1 := "2"
	nStr2 := "3"
	nStr3 := "5"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_05() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_06(t *testing.T) {
	nStr1 := "2"
	nStr2 := "0.0"
	nStr3 := "2.0"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_06() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_07(t *testing.T) {
	nStr1 := "0"
	nStr2 := "0.0"
	nStr3 := "0.0"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_07() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_08(t *testing.T) {
	nStr1 := "-6"
	nStr2 := "67.521"
	nStr3 := "61.521"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_08() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_09(t *testing.T) {
	nStr1 := "67.521"
	nStr2 := "-6"
	nStr3 := "61.521"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_09() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_10(t *testing.T) {
	nStr1 := "-67.521"
	nStr2 := "67.521"
	nStr3 := "0.000"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)
	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_10() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_AddNumStrs_11(t *testing.T) {
	nStr1 := "67.521"
	nStr2 := "-67.521"
	nStr3 := "0.000"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.AddNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)
	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_AddNumStrs_10() - ")

	if err != nil {
		t.Errorf("Error returned from nDto.IsNumStrDtoValid(&nResult). Error= %v", err)
	}

}

func TestNumStrDto_CompareAbsoluteVals_01(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "-9211.40"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_02(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "9211.40"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_03(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "12567.218956"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_04(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "12567.218956"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_05(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "-12567.218956"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_06(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "-567.21"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_07(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "567.21"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareAbsoluteVals_08(t *testing.T) {
	nStr1 := "567.21"
	nStr2 := "1567.21"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareAbsoluteVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Absolute Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_01(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "-9211.40"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_02(t *testing.T) {
	nStr1 := "12567.218956"
	nStr2 := "9211.40"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_03(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "9211.40"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_04(t *testing.T) {
	nStr1 := "12567.218956"
	nStr2 := "-9211.40"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_05(t *testing.T) {
	nStr1 := "12567.218956"
	nStr2 := "-12567.218956"
	expectedCompare := 1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_06(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "12567.218956"
	expectedCompare := -1
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_07(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "-12567.218956"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_CompareSignedVals_08(t *testing.T) {
	nStr1 := "12567.218956"
	nStr2 := "12567.218956"
	expectedCompare := 0
	nDto := NumStrDto{}.New()
	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)

	cmpSgn := nDto.CompareSignedVals(&n1, &n2)

	if cmpSgn != expectedCompare {
		t.Errorf("Compared Signed Values n1= '%v' and n2='%v'. Expected Comparison= '%v'. Instead got '%v'", nStr1, nStr2, expectedCompare, cmpSgn)
	}

}

func TestNumStrDto_FormatForMathOps_01(t *testing.T) {
	nStr1 := "-12567.218956"
	nStr2 := "-9211.40"
	nStr3 := "-12567.218956"
	nStr4 := "-09211.400000"
	expectedCompare := 1

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if false != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.NumStrOut != n1OutDto.NumStrOut {
		t.Errorf("Expected n1OutDto.NumStrOut= '%v'. Instead got '%v'", nOut1.NumStrOut, n1OutDto.NumStrOut)
	}

	if nOut2.NumStrOut != n2OutDto.NumStrOut {
		t.Errorf("Expected n2OutDto.NumStrOut= '%v'. Instead got '%v'", nOut2.NumStrOut, n2OutDto.NumStrOut)
	}

	if nOut1.SignVal != n1OutDto.SignVal {
		t.Errorf("Expected n1OutDto.SignVal= '%v'. Instead got '%v'", nOut1.SignVal, n1OutDto.SignVal)
	}

	if nOut2.SignVal != n2OutDto.SignVal {
		t.Errorf("Expected n1OutDto.SignVal= '%v'. Instead got '%v'", nOut2.SignVal, n2OutDto.SignVal)
	}

	if nOut1.Precision != n1OutDto.Precision {
		t.Errorf("Expected n1OutDto.Precision= '%v'. Instead got '%v'", nOut1.Precision, n1OutDto.Precision)
	}

	if nOut2.Precision != n2OutDto.Precision {
		t.Errorf("Expected n2OutDto.Precision= '%v'. Instead got '%v'", nOut2.Precision, n2OutDto.Precision)
	}

}

func TestNumStrDto_FormatForMathOps_02(t *testing.T) {
	nStr1 := "-9211.40"
	nStr2 := "-12567.218956"
	nStr3 := "-12567.218956"
	nStr4 := "-09211.400000"
	expectedCompare := 1

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if true != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", true, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.NumStrOut != n1OutDto.NumStrOut {
		t.Errorf("Expected n1OutDto.NumStrOut= '%v'. Instead got '%v'", nOut1.NumStrOut, n1OutDto.NumStrOut)
	}

	if nOut2.NumStrOut != n2OutDto.NumStrOut {
		t.Errorf("Expected n2OutDto.NumStrOut= '%v'. Instead got '%v'", nOut2.NumStrOut, n2OutDto.NumStrOut)
	}

	if nOut1.SignVal != n1OutDto.SignVal {
		t.Errorf("Expected n1OutDto.SignVal= '%v'. Instead got '%v'", nOut1.SignVal, n1OutDto.SignVal)
	}

	if nOut2.SignVal != n2OutDto.SignVal {
		t.Errorf("Expected n1OutDto.SignVal= '%v'. Instead got '%v'", nOut2.SignVal, n2OutDto.SignVal)
	}

	if nOut1.Precision != n1OutDto.Precision {
		t.Errorf("Expected n1OutDto.Precision= '%v'. Instead got '%v'", nOut1.Precision, n1OutDto.Precision)
	}

	if nOut2.Precision != n2OutDto.Precision {
		t.Errorf("Expected n2OutDto.Precision= '%v'. Instead got '%v'", nOut2.Precision, n2OutDto.Precision)
	}

}

func TestNumStrDto_FormatForMathOps_03(t *testing.T) {
	nStr1 := "-6"
	nStr2 := "67.521"
	nStr3 := "67.521"
	nStr4 := "-06.000"
	expectedCompare := 1

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if true != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", true, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.NumStrOut != n1OutDto.NumStrOut {
		t.Errorf("Expected n1OutDto.NumStrOut= '%v'. Instead got '%v'", nOut1.NumStrOut, n1OutDto.NumStrOut)
	}

	if nOut2.NumStrOut != n2OutDto.NumStrOut {
		t.Errorf("Expected n2OutDto.NumStrOut= '%v'. Instead got '%v'", nOut2.NumStrOut, n2OutDto.NumStrOut)
	}

	if nOut1.SignVal != n1OutDto.SignVal {
		t.Errorf("Expected n1OutDto.SignVal= '%v'. Instead got '%v'", nOut1.SignVal, n1OutDto.SignVal)
	}

	if nOut2.SignVal != n2OutDto.SignVal {
		t.Errorf("Expected n1OutDto.SignVal= '%v'. Instead got '%v'", nOut2.SignVal, n2OutDto.SignVal)
	}

	if nOut1.Precision != n1OutDto.Precision {
		t.Errorf("Expected n1OutDto.Precision= '%v'. Instead got '%v'", nOut1.Precision, n1OutDto.Precision)
	}

	if nOut2.Precision != n2OutDto.Precision {
		t.Errorf("Expected n2OutDto.Precision= '%v'. Instead got '%v'", nOut2.Precision, n2OutDto.Precision)
	}

}

func TestNumStrDto_FormatForMathOps_04(t *testing.T) {
	nStr1 := "67.521"
	nStr2 := "-6"
	nStr3 := "67.521"
	nStr4 := "-06.000"
	expectedCompare := 1

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if false != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.NumStrOut != n1OutDto.NumStrOut {
		t.Errorf("Expected n1OutDto.NumStrOut= '%v'. Instead got '%v'", nOut1.NumStrOut, n1OutDto.NumStrOut)
	}

	if nOut2.NumStrOut != n2OutDto.NumStrOut {
		t.Errorf("Expected n2OutDto.NumStrOut= '%v'. Instead got '%v'", nOut2.NumStrOut, n2OutDto.NumStrOut)
	}

	if nOut1.SignVal != n1OutDto.SignVal {
		t.Errorf("Expected n1OutDto.SignVal= '%v'. Instead got '%v'", nOut1.SignVal, n1OutDto.SignVal)
	}

	if nOut2.SignVal != n2OutDto.SignVal {
		t.Errorf("Expected n1OutDto.SignVal= '%v'. Instead got '%v'", nOut2.SignVal, n2OutDto.SignVal)
	}

	if nOut1.Precision != n1OutDto.Precision {
		t.Errorf("Expected n1OutDto.Precision= '%v'. Instead got '%v'", nOut1.Precision, n1OutDto.Precision)
	}

	if nOut2.Precision != n2OutDto.Precision {
		t.Errorf("Expected n2OutDto.Precision= '%v'. Instead got '%v'", nOut2.Precision, n2OutDto.Precision)
	}

}

func TestNumStrDto_FormatForMathOps_05(t *testing.T) {
	nStr1 := "-67.521"
	nStr2 := "6"
	nStr3 := "-67.521"
	nStr4 := "06.000"
	expectedCompare := 1
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if false != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.NumStrOut != n1OutDto.NumStrOut {
		t.Errorf("Expected n1OutDto.NumStrOut= '%v'. Instead got '%v'", nOut1.NumStrOut, n1OutDto.NumStrOut)
	}

	if nOut2.NumStrOut != n2OutDto.NumStrOut {
		t.Errorf("Expected n2OutDto.NumStrOut= '%v'. Instead got '%v'", nOut2.NumStrOut, n2OutDto.NumStrOut)
	}

	if nOut1.SignVal != n1OutDto.SignVal {
		t.Errorf("Expected n1OutDto.SignVal= '%v'. Instead got '%v'", nOut1.SignVal, n1OutDto.SignVal)
	}

	if nOut2.SignVal != n2OutDto.SignVal {
		t.Errorf("Expected n1OutDto.SignVal= '%v'. Instead got '%v'", nOut2.SignVal, n2OutDto.SignVal)
	}

	if nOut1.Precision != n1OutDto.Precision {
		t.Errorf("Expected n1OutDto.Precision= '%v'. Instead got '%v'", nOut1.Precision, n1OutDto.Precision)
	}

	if nOut2.Precision != n2OutDto.Precision {
		t.Errorf("Expected n2OutDto.Precision= '%v'. Instead got '%v'", nOut2.Precision, n2OutDto.Precision)
	}

}

func TestNumStrDto_FormatForMathOps_06(t *testing.T) {
	nStr1 := "-67.521"
	nStr2 := "67.521"
	nStr3 := "-67.521"
	nStr4 := "67.521"
	expectedCompare := 0
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nOut1, _ := nDto.ParseNumStr(nStr3)
	nOut2, _ := nDto.ParseNumStr(nStr4)

	n1OutDto, n2OutDto, compare, isReversed, err := nDto.FormatForMathOps(n1, n2)

	if err != nil {
		t.Errorf("Error returned from nDto.FormatForMathOps(n1, n2). Error= %v", err)
	}

	if false != isReversed {
		t.Errorf("Expected isReverse = '%v'. Instead got '%v'", false, isReversed)
	}

	if compare != expectedCompare {
		t.Errorf("Expected compare result = '%v'. Instead got '%v'", expectedCompare, compare)
	}

	if nOut1.NumStrOut != n1OutDto.NumStrOut {
		t.Errorf("Expected n1OutDto.NumStrOut= '%v'. Instead got '%v'", nOut1.NumStrOut, n1OutDto.NumStrOut)
	}

	if nOut2.NumStrOut != n2OutDto.NumStrOut {
		t.Errorf("Expected n2OutDto.NumStrOut= '%v'. Instead got '%v'", nOut2.NumStrOut, n2OutDto.NumStrOut)
	}

	if nOut1.SignVal != n1OutDto.SignVal {
		t.Errorf("Expected n1OutDto.SignVal= '%v'. Instead got '%v'", nOut1.SignVal, n1OutDto.SignVal)
	}

	if nOut2.SignVal != n2OutDto.SignVal {
		t.Errorf("Expected n1OutDto.SignVal= '%v'. Instead got '%v'", nOut2.SignVal, n2OutDto.SignVal)
	}

	if nOut1.Precision != n1OutDto.Precision {
		t.Errorf("Expected n1OutDto.Precision= '%v'. Instead got '%v'", nOut1.Precision, n1OutDto.Precision)
	}

	if nOut2.Precision != n2OutDto.Precision {
		t.Errorf("Expected n2OutDto.Precision= '%v'. Instead got '%v'", nOut2.Precision, n2OutDto.Precision)
	}

}

func TestNumStrDto_CopyIn_01(t *testing.T) {

	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	nDto := NumStrDto{}.New()

	nDto.CopyIn(n1)

	s := nDto.NumStrOut

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if !nDto.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue)
	}

	if nDto.Precision != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_CopyIn_02(t *testing.T) {

	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	nDto := NumStrDto{}.New()

	nDto.CopyIn(n1)

	s := nDto.NumStrOut

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if !nDto.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue)
	}

	if nDto.Precision != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_CopyOut_01(t *testing.T) {

	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	nDto := n1.CopyOut()

	s := nDto.NumStrOut

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if !nDto.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue)
	}

	if nDto.Precision != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_CopyOut_02(t *testing.T) {

	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	nDto := n1.CopyOut()

	s := nDto.NumStrOut

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if !nDto.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue)
	}

	if nDto.Precision != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_GetAbsoluteBigInt_01(t *testing.T) {

	nStr := "-123.456"
	absNumStr := "123456"
	expected, isOk := big.NewInt(0).SetString(absNumStr, 10)

	if !isOk {
		t.Errorf("big.SetString(absNumStr,10) Failed!. absNumStr= '%v'", absNumStr)
	}

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	absBigInt, err := n1.GetAbsoluteBigInt()

	if absBigInt.Cmp(expected) != 0 {
		t.Errorf("Expected absBigInt= %v . Instead got, %v", expected.String(), absBigInt.String())
	}

}

func TestNumStrDto_GetSignedBigInt_01(t *testing.T) {

	nStr := "-123.456"
	signedNumStr := "-123456"
	expected, isOk := big.NewInt(0).SetString(signedNumStr, 10)

	if !isOk {
		t.Errorf("big.SetString(signedNumStr,10) Failed!. signedNumStr= '%v'", signedNumStr)
	}

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	signedBigInt, err := n1.GetSignedBigInt()

	if signedBigInt.Cmp(expected) != 0 {
		t.Errorf("Expected signedBigInt= %v . Instead got, %v", expected.String(), signedBigInt.String())
	}

}

func TestNumStrDto_MultiplyNumStrs_01(t *testing.T) {
	nStr1 := "35.123456"
	nStr2 := "47.9876514"
	nStr3 := "1685.4921624912384"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_02(t *testing.T) {
	nStr1 := "35.123456"
	nStr2 := "-47.9876514"
	nStr3 := "-1685.4921624912384"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_03(t *testing.T) {
	nStr1 := "-35.123456"
	nStr2 := "-47.9876514"
	nStr3 := "1685.4921624912384"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_04(t *testing.T) {
	nStr1 := "57"
	nStr2 := "123"
	nStr3 := "7011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_05(t *testing.T) {
	nStr1 := "57"
	nStr2 := "-123"
	nStr3 := "-7011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_06(t *testing.T) {
	nStr1 := "-57"
	nStr2 := "123"
	nStr3 := "-7011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_07(t *testing.T) {
	nStr1 := "-57"
	nStr2 := "-123"
	nStr3 := "7011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_08(t *testing.T) {
	nStr1 := "0"
	nStr2 := "123"
	nStr3 := "0"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_09(t *testing.T) {
	nStr1 := "57"
	nStr2 := "0"
	nStr3 := "0"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_10(t *testing.T) {
	nStr1 := "-57"
	nStr2 := "0"
	nStr3 := "0"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_11(t *testing.T) {
	nStr1 := "57"
	nStr2 := "0.123"
	nStr3 := "7.011"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_12(t *testing.T) {
	nStr1 := "62.1234567890123"
	nStr2 := "3.12345678901234"
	nStr3 := "194.039932864555212496281111782"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_13(t *testing.T) {
	nStr1 := "-62.1234567890123"
	nStr2 := "3.12345678901234"
	nStr3 := "-194.039932864555212496281111782"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_MultiplyNumStrs_14(t *testing.T) {
	nStr1 := "-62.1234567890123"
	nStr2 := "-3.12345678901234"
	nStr3 := "194.039932864555212496281111782"
	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nExpected, _ := nDto.ParseNumStr(nStr3)

	nResult, err := nDto.MultiplyNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.MultiplyNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nResult.NumStrOut
	expected := nExpected.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nResult.AbsIntRunes)
	expected = string(nExpected.AbsIntRunes)

	if expected != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", expected, s)

	}

	s = string(nResult.AbsFracRunes)
	expected = string(nExpected.AbsFracRunes)

	if expected != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", expected, s)
	}

	if nExpected.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nExpected.SignVal, nResult.SignVal)
	}

	if nExpected.HasNumericDigits != nResult.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= '%v'. Instead, got '%v'", nExpected.HasNumericDigits, nResult.HasNumericDigits)
	}

	if nExpected.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nExpected.IsFractionalValue, nResult.IsFractionalValue)
	}

	if nExpected.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nExpected.Precision, nResult.Precision)

	}

	err = nDto.IsNumStrDtoValid(&nResult, "TestNumStrDto_MultiplyNumStrs_01() - ")

	if err != nil {
		t.Errorf("Resulting NumStr is INVALD. Error= %v", err)
	}

}

func TestNumStrDto_ParseNumStr_01(t *testing.T) {
	nStr := "123.456"
	iStr := "123"
	fracStr := "456"
	signVal := 1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if !nDto.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue)
	}

	if nDto.Precision != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_ParseNumStr_02(t *testing.T) {
	nStr := "123456"
	iStr := "123456"
	fracStr := ""
	signVal := 1
	precision := uint(0)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != false {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", false, nDto.IsFractionalValue)
	}

	if nDto.Precision != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_ParseNumStr_03(t *testing.T) {
	nStr := "-123456"
	iStr := "123456"
	fracStr := ""
	signVal := -1
	precision := uint(0)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != false {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", false, nDto.IsFractionalValue)
	}

	if nDto.Precision != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_ParseNumStr_04(t *testing.T) {
	nStr := "-123.456"
	iStr := "123"
	fracStr := "456"
	signVal := -1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue)
	}

	if nDto.Precision != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_ParseNumStr_05(t *testing.T) {
	nStr := "-000.000"
	nStrOut := "0.000"
	iStr := "0"
	fracStr := "000"
	signVal := 1
	precision := uint(3)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(nStr)

	if err != nil {
		t.Errorf("Received error from NumStrDto.ParseNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	s := nDto.NumStrOut

	if s != nStrOut {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStrOut, s)
	}

	s = string(nDto.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != true {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got %v", true, nDto.IsFractionalValue)
	}

	if nDto.Precision != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_ParseSignedBigInt_01(t *testing.T) {

	signedAbsNumStr := "-123456789"
	absAllNumStr := "123456789"
	nStr := "-123456.789"
	iStr := "123456"
	fracStr := "789"
	precision := uint(3)
	signVal := -1
	sBigInt, isOk := big.NewInt(0).SetString(signedAbsNumStr, 10)

	if !isOk {
		t.Errorf("BigInt.SetString(signedAbsNumStr,10) conversion failed! nStr= '%v' ", signedAbsNumStr)
	}

	n1, err := NumStrDto{}.NewPtr().ParseSignedBigInt(sBigInt, precision)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseSignedBigInt(sBigInt, precision). sBigInt= '%v' Error= %v", sBigInt.String(), err)
	}

	nDto := n1.CopyOut()

	s := nDto.NumStrOut

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.AbsAllNumRunes)

	if absAllNumStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", absAllNumStr, s)

	}

	s = string(nDto.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if !nDto.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue)
	}

	if nDto.Precision != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_ParseSignedBigInt_02(t *testing.T) {

	signedAbsNumStr := "123456789"
	absAllNumStr := "123456789"
	nStr := "123456.789"
	iStr := "123456"
	fracStr := "789"
	precision := uint(3)
	signVal := 1
	sBigInt, isOk := big.NewInt(0).SetString(signedAbsNumStr, 10)

	if !isOk {
		t.Errorf("BigInt.SetString(signedAbsNumStr,10) conversion failed! nStr= '%v' ", signedAbsNumStr)
	}

	n1, err := NumStrDto{}.NewPtr().ParseSignedBigInt(sBigInt, precision)

	if err != nil {
		t.Errorf("Received error from n1 NumStrDto.ParseSignedBigInt(sBigInt, precision). sBigInt= '%v' Error= %v", sBigInt.String(), err)
	}

	nDto := n1.CopyOut()

	s := nDto.NumStrOut

	if s != nStr {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", nStr, s)
	}

	s = string(nDto.AbsAllNumRunes)

	if absAllNumStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", absAllNumStr, s)

	}

	s = string(nDto.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != signVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", signVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if !nDto.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= 'true'. Instead, got %v", nDto.IsFractionalValue)
	}

	if nDto.Precision != precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_ScaleNumStr_01(t *testing.T) {
	nStr := "123456"
	scale := uint(3)
	expected := "123456.000"
	iStr := "123456"
	fracStr := "000"
	signVal := 1

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, scale, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if scale != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", scale, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != iStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", iStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != fracStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", fracStr, s)
	}

}

func TestNumStrDto_ScaleNumStr_02(t *testing.T) {
	nStr := "-123456"
	precision := uint(3)
	expected := "-123456.000"
	iStr := "123456"
	fracStr := "000"
	signVal := -1

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != iStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", iStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != fracStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", fracStr, s)
	}

}

func TestNumStrDto_ScaleNumStr_03(t *testing.T) {
	nStr := "123456"
	precision := uint(9)
	expected := "123456.000000000"
	iStr := "123456"
	fracStr := "000000000"
	signVal := 1

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != iStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", iStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != fracStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", fracStr, s)
	}

}

func TestNumStrDto_ScaleNumStr_04(t *testing.T) {
	nStr := "123456.7890"
	precision := uint(2)
	expected := "123456.79"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "79"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, true)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrDto_ScaleNumStr_05(t *testing.T) {
	nStr := "123456.789"
	precision := uint(5)
	expected := "123456.78900"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "78900"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, true)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrDto_ScaleNumStr_06(t *testing.T) {
	nStr := "-123456.789"
	precision := uint(5)
	expected := "-123456.78900"
	signVal := -1
	absIntRuneStr := "123456"
	absFracRuneStr := "78900"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, true)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}
func TestNumStrDto_ScaleNumStr_07(t *testing.T) {
	nStr := "123456.789"
	precision := uint(3)
	expected := "123456.789"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	nsDto, err := NumStrDto{}.NewPtr().SetPrecision(nStr, precision, true)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrUtility_ScaleAbsoluteValStr_01(t *testing.T) {
	nStr := "123456"
	precision := uint(3)
	expected := "123.456"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "456"

	nsDto, err := NumStrDto{}.NewPtr().ScaleAbsoluteValStr(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrUtility_ScaleAbsoluteValStr_02(t *testing.T) {
	nStr := "123.456"
	precision := uint(4)
	expected := "12.3456"
	signVal := 1
	absIntRuneStr := "12"
	absFracRuneStr := "3456"

	nsDto, err := NumStrDto{}.NewPtr().ScaleAbsoluteValStr(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrUtility_ScaleAbsoluteValStr_03(t *testing.T) {
	nStr := "-123456"
	precision := uint(3)
	expected := "-123.456"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "456"

	nsDto, err := NumStrDto{}.NewPtr().ScaleAbsoluteValStr(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrUtility_ScaleAbsoluteValStr_04(t *testing.T) {
	nStr := "-123.456"
	precision := uint(4)
	expected := "-12.3456"
	signVal := -1
	absIntRuneStr := "12"
	absFracRuneStr := "3456"

	nsDto, err := NumStrDto{}.NewPtr().ScaleAbsoluteValStr(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrUtility_ScaleAbsoluteValStr_05(t *testing.T) {
	nStr := "123456"
	precision := uint(7)
	expected := "0.0123456"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "0123456"

	nsDto, err := NumStrDto{}.NewPtr().ScaleAbsoluteValStr(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrUtility_ScaleAbsoluteValStr_06(t *testing.T) {
	nStr := "123.456"
	precision := uint(3)
	expected := "123.456"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "456"

	nsDto, err := NumStrDto{}.NewPtr().ScaleAbsoluteValStr(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, 3). nStr= '%v'. Error= %v", nStr, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='123'. Instead, got %v.", s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='456'. Instead, got %v", s)
	}

}

func TestNumStrUtility_ScaleAbsoluteValStr_07(t *testing.T) {
	nStr := "123456"
	precision := uint(4)
	expected := "12.3456"
	signVal := 1
	absIntRuneStr := "12"
	absFracRuneStr := "3456"

	nsDto, err := NumStrDto{}.NewPtr().ScaleAbsoluteValStr(nStr, precision, false)

	if err != nil {
		t.Errorf("Received error from nsu.SetPrecision(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if precision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", precision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_01(t *testing.T) {

	nStr := "123456.789"
	precision := uint(3)
	outPrecision := uint(6)
	expected := "123.456789"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_02(t *testing.T) {

	nStr := "123456.789"
	precision := uint(2)
	outPrecision := uint(5)
	expected := "1234.56789"
	signVal := 1
	absIntRuneStr := "1234"
	absFracRuneStr := "56789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_03(t *testing.T) {

	nStr := "123456.789"
	precision := uint(6)
	outPrecision := uint(9)
	expected := "0.123456789"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "123456789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_04(t *testing.T) {

	nStr := "123456789"
	precision := uint(6)
	outPrecision := uint(6)
	expected := "123.456789"
	signVal := 1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_05(t *testing.T) {

	nStr := "123"
	precision := uint(5)
	outPrecision := uint(5)
	expected := "0.00123"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := "00123"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_06(t *testing.T) {

	nStr := "123456.789"
	precision := uint(0)
	outPrecision := uint(3)
	expected := "123456.789"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_07(t *testing.T) {

	nStr := "-123456.789"
	precision := uint(3)
	outPrecision := uint(6)
	expected := "-123.456789"
	signVal := -1
	absIntRuneStr := "123"
	absFracRuneStr := "456789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionLeft_08(t *testing.T) {

	nStr := "0"
	precision := uint(3)
	outPrecision := uint(0)
	expected := "0"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_01(t *testing.T) {

	nStr := "123456.789"
	precision := uint(3)
	outPrecision := uint(0)
	expected := "123456789"
	signVal := 1
	absIntRuneStr := "123456789"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_02(t *testing.T) {

	nStr := "123456.789"
	precision := uint(2)
	outPrecision := uint(1)
	expected := "12345678.9"
	signVal := 1
	absIntRuneStr := "12345678"
	absFracRuneStr := "9"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_03(t *testing.T) {

	nStr := "123456.789"
	precision := uint(6)
	outPrecision := uint(0)
	expected := "123456789000"
	signVal := 1
	absIntRuneStr := "123456789000"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_04(t *testing.T) {

	nStr := "123456789"
	precision := uint(6)
	outPrecision := uint(0)
	expected := "123456789000000"
	signVal := 1
	absIntRuneStr := "123456789000000"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_05(t *testing.T) {

	nStr := "0"
	precision := uint(3)
	outPrecision := uint(0)
	expected := "0"
	signVal := 1
	absIntRuneStr := "0"
	absFracRuneStr := ""

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_ShiftPrecisionRight_06(t *testing.T) {

	nStr := "123456.789"
	precision := uint(0)
	outPrecision := uint(3)
	expected := "123456.789"
	signVal := 1
	absIntRuneStr := "123456"
	absFracRuneStr := "789"

	nsDto, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(nStr, precision)

	if err != nil {
		t.Errorf("Received error from nsu.ShiftPrecisionLeft(nStr, precision). nStr= '%v' precision= '%v'. Error= %v", nStr, precision, err)
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut='%v'. Instead, got %v.", expected, nsDto.NumStrOut)
	}

	if outPrecision != nsDto.Precision {
		t.Errorf("Expected precision='%v'. Instead, got %v.", outPrecision, nsDto.Precision)
	}

	if signVal != nsDto.SignVal {
		t.Errorf("Expected signVal='%v'. Instead, got %v.", signVal, nsDto.SignVal)
	}

	if !nsDto.IsValid {
		t.Errorf("Expected isValid='true'. Instead, got %v.", nsDto.IsValid)
	}

	if !nsDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigits='true'. Instead, got %v.", nsDto.HasNumericDigits)
	}

	s := string(nsDto.AbsIntRunes)

	if s != absIntRuneStr {
		t.Errorf("Expected AbsIntRunes='%v'. Instead, got %v.", absIntRuneStr, s)
	}

	s = string(nsDto.AbsFracRunes)

	if s != absFracRuneStr {
		t.Errorf("Expected AbsFracRunes='%v'. Instead, got %v", absFracRuneStr, s)
	}

}

func TestNumStrDto_SubtractNumStrs_01(t *testing.T) {
	nStr1 := "67.521"
	nStr2 := "-6"
	nStr3 := "73.521"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nResult, _ := nDto.ParseNumStr(nStr3)

	nDto, err := nDto.SubtractNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nDto.NumStrOut
	expected := nResult.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nDto.AbsIntRunes)
	iStr := string(nResult.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)
	fracStr := string(nResult.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nResult.SignVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nResult.IsFractionalValue, nDto.IsFractionalValue)
	}

	if nDto.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nResult.Precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_SubtractNumStrs_02(t *testing.T) {
	nStr1 := "-67.521"
	nStr2 := "6"
	nStr3 := "-73.521"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nResult, _ := nDto.ParseNumStr(nStr3)

	nDto, err := nDto.SubtractNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nDto.NumStrOut
	expected := nResult.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nDto.AbsIntRunes)
	iStr := string(nResult.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)
	fracStr := string(nResult.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nResult.SignVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nResult.IsFractionalValue, nDto.IsFractionalValue)
	}

	if nDto.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nResult.Precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_SubtractNumStrs_03(t *testing.T) {
	nStr1 := "67.521"
	nStr2 := "691.1"
	nStr3 := "-623.579"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nResult, _ := nDto.ParseNumStr(nStr3)

	nDto, err := nDto.SubtractNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nDto.NumStrOut
	expected := nResult.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nDto.AbsIntRunes)
	iStr := string(nResult.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)
	fracStr := string(nResult.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nResult.SignVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nResult.IsFractionalValue, nDto.IsFractionalValue)
	}

	if nDto.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nResult.Precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_SubtractNumStrs_04(t *testing.T) {

	nStr1 := "691.1"
	nStr2 := "67.521"
	nStr3 := "623.579"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nResult, _ := nDto.ParseNumStr(nStr3)

	nDto, err := nDto.SubtractNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nDto.NumStrOut
	expected := nResult.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nDto.AbsIntRunes)
	iStr := string(nResult.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)
	fracStr := string(nResult.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nResult.SignVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nResult.IsFractionalValue, nDto.IsFractionalValue)
	}

	if nDto.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nResult.Precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_SubtractNumStrs_05(t *testing.T) {

	nStr1 := "691.1"
	nStr2 := "0"
	nStr3 := "691.1"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nResult, _ := nDto.ParseNumStr(nStr3)

	nDto, err := nDto.SubtractNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nDto.NumStrOut
	expected := nResult.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nDto.AbsIntRunes)
	iStr := string(nResult.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)
	fracStr := string(nResult.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nResult.SignVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nResult.IsFractionalValue, nDto.IsFractionalValue)
	}

	if nDto.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nResult.Precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}

func TestNumStrDto_SubtractNumStrs_06(t *testing.T) {

	nStr1 := "0"
	nStr2 := "691.1"
	nStr3 := "-691.1"

	nDto := NumStrDto{}.New()

	n1, _ := nDto.ParseNumStr(nStr1)
	n2, _ := nDto.ParseNumStr(nStr2)
	nResult, _ := nDto.ParseNumStr(nStr3)

	nDto, err := nDto.SubtractNumStrs(n1, n2)

	if err != nil {
		t.Errorf("nDto.AddNumStrs(n1, n2) returned an error. Error= %v", err)
	}

	s := nDto.NumStrOut
	expected := nResult.NumStrOut

	if s != expected {
		t.Errorf("Expected NumStrOut = '%v'. Instead, got %v", expected, s)
	}

	s = string(nDto.AbsIntRunes)
	iStr := string(nResult.AbsIntRunes)

	if iStr != s {
		t.Errorf("Expected AbsIntRunes = '%v'. Instead, got %v", iStr, s)

	}

	s = string(nDto.AbsFracRunes)
	fracStr := string(nResult.AbsFracRunes)

	if fracStr != s {
		t.Errorf("Expected AbsFracRunes = '%v'. Instead, got %v", fracStr, s)
	}

	if nDto.SignVal != nResult.SignVal {
		t.Errorf("Expected SignVal= '%v'. Instead, got %v", nResult.SignVal, nDto.SignVal)
	}

	if !nDto.HasNumericDigits {
		t.Errorf("Expected HasNumericDigist= 'true'. Instead, got %v", nDto.HasNumericDigits)
	}

	if nDto.IsFractionalValue != nResult.IsFractionalValue {
		t.Errorf("Expected IsFractionalValue= '%v'. Instead, got '%v'", nResult.IsFractionalValue, nDto.IsFractionalValue)
	}

	if nDto.Precision != nResult.Precision {
		t.Errorf("Expected precision= '%v'. Instead, got %v", nResult.Precision, nDto.Precision)

	}

	if !nDto.IsValid {
		t.Errorf("Expected IsValid= 'true'. Instead, got %v", nDto.IsValid)
	}

}
