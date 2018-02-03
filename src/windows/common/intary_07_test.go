package common

import (

	"testing"
)



func TestIntAry_SetIntAryWithInt32_01(t *testing.T) {

	num := int32(123456789)

	eNumStr := "123456.789"
	ePrecision := uint(3)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt32(num, ePrecision)

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

func TestIntAry_SetIntAryWithInt32_02(t *testing.T) {

	num := int32(-123456789)

	eNumStr := "-12345.6789"
	ePrecision := uint(4)
	eSignVal := -1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt32(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt32(num, ePrecision). num= %v  ePrecision= %v  .", num, ePrecision)
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

func TestIntAry_SetIntAryWithInt32_03(t *testing.T) {

	num := int32(0)

	eNumStr := "0.0000"
	ePrecision := uint(4)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt32(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt32(num, ePrecision). num= %v  ePrecision= %v  .", num, ePrecision)
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

func TestIntAry_SetIntAryWithInt32_04(t *testing.T) {

	num := int32(32)

	eNumStr := "0.0032"
	ePrecision := uint(4)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt32(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt32(num, ePrecision). num= %v  ePrecision= %v  .", num, ePrecision)
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

func TestIntAry_SetIntAryWithInt32_05(t *testing.T) {

	num := int32(-32)

	eNumStr := "-32"
	ePrecision := uint(0)
	eSignVal := -1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt32(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt32(num, ePrecision). num= %v  ePrecision= %v  .", num, ePrecision)
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

func TestIntAry_SetIntAryWithInt32_06(t *testing.T) {

	num := int32(32)

	eNumStr := "0.32"
	ePrecision := uint(2)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt32(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt32(num, ePrecision). num= %v  ePrecision= %v  .", num, ePrecision)
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

func TestIntAry_SetIntAryWithUint64_01(t *testing.T) {

	num := uint64(123456789)

	eNumStr := "123456.789"
	ePrecision := uint(3)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithUint64(num, ePrecision, eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithUint64(num, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", num, ePrecision, eSignVal)
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

func TestIntAry_SetIntAryWithUint64_02(t *testing.T) {

	num := uint64(123456789)

	eNumStr := "-12345.6789"
	ePrecision := uint(4)
	eSignVal := -1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithUint64(num, ePrecision, eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithUint64(num, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", num, ePrecision, eSignVal)
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

func TestIntAry_SetIntAryWithUint64_03(t *testing.T) {

	num := uint64(0)

	eNumStr := "0.0000"
	ePrecision := uint(4)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithUint64(num, ePrecision, eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithUint64(num, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", num, ePrecision, eSignVal)
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

func TestIntAry_SetIntAryWithUint64_04(t *testing.T) {

	num := uint64(32)

	eNumStr := "0.0032"
	ePrecision := uint(4)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithUint64(num, ePrecision, eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithUint64(num, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", num, ePrecision, eSignVal)
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

func TestIntAry_SetIntAryWithUint64_05(t *testing.T) {

	num := uint64(32)

	eNumStr := "-32"
	ePrecision := uint(0)
	eSignVal := -1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithUint64(num, ePrecision, eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithUint64(num, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", num, ePrecision, eSignVal)
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

func TestIntAry_SetIntAryWithUint64_06(t *testing.T) {

	num := uint64(32)

	eNumStr := "0.32"
	ePrecision := uint(2)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithUint64(num, ePrecision, eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithUint64(num, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", num, ePrecision, eSignVal)
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

func TestIntAry_SetIntAryWithInt64_01(t *testing.T) {

	num := int64(123456789)

	eNumStr := "123456.789"
	ePrecision := uint(3)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt64(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt64(num, ePrecision). num= %v  ePrecision= %v  .", num, ePrecision)
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

func TestIntAry_SetIntAryWithInt64_02(t *testing.T) {

	num := int64(-123456789)

	eNumStr := "-12345.6789"
	ePrecision := uint(4)
	eSignVal := -1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt64(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt64(num, ePrecision). num= %v  ePrecision= %v .", num, ePrecision)
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

func TestIntAry_SetIntAryWithInt64_03(t *testing.T) {

	num := int64(0)

	eNumStr := "0.0000"
	ePrecision := uint(4)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt64(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt64(num, ePrecision). num= %v  ePrecision= %v  .", num, ePrecision)
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

func TestIntAry_SetIntAryWithInt64_04(t *testing.T) {

	num := int64(32)

	eNumStr := "0.0032"
	ePrecision := uint(4)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt64(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt64(num, ePrecision, eSignVal). num= %v  ePrecision= %v  .", num, ePrecision)
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

func TestIntAry_SetIntAryWithInt64_05(t *testing.T) {

	num := int64(-32)

	eNumStr := "-32"
	ePrecision := uint(0)
	eSignVal := -1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt64(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt64(num, ePrecision). num= %v  ePrecision= %v  .", num, ePrecision)
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

func TestIntAry_SetIntAryWithInt64_06(t *testing.T) {

	num := int64(32)

	eNumStr := "0.32"
	ePrecision := uint(2)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithInt64(num, ePrecision)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithInt64(num, ePrecision). num= %v  ePrecision= %v  .", num, ePrecision)
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

func TestIntAry_SetIntAryWithIntAry_01(t *testing.T) {
	iAry := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	eNumStr := "123456.789"
	ePrecision := uint(3)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", iAry, ePrecision, eSignVal)
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

	iaStats := ia.GetIntAryStats()

	resultAry, iaLen := ia.GetIntAry()

	if iaStats.IntAryLen != iaLen {
		t.Errorf("Expected iaStats.IntAryLen == results array length. Instead, iaStats.IntAryLen='%v' and result int array length='%v'", iaStats.IntAryLen, iaLen)
	}

	for i:=0; i < iaLen; i++ {
		if uint8(iAry[i]) != resultAry[i] {
			t.Errorf("Expected uint8(iAry[i])==resultAry[i]. Instead i='%v' uint8(iAry[i])='%v' resultAry[i]='%v' ", i, iAry[i], resultAry[i])
			return
		}
	}

}

func TestIntAry_SetIntAryWithIntAry_02(t *testing.T) {
	iAry := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	eNumStr := "-12345.6789"
	ePrecision := uint(4)
	eSignVal := -1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", iAry, ePrecision, eSignVal)
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

	iaStats := ia.GetIntAryStats()

	resultAry, iaLen := ia.GetIntAry()

	if iaStats.IntAryLen != iaLen {
		t.Errorf("Expected iaStats.IntAryLen == results array length. Instead, iaStats.IntAryLen='%v' and result int array length='%v'", iaStats.IntAryLen, iaLen)
	}

	for i:=0; i < iaLen; i++ {
		if uint8(iAry[i]) != resultAry[i] {
			t.Errorf("Expected uint8(iAry[i])==resultAry[i]. Instead i='%v' uint8(iAry[i])='%v' resultAry[i]='%v' ", i, iAry[i], resultAry[i])
			return
		}
	}


}

func TestIntAry_SetIntAryWithIntAry_03(t *testing.T) {
	iAry := []int{3, 2}
	iOutAry := []int{0,0,0,3,2}
	eNumStr := "0.0032"
	ePrecision := uint(4)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", iAry, ePrecision, eSignVal)
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

	iaStats := ia.GetIntAryStats()

	resultAry, iaLen := ia.GetIntAry()

	if iaStats.IntAryLen != iaLen {
		t.Errorf("Expected iaStats.IntAryLen == results array length. Instead, iaStats.IntAryLen='%v' and result int array length='%v'", iaStats.IntAryLen, iaLen)
	}

	for i:=0; i < iaLen; i++ {
		if uint8(iOutAry[i]) != resultAry[i] {
			t.Errorf("Expected uint8(iAry[i])==resultAry[i]. Instead i='%v' uint8(iAry[i])='%v' resultAry[i]='%v' ", i, iAry[i], resultAry[i])
			return
		}
	}

}

func TestIntAry_SetIntAryWithIntAry_04(t *testing.T) {
	iAry := []int{3, 2}
	iOutAry := []int{0,3,2}
	eNumStr := "0.32"
	ePrecision := uint(2)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithIntAry(iAry, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", iAry, ePrecision, eSignVal)
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

	iaStats := ia.GetIntAryStats()

	resultAry, iaLen := ia.GetIntAry()

	if iaStats.IntAryLen != iaLen {
		t.Errorf("Expected iaStats.IntAryLen == results array length. Instead, iaStats.IntAryLen='%v' and result int array length='%v'", iaStats.IntAryLen, iaLen)
	}

	for i:=0; i < iaLen; i++ {
		if uint8(iOutAry[i]) != resultAry[i] {
			t.Errorf("Expected uint8(iAry[i])==resultAry[i]. Instead i='%v' uint8(iAry[i])='%v' resultAry[i]='%v' ", i, iAry[i], resultAry[i])
			return
		}
	}

}

func TestIntAry_SetIntAryWithIntAryObj_01(t *testing.T) {

	numStr1 := "0000589432.607528000"
	numStr2 := "-028.3700"

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(numStr1)

	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(numStr2)

	err := ia2.SetIntAryWithIntAryObj(&ia, false)

	if err != nil {
		t.Errorf("Received Error from ia2.SetIntAryWithIntAryObj(&ia, false). Error= %v", err)
	}

	if !ia.Equals(&ia2) {
		t.Error("ia is NOT EQUAL to ia2!!!")
	}

}

func TestIntAry_SetIntAryWithIntAryObj_02(t *testing.T) {

	numStr1 := "0000589432.607528000"
	numStr2 := "-028.3700"

	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(numStr1)

	ia2 := IntAry{}.New()
	ia2.SetIntAryWithNumStr(numStr2)

	err := ia2.SetIntAryWithIntAryObj(&ia, true)

	if err != nil {
		t.Errorf("Received Error from ia2.SetIntAryWithIntAryObj(&ia, true). Error= %v", err)
	}

	if !ia.Equals(&ia2) {
		t.Error("ia is NOT EQUAL to ia2 !!!")
	}

	if !ia.BackUp.Equals(&ia2.BackUp) {
		t.Error("ia.BackUp is NOT EQUAL to ia2.BackUp !!!")
	}

}

func TestIntAry_SetIntAryWithUint8Ary_01(t *testing.T) {
	iAry := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	eNumStr := "123456.789"
	ePrecision := uint(3)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithUint8Ary(iAry, ePrecision, eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithUint8Ary(iAry, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", iAry, ePrecision, eSignVal)
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

	iaStats := ia.GetIntAryStats()

	resultAry, iaLen := ia.GetIntAry()

	if iaStats.IntAryLen != iaLen {
		t.Errorf("Expected iaStats.IntAryLen == results array length. Instead, iaStats.IntAryLen='%v' and result int array length='%v'", iaStats.IntAryLen, iaLen)
	}

	for i:=0; i < iaLen; i++ {
		if iAry[i] != resultAry[i] {
			t.Errorf("Expected iAry[i]==resultAry[i]. Instead i='%v' iAry[i]='%v' resultAry[i]='%v' ", i, iAry[i], resultAry[i])
			return
		}
	}

}

func TestIntAry_SetIntAryWithUint8Ary_02(t *testing.T) {
	iAry := []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9}
	eNumStr := "-12345.6789"
	ePrecision := uint(4)
	eSignVal := -1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithUint8Ary(iAry, ePrecision, eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithUint8Ary(iAry, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", iAry, ePrecision, eSignVal)
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

	iaStats := ia.GetIntAryStats()

	resultAry, iaLen := ia.GetIntAry()

	if iaStats.IntAryLen != iaLen {
		t.Errorf("Expected iaStats.IntAryLen == results array length. Instead, iaStats.IntAryLen='%v' and result int array length='%v'", iaStats.IntAryLen, iaLen)
	}

	for i:=0; i < iaLen; i++ {
		if iAry[i] != resultAry[i] {
			t.Errorf("Expected iAry[i]==resultAry[i]. Instead i='%v' iAry[i]='%v' resultAry[i]='%v' ", i, iAry[i], resultAry[i])
			return
		}
	}


}

func TestIntAry_SetIntAryWithUint8Ary_03(t *testing.T) {
	iAry := []uint8{3, 2}
	iOutAry := []uint8{0,0,0,3,2}
	eNumStr := "0.0032"
	ePrecision := uint(4)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithUint8Ary(iAry, ePrecision, eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithUint8Ary(iAry, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", iAry, ePrecision, eSignVal)
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

	iaStats := ia.GetIntAryStats()

	resultAry, iaLen := ia.GetIntAry()

	if iaStats.IntAryLen != iaLen {
		t.Errorf("Expected iaStats.IntAryLen == results array length. Instead, iaStats.IntAryLen='%v' and result int array length='%v'", iaStats.IntAryLen, iaLen)
	}

	for i:=0; i < iaLen; i++ {
		if iOutAry[i] != resultAry[i] {
			t.Errorf("Expected iAry[i]==resultAry[i]. Instead i='%v' uint8(iAry[i])='%v' resultAry[i]='%v' ", i, iAry[i], resultAry[i])
			return
		}
	}

}

func TestIntAry_SetIntAryWithUint8Ary_04(t *testing.T) {
	iAry := []uint8{3, 2}
	iOutAry := []uint8{0,3,2}
	eNumStr := "0.32"
	ePrecision := uint(2)
	eSignVal := 1

	ia := IntAry{}.New()

	err := ia.SetIntAryWithUint8Ary(iAry, ePrecision, eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetIntAryWithUint8Ary(iAry, ePrecision, eSignVal). num= %v  ePrecision= %v  eSignVal= %v", iAry, ePrecision, eSignVal)
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

	iaStats := ia.GetIntAryStats()

	resultAry, iaLen := ia.GetIntAry()

	if iaStats.IntAryLen != iaLen {
		t.Errorf("Expected iaStats.IntAryLen == results array length. Instead, iaStats.IntAryLen='%v' and result int array length='%v'", iaStats.IntAryLen, iaLen)
	}

	for i:=0; i < iaLen; i++ {
		if iOutAry[i] != resultAry[i] {
			t.Errorf("Expected iAry[i]==resultAry[i]. Instead i='%v' iAry[i]='%v' resultAry[i]='%v' ", i, iAry[i], resultAry[i])
			return
		}
	}

}


func TestIntAry_SetSign_01(t *testing.T) {

	nStr := "123.456"

	eStr := "-123.456"
	eSignVal := -1
	ePrecision := 3

	ia, err:= IntAry{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewNumStr(nStr). Error= %v", err)
	}

	err = ia.SetSign(eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetSign(eSignVal). Error= %v", err)
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

func TestIntAry_SetSign_02(t *testing.T) {

	nStr := "123.456"

	eStr := "123.456"
	eSignVal := 1
	ePrecision := 3

	ia, err:= IntAry{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewNumStr(nStr). Error= %v", err)
	}

	err = ia.SetSign(eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetSign(eSignVal). Error= %v", err)
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

func TestIntAry_SetSign_03(t *testing.T) {

	nStr := "-123.456"

	eStr := "123.456"
	eSignVal := 1
	ePrecision := 3

	ia, err:= IntAry{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewNumStr(nStr). Error= %v", err)
	}

	err = ia.SetSign(eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetSign(eSignVal). Error= %v", err)
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

func TestIntAry_SetSign_04(t *testing.T) {

	nStr := "-123.456"

	eStr := "-123.456"
	eSignVal := -1
	ePrecision := 3

	ia, err:= IntAry{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewNumStr(nStr). Error= %v", err)
	}

	err = ia.SetSign(eSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetSign(eSignVal). Error= %v", err)
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

func TestIntAry_SetSign_05(t *testing.T) {

	nStr := "0.0"

	eStr := "0.0"
	attemptedSignVal := -1
	eSignVal := 1
	ePrecision := 1

	ia, err:= IntAry{}.NewNumStr(nStr)

	if err != nil {
		t.Errorf("Error returned from intAry{}.NewNumStr(nStr). Error= %v", err)
	}

	err = ia.SetSign(attemptedSignVal)

	if err != nil {
		t.Errorf("Error returned from ia.SetSign(eSignVal). Error= %v", err)
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


func TestIntAry_SetWithNumStr_01(t *testing.T) {
	nStr := "123.456"
	eIAry := []uint8{1, 2, 3, 4, 5, 6}
	lEArray := len(eIAry)

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	//ia.SetIntAryLength()

	if ia.GetNumStr() != nStr {
		t.Errorf("Error: Expected numStr= '%v'. Instead received numStr= '%v'", nStr, ia.GetNumStr())
	}

	if ia.GetPrecision() != 3 {
		t.Errorf("Error: Expected precision= '%v'. Instead received precision= '%v'", 3, ia.GetPrecision())
	}

	if ia.GetSign() != 1 {
		t.Errorf("Error: Expected signVal= '%v'. Instead received signVal= '%v'", 1, ia.GetSign())
	}

	if lEArray != ia.GetIntAryLength() {
		t.Errorf("Error: Expected IntArray Length= '%v'. Instead received IntArry Length= '%v'", lEArray, ia.GetIntAryLength())
	}

	for i := 0; i < lEArray; i++ {

		if element, _ := ia.GetIntAryElement(i) ; eIAry[i] != element {
	t.Errorf("Error: Expected intAry Array element does NOT match ia IntAry element! Index='%v'", i)
			return

		}
	}
}

func TestIntAry_SetWithNumStr_02(t *testing.T) {
	nStr := "-12345.9"
	eIAry := []uint8{1, 2, 3, 4, 5, 9}
	lEArray := len(eIAry)
	ePrecision := 1
	eSignVal := -1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStr)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.SetIntAryLength()

	if ia.GetNumStr() != nStr {
		t.Errorf("Error: Expected numStr= '%v'. Instead received numStr= '%v'", nStr, ia.GetNumStr())
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

			t.Errorf("Error: Expected intAry Array does NOT match ia IntAry element! Index='%v' ", i)
			return

		}
	}
}

func TestIntAry_SetWithNumStr_03(t *testing.T) {
	nStrRaw := "-123  45.9"
	nStr := "-12345.9"
	eIAry := []uint8{1, 2, 3, 4, 5, 9}
	lEArray := len(eIAry)
	ePrecision := 1
	eSignVal := -1

	ia := IntAry{}.New()
	err := ia.SetIntAryWithNumStr(nStrRaw)

	if err != nil {
		t.Errorf("Received Error from ia.SetIntAryWithNumStr(nStr). nStr= '%v' Error= %v", nStr, err)
	}

	ia.SetIntAryLength()

	if ia.GetNumStr() != nStr {
		t.Errorf("Error: Expected numStr= '%v'. Instead received numStr= '%v'", nStr, ia.GetNumStr())
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

			t.Errorf("Error: Expected intAry Array does NOT match ia IntAry element! index='%v'", i)
			return

		}
	}

}

func TestIntAry_SetPrecision_01(t *testing.T) {
	nStr1 := "99.995"
	expected := "99.99"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.SetPrecision(precision, false)
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

func TestIntAry_SetPrecision_02(t *testing.T) {
	nStr1 := "99.995"
	expected := "100.00"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.SetPrecision(precision, true)
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

func TestIntAry_SetPrecision_03(t *testing.T) {
	nStr1 := "-0"
	expected := "0.00"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.SetPrecision(precision, true)
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

func TestIntAry_SetPrecision_04(t *testing.T) {
	nStr1 := "-999.995"
	expected := "-1000.00"
	precision := 2
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.SetPrecision(precision, true)
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

func TestIntAry_SetPrecision_05(t *testing.T) {
	nStr1 := "-999.995"
	expected := "-999.995"
	precision := 3
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.SetPrecision(precision, true)
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

func TestIntAry_SetPrecision_06(t *testing.T) {
	nStr1 := "-999.995"
	expected := "-999.995000"
	precision := 6
	ia := IntAry{}.New()
	ia.SetIntAryWithNumStr(nStr1)
	err := ia.SetPrecision(precision, true)
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
