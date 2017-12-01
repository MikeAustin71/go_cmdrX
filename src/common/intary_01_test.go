package common

import (
	"testing"
)


func TestIntAry_AddMultipleToThis_01(t *testing.T) {

	nStr1 := "457.3"
	nStr2 := "82.975"
	nStr3 := "94"
	nStr4 := "697.21589"
	nStr5 := "9648.37"

	expected := "10979.86089"

	ia0 := IntAry{}.New()
	ia0.SetIntAryToZero(0)

	ia1 := IntAry{}.New()
	ia1.SetIntAryWithNumStr(nStr1)

	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(nStr2)

	ia3 := IntAry{}.New()
	ia3.SetIntAryWithNumStr(nStr3)

	ia4 := IntAry{}.New()
	ia4.SetIntAryWithNumStr(nStr4)

	ia5 := IntAry{}.New()
	ia5.SetIntAryWithNumStr(nStr5)

	ia0.AddMultipleToThis(&ia1, &ia2, &ia3, &ia4, &ia5)

	if expected != ia0.GetNumStr() {
		t.Errorf("Error: Expected ia.GetNumStr()= '%v' .  Instead, ia.GetNumStr()= '%v' . ", expected, ia0.GetNumStr())
	}

}

func TestIntAry_AddToThis_01(t *testing.T) {
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "457.3"
	nStr2 := "22.2"
	expected := "479.5"
	eIAry := []uint8{4, 7, 9, 5}
	lEArray := len(eIAry)
	ePrecision := 1
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from mOps.AddN1N2(). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	ia1Stats := ia1.GetIntAryStats()

	if lEArray != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1Stats.IntAryLen)
	}

	for i := 0; i < lEArray; i++ {

		if element, _ := ia1.GetIntAryElement(i);  eIAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 IntAry Element! ")
			return

		}
	}

}

func TestIntAry_AddToThis_02(t *testing.T) {
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "457.325"
	nStr2 := "-22.2"
	expected := "435.125"
	eIAry := []uint8{4, 3, 5, 1, 2, 5}
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	ia1Stats := ia1.GetIntAryStats()

	if lEArray != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1Stats.IntAryLen)
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 Int Array Element! ")
			return

		}
	}

}

func TestIntAry_AddToThis_03(t *testing.T) {
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-457.325"
	nStr2 := "22.2"
	expected := "-435.125"
	eIAry := []uint8{4, 3, 5, 1, 2, 5}
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	ia1Stats := ia1.GetIntAryStats()

	if lEArray != ia1Stats.IntAryLen {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1Stats.IntAryLen)
	}

	var element uint8

	for i := 0; i < lEArray; i++ {
		if element, _ = ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Error("Error: Expected intAry Array does NOT match ia1 Int Array Element! ")
			return

		}
	}
}

func TestIntAry_AddToThis_04(t *testing.T) {
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-457.325"
	nStr2 := "-22.2"
	expected := "-479.525"
	eIAry := []uint8{4, 7, 9, 5, 2, 5}
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia1 integer array element! index='%v' ", i)
			return

		}
	}
}

func TestIntAry_AddToThis_05(t *testing.T) {
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "0.000"
	nStr2 := "-22.2"
	expected := "-22.200"
	eIAry := []uint8{2, 2, 2, 0, 0}
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i) ; eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia1 integer array element. index='%v' ", i)
			return

		}
	}
}

func TestIntAry_AddToThis_06(t *testing.T) {
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "0.000"
	nStr2 := "0"
	expected := "0.000"
	eIAry := []uint8{0, 0, 0, 0}
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia1 integer array element. index='%v' ", i)
			return

		}
	}
}

func TestIntAry_AddToThis_07(t *testing.T) {
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "99.225"
	nStr2 := "-99.1"
	expected := "0.125"
	eIAry := []uint8{0, 1, 2, 5}
	lEArray := len(eIAry)
	ePrecision := 3
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {

		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element  {
			t.Errorf("Error: Expected intAry Array does NOT match ia1 integer array element. index='%v'", i)
		}
	}
}

func TestIntAry_AddToThis_08(t *testing.T) {
	// N1 > N2 + and +
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "350"
	nStr2 := "122"
	expected := "472"
	eIAry := []uint8{4, 7, 2}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {
			t.Errorf("Error: Expected intAry Array does NOT match ia1 integer array element! index='%v' ", i)
		}
	}
}

func TestIntAry_AddToThis_09(t *testing.T) {
	// N1 > N2 - and +
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-350"
	nStr2 := "122"
	expected := "-228"
	eIAry := []uint8{2, 2, 8}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_AddToThis_10(t *testing.T) {
	// N1 > N2 - and -
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-350"
	nStr2 := "-122"
	expected := "-472"
	eIAry := []uint8{4, 7, 2}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_AddToThis_11(t *testing.T) {
	// N1 > N2 + and -
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "350"
	nStr2 := "-122"
	expected := "228"
	eIAry := []uint8{2, 2, 8}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v' ", i)
			return

		}
	}
}

func TestIntAry_AddToThis_12(t *testing.T) {
	// N1 > N2  350 + 0
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "350"
	nStr2 := "0"
	expected := "350"
	eIAry := []uint8{3, 5, 0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_AddToThis_13(t *testing.T) {
	// N1 > N2  -350 + 0
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-350"
	nStr2 := "0"
	expected := "-350"
	eIAry := []uint8{3, 5, 0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_AddToThis_14(t *testing.T) {
	// N2 > N1  + and +
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "122"
	nStr2 := "350"
	expected := "472"
	eIAry := []uint8{4, 7, 2}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v' ", i)
			return

		}
	}
}

func TestIntAry_AddToThis_15(t *testing.T) {
	// N2 > N1  - and +
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-122"
	nStr2 := "350"
	expected := "228"
	eIAry := []uint8{2, 2, 8}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_AddToThis_16(t *testing.T) {
	// N2 > N1  - and -
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-122"
	nStr2 := "-350"
	expected := "-472"
	eIAry := []uint8{4, 7, 2}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_AddToThis_17(t *testing.T) {
	// N2 > N1  + and -
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "122"
	nStr2 := "-350"
	expected := "-228"
	eIAry := []uint8{2, 2, 8}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'",i)
			return

		}
	}
}

func TestIntAry_AddToThis_18(t *testing.T) {
	// N2 > N1  0 and +350
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "0"
	nStr2 := "350"
	expected := "350"
	eIAry := []uint8{3, 5, 0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_AddToThis_19(t *testing.T) {
	// N2 > N1  0 and -350
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "0"
	nStr2 := "-350"
	expected := "-350"
	eIAry := []uint8{3, 5, 0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_AddToThis_20(t *testing.T) {
	// N1 == N2  +122 and +122
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "122"
	nStr2 := "122"
	expected := "244"
	eIAry := []uint8{2, 4, 4}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_AddToThis_21(t *testing.T) {
	// N1 == N2  -122 and +122
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-122"
	nStr2 := "122"
	expected := "0"
	eIAry := []uint8{0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}

}

func TestIntAry_AddToThis_22(t *testing.T) {
	// N1 == N2  -122 and -122
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "-122"
	nStr2 := "-122"
	expected := "-244"
	eIAry := []uint8{2, 4, 4}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_AddToThis_23(t *testing.T) {
	// N1 == N2  +122 and -122
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "122"
	nStr2 := "-122"
	expected := "0"
	eIAry := []uint8{0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_AddToThis_24(t *testing.T) {
	// N1 == N2  0 and 0
	ia1 := IntAry{}.New()
	ia2 := IntAry{}.New()
	nStr1 := "0"
	nStr2 := "0"
	expected := "0"
	eIAry := []uint8{0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia1.SetIntAryWithNumStr(nStr1)

	ia2.SetIntAryWithNumStr(nStr2)

	err := ia1.AddToThis(&ia2)

	if err != nil {
		t.Errorf("Received Error from ia1.AddToThis(&ia2, true). nStr1= '%v' nStr2= '%v' Error= %v", nStr1, nStr2, err)
	}

	s := ia1.GetNumStr()

	if s != expected {
		t.Errorf("Expected IFinal.GetNumStr()= '%v'. Instead got IFinal.Numstr= '%v' ", expected, s)
	}

	if ia1.GetPrecision() != ePrecision {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", ePrecision, ia1.GetPrecision())
	}

	if eSignVal != ia1.GetSign() {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", eSignVal, ia1.GetSign())
	}

	if lEArray != ia1.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia1.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {
		if element, _ := ia1.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia integer array element! index='%v'", i)
			return

		}
	}
}
