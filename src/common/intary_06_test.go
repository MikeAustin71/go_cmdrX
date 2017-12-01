package common

import (
	"math/big"
	"testing"
)

func TestIntAry_OptimizeIntArrayLen_07(t *testing.T) {
	nStr1 := "0"
	expected := "0"
	ePrecision := 0
	eLen := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true)
	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v'. Instead, ia.GetPrecision()= '%v' .", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}

func TestIntAry_OptimizeIntArrayLen_08(t *testing.T) {
	nStr1 := "-00579.000000000"
	expected := "-579"
	ePrecision := 0
	eLen := 3

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true)
	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v'. Instead, ia.GetPrecision()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}

func TestIntAry_OptimizeIntArrayLen_09(t *testing.T) {
	nStr1 := "-00579.123000000"
	expected := "-579.123"
	ePrecision := 3
	eLen := 6

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(true)
	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.GetNumStr()= '%v'", expected, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v'. Instead, ia.GetPrecision()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eLen != ia.GetIntAryLength() {
		t.Errorf("Expected ia IntAryLen= '%v' .   Instead, ia IntAryLen= '%v' .", eLen, ia.GetIntAryLength())
	}

}


func TestIntAry_Pow_01(t *testing.T) {

	nStr1 := "3"
	power := 3
	eStr := "27"
	eSignVal := 1
	ePrecision := 0

	ia := IntAry{}.New()

	ia.SetIntAryWithNumStr(nStr1)

	err := ia.Pow(power, ePrecision, -1)

	if err != nil {
		t.Errorf("Error returned from ia.Pow(power). Error= %v", err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr() == %v . Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v . Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetSign() == %v . Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_Pow_02(t *testing.T) {

	nStr1 := "-3"
	power := 3
	eStr := "-27"
	eSignVal := -1
	ePrecision := 0

	ia := IntAry{}.New()

	ia.SetIntAryWithNumStr(nStr1)

	err := ia.Pow(power, ePrecision, -1)

	if err != nil {
		t.Errorf("Error returned from ia.Pow(power). Error= %v", err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr() == %v . Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v . Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetSign() == %v . Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_Pow_03(t *testing.T) {

	nStr1 := "-3.97"
	power := 4
	eStr := "248.40596881"
	eSignVal := 1
	ePrecision := 8

	ia := IntAry{}.New()

	ia.SetIntAryWithNumStr(nStr1)

	err := ia.Pow(power, ePrecision, -1)

	if err != nil {
		t.Errorf("Error returned from ia.Pow(power). Error= %v", err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr() == %v . Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v . Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetSign() == %v . Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_Pow_04(t *testing.T) {

	nStr1 := "-3.97"
	power := 3
	eStr := "-62.570773"
	eSignVal := -1
	ePrecision := 6

	ia := IntAry{}.New()

	ia.SetIntAryWithNumStr(nStr1)

	err := ia.Pow(power, ePrecision, -1)

	if err != nil {
		t.Errorf("Error returned from ia.Pow(power). Error= %v", err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr() == %v . Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v . Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetSign() == %v . Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_Pow_05(t *testing.T) {

	nStr1 := "4"
	power := -3
	eStr := "0.015625"
	eSignVal := 1
	maxPrecision:= 10
	ePrecision := 6

	ia := IntAry{}.New()

	ia.SetIntAryWithNumStr(nStr1)

	err := ia.Pow(power, maxPrecision, -1)

	if err != nil {
		t.Errorf("Error returned from ia.Pow(power). Error= %v", err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr() == %v . Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v . Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v . Instead ia.GetPrecision() == %v", ePrecision, ia.GetPrecision())
	}

}


func TestIntAry_Pow_06(t *testing.T) {

	nStr1 := "92"
	power := -8
	eStr := "0.00000000000000019484864106545884795863650193677353127108077765624198799946139576907821793952688449765949765741908569104624237348062845911009040417592239998725345624987067642919240414862531015586495108191971721439283981366869555221526775365563032086453618016763974266053775613202784420219175865536764500972234932644532359373163382073933684236312884196"
	eSignVal := 1
	maxPrecision:= 350
	ePrecision := 350

	ia := IntAry{}.New()

	ia.SetIntAryWithNumStr(nStr1)

	err := ia.Pow(power, maxPrecision, -1)

	if err != nil {
		t.Errorf("Error returned from ia.Pow(power). Error= %v", err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr() == %v . Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v . Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetSign() == %v . Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}


func TestIntAry_ResetFromBackUp_01(t *testing.T) {
	nStr1 := "99.4564"
	nStr2 := "-5034.123"
	eSign := 1
	ePrecsion := 4
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.CopyToBackUp()
	ia.SetIntAryWithNumStr(nStr2)

	if nStr2 != ia.GetNumStr() {
		t.Errorf("Expected ia.NumStr2= '%v' .  Instead, ia.NumStr2= '%v' .", nStr2, ia.GetNumStr())
	}

	ia.ResetFromBackUp()

	if nStr1 != ia.GetNumStr() {
		t.Errorf("After Reset - Expected ia.NumStr1= '%v' .  Instead, ia.NumStr1= '%v' .", nStr1, ia.GetNumStr())
	}

	if ia.GetSign() != eSign {
		t.Errorf("After Reset, Expected ia.GetSign()= '%v'. Instead, ia.GetSign()= '%v' ", eSign, ia.GetSign())
	}

	if ia.GetPrecision() != ePrecsion {
		t.Errorf("After Reset, Expected ia.GetPrecision()= '%v'. Instead, ia.GetPrecision()= '%v' .", ePrecsion, ia.GetPrecision())
	}

}

func TestIntAry_RoundToPrecision_01(t *testing.T) {
	nStr1 := "999.9952"
	expected := "1000.00"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if ia.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.GetPrecision())
	}

}

func TestIntAry_RoundToPrecision_02(t *testing.T) {
	nStr1 := "-999.9952"
	expected := "-1000.00"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if ia.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.GetPrecision())
	}

}

func TestIntAry_RoundToPrecision_03(t *testing.T) {
	nStr1 := "352.954"
	expected := "352.95"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if ia.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.GetPrecision())
	}

}

func TestIntAry_RoundToPrecision_04(t *testing.T) {
	nStr1 := "-352.954"
	expected := "-352.95"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if ia.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.GetPrecision())
	}

}

func TestIntAry_RoundToPrecision_05(t *testing.T) {
	nStr1 := "-352.954"
	expected := "-352.95"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if ia.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.GetPrecision())
	}

}

func TestIntAry_RoundToPrecision_06(t *testing.T) {
	nStr1 := "999.99"
	expected := "999.99"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if ia.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.GetPrecision())
	}

}

func TestIntAry_RoundToPrecision_07(t *testing.T) {
	nStr1 := "999.99"
	expected := "999.99000"
	precision := 5
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if ia.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.GetPrecision())
	}

}

func TestIntAry_RoundToPrecision_08(t *testing.T) {
	nStr1 := "0.000"
	expected := "0.00"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if ia.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.GetPrecision())
	}

}

func TestIntAry_RoundToPrecision_09(t *testing.T) {
	nStr1 := "999.995"
	expected := "999.995"
	precision := 3
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if ia.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.GetPrecision())
	}

}

func TestIntAry_RoundToPrecision_10(t *testing.T) {
	nStr1 := "51.821810080312709972402243542716917"
	expected := "51.821810080312709972402243542717"
	precision := 30
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.RoundToPrecision(precision)
	if err != nil {
		t.Errorf("Received Error from ia.RoundToPrecision(precision). precision '%v' Error:= %v", precision, err)
	}

	s := ia.GetNumStr()
	if expected != s {
		t.Errorf("Error. Expected numStr= '%v'. Instead, got numStr='%v'\n", expected, s)
	}

	if ia.GetPrecision() != precision {
		t.Errorf("Error. Expected precision= '%v'. Instead, got precision='%v'\n", precision, ia.GetPrecision())
	}

}

func TestIntAry_SetAbsoluteValueThis_01(t *testing.T) {

	nStr1 := "-987.652"
	eStr := "987.652"

	ia, _ := IntAry{}.NewNumStr(nStr1)

	if nStr1 != ia.GetNumStr() {
		t.Errorf("Expected intialized NumStr= %v .  Instead, NumStr= %v  .", nStr1, ia.GetNumStr())
	}

	ia.SetAbsoluteValueThis()

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected Absolute Value NumStr= %v .  Instead, NumStr= %v  .", eStr, ia.GetNumStr())
	}
}

func TestIntAry_SetAbsoluteValueThis_02(t *testing.T) {

	nStr1 := "987.652"
	eStr := "987.652"

	ia, _ := IntAry{}.NewNumStr(nStr1)

	if nStr1 != ia.GetNumStr() {
		t.Errorf("Expected intialized NumStr= %v .  Instead, NumStr= %v  .", nStr1, ia.GetNumStr())
	}

	ia.SetAbsoluteValueThis()

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected Absolute Value NumStr= %v .  Instead, NumStr= %v  .", eStr, ia.GetNumStr())
	}
}

func TestIntAry_SetCurrencySymbol_01(t *testing.T) {

	ia, _ := IntAry{}.NewNumStr("50.37")

	var poundSym  rune

	poundSym = '\U000000a3'

	ia.SetCurrencySymbol(poundSym)

	curSymbol := ia.GetCurrencySymbol()

	if poundSym != curSymbol {
		t.Errorf("Error: Expected Currency Symbol= '%v'. Instead, received Currency Symbol= '%v'", poundSym, curSymbol)
	}

}

func TestIntAry_SetDecimalSeparator_01(t *testing.T) {

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

func TestIntAry_SetEqualArrayLengths_01(t *testing.T) {
	nStr1 := "3536.123456"
	eNStr1 := "3536.123456"
	ePrecision1 := 6
	eSignVal1 := 1

	nStr2 := "12.14"
	eNStr2 := "0012.140000"
	ePrecision2 := 6
	eSignVal2 := 1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia1.SetEqualArrayLengths(&ia2)

	if ia1.GetNumStr() != eNStr1 {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' . Instead, ia1.GetNumStr()= '%v' .", eNStr1, ia1.GetNumStr())
	}

	if ia1.GetSign() != eSignVal1 {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", eSignVal1, ia1.GetSign())
	}

	if ia1.GetPrecision() != ePrecision1 {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", ePrecision1, ia1.GetPrecision())
	}

	if ia2.GetNumStr() != eNStr2 {
		t.Errorf("Error - Expected ia2.GetNumStr()= '%v' . Instead, ia2.GetNumStr()= '%v' .", eNStr2, ia2.GetNumStr())
	}

	if ia2.GetSign() != eSignVal2 {
		t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", eSignVal2, ia2.GetSign())
	}

	if ia2.GetPrecision() != ePrecision2 {
		t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", ePrecision2, ia2.GetPrecision())
	}

}

func TestIntAry_SetEqualArrayLengths_02(t *testing.T) {
	nStr1 := "12.14"
	eNStr1 := "0012.140000"
	ePrecision1 := 6
	eSignVal1 := 1

	nStr2 := "3536.123456"
	eNStr2 := "3536.123456"
	ePrecision2 := 6
	eSignVal2 := 1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia1.SetEqualArrayLengths(&ia2)

	if ia1.GetNumStr() != eNStr1 {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' . Instead, ia1.GetNumStr()= '%v' .", eNStr1, ia1.GetNumStr())
	}

	if ia1.GetSign() != eSignVal1 {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", eSignVal1, ia1.GetSign())
	}

	if ia1.GetPrecision() != ePrecision1 {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", ePrecision1, ia1.GetPrecision())
	}

	if ia2.GetNumStr() != eNStr2 {
		t.Errorf("Error - Expected ia2.GetNumStr()= '%v' . Instead, ia2.GetNumStr()= '%v' .", eNStr2, ia2.GetNumStr())
	}

	if ia2.GetSign() != eSignVal2 {
		t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", eSignVal2, ia2.GetSign())
	}

	if ia2.GetPrecision() != ePrecision2 {
		t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", ePrecision2, ia2.GetPrecision())
	}

}

func TestIntAry_SetEqualArrayLengths_03(t *testing.T) {
	nStr1 := "3536.123456"
	eNStr1 := "3536.123456"
	ePrecision1 := 6
	eSignVal1 := 1

	nStr2 := "-12.14"
	eNStr2 := "-0012.140000"
	ePrecision2 := 6
	eSignVal2 := -1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia1.SetEqualArrayLengths(&ia2)

	if ia1.GetNumStr() != eNStr1 {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' . Instead, ia1.GetNumStr()= '%v' .", eNStr1, ia1.GetNumStr())
	}

	if ia1.GetSign() != eSignVal1 {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", eSignVal1, ia1.GetSign())
	}

	if ia1.GetPrecision() != ePrecision1 {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", ePrecision1, ia1.GetPrecision())
	}

	if ia2.GetNumStr() != eNStr2 {
		t.Errorf("Error - Expected ia2.GetNumStr()= '%v' . Instead, ia2.GetNumStr()= '%v' .", eNStr2, ia2.GetNumStr())
	}

	if ia2.GetSign() != eSignVal2 {
		t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", eSignVal2, ia2.GetSign())
	}

	if ia2.GetPrecision() != ePrecision2 {
		t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", ePrecision2, ia2.GetPrecision())
	}

}

func TestIntAry_SetEqualArrayLengths_04(t *testing.T) {
	nStr1 := "-12.14"
	eNStr1 := "-0012.140000"
	ePrecision1 := 6
	eSignVal1 := -1

	nStr2 := "3536.123456"
	eNStr2 := "3536.123456"
	ePrecision2 := 6
	eSignVal2 := 1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia1.SetEqualArrayLengths(&ia2)

	if ia1.GetNumStr() != eNStr1 {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' . Instead, ia1.GetNumStr()= '%v' .", eNStr1, ia1.GetNumStr())
	}

	if ia1.GetSign() != eSignVal1 {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", eSignVal1, ia1.GetSign())
	}

	if ia1.GetPrecision() != ePrecision1 {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", ePrecision1, ia1.GetPrecision())
	}

	if ia2.GetNumStr() != eNStr2 {
		t.Errorf("Error - Expected ia2.GetNumStr()= '%v' . Instead, ia2.GetNumStr()= '%v' .", eNStr2, ia2.GetNumStr())
	}

	if ia2.GetSign() != eSignVal2 {
		t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", eSignVal2, ia2.GetSign())
	}

	if ia2.GetPrecision() != ePrecision2 {
		t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", ePrecision2, ia2.GetPrecision())
	}

}

func TestIntAry_SetEqualArrayLengths_05(t *testing.T) {
	nStr1 := "-123456.143456"
	eNStr1 := "-123456.143456"
	ePrecision1 := 6
	eSignVal1 := -1

	nStr2 := "353678.123456"
	eNStr2 := "353678.123456"
	ePrecision2 := 6
	eSignVal2 := 1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia1.SetEqualArrayLengths(&ia2)

	if ia1.GetNumStr() != eNStr1 {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' . Instead, ia1.GetNumStr()= '%v' .", eNStr1, ia1.GetNumStr())
	}

	if ia1.GetSign() != eSignVal1 {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", eSignVal1, ia1.GetSign())
	}

	if ia1.GetPrecision() != ePrecision1 {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", ePrecision1, ia1.GetPrecision())
	}

	if ia2.GetNumStr() != eNStr2 {
		t.Errorf("Error - Expected ia2.GetNumStr()= '%v' . Instead, ia2.GetNumStr()= '%v' .", eNStr2, ia2.GetNumStr())
	}

	if ia2.GetSign() != eSignVal2 {
		t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", eSignVal2, ia2.GetSign())
	}

	if ia2.GetPrecision() != ePrecision2 {
		t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", ePrecision2, ia2.GetPrecision())
	}

}

func TestIntAry_SetEqualArrayLengths_06(t *testing.T) {
	nStr1 := "0.00"
	eNStr1 := "0.00"
	ePrecision1 := 2
	eSignVal1 := 1

	nStr2 := "0"
	eNStr2 := "0.00"
	ePrecision2 := 2
	eSignVal2 := 1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia1.SetEqualArrayLengths(&ia2)

	if ia1.GetNumStr() != eNStr1 {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' . Instead, ia1.GetNumStr()= '%v' .", eNStr1, ia1.GetNumStr())
	}

	if ia1.GetSign() != eSignVal1 {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", eSignVal1, ia1.GetSign())
	}

	if ia1.GetPrecision() != ePrecision1 {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", ePrecision1, ia1.GetPrecision())
	}

	if ia2.GetNumStr() != eNStr2 {
		t.Errorf("Error - Expected ia2.GetNumStr()= '%v' . Instead, ia2.GetNumStr()= '%v' .", eNStr2, ia2.GetNumStr())
	}

	if ia2.GetSign() != eSignVal2 {
		t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", eSignVal2, ia2.GetSign())
	}

	if ia2.GetPrecision() != ePrecision2 {
		t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", ePrecision2, ia2.GetPrecision())
	}

}

func TestIntAry_SetEqualArrayLengths_07(t *testing.T) {
	nStr1 := "0"
	eNStr1 := "0.00"
	ePrecision1 := 2
	eSignVal1 := 1

	nStr2 := "0.00"
	eNStr2 := "0.00"
	ePrecision2 := 2
	eSignVal2 := 1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia1.SetEqualArrayLengths(&ia2)

	if ia1.GetNumStr() != eNStr1 {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' . Instead, ia1.GetNumStr()= '%v' .", eNStr1, ia1.GetNumStr())
	}

	if ia1.GetSign() != eSignVal1 {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", eSignVal1, ia1.GetSign())
	}

	if ia1.GetPrecision() != ePrecision1 {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' . Instead, ia1.GetSign()= '%v' .", ePrecision1, ia1.GetPrecision())
	}

	if ia2.GetNumStr() != eNStr2 {
		t.Errorf("Error - Expected ia2.GetNumStr()= '%v' . Instead, ia2.GetNumStr()= '%v' .", eNStr2, ia2.GetNumStr())
	}

	if ia2.GetSign() != eSignVal2 {
		t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", eSignVal2, ia2.GetSign())
	}

	if ia2.GetPrecision() != ePrecision2 {
		t.Errorf("Error - Expected ia2.GetSign()= '%v' . Instead, ia2.GetSign()= '%v' .", ePrecision2, ia2.GetPrecision())
	}

}

func TestIntAry_SetIntAryWithFloat_01(t *testing.T) {
	num := float32(475.895)
	expectedNumStr := "475.895"
	precision := 3
	signVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithFloat32(num, precision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithFloat32(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
	}

	if expectedNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' .  Instead, ia.GetPrecision()= '%v' .", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
	}
}

func TestIntAry_SetIntAryWithFloat_02(t *testing.T) {
	num := float32(475.895)
	expectedNumStr := "475.8950"
	precision := 4
	signVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithFloat32(num, precision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithFloat32(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
	}

	if expectedNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' .  Instead, ia.GetPrecision()= '%v' .", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
	}
}

func TestIntAry_SetIntAryWithFloat_03(t *testing.T) {
	num := float32(475.895600)
	expectedNumStr := "475.8956"
	precision := 4
	signVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithFloat32(num, -1)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithFloat32(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
	}

	if expectedNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' .  Instead, ia.GetPrecision()= '%v' .", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
	}
}

func TestIntAry_SetIntAryWithFloat_04(t *testing.T) {
	num := float32(-475.8956)
	expectedNumStr := "-475.8956"
	precision := 4
	signVal := -1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithFloat32(num, -1)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithFloat32(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
	}

	if expectedNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' .  Instead, ia.GetPrecision()= '%v' .", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
	}
}

func TestIntAry_SetIntAryWithFloat64_01(t *testing.T) {
	num := float64(4756625.895)
	expectedNumStr := "4756625.895"
	precision := 3
	signVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithFloat64(num, precision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithFloat64(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
	}

	if expectedNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' .  Instead, ia.GetPrecision()= '%v' .", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
	}
}

func TestIntAry_SetIntAryWithFloat64_02(t *testing.T) {
	num := float64(4756625.895)
	expectedNumStr := "4756625.8950"
	precision := 4
	signVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithFloat64(num, precision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithFloat64(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
	}

	if expectedNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' .  Instead, ia.GetPrecision()= '%v' .", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
	}
}

func TestIntAry_SetIntAryWithFloat64_03(t *testing.T) {
	num := float64(475.895600)
	expectedNumStr := "475.8956"
	precision := 4
	signVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithFloat64(num, -1)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithFloat64(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
	}

	if expectedNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' .  Instead, ia.GetPrecision()= '%v' .", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
	}
}

func TestIntAry_SetIntAryWithFloat64_04(t *testing.T) {
	num := float64(-475.8956)
	expectedNumStr := "-475.8956"
	precision := 4
	signVal := -1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithFloat64(num, -1)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithFloat64(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
	}

	if expectedNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' .  Instead, ia.GetPrecision()= '%v' .", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
	}
}

func TestIntAry_SetIntAryWithFloatBig_01(t *testing.T) {
	num := big.NewFloat(float64(4756625.895))
	expectedNumStr := "4756625.895"
	precision := 3
	signVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithFloatBig(num, precision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithFloatBig(num). num= %v Error= %v", num, err)
	}

	if expectedNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' .  Instead, ia.GetPrecision()= '%v' .", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
	}
}

func TestIntAry_SetIntAryWithFloatBig_02(t *testing.T) {
	num := big.NewFloat(float64(4756625.8950))
	expectedNumStr := "4756625.895"
	precision := 3
	signVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithFloatBig(num, precision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithFloatBig(num). num= %v  Error= %v", num, err)
	}

	if expectedNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' .  Instead, ia.GetPrecision()= '%v' .", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
	}
}

func TestIntAry_SetIntAryWithFloatBig_03(t *testing.T) {
	num := big.NewFloat(float64(475.895600))
	expectedNumStr := "475.8956"
	precision := 4
	signVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithFloatBig(num, 4)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithFloatBig(num). num= %v  Error= %v", num.String(), err)
	}

	if expectedNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' .  Instead, ia.GetPrecision()= '%v' .", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
	}
}

func TestIntAry_SetIntAryWithFloatBig_04(t *testing.T) {

	num :=  big.NewFloat(float64(-475.8956))
	expectedNumStr := "-475.8956"
	precision := 4
	signVal := -1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithFloatBig(num, precision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithFloatBig(num). num= %v Error= %v", num.String(), err)
	}

	if expectedNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' .  Instead, ia.GetPrecision()= '%v' .", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
	}
}

func TestIntAry_SetIntAryWithFloatBig_05(t *testing.T) {
	num := big.NewFloat(float64(475.895600))
	expectedNumStr := "475.8956"
	precision := 4
	signVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithFloatBig(num, -1)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithFloatBig(num). num= %v  Error= %v", num.String(), err)
	}

	if expectedNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' .  Instead, ia.GetPrecision()= '%v' .", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
	}
}


func TestIntAry_SetIntAryWithFloatBig_06(t *testing.T) {
	num := big.NewFloat(float64(4756625.8957))
	expectedNumStr := "4756625.896"
	precision := 3
	signVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithFloatBig(num, precision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithFloatBig(num). num= %v Error= %v", num, err)
	}

	if expectedNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' .", expectedNumStr, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' .  Instead, ia.GetPrecision()= '%v' .", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' .  Instead, ia.GetSign()= '%v' .", signVal, ia.GetSign())
	}
}


func TestIntAry_SetIntAryWithBigInt_01(t *testing.T) {

	num := big.NewInt(int64(123456789))

	eNumStr := "123456.789"
	ePrecision := uint(3)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithBigInt(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithBigInt(num, ePrecision). num= %v  ePrecision= %v  eSignVal= %v", num.String(), ePrecision, eSignVal)
	}

	if eNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
	}

	if int(ePrecision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' . Instead, ia.GetPrecision()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
	}

}

func TestIntAry_SetIntAryWithBigInt_02(t *testing.T) {

	num := big.NewInt(int64(-123456789))

	eNumStr := "-123456.789"
	ePrecision := uint(3)
	eSignVal := -1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithBigInt(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithBigInt(num, ePrecision). num= %v  ePrecision= %v  .", num.String(), ePrecision)
	}

	if eNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
	}

	if int(ePrecision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' . Instead, ia.GetPrecision()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
	}

}

func TestIntAry_SetIntAryWithBigInt_03(t *testing.T) {

	num := big.NewInt(int64(0))

	eNumStr := "0.000"
	ePrecision := uint(3)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithBigInt(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithBigInt(num, ePrecision). num= %v  ePrecision= %v  .", num.String(), ePrecision)
	}

	if eNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
	}

	if int(ePrecision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' . Instead, ia.GetPrecision()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
	}

}

func TestIntAry_SetIntAryWithInt_01(t *testing.T) {

	num := 123456789

	eNumStr := "123456.789"
	ePrecision := uint(3)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt(num, ePrecision). num= %v  ePrecision= %v .", num, ePrecision)
	}

	if eNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
	}

	if int(ePrecision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' . Instead, ia.GetPrecision()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
	}

}

func TestIntAry_SetIntAryWithInt_02(t *testing.T) {

	num := -123456789

	eNumStr := "-12345.6789"
	ePrecision := uint(4)
	eSignVal := -1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt(num, ePrecision). num= %v  ePrecision= %v  .", num, ePrecision)
	}

	if eNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
	}

	if int(ePrecision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' . Instead, ia.GetPrecision()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
	}

}

func TestIntAry_SetIntAryWithInt_03(t *testing.T) {

	num := 0

	eNumStr := "0.0000"
	ePrecision := uint(4)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt(num, ePrecision). num= %v  ePrecision= %v  .", num, ePrecision)
	}

	if eNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
	}

	if int(ePrecision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' . Instead, ia.GetPrecision()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
	}

}

func TestIntAry_SetIntAryWithInt_04(t *testing.T) {

	num := 32

	eNumStr := "0.0032"
	ePrecision := uint(4)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt(num, ePrecision). num= %v  ePrecision= %v  .", num, ePrecision)
	}

	if eNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
	}

	if int(ePrecision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' . Instead, ia.GetPrecision()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
	}

}

func TestIntAry_SetIntAryWithInt_05(t *testing.T) {

	num := -32

	eNumStr := "-32"
	ePrecision := uint(0)
	eSignVal := -1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt(num, ePrecision). num= %v  ePrecision= %v  .", num, ePrecision)
	}

	if eNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
	}

	if int(ePrecision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' . Instead, ia.GetPrecision()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
	}

}

func TestIntAry_SetIntAryWithInt_06(t *testing.T) {

	num := 32

	eNumStr := "0.32"
	ePrecision := uint(2)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt(num, ePrecision). num= %v  ePrecision= %v  .", num, ePrecision)
	}

	if eNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
	}

	if int(ePrecision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision()= '%v' . Instead, ia.GetPrecision()= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v' . Instead, ia.GetSign()= '%v'", eSignVal, ia.GetSign())
	}

}
