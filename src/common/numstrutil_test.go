package common

import (
	"fmt"

	"testing"
)

/*
	These tests are associated with numstrutility.go.

	The source code repository for numstrutility.go is located at :
					https://github.com/MikeAustin71/numstrutility.git

 */

func TestNumStrUtility_DlimDecStr(t *testing.T) {

	ns := NumStrUtility{}
	n := "1234567890"
	expected := "1,234,567,890"

	result := ns.DlimDecCurrStr(n, ',', '.', '$')

	if result != expected {
		t.Errorf("Expected result = %v; instead got: %v", expected, result)
	}

}

func TestNumStrUtility_DlimDecStr_With_DecimalCurrency(t *testing.T) {
	ns := NumStrUtility{}
	n := "$1234567890.25"

	result := ns.DlimDecCurrStr(n, ',', '.', '$')
	expected := "$1,234,567,890.25"

	if result != expected {
		t.Errorf("Expected result = %v; instead got: %v", expected, result)
	}

}

func TestNumStrUtility_DnumStr(t *testing.T) {
	ns := NumStrUtility{}
	n := "1234567890"
	expected := "1,234,567,890"

	result := ns.DnumStr(n, ',')

	if result != expected {
		t.Errorf("Expected result = %v; instead got: %v", expected, result)
	}

}

func TestNumStrUtility_DNumI64(t *testing.T) {
	ns := NumStrUtility{}
	n := int64(1234567890)
	expected := "1,234,567,890"

	result := ns.DLimI64(n, ',')

	if result != expected {
		t.Errorf("Expected result = %v; instead got: %v", expected, result)
	}

}

func TestNumStrUtility_DNumI64_EvenThousands(t *testing.T) {
	ns := NumStrUtility{}
	n := int64(123456)
	expected := "123,456"

	result := ns.DLimI64(n, ',')

	if result != expected {
		t.Errorf("Expected result = %v; instead got: %v", expected, result)
	}

}

func TestNumStrUtility_DLimInt(t *testing.T) {
	ns := NumStrUtility{}
	n := int(1234567)
	expected := "1,234,567"

	result := ns.DLimInt(n, ',')

	if result != expected {
		t.Errorf("Expected result = %v; instead got: %v", expected, result)
	}
}

func TestNumStrUtility_ConvertStrToInt64_01(t *testing.T) {
	numStr := "-12314617914"

	ns := NumStrUtility{}

	result, err := ns.ConvertNumStrToInt64(numStr)

	if err != nil {
		t.Errorf("Received Error from ns.ConvertNumStrToInt64(%v).", numStr)
	}

	s, err := ns.ConvertInt64ToStr(result)

	if err != nil {
		t.Errorf("Received Error from ns.ConvertInt64ToStr(%v) ", result)
	}

	if s != numStr {
		t.Errorf("Error. Expected value %v. Instead, got %v", numStr, s)
	}
}

func TestNumStrUtility_ConvertStrToInt64_02(t *testing.T) {
	numStr := "+12314617914"

	ns := NumStrUtility{}

	result, err := ns.ConvertNumStrToInt64(numStr)

	if err != nil {
		t.Errorf("Received Error from ns.ConvertNumStrToInt64(%v).", numStr)
	}

	s, err := ns.ConvertInt64ToStr(result)

	if err != nil {
		t.Errorf("Received Error from ns.ConvertInt64ToStr(%v) ", result)
	}

	s = "+" + s

	if s != numStr {
		t.Errorf("Error. Expected value %v. Instead, got %v", numStr, s)
	}
}

func TestNumStrUtility_ConvertStrToInt64_03(t *testing.T) {

	numStr := "12314617914"

	ns := NumStrUtility{}

	result, err := ns.ConvertNumStrToInt64(numStr)

	if err != nil {
		t.Errorf("Received Error from ns.ConvertNumStrToInt64(%v).", numStr)
	}

	s, err := ns.ConvertInt64ToStr(result)

	if err != nil {
		t.Errorf("Received Error from ns.ConvertInt64ToStr(%v) ", result)
	}

	if s != numStr {
		t.Errorf("Error. Expected value %v. Instead, got %v", numStr, s)
	}
}

func TestNumStrUtility_ConvertStrToIntNumStr(t *testing.T) {
	str := "-12,314,617,914"
	expected := "-12314617914"
	ns := NumStrUtility{}
	ns.ThousandsSeparator = ','

	numStr := ns.ConvertStrToIntNumStr(str)

	if expected != numStr {
		t.Errorf("Expected number string %v. Instead got %v", expected, numStr)
	}

}

func TestNumStrUtility_SetCountryAndCurrency(t *testing.T) {

	ns := NumStrUtility{}

	err := ns.SetCountryAndCurrency("United States")

	if err != nil {
		t.Errorf("Received Error from ns.SetCountryAndCurrency(\"United States\"). Error: %v ", err.Error())
	}

	if ns.CurrencySymbol != '$' {
		t.Errorf("Expected '$' currency symbol, instead got %v.", ns.CurrencySymbol)
	}

	if ns.Nation != "United States" {
		t.Errorf("Expected Nation = 'United States'. Instead, got %v.", ns.Nation)
	}

	if ns.ThousandsSeparator != ',' {
		t.Errorf("Expected thousands separator = ','. Instead, got %v.", ns.ThousandsSeparator)
	}

	if ns.DecimalSeparator != '.' {
		t.Errorf("Expected decimal separator = ','. Instead, got %v.", ns.DecimalSeparator)
	}

}

func TestNumStrUtility_ConvertInt64ToFractionalValue(t *testing.T) {

	ns := NumStrUtility{}

	i64 := int64(123456)

	f64, err := ns.ConvertInt64ToFractionalValue(i64)

	if err != nil {
		t.Errorf("Error received from ns.ConvertInt64ToFractionalValue(i64). Error: %v", err.Error())
		return
	}

	result := fmt.Sprintf("%v", f64)
	expected := "0.123456"

	if result != expected {
		t.Errorf("Error on ConvertInt64ToFractionalValue(). Expected %v. Instead, got %v.", expected, result)
	}
}

func TestNumStrUtility_ParseNumString_01(t *testing.T) {

	rawStr := "123456.654321"
	expected := "123456.654321"
	ns := NumStrUtility{}

	nsDto, err := ns.ParseNumString(rawStr)

	if err != nil {
		t.Errorf("Error from ns.ParseNumString(rawStr). rawStr= '%v'. Error= %v", rawStr, err.Error())
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut= '%v'. Instead, got %v", expected, nsDto.NumStrOut)
	}

	if nsDto.SignVal != 1 {
		t.Errorf("Expected signVal=1. Instead, got %v", nsDto.SignVal)
	}

	if nsDto.Precision != 6 {
		t.Errorf("Expected precision=6. Instead, got %v", nsDto.Precision)
	}

	if nsDto.HasNumericDigits == false {
		t.Errorf("Expected HasNumericDigits=true. Instead, got %v", nsDto.HasNumericDigits)
	}

	if nsDto.IsFractionalValue == false {
		t.Errorf("Expected IsFractionalValue=true. Instead, got %v", nsDto.IsFractionalValue)
	}

}

func TestNumStrUtility_ParseNumString_02(t *testing.T) {

	rawStr := "-123456.654321"
	expected := "-123456.654321"
	ns := NumStrUtility{}

	nsDto, err := ns.ParseNumString(rawStr)

	if err != nil {
		t.Errorf("Error from ns.ParseNumString(rawStr). rawStr= '%v'. Error= %v", rawStr, err.Error())
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut= '%v'. Instead, got %v", expected, nsDto.NumStrOut)
	}

	if nsDto.SignVal != -1 {
		t.Errorf("Expected signVal= -1. Instead, got %v", nsDto.SignVal)
	}

	if nsDto.Precision != 6 {
		t.Errorf("Expected precision=6. Instead, got %v", nsDto.Precision)
	}

	if nsDto.HasNumericDigits == false {
		t.Errorf("Expected HasNumericDigits=true. Instead, got %v", nsDto.HasNumericDigits)
	}

	if nsDto.IsFractionalValue == false {
		t.Errorf("Expected IsFractionalValue=true. Instead, got %v", nsDto.IsFractionalValue)
	}

}

func TestNumStrUtility_ParseNumString_03(t *testing.T) {
	rawStr := "Nothing"
	expected := "0"
	ns := NumStrUtility{}

	nsDto, err := ns.ParseNumString(rawStr)

	if err != nil {
		t.Errorf("Error from ns.ParseNumString(rawStr). rawStr= '%v'. Error= %v", rawStr, err.Error())
	}

	if nsDto.NumStrOut != expected {
		t.Errorf("Expected NumStrOut= '%v'. Instead, got %v", expected, nsDto.NumStrOut)
	}

	if nsDto.HasNumericDigits == false {
		t.Errorf("Expected HasNumericDigits=true. Instead, got %v", nsDto.HasNumericDigits)
	}

}