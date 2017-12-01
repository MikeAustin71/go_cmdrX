package common

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

/*
	The source code repository for numstrdto.go is located at:
			https://github.com/MikeAustin71/numstrdto.git
 */

// Source Currency Info
// https://gist.github.com/bzerangue/5484121
// http://symbologic.info/currency.htm
// http://www.xe.com/symbols.php

var NumStrCurrencySymbols = []rune{
	'\U00000024', // Australia Dollar 								 0
	'\U00008236', // Brazil Real											 1
	'\U00000024', // Canada Dollar 										 2
	'\U000000a5', // China Yuan												 3
	'\U00000024', // Colombia Peso										 4
	'\U0004b10d', // Czech Republic Koruna						 5
	'\U000000a3', // Egypt Pound											 6
	'\U000020ac', // Euro    													 7
	'\U00070116', // Hungary Forint										 8
	'\U00107114', // Iceland Krona										 9
	'\U00082112', // Indonesia Rupiah									10
	'\U000020aa', // Israel Shekel  									11
	'\U000000a5', // Japan Yen  											12
	'\U000020a9', // Korea Won  											13
	'\U0000524d', // Malaysia Ringgit									14
	'\U00000024', // Mexico Peso  										15
	'\U00006b72', // Norway Krone											16
	'\U00000192', // Netherlands Antilles Guilder			17
	'\U000020a8', // Pakistan Rupee 									18
	'\U000020bd', // Russian Ruble  									19
	'\U0000fdfc', // Saudi Arabia Riyal 							20
	'\U00000082', // South Africa Rand								21
	'\U00006b72', // Sweden Krona											22
	'\U000020a3', // Switzerland Franc								23
	'\U00000024', // Taiwan New Dollar								24
	'\U000020ba', // TURKISH LIRA											25
	'\U00066115', // Venezuela Bolivar								26
	'\U00008363', // Viet Nam Dong										27
	'\U00000024', // United States Dollar  						28
	'\U000000a3', // United Kingdom Pound (£)					29
	'\U000020a3', // French Franc  						        30
	'\U000020a4', // Italy Lira  						          31
	'\U000020bf', // Bitcoin  						            32
	'\U000000a2'} // United States Cent		            33

type NumStrDto struct {
	IsValid            bool
	SignVal            int
	AbsAllNumRunes     []rune
	AbsIntRunes        []rune
	AbsFracRunes       []rune
	Precision          uint
	IsFractionalValue  bool
	HasNumericDigits   bool
	ThousandsSeparator rune
	DecimalSeparator   rune
	CurrencySymbol     rune
	NumStrIn           string
	NumStrOut          string
}

// New - Used to create NumStrDto types.
// This message initializes the NumStrDto
// fields. This method will return the newly
// create type (not a pointer to the type).
// Example:
// n := NumStrDto{}.New()
// n2, err := n.ParseNumStr("123.456")
//
// Compare this method of object creation
// with that shown in the NewPtr() method.
func (nDto NumStrDto) New() NumStrDto {
	n := NumStrDto{}
	n.Empty()

	return n
}

// NewPtr - Used to create and initialize
// NumStrDto fields.  This method returns
// a pointer to the newly created NumStrDto
// type. As such, this method may be used
// to streamline the initialization process.
// Example:
// n, err := NumStrDto{}.NewPtr().ParseNumStr("123.456")
func (nDto NumStrDto) NewPtr() *NumStrDto {
	n := NumStrDto{}
	n.Empty()
	return &n
}

// AddNumStrs - Adds the values represented by two NumStrDto objects and
// returns the result as an NumStrDto.
func (nDto *NumStrDto) AddNumStrs(n1Dto NumStrDto, n2Dto NumStrDto) (NumStrDto, error) {

	n1DtoSetup, n2DtoSetup, _, _, err := nDto.FormatForMathOps(n1Dto, n2Dto)

	if err != nil {
		return NumStrDto{}, fmt.Errorf("AddNumStrs() - Error returned from nDto.FormatForMathOps(n1Dto, n2Dto). Error= %v", err)
	}

	newSignVal := n1DtoSetup.SignVal

	if n1DtoSetup.SignVal != n2DtoSetup.SignVal {
		n1DtoSetup.SetSignValue(1)
		n2DtoSetup.SetSignValue(1)
		nDtoOut, err := nDto.SubtractNumStrs(n1DtoSetup, n2DtoSetup)

		if err != nil {
			return NumStrDto{}, fmt.Errorf("AddNumStrs() - Error returned from nDto.SubtractNumStrs(n1DtoSetup, n2DtoSetup). Error= %v", err)
		}

		if nDto.IsNumStrZeroValue(&nDtoOut) {
			newSignVal = 1
		}

		nDtoOut.SetSignValue(newSignVal)

		return nDtoOut, nil
	}

	precision := n1DtoSetup.Precision
	lenN1AllRunes := len(n1DtoSetup.AbsAllNumRunes)

	n3IntAry := make([]int, lenN1AllRunes+1)
	carry := 0
	n1 := 0
	n2 := 0
	n3 := 0

	for j := lenN1AllRunes - 1; j >= 0; j-- {

		n1 = int(n1DtoSetup.AbsAllNumRunes[j]) - 48
		n2 = int(n2DtoSetup.AbsAllNumRunes[j]) - 48

		n3 = n1 + n2 + carry

		carry = 0

		if n3 > 9 {
			n3 = n3 - 10
			carry = 1
		}

		n3IntAry[j+1] = n3

	}

	if carry > 0 {
		n3IntAry[0] = carry
	}

	return nDto.FindIntArraySignificantDigitLimits(n3IntAry, precision, newSignVal)

}

func (nDto *NumStrDto) CompareSignedVals(n1Dto, n2Dto *NumStrDto) int {

	cmpAbs := nDto.CompareAbsoluteVals(n1Dto, n2Dto)

	if cmpAbs == 0 {

		if n1Dto.SignVal == n2Dto.SignVal {
			return 0
		} else {
			// n1Dto.SignVal != n2Dto.SignVal
			if n1Dto.SignVal == 1 {
				return 1
			}

			// n2Dto.SignVal must == 1
			return -1

		}

	}

	if cmpAbs == 1 {

		if n1Dto.SignVal == n2Dto.SignVal {

			if n1Dto.SignVal == 1 {
				return 1
			}

			// must be n1Dto.SignVal == n2Dto.SignVal && n1Dto.SignVal == -1

			return -1

		}

		// must be n1Dto.SignVal != n2Dto.SignVal
		if n1Dto.SignVal == 1 {
			return 1
		} else {
			// must be n2Dto.SignVal == 1
			return -1
		}
	}

	// cmpAbs == -1

	if n2Dto.SignVal == n1Dto.SignVal {

		if n2Dto.SignVal == 1 {
			// n1Dto.SignVal && n2Dto.SignVal must equal 1
			return -1
		} else {
			// n1Dto.SignVal && n2Dto.SignVal must equal -1
			return 1
		}

	}

	// must be n2Dto.SignVal != n1Dto.SignVal

	if n2Dto.SignVal == -1 {
		return 1
	}

	// must be n2Dto.SignVal == 1
	return -1
}

// CompareAbsoluteVals - compares the absolute numeric values
// of two NumStrDto objects. The signs (+ or -) of the two
// compared numeric values are ignored. Only the absolute
// numeric values are compared.
// Return Values:
// -1 = n1Dto is less than n2Dto
//  0 = n1Dto is equal to n2Dto
//  1 = n1Dto is greater than n2Dto
//
// Examples:
// 	n1        			n2           	Result
// 	-9691.23				91.245				 	 1
//  9691.23					91.245					 1
//  -5							82							-1
//   5							 5							 0
//
func (nDto *NumStrDto) CompareAbsoluteVals(n1Dto, n2Dto *NumStrDto) int {
	lenN1IntRunes := len(n1Dto.AbsIntRunes)
	lenN2IntRunes := len(n2Dto.AbsIntRunes)

	isN1Zero := nDto.IsNumStrZeroValue(n1Dto)
	isN2Zero := nDto.IsNumStrZeroValue(n2Dto)

	if !isN1Zero && isN2Zero {
		return 1
	}

	if isN1Zero && !isN2Zero {
		return -1
	}

	if isN1Zero && isN2Zero {
		return 0
	}

	if lenN1IntRunes > lenN2IntRunes {
		return 1
	}

	if lenN1IntRunes < lenN2IntRunes {
		return -1
	}

	// lenN1IntRunes Must Be Equal to lenN2IntRunes

	for i := 0; i < lenN1IntRunes; i++ {
		n1 := n1Dto.AbsIntRunes[i] - 48
		n2 := n2Dto.AbsIntRunes[i] - 48

		if n1 > n2 {
			return 1
		}

		if n1 < n2 {
			return -1
		}
	}

	// All the integers are equal
	lenN1FracRunes := len(n1Dto.AbsFracRunes)
	lenN2FracRunes := len(n2Dto.AbsFracRunes)

	lenFracRunesToTest := lenN1FracRunes

	if lenN2FracRunes < lenN1FracRunes {
		lenFracRunesToTest = lenN2FracRunes
	}

	for j := 0; j < lenFracRunesToTest; j++ {
		n1 := n1Dto.AbsFracRunes[j] - 48
		n2 := n2Dto.AbsFracRunes[j] - 48
		if n1 > n2 {
			return 1
		}

		if n1 < n2 {
			return -1
		}

	}

	if lenN1FracRunes > lenN2FracRunes {
		return 1
	}

	if lenN1FracRunes < lenN2FracRunes {
		return -1
	}

	return 0
}

// CopyOut - Creates a copy of the current
// NumStrDto fields and returns a completely
// new instance of NumStrDto
func (nDto *NumStrDto) CopyOut() NumStrDto {
	nOut := NumStrDto{}

	nOut.SignVal = nDto.SignVal
	nOut.AbsAllNumRunes = nDto.AbsAllNumRunes
	nOut.AbsIntRunes = nDto.AbsIntRunes
	nOut.AbsFracRunes = nDto.AbsFracRunes
	nOut.Precision = nDto.Precision
	nOut.IsFractionalValue = nDto.IsFractionalValue
	nOut.HasNumericDigits = nDto.HasNumericDigits
	nOut.NumStrIn = nDto.NumStrIn
	nOut.NumStrOut = nDto.NumStrOut
	nOut.ThousandsSeparator = nDto.ThousandsSeparator
	nOut.DecimalSeparator = nDto.DecimalSeparator
	nOut.CurrencySymbol = nDto.CurrencySymbol
	nOut.IsValid = nDto.IsValid

	return nOut
}

// CopyIn - Receives an incoming NumStrDto object
// and copies the information to the current NumStrDto
// data fields.
func (nDto *NumStrDto) CopyIn(nInDto NumStrDto) {

	nDto.Empty()

	nDto.SignVal = nInDto.SignVal
	nDto.AbsAllNumRunes = nInDto.AbsAllNumRunes
	nDto.AbsIntRunes = nInDto.AbsIntRunes
	nDto.AbsFracRunes = nInDto.AbsFracRunes
	nDto.Precision = nInDto.Precision
	nDto.IsFractionalValue = nInDto.IsFractionalValue
	nDto.HasNumericDigits = nInDto.HasNumericDigits
	nDto.NumStrIn = nInDto.NumStrIn
	nDto.NumStrOut = nInDto.NumStrOut
	nDto.ThousandsSeparator = nInDto.ThousandsSeparator
	nDto.DecimalSeparator = nInDto.DecimalSeparator
	nDto.CurrencySymbol = nInDto.CurrencySymbol
	nDto.IsValid = nInDto.IsValid

}

// Empty - Sets all the fields in the NumStrDto
// to their initial or zero state.
func (nDto *NumStrDto) Empty() {
	nDto.IsValid = true
	nDto.SignVal = 0
	nDto.AbsAllNumRunes = []rune{}
	nDto.AbsIntRunes = []rune{}
	nDto.AbsFracRunes = []rune{}
	nDto.Precision = 0
	nDto.IsFractionalValue = false
	nDto.HasNumericDigits = false
	nDto.NumStrIn = ""
	nDto.NumStrOut = ""

	if nDto.ThousandsSeparator == 0 {
		nDto.ThousandsSeparator = ','
	}

	if nDto.DecimalSeparator == 0 {
		nDto.DecimalSeparator = '.'
	}

	if nDto.CurrencySymbol == 0 {
		nDto.CurrencySymbol = '$'
	}


}

// FindIntArraySignificantDigitLimits - Receives an array of integers and converts them
// to a number string conisting of significant digits. Leading and trailing zeros are
// eliminated. See Method: FindNumStrSignificantDigitLimits()
func (nDto *NumStrDto) FindIntArraySignificantDigitLimits(intArray []int, precision uint, signVal int) (NumStrDto, error) {

	lenIntArray := len(intArray)

	var absNumStr []rune

	for i := 0; i < lenIntArray; i++ {
		absNumStr = append(absNumStr, rune(intArray[i]+48))
	}

	return nDto.FindNumStrSignificantDigitLimits(absNumStr, precision, signVal)
}

// FindSignificantDigitLimits - Analyzes an array of characters which constitute a number string
// are returns the significant digits.
// Example:
// absAllRunes  precision signVal			Result
// 001236700			4					1					123.67
// 000006700			4					1					  0.67
// 001230000			4					1					123.0
func (nDto *NumStrDto) FindNumStrSignificantDigitLimits(absAllRunes []rune, precision uint, signVal int) (NumStrDto, error) {
	iPrecision := int(precision)
	firstIntIdx := -1
	lastIntIdx := -1
	lastFracIdx := -1

	isFractional := false

	if iPrecision > 0 {
		isFractional = true
	}

	lenAbsAllRunes := len(absAllRunes)
	lenAbsFracRunes := iPrecision
	lenAbsIntRunes := lenAbsAllRunes - lenAbsFracRunes

	if lenAbsAllRunes < 1 {
		return NumStrDto{}, errors.New("FindSignificantDigitLimits() - Error: absAllRunes has ZERO length!")
	}

	for i := 0; i < lenAbsAllRunes; i++ {

		if i < lenAbsIntRunes {

			if firstIntIdx == -1 && absAllRunes[i] > '0' && absAllRunes[i] <= '9' {
				firstIntIdx = i
			}

			lastIntIdx = i
		}

		if isFractional && i >= lenAbsIntRunes && absAllRunes[i] > '0' && absAllRunes[i] <= '9' {
			lastFracIdx = i
		}

	}

	if firstIntIdx == -1 {
		firstIntIdx = lastIntIdx
	}

	if isFractional && lastFracIdx == -1 {
		lastFracIdx = lenAbsIntRunes
	}

	numStrOut := ""

	if signVal < 0 {
		numStrOut = "-"
	}

	numStrOut += string(absAllRunes[firstIntIdx : lastIntIdx+1])
	if isFractional {
		numStrOut += string(nDto.DecimalSeparator)
		numStrOut += string(absAllRunes[lastIntIdx+1 : lastFracIdx+1])
	}

	nOutDto, err := nDto.ParseNumStr(numStrOut)

	if err != nil {
		return NumStrDto{}, fmt.Errorf("FindSignificantDigitLimits() - Error retuned from nDto.ParseNumStr(numStrOut). numStrOut= '%v' Error= %v", numStrOut, err)
	}

	return nOutDto, nil
}

// FormatForMathOps - receives two NumStrDto objects and converts their strings
// such that both have the same number of integer and fractional digits. This will
// facilitate the performance of string based math operations such as addition and
// subtraction.
//
// The return values represent the formatted NumStrDto objects. The first NumStrDto
// returned always contains the larger absolute value. The second NumStrDto always
// contains the absolute numeric value which is less than or equal to the first
// NumStrDto object returned.
//
// The third parameter returned by this method is an int which will always be set to
// 1 or 0. 1 indicates that the absolute value of the first NumStrDto returned by
// this method is greater than the second NumStrDto returned by this method. If
// the int value returned is zero, it signals that the absolute values
// (not the signed values) of both returned NumStrDto objects are equal.
func (nDto *NumStrDto) FormatForMathOps(n1Dto, n2Dto NumStrDto) (n1DtoOut NumStrDto, n2DtoOut NumStrDto, compare int, isOrderReversed bool, err error) {

	lenN1AllRunes := 0
	lenN1IntRunes := 0
	lenN1FracRunes := 0
	lenN2AllRunes := 0
	lenN2IntRunes := 0
	lenN2FracRunes := 0

	err = nDto.IsNumStrDtoValid(&n1Dto, "FormatForMathOps() - ")

	if err != nil {
		return NumStrDto{}, NumStrDto{}, 0, false, err
	}

	err = nDto.IsNumStrDtoValid(&n2Dto, "FormatForMathOps() - ")

	if err != nil {
		return NumStrDto{}, NumStrDto{}, 0, false, err
	}

	compare = nDto.CompareAbsoluteVals(&n1Dto, &n2Dto)

	if compare == 1 {
		n1DtoOut = n1Dto.CopyOut()
		n2DtoOut = n2Dto.CopyOut()
	} else if compare == -1 {
		n1DtoOut = n2Dto.CopyOut()
		n2DtoOut = n1Dto.CopyOut()
		isOrderReversed = true
		compare = 1
	} else {
		// compare must be zero
		n1DtoOut = n1Dto.CopyOut()
		n2DtoOut = n2Dto.CopyOut()
	}

	if n1DtoOut.Precision > n2DtoOut.Precision {

		deltaPrecision := n1DtoOut.Precision - n2DtoOut.Precision

		for i := uint(0); i < deltaPrecision; i++ {
			n2DtoOut.AbsAllNumRunes = append(n2DtoOut.AbsAllNumRunes, '0')
			n2DtoOut.AbsFracRunes = append(n2DtoOut.AbsFracRunes, '0')
		}

		lenN2AllRunes = len(n2DtoOut.AbsAllNumRunes)
		lenN2IntRunes = len(n2DtoOut.AbsIntRunes)
		lenN2FracRunes = len(n2DtoOut.AbsFracRunes)

		n2DtoOut.Precision = n1DtoOut.Precision
		err = n2DtoOut.ResetNumStrOut()

		if err != nil {
			return NumStrDto{}, NumStrDto{}, 0, false, err
		}

		lenN1AllRunes = len(n1DtoOut.AbsAllNumRunes)
		lenN1IntRunes = len(n1DtoOut.AbsIntRunes)
		lenN1FracRunes = len(n1DtoOut.AbsFracRunes)

	} else if n1DtoOut.Precision < n2DtoOut.Precision {

		deltaPrecision := n2DtoOut.Precision - n1DtoOut.Precision

		for i := uint(0); i < deltaPrecision; i++ {
			n1DtoOut.AbsAllNumRunes = append(n1DtoOut.AbsAllNumRunes, '0')
			n1DtoOut.AbsFracRunes = append(n1DtoOut.AbsFracRunes, '0')
		}

		lenN1AllRunes = len(n1DtoOut.AbsAllNumRunes)
		lenN1IntRunes = len(n1DtoOut.AbsIntRunes)
		lenN1FracRunes = len(n1DtoOut.AbsFracRunes)

		n1DtoOut.Precision = n2DtoOut.Precision
		err = n1DtoOut.ResetNumStrOut()

		if err != nil {
			return NumStrDto{}, NumStrDto{}, 0, false, err
		}

		lenN2AllRunes = len(n2DtoOut.AbsAllNumRunes)
		lenN2IntRunes = len(n2DtoOut.AbsIntRunes)
		lenN2FracRunes = len(n2DtoOut.AbsFracRunes)

	} else {
		// n1DtoOut.Precision == n2DtoOut.Precision

		lenN1AllRunes = len(n1DtoOut.AbsAllNumRunes)
		lenN1IntRunes = len(n1DtoOut.AbsIntRunes)
		lenN1FracRunes = len(n1DtoOut.AbsFracRunes)

		lenN2AllRunes = len(n2DtoOut.AbsAllNumRunes)
		lenN2IntRunes = len(n2DtoOut.AbsIntRunes)
		lenN2FracRunes = len(n2DtoOut.AbsFracRunes)

	}

	if lenN2IntRunes > lenN1IntRunes {

		var absAllRunes []rune
		var absIntRunes []rune
		deltaRunes := lenN2IntRunes - lenN1IntRunes
		for i := 0; i < deltaRunes; i++ {
			absAllRunes = append(absAllRunes, '0')
			absIntRunes = append(absIntRunes, '0')
		}

		for j := 0; j < lenN1AllRunes; j++ {
			absAllRunes = append(absAllRunes, n1DtoOut.AbsAllNumRunes[j])

			if j < lenN1IntRunes {
				absIntRunes = append(absIntRunes, n1DtoOut.AbsIntRunes[j])
			}

		}

		n1DtoOut.AbsAllNumRunes = absAllRunes
		n1DtoOut.AbsIntRunes = absIntRunes
		lenN1AllRunes = len(n1DtoOut.AbsAllNumRunes)
		lenN1IntRunes = len(n1DtoOut.AbsIntRunes)

		err = n1DtoOut.ResetNumStrOut()

		if err != nil {
			return NumStrDto{}, NumStrDto{}, 0, false, err
		}

	} else if lenN1IntRunes > lenN2IntRunes {

		var absAllRunes []rune
		var absIntRunes []rune
		deltaRunes := lenN1IntRunes - lenN2IntRunes
		for i := 0; i < deltaRunes; i++ {
			absAllRunes = append(absAllRunes, '0')
			absIntRunes = append(absIntRunes, '0')
		}

		for j := 0; j < lenN2AllRunes; j++ {
			absAllRunes = append(absAllRunes, n2DtoOut.AbsAllNumRunes[j])

			if j < lenN2IntRunes {
				absIntRunes = append(absIntRunes, n2DtoOut.AbsIntRunes[j])
			}

		}

		n2DtoOut.AbsAllNumRunes = absAllRunes
		n2DtoOut.AbsIntRunes = absIntRunes
		lenN2AllRunes = len(n2DtoOut.AbsAllNumRunes)
		lenN2IntRunes = len(n2DtoOut.AbsIntRunes)

		err := n2DtoOut.ResetNumStrOut()

		if err != nil {
			return NumStrDto{}, NumStrDto{}, 0, false, err
		}

	}

	if lenN1AllRunes != lenN2AllRunes {
		return NumStrDto{}, NumStrDto{}, 0, false, fmt.Errorf("FormatForMathOps() - n1 and n2 AllNumRune arrays are NOT equal in length. n1 length= '%v' n2 length= '%v'", lenN1AllRunes, lenN2AllRunes)
	}

	if lenN1IntRunes != lenN2IntRunes {
		return NumStrDto{}, NumStrDto{}, 0, false, fmt.Errorf("FormatForMathOps() - n1 and n2 IntRunes arrays are NOT equal in length. n1 length= '%v' n2 length= '%v'", lenN1IntRunes, lenN2IntRunes)
	}

	if lenN1FracRunes != lenN2FracRunes {
		return NumStrDto{}, NumStrDto{}, 0, false, fmt.Errorf("FormatForMathOps() - n1 and n2 FracRunes arrays are NOT equal in length. n1 length= '%v' n2 length= '%v'", lenN1FracRunes, lenN2FracRunes)
	}

	if n1DtoOut.Precision != n2DtoOut.Precision {
		return NumStrDto{}, NumStrDto{}, 0, false, fmt.Errorf("FormatForMathOps() - n1 and n2 FracRunes arrays are NOT equal in length. n1 length= '%v' n2 length= '%v'", lenN1FracRunes, lenN2FracRunes)

	}

	err = nDto.IsNumStrDtoValid(&n1DtoOut, "FormatForMathOps() n1Out - ")

	if err != nil {
		return NumStrDto{}, NumStrDto{}, 0, false, err
	}

	err = nDto.IsNumStrDtoValid(&n2DtoOut, "FormatForMathOps() n2Out - ")

	if err != nil {
		return NumStrDto{}, NumStrDto{}, 0, false, err
	}

	return n1DtoOut, n2DtoOut, compare, isOrderReversed, nil
}

// GetRationalNumber - returns the sign value of the number string, plus the
// numeric value of the number string expressed as a Rational Number.
//
// This method will return an error if the NumStrDto fields are not properly
// initialized and populated.
func (nDto *NumStrDto) GetRationalNumber() (int, *big.Rat, error) {

	ratZero := big.NewRat(0, 1)

	lenAbsAllNums := len(nDto.AbsAllNumRunes)
	lenAbsIntRunes := len(nDto.AbsIntRunes)
	lenAbsFracRunes := len(nDto.AbsFracRunes)

	if !nDto.IsValid || nDto.SignVal == 0 || len(nDto.AbsAllNumRunes) == 0 ||
		lenAbsAllNums != lenAbsIntRunes+lenAbsFracRunes {
		s := "GetAbsoluteBigInt() - The existing NumStrDto is corrupted or improperly initialized. Re-initialize the NumStrDto object and try again."
		return 0, ratZero, errors.New(s)

	}

	signVal := nDto.SignVal

	absInt, isOk := big.NewInt(0).SetString(string(nDto.AbsAllNumRunes), 10)

	if !isOk {
		return 0, ratZero, fmt.Errorf("GetRationalNumber() - Conversion of nDto.AbsAllNumRunes to big.Int Failed! nDto.AbsIntRunes= '%v'", nDto.AbsAllNumRunes)
	}

	base10 := big.NewInt(10)

	bigPrecision := big.NewInt(int64(nDto.Precision))

	scaleFactor := big.NewInt(0).Exp(base10, bigPrecision, nil)

	rationalNum := big.NewRat(1, 1).SetFrac(absInt, scaleFactor)

	return signVal, rationalNum, nil
}

// GetAbsoluteBigInt - Returns the absolute value of all numeric
// digits in the number string (nDto.AbsAllNumRunes). As such,
// Fractional digits to the right of the decimal are included
// in the consolidate integer number. All of the numeric digits
// in the number string are therefore returned as a *big.Int
// This method will fail if the NumStrDto has not been properly
// initialized with a valid number string.
func (nDto *NumStrDto) GetAbsoluteBigInt() (*big.Int, error) {

	lenAbsAllNums := len(nDto.AbsAllNumRunes)
	lenAbsIntRunes := len(nDto.AbsIntRunes)
	lenAbsFracRunes := len(nDto.AbsFracRunes)

	if !nDto.IsValid || nDto.SignVal == 0 || len(nDto.AbsAllNumRunes) == 0 ||
		lenAbsAllNums != lenAbsIntRunes+lenAbsFracRunes {
		s := "GetAbsoluteBigInt() - The existing NumStrDto is corrupted or improperly initialized. Re-initialize the NumStrDto object and try again."
		return big.NewInt(0), errors.New(s)

	}

	bigZero := big.NewInt(0)

	strAbsAllRunes := string(nDto.AbsAllNumRunes)

	absBigInt, isOk := bigZero.SetString(strAbsAllRunes, 10)

	if !isOk {
		s := fmt.Sprintf("GetAbsoluteBigInt() - Conversion of nDto.AbsAllNumRunes to *big.Int Failed!. nDto.AbsAllNumRunes= '%v'", strAbsAllRunes)
		return big.NewInt(0), errors.New(s)

	}

	return absBigInt, nil

}

// GetScaleFactor - returns the scale factor for this number
// string. Scale factor is defined by 10 raised to the power
// of nDto.Precision.  nDto.Precision is the number of
// digits to the right of the decimal point.
//
// This method will fail if the NumStrDto has not been properly
// initialized with a valid number string.
func (nDto *NumStrDto) GetScaleFactor() (*big.Int, error) {

	lenAbsAllNums := len(nDto.AbsAllNumRunes)
	lenAbsIntRunes := len(nDto.AbsIntRunes)
	lenAbsFracRunes := len(nDto.AbsFracRunes)

	if !nDto.IsValid || nDto.SignVal == 0 || len(nDto.AbsAllNumRunes) == 0 ||
		lenAbsAllNums != lenAbsIntRunes+lenAbsFracRunes {
		s := "GetScaleFactor() - The existing NumStrDto is corrupted or improperly initialized. Re-initialize the NumStrDto object and try again."
		return big.NewInt(0), errors.New(s)

	}

	if nDto.Precision == 0 {
		return big.NewInt(int64(1)), nil
	}

	base10 := big.NewInt(0).SetInt64(int64(10))

	bigPrecision := big.NewInt(0).SetInt64(int64(nDto.Precision))

	scaleFactor := big.NewInt(0).Exp(base10, bigPrecision, nil)

	return scaleFactor, nil

}

// GetSignedBigInt - returns the signed *big.Int representing
// the NumStrDto.NumStrOut. This method will fail if the NumStrDto
// has not been properly initialized with a valid number string.
func (nDto *NumStrDto) GetSignedBigInt() (*big.Int, error) {

	lenAbsAllNums := len(nDto.AbsAllNumRunes)
	lenAbsIntRunes := len(nDto.AbsIntRunes)
	lenAbsFracRunes := len(nDto.AbsFracRunes)

	if !nDto.IsValid || nDto.SignVal == 0 || len(nDto.AbsAllNumRunes) == 0 ||
		lenAbsAllNums != lenAbsIntRunes+lenAbsFracRunes {
		s := "GetSignedBigInt() - The existing NumStrDto is corrupted or improperly initialized. Re-initialize the NumStrDto object and try again."
		return big.NewInt(0), errors.New(s)

	}

	absBigInt, err := nDto.GetAbsoluteBigInt()

	if err != nil {
		s := fmt.Sprintf("GetSignedBigInt() - Error returned from nDto.GetAbsoluteBigInt(). Error= %v", err)
		return big.NewInt(0), errors.New(s)
	}

	if nDto.SignVal < 0 {
		signedBigInt := big.NewInt(0).Neg(absBigInt)
		return signedBigInt, nil
	}

	return absBigInt, nil
}

// GetZeroNumStr - returns a new NumStrDto initialized
// to zero value. If the parameter numFracDigits is set
// to a value greater than zero, then an equal number of
// zero characters will be added to the right of the
// decimal point.
// Examples:
// numFracDigits		Results NumStrOut
//	0									"0"
//	2									"0.00"
func (nDto *NumStrDto) GetZeroNumStr(numFracDigits uint) NumStrDto {

	// Set defaults for thousands separators,
	// decimal separators and currency Symbols
	if nDto.ThousandsSeparator == 0 {
		nDto.ThousandsSeparator = ','
	}

	if nDto.DecimalSeparator == 0 {
		nDto.DecimalSeparator = '.'
	}

	if nDto.CurrencySymbol == 0 {
		nDto.CurrencySymbol = '$'
	}

	n2Dto := NumStrDto{}.New()
	n2Dto.SignVal = 1
	n2Dto.ThousandsSeparator = nDto.ThousandsSeparator
	n2Dto.DecimalSeparator = nDto.DecimalSeparator
	n2Dto.CurrencySymbol = nDto.CurrencySymbol
	n2Dto.SignVal = 1
	n2Dto.IsFractionalValue = false
	n2Dto.Precision = 0
	n2Dto.HasNumericDigits = true
	n2Dto.AbsAllNumRunes = append(n2Dto.AbsAllNumRunes, '0')
	n2Dto.AbsIntRunes = append(n2Dto.AbsIntRunes, '0')
	n2Dto.AbsFracRunes = []rune{}
	n2Dto.NumStrOut = "0"

	if numFracDigits > 0 {
		n2Dto.IsFractionalValue = true
		n2Dto.NumStrOut = "0."

		for i := uint(0); i < numFracDigits; i++ {
			n2Dto.AbsAllNumRunes = append(n2Dto.AbsAllNumRunes, '0')
			n2Dto.AbsFracRunes = append(n2Dto.AbsFracRunes, '0')
			n2Dto.NumStrOut += "0"
		}

		n2Dto.Precision = uint(numFracDigits)
	}

	n2Dto.IsValid = true

	return n2Dto

}

func (nDto *NumStrDto) IsNumStrZeroValue(numDto *NumStrDto) bool {

	lenAbsAllNumRunes := len(numDto.AbsAllNumRunes)

	isZeroVal := true

	for i := 0; i < lenAbsAllNumRunes; i++ {
		if numDto.AbsAllNumRunes[i] != '0' {
			isZeroVal = false
		}
	}

	return isZeroVal
}

func (nDto *NumStrDto) IsNumStrDtoValid(numDto *NumStrDto, errName string) error {

	// Set defaults for thousands separators,
	// decimal separators and currency Symbols
	if numDto.ThousandsSeparator == 0 {
		numDto.ThousandsSeparator = ','
	}

	if numDto.DecimalSeparator == 0 {
		numDto.DecimalSeparator = '.'
	}

	if numDto.CurrencySymbol == 0 {
		numDto.CurrencySymbol = '$'
	}

	numDto.IsValid = false

	lenAbsAllNumRunes := len(numDto.AbsAllNumRunes)
	lenAbsFracRunes := len(numDto.AbsFracRunes)
	lenAbsIntRunes := len(numDto.AbsIntRunes)

	if lenAbsAllNumRunes > 0 {
		numDto.HasNumericDigits = true
	}

	if lenAbsFracRunes > 0 {
		numDto.IsFractionalValue = true
	}

	// Validate n2Dto object
	if lenAbsAllNumRunes != lenAbsIntRunes+lenAbsFracRunes {

		s1 := string(numDto.AbsAllNumRunes)
		s2 := string(numDto.AbsIntRunes)
		s3 := string(numDto.AbsFracRunes)

		return fmt.Errorf("%v - Length of Int Runes + Frac Runes does NOT equal len of All Runes. AllRunes= '%v' IntRunes= '%v' FracRunes= '%v' ", errName, s1, s2, s3)
	}

	if lenAbsFracRunes != int(numDto.Precision) {
		return fmt.Errorf("%v - Length of Frac Runes does NOT equal Precision.", errName)
	}

	if lenAbsFracRunes > 0 && lenAbsIntRunes == 0 {
		return fmt.Errorf("%v - Length of Frac Runes 1 or greater and length of Int Runes is ZERO!.", errName)
	}

	if numDto.SignVal != 1 && numDto.SignVal != -1 {
		return fmt.Errorf("%v - Sign Value is INVALID. Should be +1 or -1. This Sign Value is %v", errName, numDto.SignVal)
	}

	checkNumStrOut := ""

	if numDto.SignVal < 0 {
		checkNumStrOut = "-"
	}

	checkNumStrOut += string(numDto.AbsIntRunes)

	if numDto.Precision > 0 {
		checkNumStrOut += string(numDto.DecimalSeparator)
		checkNumStrOut += string(numDto.AbsFracRunes)
	}

	if checkNumStrOut != numDto.NumStrOut {
		return fmt.Errorf("%v - numDto.NumStrOut is incorrect!.", errName)
	}

	hasNonNumericChars := false

	for i := 0; i < lenAbsAllNumRunes; i++ {

		if numDto.AbsAllNumRunes[i] < '0' || numDto.AbsAllNumRunes[i] > '9' {
			hasNonNumericChars = true
			break
		}

		if i < lenAbsIntRunes &&
			(numDto.AbsIntRunes[i] < '0' || numDto.AbsIntRunes[i] > '9') {
			hasNonNumericChars = true
			break
		}

		if i >= lenAbsIntRunes &&
			(numDto.AbsFracRunes[i - lenAbsIntRunes] < '0' || numDto.AbsFracRunes[i - lenAbsIntRunes] > '9') {
			hasNonNumericChars = true
			break
		}
	}

	if hasNonNumericChars {
		return fmt.Errorf("%v - This NumStrDto contains Non-Numeric Characters and is INVALID!", errName)
	}

	numDto.IsValid = true

	return nil
}

func (nDto *NumStrDto) MultiplyNumStrs(n1Dto NumStrDto, n2Dto NumStrDto) (NumStrDto, error) {

	if err := nDto.IsNumStrDtoValid(&n1Dto, "MultiplyNumStrs() - "); err != nil {
		return NumStrDto{}, fmt.Errorf("MultiplyNumStrs() - n1Dto, first NumStrDto is invalid! Error= %v", err)
	}

	if err := nDto.IsNumStrDtoValid(&n2Dto, "MultiplyNumStrs() - "); err != nil {
		return NumStrDto{}, fmt.Errorf("MultiplyNumStrs() - n2Dto, second NumStrDto is invalid! Error= %v", err)
	}

	lenN1AbsAllRunes := len(n1Dto.AbsAllNumRunes)
	lenN2AbsAllRunes := len(n2Dto.AbsAllNumRunes)

	var n1Setup NumStrDto
	var n2Setup NumStrDto

	if lenN2AbsAllRunes > lenN1AbsAllRunes {
		n1Setup = n2Dto.CopyOut()
		n2Setup = n1Dto.CopyOut()
	} else if lenN1AbsAllRunes > lenN2AbsAllRunes {
		n1Setup = n1Dto.CopyOut()
		n2Setup = n2Dto.CopyOut()
	} else {
		// Must be lenN1AbsAllRunes == lenN2AbsAllRunes
		n1Setup = n1Dto.CopyOut()
		n2Setup = n2Dto.CopyOut()

	}


	newPrecision := n1Setup.Precision + n2Setup.Precision
	newSignVal := 1

	if n1Setup.SignVal == n2Setup.SignVal {
		newSignVal = 1
	} else {
		// Must be n1Setup.SignVal != n2Setup.SignVal
		newSignVal = -1
	}


	lenN1AbsAllRunes = len(n1Setup.AbsAllNumRunes)
	lenN2AbsAllRunes = len(n2Setup.AbsAllNumRunes)
	lenLevels := lenN2AbsAllRunes
	lenNumPlaces := (lenN1AbsAllRunes + lenN2AbsAllRunes) + 1

	intMAry := make([][]int, lenLevels)

	for i := 0; i < lenLevels; i++ {
		intMAry[i] = make([]int, lenNumPlaces)
	}

	intFinalAry := make([]int, lenNumPlaces+1)

	carry := 0
	levels := 0
	place := 0
	n1 := 0
	n2 := 0
	n3 := 0
	n4 := 0
	for i := lenN2AbsAllRunes - 1; i >= 0; i-- {

		place = (lenNumPlaces - 1) - levels

		for j := lenN1AbsAllRunes - 1; j >= 0; j-- {

			n1 = int(n1Setup.AbsAllNumRunes[j]) - 48
			n2 = int(n2Setup.AbsAllNumRunes[i]) - 48
			n3 = (n1 * n2) + carry
			n4 = int(math.Mod(float64(n3), float64(10.00)))

			intMAry[levels][place] = n4

			carry = int(n3 / 10)

			place--
		}

		intMAry[levels][place] = carry
		carry = 0
		levels++
	}

	carry = 0
	n1 = 0
	n2 = 0
	n3 = 0
	n4 = 0
	for i := 0; i < lenLevels; i++ {
		for j := lenNumPlaces - 1; j >= 0; j-- {

			n1 = intFinalAry[j+1]
			n2 = intMAry[i][j]
			n3 = n1 + n2 + carry
			n4 = 0

			if n3 > 9 {
				n4 = int(math.Mod(float64(n3), float64(10.0)))
				carry = n3 / 10

			} else {
				n4 = n3
				carry = 0
			}

			intFinalAry[j+1] = n4
		}

		if carry > 0 {
			intFinalAry[0] = carry
		}

	}

	numStrOut, err := nDto.FindIntArraySignificantDigitLimits(intFinalAry, newPrecision, newSignVal)

	if err != nil {
		return NumStrDto{}, fmt.Errorf("MultiplyNumStrs() - Error returned from nDto.FindIntArraySignificantDigitLimits(intFinalAry,newPrecision, newSignVal). Error= %v", err)
	}

	return numStrOut, nil
}

// ParseSignedBigInt - receives a signed *Big Int number and precision parameter. It then
// generates and returns a new NumStrDto type.
func (nDto *NumStrDto) ParseSignedBigInt(signedBigInt *big.Int, precision uint) (NumStrDto, error) {
	bigZero := big.NewInt(0)

	// Set defaults for thousands separators,
	// decimal separators and currency Symbols
	if nDto.ThousandsSeparator == 0 {
		nDto.ThousandsSeparator = ','
	}

	if nDto.DecimalSeparator == 0 {
		nDto.DecimalSeparator = '.'
	}

	if nDto.CurrencySymbol == 0 {
		nDto.CurrencySymbol = '$'
	}

	n2Dto := NumStrDto{}.New()
	n2Dto.NumStrIn = signedBigInt.String()

	if signedBigInt.Cmp(bigZero) == 0 {
		return nDto.GetZeroNumStr(0), nil
	}

	if precision == 0 {

		return nDto.ParseNumStr(signedBigInt.String())
	}

	signVal := 1

	if signedBigInt.Cmp(bigZero) < 0 {
		signVal = -1
	}

	absBigInt := big.NewInt(0).Abs(signedBigInt)

	n2Dto.SignVal = signVal
	n2Dto.Precision = precision

	absAllNumRunes := []rune(string(absBigInt.String()))
	lenAbsAllNumRunes := len(absAllNumRunes)
	iSpecPrecision := int(precision)
	if iSpecPrecision >= lenAbsAllNumRunes {
		deltaPrecision := (iSpecPrecision - lenAbsAllNumRunes) + 1
		for i := 0; i < deltaPrecision; i++ {
			n2Dto.AbsAllNumRunes = append(n2Dto.AbsAllNumRunes, '0')
		}
	}

	for i := 0; i < lenAbsAllNumRunes; i++ {
		n2Dto.AbsAllNumRunes = append(n2Dto.AbsAllNumRunes, absAllNumRunes[i])
	}

	lenAbsAllNumRunes = len(n2Dto.AbsAllNumRunes)
	lenAbsIntNumRunes := lenAbsAllNumRunes - iSpecPrecision

	for j := 0; j < lenAbsAllNumRunes; j++ {
		if j < lenAbsIntNumRunes {
			n2Dto.AbsIntRunes = append(n2Dto.AbsIntRunes, n2Dto.AbsAllNumRunes[j])
			n2Dto.HasNumericDigits = true
		} else {
			n2Dto.AbsFracRunes = append(n2Dto.AbsFracRunes, n2Dto.AbsAllNumRunes[j])
			n2Dto.IsFractionalValue = true
		}
	}

	lenAbsIntNumRunes = len(n2Dto.AbsIntRunes)
	lenAbsFracNumRunes := len(n2Dto.AbsFracRunes)

	if lenAbsAllNumRunes != lenAbsIntNumRunes+lenAbsFracNumRunes {
		return NumStrDto{}, fmt.Errorf("ParseSignedBigInt() lenAbsAllNumRunes != lenAbsIntNumRunes + lenAbsFracNumRunes. lenAbsAllNumRunes= '%v' lenAbsIntNumRunes= '%v' lenAbsFracNumRunes= '%v'", lenAbsAllNumRunes, lenAbsIntNumRunes, lenAbsFracNumRunes)
	}

	n2Dto.NumStrOut = ""

	if n2Dto.SignVal < 0 {
		n2Dto.NumStrOut = "-"
	}

	n2Dto.NumStrOut += string(n2Dto.AbsIntRunes)

	if lenAbsFracNumRunes > 0 {
		n2Dto.NumStrOut += "."
		n2Dto.NumStrOut += string(n2Dto.AbsFracRunes)
	}

	err := nDto.IsNumStrDtoValid(&n2Dto, "ParseSignedBigInt() - ")

	if err != nil {
		return NumStrDto{}, err
	}

	n2Dto.IsValid = true

	return n2Dto, nil

}

// ParseNumStr - receives a raw string and converts to a properly
// formatted number string. The string is returned via a NumStrDto type.
// Returned number strings may consist of a leading negative sign ('-')
// numeric digits and may include a decimal separator ('.'). The NumStrDto
// breaks the string down into Sign, Integer and Fractional components.
func (nDto *NumStrDto) ParseNumStr(str string) (NumStrDto, error) {

	if len(str) == 0 {
		return NumStrDto{}, errors.New("ParseNumStr() Received zero length number string!")
	}

	// Set defaults for thousands separators,
	// decimal separators and currency Symbols
	if nDto.ThousandsSeparator == 0 {
		nDto.ThousandsSeparator = ','
	}

	if nDto.DecimalSeparator == 0 {
		nDto.DecimalSeparator = '.'
	}

	if nDto.CurrencySymbol == 0 {
		nDto.CurrencySymbol = '$'
	}

	n2Dto := NumStrDto{}.New()

	n2Dto.NumStrIn = str

	n2Dto.SignVal = 1
	n2Dto.ThousandsSeparator = nDto.ThousandsSeparator
	n2Dto.DecimalSeparator = nDto.DecimalSeparator
	n2Dto.CurrencySymbol = nDto.CurrencySymbol
	baseRunes := []rune(n2Dto.NumStrIn)
	lBaseRunes := len(baseRunes)
	isStartRunes := false
	isEndRunes := false
	lCurRunes := len(NumStrCurrencySymbols)
	isSkip := false

	for i := 0; i < lBaseRunes && isEndRunes == false; i++ {

		for j := 0; j < lCurRunes; j++ {
			if baseRunes[i] == NumStrCurrencySymbols[j] {
				isSkip = true
				break
			}
		}

		if isSkip == true || baseRunes[i] == '+' ||
			baseRunes[i] == ' ' || baseRunes[i] == ',' ||
			baseRunes[i] == '$' ||
			baseRunes[i] == n2Dto.ThousandsSeparator ||
			baseRunes[i] == n2Dto.CurrencySymbol {

			isSkip = false
			continue

		}

		if baseRunes[i] == '-' &&
			isStartRunes == false && isEndRunes == false &&
			i+1 < lBaseRunes &&
			((baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9') ||
				(baseRunes[i+1] == '.' || baseRunes[i+1] == n2Dto.DecimalSeparator)) {

			n2Dto.SignVal = -1
			isStartRunes = true
			continue

		} else if isEndRunes == false &&
			baseRunes[i] >= '0' && baseRunes[i] <= '9' {

			n2Dto.AbsAllNumRunes = append(n2Dto.AbsAllNumRunes, baseRunes[i])
			isStartRunes = true
			n2Dto.HasNumericDigits = true

			if n2Dto.IsFractionalValue {
				n2Dto.AbsFracRunes = append(n2Dto.AbsFracRunes, baseRunes[i])
			} else {
				n2Dto.AbsIntRunes = append(n2Dto.AbsIntRunes, baseRunes[i])
			}

		} else if isEndRunes == false &&
			i+1 < lBaseRunes &&
			baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9' &&
			(baseRunes[i] == '.' || baseRunes[i] == n2Dto.DecimalSeparator) {

			n2Dto.IsFractionalValue = true
			continue

		} else if isStartRunes && !isEndRunes {

			isEndRunes = true

		}

	}

	lenAbsAllNumRunes := len(n2Dto.AbsAllNumRunes)

	if lenAbsAllNumRunes == 0 {
		nZeroNumStr := nDto.GetZeroNumStr(0)
		return nZeroNumStr, nil
	}

	lenAbsIntNumRunes := len(n2Dto.AbsIntRunes)
	if lenAbsIntNumRunes == 0 {
		n2Dto.AbsIntRunes = append(n2Dto.AbsIntRunes, '0')
	}

	lenAbsAllNumRunes = len(n2Dto.AbsAllNumRunes)
	lenAbsIntNumRunes = len(n2Dto.AbsIntRunes)
	lenAbsFracNumRunes := len(n2Dto.AbsFracRunes)

	isZeroVal := true

	for i := 0; i < lenAbsAllNumRunes; i++ {

		if n2Dto.AbsAllNumRunes[i] != '0' {
			isZeroVal = false
		}

	}

	if isZeroVal {
		nZeroDto := nDto.GetZeroNumStr(uint(lenAbsFracNumRunes))
		return nZeroDto, nil
	}

	if n2Dto.SignVal < 0 {
		n2Dto.NumStrOut = "-"
	}

	n2Dto.NumStrOut += string(n2Dto.AbsIntRunes)

	if n2Dto.IsFractionalValue {
		n2Dto.Precision = uint(len(n2Dto.AbsFracRunes))
		n2Dto.NumStrOut += string(nDto.DecimalSeparator)
		n2Dto.NumStrOut += string(n2Dto.AbsFracRunes)
	}

	if lenAbsAllNumRunes != lenAbsIntNumRunes+lenAbsFracNumRunes {
		n2Dto.AbsAllNumRunes = []rune{}
		newLenAbsAllNumRunes := lenAbsIntNumRunes + lenAbsFracNumRunes

		for i := 0; i < newLenAbsAllNumRunes; i++ {
			if i < lenAbsIntNumRunes {
				n2Dto.AbsAllNumRunes = append(n2Dto.AbsAllNumRunes, n2Dto.AbsIntRunes[i])
			} else {
				n2Dto.AbsAllNumRunes = append(n2Dto.AbsAllNumRunes, n2Dto.AbsFracRunes[i-lenAbsIntNumRunes])
			}
		}

		lenAbsAllNumRunes = len(n2Dto.AbsAllNumRunes)
	}

	// Validate n2Dto object
	err := nDto.IsNumStrDtoValid(&n2Dto, "ParseNumStr() - ")

	if err != nil {
		return NumStrDto{}, err
	}

	n2Dto.IsValid = true

	return n2Dto, nil
}

func (nDto *NumStrDto) ScaleNumStr(signedNumStr string, precision int, roundResult bool) (NumStrDto, error) {

	if precision < 0 {
		precision = precision * -1
	}

	return nDto.ShiftPrecisionLeft(signedNumStr, uint(precision))
}

// ShiftPrecisionLeft - Shifts the existing precision of a number string. The position of
// the decimal point is shifted 'precision' positions to the left.
//
// This is equivalent to: result = signedNumStr / 10^precision or signedNumStr divided
// by 10 raised to the power of precision.
//
// Examples:
// signedNumStr			precision				Result
//	"123456.789"				3						"123.456789"
//	"123456.789"				2						"1234.56789"
//	"123456.789"	   		6					  "0.123456789"
//	"123456789"					6						"123.456789"
//	"123"								5						"0.00123"
//  "0"									3						"0"
// "123456.789"					0						"123456.789"		- Zero has no effect on original number string
func (nDto *NumStrDto) ShiftPrecisionLeft(signedNumStr string, precision uint) (NumStrDto, error) {

	if len(signedNumStr) == 0 {
		return NumStrDto{}, errors.New("ShiftPrecisionLeft() Received zero length number string!")
	}

	// Set defaults for thousands separators,
	// decimal separators and currency Symbols
	if nDto.ThousandsSeparator == 0 {
		nDto.ThousandsSeparator = ','
	}

	if nDto.DecimalSeparator == 0 {
		nDto.DecimalSeparator = '.'
	}

	if nDto.CurrencySymbol == 0 {
		nDto.CurrencySymbol = '$'
	}

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(signedNumStr)

	if err != nil {
		return NumStrDto{}, fmt.Errorf("ShiftPrecisionLeft() - Received Error from NumStrDto.ParseNumStr(signedNumStr). str= '%v' Error= %v", signedNumStr, err)
	}

	n2 := NumStrDto{}.New()

	n2.ThousandsSeparator = nDto.ThousandsSeparator
	n2.DecimalSeparator = nDto.DecimalSeparator
	n2.CurrencySymbol = nDto.CurrencySymbol
	n2.SignVal = n1.SignVal
	n2.Precision = precision + n1.Precision
	n2.NumStrIn = n1.NumStrIn
	iTotalSpecPrecision := int(n2.Precision)
	lenAbsAllNumRunes := len(n1.AbsAllNumRunes)
	lenAbsIntRunes := len(n1.AbsIntRunes)
	lenAbsFracRunes := len(n1.AbsFracRunes)

	if nDto.IsNumStrZeroValue(&n1) {

		return nDto.GetZeroNumStr(0), nil
	}

	if iTotalSpecPrecision == lenAbsAllNumRunes {

		n2.AbsAllNumRunes = append(n2.AbsAllNumRunes, '0')

	} else if iTotalSpecPrecision > lenAbsAllNumRunes {

		deltaPrecision := iTotalSpecPrecision - lenAbsAllNumRunes + 1

		for i := 0; i < deltaPrecision; i++ {
			n2.AbsAllNumRunes = append(n2.AbsAllNumRunes, '0')
		}

	}

	for j := 0; j < lenAbsAllNumRunes; j++ {
		n2.AbsAllNumRunes = append(n2.AbsAllNumRunes, n1.AbsAllNumRunes[j])
	}

	lenAbsAllNumRunes = len(n2.AbsAllNumRunes)
	lenAbsFracRunes = iTotalSpecPrecision
	lenAbsIntRunes = lenAbsAllNumRunes - lenAbsFracRunes

	if lenAbsIntRunes <= 0 {
		return NumStrDto{}, fmt.Errorf("ShiftPrecisionLeft() - Calculated number of integer digits is less than or equal to ZERO. lenAbsIntRunes= '%v' ", lenAbsIntRunes)
	}

	for i := 0; i < lenAbsAllNumRunes; i++ {

		if i < lenAbsIntRunes {
			n2.AbsIntRunes = append(n2.AbsIntRunes, n2.AbsAllNumRunes[i])
		} else {
			n2.AbsFracRunes = append(n2.AbsFracRunes, n2.AbsAllNumRunes[i])
		}
	}

	lenAbsFracRunes = len(n2.AbsFracRunes)

	if n2.SignVal < 0 {
		n2.NumStrOut = "-"
	}

	n2.NumStrOut += string(n2.AbsIntRunes)

	if lenAbsFracRunes > 0 {
		n2.NumStrOut += "."
		n2.NumStrOut += string(n2.AbsFracRunes)
	}

	err = nDto.IsNumStrDtoValid(&n2, "ShiftPrecisionLeft()")

	if err != nil {
		return NumStrDto{}, err
	}

	n2.IsValid = true

	return n2, nil
}

// ShiftPrecisionRight - Shifts the existing precision of a number string. The position of
// the decimal point is shifted 'precision' positions to the right.
//
// This is equivalent to: result = signedNumStr X 10^precision or signedNumStr Multiplied
// by 10 raised to the power of precision.
//
// Examples:
// signedNumStr			precision			Result
// "123456.789"				3						"123456789"
// "123456.789"				2						"12345678.9"
// "123456.789"       6					  "123456789000"
// "123456789"				6						"123456789000000"
// "0"								3						"0"
// "123456.789"				0						"123456.789"		- Zero has no effect on original number string
func (nDto *NumStrDto) ShiftPrecisionRight(signedNumStr string, precision uint) (NumStrDto, error) {

	if len(signedNumStr) == 0 {
		return NumStrDto{}, errors.New("ShiftPrecisionRight() Received zero length number string!")
	}

	// Set defaults for thousands separators,
	// decimal separators and currency Symbols
	if nDto.ThousandsSeparator == 0 {
		nDto.ThousandsSeparator = ','
	}

	if nDto.DecimalSeparator == 0 {
		nDto.DecimalSeparator = '.'
	}

	if nDto.CurrencySymbol == 0 {
		nDto.CurrencySymbol = '$'
	}

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(signedNumStr)

	if err != nil {
		return NumStrDto{}, fmt.Errorf("ShiftPrecisionRight() - Received Error from NumStrDto.ParseNumStr(signedNumStr). str= '%v' Error= %v", signedNumStr, err)
	}

	n2 := NumStrDto{}.New()

	iTotalSpecPrecision := 0
	iPrecision := int(precision)
	iN1Precision := int(n1.Precision)

	if iN1Precision > 0 && iPrecision < iN1Precision {
		iTotalSpecPrecision = iN1Precision - iPrecision
	} else {
		iTotalSpecPrecision = 0
	}

	n2.ThousandsSeparator = nDto.ThousandsSeparator
	n2.DecimalSeparator = nDto.DecimalSeparator
	n2.CurrencySymbol = nDto.CurrencySymbol
	n2.SignVal = n1.SignVal
	n2.Precision = uint(iTotalSpecPrecision)
	n2.NumStrIn = n1.NumStrIn

	lenAbsAllNumRunes := len(n1.AbsAllNumRunes)

	if nDto.IsNumStrZeroValue(&n1) {

		return nDto.GetZeroNumStr(0), nil
	}

	if int(precision) > int(n1.Precision) {

		for i := 0; i < lenAbsAllNumRunes; i++ {
			n2.AbsAllNumRunes = append(n2.AbsAllNumRunes, n1.AbsAllNumRunes[i])
		}

		deltaPrecision := int(precision) - int(n1.Precision)

		for i := 0; i < deltaPrecision; i++ {
			n2.AbsAllNumRunes = append(n2.AbsAllNumRunes, '0')
		}

	} else {

		for i := 0; i < lenAbsAllNumRunes; i++ {
			n2.AbsAllNumRunes = append(n2.AbsAllNumRunes, n1.AbsAllNumRunes[i])
		}

	}

	lenAbsAllNumRunes = len(n2.AbsAllNumRunes)
	lenAbsFracRunes := iTotalSpecPrecision
	lenAbsIntRunes := lenAbsAllNumRunes - lenAbsFracRunes

	if lenAbsIntRunes <= 0 {
		return NumStrDto{}, fmt.Errorf("ShiftPrecisionRight() - Calculated number of integer digits is less than or equal to ZERO. lenAbsIntRunes= '%v' ", lenAbsIntRunes)
	}

	for i := 0; i < lenAbsAllNumRunes; i++ {

		if i < lenAbsIntRunes {
			n2.AbsIntRunes = append(n2.AbsIntRunes, n2.AbsAllNumRunes[i])
		} else {
			n2.AbsFracRunes = append(n2.AbsFracRunes, n2.AbsAllNumRunes[i])
		}
	}

	lenAbsFracRunes = len(n2.AbsFracRunes)

	if n2.SignVal < 0 {
		n2.NumStrOut = "-"
	}

	n2.NumStrOut += string(n2.AbsIntRunes)

	if lenAbsFracRunes > 0 {
		n2.NumStrOut += "."
		n2.NumStrOut += string(n2.AbsFracRunes)
	}

	err = nDto.IsNumStrDtoValid(&n2, "ShiftPrecisionRight()")

	if err != nil {
		return NumStrDto{}, err
	}

	n2.IsValid = true

	return n2, nil
}

// TODO - Needs more testing
// ScaleAbsoluteValStr - Takes the absolute value of a signed number string parameter
// and applies the specified precision. If set to 'true', the boolean input parameter
// 'roundResult' will apply rounding in cases where the value of 'precision' is less
// less than the existing number of digits to the right of the decimal.
//
// Examples:
//    --------- Input Parameters ---------
//    signedNumStr  Precision    Round Result    Result
// 		"123.456"     7            false           "0.0123456"
//    "123456"      7					   false           "0.0123456"
//    "123.456      2         	 true            "123.46"
//    "123456"      3            false           "123.456"
//    "123456"      5            false           "1.23456"
//    "123.456      3            false           "123.456"
//    "123.456"     4            false           "12.3456"
func (nDto *NumStrDto) ScaleAbsoluteValStr(signedNumStr string, precision uint, roundResult bool) (NumStrDto, error) {

	if len(signedNumStr) == 0 {
		return NumStrDto{}, errors.New("ScaleAbsoluteValStr() Received zero length number string!")
	}

	// Set defaults for thousands separators,
	// decimal separators and currency Symbols
	if nDto.ThousandsSeparator == 0 {
		nDto.ThousandsSeparator = ','
	}

	if nDto.DecimalSeparator == 0 {
		nDto.DecimalSeparator = '.'
	}

	if nDto.CurrencySymbol == 0 {
		nDto.CurrencySymbol = '$'
	}

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(signedNumStr)

	if err != nil {
		return NumStrDto{}, fmt.Errorf("ScaleAbsoluteValStr() - Received Error from NumStrDto.ParseNumStr(signedNumStr). str= '%v' Error= %v", signedNumStr, err)
	}

	n2 := NumStrDto{}.New()
	n2.NumStrIn = signedNumStr
	n2.ThousandsSeparator = nDto.ThousandsSeparator
	n2.DecimalSeparator = nDto.DecimalSeparator
	n2.CurrencySymbol = nDto.CurrencySymbol

	n2.SignVal = n1.SignVal
	n2.Precision = precision

	lenN1AbsFracRunes := len(n1.AbsFracRunes)
	lenN1AbsIntRunes := len(n1.AbsIntRunes)
	lenN1AbsAllNumRunes := len(n1.AbsAllNumRunes)
	iSpecPrecision := int(precision)

	if roundResult && lenN1AbsFracRunes > 0 &&
		iSpecPrecision < lenN1AbsFracRunes {

		absAllNumsToRound, isOk := big.NewInt(0).SetString(string(n1.AbsAllNumRunes), 10)

		if !isOk {
			return NumStrDto{}, fmt.Errorf("ScaleAbsoluteValStr()- Error: Failed to convert string to big.Int(). big.Int.SetString(n1.AbsAllNumRunes). n1.AbsAllNumRunes='%v' ", string(n1.AbsAllNumRunes))
		}

		bigDeltaPrecision := big.NewInt(int64(lenN1AbsFracRunes - iSpecPrecision - 1))
		base10 := big.NewInt(int64(10))
		roundUp5 := big.NewInt(int64(5))
		roundScaleFactor := big.NewInt(0).Exp(base10, bigDeltaPrecision, nil)
		roundUpNum := big.NewInt(0).Mul(roundUp5, roundScaleFactor)
		roundedAbsAllNums := big.NewInt(0).Add(absAllNumsToRound, roundUpNum)
		actualDeltaPrecision := big.NewInt(int64(lenN1AbsFracRunes - iSpecPrecision))
		actualDeltaScaleFactor := big.NewInt(0).Exp(base10, actualDeltaPrecision, nil)
		actualAbsAllNums := big.NewInt(0).Div(roundedAbsAllNums, actualDeltaScaleFactor)
		n1.AbsAllNumRunes = []rune{}
		n1.AbsIntRunes = []rune{}
		n1.AbsFracRunes = []rune{}
		n1.AbsAllNumRunes = []rune(string(actualAbsAllNums.String()))
		lenN1AbsAllNumRunes = len(n1.AbsAllNumRunes)

		for i := 0; i < lenN1AbsAllNumRunes; i++ {

			if i < lenN1AbsIntRunes {
				n1.AbsIntRunes = append(n1.AbsIntRunes, n1.AbsAllNumRunes[i])
			} else {
				n1.AbsFracRunes = append(n1.AbsFracRunes, n1.AbsAllNumRunes[i])
			}
		}

		lenN1AbsIntRunes = len(n1.AbsIntRunes)
		lenN1AbsFracRunes = len(n1.AbsFracRunes)

		if lenN1AbsAllNumRunes != (lenN1AbsIntRunes + lenN1AbsFracRunes) {

			return NumStrDto{}, fmt.Errorf("ScaleAbsoluteValStr()- Error on Rounding. lenN1AbsAllNumRunes != (lenN1AbsIntRunes + lenN1AbsFracRunes). lenN1AbsAllNumRunes= '%v' lenN1AbsIntRunes= '%v' lenN1AbsFracRunes= '%v'", lenN1AbsAllNumRunes, lenN1AbsIntRunes, lenN1AbsFracRunes)
		}

	}

	if iSpecPrecision >= lenN1AbsAllNumRunes {
		deltaPrecision := (iSpecPrecision - lenN1AbsAllNumRunes) + 1
		for i := 0; i < deltaPrecision; i++ {
			n2.AbsAllNumRunes = append(n2.AbsAllNumRunes, '0')
		}
	}

	for i := 0; i < lenN1AbsAllNumRunes; i++ {
		n2.AbsAllNumRunes = append(n2.AbsAllNumRunes, n1.AbsAllNumRunes[i])
	}

	lenAbsAllNumRunes := len(n2.AbsAllNumRunes)
	lenAbsIntNumRunes := lenAbsAllNumRunes - iSpecPrecision

	for j := 0; j < lenAbsAllNumRunes; j++ {
		if j < lenAbsIntNumRunes {
			n2.AbsIntRunes = append(n2.AbsIntRunes, n2.AbsAllNumRunes[j])
			n2.HasNumericDigits = true
		} else {
			n2.AbsFracRunes = append(n2.AbsFracRunes, n2.AbsAllNumRunes[j])
			n2.IsFractionalValue = true
		}
	}

	lenAbsIntNumRunes = len(n2.AbsIntRunes)
	lenAbsFracNumRunes := len(n2.AbsFracRunes)

	if lenAbsAllNumRunes != lenAbsIntNumRunes+lenAbsFracNumRunes {
		return NumStrDto{}, fmt.Errorf("ScaleAbsoluteValStr() Error: lenAbsAllNumRunes != lenAbsIntNumRunes + lenAbsFracNumRunes. lenAbsAllNumRunes= '%v' lenAbsIntNumRunes= '%v' lenAbsFracNumRunes= '%v'", lenAbsAllNumRunes, lenAbsIntNumRunes, lenAbsFracNumRunes)
	}

	if lenAbsFracNumRunes != int(n2.Precision) {
		return NumStrDto{}, fmt.Errorf("ScaleAbsoluteValStr() Error: lenAbsAllNumRunes != lenAbsIntNumRunes + lenAbsFracNumRunes. lenAbsAllNumRunes= '%v' lenAbsIntNumRunes= '%v' lenAbsFracNumRunes= '%v'", lenAbsAllNumRunes, lenAbsIntNumRunes, lenAbsFracNumRunes)
	}

	n2.NumStrOut = ""

	if n2.SignVal < 0 {
		n2.NumStrOut = "-"
	}

	n2.NumStrOut += string(n2.AbsIntRunes)

	if lenAbsFracNumRunes > 0 {
		n2.NumStrOut += string(nDto.DecimalSeparator)
		n2.NumStrOut += string(n2.AbsFracRunes)
	}

	n2.IsValid = true

	return n2, nil
}

// TODO - Needs more testing
// SetPrecision - parses the incoming number string and applies the designated 'precision'. 'precision'
// determines the number of digits to the right of the decimal place. The boolean parameter 'roundResult'
// is used to apply rounding in those cases where 'precision' dictates a reduction in the number of
// digits to the right of the decimal place.
//
// Examples:
// ----------_- Input Parameters --_---------			Result
// signedNumStr			precision			roundResult
// "123456"				  7							false						"123456.0000000"
// "123.456"				2							true						"123.46"
// "123.456         5             false						"123.45600"
func (nDto *NumStrDto) SetPrecision(signedNumStr string, precision uint, roundResult bool) (NumStrDto, error) {

	if len(signedNumStr) == 0 {
		return NumStrDto{}, errors.New("SetPrecision() Received zero length number string!")
	}

	// Set defaults for thousands separators,
	// decimal separators and currency Symbols
	if nDto.ThousandsSeparator == 0 {
		nDto.ThousandsSeparator = ','
	}

	if nDto.DecimalSeparator == 0 {
		nDto.DecimalSeparator = '.'
	}

	if nDto.CurrencySymbol == 0 {
		nDto.CurrencySymbol = '$'
	}

	n0 := NumStrDto{}.New()
	n0.ThousandsSeparator = nDto.ThousandsSeparator
	n0.DecimalSeparator = nDto.DecimalSeparator
	n0.CurrencySymbol = nDto.CurrencySymbol

	n1, err := n0.ParseNumStr(signedNumStr)

	if err != nil {
		return NumStrDto{}, fmt.Errorf("SetPrecision()- Error returned from ns.ParseNumString(signedNumStr). signedNumStr='%v' Error= %v", signedNumStr, err)
	}

	n2 := NumStrDto{}.New()

	n2.SignVal = n1.SignVal
	n2.Precision = precision
	n2.ThousandsSeparator = nDto.ThousandsSeparator
	n2.DecimalSeparator = nDto.DecimalSeparator
	n2.CurrencySymbol = nDto.CurrencySymbol
	n2.HasNumericDigits = true
	n2.NumStrIn = signedNumStr
	iSpecPrecision := int(precision)
	lenN1AbsAllNumRunes := len(n1.AbsAllNumRunes)
	lenN1AbsIntRunes := len(n1.AbsIntRunes)
	lenN1AbsFracRunes := len(n1.AbsFracRunes)
	totalRunes := 0

	if roundResult && lenN1AbsFracRunes > 0 &&
		iSpecPrecision < lenN1AbsFracRunes {

		absAllNumsToRound, isOk := big.NewInt(0).SetString(string(n1.AbsAllNumRunes), 10)

		if !isOk {
			return NumStrDto{}, fmt.Errorf("SetPrecision()- Error: Failed to convert string to big.Int(). big.Int.SetString(n1.AbsAllNumRunes). n1.AbsAllNumRunes='%v' ", string(n1.AbsAllNumRunes))
		}

		bigDeltaPrecision := big.NewInt(int64(lenN1AbsFracRunes - iSpecPrecision - 1))
		base10 := big.NewInt(int64(10))
		roundUp5 := big.NewInt(int64(5))
		roundScaleFactor := big.NewInt(0).Exp(base10, bigDeltaPrecision, nil)
		roundUpNum := big.NewInt(0).Mul(roundUp5, roundScaleFactor)
		roundedAbsAllNums := big.NewInt(0).Add(absAllNumsToRound, roundUpNum)
		actualDeltaPrecision := big.NewInt(int64(lenN1AbsFracRunes - iSpecPrecision))
		actualDeltaScaleFactor := big.NewInt(0).Exp(base10, actualDeltaPrecision, nil)
		actualAbsAllNums := big.NewInt(0).Div(roundedAbsAllNums, actualDeltaScaleFactor)
		n1.AbsAllNumRunes = []rune{}
		n1.AbsIntRunes = []rune{}
		n1.AbsFracRunes = []rune{}
		n1.AbsAllNumRunes = []rune(string(actualAbsAllNums.String()))
		lenN1AbsAllNumRunes = len(n1.AbsAllNumRunes)

		for i := 0; i < lenN1AbsAllNumRunes; i++ {

			if i < lenN1AbsIntRunes {
				n1.AbsIntRunes = append(n1.AbsIntRunes, n1.AbsAllNumRunes[i])
			} else {
				n1.AbsFracRunes = append(n1.AbsFracRunes, n1.AbsAllNumRunes[i])
			}
		}

		lenN1AbsIntRunes = len(n1.AbsIntRunes)
		lenN1AbsFracRunes = len(n1.AbsFracRunes)

		if lenN1AbsAllNumRunes != (lenN1AbsIntRunes + lenN1AbsFracRunes) {

			return NumStrDto{}, fmt.Errorf("SetPrecision()- Error on Rounding. lenN1AbsAllNumRunes != (lenN1AbsIntRunes + lenN1AbsFracRunes). lenN1AbsAllNumRunes= '%v' lenN1AbsIntRunes= '%v' lenN1AbsFracRunes= '%v'", lenN1AbsAllNumRunes, lenN1AbsIntRunes, lenN1AbsFracRunes)
		}

	}

	if lenN1AbsIntRunes == 0 {
		n2.AbsAllNumRunes = append(n2.AbsAllNumRunes, '0')
		n2.AbsIntRunes = append(n2.AbsIntRunes, '0')
	}

	totalRunes = lenN1AbsIntRunes + iSpecPrecision

	for i := 0; i < totalRunes; i++ {

		if i < lenN1AbsAllNumRunes {
			n2.AbsAllNumRunes = append(n2.AbsAllNumRunes, n1.AbsAllNumRunes[i])
		} else {
			n2.AbsAllNumRunes = append(n2.AbsAllNumRunes, '0')
		}

		if i < lenN1AbsIntRunes {

			n2.AbsIntRunes = append(n2.AbsIntRunes, n1.AbsAllNumRunes[i])

		} else {

			if i < lenN1AbsAllNumRunes {
				n2.AbsFracRunes = append(n2.AbsFracRunes, n1.AbsAllNumRunes[i])
			} else {
				n2.AbsFracRunes = append(n2.AbsFracRunes, '0')
			}
		}
	}

	n2.NumStrOut = ""

	if n2.SignVal < 0 {
		n2.NumStrOut = "-"
	}

	n2.NumStrOut += string(n2.AbsIntRunes)

	if n2.Precision > 0 {
		n2.NumStrOut += string(n2.DecimalSeparator)
		n2.NumStrOut += string(n2.AbsFracRunes)
		n2.IsFractionalValue = true
	}

	err = nDto.IsNumStrDtoValid(&n2, "SetPrecision()- ")

	if err != nil {

		return NumStrDto{}, err
	}

	n2.IsValid = true

	return n2, nil
}

// SetSignValue - Sets the sign of the numeric value
// for the current NumStrDto. Only two values are
// allowed: +1 and -1. If any other value is passed
// an error is thrown
func (nDto *NumStrDto) SetSignValue(newSignVal int) error {

	if err := nDto.IsNumStrDtoValid(nDto, "ChangeSignValue() - "); err != nil {
		return err
	}

	if newSignVal != -1 && newSignVal != 1 {
		return fmt.Errorf("SetSignValue() invalid sign value passed. Sign must be +1 or -1. This sign value= %v", newSignVal)
	}

	nDto.SignVal = newSignVal

	return nDto.ResetNumStrOut()
}

// ResetNumStrOut - Re-creates the NumStrOut field using
// the current 'Precision' value.
func (nDto *NumStrDto) ResetNumStrOut() error {

	nDto.NumStrOut = ""

	if nDto.SignVal < 0 {
		nDto.NumStrOut = "-"
	}

	nDto.NumStrOut += string(nDto.AbsIntRunes)

	if nDto.Precision > 0 {
		nDto.NumStrOut += string(nDto.DecimalSeparator)
		nDto.NumStrOut += string(nDto.AbsFracRunes)
	}

	return nDto.IsNumStrDtoValid(nDto, "ResetNumStrOut()")
}

// SubtractNumStrs - Subtracts the numeric values represented by two NumStrDto
// objects.
func (nDto *NumStrDto) SubtractNumStrs(n1Dto, n2Dto NumStrDto) (NumStrDto, error) {

	n1NumDto, n2NumDto, compare, isReversed, err := nDto.FormatForMathOps(n1Dto, n2Dto)

	if err != nil {
		return NumStrDto{}, fmt.Errorf("SubtractNumStrs() - Error from nDto.FormatForMathOps(n1Dto, n2Dto). Error= %v", err)
	}

	if compare == 0 {
		return nDto.GetZeroNumStr(n1NumDto.Precision), nil
	}

	newSignVal := n1NumDto.SignVal
	precision := n1NumDto.Precision

	if n1NumDto.SignVal != n2NumDto.SignVal {

		err = n1NumDto.SetSignValue(1)

		if err != nil {
			return NumStrDto{}, fmt.Errorf("SubtractNumStrs() - Error from n1NumDto.SetSignValue(1). Error= %v", err)
		}

		err = n2NumDto.SetSignValue(1)

		if err != nil {
			return NumStrDto{}, fmt.Errorf("SubtractNumStrs() - Error from n2NumDto.SetSignValue(1). Error= %v", err)
		}

		nOutDto, err := nDto.AddNumStrs(n1NumDto, n2NumDto)

		if err != nil {
			return NumStrDto{}, fmt.Errorf("SubtractNumStrs() - Error from nDto.AddNumStrs(n1NumDto, n2NumDto). Error= %v", err)
		}

		nOutDto.SetSignValue(newSignVal)

		return nOutDto, nil
	}

	// Change sign for subtraction
	newSignVal = n1NumDto.SignVal

	if isReversed {
		newSignVal = newSignVal * -1
	}

	lenN1AllRunes := len(n1NumDto.AbsAllNumRunes)

	n1IntAry := make([]int, lenN1AllRunes)
	n2IntAry := make([]int, lenN1AllRunes)
	n3IntAry := make([]int, lenN1AllRunes)

	for i := 0; i < lenN1AllRunes; i++ {

		n1IntAry[i] = int(n1NumDto.AbsAllNumRunes[i]) - 48
		n2IntAry[i] = int(n2NumDto.AbsAllNumRunes[i]) - 48

	}

	carry := 0
	n1 := 0
	n2 := 0
	n3 := 0
	// Main Subtraction Routine
	for j := lenN1AllRunes - 1; j >= 0; j-- {

		n1 = n1IntAry[j]
		n2 = n2IntAry[j]
		n3 = 0

		if n1-carry-n2 < 0 {
			n1 += 10
			n3 = n1 - n2 - carry
			carry = 1
		} else {
			n3 = n1 - n2 - carry
			carry = 0
		}

		n3IntAry[j] = n3

	}

	nOutDto, err := nDto.FindIntArraySignificantDigitLimits(n3IntAry, precision, newSignVal)

	if err != nil {
		return NumStrDto{}, fmt.Errorf("SubtractNumStrs() - Error from final nDto.FindIntArraySignificantDigitLimits(n3IntAry, precision, newSignVal). precision='%v' newSignVal='%v' Error= %v", precision, newSignVal, err)
	}

	return nOutDto, nil
}
