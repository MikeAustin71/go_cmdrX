package common

import (
	"testing"
)


func TestIntAry_CompareSignedValues_01(t *testing.T) {

	nStr1 := "451.3"
	nStr2 := "451.2"
	expectedCompare := 1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_02(t *testing.T) {

	nStr1 := "45.13"
	nStr2 := "451.2"
	expectedCompare := -1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_03(t *testing.T) {

	nStr1 := "45.13975"
	nStr2 := "-451.21"
	expectedCompare := 1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_04(t *testing.T) {

	nStr1 := "0"
	nStr2 := "0.00"
	expectedCompare := 0

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_05(t *testing.T) {

	nStr1 := "-625.414"
	nStr2 := "-625.413"
	expectedCompare := -1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_06(t *testing.T) {

	nStr1 := "625.414"
	nStr2 := "625.413000"
	expectedCompare := 1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_07(t *testing.T) {

	nStr1 := "625.413"
	nStr2 := "625.413001"
	expectedCompare := -1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_08(t *testing.T) {

	nStr1 := "625.413"
	nStr2 := "625.413000"
	expectedCompare := 0

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_09(t *testing.T) {

	nStr1 := "625.413"
	nStr2 := "00625.413000"
	expectedCompare := 0

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_10(t *testing.T) {

	nStr1 := "625.413"
	nStr2 := "-00625.413000"
	expectedCompare := 1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_11(t *testing.T) {

	nStr1 := "-625.413"
	nStr2 := "-00625.413000"
	expectedCompare := 0

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareSignedValues_12(t *testing.T) {

	nStr1 := "-625.413"
	nStr2 := "00625.413000"
	expectedCompare := -1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareSignedValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareAbsoluteValues_01(t *testing.T) {

	nStr1 := "45.13"
	nStr2 := "451.2"
	expectedCompare := -1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareAbsoluteValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareAbsoluteValues_02(t *testing.T) {

	nStr1 := "45.13"
	nStr2 := "-451.2"
	expectedCompare := -1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareAbsoluteValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareAbsoluteValues_03(t *testing.T) {

	nStr1 := "-45.13"
	nStr2 := "45.13"
	expectedCompare := 0

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareAbsoluteValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareAbsoluteValues_04(t *testing.T) {

	nStr1 := "-45.13000"
	nStr2 := "45.13"
	expectedCompare := 0

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareAbsoluteValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareAbsoluteValues_05(t *testing.T) {

	nStr1 := "-45.13001"
	nStr2 := "45.13"
	expectedCompare := 1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareAbsoluteValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CompareAbsoluteValues_06(t *testing.T) {

	nStr1 := "-45.1300"
	nStr2 := "452"
	expectedCompare := -1

	ia1 := IntAry{}.New()

	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()

	ia2.SetIntAryWithNumStr(nStr2)

	actualCompare := ia1.CompareAbsoluteValues(&ia2)

	if actualCompare != expectedCompare {
		t.Errorf("Error. Expected Comparison Result= '%v'. Instead, Comparison Result= '%v'", expectedCompare, actualCompare)
	}

}

func TestIntAry_CopyIn_01(t *testing.T) {

	nStr1 := "00100.1230"
	nStr2 := "-0000052.795813000"

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	if nStr1 != ia.GetNumStr() {
		t.Errorf("Error ia.SetIntAryWithNumStr(nStr1). Expected ia.GetNumStr()= '%v'  .   Instead, ia.GetNumStr()= '%v' .", nStr1, ia.GetNumStr())
	}

	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)

	if nStr2 != ia2.GetNumStr() {
		t.Errorf("Error ia2.SetIntAryWithNumStr(nStr2). Expected ia.GetNumStr()= '%v'  .   Instead, ia.GetNumStr()= '%v' .", nStr2, ia2.GetNumStr())
	}

	ia.CopyIn(&ia2, false)

	if ia2.GetNumStr() != ia.GetNumStr() {
		t.Errorf("Error Expected ia.GetNumStr()= '%v'  .   Instead, ia.GetNumStr()= '%v' .", ia2.GetNumStr(), ia.GetNumStr())
	}

	if ia2.GetSign() != ia.GetSign() {
		t.Errorf("Error Expcted ia.GetSign()= '%v'  .   Instead, ia.GetSign()= '%v' .", ia2.GetSign(), ia.GetSign())
	}

	if ia2.GetPrecision() != ia.GetPrecision() {
		t.Errorf("Error Expcted ia.GetPrecision()= '%v'  .   Instead, ia.GetPrecision()= '%v' .", ia2.GetPrecision(), ia.GetPrecision())
	}

	if ia2.GetIntAryLength() != ia.GetIntAryLength() {
		t.Errorf("Error Expcted ia integer array length= '%v'  .   Instead, ia integer array length= '%v' .", ia2.GetIntAryLength(), ia.GetIntAryLength())
	}

	ia2Stats := ia2.GetIntAryStats()
	iaStats := ia.GetIntAryStats()

	if ia2Stats.IntegerLen != iaStats.IntegerLen {
		t.Errorf("Error Expected ia2 Integer Length == ia Integer Length. Instead ia2 IntegerLen= '%v'  and ia IntegerLen= '%v' .", ia2Stats.IntegerLen, iaStats.IntegerLen)
	}

	if ia2Stats.SignificantIntegerLen != iaStats.SignificantIntegerLen {
		t.Errorf("Error Expected ia2 Signigicant Integer Length == ia Significant Integer Length. Instead ia2 SignificantIntegerLen= '%v'  and ia SignificantIntegerLen= '%v' .", ia2Stats.SignificantIntegerLen, iaStats.SignificantIntegerLen)
	}

	if ia2Stats.SignificantFractionLen != iaStats.SignificantFractionLen {
		t.Errorf("Error Expected ia2 SignificantFractionLen == ia SignificantFractionLen. Intead, ia2 SignificantFractionLen= '%v' and ia SignificantFractionLen= '%v' .", ia2Stats.SignificantFractionLen, iaStats.SignificantFractionLen)
	}

	if ia2Stats.FirstDigitIdx != iaStats.FirstDigitIdx {
		t.Errorf("Error Expected ia2 FirstDigitIdx == ia FirstDigitIdx. Intead,  ia2 FirstDigitIdx= '%v' and ia FirstDigitIdx= '%v' .", ia2Stats.FirstDigitIdx, iaStats.FirstDigitIdx)
	}

	if ia2Stats.LastDigitIdx != iaStats.LastDigitIdx {
		t.Errorf("Error Expected ia2 LastDigitIdx == ia LastDigitIdx. Instead, ia2 LastDigitIdx= '%v'  and ia LastDigitIdx= '%v' .", ia2Stats.LastDigitIdx, iaStats.LastDigitIdx)
	}

	if ia2Stats.IsZeroValue != iaStats.IsZeroValue {
		t.Errorf("Error Expected ia2 IsZeroValue == ia IsZeroValue. Instead, ia2 IsZeroValue= '%v' and ia IsZeroValue= '%v' .", ia2Stats.IsZeroValue, iaStats.IsZeroValue)
	}

	if ia2Stats.IsIntegerZeroValue != iaStats.IsIntegerZeroValue {
		t.Errorf("Error Expected ia2 IsIntegerZeroValue == ia IsIntegerZeroValue. Instead,  ia2 IsIntegerZeroValue= '%v' and ia IsIntegerZeroValue= '%v' .", ia2Stats.IsIntegerZeroValue, iaStats.IsIntegerZeroValue)
	}

	if ia2.GetDecimalSeparator() != '.' {
		t.Errorf("Error Expcted ia2 DecimalSeparator= '%v'  .   Instead, ia2 DecimalSeparator= '%v' .", '.', ia.GetDecimalSeparator())
	}

	for i := 0; i < ia.GetIntAryLength(); i++ {

		ia2Element, _ := ia2.GetIntAryElement(i)
		iaElement, _ := ia.GetIntAryElement(i)
		if ia2Element != iaElement {
			t.Errorf("Error Expcted ia2 integer array element [i]= '%v'  .   Instead, ia integer array element [i]= '%v' .", ia2Element, iaElement)
		}
	}
}

func TestIntAry_CopyIn_02(t *testing.T) {

	nStr1 := "00100.1230"
	nStr2 := "3594176.123459"

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.CopyToBackUp()

	if nStr1 != ia.GetNumStr() {
		t.Errorf("Error ia.SetIntAryWithNumStr(nStr1). Expected ia.GetNumStr()= '%v'  .   Instead, ia.GetNumStr()= '%v' .", nStr1, ia.GetNumStr())
	}

	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)
	ia2.CopyToBackUp()

	if nStr2 != ia2.GetNumStr() {
		t.Errorf("Error ia2.SetIntAryWithNumStr(nStr2). Expected ia.GetNumStr()= '%v'  .   Instead, ia.GetNumStr()= '%v' .", nStr2, ia2.GetNumStr())
	}

	ia.CopyIn(&ia2, true)

	if ia2.GetNumStr() != ia.GetNumStr() {
		t.Errorf("Error Expected ia.GetNumStr()= '%v'  .   Instead, ia.GetNumStr()= '%v' .", ia2.GetNumStr(), ia.GetNumStr())
	}

	if ia2.GetSign() != ia.GetSign() {
		t.Errorf("Error Expcted ia.GetSign()= '%v'  .   Instead, ia.GetSign()= '%v' .", ia2.GetSign(), ia.GetSign())
	}

	if ia2.GetPrecision() != ia.GetPrecision() {
		t.Errorf("Error Expcted ia.GetPrecision()= '%v'  .   Instead, ia.GetPrecision()= '%v' .", ia2.GetPrecision(), ia.GetPrecision())
	}

	iaStats := ia.GetIntAryStats()

	ia2Stats := ia2.GetIntAryStats()

	if ia2Stats.IntAryLen != iaStats.IntAryLen {
		t.Errorf("Error Expected ia2 IntAryLen == ia IntAryLen. Instead, ia2 integer array length = '%v' and ia integer array length = '%v' .", ia2Stats.IntAryLen, iaStats.IntAryLen)
	}

	if ia2Stats.IntegerLen != iaStats.IntegerLen {
		t.Errorf("Error Expected ai2 IntegerLen == ia IntegerLen. Instead, ia2 IntegerLen= '%v' and ia IntegerLen= '%v' .", ia2Stats.IntegerLen, iaStats.IntegerLen)
	}

	if ia2Stats.SignificantIntegerLen != iaStats.SignificantIntegerLen {
		t.Errorf("Error Expected ia2 SignificantIntegerLen == ia SignificantIntegerLen. Instead, ia2.SignificantIntegerLen= '%v' and ia SignificantIntegerLen= '%v' .", ia2Stats.SignificantIntegerLen, iaStats.SignificantIntegerLen)
	}

	if ia2Stats.SignificantFractionLen != iaStats.SignificantFractionLen {
		t.Errorf("Error Expected ia2 SignificantFractionLen == ia SignificantFractionLen. Instead, ia2 SignificantFractionLen= '%v' and ia SignificantFractionLen= '%v' .", ia2Stats.SignificantFractionLen, iaStats.SignificantFractionLen)
	}

	if ia2Stats.FirstDigitIdx != iaStats.FirstDigitIdx {
		t.Errorf("Error Expected ia2 FirstDigitIdx == ia FirstDigitIdx. Instead, ia2 FirstDigitIdx= '%v' and ia FirstDigitIdx= '%v' .", ia2Stats.FirstDigitIdx, iaStats.FirstDigitIdx)
	}

	if ia2Stats.LastDigitIdx != iaStats.LastDigitIdx {
		t.Errorf("Error Expected ia2 LastDigitIdx == ia LastDigitIdx.  Instead, ia2 LastDigitIdx= '%v' and ia LastDigitIdx= '%v' .", ia2Stats.LastDigitIdx, iaStats.LastDigitIdx)
	}

	if ia2Stats.IsZeroValue != iaStats.IsZeroValue {
		t.Errorf("Error Expected ia2 IsZeroValue == ia IsZeroValue.  Instead, ia2 IsZeroValue= '%v' and ia IsZeroValue= '%v' .", ia2Stats.IsZeroValue, iaStats.IsZeroValue)
	}

	if ia2Stats.IsIntegerZeroValue != iaStats.IsIntegerZeroValue {
		t.Errorf("Error Expected ia2 IsIntegerZeroValue == ia IsIntegerZeroValue.  Instead, ia2 IsIntegerZeroValue= '%v' and ia IsIntegerZeroValue= '%v' .", ia2Stats.IsIntegerZeroValue, iaStats.IsIntegerZeroValue)
	}

	if ia2Stats.DecimalSeparator != '.' {
		t.Errorf("Error Expected ia2 DecimalSeparator= '.'  .   Instead, ia2 DecimalSeparator= '%v' .", ia2Stats.DecimalSeparator)
	}

	if iaStats.DecimalSeparator != '.' {
		t.Errorf("Error Expected ia DecimalSeparator= '.'  .   Instead, ia DecimalSeparator= '%v' .",  iaStats.DecimalSeparator)
	}

	for i := 0; i < ia.GetIntAryLength(); i++ {

		ia2Element, _ := ia2.GetIntAryElement(i)
		iaElement, _ := ia.GetIntAryElement(i)

		if ia2Element != iaElement {
			t.Errorf("Error Expected ia2 integer array element [i]= '%v'  .   Instead, ia integer array element [i]= '%v' .", ia2Element, iaElement)
		}
	}

	if !ia.BackUp.Equals(&ia2.BackUp) {
		t.Error("Error Expected ia.Backup to Equal ia2.Backup.   ")
	}
}

func TestIntAry_CopyToBackUp_01(t *testing.T) {
	nStr1 := "100.123"
	nStr2 := "-52.795813"
	expected := "100.123"

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.CopyToBackUp()
	ia.SetIntAryWithNumStr(nStr2)

	if nStr2 != ia.GetNumStr() {
		t.Errorf("Error - Expected ia.GetNumStr()= '%v' . Instead, ia.GetNumStr()= '%v' .", nStr2, ia.GetNumStr())
	}

	str := ia.BackUp.GetNumStr()

	if expected != str {
		t.Errorf("Error - Expected ia.BackUp.GetNumStr()= '%v' .  Instead, ia.BackUp.GetNumStr()= '%v' .", expected, str)
	}

}

func TestIntAry_DecrementIntegerOne_01(t *testing.T) {
	nStr1 := "100.123"
	expected := "-100.123"
	cycles := 200
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	for i := 0; i < cycles; i++ {
		ia.DecrementIntegerOne()
	}

	if expected != ia.GetNumStr() {
		t.Errorf("Error - Expected numStr= '%v'. Instead, numStr= '%v'", expected, ia.GetNumStr())
	}
}

func TestIntAry_DecrementIntegerOne_02(t *testing.T) {
	nStr1 := "2000"
	expected := "-2000"
	cycles := 4000
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	for i := 0; i < cycles; i++ {
		ia.DecrementIntegerOne()
	}

	if expected != ia.GetNumStr() {
		t.Errorf("Error - Expected numStr= '%v'. Instead, numStr= '%v'", expected, ia.GetNumStr())
	}

}

func TestIntAry_DecrementIntegerOne_03(t *testing.T) {
	nStr1 := "2000.123"
	expected := "-2000.123"
	cycles := 4000
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)

	for i := 0; i < cycles; i++ {
		ia.DecrementIntegerOne()
	}

	if expected != ia.GetNumStr() {
		t.Errorf("Error - Expected numStr= '%v'. Instead, numStr= '%v'", expected, ia.GetNumStr())
	}

}

func TestIntAry_DivideByInt64_01(t *testing.T) {
	nStr1 := "579"
	expected := "17.545454545454545454545454545455"
	maxPrecision := uint64(30)
	divisor := int64(33)

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.DivideByInt64(divisor, maxPrecision)

	if err != nil {
		t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
	}

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.Numstr= '%v'", expected, ia.GetNumStr())
	}

	if ia.GetPrecision() != int(maxPrecision) {
		t.Errorf("Expected ia.GetPrecision()= '%v'. Instead, ia.GetPrecision()= '%v'.", maxPrecision, ia.GetPrecision())
	}

}

func TestIntAry_DivideByInt64_02(t *testing.T) {
	nStr1 := "579"
	maxPrecision := uint64(0)
	divisor := int64(0)

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.DivideByInt64(divisor, maxPrecision)

	if err == nil {
		t.Errorf("Expected Divide By Zero Error. No Error Presented. Error = %v", err)
	}

}

func TestIntAry_DivideByInt64_03(t *testing.T) {
	nStr1 := "4"
	expected := "2"
	maxPrecision := uint64(0)
	divisor := int64(2)

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.DivideByInt64(divisor, maxPrecision)

	if err != nil {
		t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
	}

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.Numstr= '%v'", expected, ia.GetNumStr())
	}

	if ia.GetPrecision() != int(maxPrecision) {
		t.Errorf("Expected ia.GetPrecision()= '%v'. Instead, ia.GetPrecision()= '%v'.", maxPrecision, ia.GetPrecision())
	}

}

func TestIntAry_DivideByInt64_04(t *testing.T) {
	nStr1 := "476"
	expected := "-14"
	maxPrecision := uint64(10)
	ePrecision := 0
	eSignVal := -1
	divisor := int64(-34)

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.DivideByInt64(divisor, maxPrecision)

	if err != nil {
		t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
	}

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.Numstr= '%v'", expected, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Expected ia.GetPrecision()= '%v'. Instead, ia.GetPrecision()= '%v'.", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v'. Instead, ia.GetPrecision()= '%v'.", eSignVal, ia.GetSign())
	}

}

func TestIntAry_DivideByInt64_05(t *testing.T) {
	nStr1 := "-476"
	expected := "-14"
	maxPrecision := uint64(10)
	ePrecision := 0
	eSignVal := -1
	divisor := int64(34)

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.DivideByInt64(divisor, maxPrecision)

	if err != nil {
		t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
	}

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.Numstr= '%v'", expected, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Expected ia.GetPrecision()= '%v'. Instead, ia.GetPrecision()= '%v'.", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v'. Instead, ia.GetPrecision()= '%v'.", eSignVal, ia.GetSign())
	}

}

func TestIntAry_DivideByInt64_06(t *testing.T) {
	nStr1 := "-476"
	expected := "14"
	maxPrecision := uint64(10)
	ePrecision := 0
	eSignVal := 1
	divisor := int64(-34)

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.DivideByInt64(divisor, maxPrecision)

	if err != nil {
		t.Errorf("Received Error from ia.DivideByInt64(divisor, maxPrecision, true ). Error= %v", err)
	}

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= '%v'. Instead, ia.Numstr= '%v'", expected, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Expected ia.GetPrecision()= '%v'. Instead, ia.GetPrecision()= '%v'.", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign()= '%v'. Instead, ia.GetPrecision()= '%v'.", eSignVal, ia.GetSign())
	}

}

func TestIntAry_DivideByTwo_01(t *testing.T) {

	nStr1 := "1959"
	expected := "979.5"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.DivideByTwo()

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected numStr= '%v'. Instead received numStr= '%v'", expected, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", precision, ia.GetPrecision())
	}

}

func TestIntAry_DivideByTwo_02(t *testing.T) {

	nStr1 := "1"
	expected := "0.5"
	precision := 1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.DivideByTwo()

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected numStr= '%v'. Instead received numStr= '%v'", expected, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", precision, ia.GetPrecision())
	}

}

func TestIntAry_DivideByTwo_03(t *testing.T) {

	nStr1 := "0"
	expected := "0"
	precision := 0

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.DivideByTwo()

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected numStr= '%v'. Instead received numStr= '%v'", expected, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", precision, ia.GetPrecision())
	}

}

func TestIntAry_DivideByTwo_04(t *testing.T) {

	nStr1 := "-2959"
	expected := "-1479.5"
	precision := 1
	signVal := -1

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.DivideByTwo()

	if expected != ia.GetNumStr() {
		t.Errorf("Error: Expected numStr= '%v'. Instead received numStr= '%v'", expected, ia.GetNumStr())
	}

	if precision != ia.GetPrecision() {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", precision, ia.GetPrecision())
	}

	if signVal != ia.GetSign() {
		t.Errorf("Error: Expected Sign Value= '%v'. Instead, Sign Value= '%v'. ", signVal, ia.GetSign())
	}

}

func TestIntAry_DivideByTenToPower_01(t *testing.T) {

	nStr := "457.3"
	power := uint(1)
	eNumStr := "45.73"
	eIAry := []uint8{4, 5, 7, 3}
	lEArray := len(eIAry)
	ePrecision := 2
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.DivideByTenToPower(power)

	if ia.GetNumStr() != eNumStr {
		t.Errorf("Error: Expected numStr= '%v'. Instead received numStr= '%v'", eNumStr, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
	}

	if lEArray != ia.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {

		if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_DivideByTenToPower_02(t *testing.T) {

	nStr := "457.3"
	power := uint(3)
	eNumStr := "0.4573"
	eIAry := []uint8{0, 4, 5, 7, 3}
	lEArray := len(eIAry)
	ePrecision := 4
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.DivideByTenToPower(power)

	if ia.GetNumStr() != eNumStr {
		t.Errorf("Error: Expected numStr= '%v'. Instead received numStr= '%v'", eNumStr, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
	}

	if lEArray != ia.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_DivideByTenToPower_03(t *testing.T) {

	nStr := "457.3"
	power := uint(7)
	eNumStr := "0.00004573"
	eIAry := []uint8{0, 0, 0, 0, 0, 4, 5, 7, 3}
	lEArray := len(eIAry)
	ePrecision := 8
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.DivideByTenToPower(power)

	if ia.GetNumStr() != eNumStr {
		t.Errorf("Error: Expected numStr= '%v'. Instead received numStr= '%v'", eNumStr, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
	}

	if lEArray != ia.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_DivideByTenToPower_04(t *testing.T) {

	nStr := "-457.3"
	power := uint(7)
	eNumStr := "-0.00004573"
	eIAry := []uint8{0, 0, 0, 0, 0, 4, 5, 7, 3}
	lEArray := len(eIAry)
	ePrecision := 8
	eSignVal := -1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.DivideByTenToPower(power)

	if ia.GetNumStr() != eNumStr {
		t.Errorf("Error: Expected numStr= '%v'. Instead received numStr= '%v'", eNumStr, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
	}

	if lEArray != ia.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_DivideByTenToPower_05(t *testing.T) {

	nStr := "0"
	power := uint(2)
	eNumStr := "0.00"
	eIAry := []uint8{0, 0, 0}
	lEArray := len(eIAry)
	ePrecision := 2
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.DivideByTenToPower(power)

	if ia.GetNumStr() != eNumStr {
		t.Errorf("Error: Expected numStr= '%v'. Instead received numStr= '%v'", eNumStr, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
	}

	if lEArray != ia.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_DivideByTenToPower_06(t *testing.T) {

	nStr := "-4573"
	power := uint(1)
	eNumStr := "-457.3"
	eIAry := []uint8{4, 5, 7, 3}
	lEArray := len(eIAry)
	ePrecision := 1
	eSignVal := -1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.DivideByTenToPower(power)

	if ia.GetNumStr() != eNumStr {
		t.Errorf("Error: Expected numStr= '%v'. Instead received numStr= '%v'", eNumStr, ia.GetNumStr())
	}

	if ia.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia.GetSign())
	}

	if lEArray != ia.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_DivideThisBy_01(t *testing.T) {
	dividend := "56234369384300"
	divisor := "24"
	eQuotient := "2343098724345.833333333333333333333"
	eSignVal := 1
	maxPrecision := 21
	ePrecision := 21

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.GetNumStr() {
		t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
	}

	if ePrecision != quotient.GetPrecision() {
		t.Errorf("Expected quotient.GetPrecision()= '%v' .  Instead, quotient.GetPrecision()= '%v'  .", ePrecision, quotient.GetPrecision())
	}

	if eSignVal != quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
	}

}

func TestIntAry_DivideThisBy_02(t *testing.T) {
	dividend := "48"
	divisor := "24"
	eSignVal := 1
	eQuotient := "2"
	maxPrecision := 21
	ePrecision := 0

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.GetNumStr() {
		t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
	}

	if ePrecision != quotient.GetPrecision() {
		t.Errorf("Expected quotient.GetPrecision()= '%v' .  Instead, quotient.GetPrecision()= '%v'  .", ePrecision, quotient.GetPrecision())
	}

	if eSignVal != quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
	}

}

func TestIntAry_DivideThisBy_03(t *testing.T) {
	dividend := "24"
	divisor := "24"
	eQuotient := "1"
	eSignVal := 1
	maxPrecision := 21
	ePrecision := 0

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.GetNumStr() {
		t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
	}

	if ePrecision != quotient.GetPrecision() {
		t.Errorf("Expected quotient.GetPrecision()= '%v' .  Instead, quotient.GetPrecision()= '%v'  .", ePrecision, quotient.GetPrecision())
	}

	if eSignVal != quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
	}

}

func TestIntAry_DivideThisBy_04(t *testing.T) {
	dividend := "0.05"
	divisor := "24"
	eQuotient := "0.00208333333333333333333333333333"
	eSignVal := 1
	maxPrecision := 32
	ePrecision := 32

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.GetNumStr() {
		t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
	}

	if ePrecision != quotient.GetPrecision() {
		t.Errorf("Expected quotient.GetPrecision()= '%v' .  Instead, quotient.GetPrecision()= '%v'  .", ePrecision, quotient.GetPrecision())
	}

	if eSignVal != quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
	}

}

func TestIntAry_DivideThisBy_05(t *testing.T) {
	dividend := "0"
	divisor := "24"
	eQuotient := "0"
	eSignVal := 1
	maxPrecision := 7
	ePrecision := 0

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.GetNumStr() {
		t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
	}

	if ePrecision != quotient.GetPrecision() {
		t.Errorf("Expected quotient.GetPrecision()= '%v' .  Instead, quotient.GetPrecision()= '%v'  .", ePrecision, quotient.GetPrecision())
	}

	if eSignVal != quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
	}

}

func TestIntAry_DivideThisBy_06(t *testing.T) {
	dividend := "48"
	divisor := "0"

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	_, err := ia1.DivideThisBy(&ia2, 15)

	if err == nil {
		t.Error("Expected an error from Divideby Zero. No Error Received!")
	}

}

func TestIntAry_DivideThisBy_07(t *testing.T) {
	dividend := "-9360"
	divisor := "24.48"
	eQuotient := "-382.35294117647058823529411764706"
	eSignVal := -1
	maxPrecision := 29
	ePrecision := 29

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.GetNumStr() {
		t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
	}

	if ePrecision != quotient.GetPrecision() {
		t.Errorf("Expected quotient.GetPrecision()= '%v' .  Instead, quotient.GetPrecision()= '%v'  .", ePrecision, quotient.GetPrecision())
	}

	if eSignVal != quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
	}
}

func TestIntAry_DivideThisBy_08(t *testing.T) {
	dividend := "-9360"
	divisor := "-24.48"
	eQuotient := "382.35294117647058823529411764706"
	eSignVal := 1
	maxPrecision := 29
	ePrecision := 29

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.GetNumStr() {
		t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
	}

	if ePrecision != quotient.GetPrecision() {
		t.Errorf("Expected quotient.GetPrecision()= '%v' .  Instead, quotient.GetPrecision()= '%v'  .", ePrecision, quotient.GetPrecision())
	}

	if eSignVal != quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
	}
}

func TestIntAry_DivideThisBy_09(t *testing.T) {
	dividend := "9360"
	divisor := "-24.48"
	eQuotient := "-382.35294117647058823529411764706"
	eSignVal := -1
	maxPrecision := 29
	ePrecision := 29

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.GetNumStr() {
		t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
	}

	if ePrecision != quotient.GetPrecision() {
		t.Errorf("Expected quotient.GetPrecision()= '%v' .  Instead, quotient.GetPrecision()= '%v'  .", ePrecision, quotient.GetPrecision())
	}

	if eSignVal != quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
	}
}

func TestIntAry_DivideThisBy_10(t *testing.T) {
	dividend := "-19260.549"
	divisor := "-246.483"
	eQuotient := "78.141490488187826340964691276883"
	eSignVal := 1
	maxPrecision := 30
	ePrecision := 30

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.GetNumStr() {
		t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
	}

	if ePrecision != quotient.GetPrecision() {
		t.Errorf("Expected quotient.GetPrecision()= '%v' .  Instead, quotient.GetPrecision()= '%v'  .", ePrecision, quotient.GetPrecision())
	}

	if eSignVal != quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
	}
}

func TestIntAry_DivideThisBy_11(t *testing.T) {
	// Testing condition where maxPrecision is
	// less than significant result.
	dividend := "1"
	divisor := "25"
	// actual result is "0.04"
	eQuotient := "0.0"
	eSignVal := 1
	maxPrecision := 1
	ePrecision := 1

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.GetNumStr() {
		t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
	}

	if ePrecision != quotient.GetPrecision() {
		t.Errorf("Expected quotient.GetPrecision()= '%v' .  Instead, quotient.GetPrecision()= '%v'  .", ePrecision, quotient.GetPrecision())
	}

	if eSignVal != quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
	}

}

func TestIntAry_DivideThisBy_12(t *testing.T) {
	// Testing condition where maxPrecision is
	// greater than than significant result.
	dividend := "1"
	divisor := "25"
	// actual result is "0.04"
	eQuotient := "0.04"
	eSignVal := 1
	maxPrecision := 10 // maxPrecision exceeds actual result
	ePrecision := 2

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.GetNumStr() {
		t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
	}

	if ePrecision != quotient.GetPrecision() {
		t.Errorf("Expected quotient.GetPrecision()= '%v' .  Instead, quotient.GetPrecision()= '%v'  .", ePrecision, quotient.GetPrecision())
	}

	if eSignVal != quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
	}

}

func TestIntAry_DivideThisBy_13(t *testing.T) {
	// Testing condition where maxPrecision is
	// set equal to -1.
	dividend := "1"
	divisor := "26"
	eQuotient := "0.0384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615384615385"
	eSignVal := 1
	maxPrecision := -1 // maxPrecision exceeds actual result
	ePrecision := 1024

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.GetNumStr() {
		t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
	}

	if ePrecision != quotient.GetPrecision() {
		t.Errorf("Expected quotient.GetPrecision()= '%v' .  Instead, quotient.GetPrecision()= '%v'  .", ePrecision, quotient.GetPrecision())
	}

	if eSignVal != quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
	}

}

func TestIntAry_DivideThisBy_14(t *testing.T) {
	// Testing condition where maxPrecision is
	// invalid
	dividend := "1"
	divisor := "25"
	maxPrecision := -10 // maxPrecision exceeds actual result

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	_, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err == nil {
		t.Error("Expected error to result from invalid 'maxPrecision' value of -10. Instead, no such error was triggered")
	}

}

func TestIntAry_DivideThisBy_15(t *testing.T) {
	// Test Condition were actual decimal places are
	// 21+ and maxPrecision is set to zero
	dividend := "56234369384300"
	divisor := "24"
	// actual eQuotient := "2343098724345.833333333333333333333"
	eQuotient := "2343098724346"
	eSignVal := 1
	maxPrecision := 0
	ePrecision := 0

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.GetNumStr() {
		t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
	}

	if ePrecision != quotient.GetPrecision() {
		t.Errorf("Expected quotient.GetPrecision()= '%v' .  Instead, quotient.GetPrecision()= '%v'  .", ePrecision, quotient.GetPrecision())
	}

	if eSignVal != quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
	}

}


func TestIntAry_DivideThisBy_16(t *testing.T) {
	// Test Condition were actual decimal places are
	// 21+ and maxPrecision is set to zero
	dividend := "1"
	divisor := "5132188731375616"
	eQuotient := "0.0000000000000001948486410654588479586365019367735312710807776562419879994613957690782179395268844976594976574190856910462423734806284591100904041759223999872534562498706764291924041486253101558649510819197172143928398136686955522152677536556303208645361801676397426605377561320278442021917586553676450097223493264453235937316338207393368423631288419620924881566999933402357520007410670136732435322811808257677794240393475046309703702321880635871858096536620946259448966209463978714360724504546027791919182097769678551824119271867445986093109132873712847086353082760161728835291168886244778340678734742887929927691392086547229915192975721487525948779793899832544746398668414935327086017026715079825557839286764777130808629545941416744923613139031065794055208370453857117267343145382178198213289314418222458170836304178717692361810139999277887517350389701068636718076181550016169308117583049414576551610275225085122626807472053468620691915924510100661007669795710586291378217910587138791895083422555361272265231787002422046566"
	eSignVal := 1
	maxPrecision := 1024
	ePrecision := 1024

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(dividend)
	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(divisor)

	quotient, err := ia1.DivideThisBy(&ia2, maxPrecision)

	if err != nil {
		t.Errorf("Error returned from ia1.DivideThisBy(&ia2, maxPrecision). Error= %v", err)
	}

	if eQuotient != quotient.GetNumStr() {
		t.Errorf("Expected quotient.GetNumStr()= '%v' .  Instead, quotient.GetNumStr()= '%v'  .", eQuotient, quotient.GetNumStr())
	}

	if ePrecision != quotient.GetPrecision() {
		t.Errorf("Expected quotient.GetPrecision()= '%v' .  Instead, quotient.GetPrecision()= '%v'  .", ePrecision, quotient.GetPrecision())
	}

	if eSignVal != quotient.GetSign() {
		t.Errorf("Error - Expected smop.Quotient.GetSign()= '%v'. Instead, smop.Quotient.GetSign()= '%v' .", eSignVal, quotient.GetSign())
	}

}
