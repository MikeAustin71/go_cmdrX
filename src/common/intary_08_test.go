package common

import (
	"testing"
)


func TestIntAry_SetSignificantDigitIdxs_01(t *testing.T) {
	nStr := ".7770"
	eAryLen := 5
	eNumStr := "0.7770"
	eIntegerLen := 1
	eSigIntegerLen := 1
	eSigFractionLen := 3
	eIsZeroValue := false
	eIsIntegerZeroValue := true
	eFirstDigitIdx := 0
	eLastDigitIdx := 3
	eSignVal := 1
	ePrecision := 4

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned from SetIntAryWithNumStr(nStr). numStr= %v ", eNumStr)
	}

	iaStats := ia.GetIntAryStats()

	if eFirstDigitIdx != iaStats.FirstDigitIdx {
		t.Errorf("Expected ia FirstDigitIdx= '%v' .  Instead, ia FirstDigitIdx= '%v' .", eFirstDigitIdx, iaStats.FirstDigitIdx)
	}

	if eLastDigitIdx != iaStats.LastDigitIdx {
		t.Errorf("Expected ia LastDigitIdx= '%v' .  Instead, ia LastDigitIdx= '%v' .", eLastDigitIdx, iaStats.LastDigitIdx)
	}

	eIsZeroValue = !eIsZeroValue
	eIsZeroValue = !eIsZeroValue

	if eIsZeroValue != iaStats.IsZeroValue {
		t.Errorf("Expected ia IsZeroValue= '%v' .  Instead, ia IsZeroValue= '%v' .", eIsZeroValue, iaStats.IsZeroValue)
	}

	eIsIntegerZeroValue = !eIsIntegerZeroValue
	eIsIntegerZeroValue = !eIsIntegerZeroValue

	if eIsIntegerZeroValue != iaStats.IsIntegerZeroValue {
		t.Errorf("Expected ia IsIntegerZeroValue= '%v' .  Instead, ia IsIntegerZeroValue= '%v' .", eIsIntegerZeroValue, iaStats.IsIntegerZeroValue)
	}

	if eAryLen != iaStats.IntAryLen {
		t.Errorf("Expected ia IntAryLen= '%v' .  Instead, ia IntAryLen= '%v' .", eAryLen, iaStats.IntAryLen)
	}

	if eIntegerLen != iaStats.IntegerLen {
		t.Errorf("Expected ia IntegerLen= '%v' .  Instead, ia IntegerLen= '%v' .", eIntegerLen, iaStats.IntegerLen)
	}

	if eSigFractionLen != iaStats.SignificantFractionLen {
		t.Errorf("Expected ia SignificantFractionLen= '%v' .  Instead, ia SignificantFractionLen= '%v' .", eSigFractionLen, iaStats.SignificantFractionLen)
	}

	if eSigIntegerLen != iaStats.SignificantIntegerLen {
		t.Errorf("Expected ia SignificantIntegerLen= '%v' .  Instead, ia SignificantIntegerLen= '%v' .", eSigIntegerLen, iaStats.SignificantIntegerLen)
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

func TestIntAry_SetSignificantDigitIdxs_02(t *testing.T) {
	nStr := "000123456.123456000"
	eAryLen := 18
	eNumStr := "000123456.123456000"
	eIntegerLen := 9
	eSigIntegerLen := 6
	eSigFractionLen := 6
	eIsZeroValue := false
	eIsIntegerZeroValue := false
	eFirstDigitIdx := 3
	eLastDigitIdx := 14
	eSignVal := 1
	ePrecision := 9

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned from SetIntAryWithNumStr(nStr). numStr= %v ", eNumStr)
	}

	iaStats := ia.GetIntAryStats()

	if eFirstDigitIdx != iaStats.FirstDigitIdx {
		t.Errorf("Expected ia FirstDigitIdx= '%v' .  Instead, ia FirstDigitIdx= '%v' .", eFirstDigitIdx, iaStats.FirstDigitIdx)
	}

	if eLastDigitIdx != iaStats.LastDigitIdx {
		t.Errorf("Expected ia LastDigitIdx= '%v' .  Instead, ia LastDigitIdx= '%v' .", eLastDigitIdx, iaStats.LastDigitIdx)
	}

	eIsZeroValue = !eIsZeroValue
	eIsZeroValue = !eIsZeroValue

	if eIsZeroValue != iaStats.IsZeroValue {
		t.Errorf("Expected ia IsZeroValue= '%v' .  Instead, ia IsZeroValue= '%v' .", eIsZeroValue, iaStats.IsZeroValue)
	}

	eIsIntegerZeroValue = !eIsIntegerZeroValue
	eIsIntegerZeroValue = !eIsIntegerZeroValue

	if eIsIntegerZeroValue != iaStats.IsIntegerZeroValue {
		t.Errorf("Expected ia IsIntegerZeroValue= '%v' .  Instead, ia IsIntegerZeroValue= '%v' .", eIsIntegerZeroValue, iaStats.IsIntegerZeroValue)
	}

	if eAryLen != iaStats.IntAryLen {
		t.Errorf("Expected ia IntAryLen= '%v' .  Instead, ia IntAryLen= '%v' .", eAryLen, iaStats.IntAryLen)
	}

	if eIntegerLen != iaStats.IntegerLen {
		t.Errorf("Expected ia IntegerLen= '%v' .  Instead, ia IntegerLen= '%v' .", eIntegerLen, iaStats.IntegerLen)
	}

	if eSigFractionLen != iaStats.SignificantFractionLen {
		t.Errorf("Expected ia SignificantFractionLen= '%v' .  Instead, ia SignificantFractionLen= '%v' .", eSigFractionLen, iaStats.SignificantFractionLen)
	}

	if eSigIntegerLen != iaStats.SignificantIntegerLen {
		t.Errorf("Expected ia SignificantIntegerLen= '%v' .  Instead, ia SignificantIntegerLen= '%v' .", eSigIntegerLen, iaStats.SignificantIntegerLen)
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

func TestIntAry_SetSignificantDigitIdxs_03(t *testing.T) {
	nStr := "-000123456.123456000"
	eAryLen := 18
	eNumStr := "-000123456.123456000"
	eIntegerLen := 9
	eSigIntegerLen := 6
	eSigFractionLen := 6
	eIsZeroValue := false
	eIsIntegerZeroValue := false
	eFirstDigitIdx := 3
	eLastDigitIdx := 14
	eSignVal := -1
	ePrecision := 9

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned from SetIntAryWithNumStr(nStr). numStr= %v ", eNumStr)
	}

	iaStats := ia.GetIntAryStats()

	if eFirstDigitIdx != iaStats.FirstDigitIdx {
		t.Errorf("Expected ia FirstDigitIdx= '%v' .  Instead, ia FirstDigitIdx= '%v' .", eFirstDigitIdx, iaStats.FirstDigitIdx)
	}

	if eLastDigitIdx != iaStats.LastDigitIdx {
		t.Errorf("Expected ia LastDigitIdx= '%v' .  Instead, ia LastDigitIdx= '%v' .", eLastDigitIdx, iaStats.LastDigitIdx)
	}

	eIsZeroValue = !eIsZeroValue
	eIsZeroValue = !eIsZeroValue

	if eIsZeroValue != iaStats.IsZeroValue {
		t.Errorf("Expected ia IsZeroValue= '%v' .  Instead, ia IsZeroValue= '%v' .", eIsZeroValue, iaStats.IsZeroValue)
	}

	eIsIntegerZeroValue = !eIsIntegerZeroValue
	eIsIntegerZeroValue = !eIsIntegerZeroValue

	if eIsIntegerZeroValue != iaStats.IsIntegerZeroValue {
		t.Errorf("Expected ia IsIntegerZeroValue= '%v' .  Instead, ia IsIntegerZeroValue= '%v' .", eIsIntegerZeroValue, iaStats.IsIntegerZeroValue)
	}

	if eAryLen != iaStats.IntAryLen {
		t.Errorf("Expected ia IntAryLen= '%v' .  Instead, ia IntAryLen= '%v' .", eAryLen, iaStats.IntAryLen)
	}

	if eIntegerLen != iaStats.IntegerLen {
		t.Errorf("Expected ia IntegerLen= '%v' .  Instead, ia IntegerLen= '%v' .", eIntegerLen, iaStats.IntegerLen)
	}

	if eSigFractionLen != iaStats.SignificantFractionLen {
		t.Errorf("Expected ia SignificantFractionLen= '%v' .  Instead, ia SignificantFractionLen= '%v' .", eSigFractionLen, iaStats.SignificantFractionLen)
	}

	if eSigIntegerLen != iaStats.SignificantIntegerLen {
		t.Errorf("Expected ia SignificantIntegerLen= '%v' .  Instead, ia SignificantIntegerLen= '%v' .", eSigIntegerLen, iaStats.SignificantIntegerLen)
	}

	if eNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
	}

	if int(ePrecision) != iaStats.Precision {
		t.Errorf("Expected ia Precision= '%v' . Instead, ia Precision= '%v'", ePrecision, iaStats.Precision)
	}

	if eSignVal != iaStats.SignVal {
		t.Errorf("Expected ia SignVal= '%v' . Instead, ia GetSignVal= '%v'", eSignVal, iaStats.SignVal)
	}

}

func TestIntAry_SetSignificantDigitIdxs_04(t *testing.T) {
	nStr := "000.123456000"
	eAryLen := 12
	eNumStr := "000.123456000"
	eIntegerLen := 3
	eSigIntegerLen := 1
	eSigFractionLen := 6
	eIsZeroValue := false
	eIsIntegerZeroValue := true
	eFirstDigitIdx := 2
	eLastDigitIdx := 8
	eSignVal := 1
	ePrecision := 9

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned from SetIntAryWithNumStr(nStr). numStr= %v ", eNumStr)
	}

	iaStats := ia.GetIntAryStats()

	if eFirstDigitIdx != iaStats.FirstDigitIdx {
		t.Errorf("Expected ia FirstDigitIdx= '%v' .  Instead, ia FirstDigitIdx= '%v' .", eFirstDigitIdx, iaStats.FirstDigitIdx)
	}

	if eLastDigitIdx != iaStats.LastDigitIdx {
		t.Errorf("Expected ia LastDigitIdx= '%v' .  Instead, ia LastDigitIdx= '%v' .", eLastDigitIdx, iaStats.LastDigitIdx)
	}

	eIsZeroValue = !eIsZeroValue
	eIsZeroValue = !eIsZeroValue

	if eIsZeroValue != iaStats.IsZeroValue {
		t.Errorf("Expected ia IsZeroValue= '%v' .  Instead, ia IsZeroValue= '%v' .", eIsZeroValue, iaStats.IsZeroValue)
	}

	eIsIntegerZeroValue = !eIsIntegerZeroValue
	eIsIntegerZeroValue = !eIsIntegerZeroValue

	if eIsIntegerZeroValue != iaStats.IsIntegerZeroValue {
		t.Errorf("Expected ia IsIntegerZeroValue= '%v' .  Instead, ia IsIntegerZeroValue= '%v' .", eIsIntegerZeroValue, iaStats.IsIntegerZeroValue)
	}

	if eAryLen != iaStats.IntAryLen {
		t.Errorf("Expected ia IntAryLen= '%v' .  Instead, ia IntAryLen= '%v' .", eAryLen, iaStats.IntAryLen)
	}

	if eIntegerLen != iaStats.IntegerLen {
		t.Errorf("Expected ia IntegerLen= '%v' .  Instead, ia IntegerLen= '%v' .", eIntegerLen, iaStats.IntegerLen)
	}

	if eSigFractionLen != iaStats.SignificantFractionLen {
		t.Errorf("Expected ia SignificantFractionLen= '%v' .  Instead, ia SignificantFractionLen= '%v' .", eSigFractionLen, iaStats.SignificantFractionLen)
	}

	if eSigIntegerLen != iaStats.SignificantIntegerLen {
		t.Errorf("Expected ia SignificantIntegerLen= '%v' .  Instead, ia SignificantIntegerLen= '%v' .", eSigIntegerLen, iaStats.SignificantIntegerLen)
	}

	if eNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
	}

	if int(ePrecision) != iaStats.Precision {
		t.Errorf("Expected iaStats.Precision= '%v' . Instead, iaStats.Precision= '%v'", ePrecision, iaStats.Precision)
	}

	if eSignVal != iaStats.SignVal {
		t.Errorf("Expected iaStats.SignVal= '%v' . Instead, iaStats.SignVal= '%v'", eSignVal, iaStats.SignVal)
	}

}

func TestIntAry_SetSignificantDigitIdxs_05(t *testing.T) {
	nStr := "256"
	eAryLen := 3
	eNumStr := "256"
	eIntegerLen := 3
	eSigIntegerLen := 3
	eSigFractionLen := 0
	eIsZeroValue := false
	eIsIntegerZeroValue := false
	eFirstDigitIdx := 0
	eLastDigitIdx := 2
	eSignVal := 1
	ePrecision := 0

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned from SetIntAryWithNumStr(nStr). numStr= %v ", eNumStr)
	}

	iaStats := ia.GetIntAryStats()

	if eFirstDigitIdx != iaStats.FirstDigitIdx {
		t.Errorf("Expected ia FirstDigitIdx= '%v' .  Instead, ia FirstDigitIdx= '%v' .", eFirstDigitIdx, iaStats.FirstDigitIdx)
	}

	if eLastDigitIdx != iaStats.LastDigitIdx {
		t.Errorf("Expected ia LastDigitIdx= '%v' .  Instead, ia LastDigitIdx= '%v' .", eLastDigitIdx, iaStats.LastDigitIdx)
	}

	eIsZeroValue = !eIsZeroValue
	eIsZeroValue = !eIsZeroValue

	if eIsZeroValue != iaStats.IsZeroValue {
		t.Errorf("Expected iaStats.IsZeroValue= '%v' .  Instead, iaStats.IsZeroValue= '%v' .", eIsZeroValue, iaStats.IsZeroValue)
	}

	eIsIntegerZeroValue = !eIsIntegerZeroValue
	eIsIntegerZeroValue = !eIsIntegerZeroValue

	if eIsIntegerZeroValue != iaStats.IsIntegerZeroValue {
		t.Errorf("Expected iaStats.IsIntegerZeroValue= '%v' .  Instead, iaStats.IsIntegerZeroValue= '%v' .", eIsIntegerZeroValue, iaStats.IsIntegerZeroValue)
	}

	if eAryLen != iaStats.IntAryLen {
		t.Errorf("Expected iaStats.IntAryLen= '%v' .  Instead, iaStats.IntAryLen= '%v' .", eAryLen, iaStats.IntAryLen)
	}

	if eIntegerLen != iaStats.IntegerLen {
		t.Errorf("Expected iaStats.IntegerLen= '%v' .  Instead, iaStats.IntegerLen= '%v' .", eIntegerLen, iaStats.IntegerLen)
	}

	if eSigFractionLen != iaStats.SignificantFractionLen {
		t.Errorf("Expected iaStats.SignificantFractionLen= '%v' .  Instead, iaStats.SignificantFractionLen= '%v' .", eSigFractionLen, iaStats.SignificantFractionLen)
	}

	if eSigIntegerLen != iaStats.SignificantIntegerLen {
		t.Errorf("Expected iaStats.SignificantIntegerLen= '%v' .  Instead, iaStats.SignificantIntegerLen= '%v' .", eSigIntegerLen, iaStats.SignificantIntegerLen)
	}

	if eNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
	}

	if int(ePrecision) != iaStats.Precision {
		t.Errorf("Expected iaStats.Precision= '%v' . Instead, iaStats.Precision= '%v'", ePrecision, iaStats.Precision)
	}

	if eSignVal != iaStats.SignVal {
		t.Errorf("Expected iaStats.SignVal= '%v' . Instead, iaStats.SignVal= '%v'", eSignVal, iaStats.SignVal)
	}

}

func TestIntAry_SetSignificantDigitIdxs_06(t *testing.T) {
	nStr := "000256"
	eAryLen := 6
	eNumStr := "000256"
	eIntegerLen := 6
	eSigIntegerLen := 3
	eSigFractionLen := 0
	eIsZeroValue := false
	eIsIntegerZeroValue := false
	eFirstDigitIdx := 3
	eLastDigitIdx := 5
	eSignVal := 1
	ePrecision := 0

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned from SetIntAryWithNumStr(nStr). numStr= %v ", eNumStr)
	}

	iaStats := ia.GetIntAryStats()

	if eFirstDigitIdx != iaStats.FirstDigitIdx {
		t.Errorf("Expected iaStats.FirstDigitIdx= '%v' .  Instead, iaStats.FirstDigitIdx= '%v' .", eFirstDigitIdx, iaStats.FirstDigitIdx)
	}

	if eLastDigitIdx != iaStats.LastDigitIdx {
		t.Errorf("Expected iaStats.LastDigitIdx= '%v' .  Instead, iaStats.LastDigitIdx= '%v' .", eLastDigitIdx, iaStats.LastDigitIdx)
	}

	eIsZeroValue = !eIsZeroValue
	eIsZeroValue = !eIsZeroValue

	if eIsZeroValue != iaStats.IsZeroValue {
		t.Errorf("Expected iaStats.IsZeroValue= '%v' .  Instead, iaStats.IsZeroValue= '%v' .", eIsZeroValue, iaStats.IsZeroValue)
	}

	eIsIntegerZeroValue = !eIsIntegerZeroValue
	eIsIntegerZeroValue = !eIsIntegerZeroValue

	if eIsIntegerZeroValue != iaStats.IsIntegerZeroValue {
		t.Errorf("Expected iaStats.IsIntegerZeroValue= '%v' .  Instead, iaStats.IsIntegerZeroValue= '%v' .", eIsIntegerZeroValue, iaStats.IsIntegerZeroValue)
	}

	if eAryLen != iaStats.IntAryLen {
		t.Errorf("Expected iaStats.IntAryLen= '%v' .  Instead, iaStats.IntAryLen= '%v' .", eAryLen, iaStats.IntAryLen)
	}

	if eIntegerLen != iaStats.IntegerLen {
		t.Errorf("Expected iaStats.IntegerLen= '%v' .  Instead, iaStats.IntegerLen= '%v' .", eIntegerLen, iaStats.IntegerLen)
	}

	if eSigFractionLen != iaStats.SignificantFractionLen {
		t.Errorf("Expected iaStats.SignificantFractionLen= '%v' .  Instead, iaStats.SignificantFractionLen= '%v' .", eSigFractionLen, iaStats.SignificantFractionLen)
	}

	if eSigIntegerLen != iaStats.SignificantIntegerLen {
		t.Errorf("Expected iaStats.SignificantIntegerLen= '%v' .  Instead, iaStats.SignificantIntegerLen= '%v' .", eSigIntegerLen, iaStats.SignificantIntegerLen)
	}

	if eNumStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v'", eNumStr, ia.GetNumStr())
	}

	if int(ePrecision) != iaStats.Precision {
		t.Errorf("Expected iaStats.Precision= '%v' . Instead, iaStats.Precision= '%v'", ePrecision, iaStats.Precision)
	}

	if eSignVal != iaStats.SignVal {
		t.Errorf("Expected iaStats.SignVal= '%v' . Instead, iaStats.SignVal= '%v'", eSignVal, iaStats.SignVal)
	}

}

func TestIntAry_SetThousandsSeparator_01(t *testing.T) {
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

func TestIntAry_SubtractFromThis_01(t *testing.T) {
	nStr1 := "900.777"
	nStr2 := "901.000"
	eNumStr := "-0.223"
	ePrecision := 3
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_02(t *testing.T) {
	nStr1 := "350"
	nStr2 := "122"
	eNumStr := "228"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_03(t *testing.T) {
	nStr1 := "-350"
	nStr2 := "122"
	eNumStr := "-472"
	ePrecision := 0
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_04(t *testing.T) {
	nStr1 := "-350"
	nStr2 := "-122"
	eNumStr := "-228"
	ePrecision := 0
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_05(t *testing.T) {
	nStr1 := "350"
	nStr2 := "-122"
	eNumStr := "472"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_06(t *testing.T) {
	nStr1 := "350"
	nStr2 := "0"
	eNumStr := "350"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_07(t *testing.T) {
	nStr1 := "-350"
	nStr2 := "0"
	eNumStr := "-350"
	ePrecision := 0
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_08(t *testing.T) {
	nStr1 := "122"
	nStr2 := "350"
	eNumStr := "-228"
	ePrecision := 0
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_09(t *testing.T) {
	nStr1 := "-122"
	nStr2 := "350"
	eNumStr := "-472"
	ePrecision := 0
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_10(t *testing.T) {
	nStr1 := "-122"
	nStr2 := "-350"
	eNumStr := "228"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_11(t *testing.T) {
	nStr1 := "122"
	nStr2 := "-350"
	eNumStr := "472"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_12(t *testing.T) {
	nStr1 := "0"
	nStr2 := "350"
	eNumStr := "-350"
	ePrecision := 0
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_13(t *testing.T) {
	nStr1 := "0"
	nStr2 := "-350"
	eNumStr := "350"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_14(t *testing.T) {
	nStr1 := "122"
	nStr2 := "122"
	eNumStr := "0"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_15(t *testing.T) {
	nStr1 := "-122"
	nStr2 := "122"
	eNumStr := "-244"
	ePrecision := 0
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_16(t *testing.T) {
	nStr1 := "-122"
	nStr2 := "-122"
	eNumStr := "0"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_17(t *testing.T) {
	nStr1 := "122"
	nStr2 := "-122"
	eNumStr := "244"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_18(t *testing.T) {
	nStr1 := "0"
	nStr2 := "0"
	eNumStr := "0"
	ePrecision := 0
	eSignVal := 1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}

func TestIntAry_SubtractFromThis_19(t *testing.T) {
	nStr1 := "1.122"
	nStr2 := "4.5"
	eNumStr := "-3.378"
	ePrecision := 3
	eSignVal := -1

	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	err := ia1.SubtractFromThis(&ia2)

	if err != nil {
		t.Errorf("Error returned from ia1.SubtractFromThis(&ia2, true). Error= %v", err)
	}

	if eNumStr != ia1.GetNumStr() {
		t.Errorf("Error - Expected ia1.GetNumStr()= '%v' .  Instead, ia1.GetNumStr()= '%v' .", eNumStr, ia1.GetNumStr())
	}

	if ePrecision != ia1.GetPrecision() {
		t.Errorf("Error - Expected ia1.GetPrecision()= '%v' .  Instead, ia1.GetPrecision()= '%v' .", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error - Expected ia1.GetSign()= '%v' .  Instead, ia1.GetSign()= '%v' .", eSignVal, ia1.GetSign())
	}

}


func TestIntAry_SubtractMultipleFromThis_01(t *testing.T) {
	nStrBase := "197.452"
	nStr1 := "1.122"
	nStr2 := "4.5"
	nStr3 := "32.148"
	nStr4 := "10.0"
	eNumStr := "149.682"
	ePrecision := 3
	eSignVal := 1

	iaBase := IntAry{}.New()
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia3 := IntAry{}.New()
	ia4 := IntAry{}.New()
	iaBase.SetIntAryWithNumStr(nStrBase)
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	ia3.SetIntAryWithNumStr(nStr3)
	ia4.SetIntAryWithNumStr(nStr4)
	err := iaBase.SubtractMultipleFromThis( &ia1, &ia2, &ia3, &ia4)

	if err != nil {
		t.Errorf("Error returned from iaBase.SubtractMultipleFromThis(true, ia...). Error= %v", err)
	}

	if eNumStr != iaBase.GetNumStr() {
		t.Errorf("Error - Expected iaBase.GetNumStr()= '%v' .  Instead, iaBase.GetNumStr()= '%v' .", eNumStr, iaBase.GetNumStr())
	}

	if ePrecision != iaBase.GetPrecision() {
		t.Errorf("Error - Expected iaBase.GetPrecision()= '%v' .  Instead, iaBase.GetPrecision()= '%v' .", ePrecision, iaBase.GetPrecision())
	}

	if eSignVal != iaBase.GetSign() {
		t.Errorf("Error - Expected iaBase.GetSign()= '%v' .  Instead, iaBase.GetSign()= '%v' .", eSignVal, iaBase.GetSign())
	}

}

func TestIntAry_SubtractMultipleFromThis_02(t *testing.T) {
	nStrBase := "197.452"
	nStr1 := "1.122"
	nStr2 := "4.5"
	nStr3 := "-32.148"
	nStr4 := "10.0"
	eNumStr := "213.978"
	ePrecision := 3
	eSignVal := 1

	iaBase := IntAry{}.New()
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia3 := IntAry{}.New()
	ia4 := IntAry{}.New()
	iaBase.SetIntAryWithNumStr(nStrBase)
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	ia3.SetIntAryWithNumStr(nStr3)
	ia4.SetIntAryWithNumStr(nStr4)
	err := iaBase.SubtractMultipleFromThis( &ia1, &ia2, &ia3, &ia4)

	if err != nil {
		t.Errorf("Error returned from iaBase.SubtractMultipleFromThis(true, ia...). Error= %v", err)
	}

	if eNumStr != iaBase.GetNumStr() {
		t.Errorf("Error - Expected iaBase.GetNumStr()= '%v' .  Instead, iaBase.GetNumStr()= '%v' .", eNumStr, iaBase.GetNumStr())
	}

	if ePrecision != iaBase.GetPrecision() {
		t.Errorf("Error - Expected iaBase.GetPrecision()= '%v' .  Instead, iaBase.GetPrecision()= '%v' .", ePrecision, iaBase.GetPrecision())
	}

	if eSignVal != iaBase.GetSign() {
		t.Errorf("Error - Expected iaBase.GetSign()= '%v' .  Instead, iaBase.GetSign()= '%v' .", eSignVal, iaBase.GetSign())
	}

}

func TestIntAry_SubtractMultipleFromThis_03(t *testing.T) {
	nStrBase := "197.452"
	nStr1 := "1.122"
	nStr2 := "4.5"
	nStr3 := "932.148"
	nStr4 := "10.0"
	eNumStr := "-750.318"
	ePrecision := 3
	eSignVal := -1

	iaBase := IntAry{}.New()
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	ia3 := IntAry{}.New()
	ia4 := IntAry{}.New()
	iaBase.SetIntAryWithNumStr(nStrBase)
	ia1.SetIntAryWithNumStr(nStr1)
	ia2.SetIntAryWithNumStr(nStr2)
	ia3.SetIntAryWithNumStr(nStr3)
	ia4.SetIntAryWithNumStr(nStr4)
	err := iaBase.SubtractMultipleFromThis( &ia1, &ia2, &ia3, &ia4)

	if err != nil {
		t.Errorf("Error returned from iaBase.SubtractMultipleFromThis(true, ia...). Error= %v", err)
	}

	if eNumStr != iaBase.GetNumStr() {
		t.Errorf("Error - Expected iaBase.GetNumStr()= '%v' .  Instead, iaBase.GetNumStr()= '%v' .", eNumStr, iaBase.GetNumStr())
	}

	if ePrecision != iaBase.GetPrecision() {
		t.Errorf("Error - Expected iaBase.GetPrecision()= '%v' .  Instead, iaBase.GetPrecision()= '%v' .", ePrecision, iaBase.GetPrecision())
	}

	if eSignVal != iaBase.GetSign() {
		t.Errorf("Error - Expected iaBase.GetSign()= '%v' .  Instead, iaBase.GetSign()= '%v' .", eSignVal, iaBase.GetSign())
	}

}
