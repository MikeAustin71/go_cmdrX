package common

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"bytes"
)

/*
	Source Code Repository
	**********************
	The source code file intary.go is located in the source code respository:

		https://github.com/MikeAustin71/intary.git

	IntAry Object
	*************
	Source file 'intary.go' contains a structure, 'IntAry', which is designed
 	to perform a variety of math operations on integer strings.

	The directory ./archives/eulersnumbercalc contains an example which calculates
	Euler's Number out to 1,000 digits.

	Dependencies
	************
	The 'IntAry' object has the following dependency:

	nthroot.go - Source Repository: https://github.com/MikeAustin71/mathhlpr.git

 */


// A fraction represented by a numerator and a denominator.
// Both numerator and denominator are of type intAry
type FracIntAry struct {
	Numerator   IntAry
	Denominator IntAry
}

// NewNumStrs - Creates a FracIntAry type by passing input parameters numerator and
// denominator as number strings.
func (fIa FracIntAry) NewNumStrs(numerator, denominator string) (FracIntAry, error) {

	var err error

	fIa2 := FracIntAry{}

	fIa2.Numerator, err = IntAry{}.NewNumStr(numerator)

	if err != nil {
		return FracIntAry{}, fmt.Errorf("FracIntAry.NewNumStrs() - Error returned from intAry{}.NewNumStr(numerator). Error= %v", err)
	}

	fIa2.Denominator, err = IntAry{}.NewNumStr(denominator)

	if err != nil {
		return FracIntAry{}, fmt.Errorf("FracIntAry.NewNumStrs() - Error returned from intAry{}.NewNumStr(denominator). Error= %v", err)
	}

	return fIa2, nil
}

// NewIntArys - Creates a type FracIntAry by passing numerator and denominator
// input parameters of type *intAry
func (fIa FracIntAry) NewIntArys(numerator, denominator *IntAry) FracIntAry {

	fIa2 := FracIntAry{}

	fIa2.Numerator = numerator.CopyOut()
	fIa2.Denominator = denominator.CopyOut()

	return fIa2
}

// GetRationalValue - Converts the fraction and returns the value as a
// big rational number (*big.Rat).
//
// Input parameter maxPrecision determines the maximum number of decimal
// places to the right of the decimal point contained in the result.
//
// If the value of maxPrecision is -1, maximum precision will default to
// 1024 decimal places. maxPrecision values less than -1 will trigger an
// error.
func (fIa *FracIntAry) GetRationalValue(maxPrecision int) (*big.Rat, error) {

	if maxPrecision < -1 {
		return big.NewRat(1, 1), fmt.Errorf("GetRationalValue() - maxPrecision is less than -1 and therefore INVALID. maxPrecision= %v", maxPrecision)
	}

	if maxPrecision == -1 {
		maxPrecision = 1024
	}

	if fIa.Numerator.GetPrecision() == 0 && fIa.Denominator.GetPrecision() == 0 {
		fRat, ok := big.NewRat(1, 1).SetString(fIa.Numerator.GetNumStr() + "/" + fIa.Denominator.GetNumStr())

		if !ok {
			return big.NewRat(1, 1), errors.New("GetRationalValue() - big.NewFloat(0).SetString(fracIa.GetNumStr()) Failed!")
		}

		return fRat, nil

	}

	newFloat, err := fIa.Numerator.DivideThisBy(&fIa.Denominator, maxPrecision)

	if err != nil {
		return big.NewRat(1, 1), fmt.Errorf("GetRationalValue() - Error returned from fIa.Numerator.DivideThisBy(&fIa.Denominator, 42). Error= %v", err)
	}

	fRat, ok := big.NewRat(1, 1).SetString(newFloat.GetNumStr())

	if !ok {
		return big.NewRat(1, 1), errors.New("GetRationalValue() - big.NewFloat(0).SetString(fracIa.GetNumStr()) Failed!")
	}

	return fRat, nil

}

type BackUpIntAry struct {
	intAry                 []uint8
	intAryLen              int
	integerLen             int
	significantIntegerLen  int
	significantFractionLen int
	firstDigitIdx          int
	lastDigitIdx           int
	isZeroValue            bool
	isIntegerZeroValue     bool
	precision              int
	signVal                int
	decimalSeparator       rune // https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
	thousandsSeparator     rune // https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
	currencySymbol         rune // See Currency Symbol References below:
															// https://gist.github.com/bzerangue/5484121
															// http://symbologic.info/currency.htm
															// http://www.xe.com/symbols.php
}

func (iBa BackUpIntAry) New() BackUpIntAry {
	iAry := BackUpIntAry{}

	iAry.intAry = []uint8{}
	iAry.intAryLen = 0
	iAry.integerLen = 0
	iAry.significantIntegerLen = 0
	iAry.significantFractionLen = 0
	iAry.firstDigitIdx = -1
	iAry.lastDigitIdx = -1
	iAry.isZeroValue = true
	iAry.isIntegerZeroValue = true
	iAry.precision = 0
	iAry.signVal = 1
	iAry.decimalSeparator = '.'
	iAry.thousandsSeparator = ','
	iAry.currencySymbol = '$'

	return iAry
}

func (iBa *BackUpIntAry) Empty() {
	iBa.intAry = []uint8{}
	iBa.intAryLen = 0
	iBa.integerLen = 0
	iBa.significantIntegerLen = 0
	iBa.significantFractionLen = 0
	iBa.firstDigitIdx = -1
	iBa.lastDigitIdx = -1
	iBa.isZeroValue = true
	iBa.isIntegerZeroValue = true
	iBa.precision = 0
	iBa.signVal = 1
	if iBa.decimalSeparator == 0 {
		iBa.decimalSeparator = '.'
	}

	if iBa.thousandsSeparator == 0 {
		iBa.thousandsSeparator = ','
	}

	if iBa.currencySymbol == 0 {
		iBa.currencySymbol = '$'
	}
}

func (iBa *BackUpIntAry) CopyIn(iBa2 *BackUpIntAry) {
	iBa2.SetInternalFlags()
	iBa.Empty()

	iBa.intAry = make([]uint8, iBa2.intAryLen)
	for i := 0; i < iBa2.intAryLen; i++ {
		iBa.intAry[i] = iBa2.intAry[i]
	}

	iBa.intAryLen = iBa2.intAryLen
	iBa.integerLen = iBa2.integerLen
	iBa.significantIntegerLen = iBa2.significantIntegerLen
	iBa.significantFractionLen = iBa2.significantFractionLen
	iBa.firstDigitIdx = iBa2.firstDigitIdx
	iBa.lastDigitIdx = iBa2.lastDigitIdx
	iBa.isZeroValue = iBa2.isZeroValue
	iBa.isIntegerZeroValue = iBa2.isIntegerZeroValue
	iBa.precision = iBa2.precision
	iBa.signVal = iBa2.signVal
	iBa.decimalSeparator = iBa2.decimalSeparator
	iBa.thousandsSeparator = iBa2.thousandsSeparator
	iBa.currencySymbol = iBa2.currencySymbol

}

func (iBa *BackUpIntAry) CopyOut() BackUpIntAry {
	iBa.SetInternalFlags()
	iAry2 := BackUpIntAry{}.New()

	iAry2.intAry = make([]uint8, iBa.intAryLen)

	for i := 0; i < iBa.intAryLen; i++ {
		iAry2.intAry[i] = iBa.intAry[i]
	}

	iAry2.intAryLen = iBa.intAryLen
	iAry2.integerLen = iBa.integerLen
	iAry2.significantIntegerLen = iBa.significantIntegerLen
	iAry2.significantFractionLen = iBa.significantFractionLen
	iAry2.firstDigitIdx = iBa.firstDigitIdx
	iAry2.lastDigitIdx = iBa.lastDigitIdx
	iAry2.isZeroValue = iBa.isZeroValue
	iAry2.isIntegerZeroValue = iBa.isIntegerZeroValue
	iAry2.precision = iBa.precision
	iAry2.signVal = iBa.signVal
	iAry2.decimalSeparator = iBa.decimalSeparator
	iAry2.thousandsSeparator = iBa.thousandsSeparator
	iAry2.currencySymbol = iBa.currencySymbol

	return iAry2
}

func (iBa *BackUpIntAry) Equals(iBa2 *BackUpIntAry) bool {
	iBa.SetInternalFlags()
	iBa2.SetInternalFlags()

	if iBa.intAryLen != iBa2.intAryLen {
		return false
	}

	for i := 0; i < iBa2.intAryLen; i++ {
		if iBa.intAry[i] != iBa2.intAry[i] {
			return false
		}
	}

	if iBa.integerLen != iBa2.integerLen ||
		iBa.significantIntegerLen != iBa2.significantIntegerLen ||
		iBa.significantFractionLen != iBa2.significantFractionLen ||
		iBa.firstDigitIdx != iBa2.firstDigitIdx ||
		iBa.lastDigitIdx != iBa2.lastDigitIdx ||
		iBa.isZeroValue != iBa2.isZeroValue ||
		iBa.isIntegerZeroValue != iBa2.isIntegerZeroValue ||
		iBa.precision != iBa2.precision ||
		iBa.signVal != iBa2.signVal ||
		iBa.decimalSeparator != iBa2.decimalSeparator ||
		iBa.thousandsSeparator != iBa2.thousandsSeparator ||
		iBa.currencySymbol != iBa2.currencySymbol	{

		return false
	}

	return true
}

func (iBa *BackUpIntAry) GetIntAryStats() IntAryStatsDto {

	iBa.SetInternalFlags()

	iStats := IntAryStatsDto{}

	iStats.IntAryLen = iBa.intAryLen
	iStats.IntegerLen = iBa.integerLen
	iStats.SignificantIntegerLen = iBa.significantIntegerLen
	iStats.SignificantFractionLen = iBa.significantFractionLen
	iStats.Precision = iBa.precision
	iStats.SignVal	= iBa.signVal
	iStats.FirstDigitIdx = iBa.firstDigitIdx
	iStats.LastDigitIdx = iBa.lastDigitIdx
	iStats.IsZeroValue = iBa.isZeroValue
	iStats.IsIntegerZeroValue = iBa.isIntegerZeroValue
	iStats.DecimalSeparator = iBa.decimalSeparator
	iStats.ThousandsSeparator = iBa.thousandsSeparator
	iStats.CurrencySymbol = iBa.currencySymbol

	return iStats

}

func (iBa *BackUpIntAry) GetNumStr() string {

	iBa.SetInternalFlags()

	if iBa.decimalSeparator == 0 {
		iBa.decimalSeparator = '.'
	}

	var buffer bytes.Buffer

	if iBa.signVal < 0 {
		buffer.WriteRune('-')
	}

	intLen := iBa.intAryLen - iBa.precision

	for i:=0 ; i < intLen; i++ {
		buffer.WriteRune(rune(iBa.intAry[i] + 48))
	}

	if iBa.precision > 0 {
		buffer.WriteRune(iBa.decimalSeparator)

		for j:= 0; j < iBa.precision; j++ {
			buffer.WriteRune(rune(iBa.intAry[intLen] + 48))
			intLen++
		}

	}

	return buffer.String()

}

func (iBa *BackUpIntAry) GetPrecision() int {
	return iBa.precision
}


func (iBa *BackUpIntAry) GetSignValue() int {
	return iBa.signVal
}



func (iBa *BackUpIntAry) SetInternalFlags() {

	iBa.intAryLen = len(iBa.intAry)

	if iBa.intAryLen == iBa.precision {
		iBa.intAry = append([]uint8{0}, iBa.intAry...)
		iBa.intAryLen++
	}

	if iBa.intAryLen < iBa.precision {

		deltaZeros := iBa.precision - iBa.intAryLen + 1
		zeroAry := make([]uint8, deltaZeros)
		iBa.intAry = append(zeroAry, iBa.intAry...)
		iBa.intAryLen += deltaZeros
	}

	iBa.firstDigitIdx = -1
	iBa.lastDigitIdx = -1

	lastIntIdx := iBa.intAryLen - iBa.precision - 1
	iBa.isZeroValue = true
	iBa.isIntegerZeroValue = true
	iBa.integerLen = iBa.intAryLen - iBa.precision

	for i := 0; i < iBa.intAryLen; i++ {
		if iBa.intAry[i] > 0 {
			iBa.isZeroValue = false

			if i < iBa.integerLen {
				iBa.isIntegerZeroValue = false
			}
		}
		// At minimum, there should be a single
		// leading zero before the decimal point.
		// Example 0.000.
		if i == lastIntIdx && iBa.intAry[i] == 0 {

			if iBa.firstDigitIdx == -1 {
				iBa.firstDigitIdx = i
			}

		}

		if iBa.intAry[i] > 0 {

			if iBa.firstDigitIdx == -1 {
				iBa.firstDigitIdx = i
			}

			iBa.lastDigitIdx = i
		}

	}

	iBa.significantIntegerLen = iBa.intAryLen - iBa.precision - iBa.firstDigitIdx

	if iBa.lastDigitIdx >= iBa.integerLen {
		iBa.significantFractionLen = iBa.precision - (iBa.lastDigitIdx - iBa.integerLen + 1)
	} else {
		iBa.significantFractionLen = 0
	}

}

func (iBa *BackUpIntAry) SetSignValue(signVal int) {

	if signVal < 0 {
		panic("BackUpIntAry.SetSignValue() - sign value less than zero!")
	}

	iBa.signVal = signVal

}

// IntAryStatsDto - used to transmit
// statistics and critical configuration
// data on an IntAry object.
type IntAryStatsDto struct {
	IntAryLen              int
	IntegerLen             int
	SignificantIntegerLen  int
	SignificantFractionLen int
	Precision              int
	SignVal                int
	FirstDigitIdx          int
	LastDigitIdx           int
	IsZeroValue            bool
	IsIntegerZeroValue     bool
	DecimalSeparator       rune
	ThousandsSeparator     rune
	CurrencySymbol         rune
}

// IntAry - Used to perform string
// based arithmetic.
//
// Dependencies: NthRootOp - nthroot.go
//
type IntAry struct {
	intAry                 []uint8
	intAryLen              int
	integerLen             int
	significantIntegerLen  int
	significantFractionLen int
	firstDigitIdx          int
	lastDigitIdx           int
	isZeroValue            bool
	isIntegerZeroValue     bool
	precision              int
	signVal                int
	decimalSeparator       rune
	thousandsSeparator     rune
	currencySymbol         rune
	BackUp                 BackUpIntAry
}

// AddToThis - Adds the value of intAry parameter ia2 to the value
// of the current intAry object.
//
// Paramters:
//
// ia2 *intAry - Incoming intAry object whose value will be subtracted
// 								from this current intAry value.
func (ia *IntAry) AddToThis(ia2 *IntAry) error {

	ia.SetEqualArrayLengths(ia2)

	if ia2.isZeroValue {
		return nil
	}

	compare := ia.CompareAbsoluteValues(ia2)

	newSignVal := ia.signVal
	doAdd := true
	isZeroResult := false
	doReverseNums := false

	if compare == 1 {
		// compare == + 1
		// Absolute Value: N1 > N2

		if ia.signVal == 1 && ia2.signVal == 1 {
			doAdd = true
			newSignVal = 1
		} else if ia.signVal == -1 && ia2.signVal == 1 {
			doAdd = false
			newSignVal = -1
		} else if ia.signVal == -1 && ia2.signVal == -1 {
			doAdd = true
			newSignVal = -1
		} else {
			// Must Be ia.signVal == 1 && ia2.signVal == -1
			doAdd = false
			newSignVal = 1
		}

	} else if compare == -1 {
		// Absolute Values: N2 > N1
		if ia.signVal == 1 && ia2.signVal == 1 {
			doAdd = true
			newSignVal = 1
		} else if ia.signVal == -1 && ia2.signVal == 1 {
			doAdd = false
			doReverseNums = true
			newSignVal = 1
		} else if ia.signVal == -1 && ia2.signVal == -1 {
			doAdd = true
			newSignVal = -1
		} else {
			// Must Be ia.signVal == 1 && ia2.signVal == -1
			doAdd = false
			doReverseNums = true
			newSignVal = -1
		}

	} else {
		// Must be compare == 0
		// Absolute Values: N1==N2
		if ia.signVal == 1 && ia2.signVal == 1 {
			doAdd = true
			newSignVal = 1
		} else if ia.signVal == -1 && ia2.signVal == 1 {
			doAdd = false
			newSignVal = 1
			isZeroResult = true
		} else if ia.signVal == -1 && ia2.signVal == -1 {
			doAdd = true
			newSignVal = -1
		} else {
			// Must Be ia.signVal == 1 && ia2.signVal == -1
			doAdd = false
			newSignVal = 1
			isZeroResult = true
		}

	}

	return ia.addToSubtractFromThis(ia2, newSignVal, doAdd, isZeroResult, doReverseNums)
}

// AddIntToThis - Adds an integer number to the value of the
// current IntAry object.
//
//
// Input Parameters:
//
//	num int 				-	The integer number to be added to the current IntAry object.
//
//  precision uint 	- The precision which should be applied to the int64 input
//										parameter to designate the number of digits to the right
//										of the decimal point. Example:  num = 123456, precision = 3
//										Result = 123.456 will be added to the current value of the
//                    the current IntAry object.
//
// Example:
//  intDigits     precision     	    result
//  946254  			   3							   946.254
//  946254				   0							   946254
//  -946254  			   3					      -946.254
//  -946254				   0						    -946254
//
func (ia *IntAry) AddIntToThis(num int, precision uint) error {

	ia2, err := IntAry{}.NewInt(num, precision)

	if err != nil {
		return err
	}

	ia.AddToThis(&ia2)

	return nil
}

// AddInt64ToThis - Adds an integer (int64) to the value of the
// current IntAry object.
//
// Input Parameters:
//
//	num int64 -	The integer number to be added to the current IntAry object.
//
//  precision uint - 	The precision which should be applied to the int64 input
//										parameter to designate the number of digits to the right
//										of the decimal point. Example:  num = 123456, precision = 3
//										Result = 123.456 will be added to the current value of the
//                    the current IntAry object.
// Example:
//  intDigits     precision     	    result
//  946254  			   3							   946.254
//  946254				   0							   946254
//  -946254  			   3					      -946.254
//  -946254				   0						    -946254
//
func (ia *IntAry) AddInt64ToThis(num int64, precision uint) error {

	ia2, err := IntAry{}.NewInt64(num, precision)

	if err != nil {
		return err
	}

	ia.AddToThis(&ia2)

	return nil
}

// AddBigIntToThis - Adds the value of the *big.Int input
// parameter, 'num' to the value of the current IntAry object.
//
// Input Parameters:
//
// num *big.Int 	- The numeric value to be added to the
//									value of the current IntAry object. Note
//									that 'num' may be positive or negative.
//
// precision uint	- 'precision' indicates the number of digits
// 									to be formatted to the right of the decimal
// 									place.
//
// Example:
//  intDigits     precision     	    result
//  946254  			   3							   946.254
//  946254				   0							   946254
//  -946254  			   3					      -946.254
//  -946254				   0						    -946254
//
// Usage:
// num := big.NewInt(123456)
// precision := uint(3)
// err := ia.AddBigIntToThis(num, precision)
//
// The result will equal the current value of the IntAry
// object plus, '123.456'
//
func (ia *IntAry) AddBigIntToThis(num *big.Int, precision uint) error {

	ia2, err := IntAry{}.NewBigInt(num, precision)

	if err != nil {
		return err
	}

	ia.AddToThis(&ia2)

	return nil
}
// AddFloat32ToThis - Adds the value of input number (float32)
// to the current value of this IntAry object.
//
// Intput Parameters:
//
// num float32 		- The number which will be added to the current
//									IntAry object.
//
// precision int	-	Input parameter 'precision' is used
//									to set the input precision of 'num'.
//
// 									Input parameter 'precision' must be set
// 									to a number greater than or equal to zero.
// 									It may also be set to a value of -1
// 									which causes the number to be formatted to
// 									the smallest number of digits to right of
// 									the decimal point.
//
// Usage:
// num := float32(123.456000)
// precision := 3 - signals that only the decimals 456 will included as input
// err := ia.AddFloat32ToThis(num, precision)
//
// This usage will result in a value of '123.456'
//
// If precision were set to -1, the resulting value
// would also be '123.456'
//
// Note: if precision is a positive value, and less than the current number
// of digits to the right of the decimal place, rounding may occur.
//
func (ia *IntAry) AddFloat32ToThis(num float32, precision int) error {

	if precision < -1 {
		return fmt.Errorf("AddFloat32ToThis() Error: Input parameter 'precision' is invalid. 'precision' must be greater than or equal to -1. 'precision'= '%v'", precision)
	}

	ia2, err := IntAry{}.NewFloat32(num, precision)

	if err != nil {
		return nil
	}

	ia.AddToThis(&ia2)

	return nil
}

// AddFloat64ToThis - Adds a floating point number (float64) to the
// current value of this IntAry object.
//
// Input Parameters:
//
//	num float64 - This float64 value will be added to the value of the
//								current IntAry. Positive and negative values are
//								accepted.
//
// precision int -	'precision' is applied to the input parameter
//									'num' (float64) to determine the number of
//									digits to the right of the decimal point which
//									will be applied to the resulting value.
//
//									Input parameter 'precision' must be set to a
// 									number greater than or equal to zero.  It may
// 									also be set to a value of -1 which causes the
// 									number to be formatted to the smallest number
// 									of digits to right of the decimal point.
//
//									Note: if precision is a positive number and less
// 									than the current number of digits to the right
// 									of the decimal place, rounding may occur.
//
func (ia *IntAry) AddFloat64ToThis(num float64, precision int) error {

	if precision < -1 {
		return fmt.Errorf("AddFloat64ToThis() Error: Input parameter 'precision' is invalid. 'precision' must be greater than or equal to -1. 'precision'= '%v'", precision)
	}

	ia2, err := IntAry{}.NewFloat64(num, precision)

	if err != nil {
		return err
	}

	ia.AddToThis(&ia2)

	return nil
}

// AddFloatBigToThis - Adds a *big.Float number to the current
// value of this IntAry object.
//
// Input Parameters:
//
//	num *big.Float	- This *big.Float value will be added to the value
// 										of the current IntAry object. Positive and negative
// 										values are accepted.
//
// precision int -	'precision' is applied to the input parameter
//									'num' (*big.Float) to determine the number of
//									digits to the right of the decimal point which
//									will be applied to the resulting value.
//
//									Input parameter 'precision' must be set to a
// 									number greater than or equal to zero.  It may
// 									also be set to a value of -1 which causes the
// 									number to be formatted to the smallest number
// 									of digits to right of the decimal point.
//
//									Note: if precision is a positive number and less
// 									than the current number of digits to the right
// 									of the decimal place, rounding may occur.
//
func (ia *IntAry) AddFloatBigToThis(num *big.Float, precision int) error {

	if precision < -1 {
		return fmt.Errorf("AddFloatBigToThis() Error: Input parameter 'precision' is invalid. 'precision' must be greater than or equal to -1. 'precision'= '%v'", precision)
	}


	ia2, err := IntAry{}.NewFloatBig(num, precision)

	if err != nil {
		return err
	}

	ia.AddToThis(&ia2)

	return nil
}

// AddMultipleToThis - Add the values of  multiple intAry objects to the current
// intAry value.
//
// convertToNumStr - boolean value determines whether the current intAry
//                   object will convert the intAry value to a number string.
//                   Set this parameter to 'false' if this method is called
//                   multiple times in order to improve performance.
//
func (ia *IntAry) AddMultipleToThis(iaMany ...*IntAry) error {

	var err error

	for _, iAry := range iaMany {

		err = ia.AddToThis(iAry)

		if err != nil {
			return fmt.Errorf("AddMultipleToThis() - Received error from ia.AddToThis(iAry, false). Error= %v", err)
		}

	}

	return nil
}

func (ia *IntAry) addToSubtractFromThis(ia2 *IntAry, newSignVal int, doAdd bool, isZeroResult bool, doReverseNums bool) error {

	if isZeroResult {
		ia.SetIntAryToZero(ia.precision)
		return nil
	}

	ia.signVal = newSignVal

	carry := 0
	n1 := 0
	n2 := 0
	n3 := 0

	for j := ia.intAryLen - 1; j >= 0; j-- {

		if doReverseNums {

			n2 = int(ia.intAry[j])
			n1 = int(ia2.intAry[j])

		} else {
			n1 = int(ia.intAry[j])
			n2 = int(ia2.intAry[j])

		}

		if doAdd {
			// doAdd == true
			// Do Addition

			n3 = n1 + n2 + carry

			if n3 > 9 {
				n3 = n1 + n2 + carry - 10
				carry = 1

			} else {
				carry = 0
			}

		} else {
			// doAdd == false
			// Do Subtraction
			n3 = n1 - n2 - carry

			if n3 < 0 {
				n3 = n1 + 10 - n2 - carry
				carry = 1
			} else {
				carry = 0
			}
		}

		ia.intAry[j] = uint8(n3)

	}

	if carry > 0 {
		ia.intAry = append([]uint8{1}, ia.intAry...)
		ia.intAryLen++
	}

	if ia.intAry[0] == 0 {
		ia.SetSignificantDigitIdxs()
		ia.intAry = ia.intAry[ia.firstDigitIdx:]
	}

	ia.SetInternalFlags()

	return nil
}

// AddArrayLengthLeft - Adds leading zeros to the intAry
func (ia *IntAry) AddArrayLengthLeft(addLen int) {

	ia.SetIntAryLength()

	newLen := addLen + ia.intAryLen
	t := make([]uint8, newLen)

	for i := 0; i < newLen; i++ {

		if i < addLen {
			t[i] = 0
		} else {
			t[i] = ia.intAry[i-addLen]
		}

	}

	ia.intAry = t
	ia.SetIntAryLength()
}

// AddArrayLengthRight - Adds trailing zeros
// to the right of the current intAry.
func (ia *IntAry) AddArrayLengthRight(addLen int) {
	ia.SetIntAryLength()

	for i := 0; i < addLen; i++ {
		ia.intAry = append(ia.intAry, 0)
	}

	ia.SetIntAryLength()
}


// AppendToIntAry - appends an integer of
// type uint8 to the internal Integer Array
// of the current IntAry object.
func (ia *IntAry) AppendToIntAry(num uint8) {
	ia.intAry = append(ia.intAry, num)
	ia.intAryLen  = len(ia.intAry)
	ia.integerLen = ia.intAryLen - ia.precision
	if num > 0 {
		ia.isZeroValue = false
	}
}

func (ia *IntAry) Ceiling() (IntAry, error) {

	err := ia.IsIntAryValid("Ceiling() - ")

	if err != nil {
		return IntAry{}, err
	}

	iAry2 := IntAry{}.New()

	intLen := ia.intAryLen - ia.precision

	intIdx := intLen - 1

	hasFracDigits, err := ia.HasFractionalDigits()

	if err != nil {
		return iAry2, err
	}

	if !hasFracDigits {
		iAry2 = ia.CopyOut()
		return iAry2, nil
	}

	if ia.signVal < 0 && hasFracDigits {

		t := make([]uint8, ia.intAryLen)

		for i := 0; i < intLen; i++ {
			t[i] = ia.intAry[i]
		}

		iAry2.intAry = t[0:]
		iAry2.intAryLen = ia.intAryLen
		iAry2.precision = ia.precision
		iAry2.signVal = ia.signVal
		return iAry2, nil
	}

	t := make([]uint8, ia.intAryLen+1)

	n1 := 0
	n2 := 0
	carry := 0
	adjFac := 1 * ia.signVal
	for i := intIdx; i >= 0; i-- {

		n1 = int(ia.intAry[i])

		if i == intIdx {
			if n1+adjFac < 0 {
				n2 = 10 + n1 + adjFac
				carry = -1
			} else if n1+adjFac > 9 {
				n2 = n1 + adjFac - 10
				carry = 1
			} else {
				n2 = n1 + adjFac
				carry = 0
			}

		} else {

			if int(n1) + carry < 0 {
				n2 = 10 + n1
				carry = -1
			} else if int(n1) + carry > 9 {
				n2 = n1 - 10
				carry = 1
			} else {
				n2 = n1 + carry
				carry = 0
			}

		}

		t[i+1] = uint8(n2)

	}

	if carry != 0 {
		t[0] = uint8(carry)
		iAry2.intAry = t[0 : ia.intAryLen+1]

	} else {

		iAry2.intAry = t[1 : ia.intAryLen+1]
	}

	iAry2.intAryLen = len(iAry2.intAry)
	iAry2.precision = ia.precision
	iAry2.signVal = ia.signVal
	return iAry2, nil
}

func (ia *IntAry) CompareSignedValues(iAry2 *IntAry) int {

	iCompare := ia.CompareAbsoluteValues(iAry2)

	if ia.isZeroValue && iAry2.isZeroValue {
		return 0
	}

	if ia.signVal != iAry2.signVal {

		if ia.signVal == 1 {
			return 1
		} else {
			return -1
		}
	}

	// Must be ia.signVal == iAry2.signVal

	if ia.signVal == 1 {
		return iCompare
	}

	// Must be ia.signVal && iAry2.signVal == -1

	return iCompare * -1

}

func (ia *IntAry) CompareAbsoluteValues(iAry2 *IntAry) int {

	ia.SetIntAryLength()
	iAry2.SetIntAryLength()

	iAry2.SetIsZeroValue()
	ia.SetIsZeroValue()

	if ia.isZeroValue && iAry2.isZeroValue {
		return 0
	}

	iaIntLen := ia.intAryLen - ia.precision
	iAry2IntLen := iAry2.intAryLen - iAry2.precision

	// Integer Lengths are Equal
	if iaIntLen == iAry2IntLen {
		for i := 0; i < iaIntLen; i++ {
			if ia.intAry[i] > iAry2.intAry[i] {
				return 1
			}

			if iAry2.intAry[i] > ia.intAry[i] {
				return -1
			}
		}
	}

	deltaStartIdx := 0

	// ia Integer Length is Greater than IAry2 Integer Length
	if iaIntLen > iAry2IntLen {
		deltaStartIdx = iaIntLen - iAry2IntLen

		for j := 0; j < iaIntLen; j++ {

			if j < deltaStartIdx {

				if ia.intAry[j] > 0 {
					return 1
				}

			} else {
				// i must be >= deltaStartIdx

				if ia.intAry[j] > iAry2.intAry[j-deltaStartIdx] {
					return 1
				}

				if iAry2.intAry[j-deltaStartIdx] > ia.intAry[j] {
					return -1
				}
			}
		}
	}

	// iAry2 Integer Length is Greater Than ia Integer Length
	if iAry2IntLen > iaIntLen {
		deltaStartIdx = iAry2IntLen - iaIntLen

		for k := 0; k < iAry2IntLen; k++ {

			if k < deltaStartIdx {
				if iAry2.intAry[k] > 0 {
					return -1
				}

			} else {
				// i must be >= deltaStartIdx

				if iAry2.intAry[k] > ia.intAry[k-deltaStartIdx] {
					return -1
				}

				if ia.intAry[k-deltaStartIdx] > iAry2.intAry[k] {
					return 1
				}
			}
		}
	}

	// If precision is zero, the intAry's are equivalent
	if ia.precision == 0 && iAry2.precision == 0 {
		return 0
	}

	// Integer Values are Equivalent. Now test
	// digits to the right of the decimal point.

	// Test fractional digits to right of decimal point
	iaFracIdx := iaIntLen
	iAry2FracIdx := iAry2IntLen
	// Test for case of Equal precision
	if ia.precision == iAry2.precision {
		for m := 0; m < ia.precision; m++ {

			if ia.intAry[iaFracIdx] > iAry2.intAry[iAry2FracIdx] {
				return 1
			}

			if iAry2.intAry[iAry2FracIdx] > ia.intAry[iaFracIdx] {
				return -1
			}

			iaFracIdx++
			iAry2FracIdx++
		}
	}

	iaFracIdx = iaIntLen
	iAry2FracIdx = iAry2IntLen
	// Test for case where ia precision Greater than iAry2 precision
	if ia.precision > iAry2.precision {

		for i := 0; i < ia.precision; i++ {

			if i < iAry2.precision {

				if ia.intAry[iaFracIdx] > iAry2.intAry[iAry2FracIdx] {
					return 1
				}

				if iAry2.intAry[iAry2FracIdx] > ia.intAry[iaFracIdx] {
					return -1
				}

				iaFracIdx++
				iAry2FracIdx++

			} else {
				if ia.intAry[iaFracIdx] > 0 {
					return 1
				}

				iaFracIdx++
			}
		}
	}

	iaFracIdx = iaIntLen
	iAry2FracIdx = iAry2IntLen
	// Test for case where iAry2 precision Greater than ia precision
	if iAry2.precision > ia.precision {

		for i := 0; i < iAry2.precision; i++ {

			if i < ia.precision {

				if ia.intAry[iaFracIdx] > iAry2.intAry[iAry2FracIdx] {
					return 1
				}

				if iAry2.intAry[iAry2FracIdx] > ia.intAry[iaFracIdx] {
					return -1
				}

				iaFracIdx++
				iAry2FracIdx++

			} else {
				if iAry2.intAry[iAry2FracIdx] > 0 {
					return -1
				}

				iAry2FracIdx++
			}
		}

	}

	// The two absolute numeric values must be equal
	return 0

}


func (ia *IntAry) CopyIn(iAry2 *IntAry, copyBackUp bool) {
	iAry2.SetInternalFlags()
	ia.Empty()

	ia.intAry = make([]uint8, iAry2.intAryLen)
	for i := 0; i < iAry2.intAryLen; i++ {
		ia.intAry[i] = iAry2.intAry[i]
	}

	ia.intAryLen = iAry2.intAryLen
	ia.integerLen = iAry2.integerLen
	ia.significantIntegerLen = iAry2.significantIntegerLen
	ia.significantFractionLen = iAry2.significantFractionLen
	ia.firstDigitIdx = iAry2.firstDigitIdx
	ia.lastDigitIdx = iAry2.lastDigitIdx
	ia.isZeroValue = iAry2.isZeroValue
	ia.isIntegerZeroValue = iAry2.isIntegerZeroValue
	ia.precision = iAry2.precision
	ia.signVal = iAry2.signVal
	ia.decimalSeparator = iAry2.decimalSeparator
	ia.thousandsSeparator = iAry2.thousandsSeparator
	ia.currencySymbol = iAry2.currencySymbol

	if copyBackUp {
		ia.BackUp = iAry2.BackUp.CopyOut()
	}
}

func (ia *IntAry) CopyOut() IntAry {
	ia.SetInternalFlags()
	iAry2 := IntAry{}.New()

	iAry2.intAry = make([]uint8, ia.intAryLen)

	for i := 0; i < ia.intAryLen; i++ {
		iAry2.intAry[i] = ia.intAry[i]
	}

	iAry2.intAryLen = ia.intAryLen
	iAry2.integerLen = ia.integerLen
	iAry2.significantIntegerLen = ia.significantIntegerLen
	iAry2.significantFractionLen = ia.significantFractionLen
	iAry2.firstDigitIdx = ia.firstDigitIdx
	iAry2.lastDigitIdx = ia.lastDigitIdx
	iAry2.isZeroValue = ia.isZeroValue
	iAry2.isIntegerZeroValue = ia.isIntegerZeroValue
	iAry2.precision = ia.precision
	iAry2.signVal = ia.signVal
	iAry2.decimalSeparator = ia.decimalSeparator
	iAry2.thousandsSeparator = ia.thousandsSeparator
	iAry2.currencySymbol = ia.currencySymbol

	iAry2.BackUp.CopyIn(&ia.BackUp)

	return iAry2
}

// CopyToBackUp - Makes a copy of the current
// intAry object and saves it to BackUp.
//
// See ResetFromBackUp() to retrieve the
// last backup copy.
func (ia *IntAry) CopyToBackUp() {

	iBa := BackUpIntAry{}.New()

	iBa.intAry = make([]uint8, ia.intAryLen)

	for i := 0; i < ia.intAryLen; i++ {
		iBa.intAry[i] = ia.intAry[i]
	}

	iBa.intAryLen = ia.intAryLen
	iBa.firstDigitIdx = ia.firstDigitIdx
	iBa.lastDigitIdx = ia.lastDigitIdx
	iBa.isZeroValue = ia.isZeroValue
	iBa.precision = ia.precision
	iBa.signVal = ia.signVal
	iBa.decimalSeparator = ia.decimalSeparator
	iBa.thousandsSeparator = ia.thousandsSeparator
	iBa.currencySymbol = ia.currencySymbol

	ia.BackUp = iBa
}

// DecrementIntegerOne - Decrements the numeric value of the current
// intAry by subtracting '1'.
//
// IMPORTANT: This method assumes that SetIntAryLength() and
// SetIsZeroValue() have already been called.

func (ia *IntAry) DecrementIntegerOne() error {

	if ia.isZeroValue || ia.isIntegerZeroValue {
		ia.signVal = -1
	}

	intLen := ia.intAryLen - ia.precision
	intIdx := intLen - 1
	lastIdx := ia.intAryLen - 1

	n1 :=0
	n2 := 0
	carry := 0

	ia.isZeroValue = true
	ia.isIntegerZeroValue = true

	for i := lastIdx; i >= 0; i-- {
		n1 = int(ia.intAry[i])

		if i > intIdx {
			//  i > intIdx
			// This must be a fractional digit
			// Retain fractional digits

			if n1 != 0 {
				ia.isZeroValue = false
			}

			continue

		} else if i == intIdx {

			n2 = n1 + (-1 * ia.signVal)

			if n2 < 0 {
				n2 = n1 + 10 - 1
				carry = 1

			} else if n2 > 9 {
				n2 = n1 + 1 - 10
				carry = 1

			} else {
				carry = 0
			}

		} else {
			// Must be i < intIdx

			n2 = n1 + ( (ia.signVal * carry) * -1)

			if n2 < 0 {
				n2 = n1 + 10 - carry
				carry = 1
			} else if n2 > 9 {
				n2 = n1 - 10 + carry
				carry = 1
			} else {
				carry = 0
			}

		}

		if n2 != 0 {
			ia.isZeroValue = false
			ia.isIntegerZeroValue = false
		}

		ia.intAry[i] = uint8(n2)

	}

	if ia.isZeroValue && carry == 0 {
		ia.signVal = 1
	}

	if carry > 0 {

		ia.intAry = append([]uint8{1}, ia.intAry...)
		ia.intAryLen++

	} else if ia.intAry[0] == 0 && intLen > 1 {
		ia.intAry = ia.intAry[1:]
		ia.intAryLen--
	}

	return nil
}

// DivideByTwo - Divides the current value of
// intAry by 2. If parameter 'convertToNumStr'
// is set to 'true', the result will be converted
// to a number string.
func (ia *IntAry) DivideByTwo() {

	ia.OptimizeIntArrayLen(false)

	if ia.isZeroValue {

		ia.SetIntAryToZero(ia.precision)

		return
	}

	n1 := uint8(0)
	n2 := uint8(0)
	carry := uint8(0)

	for i := 0; i < ia.intAryLen; i++ {

		n1 = ia.intAry[i] + carry
		n2 = n1 / 2
		carry = (n1 - (n2 * 2)) * 10
		ia.intAry[i] = n2

	}

	if carry > 0 {
		ia.intAry = append(ia.intAry, 5)
		ia.intAryLen++
		ia.precision++
	}

	if ia.intAry[0] == 0 {
		ia.SetSignificantDigitIdxs()
		ia.intAry = ia.intAry[ia.firstDigitIdx:]
		ia.SetIntAryLength()
	}
}

// DivideByInt64 - Divide the current value of the intAry
// by an int64 'divisor' parameter passed to the method.
func (ia *IntAry) DivideByInt64(divisor int64, maxPrecision uint64) error {

	if divisor == 0 {
		return errors.New("'divisor' Equals Zero. Cannot divide by zero")
	}

	ia.OptimizeIntArrayLen(false)

	if ia.isZeroValue {

		ia.SetIntAryToZero(ia.precision)

		return nil
	}

	dSignVal := 1

	if divisor < 0 {
		dSignVal = -1
		divisor = divisor * -1
	}

	ia.signVal = dSignVal * ia.signVal

	n1 := int64(0)
	n2 := int64(0)
	carry := int64(0)
	iMaxPrecision := int(maxPrecision) + 1
	newAryLen := ia.intAryLen
	intAryLen := ia.intAryLen - ia.precision
	precisionCnt := 0

	for i := 0; i < newAryLen; i++ {

		if i >= intAryLen {
			precisionCnt++
		}

		if i < ia.intAryLen {
			n1 = int64(ia.intAry[i]) + carry
		} else {
			n1 = int64(0) + carry
		}

		n2 = n1 / divisor
		carry = (n1 - (n2 * divisor)) * 10

		if i < ia.intAryLen {
			ia.intAry[i] = uint8(n2)
		} else {
			ia.intAry = append(ia.intAry, uint8(n2))
		}

		if i == newAryLen-1 &&
			carry > 0 && precisionCnt <= iMaxPrecision {

			newAryLen++

		}

	}

	ia.precision = precisionCnt

	ia.intAryLen = newAryLen

	if precisionCnt >= iMaxPrecision {
		iMaxPrecision--
		ia.RoundToPrecision(iMaxPrecision)
	}

	if ia.intAry[0] == 0 {
		ia.SetSignificantDigitIdxs()
		ia.intAry = ia.intAry[ia.firstDigitIdx:]
		ia.SetIntAryLength()
	}

	return nil
}

// DivideByTenToPower - Divide intAry value by
// 10 raised to the power of the 'power'
// parameter.
//
// If parameter 'convertToNumStr' is set to 'true', the
// result will be automatically converted to a number string.
func (ia *IntAry) DivideByTenToPower(power uint) {

	if power == 0 {
		return
	}

	ia.precision += int(power)
	ia.intAryLen = len(ia.intAry)
	newLen := ia.precision + 1

	if ia.intAryLen < newLen {

		t := make([]uint8, newLen)

		deltaLen := newLen - ia.intAryLen

		for i := 0; i < newLen; i++ {

			if i < deltaLen {
				t[i] = 0
			} else {
				t[i] = ia.intAry[i-deltaLen]
			}

		}

		ia.intAry = make([]uint8, newLen)
		for i := 0; i < newLen; i++ {
			ia.intAry[i] = t[i]
		}

		ia.intAryLen = newLen
	}

}

// DivideThisBy - Divides the current value of intAry by the parameter iAry2.
// The result of this division is returned as an intAry.
//
// Maximum precision of the division result is controlled by the input
// parameter, 'maxPrecision'.
//
// If 'maxPrecision' is greater than or equal to zero ('0'),
// the number of digits to the right of the decimal place will
// not exceed 'maxPrecision'.

// 'maxPrecision' is set equal to minus one ('-1'), will be set
// to a maximum of 1,024 digits to the right of the decimal
// point.
//
func (ia *IntAry) DivideThisBy(iAry2 *IntAry, maxPrecision int) (IntAry, error) {

	ia.SetInternalFlags()
	iAry2.SetInternalFlags()

	if iAry2.isZeroValue {
		return IntAry{}.New(), errors.New("Error: divide by zero")
	}

	if maxPrecision < -1 {
		return IntAry{}.New(), errors.New("Error: Input parameter 'maxPrecision' is INVALID. 'maxPrecision' is less than -1")
	}

	if maxPrecision == -1 {
		maxPrecision = 1024
	}

	quotient := IntAry{}.New()
	quotient.SetIntAryToZero(0)

	if ia.isZeroValue {
		return quotient, nil
	}

	trialDividend := ia.CopyOut()

	divisor := iAry2.CopyOut()

	tensCount := IntAry{}.New()
	tensCount.SetIntAryToOne(0)

	newSignVal := 1

	if trialDividend.signVal != divisor.signVal {
		newSignVal = -1
	}

	if trialDividend.signVal == -1 {
		trialDividend.signVal = 1
	}

	if divisor.signVal == -1 {
		divisor.signVal = 1
	}

	dividendMag := trialDividend.GetMagnitude()
	divisorMag := divisor.GetMagnitude()
	deltaMag := uint(0)
	incrementVal := IntAry{}.New()
	incrementVal = divisor.CopyOut()

	if dividendMag > divisorMag {
		deltaMag = uint(dividendMag - divisorMag)
		tensCount.MultiplyByTenToPower(deltaMag)
		incrementVal.MultiplyThisBy(&tensCount, -1)

	} else if divisorMag > dividendMag {
		deltaMag = uint(divisorMag - dividendMag)
		trialDividend.MultiplyByTenToPower(deltaMag)
		tensCount.DivideByTenToPower(deltaMag)

	}

	compare := 0
	precisionCutOff := maxPrecision + dividendMag + 1

	for true {

		if quotient.precision >= precisionCutOff {
			quotient.RoundToPrecision(maxPrecision)
			quotient.OptimizeIntArrayLen(true)
			quotient.signVal = newSignVal
			return quotient, nil
		}

		compare = incrementVal.CompareAbsoluteValues(&trialDividend)

		if compare == 0 {
			// incrementalVal is equal to trialDividend
			quotient.AddToThis(&tensCount)
			quotient.RoundToPrecision(maxPrecision)
			quotient.OptimizeIntArrayLen(true)
			quotient.signVal = newSignVal
			return quotient, nil

		} else if compare == -1 {
			// incrementalVal < trialDividend
			quotient.AddToThis(&tensCount)

			// Calc Remainder
			trialDividend.SubtractFromThis(&incrementVal)

			continue

		} else {
			// Must Be compare == 1
			// incrementalVal > trialDividend

			tensCount.DivideByTenToPower(1)
			incrementVal.DivideByTenToPower(1)
		}

	}

	return quotient, nil
}

// Empty - Basically resets all the fields of the intAry
// structure to their 'zero' values.
func (ia *IntAry) Empty() {
	ia.intAry = []uint8{}
	ia.intAryLen = 0
	ia.integerLen = 0
	ia.significantIntegerLen = 0
	ia.significantFractionLen = 0
	ia.firstDigitIdx = -1
	ia.lastDigitIdx = -1
	ia.isZeroValue = true
	ia.isIntegerZeroValue = true
	ia.precision = 0
	ia.signVal = 1

	if ia.decimalSeparator == 0 {
		ia.decimalSeparator = '.'
	}

	if ia.currencySymbol == 0 {
		ia.currencySymbol = '$'
	}

	if ia.thousandsSeparator == 0 {
		ia.thousandsSeparator = ','
	}
}

// EmptyBackUp - Deletes the values
// currently stored as backup to the
// current intAry object.
func (ia *IntAry) EmptyBackUp() {
	ia.BackUp = BackUpIntAry{}.New()
}

// Equals - Returns 'true' if all field values
// of the current intAry object are equal to all
// field values of the input parameter intAry object,
// 'iAry2'.
func (ia *IntAry) Equals(iAry2 *IntAry) bool {
	ia.SetInternalFlags()
	iAry2.SetInternalFlags()

	if ia.intAryLen != iAry2.intAryLen {

		return false
	}

	for i := 0; i < iAry2.intAryLen; i++ {
		if ia.intAry[i] != iAry2.intAry[i] {
			return false
		}
	}

	if ia.integerLen != iAry2.integerLen ||
		ia.significantIntegerLen != iAry2.significantIntegerLen ||
		ia.significantFractionLen != iAry2.significantFractionLen ||
		ia.firstDigitIdx != iAry2.firstDigitIdx ||
		ia.lastDigitIdx != iAry2.lastDigitIdx ||
		ia.isZeroValue != iAry2.isZeroValue ||
		ia.isIntegerZeroValue != iAry2.isIntegerZeroValue ||
		ia.precision != iAry2.precision ||
		ia.signVal != iAry2.signVal ||
		ia.decimalSeparator != iAry2.decimalSeparator ||
		ia.thousandsSeparator != iAry2.thousandsSeparator ||
		ia.currencySymbol != iAry2.currencySymbol {

		return false
	}

	return true
}

// Floor - Math 'Floor' function. Finds the
// integer number which is less than or
// equal to the value of the current intAry.
func (ia *IntAry) Floor() (IntAry, error) {

	err := ia.IsIntAryValid("Floor() - ")

	if err != nil {
		return IntAry{}, err
	}

	iAry2 := IntAry{}.New()

	if ia.isZeroValue {
		iAry2.SetIntAryToZero(ia.precision)
	}

	hasFracDigits, err := ia.HasFractionalDigits()

	if err != nil {
		return iAry2, err
	}

	if !hasFracDigits {
		// There are NO non-zero digits to the
		// right of the decimal place
		iAry2 = ia.CopyOut()
		return iAry2, nil
	}

	intLen := ia.intAryLen - ia.precision
	intIdx := intLen - 1

	if hasFracDigits && ia.signVal > 0 {
		// There ARE non-zero digits to the right of the
		// decimal place
		t := make([]uint8, ia.intAryLen)

		for i := 0; i < intLen; i++ {
			t[i] = ia.intAry[i]
		}

		iAry2.intAry = t[0:]
		iAry2.intAryLen = ia.intAryLen
		iAry2.precision = ia.precision
		iAry2.signVal = ia.signVal
		return iAry2, nil
	}

	// The number has non-zero digits to
	// the right of the decimal place and
	// the number sign is minus (- or ia.signVal = -1)

	t := make([]uint8, ia.intAryLen+1)

	n1 := uint8(0)
	n2 := uint8(0)
	carry := uint8(0)

	for i := intIdx; i >= 0; i-- {

		n1 = ia.intAry[i]

		if i == intIdx {

			if n1+1 > 9 {
				n2 = n1 + 1 - 10
				carry = 1
			} else {
				n2 = n1 + 1
				carry = 0
			}

		} else {

			if n1+carry > 9 {
				n2 = n1 + carry - 10
				carry = 1
			} else {
				n2 = n1 + carry
				carry = 0
			}
		}

		t[i+1] = n2

	}

	if carry != 0 {
		t[0] = carry
		iAry2.intAry = t[0 : ia.intAryLen+1]

	} else {

		iAry2.intAry = t[1 : ia.intAryLen+1]
	}

	iAry2.precision = ia.precision
	iAry2.signVal = ia.signVal
	iAry2.intAryLen = len(iAry2.intAry)
	iAry2.SetIsZeroValue()

	return iAry2, nil
}

// GetAbsoluteValue - Returns an intAry which represents
// the Absolute Value of the current intAry
func (ia *IntAry) GetAbsoluteValue() IntAry {

	absIa := ia.CopyOut()

	absIa.SetAbsoluteValueThis()

	return absIa

}

// GetBigInt - Returns the current value of
// this intAry object as a big integer
// (*big.Int)
func (ia *IntAry) GetBigInt() *big.Int {

	result := big.NewInt(0).SetInt64(0)
	big10 := big.NewInt(0).SetInt64(10)

	ia.SetInternalFlags()

	for i := 0; i < ia.intAryLen; i++ {
		result = big.NewInt(0).Mul(result, big10)
		result = big.NewInt(0).Add(result, big.NewInt(0).SetInt64(int64(ia.intAry[i])))

	}

	if ia.signVal == -1 {

		bigMinusOne := big.NewInt(0).SetInt64(-1)

		result = big.NewInt(0).Mul(result, bigMinusOne)
	}

	return result
}

// GetCurrencySymbol - returns a type 'rune'
// which represents the setting for currency
// symbol in the current IntAry object.
//
// To set the value of currency symbol, see
// the method SetCurrencySymbol().
func (ia *IntAry) GetCurrencySymbol() rune {
	return ia.currencySymbol
}


// GetDecimalSeparator - returns a type 'rune'
// which represents the setting for decimal
// separator in the current IntAry object.
//
// A decimal separator separates integer digits
// from fractional digits which the value of
// the current IntAry object is expressed as a
// string.
//
// To set the value of decimal separator, see
// method SetDecimalSeparator().
func (ia *IntAry) GetDecimalSeparator() rune {
	return ia.decimalSeparator
}

// GetFractionalDigits - Examines the current intAry and
// returns another intAry consisting of the fractional
// digits to the right of the decimal point from the
// current intAry object.
//
// Note: The Sign Value of the returned int Ary is always
// positive or +1.
//
// The return intAry will display fractional digits with
// a leading integer digit of zero. Example '0.5678'
func (ia *IntAry) GetFractionalDigits() (IntAry, error) {

	err := ia.IsIntAryValid("GetFractionalDigits() - ")

	if err != nil {
		return IntAry{}, err
	}

	iAry2 := IntAry{}.New()

	iAry2.SetIntAryToZero(0)

	if ia.precision == 0 {
		return iAry2, nil
	}

	fracIdx := ia.intAryLen - ia.precision

	iAry2.intAry = make([]uint8, ia.precision+1)
	idx := 1
	for i := fracIdx; i < ia.intAryLen; i++ {
		iAry2.intAry[idx] = ia.intAry[i]
		idx++
	}

	iAry2.precision = ia.precision
	iAry2.signVal = 1
	iAry2.SetInternalFlags()

	return iAry2, nil
}

// GetIntegerDigits - Examines the current intAry object
// and returns a new intAry consisting of only the integer
// digits to the left of the decimal point in the current
// intAry object.
func (ia *IntAry) GetIntegerDigits() (IntAry, error) {

	err := ia.IsIntAryValid("GetFractionalDigits() - ")

	if err != nil {
		return IntAry{}, err
	}

	iAry2 := IntAry{}.New()

	if ia.isZeroValue {
		iAry2.SetIntAryToZero(0)
		return iAry2, nil
	}

	intLen := ia.intAryLen - ia.precision

	iAry2.intAry = make([]uint8, intLen)

	for i := 0; i < intLen; i++ {
		iAry2.intAry[i] = ia.intAry[i]
	}

	iAry2.signVal = ia.signVal
	iAry2.precision = 0

	iAry2.SetInternalFlags()

	if iAry2.isZeroValue {
		iAry2.signVal = 1
	}

	return iAry2, nil

}

// GetInt - returns the value of the current
// intAry as a 32-bit integer. If the intAry
// exceeds the maximum or minimum values for
// int, an error will be thrown.
//
// Minimum and Maximum Values for 32-bit Integer
// (int32 & int Types): -2147483648 to 2147483647 .
// Anything out this range will generate an error.
//
// Reference: https://golang.org/ref/spec#Numeric_types
//
func (ia *IntAry) GetInt() (int, error) {

	maxInt := big.NewInt(0).SetInt64(int64(math.MaxInt32))

	minInt := big.NewInt(0).SetInt64(int64(math.MinInt32))

	result := ia.GetBigInt()

	compare := result.Cmp(maxInt)

	if compare == 1 {
		return int(0), errors.New("error: the value of this intAry object exceeds the maximum allowable value for the int type")
	}

	compare = result.Cmp(minInt)

	if compare == -1 {
		return int(0), errors.New("error: the value of this intAry object is less than the minimum allowable value for the int type")
	}

	return int(result.Int64()), nil

}

// GetInt64 - Returns the value of the current
// intAry as an int64 value. If the intAry exceeds
// the maximum or minimum values for int64, an error will be
// thrown.
//
// Minimum and Maximum Values for 64-bit Integer (int64 Type):
// -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
// Anything outside this range will generate an error.
//
// Reference: https://golang.org/ref/spec#Numeric_types
func (ia *IntAry) GetInt64() (int64, error) {

	maxI64 := big.NewInt(0).SetInt64(math.MaxInt64)

	minI64 := big.NewInt(0).SetInt64(math.MinInt64)

	result := ia.GetBigInt()

	compare := result.Cmp(maxI64)

	if compare == 1 {
		return int64(0), errors.New("error: the value of this intAry object exceeds the maximum allowable value for the int64 type")
	}

	compare = result.Cmp(minI64)

	if compare == -1 {
		return int64(0), errors.New("error: the value of this intAry object is less than the minimum allowable value for the int64 type")
	}

	return result.Int64(), nil
}


// GetIntAryElement - Returns an element of the
// internal integer array maintained by the
// current IntAry object.
//
// Input Parameter:
// index int - 	The element returned is based on the
// 							integer index passed to the method. If
//							the index is outside the range of
//							the internal array, an error is returned.
//
// Return Value:	If the 'index' passed to the method is valid,
//								the method will return an integer of type
//								uint8 representing the value of the internal
//								integer array at the specified index.
func (ia *IntAry) GetIntAryElement(index int) (uint8, error) {


	if index < 0 || index > (ia.intAryLen - 1) {
		return 0, fmt.Errorf("Error: GetIntAryElement(index int) - Index is INVALID! index Out Of Array Bounds! index= '%v'", index)
	}

	result := ia.intAry[index]

	return result, nil

}

// ************************************************
// BE CAREFUL!! This returns a reference (pointer)
// to the internal integer array for this IntAry
// object.
// ************************************************
// GetIntAry returns the internal integer
// array used by the IntAry object.
//
// Return parameters:
// []uint8 = an array of unsigned 8-bit integers
//
// int = The length of the []uint8 array returned
// 			 above.
//
// ************************************************
//			BE CAREFUL!!!!!!!
// ************************************************
// Be careful!! - This returns a slice and therefore
// a reference (i.e. pointer) to the internal array
// for this IntAry object. If you alter this array
// externally, it will alter the internal array
// values.
//
// IMPORTANT: If you alter the returned array, the
// current IntAry object may become invalid. After
// changing the returned array, Call method
// SetInternalFlags() to ensure that the IntAry
// object correctly reflects the changed value.
//
// For a an array which can be safely manipulated,
// see method GetIntAryDeepCopy().
//
// To Append Elements to the internal integer array,
// see method AppendToIntAry()
func (ia *IntAry) GetIntAry() ([]uint8, int) {

	ia.SetInternalFlags()

	return ia.intAry, ia.intAryLen
}

// GetIntAryDeepCopy - returns a deep copy of the
// internal integer array maintained by this IntAry
// object. Unlike the array returned by method
// GetIntAry(), the array returned by this method
// is not a reference or pointer.
//
// This means that the array returned by this method
// may altered externally without changing the value
// of the original internal array maintained by this
// IntAry object.
func (ia *IntAry) GetIntAryDeepCopy() ([]uint8, int) {
	ia.SetInternalFlags()
	ary := make([]uint8, ia.intAryLen)

	for i:= 0; i < ia.intAryLen; i++ {
		ary[i] = ia.intAry[i]
	}

	return ary, ia.intAryLen
}

// GetIntAryLength - returns the length of the
// internal integer array maintained by the current
// IntAry object.
func (ia *IntAry) GetIntAryLength() int {
	ia.SetInternalFlags()
	return ia.intAryLen
}

// GetIntAryStats - returns a series of descriptive
// statistics about the internal Integer Array
// maintained by the current IntAry object.
//
// The method returns an IntAryStatsDto object which
// includes the following information on the current
// IntAry object:
//
// IntAryLen  - Type int: Length of the internal integer array
//
// IntegerLen - Type int: Number of integer digits in the current
//							IntAry value.
//
// SignificantIntegerLen - 	Type int: The number of non-zero integer
//													digits in the current IntAry value
//
// SignificantFractionLen - Type int: The number of non-zero digits
//													to the right of the decimal point in
//													the current IntAry value
//
// Precision							- Type int: The number of digits to the right
//													of the decimal point in the current
//													IntAry value
//
// SignVal                - Type int: Either +1 or -1 indicating the sign
//													of the current IntAry value
//
// FirstDigitIdx          - Type int: The IntAry index of the first
//													non-zero digit in the current
//													internal integer array
//
// LastDigitIdx          -  Type int: The IntAry index of the last non-zero
//													digit in the internal integer array
//
// IsZeroValue           -	A boolean value indicating whether the
//													current IntAry value is zero
//
// IsIntegerZeroValue    - 	A boolean value indicating whether integer
//													value to the left of the decimal point is
//													zero
//
// DecimalSeparator      -  Type rune: The character used to separate
//													integer and fractional elements of the current
//													IntAry value when presented in string format
//
// ThousandsSeparator    -	Type rune: The character used to separate
//													thousands when the current IntAry value is
//													presented in string format
//
// CurrencySymbol        -  Type rune: The character used to designate
//													currency when the current IntAry value is
//													presented as a currency string.
func (ia *IntAry) GetIntAryStats() IntAryStatsDto {
	iStats := IntAryStatsDto{}
	ia.SetInternalFlags()

	iStats.IntAryLen = ia.intAryLen
	iStats.IntegerLen = ia.integerLen
	iStats.SignificantIntegerLen = ia.significantIntegerLen
	iStats.SignificantFractionLen = ia.significantFractionLen
	iStats.Precision = ia.precision
	iStats.SignVal = ia.signVal
	iStats.FirstDigitIdx =  ia.firstDigitIdx
	iStats.LastDigitIdx = ia.lastDigitIdx
	iStats.IsZeroValue = ia.isZeroValue
	iStats.IsIntegerZeroValue = ia.isIntegerZeroValue
	iStats.DecimalSeparator = ia.decimalSeparator
	iStats.ThousandsSeparator = ia.thousandsSeparator
	iStats.CurrencySymbol = ia.currencySymbol

	return iStats
}

// GetNumStr - returns the current value
// of this intAry object as a number string.
func (ia *IntAry) GetNumStr() string {

	if ia.decimalSeparator == 0 {
		ia.decimalSeparator = '.'
	}

	ia.SetInternalFlags()

	var buffer bytes.Buffer

	if ia.signVal < 0 {
		buffer.WriteRune('-')
	}

	intLen := ia.intAryLen - ia.precision

	for i:=0 ; i < intLen; i++ {
		buffer.WriteRune(rune(ia.intAry[i] + 48))
	}

	if ia.precision > 0 {
		buffer.WriteRune(ia.decimalSeparator)

		for j:= 0; j < ia.precision; j++ {
			buffer.WriteRune(rune(ia.intAry[intLen] + 48))
			intLen++
		}

	}

	return buffer.String()
	// ia.ConvertIntAryToNumStr()
	// return ia.numStr
}

func (ia *IntAry) GetMagnitude() int {
	ia.SetSignificantDigitIdxs()
	return ia.intAryLen - ia.precision - ia.firstDigitIdx

}

// GetNthRootOfThis - Returns an intAry object equal to the 'nth Root'
// of the current intAry.
//
// Input Parameters:
// 'nthRoot' - uint value which specifies the root to calculate
// 'maxPrecision' - uint value which specifies the number of digits
//									to the right of the decimal place in the result.
func (ia *IntAry) GetNthRootOfThis(nthRoot, maxPrecision uint) (IntAry, error) {
	nthRt := NthRootOp{}

	return nthRt.GetNthRootIntAry(ia, nthRoot, maxPrecision)
}

// GetPrecision - returns the precision value
// for the current intAry object. 'precision'
// represents the number of digits to the right
// of the decimal point.
//
// The value of 'precision' returned by this
// method should always be >= zero (greater than
// or equal to zero '0' )
func (ia *IntAry) GetPrecision() int {
	return ia.precision
}

// GetScaleFactorBigInt - Returns a pointer to a Big Integer
// (*big.Int) which specifies the scale factor associated
// with this IntAry value.
func (ia *IntAry) GetScaleFactor() (*big.Int, error) {
	ia.SetInternalFlags()

	err := ia.IsIntAryValid("GetScaleFactorBigInt() - ")

	if err != nil {
		return big.NewInt(0), errors.New("This intAry is invalid")
	}

	if ia.precision == 0 {
		return big.NewInt(int64(1)), nil
	}

	base10 := big.NewInt(0).SetInt64(int64(10))

	bigPrecision := big.NewInt(0).SetInt64(int64(ia.precision))

	scaleFactor := big.NewInt(0).Exp(base10, bigPrecision, nil)

	return scaleFactor, nil

}

// GetSign - returns the sign of the current
// intAry object as an integer value.
//
// The sign value returned by this method should
// always be one of two values: +1 or -1 .
func (ia *IntAry) GetSign() int {
	return ia.signVal
}

// GetSquareRootOfThis - Returns an intAry object equal to the 'square root'
// of the current intAry.
//
// Input Parameters:
// 'maxPrecision' - uint value which specifies the number of digits
//									to the right of the decimal place in the result.
func (ia *IntAry) GetSquareRootOfThis(maxPrecision uint) (IntAry, error) {
	nthRt := NthRootOp{}

	return nthRt.GetSquareRootIntAry(ia, maxPrecision)
}

// GetThousandsSeparator - returns a value of type 'rune'
// which represents the thousands separator associated
// with the current IntAry object.
//
// The thousands separator is used to separate thousands
// in the integer digits of this IntAry value when that
// value is expressed as a string.
//
// In the US, the thousands separator is typically the comma
// character (','). In this example, the thousands separator
// is a comma: Example - '1,000,000,000'.
//
// To set the value of thousands separator, see the method
// SetThousandsSeparator().
func (ia *IntAry) GetThousandsSeparator() rune {
	return ia.thousandsSeparator
}

// HasFractionalDigits - This method examines the
// current intAry object to determine if there
// are non-zero digits to the right of the decimal
// place. If all digits to the right of the decimal
// place are zero, this method returns 'false'
//
// If non-zero digits are present to the right of the
// decimal place, the method returns 'true'.
func (ia *IntAry) HasFractionalDigits() (bool, error) {

	err := ia.IsIntAryValid("HasFractionalDigits() - ")

	if err != nil {
		return false, err
	}

	if ia.precision == 0 {
		return false, nil
	}

	ia.SetIntAryLength()

	intLen := ia.intAryLen - ia.precision

	if intLen < 1 {
		return false, fmt.Errorf("HasFractionalDigits() Error - Int Array integer length is less than 1. intLen= '%v'", intLen)
	}

	for i := intLen; i < ia.intAryLen; i++ {
		if ia.intAry[i] > 0 {
			return true, nil
		}
	}

	return false, nil
}

// IncrementIntegerOne - Increment the value of the
// current intAry by adding '1'
func (ia *IntAry) IncrementIntegerOne() error {

	if ia.isZeroValue || ia.isIntegerZeroValue {
		ia.signVal = 1
	}

	intLen := ia.intAryLen - ia.precision
	intIdx := intLen - 1
	lastIdx := ia.intAryLen - 1

	n1 := 0
	n2 := 0
	carry := 0

	ia.isZeroValue = true
	ia.isIntegerZeroValue = true

	for i := lastIdx; i >= 0; i-- {
		n1 = int(ia.intAry[i])

		if i > intIdx {
			//  i > intIdx
			// This must be a fractional digit
			// Retain fractional digits

			if n1 != 0 {
				ia.isZeroValue = false
			}

			continue

		} else if i == intIdx {

			n2 = n1 + (1 * ia.signVal)

			if n2 < 0 {
				n2 = n1 + 10 - 1
				carry = 1

			} else if n2 > 9 {
				n2 = n1 + 1 - 10
				carry = 1

			} else {
				carry = 0
			}

		} else {
			// Must be i < intIdx

			n2 = n1 + ((ia.signVal * carry) * 1)

			if n2 < 0 {
				n2 = n1 + 10 - carry
				carry = 1
			} else if n2 > 9 {
				n2 = n1 - 10 + carry
				carry = 1
			} else {
				carry = 0
			}

		}

		if n2 != 0 {
			ia.isZeroValue = false
			ia.isIntegerZeroValue = false
		}

		ia.intAry[i] = uint8(n2)

	}

	if ia.isZeroValue && carry == 0 {
		ia.signVal = 1
	}

	if carry > 0 {

		ia.intAry = append([]uint8{1}, ia.intAry...)
		ia.intAryLen++

	} else if ia.intAry[0] == 0 && intLen > 1 {
		ia.intAry = ia.intAry[1:]
		ia.intAryLen--
	}

	return nil
}

// IsIntAryValid - Examines the current intAry and returns
// an error if the intAry object is found to be invalid.
func (ia *IntAry) IsIntAryValid(errName string) error {

	ia.SetInternalFlags()

	if ia.signVal != -1 && ia.signVal != 1 {
		return fmt.Errorf("%v Sign Value is INVALID! Sign Value= '%v'", errName, ia.signVal)
	}

	if ia.precision < 0 {
		return fmt.Errorf("%v precision Value is INVALID! Sign Value= '%v'", errName, ia.precision)
	}


	if ia.precision >= ia.intAryLen {
		return fmt.Errorf("%v error: precision is greater than or equal to IntArray length - ia.precision= %v  ia.intAryLen= %v ", errName, ia.precision, ia.intAryLen)
	}

	if ia.integerLen == 0 {
		return fmt.Errorf("%v error: integer length is zero - missing leading integer zero", errName)
	}

	return nil
}

// Inverse - Returns the inverse of the the current intAry's
// value.
//
// Input Parameter:
//	maxPrecision int -	determines the number of digits to the
//  										right of the decimal point in the result.
//
//											Note: if 'maxPrecision' is set equal to negative
//														one (-1), the maximum number of decimals is
//														set to 1024 digits to the right of the decimal
//														place
func (ia *IntAry) Inverse(maxPrecision int) (IntAry, error) {

	if maxPrecision == -1 {
		maxPrecision = 1024
	}

	if maxPrecision < 0 {
		return IntAry{}.New(), errors.New("InverseOfThis() - Input parameter 'maxPrecision' is INVALID. 'maxPrecision' cannot be less than zero.")
	}

	iaOne, err := IntAry{}.NewInt(1, 0)

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("InverseOfThis() - Error returned from intAry{}.NewInt(1, 0). Error= %v", err)
	}

	iaInverse, err := iaOne.DivideThisBy(ia, maxPrecision)

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("InverseOfThis() - Error returned from iaOne.DivideThisBy(ia, maxPrecision). Error= %v", err)
	}

	return iaInverse, nil
}

// MultiplyByTwoToPower Multiply the existing value
// of the intAry by 2 to the power of the passed in
// parameter.
//
func (ia *IntAry) MultiplyByTwoToPower(power uint) {

	ia.SetIntAryLength()

	if power == 0 {
		return
	}

	for h := uint(0); h < power; h++ {
		n1 := uint8(0)
		carry := uint8(0)

		for i := ia.intAryLen - 1; i >= 0; i-- {

			n1 = (ia.intAry[i] * 2) + carry

			if n1 > 9 {
				n1 = n1 - 10
				carry = 1
			} else {
				carry = 0
			}

			ia.intAry[i] = n1
		}

		if carry > 0 {
			ia.intAry = append([]uint8{1}, ia.intAry...)
			ia.intAryLen++
		}

	}

}

// MultiplyByTenToPower - The value of intAry is multiplied
// by 10 to the power of the passed in parameter.
func (ia *IntAry) MultiplyByTenToPower(power uint) {

	if power == 0 {
		return
	}
	for i := uint(0); i < power; i++ {

		if ia.precision > 0 {
			ia.precision--
			continue
		}

		ia.intAry = append(ia.intAry, 0)
	}

	ia.intAryLen = len(ia.intAry)

	if ia.precision < 0 {
		ia.precision = 0
	}

}

func (ia *IntAry) MultiplyThisBy(ia2 *IntAry, maxPrecision int) error {

	return ia.Multiply(ia, ia2, ia, maxPrecision)


}

// Multiply - Multiplies intAry ia1 by intAry ia2 and stores the multiplication result
// in intAry iaResult.
//
// Parameters
// ==========
//
// 'ia1' - 				Pointer to an intAry object. In this multiplication operation, 'ia1'
// 								is the multiplicand.
//
// 'ia2' - 				Pointer to an intAry object. In this multiplication operation, 'ia2'
// 								is the multiplier.
//
//  'iaResult' -	Pointer to an intAry object which will be populated with the result
//								of the multiplication operation. The multiplication operation is achieved
//								by multiplying 'ia1' by 'ia2'.
//
// 'maxResultPrecision' int -
//								'maxResultPrecision' will determine the maximum
// 								number of digits to the right of the decimal
//								place in the result.
//
//								Valid values are -1 and values >= zero ('0')
//        				Values less than -1 will trigger an error.
//
//								A value of -1 signals that no limit will be placed on
//								the number of decimals places to right of the decimal
//								point in the result.
//
func (ia *IntAry) Multiply(ia1, ia2, iaResult *IntAry, maxResultPrecision int) error {

	if maxResultPrecision < -1 {
		return fmt.Errorf("Error: Parameter 'maxResultPrecision' is less than -1. maxResultPrecision= %v", maxResultPrecision)
	}

	ia1.SetInternalFlags()
	ia2.SetInternalFlags()


	if ia1.isZeroValue || ia2.isZeroValue {
		iaResult.SetIntAryToZero(ia2.GetPrecision())
		return nil
	}

	newSignVal := 1

	if ia1.signVal != ia2.signVal {
		newSignVal = -1
	}

	newPrecision := ia1.precision + ia2.precision

	newIntAryLen := ia1.intAryLen + ia2.intAryLen

	resultAry := make([]uint8, newIntAryLen)

	carry := uint8(0)
	multiplier := uint8(0)
	multiplicand := uint8(0)
	product := uint8(0)
	resultIdx := 0
	offset := 0

	for i := ia2.intAryLen -1; i >=0 ; i-- {
		multiplicand = ia2.intAry[i]
		offset++
		nextResultIdx := newIntAryLen - offset

		for j := ia1.intAryLen - 1; j >= 0 ; j-- {

			multiplier = ia1.intAry[j]

			product = multiplier * multiplicand

			resultIdx = nextResultIdx

			resultAry[resultIdx]+= product


			for resultAry[resultIdx] > 9 {
				carry = resultAry[resultIdx] / 10
				resultAry[resultIdx] = resultAry[resultIdx] - (carry * 10)

				resultIdx--

				resultAry[resultIdx]+= carry

			}

			carry = 0
			nextResultIdx--
		}

	}


	if newIntAryLen - newPrecision > 1 && resultAry[0] == 0 {

		iaResult.intAry = resultAry[1:]
		newIntAryLen--
	} else {
		iaResult.intAry = resultAry
	}

	iaResult.precision = newPrecision
	iaResult.signVal = newSignVal
	iaResult.intAryLen = newIntAryLen
	iaResult.integerLen = newIntAryLen - newPrecision
	iaResult.isZeroValue = false

	if maxResultPrecision > -1 && maxResultPrecision < newPrecision {
		iaResult.SetPrecision(maxResultPrecision, true)
	}

	return nil
}

// New - Creates a new blank intAry object.
// Usage: ia := intAry{}.New()
func (ia IntAry) New() IntAry {
	iAry := IntAry{}
	iAry.intAry = []uint8{}
	iAry.intAryLen = 0
	iAry.integerLen = 0
	iAry.significantIntegerLen = 0
	iAry.significantFractionLen = 0
	iAry.firstDigitIdx = -1
	iAry.lastDigitIdx = -1
	iAry.isZeroValue = true
	iAry.isIntegerZeroValue = true
	iAry.precision = 0
	iAry.signVal = 1
	iAry.decimalSeparator = '.'
	iAry.thousandsSeparator = ','
	iAry.currencySymbol = '$'
	iAry.BackUp = BackUpIntAry{}.New()

	return iAry
}

// NewBigInt - Creates a new intAry object initialized
// to the value of input parameter 'num' which is passed
// as type '*big.Int'. Note that 'num' may be a positive
// or negative number.
//
// Input parameter 'precision' indicates the number of digits
// to be formatted to the right of the decimal place.
//
// Example:
//  intDigits     precision     	    result
//  946254  			   3							   946.254
//  946254				   0							   946254
//  -946254  			   3					      -946.254
//  -946254				   0						    -946254
//
// Usage:
// num := big.NewInt(123456)
// precision := uint(3)
// ia, err := intAry{}.NewBigInt(num, precision)
//
func (ia IntAry) NewBigInt(num *big.Int, precision uint) (IntAry, error) {

	iAry := IntAry{}.New()
	err := iAry.SetIntAryWithBigInt(num, precision)

	if err != nil {
		return IntAry{}, err
	}

	return iAry, nil

}

// NewFloat32 - Creates a new intAry object initialized
// to the value of input parameter 'num' which is passed
// as type 'float32'. Input parameter 'precision' is used
// to set the input precision of 'num'.
//
// Input parameter 'precision' must be set to a number greater than or equal
// to zero.  It may also be set to a value of -1 which causes the number
// to be formatted to the smallest number of digits to right of the decimal
// point.
//
// Usage:
// num := float32(123.456000)
// precision := 3 - signals that only the decimals 456 will included as input
// ia, err := intAry{}.NewFloat32(num, precision)
//
// The value resulting from the usage example above would be '123.456'.
// If the value of 'precision' was set to -1, the resulting value would
// also be '123.456'
//
// Note: if precision is a positive value, and less than the current number
// of digits to the right of the decimal place, rounding may occur.
//
func (ia IntAry) NewFloat32(num float32, precision int) (IntAry, error) {

	iAry := IntAry{}.New()

	if precision < -1 {
		return iAry, fmt.Errorf("NewFloat32() Error: 'precision' INVALID. 'precision' must be greater than or equal to -1. precision='%v'", precision)
	}


	err := iAry.SetIntAryWithFloat32(num, precision)

	if err != nil {
		return IntAry{}, err
	}

	return iAry, nil
}

// NewFloat64 - Creates a new intAry object initialized
// to the value of input parameter 'num' which is passed
// as type 'float64'. Input parameter 'precision' is used
// to set the input precision of 'num'.
//
// Input parameter 'precision' must be set to a number greater than or equal
// to zero.  It may also be set to a value of -1 which causes the number
// to be formatted to the smallest number of digits to right of the decimal
// point.
//
// Input parameter 'precision' must be set to a
// 									number greater than or equal to zero.  It may
// 									also be set to a value of -1 which causes the
// 									number to be formatted to the smallest number
// 									of digits to right of the decimal point.
//
// Usage:
// num := float64(123.456000)
// precision := 3 - signals that only the decimals 456 will included as input
// ia, err := intAry{}.NewFloat64(num, precision)
//
func (ia IntAry) NewFloat64(num float64, precision int) (IntAry, error) {

	iAry := IntAry{}.New()

	if precision < -1 {
		return iAry, fmt.Errorf("NewFloat64() Error: 'precision' INVALID. 'precision' must be greater than or equal to -1. precision='%v'", precision)
	}


	err := iAry.SetIntAryWithFloat64(num, precision)

	if err != nil {
		return IntAry{}, err
	}

	return iAry, nil
}


// NewFloatBig - Creates a new intAry object initialized
// to the value of input parameter 'num' which is passed
// as type '*big.Float'.
//
// Input Parameters:
//
// num *big.Float -	This floating value will be used to
//									to set the value of a new IntAry object
//
// precision int -	'precision' is applied to the input parameter
//									'num' (*big.Float) to determine the number of
//									digits to the right of the decimal point which
//									will be applied to the resulting value.
//
//									Input parameter 'precision' must be set to a
// 									number greater than or equal to zero.  It may
// 									also be set to a value of -1 which causes the
// 									number to be formatted to the smallest number
// 									of digits to right of the decimal point.
//
// Usage:
// num, err := big.NewFloatBig(123.4560, 3)
//
// Result:
// num = '123.456 - Note: precision parameter may result in rounding.
func (ia IntAry) NewFloatBig(num *big.Float, precision int) (IntAry, error) {

	iAry := IntAry{}.New()

	if precision < -1 {
		return iAry, fmt.Errorf("NewFloatBig() Error: 'precision' INVALID. 'precision' must be greater than or equal to -1. precision='%v'", precision)
	}

	err := iAry.SetIntAryWithFloatBig(num, precision)

	if err != nil {
		return IntAry{}, err
	}

	return iAry, nil
}

// NewInt32 - Creates a new intAry object initialized
// to the value of input parameter 'num' which is passed
// as type 'int32'.
//
// Input parameter 'precision' indicates the number of digits
// to be formatted to the right of the decimal place.
//
// Usage:
// num := int32(123456)
// precision := uint(3)
// ia, err := intAry{}.NewInt32(num, precision)
//
func (ia IntAry) NewInt32(num int32, precision uint) (IntAry, error) {

	iAry := IntAry{}.New()
	err := iAry.SetIntAryWithInt32(num, precision)

	if err != nil {
		return IntAry{}, err
	}

	return iAry, nil

}

// NewInt64 - Creates a new intAry object initialized
// to the value of input parameter 'num' which is passed
// as type 'int64'.
//
// Input parameter 'precision' indicates the number of digits
// to be formatted to the right of the decimal place.
//
// Usage:
// num := int64(123456)
// precision := uint(3)
// ia, err := intAry{}.NewInt64(num, precision)
//
func (ia IntAry) NewInt64(num int64, precision uint) (IntAry, error) {

	iAry := IntAry{}.New()
	err := iAry.SetIntAryWithInt64(num, precision)

	if err != nil {
		return IntAry{}, err
	}

	return iAry, nil

}

// NewInt - Creates a new intAry object initialized
// to the value of input parameter 'num' which is passed
// as type 'int32'.
//
// Input parameter 'precision' indicates the number of digits
// to be formatted to the right of the decimal place.
//
// Usage:
// num := int(123456)
// precision := uint(3)
// ia, err := intAry{}.NewInt32(num, precision)
//
func (ia IntAry) NewInt(num int, precision uint) (IntAry, error) {

	iAry := IntAry{}.New()
	err := iAry.SetIntAryWithInt(num, precision)

	if err != nil {
		return IntAry{}, err
	}

	return iAry, nil
}

// NewNumStr - Creates a new intAry object initialized
// to the value of input parameter 'num' which is passed
// as type 'string'.
//
// Usage: ia := intAry{}.NewNumStr("123.456")
func (ia IntAry) NewNumStr(num string) (IntAry, error) {

	iAry := IntAry{}.New()
	err := iAry.SetIntAryWithNumStr(num)

	if err != nil {
		return IntAry{}, err
	}

	return iAry, nil

}

// OptimizeIntArrayLen - Eliminates Leading
// zeros from the front or integer portion
// of the integer string.
//
// If parameter 'optimizeFracDigits' is set
// equal to 'true', trailing zeros to the
// right of the decimal place will also be
// eliminated.
//
func (ia *IntAry) OptimizeIntArrayLen(optimizeFracDigits bool) {

	ia.SetInternalFlags()

	if ia.isZeroValue {
		return
	}

	integerLen := ia.intAryLen - ia.precision - ia.firstDigitIdx

	if optimizeFracDigits {

		ia.intAry = ia.intAry[ia.firstDigitIdx: ia.lastDigitIdx+1]
		ia.intAryLen = ia.lastDigitIdx - ia.firstDigitIdx + 1
	} else {
		ia.intAry = ia.intAry[ia.firstDigitIdx:]
		ia.intAryLen = ia.intAryLen - ia.firstDigitIdx
	}

	ia.precision = ia.intAryLen - integerLen

}

// Pow - Raises the value of the current intAry
// to the power designated by the input parameter 'power'.
// This method calls ia.pwrByTwos() which raises
// a number to a specified power using the
// exponentiation by squaring algorithm.
//
// Input Parameters:
// =================
// 'power' int -
// 				The input parameter 'power' may be either
// 				a positive or negative integer.
//
// 'maxResultPrecision' int -
//				'maxResultPrecision' will determine the maximum
// 				number of digits to the right of the decimal
//				place in the result.
//
//				Valid values are -1 and values >= zero ('0')
//        Values less than -1 will trigger an error.
//
//				A value of -1 signals that no limit will be placed on
//				the number of decimals places to right of the decimal
//				point in the result.
//
//	'internalPrecision' int -
// 				'internalPrecision' will control the number of digits of
//				accuracy to the right of the decimal point maintained by
//				internal multiplication operations used in raising the intAry
//				value to the designated power.
//
//				Valid values are -1 and values >= zero ('0')
//        Values less than -1 will trigger an error.
//
//				A value of -1 signals that no limit will be placed on
//				the number of decimals places to right of the decimal
//				point during internal multiplication operations.
//
func (ia *IntAry) Pow(power, maxResultPrecision, internalPrecision int) error {

	pwr := big.NewInt(int64(power))
	return ia.pwrByTwos(pwr, maxResultPrecision, internalPrecision)

}

// PowThisSquared - Raises the value of the current intAry object
// to a power of '2'.  Essentially the new value of this intAry object
// is equal to the original value squared.
func (ia *IntAry)PowThisSquared() error {

	return ia.MultiplyThisBy(ia, -1)
}


// PowByTwos - Raises the value of the current intAry to the power indicated by the parameter,
// 'power'.
//
// Input Parameters:
// =================
// 'power' int -
// 				The input parameter 'power' may be either
// 				a positive or negative integer.
//
// 'maxResultPrecision' int -
//				'maxResultPrecision' will determine the maximum
// 				number of digits to the right of the decimal
//				place in the result.
//
//				Valid values are -1 and values >= zero ('0')
//        Values less than -1 will trigger an error.
//
//				A value of -1 signals that no limit will be placed on
//				the number of decimals places to right of the decimal
//				point in the result.
//
//	'internalPrecision' int -
// 				'internalPrecision' will control the number of digits of
//				accuracy to the right of the decimal point maintained by
//				internal multiplication operations used in raising the intAry
//				value to the designated power.
//
//				Valid values are -1 and values >= zero ('0')
//        Values less than -1 will trigger an error.
//
//				A value of -1 signals that no limit will be placed on
//				the number of decimals places to right of the decimal
//				point during internal multiplication operations.
//
func (ia *IntAry) PowByTwos(power *big.Int, maxResultPrecision, internalPrecision int) error {


	return ia.pwrByTwos(power, maxResultPrecision, internalPrecision)
}


// pwrByTwos - Raises a *big.Int 'base', to the specified 'power'
// using the Exponentiation by squaring algorithm.
// See:
// https://en.wikipedia.org/wiki/Exponentiation_by_squaring
// https://en.wikipedia.org/wiki/Exponentiation_by_squaring#Computation_by_powers_of_2
// This method is based on revised code taken in part from Ye Lin Aung.
// https://stackoverflow.com/questions/30182129/calculating-large-exponentiation-in-golang
// Algorithm modified by Mike Rapp to achieve improved performance.
//
// Input Parameters:
// =================
// 'power' int -
// 				The input parameter 'power' may be either
// 				a positive or negative integer.
//
// 'maxResultPrecision' int -
//				'maxResultPrecision' will determine the maximum
// 				number of digits to the right of the decimal
//				place in the result.
//
//				Valid values are -1 and values >= zero ('0')
//        Values less than -1 will trigger an error.
//
//				A value of -1 signals that no limit will be placed on
//				the number of decimals places to right of the decimal
//				point in the result.
//
//	'internalPrecision' int -
// 				'internalPrecision' will control the number of digits of
//				accuracy to the right of the decimal point maintained by
//				internal multiplication operations used in raising the intAry
//				value to the designated power.
//
//				Valid values are -1 and values >= zero ('0')
//        Values less than -1 will trigger an error.
//
//				A value of -1 signals that no limit will be placed on
//				the number of decimals places to right of the decimal
//				point during internal multiplication operations.
//
func (ia *IntAry) pwrByTwos(power *big.Int, maxResultPrecision, internalPrecision int ) error {

	if maxResultPrecision < -1 {
		return fmt.Errorf("Error: Parameter maxResultPrecision is less than -1. maxResultPrecision= %v", maxResultPrecision)
	}

	if internalPrecision < -1 {
		return fmt.Errorf("Error: Parameter internalPrecision is less than -1. internalPrecision= %v", internalPrecision)
	}


	ia.SetInternalFlags()
	tPower := big.NewInt(0).Set(power)
	one := big.NewInt(1)
	zero := big.NewInt(0)
	two := big.NewInt(2)

	if ia.isZeroValue {
		ia.SetIntAryToZero(0)
		return nil
	}

	if tPower.Cmp(two) == -1 {

		if tPower.Cmp(zero) == -1 {

			ia2, err := ia.Inverse(internalPrecision)

			if err != nil {
				return fmt.Errorf("intAry.pwrByTwos = Error returned from ia.Inverse(-1): Error= %v", err)
			}

			ia.CopyIn(&ia2, false)
			tPower = big.NewInt(0).Mul(tPower, big.NewInt(-1))

		} else if tPower.Cmp(one) == 0 {
			// no change in value. x^1 == x
			return nil
		} else if tPower.Cmp(zero) == 0 {
			ia.SetIntAryToOne(0)
			return nil
		}

	}

	tBase := ia.CopyOut()

	ia.SetIntAryToOne(0)

	for tPower.Cmp(zero) == 1 {
		//temp, _:= intAry{}.NewNumStr("0")

		if big.NewInt(0).Mod(tPower, two).Cmp(one) == 0 {
			//temp = big.NewInt(0).Mul(result, tBase)
			//result = big.NewInt(0).Set(temp)
			err := ia.MultiplyThisBy(&tBase, internalPrecision)
			//fmt.Println("ia Precision = ", ia.GetPrecision())

			if err != nil {
				return fmt.Errorf("intAry.pwrByTwos() - Error From result.MultiplyThisBy(&tBase, true). Error= %v", err)
			}

			if tPower.Cmp(one) == 0 {
				if maxResultPrecision > -1 && maxResultPrecision < ia.GetPrecision() {
					ia.SetPrecision(maxResultPrecision, true)
				}

				return nil
			}
		}


		err:= tBase.MultiplyThisBy(&tBase, internalPrecision)
		if err != nil {
			return fmt.Errorf("intAry.pwrByTwos() - Error From tBase.MultiplyThisBy(&temp, true). Error= %v", err)
		}


		tPower = big.NewInt(0).Div(tPower, two)
	}

	if maxResultPrecision == -1 {
		return nil
	}

	if maxResultPrecision > -1 && maxResultPrecision < ia.GetPrecision()  {
		ia.SetPrecision(maxResultPrecision, true)
	}

	return nil
}

func (ia *IntAry) pwrByOnes(power, maxPrecision int) error {
	ia.SetInternalFlags()

	if ia.isZeroValue {
		return nil
	}

	if power == 0 {
		ia.SetIntAryToOne(ia.precision)
		return nil
	}

	if power == 1 {
		return nil
	}

	oldPowerSignVal := 1

	if power < 0 {
		power = power * -1
		oldPowerSignVal = -1
	}

	power--

	multiplier := ia.CopyOut()

	for i := 0; i < power; i++ {
		ia.MultiplyThisBy(&multiplier, -1)
	}

	if oldPowerSignVal == 1 {

		return nil
	}

	// Power must be negative.
	iaInv, err := ia.Inverse(maxPrecision)

	if err != nil {
		fmt.Errorf("Pow() - Error returned from ia.Inverse(maxPrecision). Error= %v ", err)
	}

	ia.CopyIn(&iaInv, false)

	return nil

}

// PrefixToIntAry - Adds an integer of type
// uint8 to the front of the existing internal
// integer array maintained by the current
// IntAry object.
func (ia *IntAry) PrefixToIntAry(num uint8) {
	ia.intAry = append([]uint8{num}, ia.intAry ...)
	ia.intAryLen  = len(ia.intAry)
	ia.integerLen = ia.intAryLen - ia.precision
	if num > 0 {
		ia.isZeroValue = false
	}
}

// ResetFromBackUp - Retrieves data from the
// last saved backup and populates the current
// intAry object.
func (ia *IntAry) ResetFromBackUp() {
	ia.BackUp.SetInternalFlags()

	ia.intAry = make([]uint8, ia.BackUp.intAryLen)
	for i := 0; i < ia.BackUp.intAryLen; i++ {
		ia.intAry[i] = ia.BackUp.intAry[i]
	}

	ia.intAryLen = ia.BackUp.intAryLen
	ia.firstDigitIdx = ia.BackUp.firstDigitIdx
	ia.lastDigitIdx = ia.BackUp.lastDigitIdx
	ia.isZeroValue = ia.BackUp.isZeroValue
	ia.precision = ia.BackUp.precision
	ia.signVal = ia.BackUp.signVal
	ia.decimalSeparator = ia.BackUp.decimalSeparator
	if ia.decimalSeparator == 0 {
		ia.decimalSeparator = '.'
	}

}

// RoundToPrecision - Rounds the value of the intAry to a precision
// specified by the 'roundToPrecision' parameter.
func (ia *IntAry) RoundToPrecision(roundToPrecision int) error {

	if roundToPrecision < 0 {
		fmt.Errorf("RoundToPrecision() - Error: roundToPrecision is less than ZERO! roundToPrecision= '%v'", roundToPrecision)
	}

	if ia.precision == 0 {
		return nil
	}

	err := ia.IsIntAryValid("RoundToPrecision() - ")

	if err != nil {
		return err
	}

	if ia.isZeroValue {
		ia.SetIntAryToZero(roundToPrecision)
		return nil
	}

	if roundToPrecision == ia.precision {
		return nil
	}

	if roundToPrecision > ia.precision {
		deltaPrecision := roundToPrecision - ia.precision

		for i := 0; i < deltaPrecision; i++ {
			ia.intAry = append(ia.intAry, 0)
		}

		ia.intAryLen = len(ia.intAry)
		ia.precision = roundToPrecision
		return nil
	}

	// roundToPrecision must be < ia.precision

	intLen := ia.intAryLen - ia.precision
	newIntAryLen := intLen + roundToPrecision
	fracIdx := intLen

	fracRoundIdx := fracIdx + roundToPrecision

	t := make([]uint8, ia.intAryLen+1)
	n1 := uint8(0)
	n2 := uint8(0)
	carry := uint8(0)

	for i := fracRoundIdx; i >= 0; i-- {

		n1 = ia.intAry[i]

		if i == fracRoundIdx {
			n2 = n1 + 5
		} else {
			n2 = n1 + carry
		}

		if n2 > 9 {
			carry = 1
			n2 = n2 - 10
		} else {
			carry = 0
		}

		t[i+1] = n2
	}

	ia.intAry = []uint8{}

	if carry > 0 {
		t[0] = carry
		ia.intAry = t[0 : newIntAryLen+1]
	} else {
		ia.intAry = t[1 : newIntAryLen+1]
	}
	ia.precision = roundToPrecision
	ia.intAryLen = len(ia.intAry)
	return nil
}

// SetAbsoluteValueThis - Converts the current
// value of this intAry object to its
// absolute value.
func (ia *IntAry) SetAbsoluteValueThis() {
	ia.SetSign(1)
}

// SetCurrencySymbol is used to set the value of the currency
// separator character. The currency separator character is used
// to designate currency when formatting the value of this IntAry
// is expressed as a string. In the US, the currency is the dollar
// character ('$'). In the following example, the '$' character
// designates a currency value. Example - '$1,000,000,000.00' .
//
// The SetCurrencySymbol method can be used to change the currency
// separator character in accordance with the customs of countries other
// than the US. For example, in England use of the pound sterling currency
// symbol ('£') would be appropriate.
//
// To examine the current setting for currency symbol, see
// the method GetCurrencySymbol().
//
// Note the default thousands separator character is the comma (',').
func (ia *IntAry) SetCurrencySymbol(currencySymbol rune) {
	ia.currencySymbol = currencySymbol
}

// SetDecimalSeparator - sets the decimal separator character
// (usually a decimal point) which is used to separate
// integer digits from fractional digits when the value
// of this IntAry is expressed as a string.
//
// For US usage, the '.' in the following example is
// the decimal separator. Example: 123.456
//
// SetDecimalSeparator can be used to change the default
// decimal separator to a value other than '.'
//
// To examine the current setting for decimal separator, see
// method GetDecimalSeparator().
//
// Note: The default decimal separator character is '.'
func (ia *IntAry) SetDecimalSeparator(decimalSeparator rune) {
	ia.decimalSeparator = decimalSeparator
}

// SetEqualArrayLengths - Compares an intAry object
// to the current intAry and ensures that the lengths
// of both IntArrays are equal.
func (ia *IntAry) SetEqualArrayLengths(iAry2 *IntAry) {
	iAry2.SetInternalFlags()
	ia.SetInternalFlags()

	iaIntLen := ia.intAryLen - ia.precision
	iAry2IntLen := iAry2.intAryLen - iAry2.precision

	if iaIntLen > iAry2IntLen {
		iAry2.AddArrayLengthLeft(iaIntLen - iAry2IntLen)
	}

	if iAry2IntLen > iaIntLen {
		ia.AddArrayLengthLeft(iAry2IntLen - iaIntLen)
	}

	if ia.precision > iAry2.precision {
		iAry2.AddArrayLengthRight(ia.precision - iAry2.precision)
		iAry2.precision = ia.precision
	}

	if iAry2.precision > ia.precision {
		ia.AddArrayLengthRight(iAry2.precision - ia.precision)
		ia.precision = iAry2.precision
	}

	return
}

// SetIntAryWithInt - Sets the value of the current intAry object
// to that of the input parameter 'intDigits', an integer of type 'int'.
//
// Input parameter 'precision' to indicate the number of
// digits to the right of the decimal place.
//
// The numeric sign (plus or minus) of the resulting intAry value
// is determined by the sign of input parameter,'intDigits'.
// Example:
//  intDigits     precision     	    result
//  946254  			   3							   946.254
//  946254				   0							   946254
//  -946254  			   3					      -946.254
//  -946254				   0						    -946254
//
func (ia *IntAry) SetIntAryWithInt(intDigits int, precision uint) error {
	quotient := 0
	mod := 0

	ia.intAry = []uint8{}
	ia.intAryLen = 0
	ia.precision = int(precision)
	ia.signVal = 1

	if intDigits < 0 {
		intDigits = intDigits * -1
		ia.signVal = -1
	}

	if intDigits == 0 {
		ia.SetIntAryToZero(int(precision))
		return nil
	}

	for true {

		if intDigits == 0 {
			break
		}

		quotient = intDigits / 10

		mod = intDigits - (quotient * 10)

		ia.intAry = append(ia.intAry, uint8(mod))
		ia.intAryLen++

		intDigits = quotient

	}

	n1 := uint8(0)
	lastIdx := ia.intAryLen - 1
	totalLen := ia.intAryLen / 2
	for i := 0; i < totalLen; i++ {
		n1 = ia.intAry[i]
		ia.intAry[i] = ia.intAry[lastIdx]
		ia.intAry[lastIdx] = n1
		lastIdx--
	}

	ia.SetInternalFlags()

	return nil
}

// SetIntAryWithInt32 - Sets the value of the current intAry object
// to that of the input parameter 'intDigits', a 32-bit integer.
//
// Input parameter 'precision' to indicate the number of
// digits to the right of the decimal place.
//
// The numeric sign (plus or minus) of the resulting intAry value
// is determined by the sign of input parameter 'intDigits'.
// Example:
//  intDigits     precision     	    result
//  946254  			   3							   946.254
//  946254				   0							   946254
//  -946254  			   3					      -946.254
//  -946254				   0						    -946254
//
func (ia *IntAry) SetIntAryWithInt32(intDigits int32, precision uint) error {
	tenI32 := int32(10)
	quotient := int32(0)
	mod := int32(0)

	ia.intAry = []uint8{}
	ia.intAryLen = 0
	ia.precision = int(precision)
	ia.signVal = 1

	if intDigits < 0 {
		intDigits = intDigits * -1
		ia.signVal = -1
	}

	if intDigits == 0 {
		ia.SetIntAryToZero(int(precision))
		return nil
	}

	for true {

		if intDigits == 0 {
			break
		}

		quotient = intDigits / tenI32

		mod = intDigits - (quotient * tenI32)

		ia.intAry = append(ia.intAry, uint8(mod))
		ia.intAryLen++

		intDigits = quotient

	}

	n1 := uint8(0)
	lastIdx := ia.intAryLen - 1
	totalLen := ia.intAryLen / 2
	for i := 0; i < totalLen; i++ {
		n1 = ia.intAry[i]
		ia.intAry[i] = ia.intAry[lastIdx]
		ia.intAry[lastIdx] = n1
		lastIdx--
	}

	ia.SetInternalFlags()

	return nil
}

// SetIntAryWithInt64 - Sets the value of the current intAry
// object to that of the input parameter 'intDigits', a 64-bit
// integer.
//
// Note: Input parameter 'precision' to indicate the number of
// digits to the right of the decimal place.
//
// The numeric sign (plus or minus) of the resulting intAry value
// is determined by the sign of input parameter 'intDigits'.
//
// Example:
//  intDigits     precision     	    result
//  946254  			   3							   946.254
//  946254				   0							   946254
//  -946254  			   3					      -946.254
//  -946254				   0						    -946254
//
func (ia *IntAry) SetIntAryWithInt64(intDigits int64, precision uint) error {

	quotient := int64(0)
	mod := int64(0)
	i64Ten := int64(10)

	ia.intAry = []uint8{}
	ia.intAryLen = 0
	ia.precision = int(precision)
	ia.signVal = 1

	if intDigits < 0 {
		intDigits = intDigits * int64(-1)
		ia.signVal = -1
	}

	if intDigits == 0 {
		ia.SetIntAryToZero(int(precision))
		return nil
	}

	for true {

		if intDigits == 0 {
			break
		}

		quotient = intDigits / i64Ten

		mod = intDigits - (quotient * i64Ten)

		ia.intAry = append(ia.intAry, uint8(mod))
		ia.intAryLen++

		intDigits = quotient

	}

	n1 := uint8(0)
	lastIdx := ia.intAryLen - 1
	totalLen := ia.intAryLen / 2
	for i := 0; i < totalLen; i++ {
		n1 = ia.intAry[i]
		ia.intAry[i] = ia.intAry[lastIdx]
		ia.intAry[lastIdx] = n1
		lastIdx--
	}

	ia.SetInternalFlags()

	return nil
}

// SetIntAryWithUint64 - Sets the value of the current intAry
// object to that of the input parameter 'intDigits', a 64-bit
// unsigned integer.
//
// Note: Input parameter 'precision' to indicate the number of
// digits to the right of the decimal place.
//
// Input parameter, 'signVal' must be set to one of two values: -1 or +1.
// 'signVal' determines the numeric sign of the resulting intAry value,
// either plus or minus.
//
// Example:
//  intDigits  precision  signVal		result
//  946254  			3					1				946.254
//  946254				0					1				946254
//  946254  			3					-1			-946.254
//  946254				0					-1			-946254
//
//
func (ia *IntAry) SetIntAryWithUint64(intDigits uint64, precision uint, signVal int) error {

	if signVal != 1 && signVal != -1 {
		return fmt.Errorf("ERROR - Input parameter must be equal to +1 or -1. Input signVal= %v", signVal)
	}

	ia.signVal = signVal

	if intDigits == 0 {
		ia.SetIntAryToZero(int(precision))
		return nil
	}

	ia.precision = int(precision)

	quotient := uint64(0)
	mod := uint64(0)
	ten := uint64(10)

	ia.intAry = []uint8{}
	ia.intAryLen = 0
	for true {

		if intDigits == 0 {
			break
		}

		quotient = intDigits / ten
		mod = intDigits - (quotient * ten)

		ia.intAry = append(ia.intAry, uint8(mod))
		ia.intAryLen++

		intDigits = quotient

	}

	n1 := uint8(0)
	lastIdx := ia.intAryLen - 1
	totalLen := ia.intAryLen / 2
	for i := 0; i < totalLen; i++ {
		n1 = ia.intAry[i]
		ia.intAry[i] = ia.intAry[lastIdx]
		ia.intAry[lastIdx] = n1
		lastIdx--
	}

	ia.SetInternalFlags()

	return nil
}

// SetIntAryWithBigInt - Sets the current value of the intAry to the value
// of input parameter 'intDigits'. The sign value (plus or minus) is taken
// from the input parameter, 'intDigits'. The precision or number of digits
// to the right of the decimal point, is determined by the input parameter,
// 'precision'.
//
// Example:
//  intDigits     precision     	    result
//  946254  			   3							   946.254
//  946254				   0							   946254
//  -946254  			   3					      -946.254
//  -946254				   0						    -946254
//
func (ia *IntAry) SetIntAryWithBigInt(intDigits *big.Int, precision uint) error {

	bigZero := big.NewInt(0)
	quotient := big.NewInt(0)
	mod := big.NewInt(0)
	big10 := big.NewInt(10)
	modx := big.NewInt(0)

	compare := bigZero.Cmp(intDigits)

	ia.intAry = []uint8{}
	ia.intAryLen = 0
	ia.precision = int(precision)
	ia.signVal = 1

	if compare == 1 {
		bigMinus1 := big.NewInt(0).SetInt64(int64(-1))
		intDigits = big.NewInt(0).Mul(intDigits, bigMinus1)
		ia.signVal = -1
	}

	if compare == 0 {
		ia.SetIntAryToZero(int(precision))
		return nil
	}

	for true {

		compare := bigZero.Cmp(intDigits)

		if compare == 0 {
			break
		}

		quotient, mod = big.NewInt(0).DivMod(intDigits, big10, modx)

		ia.intAry = append(ia.intAry, uint8(mod.Int64()))
		ia.intAryLen++

		intDigits.Set(quotient)

	}

	n1 := uint8(0)
	lastIdx := ia.intAryLen - 1
	totalLen := ia.intAryLen / 2
	for i := 0; i < totalLen; i++ {
		n1 = ia.intAry[i]
		ia.intAry[i] = ia.intAry[lastIdx]
		ia.intAry[lastIdx] = n1
		lastIdx--
	}

	ia.SetInternalFlags()

	return nil
}

// SetIntAryWithFloat32 - Sets the current value of the intAry based on the
// input parameter, 'floatNum'. 'floatNum is of type float32, a 32-bit
// floating point number.
//
// Input parameter 'precision' must be set to a number greater than or equal
// to zero.  It may also be set to a value of -1 which causes the number
// to be formatted to the smallest number of digits to right of the decimal
// point.
//
// Usage:
// num := float32(123.456000)
// precision := 3 - signals that only the decimals 456 will included as input
// err := ia.SetIntAryWithFloat32(num, precision)
//
// The value resulting from the usage example above would be '123.456'.
// If the value of 'precision' was set to -1, the resulting value would
// also be '123.456'
//
// Note: if precision is a positive value, and less than the current number
// of digits to the right of the decimal place, rounding may occur.
//
func (ia *IntAry) SetIntAryWithFloat32(floatNum float32, precision int) error {

	if precision < -1 {
		return fmt.Errorf("SetIntAryWithFloat32() - Input Parameter 'precision' invalid. 'precision' must be greater than or equal to -1. precison= %v", precision)
	}

	numStr := strconv.FormatFloat(float64(floatNum), 'f', -1, 32)

	err := ia.SetIntAryWithNumStr(numStr)

	if err != nil {
		return fmt.Errorf("SetIntAryWithFloat32() - Error returned from ia.SetIntAryWithNumStr(numStr). numStr= '%v'  Error= '%v'", numStr, err)
	}

	if precision > -1 {

		ia.SetPrecision(precision, true)

		if err != nil {
			return fmt.Errorf("SetIntAryWithFloat32() - Error returned from ia.SetPrecision(precision, true). precision='%v' Error= '%v'", precision, err)
		}

	}

	return nil
}

// SetIntAryWithFloat32 - Sets the current value of the intAry based on the
// input parameter, 'floatNum'. 'floatNum' is of type float64, a 64-bit
// floating point number.
//
// Input parameter 'precision' must be set to a number greater than or equal
// to zero.  It may also be set to a value of -1 which causes the number
// to be formatted to the smallest number of digits to right of the decimal
// point.
//
// Note: if precision is a positive value, and less than the current number
// of digits to the right of the decimal place, rounding	may occur.
func (ia *IntAry) SetIntAryWithFloat64(floatNum float64, precision int) error {

	if precision < -1 {
		return fmt.Errorf("SetIntAryWithFloat64() Error: Invalid input parameter 'precision'. 'precision' must be greater than or equal to -1. precision= '%v'", precision)
	}

	numStr :=  strconv.FormatFloat(floatNum, 'f', -1, 64)

	err := ia.SetIntAryWithNumStr(numStr)

	if err != nil {
		return fmt.Errorf("SetIntAryWithFloat64() - ia.SetIntAryWithNumStr(numStr) returned error. numStr='%v' Error='%v'",numStr, err)
	}

	if precision > -1 {
		err := ia.SetPrecision(precision, true)

		if err != nil {
			return fmt.Errorf("SetIntAryWithFloat64() - Error returned from ia.SetPrecision(precision, true). precision='%v' Error= '%v'", precision, err)
		}

	}


	return nil
}

// SetIntAryWithFloatBig - Sets the current value of the intAry based on the
// input parameter, 'floatNum'.  'floatNum' is of type *big.Float.
//
// Input Parameters:
//
// floatNum *big.Float - 	The floating point value to which the current IntAry object
// 											 	will be set.
//
// precision int 			 - 	The number of digits to the right of the decimal point
//												which will be used to set the value of the current IntAry
//												object
//
// 												Input parameter 'precision' must be set to a number greater
// 												than or equal to zero.  It may also be set to a value of -1
// 												which causes the number to be formatted to the smallest number
// 												of digits to right of the decimal point.
//
// 												Note: if precision is a positive value, and less than the current
// 												number of digits to the right of the decimal place, rounding may occur.
//
func (ia *IntAry) SetIntAryWithFloatBig(floatNum *big.Float, precision int) error {

	if precision < -1 {
		return fmt.Errorf("SetIntAryWithFloatBig() Error: Invalid input parameter 'precision'. 'precision' must be greater than -1. precision= '%v'", precision)
	}

	var numStr string

	if precision == -1 {
		numStr = floatNum.Text('f', precision)
	} else {
		numStr = floatNum.Text('f', precision+1)
	}

	err := ia.SetIntAryWithNumStr(numStr)

	if err != nil {
		return fmt.Errorf("SetIntAryWithFloatBig() - Error returned from ia.SetIntAryWithNumStr(numStr). numStr='%v'  Error= '%v'", numStr, err)
	}

	if precision > -1 {
		err := ia.SetPrecision(precision, true)

		if err != nil {
			return fmt.Errorf("SetIntAryWithFloatBig() - Error returned from ia.SetPrecision(precision, true). precision='%v' Error= '%v'", precision, err)
		}

	}

	return nil
}

// SetIntAryWithIntAry - Sets the value of the current intAry based on an array of
// integers i.e. []int.
//
// Input parameter 'precision' will determine the number of digits to the right
// of the decimal place.
//
// Input parameter 'signVal' must be either +1 or -1 indicating the sign of the
// number represented by the integer array. Note: If signVal is not equal to +1 or -1,
// an error is generated.
func (ia *IntAry) SetIntAryWithIntAry(iAry2 []int, precision uint, signVal int) error {

	if signVal != 1 && signVal != -1 {
		return fmt.Errorf("SetIntAryWithIntAry() - Error: signVal parameter is INVALID! signVal must be -1 or +1. signVal='%v'", signVal)
	}

	lIAry2 := len(iAry2)
	ia.intAry = make([]uint8, lIAry2)
	for i := 0; i < lIAry2; i++ {
		ia.intAry[i] = uint8(iAry2[i])
	}

	ia.intAryLen = lIAry2
	ia.precision = int(precision)
	ia.signVal = signVal
	ia.SetInternalFlags()
	return nil
}

// SetIntAryWithUint8Ary - is designed to set the value of the current IntAry
// object by passing in a pointer to an unsigned integer array ([]uint8). []uint8
// is the type of array native to the IntAry object.
//
// Input parameter 'precision' will determine the number of digits to the right
// of the decimal place.
//
// Input parameter 'signVal' must be either +1 or -1 indicating the sign of the
// number represented by the integer array. Note: If signVal is not equal to +1 or -1,
// an error is generated.
func (ia *IntAry) SetIntAryWithUint8Ary( iAry2 []uint8,precision uint, signVal int) error {

	if signVal != 1 && signVal != -1 {
		return fmt.Errorf("SetIntAryWithUint8Ary() - Error: signVal parameter is INVALID! signVal must be -1 or +1. signVal='%v'", signVal)
	}

	lIAry2 := len(iAry2)
	ia.intAry = make([]uint8, lIAry2)

	for i := 0; i < lIAry2; i++ {
		ia.intAry[i] = uint8(iAry2[i])
	}

	ia.intAryLen = lIAry2
	ia.precision = int(precision)
	ia.signVal = signVal
	ia.SetInternalFlags()

	if ia.isIntegerZeroValue && ia.integerLen > 1 {
		ia.OptimizeIntArrayLen(false)
	}

	return nil
}


func (ia *IntAry) SetIntAryWithIntAryObj(iAry2 *IntAry, copyBackup bool) error {

	err := iAry2.IsIntAryValid("SetIntAryWithIntAryObj()")

	if err != nil {
		return fmt.Errorf("SetIntAryWithIntAryObj - Input parameter iAry2 is INVALID! Error= %v", err)
	}

	ia.CopyIn(iAry2, copyBackup)

	return nil
}

// SetIntAryWithNumStr - receives a raw number string and sets the
// fields of the internal intAry structure to the appropriate
// values.
func (ia *IntAry) SetIntAryWithNumStr(str string) error {

	if len(str) == 0 {
		return errors.New("SetIntAryWithNumStr() received zero length number string")
	}

	ia.Empty()

	if ia.decimalSeparator == 0 {
		ia.decimalSeparator = '.'
	}

	if ia.thousandsSeparator == 0 {
		ia.thousandsSeparator = ','
	}

	if ia.currencySymbol == 0 {
		ia.currencySymbol = '$'
	}

	ia.signVal = 1
	baseRunes := []rune(str)
	lBaseRunes := len(baseRunes)
	isStartRunes := false
	isEndRunes := false
	isFractionalValue := false

	for i := 0; i < lBaseRunes && isEndRunes == false; i++ {

		if baseRunes[i] == '+' ||
			baseRunes[i] == ' ' ||
			baseRunes[i] == ia.thousandsSeparator ||
			baseRunes[i] == ia.currencySymbol {

			continue

		}

		if baseRunes[i] == ',' && ia.decimalSeparator != ',' {
			continue
		}

		if isStartRunes == true &&
			isEndRunes == false &&
			isFractionalValue &&
			baseRunes[i] == ia.decimalSeparator {

			continue
		}

		if baseRunes[i] == '-' &&
			isStartRunes == false && isEndRunes == false &&
			i+1 < lBaseRunes &&
			((baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9') ||
				baseRunes[i+1] == ia.decimalSeparator) {

			ia.signVal = -1
			isStartRunes = true
			continue

		} else if isEndRunes == false &&
			baseRunes[i] >= '0' && baseRunes[i] <= '9' {

			ia.intAry = append(ia.intAry, uint8(baseRunes[i]-48))
			isStartRunes = true

			if isFractionalValue {
				ia.precision++
			}

		} else if isEndRunes == false &&
			i+1 < lBaseRunes &&
			baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9' &&
			baseRunes[i] == ia.decimalSeparator {

			isFractionalValue = true
			continue

		} else if isStartRunes && !isEndRunes {

			isEndRunes = true

		}

	}

	ia.SetSignificantDigitIdxs()

	if ia.intAryLen == 0 || ia.isZeroValue {
		ia.SetIntAryToZero(ia.precision)
		return nil
	}


	// Validate intAry object
	err := ia.IsIntAryValid("SetIntAryWithNumStr() - ")

	if err != nil {
		return err
	}


	return nil
}

func (ia *IntAry) SetIntAryLength() {
	ia.intAryLen = len(ia.intAry)
}

// SetIsZeroValue - Analyzes the value
// of the intAry and sets a flag
// if the value of intAry evaluates
// to zero.
func (ia *IntAry) SetIsZeroValue() {
	ia.intAryLen = len(ia.intAry)

	ia.isZeroValue = true
	ia.isIntegerZeroValue = true

	intLen := ia.intAryLen - ia.precision

	for i := 0; i < ia.intAryLen; i++ {

		if i < intLen && ia.intAry[i] > 0 {
			ia.isIntegerZeroValue = false
		}

		if ia.intAry[i] > 0 {
			ia.isZeroValue = false
			return
		}
	}

	// ia.isZeroValue == true
	// signVal must be 1
	ia.signVal = 1
}

// SetIntAryToOne - Sets the value of the intAry object to one ('1').
func (ia *IntAry) SetIntAryToOne(precision int) error {

	if precision < 0 {
		return fmt.Errorf("SetIntAryToOne() - Error: precision is less than ZERO! precision= '%v'", precision)
	}

	ia.intAryLen = 1 + precision
	ia.precision = precision
	ia.intAry = make([]uint8, ia.intAryLen)
	ia.intAry[0] = 1
	ia.signVal = 1
	ia.isZeroValue = false
	ia.isIntegerZeroValue = false
	ia.firstDigitIdx = 0
	ia.lastDigitIdx = 0

	if ia.decimalSeparator == 0 {
		ia.decimalSeparator = '.'
	}

	return nil
}

// SetIntAryToTen - Sets the value of the intAry object to Ten ('10')
func (ia *IntAry) SetIntAryToTen(precision int) error {

	if precision < 0 {
		return fmt.Errorf("SetIntAryToOne() - Error: precision is less than ZERO! precision= '%v'", precision)
	}

	ia.intAryLen = 2 + precision
	ia.precision = precision
	ia.intAry = make([]uint8, ia.intAryLen)
	ia.intAry[0] = 1
	ia.signVal = 1
	ia.isZeroValue = false
	ia.isIntegerZeroValue = false
	ia.firstDigitIdx = 0
	ia.lastDigitIdx = 0

	if ia.decimalSeparator == 0 {
		ia.decimalSeparator = '.'
	}

	return nil
}

// SetIntAryToZero - Sets the value of the intAry object to Zero ('0').
func (ia *IntAry) SetIntAryToZero(precision int) error {

	if precision < 0 {
		return fmt.Errorf("SetIntAryToOne() - Error: precision is less than ZERO! precision= '%v'", precision)
	}

	ia.intAryLen = 1 + precision
	ia.precision = precision
	ia.intAry = make([]uint8, ia.intAryLen)
	ia.signVal = 1

	if ia.decimalSeparator == 0 {
		ia.decimalSeparator = '.'
	}

	ia.SetInternalFlags()

	return nil
}

// SetInternalFlags - Sets Array Lengths and
// test for zero values
func (ia *IntAry) SetInternalFlags() {
	ia.SetSignificantDigitIdxs()
}


// SetTruncateToPrecision - Truncates the existing
// value to precision specified by the 'precision'
// parameter. No rounding is performed.
//
// If 'precision' is zero, the Int Ary value will
// be truncated to an integer value with no
// fractional digits.
//
// If 'precision' is set to a value less than zero,
// an error will be returned.
//
// If 'precision' is greater than the existing precision,
// trailing zeros will be added

func (ia *IntAry) SetPrecision(precision int, roundResult bool) error {

	if precision < 0 {
		return fmt.Errorf("SetPrecision() - Error: 'precision' value is less than ZERO! precision= '%v'", precision)
	}

	err := ia.IsIntAryValid("SetPrecision() - ")

	if err != nil {
		return err
	}

	if ia.isZeroValue {
		ia.SetIntAryToZero(precision)
		return nil
	}

	if precision == ia.precision {
		return nil
	}

	if precision > ia.precision {
		deltaPrecision := precision - ia.precision

		for i := 0; i < deltaPrecision; i++ {
			ia.intAry = append(ia.intAry, 0)
		}

		ia.precision = precision
		ia.intAryLen = len(ia.intAry)
		return nil
	}

	// Must ia.precision > precision

	if roundResult {
		ia.RoundToPrecision(precision)
		return nil
	}

	intLen := ia.intAryLen - ia.precision
	newAryLen := intLen + precision
	ia.intAry = ia.intAry[0:newAryLen]

	ia.intAryLen = newAryLen
	ia.precision = precision

	return nil
}

// SetSign - Can be used to change the sign value
// of the current intAry value. The new sign value
// will be set according the input parameter, 'signVal'.
//
// 'signVal' has only two valid values, -1 or +1. If
// any value other than -1 or +1 is detected, an error
// will be thrown.
func (ia *IntAry) SetSign(signVal int) error {

	ia.SetInternalFlags()

	if ia.isZeroValue {
		return nil
	}

	if signVal != -1 && signVal != 1 {
		return fmt.Errorf("SetSign() - Input parameter 'signVal' is INVALID. ")
	}

	if signVal != ia.signVal {
		ia.signVal = signVal
	}

	return nil
}

// SetSignificantDigitIdxs - Finds the first
// significant digit (the first numeric digit
// greater than zero) and sets index value in
// the local field variable, 'firstDigitIdx'.
//
// In addition, this method also identifies the
// Last Significant Digit (the last non-zero value
// in the intAry) and records that index in the
// local field variable, 'lastDigitIdx'.
func (ia *IntAry) SetSignificantDigitIdxs() {

	ia.intAryLen = len(ia.intAry)

	if ia.intAryLen == ia.precision {
		ia.intAry = append([]uint8{0}, ia.intAry...)
		ia.intAryLen++
	}

	if ia.intAryLen < ia.precision {

		deltaZeros := ia.precision - ia.intAryLen + 1
		zeroAry := make([]uint8, deltaZeros)
		ia.intAry = append(zeroAry, ia.intAry...)
		ia.intAryLen += deltaZeros
	}

	ia.firstDigitIdx = -1
	ia.lastDigitIdx = -1

	ia.integerLen = 0
	ia.significantIntegerLen = 0
	ia.significantFractionLen = 0

	lastIntIdx := ia.intAryLen - ia.precision - 1
	ia.isZeroValue = true
	ia.isIntegerZeroValue = true
	ia.integerLen = ia.intAryLen - ia.precision

	for i := 0; i < ia.intAryLen; i++ {
		if ia.intAry[i] > 0 {
			ia.isZeroValue = false

			if i < ia.integerLen {
				ia.isIntegerZeroValue = false
			}
		}

		// At minimum, there should be a single
		// leading zero before the decimal point.
		// Example 0.000.
		if i == lastIntIdx && ia.intAry[i] == 0 {

			if ia.firstDigitIdx == -1 {
				ia.firstDigitIdx = i
			}

		}

		if ia.intAry[i] > 0 {

			if ia.firstDigitIdx == -1 {
				ia.firstDigitIdx = i
			}

			ia.lastDigitIdx = i
		}

	}

	ia.significantIntegerLen = ia.intAryLen - ia.precision - ia.firstDigitIdx

	if ia.lastDigitIdx >= ia.integerLen {
		ia.significantFractionLen = ia.lastDigitIdx - ia.integerLen + 1
	} else {
		ia.significantFractionLen = 0
	}
}

// SetThousandsSeparator is used to set the value of the thousands
// separator character. The thousands separator character is used
// to separate thousands when the value of this IntAry object is
// expressed as a string. In the US, the thousands separator character
// is a comma. In the following example, the ',' character separates
// thousands in the integer number: Example - '1,000,000,000'.
//
// The SetThousandsSeparator method can be used to change the thousands
// separator character in accordance with the customs of countries other
// than the US.
//
// To examine the current setting for thousands separator, see
// method GetIntAryStats().
//
// Note the default thousands separator character is the comma (',').
func (ia *IntAry) SetThousandsSeparator(thousandsSeparator rune) {
	ia.thousandsSeparator = thousandsSeparator
}

// SubtractFromThis - Subtracts the value of parameter
// 'ia2' from the current intAry object.
// Paramters:
//
// ia2 *intAry - Incoming intAry object whose value will be subtracted
// 								from this current intAry value.
//
// convertToNumStr - boolean value determines whether the current intAry
//                   object will convert the intAry value to a number string.
//                   Set this parameter to 'false' if this method is called
//                   multiple times in order to improve performance.
func (ia *IntAry) SubtractFromThis(ia2 *IntAry) error {

	ia.SetEqualArrayLengths(ia2)

	if ia.isZeroValue && ia2.isZeroValue {
		ia.SetIntAryToZero(ia.precision)
		return nil
	}

	compare := ia.CompareAbsoluteValues(ia2)
	isZeroResult := false

	// Largest Value in now in N1 slot
	newSignVal := ia.signVal
	doAdd := false
	doReverseNums := false

	if compare == 1 {
		// compare == + 1
		// Absolute Value: N1 > N2

		if ia.signVal == 1 && ia2.signVal == 1 {
			doAdd = false
			newSignVal = 1
		} else if ia.signVal == -1 && ia2.signVal == 1 {
			doAdd = true
			newSignVal = -1
		} else if ia.signVal == -1 && ia2.signVal == -1 {
			doAdd = false
			newSignVal = -1
		} else {
			// Must Be ia.signVal == 1 && ia2.signVal == -1
			doAdd = true
			newSignVal = 1
		}

	} else if compare == -1 {
		// Absolute Values: N2 > N1
		if ia.signVal == 1 && ia2.signVal == 1 {
			doAdd = false
			doReverseNums = true
			newSignVal = -1
		} else if ia.signVal == -1 && ia2.signVal == 1 {
			doAdd = true
			newSignVal = -1
		} else if ia.signVal == -1 && ia2.signVal == -1 {
			doAdd = false
			doReverseNums = true
			newSignVal = 1
		} else {
			// Must Be ia.signVal == 1 && ia2.signVal == -1
			doAdd = true
			newSignVal = 1
		}

	} else {
		// Must be compare == 0
		// Absolute Values: N1==N2
		if ia.signVal == 1 && ia2.signVal == 1 {
			doAdd = false
			newSignVal = 1
			isZeroResult = true
		} else if ia.signVal == -1 && ia2.signVal == 1 {
			doAdd = true
			newSignVal = -1
		} else if ia.signVal == -1 && ia2.signVal == -1 {
			doAdd = false
			newSignVal = 1
			isZeroResult = true
		} else {
			// Must Be ia.signVal == 1 && ia2.signVal == -1
			doAdd = true
			newSignVal = 1
		}

	}

	return ia.addToSubtractFromThis(ia2, newSignVal, doAdd, isZeroResult, doReverseNums)
}

// SubtractMultipleFromThis - This method will subtract multiple intAry values from the
// current intAry value. There are two input parameters:
//
// convertToNumStr bool - If true the result will be converted to a number string after
//  											the final subtraction operation.
//
// iaMany ...*intAry - An unlimited series of pointers to intAry objects which will be
// 										 subtracted from the current intAry Value.
//
func (ia *IntAry) SubtractMultipleFromThis(iaMany ...*IntAry) error {
	var err error

	for _, iAry := range iaMany {

		err = ia.SubtractFromThis(iAry)

		if err != nil {
			return fmt.Errorf("SubtractMultipleFromThis() - Received error from ia.SubtractFromThis(iAry, false). Error= %v", err)
		}
	}

	return nil
}
