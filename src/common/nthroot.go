package common

import (
	"errors"
	"fmt"
	"math/big"
)

// nthroot.go
//
// Used to extract square roots and nth roots of positive and negative
// real numbers. Currently nth roots may only be passed as integer values.
//
// The technique employed to calculate nth roots is known as the
// "shifting nth-root algorithm".
//
// This source file is located in source code repository:
//
// https://github.com/MikeAustin71/mathhlpr.git
//
//
// See: https://en.wikipedia.org/wiki/Shifting_nth_root_algorithm
//
// Dependencies: intAry - intary.go
//
type NthRootOp struct {
	NthRoot            int
	OriginalNum        *IntAry
	BaseNumBundles     [][]int
	LenBaseNumBundles  int
	BaseNumBundlesIdx  int
	ResultAry          IntAry
	ResultIdx          int
	ResultPrecision    int
	RequestedPrecision int
	BigOne             *big.Int
	Big10              *big.Int
	Big10ToNthPower    *big.Int
	BigZero            *big.Int
	Big3               *big.Int
	Y                  *big.Int // Root Extracted thusfar
	YPrime             *big.Int // Next Value of Y
	Minuend            *big.Int
	Subtrahend         *big.Int
	R                  *big.Int // Let R be the remainder
	RPrime             *big.Int // Let RPrime be the new value of r for next iteration
	BaseNum            *big.Int // Base Number System - always 10
	Alpha              *big.Int // Next n-digits of the radicand
	Beta               *big.Int // Next Digit of the root
}

func (nthrt *NthRootOp) Empty() {
	nthrt.NthRoot = 0
	nthrt.OriginalNum = nil
	nthrt.BaseNumBundles = make([][]int, 0, 500)
	nthrt.LenBaseNumBundles = 0
	nthrt.BaseNumBundlesIdx = 0
	nthrt.ResultAry = IntAry{}.New()
	nthrt.ResultIdx = 0
	nthrt.ResultPrecision = 0
	nthrt.RequestedPrecision = 0
	nthrt.Big10 = big.NewInt(0)
	nthrt.Big10ToNthPower = big.NewInt(0)
}

// GetNthRootFloat32 - Calculates the Nth Root of a positive real number ('num')
// passed to the method as a type float32. In addition, the caller must supply
// input parameters for 'precision', 'nthRoot' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the float32 parameter, 'num', which
// will be input and positioned to the right of the decimal place.
//
// Nth root specifies the root which will be calculated for parameter, 'num'. Example,
// square root, cube root, 4th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetNthRootFloat32(num float32, precision, nthRoot, maxPrecision uint) (IntAry, error) {

	ai := IntAry{}.New()

	err := ai.SetIntAryWithFloat32(num, int(precision))

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.GetNthRootFloat32() - Error ai.SetIntAryWithFloat32(num, precision) num= %v, precision=%v Error= %v ", num, precision, err)
	}

	return nthrt.GetNthRootIntAry(&ai, nthRoot, maxPrecision)

}

// GetNthRootFloat64 - Calculates the Nth Root of a positive real number ('num')
// passed to the method as a type float64. In addition, the caller must supply
// input parameters for 'precision', 'nthRoot' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the float64 parameter, 'num', which
// will be input and positioned to the right of the decimal place.
//
// Nth root specifies the root which will be calculated for parameter, 'num'. Example,
// square root, cube root, 4th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetNthRootFloat64(num float64, precision, nthRoot, maxPrecision uint) (IntAry, error) {

	ai := IntAry{}.New()

	err := ai.SetIntAryWithFloat64(num, int(precision))

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.GetNthRootFloat64() - Error ai.SetIntAryWithFloat64(num, precision) num= %v, precision=%v Error= %v ", num, precision, err)
	}

	return nthrt.GetNthRootIntAry(&ai, nthRoot, maxPrecision)
}

// GetNthRootBigFloat - Calculates the Nth Root of a positive real number ('num')
// passed to the method as type Big Float (*big.Float). In addition, the caller
// must supply input parameters for the 'nthRoot' and 'maxPrecision'.
//
// Nth root specifies the root which will be calculated for parameter, 'num'. Example,
// square root, cube root, 4th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.

// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetNthRootBigFloat(num *big.Float, nthRoot, maxPrecision uint) (IntAry, error) {

	ai := IntAry{}.New()

	err := ai.SetIntAryWithFloatBig(num, -1)

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.SetIntAryWithFloatBig() - Error ai.SetIntAryWithFloatBig(num) num= %v Error= %v ", num, err)
	}

	return nthrt.GetNthRootIntAry(&ai, nthRoot, maxPrecision)
}

// GetNthRootInt - Calculates the Nth Root of a positive real number ('num')
// passed to the method as a type int. In addition, the caller must supply
// input parameters for 'precision', 'nthRoot' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the int64 parameter, 'num', which
// will be positioned to the right of the decimal place.
//
// Nth root specifies the root which will be calculated for parameter, 'num'. Example,
// square root, cube root, 4th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetNthRootInt(num int, precision, nthRoot, maxPrecision uint) (IntAry, error) {

	ai := IntAry{}.New()

	err := ai.SetIntAryWithInt(num, precision)

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.GetNthRootIn() - Error ai.SetIntAryWithInt32(num, precision) num= %v, precision=%v Error= %v ", num, precision, err)
	}

	return nthrt.GetNthRootIntAry(&ai, nthRoot, maxPrecision)
}

// GetNthRootInt64 - Calculates the Nth Root of a positive real number ('num')
// passed to the method as a type int64. In addition, the caller must supply
// input parameters for 'precision', 'nthRoot' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the int64 parameter, 'num', which
// will be positioned to the right of the decimal place.
//
// Nth root specifies the root which will be calculated for parameter, 'num'. Example,
// square root, cube root, 4th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetNthRootInt64(num int64, precision, nthRoot, maxPrecision uint) (IntAry, error) {

	ai := IntAry{}.New()

	err := ai.SetIntAryWithInt64(num, precision)

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.GetNthRootInt64() - Error ai.SetIntAryWithInt32(num, precision) num= %v, precision=%v Error= %v ", num, precision, err)
	}

	return nthrt.GetNthRootIntAry(&ai, nthRoot, maxPrecision)
}

// GetNthRootBigInt - Calculates the Nth Root of a positive real number ('num')
// passed to the method as type Big Int (*big.Int). In addition, the caller
// must supply input parameters for 'precision', 'nthRoot' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the *big.Int parameter, 'num', which
// will be positioned to the right of the decimal place.
//
// Nth root specifies the root which will be calculated for parameter, 'num'. Example,
// square root, cube root, 4th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetNthRootBigInt(num *big.Int, precision, nthRoot, maxPrecision uint) (IntAry, error) {
	ai := IntAry{}.New()

	err := ai.SetIntAryWithBigInt(num, precision)

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.GetNthRootBigInt() - Error ai.SetIntAryWithBigInt(num, precision) num= %v, precision=%v Error= %v ", num, precision, err)
	}

	return nthrt.GetNthRootIntAry(&ai, nthRoot, maxPrecision)
}

// GetNthRootIntAry  - Calculates the Nth Root of a real number ('num')
// passed to the method as a pointer to type intAry.  In addition, the caller must supply
// input parameters for 'nthRoot' and 'maxPrecision'.
//
// 'nthRoot' specifies the root which will be calculated for parameter, 'num'. Example,
// square root, cube root, 4th root, 9th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetNthRootIntAry(num *IntAry, nthRoot, maxPrecision uint) (IntAry, error) {

	err := nthrt.initializeAndExtract(num, nthRoot, maxPrecision)

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.GetNthRootIntAry() Error returned from initializeAndExtract. Error= %v", err)
	}

	return nthrt.ResultAry.CopyOut(), nil
}

// GetSquareRootFloat32 - Calculates the Square Root of a positive real number ('num')
// passed to the method as a type float32. In addition, the caller must supply
// input parameters for 'precision' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the float32 parameter, 'num', which
// will be input and positioned to the right of the decimal place.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootFloat32(num float32, precision, maxPrecision uint) (IntAry, error) {

	ai := IntAry{}.New()

	err := ai.SetIntAryWithFloat32(num, int(precision))

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.GetSquareRootFloat32() - Error ai.SetIntAryWithFloat32(num, precision) num= %v, precision=%v Error= %v ", num, precision, err)
	}

	return nthrt.GetSquareRootIntAry(&ai, maxPrecision)

}

// GetSquareRootFloat64 - Calculates the Square Root of a positive real number ('num')
// passed to the method as a type float64. In addition, the caller must supply
// input parameter 'precision' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the float64 parameter, 'num', which
// will be input and positioned to the right of the decimal place.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootFloat64(num float64, precision, maxPrecision uint) (IntAry, error) {

	ai := IntAry{}.New()

	err := ai.SetIntAryWithFloat64(num, int(precision))

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.GetSquareRootFloat64() - Error ai.SetIntAryWithFloat64(num, precision) num= %v, precision=%v Error= %v ", num, precision, err)
	}

	return nthrt.GetSquareRootIntAry(&ai, maxPrecision)
}

// GetSquareRootBigFloat - Calculates the Square Root of a positive real number ('num')
// passed to the method as type Big Float (*big.Float). In addition, the caller
// must supply input parameter for 'maxPrecision'.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.

// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootBigFloat(num *big.Float, maxPrecision uint) (IntAry, error) {

	ai := IntAry{}.New()

	err := ai.SetIntAryWithFloatBig(num, -1)

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.GetSquareRootBigFloat() - Error ai.SetIntAryWithFloatBig(num) num= %v Error= %v ", num, err)
	}

	return nthrt.GetSquareRootIntAry(&ai, maxPrecision)
}

// GetSquareRootInt - Calculates the Square Root of a positive real number ('num')
// passed to the method as type Int. In addition, the caller must supply input
// parameters for 'precision' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the 'int' parameter, 'num', which
// will be positioned to the right of the decimal place.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootInt(num int, precision, maxPrecision uint) (IntAry, error) {
	ai := IntAry{}.New()

	err := ai.SetIntAryWithInt(num, precision)

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.GetSquareRootInt() - Error ai.SetIntAryWithInt(num, precision) num= %v, precision=%v Error= %v ", num, precision, err)
	}

	return nthrt.GetSquareRootIntAry(&ai, maxPrecision)
}

// GetSquareRootInt32 - Calculates the Square Root of a positive real number ('num')
// passed to the method as type Int32. In addition, the caller must supply input
// parameters for 'precision' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the 'int32' parameter, 'num', which
// will be positioned to the right of the decimal place.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootInt32(num int32, precision, maxPrecision uint) (IntAry, error) {
	ai := IntAry{}.New()

	err := ai.SetIntAryWithInt32(num, precision)

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.GetSquareRootInt32() - Error ai.SetIntAryWithInt32(num, precision) num= %v, precision=%v Error= %v ", num, precision, err)
	}

	return nthrt.GetSquareRootIntAry(&ai, maxPrecision)
}

// GetSquareRootInt64 - Calculates the Square Root of a positive real number ('num')
// passed to the method as type Int64. In addition, the caller must supply input
// parameters for 'precision' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the 'int64' parameter, 'num', which
// will be positioned to the right of the decimal place.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootInt64(num int64, precision, maxPrecision uint) (IntAry, error) {
	ai := IntAry{}.New()

	err := ai.SetIntAryWithInt64(num, precision)

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.GetSquareRootInt64() - Error ai.SetIntAryWithInt64(num, precision) num= %v, precision=%v Error= %v ", num, precision, err)
	}

	return nthrt.GetSquareRootIntAry(&ai, maxPrecision)
}

// GetSquareRootBigInt - Calculates the Square Root of a positive real number ('num')
// passed to the method as a pointer to type Big Int (*big.Int). In addition, the caller
// must supply input parameters for 'precision' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the *big.Int parameter, 'num', which
// will be positioned to the right of the decimal place.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootBigInt(num *big.Int, precision, maxPrecision uint) (IntAry, error) {
	ai := IntAry{}.New()

	err := ai.SetIntAryWithBigInt(num, precision)

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.GetSquareRootBigInt() - Error ai.SetIntAryWithBigInt(num, precision) num= %v, precision=%v Error= %v ", num, precision, err)
	}

	return nthrt.GetSquareRootIntAry(&ai, maxPrecision)
}

// GetSquareRootIntAry - Calculates the square Root of a positive real number ('num')
// passed to the method as a pointer to type intAry. In addition, the caller
// must supply input parameters for 'maxPrecision'.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootIntAry(num *IntAry, maxPrecision uint) (IntAry, error) {

	err := nthrt.initializeAndExtract(num, 2, maxPrecision)

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("NthRootOp.GetNthRootIntAry() Error returned from initializeAndExtract. Error= %v", err)
	}

	return nthrt.ResultAry.CopyOut(), nil

}

// SetNthRootIntAry  - Calculates the Nth Root of a positive real number ('num')
// passed to the method as a pointer to type intAry.  In addition, the caller must supply
// input parameters for 'nthRoot' and 'maxPrecision'.
//
// The difference between this method, 'SetNthRootIntAry' and 'GetNthRootIntAry' is in
// the return value.  This method, 'SetNthRootIntAry' does not return the result. Instead,
// the calculation result is stored in the NthRootOp intAry Object, NthRootOp.ResultAry.
// This method is primarily for use by other low level routines seeking to improve performance
// by avoiding the return of a new intAry object.
//
// Nth root specifies the root which will be calculated for parameter, 'num'. Example,
// square root, cube root, 4th root, 9th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is stored in the NthRootOp field, 'NthRootOp.ResultAry'.
// 'NthRootOp.ResultAry' is an intAry Object.
//
// Note: A negative 'num' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) SetNthRootIntAry(num *IntAry, nthRoot, maxPrecision uint) error {

	err := nthrt.initializeAndExtract(num, nthRoot, maxPrecision)

	if err != nil {
		return fmt.Errorf("NthRootOp.SetNthRootIntAry() Error returned from initializeAndExtract. Error= %v", err)
	}

	return nil
}

func (nthrt *NthRootOp) initializeAndExtract(num *IntAry, nthRoot, maxPrecision uint) error {

	if nthRoot == 0 {
		nthrt.ResultAry.SetIntAryToOne(int(maxPrecision))
		return nil
	}

	err := nthrt.initialize(num, nthRoot, maxPrecision)

	if err != nil {
		return fmt.Errorf("NthRootOp.initializeAndExtract() Error returned from initialization. Error= %v", err)
	}

	err = nthrt.doRootExtraction()

	if err != nil {
		return fmt.Errorf("NthRootOp.initializeAndExtract() - Error returned from nthrt.doRootExtraction() - %v", err)
	}

	return nil

}

// initialize - Initializes the data fields of the NthRootOp structure and validates the
// original number passed to the Nth Root Calculation.
func (nthrt *NthRootOp) initialize(originalNum *IntAry, nthRoot, maxPrecision uint) error {

	err := originalNum.IsIntAryValid("NthRootOp.initialize() originalNum Invalid - ")

	if err != nil {
		return err
	}

	if nthRoot == 1 {
		return fmt.Errorf("NthRootOp.initialize() - Input Parameter 'nthRoot' INVALID! 'nthRoot' cannot equal 1. nthRoot= %v", nthRoot)
	}

	if originalNum.GetSign() == -1 {

		isNthRootEven := nthRoot / 2

		if isNthRootEven*2 == nthRoot {
			return fmt.Errorf("NthRootOp.initialize() - Cannot compute nthRoot of a negative number when nthRoot is even. nthRoot can only be extracted from negative numbers when nthRoot is odd. Original Number= %v  nthRoot= %v", originalNum.GetNumStr(), nthRoot)
		}

	}

	nthrt.NthRoot = int(nthRoot)
	nthrt.OriginalNum = originalNum
	nthrt.RequestedPrecision = int(maxPrecision)

	err = nthrt.bundleInts()

	if err != nil {
		return fmt.Errorf("NthRootOp.initialize() - Error returned from nthrt.bundleInts(). Error= %v", err)
	}

	err = nthrt.bundleFracs()

	if err != nil {
		return fmt.Errorf("NthRootOp.initialize() - Error returned from nthrt.bundleFracs(). Error= %v", err)
	}

	err = nthrt.calcPrecision()

	if err != nil {
		return fmt.Errorf("NthRootOp.initialize() - Error returned from nthrt.calcPrecision(). Error= %v", err)
	}

	// Set constants for calculations

	nthrt.BigOne = big.NewInt(1)
	nthrt.Big3 = big.NewInt(3)
	nthrt.Big10 = big.NewInt(10)
	nthrt.Big10ToNthPower = nthrt.bigPower(nthrt.Big10, uint(nthrt.NthRoot))
	nthrt.BigZero = big.NewInt(0)
	nthrt.ResultAry = IntAry{}.New()
	nthrt.ResultAry.SetPrecision(nthrt.RequestedPrecision, false)
	nthrt.ResultAry.SetSign(originalNum.GetSign())
	nthrt.Y = big.NewInt(0)
	nthrt.YPrime = big.NewInt(0)
	nthrt.R = big.NewInt(0)
	nthrt.RPrime = big.NewInt(0)
	nthrt.BaseNum = big.NewInt(10)
	nthrt.Alpha = big.NewInt(0)
	nthrt.Beta = big.NewInt(0)
	return nil
}

// bundleInts - computes the bundle size and creates the
// bundle array elements necessary to process the integer
// digits in the original number.
func (nthrt *NthRootOp) bundleInts() error {

	intNums, err := nthrt.OriginalNum.GetIntegerDigits()

	intNumsStats := intNums.GetIntAryStats()

	if err != nil {
		return fmt.Errorf("NthRootOp.bundleInts() - Error= %v", err)
	}

	if intNumsStats.IntAryLen < 1 {
		return errors.New("NthRootOp.bundleInts() intNums array length less than 1")
	}

	bundleSize := 0

	if intNumsStats.IntAryLen <= nthrt.NthRoot {

		bundleSize = 1

	} else {

		bundleSize = intNumsStats.IntAryLen / nthrt.NthRoot

		if intNumsStats.IntAryLen > ((intNumsStats.IntAryLen / nthrt.NthRoot) * nthrt.NthRoot) {

			bundleSize++

		}

	}

	nthrt.BaseNumBundles = make([][]int, bundleSize)

	bundleIdx := bundleSize - 1
	intAryIdx := 0
	intAry, _ := intNums.GetIntAry()
	for j := intNumsStats.IntAryLen - 1; j >= 0; j -= nthrt.NthRoot {
		bundle := make([]int, nthrt.NthRoot)
		intAryIdx = j
		for k := nthrt.NthRoot - 1; k >= 0; k-- {
			if intAryIdx < 0 {
				break
			}

			bundle[k] = int(intAry[intAryIdx])
			intAryIdx--
		}

		nthrt.BaseNumBundles[bundleIdx] = append(nthrt.BaseNumBundles[bundleIdx], bundle...)

		bundleIdx--
	}

	return nil
}

// bundleFracs - computes the numeric bundle size and creates
// the bundle array elements required to process the fractional
// digits in the original number.
func (nthrt *NthRootOp) bundleFracs() error {

	if nthrt.OriginalNum.GetPrecision() < 1 {
		return nil
	}

	fracNums, err := nthrt.OriginalNum.GetFractionalDigits()

	iFracNumStats := fracNums.GetIntAryStats()
	iFAry, _ := fracNums.GetIntAry()

	if err != nil {
		return fmt.Errorf("NthRootOp.bundleFracs() Error Returned from nthrt.OriginalNum.GetFractionalDigits() - Error= %v", err)
	}

	for i := 1; i < iFracNumStats.IntAryLen; i += nthrt.NthRoot {

		bundle := make([]int, nthrt.NthRoot)

		for j := 0; j < nthrt.NthRoot; j++ {

			if i+j < iFracNumStats.IntAryLen {
				bundle[j] = int(iFAry[i+j])
			}
		}

		nthrt.BaseNumBundles = append(nthrt.BaseNumBundles, bundle)

	}

	return nil
}

// calcPrecision - calculates the bundle size and creates
// the bundle array elements necessary to process the nth
// root to the requested number of decimal places to the
// right of the decimal point.
func (nthrt *NthRootOp) calcPrecision() error {

	if nthrt.OriginalNum.GetPrecision() < 0 {
		return fmt.Errorf("NthRootOp.calcPrecision() - Existing Precision is less than zero! OriginalNum.Precision= %v", nthrt.OriginalNum.GetPrecision())
	}

	existingPrecision := nthrt.OriginalNum.GetPrecision() / nthrt.NthRoot

	if nthrt.OriginalNum.GetPrecision() != existingPrecision*nthrt.NthRoot {
		existingPrecision++
	}

	bundle := make([]int, nthrt.NthRoot)

	if nthrt.RequestedPrecision <= existingPrecision {

		// nthrt.BaseNumBundles = append(nthrt.BaseNumBundles, bundle)

		nthrt.RequestedPrecision = existingPrecision

		//return nil
	}

	nthrt.RequestedPrecision++

	deltaPrecision := nthrt.RequestedPrecision - existingPrecision

	for i := 0; i < deltaPrecision; i++ {
		nthrt.BaseNumBundles = append(nthrt.BaseNumBundles, bundle)
	}

	nthrt.LenBaseNumBundles = len(nthrt.BaseNumBundles)

	nthrt.ResultPrecision = nthrt.RequestedPrecision

	return nil
}

func (nthrt *NthRootOp) doRootExtraction() error {

	nthrt.Y = big.NewInt(0)
	nthrt.Minuend = big.NewInt(0)
	nthrt.Subtrahend = big.NewInt(0)
	nthrt.R = big.NewInt(0)

	for i := 0; i < nthrt.LenBaseNumBundles; i++ {

		nthrt.findNextRoot(i)

	}

	nthrt.ResultAry.SetSign(nthrt.OriginalNum.GetSign())
	nthrt.ResultAry.OptimizeIntArrayLen(false)
	err := nthrt.ResultAry.RoundToPrecision(nthrt.RequestedPrecision - 1)

	if err != nil {
		return fmt.Errorf("doRootExtraction() - Error returned from sqrt.ResultAry.RoundToPrecision(sqrt.ResultPrecision). nthrt.ResultPrecision= %v  Error= %v", nthrt.ResultPrecision, err)
	}

	return nil
}

func (nthrt *NthRootOp) findNextRoot(bundleIdx int) {

	bundle := nthrt.getBundleBigInt(bundleIdx)

	// alpha = next n-digits of radicand
	nthrt.Alpha = big.NewInt(0).Set(bundle)

	nthrt.RPrime = big.NewInt(-1)

	// nthrt.R already set
	// n = nthrt.NthRoot already set
	// nthrt.Y already set
	// nthrt.BaseNum = 10
	// nthrt.Big10ToNthPower = BaseNum^n

	itatr := big.NewInt(9)
	term_1a := big.NewInt(0)
	term_1b := big.NewInt(0)

	term_2a1 := big.NewInt(0)
	term_2a2 := big.NewInt(0)
	term_2a := big.NewInt(0)

	term_2b := big.NewInt(0)
	term_2b1 := big.NewInt(0)
	term_2b2 := big.NewInt(0)

	term_1a = big.NewInt(0).Mul(nthrt.Big10ToNthPower, nthrt.R)
	term_1b = big.NewInt(0).Set(nthrt.Alpha)
	nthrt.Minuend = big.NewInt(0).Add(term_1a, term_1b)

	for itatr.Cmp(nthrt.BigZero) > -1 &&
		nthrt.RPrime.Cmp(nthrt.BigZero) == -1 {

		nthrt.Beta = big.NewInt(0).Set(itatr)
		nthrt.YPrime = big.NewInt(0).Mul(nthrt.Y, nthrt.Big10)
		nthrt.YPrime = big.NewInt(0).Add(nthrt.YPrime, nthrt.Beta)

		term_2a1 = big.NewInt(0).Mul(nthrt.BaseNum, nthrt.Y)
		term_2a2 = big.NewInt(0).Add(term_2a1, nthrt.Beta)
		term_2a = nthrt.bigPower(term_2a2, uint(nthrt.NthRoot))

		term_2b1 = big.NewInt(0).Set(nthrt.Big10ToNthPower)
		term_2b2 = nthrt.bigPower(nthrt.Y, uint(nthrt.NthRoot))
		term_2b = big.NewInt(0).Mul(term_2b1, term_2b2)

		nthrt.Subtrahend = big.NewInt(0).Sub(term_2a, term_2b)

		nthrt.RPrime = big.NewInt(0).Sub(nthrt.Minuend, nthrt.Subtrahend)

		itatr = big.NewInt(0).Sub(itatr, nthrt.BigOne)
	}

	nthrt.R = big.NewInt(0).Set(nthrt.RPrime)
	nthrt.Y = big.NewInt(0).Set(nthrt.YPrime)
	nthrt.ResultAry.AppendToIntAry(uint8(nthrt.Beta.Int64()))

}

func (nthrt *NthRootOp) bigPower(bigNum *big.Int, power uint) *big.Int {

	bigResult := big.NewInt(1)

	for i := uint(0); i < power; i++ {
		bigResult = big.NewInt(0).Mul(bigResult, bigNum)
	}

	return bigResult
}

func (nthrt *NthRootOp) getBundleBigInt(idx int) *big.Int {

	bigBundleVal := big.NewInt(0)

	for i := 0; i < nthrt.NthRoot; i++ {

		bigBundleVal = big.NewInt(0).Mul(bigBundleVal, nthrt.Big10)
		bigBundleVal = big.NewInt(0).Add(bigBundleVal, big.NewInt(0).SetInt64(int64(nthrt.BaseNumBundles[idx][i])))

	}

	return bigBundleVal
}
