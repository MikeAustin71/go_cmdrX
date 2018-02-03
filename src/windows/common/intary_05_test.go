package common

import (
	"testing"
	"math/big"
)

func TestIntAry_MultiplyByTenToPower_01(t *testing.T) {

	nStr := "457.3"
	eNumStr := "45730"
	eIAry := []uint8{4, 5, 7, 3, 0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(2)

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

			t.Errorf("Error: Expected intAry Array does NOT match ia intAry! index='%v'", i)
			return

		}
	}

}

func TestIntAry_MultiplyByTenToPower_02(t *testing.T) {

	nStr := "457.3"
	power := uint(2)
	eNumStr := "45730"
	eIAry := []uint8{4, 5, 7, 3, 0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power)

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

			t.Errorf("Error: Expected intAry Array does NOT match ia IntAry! Index='%v'", i)
			return

		}
	}
}

func TestIntAry_MultiplyByTenToPower_03(t *testing.T) {

	nStr := "457.3"
	power := uint(10)
	eNumStr := "4573000000000"
	eIAry := []uint8{4, 5, 7, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power)

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
		if element, _:= ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia IntAry element! index='%v'", i)
			return

		}
	}

}

func TestIntAry_MultiplyByTenToPower_04(t *testing.T) {

	nStr := "457.3"
	power := uint(0)
	eNumStr := "457.3"
	eIAry := []uint8{4, 5, 7, 3}
	lEArray := len(eIAry)
	ePrecision := 1
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power)

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
		if element, _:= ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia IntAry element! index='%v'", i)
			return

		}
	}
}

func TestIntAry_MultiplyByTenToPower_05(t *testing.T) {

	nStr := "-457.3"
	power := uint(1)
	eNumStr := "-4573"
	eIAry := []uint8{4, 5, 7, 3}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := -1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power)

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
		if element, _:= ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia IntAry element! index='%v' ", i)
			return

		}
	}
}

func TestIntAry_MultiplyByTenToPower_06(t *testing.T) {

	nStr := "0"
	power := uint(2)
	eNumStr := "000"
	eIAry := []uint8{0, 0, 0}
	lEArray := len(eIAry)
	ePrecision := 0
	eSignVal := 1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.MultiplyByTenToPower(power)

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
		if element, _:= ia.GetIntAryElement(i); eIAry[i] != element {

			t.Errorf("Error: Expected intAry Array does NOT match ia IntAry element! index='%v' ", i)
			return

		}
	}
}

func TestIntAry_MultiplyByTwoToPower_01(t *testing.T) {
	nStr1 := "23"
	expected := "12058624"
	power := uint(19)
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.MultiplyByTwoToPower(power)

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= %v. Instead, ia.GetNumStr()= %v", expected, ia.GetNumStr())
	}

}

func TestIntAry_MultiplyByTwoToPower_02(t *testing.T) {
	nStr1 := "23"
	expected := "23"
	power := uint(0)
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.MultiplyByTwoToPower(power)

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= %v. Instead, ia.GetNumStr()= %v", expected, ia.GetNumStr())
	}

}

func TestIntAry_MultiplyByTwoToPower_03(t *testing.T) {
	nStr1 := "23"
	expected := "46"
	power := uint(1)
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.MultiplyByTwoToPower(power)

	if expected != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()= %v. Instead, ia.GetNumStr()= %v", expected, ia.GetNumStr())
	}

}


func TestIntAry_MultiplyThisBy_01(t *testing.T) {

	this, _ := IntAry{}.NewNumStr("96.8524")
	ia2, _ := IntAry{}.NewNumStr("8574.21396845")
	expected := "830433.200957906780"

	this.MultiplyThisBy(&ia2, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

}

func TestIntAry_MultiplyThisBy_02(t *testing.T) {

	this, _ := IntAry{}.NewNumStr("-96.8524")
	ia2, _ := IntAry{}.NewNumStr("8574.21396845")
	expected := "-830433.200957906780"

	this.MultiplyThisBy(&ia2, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

}

func TestIntAry_MultiplyThisBy_03(t *testing.T) {

	this, _ := IntAry{}.NewNumStr("96.8524")
	ia2, _ := IntAry{}.NewNumStr("-8574.21396845")
	expected := "-830433.200957906780"

	this.MultiplyThisBy(&ia2, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

}

func TestIntAry_MultiplyThisBy_04(t *testing.T) {

	this, _ := IntAry{}.NewNumStr("-96.8524")
	ia2, _ := IntAry{}.NewNumStr("-8574.21396845")
	expected := "830433.200957906780"

	this.MultiplyThisBy(&ia2, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

}

func TestIntAry_MultiplyThisBy_05(t *testing.T) {

	this, _ := IntAry{}.NewNumStr("96.8524")
	ia2, _ := IntAry{}.NewNumStr("8574.21396845")
	expected := "830433.200958"
	maxPrecision := 6

	this.MultiplyThisBy(&ia2, maxPrecision)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

	if maxPrecision != this.GetPrecision() {
		t.Errorf("Expected this.GetPrecision()= %v.  Instead, this.GetPrecision()= %v", maxPrecision, this.GetPrecision())
	}

}


func TestIntAry_MultiplyThisBy_06(t *testing.T) {

	this, _ := IntAry{}.NewNumStr("0")
	ia2, _ := IntAry{}.NewNumStr("-8574.21396845")
	expected := "0.00000000"

	this.MultiplyThisBy(&ia2, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

}

func TestIntAry_MultiplyThisBy_07(t *testing.T) {

	this, _ := IntAry{}.NewNumStr("96.8524")
	ia2, _ := IntAry{}.NewNumStr("0")
	expected := "0"

	this.MultiplyThisBy(&ia2, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

}

func TestIntAry_MultiplyThisBy_08(t *testing.T) {

	this, _ := IntAry{}.NewNumStr("40")
	ia2, _ := IntAry{}.NewNumStr("2")
	expected := "80"

	this.MultiplyThisBy(&ia2, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

	thisStats := this.GetIntAryStats()

	if thisStats.IntAryLen != 2 {
		t.Errorf("Expected this this IntAryLen = 2 .  Instead this IntAryLen= %v .", thisStats.IntAryLen)
	}

	if thisStats.IntegerLen != 2 {
		t.Errorf("Expected this this IntegerLen = 2 .  Instead this IntegerLen= %v .", thisStats.IntegerLen)
	}

	if this.GetPrecision() != 0 {
		t.Errorf("Expected this this.GetPrecision() = 0 .  Instead this.GetPrecision= %v .", this.GetPrecision)
	}

}

func TestIntAry_MultiplyThisBy_09(t *testing.T) {

	this, _ := IntAry{}.NewNumStr("999.99")
	ia2, _ := IntAry{}.NewNumStr("99.9")
	expected := "99899.001"

	this.MultiplyThisBy(&ia2, -1)

	resultStr := this.GetNumStr()

	if expected != resultStr {
		t.Errorf("Expected this.GetNumStr()= %v.  Instead, this.GetNumStr()= %v", expected, resultStr)
	}

	thisStats := this.GetIntAryStats()

	if thisStats.IntAryLen != 8 {
		t.Errorf("Expected this this IntAryLen = 8 .  Instead this IntAryLen= %v .", thisStats.IntAryLen)
	}

	if thisStats.IntegerLen != 5 {
		t.Errorf("Expected this this IntegerLen = 5 .  Instead this IntegerLen= %v .", thisStats.IntegerLen)
	}

	if this.GetPrecision() != 3 {
		t.Errorf("Expected this this.GetPrecision() = 3 .  Instead this.GetPrecision= %v .", this.GetPrecision)
	}

}


func TestIntAry_NewNumStr_01(t *testing.T) {
	nStr1 := "579.123456000"
	ePrecision := 9
	eSignVal := 1

	ia, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewNumStr(nStr1). Error= %v", err)
	}

	if nStr1 != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", nStr1, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewNumStr_02(t *testing.T) {
	nStr1 := "-579.123456000"
	ePrecision := 9
	eSignVal := -1

	ia, err := IntAry{}.NewNumStr(nStr1)

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewNumStr(nStr1). Error= %v", err)
	}

	if nStr1 != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", nStr1, ia.GetNumStr())
	}

	if ePrecision != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", ePrecision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewBigInt_01(t *testing.T) {
	num := big.NewInt(123456)
	precision := uint(3)
	ia, err := IntAry{}.NewBigInt(num, precision)
	eStr := "123.456"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewBigInt(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewBigInt_02(t *testing.T) {
	num := big.NewInt(-123456)
	precision := uint(3)
	ia, err := IntAry{}.NewBigInt(num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewBigInt(num, precision). num= %v  precision= %v Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewFloatBig_01(t *testing.T) {
	num := big.NewFloat(123.456)
	precision := 3
	ia, err := IntAry{}.NewFloatBig(num, 3)
	eStr := "123.456"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloatBig(num). num= %v Error= %v", num, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewFloatBig_02(t *testing.T) {
	num := big.NewFloat(-123.456)
	precision := 3
	ia, err := IntAry{}.NewFloatBig(num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloatBig(num). num= %v Error= %v", num, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewFloatBig_03(t *testing.T) {
	num := big.NewFloat(123.456000)
	precision := 3
	ia, err := IntAry{}.NewFloatBig(num, -1)
	eStr := "123.456"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloatBig(num). num= %v Error= %v", num, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewFloatBig_04(t *testing.T) {
	num := big.NewFloat(123.456700)
	precision := 3
	ia, err := IntAry{}.NewFloatBig(num, 3)
	eStr := "123.457"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloatBig(num). num= %v Error= %v", num, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}


func TestIntAry_NewFloat64_01(t *testing.T) {
	num := float64(123.456111)
	precision := 3
	ia, err := IntAry{}.NewFloat64(num, precision)
	eStr := "123.456"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloat64(num, precision). num= %v precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}
}

func TestIntAry_NewFloat64_02(t *testing.T) {
	num := float64(-123.456111)
	precision := 3
	ia, err := IntAry{}.NewFloat64(num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloat64(num, precision). num= %v precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}
}

func TestIntAry_NewFloat32_01(t *testing.T) {
	num := float64(123.456111)
	precision := 3
	ia, err := IntAry{}.NewFloat64(num, precision)
	eStr := "123.456"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloat32(num, precision). num= %v precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}
}

func TestIntAry_NewFloat32_02(t *testing.T) {
	num := float64(-123.456111)
	precision := 3
	ia, err := IntAry{}.NewFloat64(num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewFloat32(num, precision). num= %v precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}
}

func TestIntAry_NewInt64_01(t *testing.T) {
	num := int64(123456)
	precision := uint(3)
	ia, err := IntAry{}.NewInt64(num, precision)
	eStr := "123.456"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewInt64(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt64_02(t *testing.T) {
	num := int64(-123456)
	precision := uint(3)
	ia, err := IntAry{}.NewInt64(num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewInt64(num, precision). num= %v  precision= %v Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt32_01(t *testing.T) {
	num := int32(123456)
	precision := uint(3)
	ia, err := IntAry{}.NewInt32(num, precision)
	eStr := "123.456"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewInt32(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt32_02(t *testing.T) {
	num := int32(-123456)
	precision := uint(3)
	ia, err := IntAry{}.NewInt32(num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewInt32(num, precision). num= %v  precision= %v Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt_01(t *testing.T) {
	num := int(123456)
	precision := uint(3)
	ia, err := IntAry{}.NewInt(num, precision)
	eStr := "123.456"
	eSignVal := 1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewInt(num, precision). num= %v  precision= %v  Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_NewInt_02(t *testing.T) {
	num := int(-123456)
	precision := uint(3)
	ia, err := IntAry{}.NewInt(num, precision)
	eStr := "-123.456"
	eSignVal := -1

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewInt32(num, precision). num= %v  precision= %v Error= %v", num, precision, err)
	}

	if eStr != ia.GetNumStr() {
		t.Errorf("Expected ia.GetNumStr()== %v  .   Instead ia.GetNumStr() == %v", eStr, ia.GetNumStr())
	}

	if int(precision) != ia.GetPrecision() {
		t.Errorf("Expected ia.GetPrecision() == %v  .   Instead ia.GetPrecision() == %v", precision, ia.GetPrecision())
	}

	if eSignVal != ia.GetSign() {
		t.Errorf("Expected ia.GetSign() == %v  .   Instead ia.GetSign() == %v", eSignVal, ia.GetSign())
	}

}

func TestIntAry_OptimizeIntArrayLen_01(t *testing.T) {
	nStr1 := "00579.123456000"
	expected := "579.123456000"
	ePrecision := 9
	eLen := 12

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	ia.OptimizeIntArrayLen(false)

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

func TestIntAry_OptimizeIntArrayLen_02(t *testing.T) {
	nStr1 := "00579.123456000"
	expected := "579.123456"
	ePrecision := 6
	eLen := 9

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

func TestIntAry_OptimizeIntArrayLen_03(t *testing.T) {
	nStr1 := "00579.000000000"
	expected := "579"
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

func TestIntAry_OptimizeIntArrayLen_04(t *testing.T) {
	nStr1 := "00000.123450000"
	expected := "0.12345"
	ePrecision := 5
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

func TestIntAry_OptimizeIntArrayLen_05(t *testing.T) {
	nStr1 := "00000.000123450000"
	expected := "0.00012345"
	ePrecision := 8
	eLen := 9

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

func TestIntAry_OptimizeIntArrayLen_06(t *testing.T) {
	nStr1 := "0.0"
	expected := "0.0"
	ePrecision := 1
	eLen := 2

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


