package common

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

/*
	The source code repository for decimal.go is located at:
			https://github.com/MikeAustin71/decimalnum.git
*/

// Decimal type is used to perform math operations to
// achieve a high degree of accuracy and uniformity when
// dealing with fractional numbers containing digits to the
// right of the decimal place.

type Decimal struct {
	isValid               bool
	signVal               int
	precision             uint
	scaleFactor           *big.Int
	thousandsSeparator    rune // defaults to ','
	decimalSeparator      rune // defaults to '.'
	currencySymbol        rune // defaults to '$'
	numStr                string
	signedAllDigitsBigInt *big.Int
}

// Add - Adds the value of the current Decimal to that of
// the incoming Decimal and returns in the result in a
// Decimal Type.
func (dec *Decimal) Add(d2 Decimal) (Decimal, error) {

	if !dec.isValid {
		return Decimal{}, errors.New("This Decimal value invalid")
	}

	if !d2.isValid {
		return Decimal{}, errors.New("Incoming Decimal invalid")
	}

	base10 := big.NewInt(10)

	if dec.precision == d2.precision {
		s3Val := big.NewInt(0).Add(dec.signedAllDigitsBigInt, d2.signedAllDigitsBigInt)

		s3Text := s3Val.Text(10)

		d4, err := dec.NumStrPrecisionToDecimal(s3Text, dec.precision, true)

		if err != nil {
			return Decimal{}, fmt.Errorf("Add() received error from nDto.NumStrPrecisionToDecimal(s3Text, nDto.precision) s3Text='%v' nDto.precision=%v. Error= %v", s3Text, dec.precision, err)
		}

		return d4, nil
	}

	if d2.precision > dec.precision {
		deltaPrecision := big.NewInt(int64(d2.precision - dec.precision))
		deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
		s1 := big.NewInt(0).Mul(dec.signedAllDigitsBigInt, deltaPrecisionScale)
		s3 := big.NewInt(0).Add(s1, d2.signedAllDigitsBigInt)
		s3Text := s3.Text(10)

		d4, err := dec.NumStrPrecisionToDecimal(s3Text, d2.precision, true)

		if err != nil {
			return Decimal{}, fmt.Errorf("Add() received error from nDto.NumStrPrecisionToDecimal(s3Text, nDto.precision) s3Text='%v' nDto.precision=%v. Error= %v", s3Text, dec.precision, err)
		}

		return d4, nil

	}

	// must be nDto.precison > d2.precision
	idp := int(dec.precision) - int(d2.precision)
	i64DeltaPrecision := int64(idp)
	deltaPrecision := big.NewInt(i64DeltaPrecision)
	deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
	s1 := big.NewInt(0).Mul(d2.signedAllDigitsBigInt, deltaPrecisionScale)
	s3 := big.NewInt(0).Add(s1, dec.signedAllDigitsBigInt)
	s3Text := s3.Text(10)

	d4, err := dec.NumStrPrecisionToDecimal(s3Text, dec.precision, true)

	if err != nil {
		return Decimal{}, fmt.Errorf("Add() received error from nDto.NumStrPrecisionToDecimal(s3Text, nDto.precision) s3Text='%v' nDto.precision=%v. Error= %v", s3Text, dec.precision, err)
	}

	return d4, nil
}

// AddToThis - adds the value of the incoming Decimal to that
// of the current Decimal object. The new total is retained
// in the current Decimal object.
func (dec *Decimal) AddToThis(d2 Decimal) error {

	dec3, err := dec.Add(d2)

	if err != nil {
		return fmt.Errorf("AddToThis() - Error received from Add(). Error= %v", err)
	}

	dec.CopyIn(dec3)

	return nil
}

// AddToThisArray - Receives an array of Decimal objects and
// adds them to the current Decimal value.
func (dec *Decimal) AddToThisArray(decs []Decimal) error {

	for _, dx := range decs {

		dec3, err := dec.Add(dx)

		if err != nil {
			return fmt.Errorf("AddToThisArray() - Error received from Add(). Error= %v", err)
		}

		dec.CopyIn(dec3)

	}

	return nil
}

// AddToThisMultiple - Receives multiple Decimal objects and
// adds them to the current Decimal value.
func (dec *Decimal) AddToThisMultiple(decs ...Decimal) error {

	for _, dx := range decs {

		dec3, err := dec.Add(dx)

		if err != nil {
			return fmt.Errorf("AddToThisMultiple() - Error received from Add(). Error= %v", err)
		}

		dec.CopyIn(dec3)

	}

	return nil
}

// AllDigitsNumStr - parses the incoming string and returns
// a pure number string consisting of all numeric digits. No
// sign characters, decimals or thousands separators are returned.
// The returned value is the absolute numeric value extracted from
// the 'numStr' input parameter.
func (dec *Decimal) AllDigitsNumStr(numStr string) (string, error) {

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(numStr)

	if err != nil {
		return "", fmt.Errorf("AllDigitsNumStr() - nDto.ParseNumStr(numStr) returned an error. numStr= '%v' Error= %v", numStr, err)
	}

	return string(nDto.AbsAllNumRunes), nil
}

// NumStrToDecimal - Creates a Decimal type from a number
// string.
func (dec *Decimal) NumStrToDecimal(str string) (Decimal, error) {

	if dec.decimalSeparator == 0 {
		dec.decimalSeparator = '.'
	}

	if dec.thousandsSeparator == 0 {
		dec.thousandsSeparator = ','
	}

	if dec.currencySymbol == 0 {
		dec.currencySymbol = '$'
	}

	n1 := NumStrDto{}.New()
	n1.ThousandsSeparator = dec.thousandsSeparator
	n1.DecimalSeparator = dec.decimalSeparator
	n1.CurrencySymbol = dec.currencySymbol

	nDto, err := n1.ParseNumStr(str)

	if err != nil {
		return Decimal{}, fmt.Errorf("NumStrToDecimal() received errror from nDto.ParseNumStr(str). str='%v' Error= %v", str, err)
	}

	d2, err := dec.MakeDecimalFromNumStrDto(nDto)

	return d2, nil

}

// CopyIn - Receives an incoming Decimal object
// as an input parameter and sets the Current Decimal
// equal to that of the incoming Decimal object.
func (dec *Decimal) CopyIn(d2 Decimal) {
	dec.Empty()

	dec.isValid = d2.isValid
	dec.signVal = d2.signVal
	dec.precision = d2.precision
	dec.scaleFactor = big.NewInt(0).Set(d2.scaleFactor)
	dec.numStr = d2.numStr
	dec.signedAllDigitsBigInt = big.NewInt(0).Set(d2.signedAllDigitsBigInt)
	dec.currencySymbol = d2.currencySymbol
	dec.thousandsSeparator = d2.thousandsSeparator
	dec.decimalSeparator = d2.decimalSeparator

}

func (dec *Decimal) CopyOut() Decimal {
	d2 := Decimal{}.New()
	d2.isValid = dec.isValid
	d2.signVal = dec.signVal
	d2.precision = dec.precision
	d2.scaleFactor = big.NewInt(0).Set(dec.scaleFactor)
	d2.numStr = dec.numStr
	d2.signedAllDigitsBigInt = big.NewInt(0).Set(dec.signedAllDigitsBigInt)
	d2.currencySymbol = dec.currencySymbol
	d2.thousandsSeparator = dec.thousandsSeparator
	d2.decimalSeparator = dec.decimalSeparator

	return d2
}

// Divide - Divides the current decimal value by the input
// parameter 'divisor'.
func (dec *Decimal) Divide(divisor Decimal, precision int) (Decimal, error) {

	base10 := big.NewInt(10)
	sDividend1 := big.NewInt(0)
	sDivisor2 := big.NewInt(0)

	if divisor.precision > dec.precision {
		deltaPrecision := big.NewInt(int64(divisor.precision - dec.precision))
		deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
		sDividend1 = big.NewInt(0).Mul(dec.signedAllDigitsBigInt, deltaPrecisionScale)
		sDivisor2 = big.NewInt(0).Set(divisor.signedAllDigitsBigInt)

	} else if dec.precision > divisor.precision {
		deltaPrecision := big.NewInt(int64(dec.precision - divisor.precision))
		deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
		sDividend1 = big.NewInt(0).Set(dec.signedAllDigitsBigInt)
		sDivisor2 = big.NewInt(0).Mul(divisor.signedAllDigitsBigInt, deltaPrecisionScale)

	} else {
		sDividend1 = big.NewInt(0).Set(dec.signedAllDigitsBigInt)
		sDivisor2 = big.NewInt(0).Set(divisor.signedAllDigitsBigInt)

	}

	rDividend := big.NewRat(1, 1).SetInt(sDividend1)
	rDivisor := big.NewRat(1, 1).SetInt(sDivisor2)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)
	numStr := rQuotient.FloatString(precision)

	return dec.NumStrToDecimal(numStr)

}

// Empty - Sets all values of the current Decimal's
// fields to their 'zero' values.
func (dec *Decimal) Empty() {
	dec.isValid = true
	dec.signVal = 1
	dec.precision = 0
	dec.scaleFactor = big.NewInt(int64(1))
	dec.numStr = "0"
	dec.signedAllDigitsBigInt = big.NewInt(0)
	dec.currencySymbol = '$'
	dec.thousandsSeparator = ','
	dec.decimalSeparator = '.'
}

// GetAbsoluteValue - returns the absolute value of the
// decimal expressed as a string. If the decimal value is
// '-123.456', this method will return '123.456'.
func (dec *Decimal) GetAbsoluteValue() (Decimal, error) {

	if !dec.isValid {
		return Decimal{}, errors.New("This Decimal data is corrupted. Please re-initialize")
	}

	d2 := dec.CopyOut()

	if d2.signVal == 1 {
		return d2, nil
	}

	d2.signVal = 1
	d2.signedAllDigitsBigInt = big.NewInt(0).Abs(dec.signedAllDigitsBigInt)

	rDividend := big.NewRat(1, 1).SetInt(d2.signedAllDigitsBigInt)
	rDivisor := big.NewRat(1, 1).SetInt(d2.scaleFactor)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)

	d2.numStr = rQuotient.FloatString(int(d2.precision))

	return d2, nil

}

// GetAbsoluteAllDigitsStr - Returns the absolute value of the Decimal integer.
// Fractions are not returned, only the string of signed numeric digits which
// constitutes the entire number. In other words, if the value of the decimal
// is '-123.456', this method will return '123456.
func (dec *Decimal) GetAbsoluteAllDigitsStr() (string, error) {

	if !dec.isValid {
		return "", errors.New("The Decimal data is corrupted. Please re-initialize!")
	}

	absVal := big.NewInt(0).Abs(dec.signedAllDigitsBigInt)

	return absVal.String(), nil
}

// GetBigFloat - returns big Float representation of the Decimal Value.
func (dec *Decimal) GetBigFloat() (*big.Float, error) {

	if !dec.isValid {
		return big.NewFloat(0.0), errors.New("The Decimal data is corrupted. Please re-initialize!")
	}

	bf, status := big.NewFloat(0.0).SetString(dec.numStr)

	if !status {
		return big.NewFloat(0.0), fmt.Errorf("SetString Failed. NumStr= %v", dec.numStr)
	}

	return bf, nil
}

// GetBigFloatString - returns a signed number string which is accurate out
// to a very large number of decimal places (40+ decimal places). In contrast,
// the *big.Float returned by GetBigFloat() is only accurate to about
// 16-17 decimal places. func (nDto *Decimal)
func (dec *Decimal) GetBigFloatString(precision uint) (string, error) {

	if !dec.isValid {
		return "", errors.New("The Decimal data is corrupted. Please re-initialize!")
	}

	rDividend := big.NewRat(1, 1).SetInt(dec.signedAllDigitsBigInt)
	rDivisor := big.NewRat(1, 1).SetInt(dec.scaleFactor)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)

	return rQuotient.FloatString(int(precision)), nil

}

// GetCurrencySymbol - Returns the Decimal's current
// value for Currency Symbol
func (dec *Decimal) GetCurrencySymbol() rune {
	return dec.currencySymbol
}

// GetDecimalSeparator - returns the Decimal's current
// value for Decimal Separator (i.e. '.')
func (dec *Decimal) GetDecimalSeparator() rune {
	return dec.decimalSeparator
}

// GetFloat32 - Returns the current value of the Decimal as a
// float32. There may be a loss of accuracy during this conversion.
// The return parameter big.Accuracy will describe the level of
// accuracy provided.
//
// See big.Accuracy:
//   Below Accuracy = -1
//   Exact Accuracy = 0
//   Above Accuracy = +1
func (dec *Decimal) GetFloat32() (float32, big.Accuracy, error) {

	if !dec.isValid {
		return float32(0), big.Accuracy(0), errors.New("The Decimal data is corrupted. Please re-initialize!")
	}

	bf, status := big.NewFloat(0.0).SetString(dec.numStr)

	if !status {
		return float32(0.0), big.Accuracy(0), fmt.Errorf("SetString Failed. NumStr= %v", dec.numStr)
	}

	f32, accuracy := bf.Float32()

	return f32, accuracy, nil

}

// GetFloat64 - Returns the current value of the Decimal as a
// float64. There may be a loss of accuracy during this conversion.
// The return parameter big.Accuracy will describe the level of
// accuracy provided.
//
// See big.Accuracy:
//   Below Accuracy = -1
//   Exact Accuracy = 0
//   Above Accuracy = +1
//
func (dec *Decimal) GetFloat64() (float64, big.Accuracy, error) {

	if !dec.isValid {
		return float64(0), big.Accuracy(0), errors.New("The Decimal data is corrupted. Please re-initialize!")
	}

	bf, status := big.NewFloat(0.0).SetString(dec.numStr)

	if !status {
		return float64(0.0), big.Accuracy(0), fmt.Errorf("SetString Failed. NumStr= %v", dec.numStr)
	}

	f64, accuracy := bf.Float64()

	return f64, accuracy, nil
}

// GetIntAry - Returns an IntAry structure initialized
// to the value of the current 'Decimal' object.
func (dec *Decimal) GetIntAry() (IntAry, error) {

	ia, err := IntAry{}.NewBigInt(dec.signedAllDigitsBigInt, dec.precision)

	if err != nil {
		return IntAry{}.New(), fmt.Errorf("GetIntAry() - Received error from IntAry{}.NewBigInt(dec.signedAllDigitsBigInt, dec.precision). Error= %v", err)
	}

	return ia, nil
}

// GetNumStr - Returns the internal value of the Decimal
// expressed as a signed numeric string. Precision is
// controlled by the Decimal's precision setting.
func (dec *Decimal) GetNumStr() string {

	return dec.numStr

}

// GetNumStrDto - returns a NumStrDto structure initialized
// to the value of the current Decimal object.
func (dec *Decimal) GetNumStrDto() (NumStrDto, error) {

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(dec.numStr)

	if err != nil {
		return NumStrDto{}, fmt.Errorf("GetNumStrDto() - Error returned from NumStrDto.ParseNumStr(dec.numStr). dec.numStr= '%v' Error= %v", dec.numStr, err)
	}

	return nDto, nil
}

// GetNthRoot - Calculates the nth root of the current Decimal value.
//
// Input parameters:
//
//  nthRoot uint - Nth root specifies the root which will be calculated for parameter, 'num'.
// 									Example, square root, cube root, 4th root, 9th root etc.
//
// maxPrecision uint -  Specifies the number of decimals to the right of the decimal place to
// 											which the Nth root will be calculated.
//
// The calculation result is returned as an IntAry object.
//
// Note: A negative Decimal value with an even nthRoot will generate an error.
func (dec *Decimal) GetNthRoot(nthRoot, maxPrecision uint) (Decimal, error) {

	if !dec.isValid {
		return Decimal{}.New(), errors.New("GetNthRoot() - The current Decimal object is INVALID. Please re-initialize.")
	}

	ia, err := dec.GetIntAry()

	if err != nil {
		return Decimal{}.New(), fmt.Errorf("GetNthRoot() - Received error from dec.GetIntAry(). Error= %v", err)
	}

	nrt := NthRootOp{}

	iaNth, err := nrt.GetNthRootIntAry(&ia, nthRoot, maxPrecision)

	if err != nil {
		return Decimal{}.New(), fmt.Errorf("GetNthRoot() - Received error from  NthRootOp.GetNthRootIntAry(). Error= %v", err)
	}

	dOut, err := dec.MakeDecimalFromIntAry(&iaNth)

	if err != nil {
		return Decimal{}.New(), fmt.Errorf("GetNthRoot() - Received error from  dec.MakeDecimalFromIntAry(&iaNth). Error= %v", err)
	}

	return dOut, nil

}

// GetPrecision - returns the Decimal's current precision
// value. The Decimal structure maintains precision as an
// unsigned integer.
func (dec *Decimal) GetPrecision() int {

	return int(dec.precision)

}

// GetRational - returns a big Rational number type which
// is capable of very high accuracy.
//
// The returned *big.Rat number is initialized to the
// current value of the Decimal object.
func (dec *Decimal) GetRational() (*big.Rat, error) {

	rDividend := big.NewRat(1, 1).SetInt(dec.signedAllDigitsBigInt)
	rDivisor := big.NewRat(1, 1).SetInt(dec.scaleFactor)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)

	return rQuotient, nil
}

// GetSign - Returns the sign of the current
// Decimal Value. Return values are one of two
// integers: +1 or -1.
func (dec *Decimal) GetSign() int {

	return dec.signVal
}

// GetRelevantPrecision - Returns an unsigned integer representing
// the number of decimal places to the right of the decimal which
// are occupied by non-zero values.
//
// Example: Value = 1.640700000.  The number of relevant decimal
// places to the right of the decimal is '4'. In the case of an
// integer number the relevant precision is zero ('0') because
// there are no digits to the right of the decimal.
func (dec *Decimal) GetRelevantPrecision() uint {

	nRunes := []rune(dec.numStr)

	lnRunes := len(nRunes)

	if lnRunes == 0 {
		return 0
	}

	isDecimal := false
	decSep := rune('.')
	decPos := -1
	lastDecPos := -1
	if dec.decimalSeparator != 0 {
		decSep = dec.decimalSeparator
	}

	for i := 0; i < lnRunes; i++ {

		if nRunes[i] == decSep {
			isDecimal = true
			decPos = i
			continue
		}

		if isDecimal {

			if nRunes[i] >= '1' && nRunes[i] <= '9' {
				lastDecPos = i
			}

		}

	}

	if decPos == -1 {
		return 0
	}

	if lastDecPos == -1 {
		return 1
	}

	return uint(lastDecPos - decPos)
}

// GetScaleVal - Returns the scale value associated with this decimal value. The
// scale value is expressed as 10 to an exponent. Example 10^2 == 100.
//
// The return scale value is of type big integer (*big.Int)
//
func (dec *Decimal) GetScaleVal() (*big.Int, error) {

	if !dec.isValid {
		return big.NewInt(0), errors.New("The Decimal data is corrupted. Please re-initialize!")

	}

	return dec.scaleFactor, nil
}

// GetSignedAllDigitsStr - Returns the Decimal's internal
// Signed All Digits Integer Value expressed as a string.
// No Fractional Digits are included, this is a signed
// integer number string. Example: The value '-123.456'
// would be returned as '-123456'
func (dec *Decimal) GetSignedAllDigitsStr() (string, error) {

	if !dec.isValid {
		return "", errors.New("The Decimal data is corrupted. Please re-initialize")

	}

	return dec.signedAllDigitsBigInt.String(), nil
}

// GetSignedAllDigitsVal - returns the Decimal value expressed as an
// integer value using type *big.Int. No factional values are included.
// For example, the value '-123.456' would be returned as the integer
// value '-123456'.  To compute the precise value of the Decimal, this
// integer value would need to be divided by the 'Precision Value'. See
// GetScaleVal() below.
func (dec *Decimal) GetSignedAllDigitsVal() (*big.Int, error) {

	if !dec.isValid {
		return big.NewInt(0), errors.New("The Decimal data is corrupted. Please re-initialize")

	}

	return dec.signedAllDigitsBigInt, nil
}

// GetSquareRoot - Returns a Decimal object equal to the square root
// of the current Decimal value.
//
// Note: If the current Decimal value is a negative value, an error will
// be generated. You cannot take the square root of a negative number.
func (dec *Decimal) GetSquareRoot(maxPrecision uint) (Decimal, error) {

	if !dec.isValid {
		return Decimal{}.New(), errors.New("GetSquareRoot() - The current Decimal object is INVALID. Please re-initialize.")
	}

	ia, err := dec.GetIntAry()

	if err != nil {
		return Decimal{}.New(), fmt.Errorf("GetSquareRoot() - Received error from dec.GetIntAry(). Error= %v", err)
	}

	nrt := NthRootOp{}

	iaSq, err := nrt.GetSquareRootIntAry(&ia, maxPrecision)

	if err != nil {
		return Decimal{}.New(), fmt.Errorf("GetSquareRoot() - Received error from  NthRootOp.GetSquareRootIntAry(). Error= %v", err)
	}

	dOut, err := dec.MakeDecimalFromIntAry(&iaSq)

	if err != nil {
		return Decimal{}.New(), fmt.Errorf("GetSquareRoot() - Received error from  dec.MakeDecimalFromIntAry(&iaSq). Error= %v", err)
	}

	return dOut, nil

}

// GetThousandsSeparator - Gets the current value of
// the Thousands Separator for the current Decimal
// object.
//
// For U.S.A. - The thousands separator is a the comma (',')
//
func (dec *Decimal) GetThousandsSeparator() rune {

	return dec.thousandsSeparator
}

// GetIsValid - returns a boolean indicating
// the current state of the Decimal information.
func (dec *Decimal) GetIsValid() bool {

	return dec.isValid
}

// Inverse - Returns the inverse of the current
// Decimal.  Inverse = 1/Current Decimal Value.
//
// Input Parameters:
// maxPrecision int - determines the number of digits
// 			to the right of the decimal point in the result.
// 			if maxPrecision is less than zero, a default precision
//      of 500 digits to the right of the decimal point is
//      applied.
//
func (dec *Decimal) Inverse(maxPrecision int) (Decimal, error) {

	if maxPrecision < 0 {
		maxPrecision = 500
	}

	rDividend := big.NewRat(1, 1).SetInt(dec.signedAllDigitsBigInt)
	rDivisor := big.NewRat(1, 1).SetInt(dec.scaleFactor)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)
	rInverse := big.NewRat(1, 1).Inv(rQuotient)

	numStr := rInverse.FloatString(maxPrecision)

	d2, err := dec.NumStrToDecimal(numStr)

	if err != nil {
		return Decimal{}, fmt.Errorf("Inverse() received error from nDto.NumStrToDecimal(numStr). Error= %v", err)
	}

	rP := d2.GetRelevantPrecision()

	err = d2.SetPrecisionTrunc(rP)

	if err != nil {
		return Decimal{}, fmt.Errorf("Inverse() received error from d2.SetPrecisionTrunc(rP). Error= %v", err)
	}

	return d2, nil
}

// IsFraction - returns a boolean value. If 'true',
// it signals that the Decimal has digits to the
// right of the decimal place. If 'false', it
// signals that the decimal value is an integer with
// no digits to the right of the decimal place.
func (dec *Decimal) IsFraction() (bool, error) {

	if !dec.isValid {
		return false, errors.New("The Decimal data is corrupted. Please re-initialize!")
	}

	if dec.precision != 0 {
		return true, nil
	}

	return false, nil
}

// MakeDecimalBigIntPrecision - This method receives a *big.Int and a precision value which
// are used to construct and return a Decimal Type. The value of 'precision' determines the
// number of Big Int digits which will be placed to the right of the decimal place.
func (dec *Decimal) MakeDecimalBigIntPrecision(iBig *big.Int, precision uint) (Decimal, error) {
	d2 := Decimal{}.New()

	baseZero := big.NewInt(0)
	base10 := big.NewInt(10)
	d2.precision = precision
	exponent := big.NewInt(int64(d2.precision))
	d2.scaleFactor = big.NewInt(0).Exp(base10, exponent, nil)
	d2.signedAllDigitsBigInt = big.NewInt(0).Set(iBig)

	if d2.signedAllDigitsBigInt.Cmp(baseZero) < 0 {
		d2.signVal = -1
	}

	rDividend := big.NewRat(1, 1).SetInt(d2.signedAllDigitsBigInt)
	rDivisor := big.NewRat(1, 1).SetInt(d2.scaleFactor)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)

	d2.numStr = rQuotient.FloatString(int(d2.precision))
	d2.thousandsSeparator = dec.thousandsSeparator
	d2.decimalSeparator = dec.decimalSeparator
	d2.currencySymbol = dec.currencySymbol
	d2.isValid = true

	return d2, nil
}

// MakeDecimalFromNumStrDto - generates a Decimal Type based on string information
// provided by the 'nDto' NumStrDto parameter.
func (dec *Decimal) MakeDecimalFromNumStrDto(nDto NumStrDto) (Decimal, error) {

	if len(nDto.AbsAllNumRunes) == 0 {

		d2 := Decimal{}.New()

		return d2, nil
	}

	d2 := Decimal{}.New()
	d2.signVal = nDto.SignVal
	d2.numStr = nDto.NumStrOut
	d2.precision = nDto.Precision
	d2.currencySymbol = dec.currencySymbol
	d2.decimalSeparator = nDto.DecimalSeparator
	d2.thousandsSeparator = dec.thousandsSeparator

	var err error

	d2.signedAllDigitsBigInt, err = nDto.GetSignedBigInt()

	if err != nil {
		return Decimal{}.New(), fmt.Errorf("MakeDecimalFromNumStrDto() - Error from nDto.GetSignedBigInt(). Error= %v", err)
	}

	d2.scaleFactor, err = nDto.GetScaleFactor()

	if err != nil {
		return Decimal{}.New(), fmt.Errorf("MakeDecimalFromNumStrDto() - Error from nDto.GetScaleFactor(). Error= %v", err)
	}

	d2.isValid = true

	return d2, nil
}

// MakeDecimalFromNumStrDto - generates a Decimal Type based on string information
// provided by the 'ia' *IntAry input parameter.
func (dec *Decimal) MakeDecimalFromIntAry(ia *IntAry) (Decimal, error) {

	var err error

	ia.SetInternalFlags()

	err = ia.IsIntAryValid("MakeDecimalFromIntAry() - ")

	if err != nil {
		return Decimal{}.New(), fmt.Errorf("IntAry Invalid - Error: %v", err)
	}

	iaStats := ia.GetIntAryStats()

	if iaStats.IsZeroValue {

		d2 := Decimal{}.New()

		return d2, nil
	}

	d2 := Decimal{}.New()
	d2.signVal = ia.GetSign()
	d2.numStr = ia.GetNumStr()
	d2.signedAllDigitsBigInt = ia.GetBigInt()
	d2.precision = uint(ia.GetPrecision())
	d2.currencySymbol = dec.currencySymbol
	d2.decimalSeparator = dec.decimalSeparator
	d2.thousandsSeparator = dec.thousandsSeparator

	d2.scaleFactor, err = ia.GetScaleFactor()

	if err != nil {
		return Decimal{}.New(), fmt.Errorf("MakeDecimalFromIntAry() - Error from ia.GetScaleFactor(). Error= %v", err)
	}

	d2.isValid = true

	return d2, nil
}

// Mul - Multiplies the incoming Decimal value, by the
// current Decimal Value and returns the result in a
// Decimal Type.
func (dec *Decimal) Mul(d2 Decimal) (Decimal, error) {

	if !dec.isValid {
		return Decimal{}, errors.New("This Decimal Value Invalid!")
	}

	if !d2.isValid {
		return Decimal{}, errors.New("Incoming Decimal Invalid!")
	}

	s3Val := big.NewInt(0).Mul(dec.signedAllDigitsBigInt, d2.signedAllDigitsBigInt)

	newPrecision := dec.precision + d2.precision

	return dec.NumStrPrecisionToDecimal(s3Val.String(), newPrecision, true)
}

// MulTotal - Multiplies the value of the incoming
// Decimal type by the value of the current or receiving
// Decimal type. The result of the multiplication is
// stored in the current Decimal type.
func (dec *Decimal) MulTotal(d2 Decimal) error {

	d3, err := dec.Mul(d2)

	if err != nil {
		return fmt.Errorf("MulTotal() Error from dec.Mul(d2). Error= %v", err)
	}

	dec.CopyIn(d3)

	return nil
}

// New - Creates and returns a Decimal type.
// The Decimal value is initialized to zero.
//
// Example Usage:
//   d := Decimal{}.New()
//
// This is the recommended procedure for creating
// a Decimal type.
func (dec Decimal) New() Decimal {

	d := Decimal{}
	d.Empty()

	return d

}

// NewPtr - Creates and returns a pointer to
// a new Decimal type. Can be used for initializing
// a Decimal Type and calling pointer methods in
// one statement.
//
// Example:
// d1, err := Decimal{}.NewPtr().NumStrPrecisionToDecimal(inStr, precision,true)
//
func (dec Decimal) NewPtr() *Decimal {
	d := Decimal{}
	d.Empty()
	return &d
}

// NewBigInt - Returns a Decimal type based on Big Int and
// precision input parameters. If an error is encountered, it
// will trigger a panic condition.
//
// The 'NewBigInt' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation
// and initialization in one step.
//
// Example: Decimal{}.NewBigInt(bigI, precision)
func (dec Decimal) NewBigInt(bigI *big.Int, precision uint) Decimal {

	d2 := Decimal{}.New()
	err := d2.SetBigInt(bigI, precision)

	if err != nil {
		panic(err)
	}

	return d2

}

// NewInt - Returns a Decimal type based on int and precision input
// parameters. If an error is encountered, it will trigger a panic
// condition.
//
// The 'NewInt' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation and
// initialization in one step.
//
// Example: Decimal{}.NewInt(123456, 3)
func (dec Decimal) NewInt(iNum int, precision uint) Decimal {
	d2 := Decimal{}.New()

	err := d2.SetInt(iNum, precision)

	if err != nil {
		panic(err)
	}

	return d2

}

// NewI64 - Returns a Decimal type based on int64 and precision
// input parameters. If an error is encountered, it will trigger
// a panic condition.
//
// The 'NewI64' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation
// and initialization in one step.
//
// Example: Decimal{}.NewI64(i64, precision)
func (dec Decimal) NewI64(i64 int64, precision uint) Decimal {
	d2 := Decimal{}.New()
	err := d2.SetInt64(i64, precision)

	if err != nil {
		panic(err)
	}

	return d2
}

// NewNumStr - Returns a Decimal type based on a number string
// input parameter. If an error is encountered, it will trigger
// a panic condition.
//
// The 'NewNumStr' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation
// and initialization in one step.
//
// Example: Decimal{}.NewNumStr('123.456')
func (dec Decimal) NewNumStr(numStr string) Decimal {
	d2 := Decimal{}.New()

	err := d2.SetNumStr(numStr)

	if err != nil {
		panic(err)
	}

	return d2

}

// NewNumStrPrecision - Returns a Decimal type based on a number string
// and a precision value received as input parameters. If an
// error is encountered, it will trigger a panic condition.
//
// The 'NewNumStrPrecision' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation
// and initialization in one step.
//
// Example: Decimal{}.NewNumStrPrecision('123456', 3)
func (dec Decimal) NewNumStrPrecision(numStr string, precision uint, roundResult bool) (Decimal, error) {

	n1, err := NumStrDto{}.NewPtr().ParseNumStr(numStr)

	if err != nil {
		return Decimal{}, fmt.Errorf("NewNumStrPrecision() Error returned from  NumStrDto.ParseNumStr(numStr). numStr= '%v' Error= %v", numStr, err)
	}

	n2, err := NumStrDto{}.NewPtr().SetPrecision(n1.NumStrOut, precision, roundResult)

	d2, err := dec.MakeDecimalFromNumStrDto(n2)

	if err != nil {
		return Decimal{}, fmt.Errorf("NewNumStrPrecision() Error returned from  Decimal.MakeDecimalFromNumStrDto(n2). dec.numStr= '%v' Error= %v", dec.numStr, err)
	}

	return d2, nil
}

// TODO More testing required NumStrPrecisionToDecimal
// NumStrPrecisionToDecimal - receives a number string and a
// precision value as parameters. This method creates a Decimal
// Type containing the converted numeric and returns it.
// For example, if passed the string ('str') '123456' and a precision
// value of '3', the resulting Decimal value would be 123.456.
//
// Example Usage:
// d := Decimal{}.New()
// d2, err := d.NumStrPrecisionToDecimal("123456", 3)
// d2 is Now Equal to 123.456
func (dec *Decimal) NumStrPrecisionToDecimal(str string, precision uint, roundResult bool) (Decimal, error) {

	if dec.decimalSeparator == 0 {
		dec.decimalSeparator = '.'
	}

	if dec.thousandsSeparator == 0 {
		dec.thousandsSeparator = ','
	}

	if dec.currencySymbol == 0 {
		dec.currencySymbol = '$'
	}

	n1 := NumStrDto{}.New()
	n1.ThousandsSeparator = dec.thousandsSeparator
	n1.DecimalSeparator = dec.decimalSeparator
	n1.CurrencySymbol = dec.currencySymbol

	nDto, err := n1.ScaleAbsoluteValStr(str, precision, true)

	if err != nil {
		return Decimal{}, fmt.Errorf("NumStrPrecisionToDecimal() Error received from nDto.ScaleAbsoluteValStr(nDto.NumStrOut, precision, true). nDto.NumStrOut='%v' precision='%v' Error= %v", nDto.NumStrOut, precision, err)
	}

	d2, err := dec.MakeDecimalFromNumStrDto(nDto)

	if err != nil {
		return Decimal{}, fmt.Errorf("NumStrPrecisionToDecimal() Error received from dec.MakeDecimalFromNumStrDto(nDto, precision). Error= %v", err)
	}

	return d2, nil

	/*
		lAbsAllRunes := len(nDto.AbsAllNumRunes)

		if lAbsAllRunes == 0 {

			outStr := "0"

			d2, err := dec.MakeDecimalFromNumStrDto(nDto, precision)

			if err != nil {
				return Decimal{}.New(), fmt.Errorf("NumStrPrecisionToDecimal() received error from d2.SetNumStr(outStr). outStr='%v' Error= %v", outStr, err)
			}

			return d2, nil
		}

		absAllRunes := []rune{}
		lAbsFracRunes := len(nDto.AbsFracRunes)
		deltaPrecision := 0

		if lAbsFracRunes > 0 {
			// This number string includes fractional digits to the right of the
			// decimal place.
			if int(precision) > lAbsFracRunes {

				deltaPrecision = int(precision) - lAbsFracRunes

				for j := 0; j < lAbsAllRunes; j++ {
					absAllRunes = append(absAllRunes, nDto.AbsAllNumRunes[j])
				}

				for i := 0; i < deltaPrecision; i++ {
					absAllRunes = append(absAllRunes, '0')
				}

			} else if int(precision) < lAbsFracRunes {

				deltaPrecision = lAbsFracRunes - int(precision)
				nDto.AbsFracRunes = nDto.AbsFracRunes[0:precision]
				lAbsAllRunes = lAbsAllRunes - deltaPrecision

				for j := 0; j < lAbsAllRunes; j++ {
					absAllRunes = append(absAllRunes, nDto.AbsAllNumRunes[j])
				}

			} else {
				// precision == lAbsFracRunes
				for j := 0; j < lAbsAllRunes; j++ {
					absAllRunes = append(absAllRunes, nDto.AbsAllNumRunes[j])
				}
			}

		} else {
			// There are no fractional digits. This is a pure, integer number string.
			if int(precision) > lAbsAllRunes {

				deltaPrecision = (int(precision) - lAbsAllRunes) + 1

				for i := 0; i < deltaPrecision; i++ {
					absAllRunes = append(absAllRunes, '0')
				}

				for j := 0; j < lAbsAllRunes; j++ {
					absAllRunes = append(absAllRunes, nDto.AbsAllNumRunes[j])
				}

			} else {

				for j := 0; j < lAbsAllRunes; j++ {
					absAllRunes = append(absAllRunes, nDto.AbsAllNumRunes[j])
				}

			}

		}

		nDto := NumStrDto{}.NewPtr().MakeAllNums(nDto.SignVal, precision, absAllRunes, dec.decimalSeparator)

		d2, err := dec.MakeDecimalFromNumStrDto(nDto, precision)

		if err != nil {
			return Decimal{}, fmt.Errorf("NumStrPrecisionToDecimal() error received from err:= d2.MakeDecimalFromNumStrDto(nDto, precision). nDto='%v' Error= %v", nDto, err)
		}

		return d2, nil
	*/
}

// Pow - raises the current Decimal to the power of
// an integer 'exponent'. The result is returned as
// a Decimal type.
//
// Input Parameters:
// maxPrecision int - Determines the number of digits to the right of the
// 			decimal point returned in the result.
// 			** If the value of maxPrecision is minus 1:
//  					-- 	For Positive exponents, the actual number of digits calculated to
//        					the right of the decimal point will be returned.
//
//      			--	For Negative exponents, a default value of 500 digits to the right
//                  of the decimal point will be returned.
//
func (dec *Decimal) Pow(exponent int, maxPrecision int) (Decimal, error) {

	expSign := 1

	if exponent < 0 {
		expSign = -1
		exponent = exponent * -1
	}

	s1Val := dec.signedAllDigitsBigInt
	ex := big.NewInt(int64(exponent))

	newPrecision := dec.precision * uint(exponent)

	result := big.NewInt(0).Exp(s1Val, ex, nil)

	d2, err := dec.NumStrPrecisionToDecimal(result.String(), newPrecision, true)

	if err != nil {
		return Decimal{}.New(), fmt.Errorf("Pow() Error from nDto.NumStrPrecisionToDecimal(result.String(), newPrecision). Error= %v.", err)
	}

	if expSign == 1 {
		if maxPrecision > -1 {
			d2.SetPrecisionRound(uint(maxPrecision))
		}

		return d2, nil
	}

	// This is a negative exponent. Take the inverse of the result.
	d3, err := d2.Inverse(maxPrecision)

	if err != nil {
		return Decimal{}.New(), fmt.Errorf("Pow() Error from d2.Inverse(). Error= %v.", err)
	}

	return d3, nil

}

// Sets the value of the current Decimal to the input parameter 'iBig'
// scaled to the value of precision. In other words, if 'iBig' is set
// to a value of '123456' and precision is set to '3', the current
// Decimal will be set to a value of '123.456'.
//
// Using the same example, a big int value of '123456' with a precision of
// zero ('0') will yield an integer of '123456'.
//
// Example usage:
// d:= Decimal{}.New()
// iBig := big.NewInt(int64(123))
// d.SetBigInt(iBig, 1)
//
func (dec *Decimal) SetBigInt(iBig *big.Int, precision uint) error {

	d2, err := dec.MakeDecimalBigIntPrecision(iBig, precision)

	if err != nil {
		return fmt.Errorf("Error from MakeDecimalBigIntPrecision(). Error = %v", err)
	}

	dec.CopyIn(d2)

	return nil

}

// SetCurrencySymbol - sets the character which serves as the
// currency symbol for this Decimal value. Currency defaults
// to '$'
func (dec *Decimal) SetCurrencySymbol(currencySymbol rune) error {
	dec.currencySymbol = currencySymbol
	return nil
}

// SetDecimalSeparator - sets the character which separates
// the number into integer and fractional components. This
// defaults to "."
func (dec *Decimal) SetDecimalSeparator(decimalSeparator rune) error {
	dec.decimalSeparator = decimalSeparator
	return nil
}

// Sets the value of the current Decimal to the input parameter 'iNum'
// scaled to the value of precision. In other words, if 'iNum' is set
// to a value of '123456' and precision is set to '3', the current
// Decimal will be set to a value of '123.456'.
//
// Example usage:
// d:= Decimal{}.New()
// d.SetInt(123456, 3)
func (dec *Decimal) SetInt(iNum int, precision uint) error {

	iBig := big.NewInt(int64(iNum))

	return dec.SetBigInt(iBig, precision)
}

// SetIntFracStrings - Sets the value of a decimal based on separate
// integer and fraction strings passed as input parameters. 'intNum'
// parameter will constitute the integer component of the Decimal value
// while 'fracNum' will be converted to the fractional component of the
// new Decimal value. 'fracNum' represents all the numeric digits to
// the right of the decimal place, while 'intNum' represents all the
// integer digits to the left of the decimal place.
//
// The parameter 'signVal will determine the Sign Value for the returned
// Decimal type. It should be set to either +1 or -1.
func (dec *Decimal) SetIntFracStrings(signVal int, intNum, fracNum string) error {

	if dec.decimalSeparator == 0 {
		dec.decimalSeparator = '.'
	}

	if dec.thousandsSeparator == 0 {
		dec.thousandsSeparator = ','
	}

	if dec.currencySymbol == 0 {
		dec.currencySymbol = '$'
	}

	if signVal < 0 {
		signVal = -1
	} else {
		signVal = 1
	}

	n1 := NumStrDto{}.New()
	n1.DecimalSeparator = dec.decimalSeparator
	n1.ThousandsSeparator = dec.thousandsSeparator
	n1.CurrencySymbol = dec.currencySymbol

	n2, err := n1.ParseNumStr(intNum)

	if err != nil {
		return fmt.Errorf("SetIntFracStrings() - Error returned from n2 ParseNumStr(intNum). intNum= '%v' Error= %v", intNum, err)
	}

	n3, err := n1.ParseNumStr(fracNum)

	if err != nil {
		return fmt.Errorf("SetIntFracStrings() - Error returned from n3 ParseNumStr(fracNum). intNum= '%v' Error= %v", fracNum, err)
	}

	numStr := ""

	if signVal < 0 {
		numStr = "-"
	}

	numStr += string(n2.AbsAllNumRunes)

	if len(n3.AbsAllNumRunes) > 0 {
		numStr += string(dec.decimalSeparator)
		numStr += string(n3.AbsAllNumRunes)
	}

	n4, err := n1.ParseNumStr(numStr)

	if err != nil {
		return fmt.Errorf("SetIntFracStrings() - Error returned from n4 ParseNumStr(numStr). intNum= '%v' Error= %v", numStr, err)
	}

	d2, err := dec.MakeDecimalFromNumStrDto(n4)

	if err != nil {
		return fmt.Errorf("SetIntFracStrings() - Error returned from dec.MakeDecimalFromNumStrDto(n4). n4.NumStrOut = '%v' Error= %v", n4.NumStrOut, err)
	}

	dec.CopyIn(d2)

	return nil
}

// Sets the value of the current Decimal to the input parameter 'i64'
// scaled to the value of precision. In other words, if 'i64' is set
// to a value of '123456' and precision is set to '3', the current
// Decimal will be set to a value of '123.456'.
//
// Using the same example, an int64 value of '123456' and a precision
// value of zero ('0') will yield an integer value of '123456'.
//
// Example usage:
// d:= Decimal{}.New()
// i64:= int64(123456)
// d.SetInt64(i64, 3)
func (dec *Decimal) SetInt64(i64 int64, precision uint) error {

	return dec.SetBigInt(big.NewInt(i64), precision)

}

// SetPrecisionTrunc - Sets the precision or
// scale of the Decimal value. Precision determines
// the number of digits displayed to the right of
// the decimal place. Note that precision is
// processed as an unsigned integer.
//
// When reducing precision, existing digits
// are ROUNDED! When increasing precision,
// additional zeros ('0') are added to the right
// of the decimal place.
func (dec *Decimal) SetPrecisionRound(precision uint) error {

	n1, err := NumStrDto{}.NewPtr().SetPrecision(dec.numStr, precision, true)

	if err != nil {
		return fmt.Errorf("SetPrecisionRound() - Received error from NumStrDto.SetPrecision(dec.numStr, precision, true). dec.numStr='%v' precision='%v' Error= %v", dec.numStr, precision, err)
	}

	d2, err := dec.MakeDecimalFromNumStrDto(n1)

	if err != nil {
		return fmt.Errorf("SetPrecisionRound() - Received error from dec.MakeDecimalFromNumStrDto(n1, precision). dec.numStr='%v' precision='%v' Error= %v", dec.numStr, precision, err)
	}

	dec.CopyIn(d2)

	return nil

}

// SetPrecisionTrunc - Sets the precision or
// scale of the Decimal value. Precision determines
// the number of digits displayed to the right of
// the decimal place. Note that precision is
// processed as an unsigned integer.
//
// When reducing precision, existing digits
// are TRUNCATED! When increasing precision,
// additional zeros ('0') are added to the right
// of the decimal place.
func (dec *Decimal) SetPrecisionTrunc(precision uint) error {

	if dec.precision == precision {
		return nil
	}

	n1, err := NumStrDto{}.NewPtr().SetPrecision(dec.numStr, precision, false)

	if err != nil {
		return fmt.Errorf("SetPrecisionTrunc() - Received error from NumStrDto.SetPrecision(dec.numStr, precision, true). dec.numStr='%v' precision='%v' Error= %v", dec.numStr, precision, err)
	}

	d2, err := dec.MakeDecimalFromNumStrDto(n1)

	if err != nil {
		return fmt.Errorf("SetPrecisionTrunc() - Received error from dec.MakeDecimalFromNumStrDto(n1, precision). dec.numStr='%v' precision='%v' Error= %v", dec.numStr, precision, err)
	}

	dec.CopyIn(d2)

	return nil
}

// SetThousandsSeparator - sets the character which serves
// as the 'thousands' separator.
//
// This values defaults to ','
func (dec *Decimal) SetThousandsSeparator(thousandsSeparator rune) error {

	if thousandsSeparator == 0 {
		thousandsSeparator = ','
	}

	dec.thousandsSeparator = thousandsSeparator
	return nil
}

// SetFloat - Sets the value of the current decimal to
// that of the passed-in float32 parameter.
//
// Example usage:
// d:= Decimal{}.New()
// f32:= float32(123.456)
// d.SetFloat(f32)
func (dec *Decimal) SetFloat(f32 float32) error {

	if dec.decimalSeparator == 0 {
		dec.decimalSeparator = '.'
	}

	if dec.thousandsSeparator == 0 {
		dec.thousandsSeparator = ','
	}

	if dec.currencySymbol == 0 {
		dec.currencySymbol = '$'
	}

	str := strconv.FormatFloat(float64(f32), 'f', -1, 32)

	d2, err := dec.NumStrToDecimal(str)

	if err != nil {
		return fmt.Errorf("SetFloat() Error from nDto.NumStrToDecimal(str). str= '%v'. Error= %v", str, err)
	}

	dec.CopyIn(d2)

	return nil
}

// SetFloat64 - Sets the value of the current decimal to
// that of the passed-in float64 parameter.
//
// Example usage:
// d:= Decimal{}.New()
// f64:= float64(123.456)
// d.SetFloat(f64)
func (dec *Decimal) SetFloat64(f64 float64) error {

	if dec.decimalSeparator == 0 {
		dec.decimalSeparator = '.'
	}

	if dec.thousandsSeparator == 0 {
		dec.thousandsSeparator = ','
	}

	if dec.currencySymbol == 0 {
		dec.currencySymbol = '$'
	}

	str := strconv.FormatFloat(f64, 'f', -1, 64)

	d2, err := dec.NumStrToDecimal(str)

	if err != nil {
		return fmt.Errorf("SetFloat64() Error from nDto.NumStrToDecimal(str). str= '%v'. Error= %v", str, err)
	}

	dec.CopyIn(d2)

	return nil
}

// SetFloatBig - Sets the value of the current Decimal to the
// passed-in *big.Float parameter.
//
// Example usage:
// d:= Decimal{}.New()
// bigFloat:= big.NewFloat(float64(123.456))
// d.SetBigFloat(bigFloat)
func (dec *Decimal) SetFloatBig(bigFloat *big.Float) error {

	if dec.decimalSeparator == 0 {
		dec.decimalSeparator = '.'
	}

	if dec.thousandsSeparator == 0 {
		dec.thousandsSeparator = ','
	}

	if dec.currencySymbol == 0 {
		dec.currencySymbol = '$'
	}

	str := bigFloat.Text('f', -1)

	d2, err := dec.NumStrToDecimal(str)

	if err != nil {
		return fmt.Errorf("SetFloatBig() Error from nDto.NumStrToDecimal(str). str= '%v'. Error= %v", str, err)
	}

	dec.CopyIn(d2)

	return nil
}

// SetNumStrPrecision - Sets the Decimal's value to a number string and
// applies the appropriate 'precision' in order to determine placement
// of the decimal point. For example, if the number string ('str') is passed
// in as "123456" with a 'precision value of '3', the result is a value of
// 123.456.
//
// Example Usage:
// d:= Decimal{}.New()
// d.SetNumStrPrecision("123456", 3)
// Resulting Decimal Value = 123.456
func (dec *Decimal) SetNumStrPrecision(str string, precision uint, roundResult bool) error {

	if dec.decimalSeparator == 0 {
		dec.decimalSeparator = '.'
	}

	if dec.thousandsSeparator == 0 {
		dec.thousandsSeparator = ','
	}

	if dec.currencySymbol == 0 {
		dec.currencySymbol = '$'
	}

	d2, err := dec.NumStrPrecisionToDecimal(str, precision, roundResult)

	if err != nil {
		return fmt.Errorf("SetNumStrPrecision() NumStrPrecisionToDecimal(str) failed. str=%v. Error= %v", str, err)
	}

	dec.CopyIn(d2)

	return nil

}

// SetNumStr - Set's the Decimal's value to the input
// parameter 'str'. For example if 'str' is set equal
// to '123.456', this method will set the Decimal's
// value to 123.456.
//
// Example Usage:
// d := Decimal{}.New()
// d.SetNumStr("123.456")
// Decimal Value = 123.456
func (dec *Decimal) SetNumStr(str string) error {

	if dec.decimalSeparator == 0 {
		dec.decimalSeparator = '.'
	}

	if dec.thousandsSeparator == 0 {
		dec.thousandsSeparator = ','
	}

	if dec.currencySymbol == 0 {
		dec.currencySymbol = '$'
	}

	d2, err := dec.NumStrToDecimal(str)

	if err != nil {
		return fmt.Errorf("SetNumStr() Error. NumStrToDecimal(str) failed. str=%v. Error= %v.", str, err)
	}

	dec.CopyIn(d2)

	return nil

}

// SetNumStrDto - Sets the value of the current Decimal type
// to the value represented by the incoming NumStrDto parameter.
func (dec *Decimal) SetNumStrDto(nDto NumStrDto) error {

	if dec.decimalSeparator == 0 {
		dec.decimalSeparator = '.'
	}

	if dec.thousandsSeparator == 0 {
		dec.thousandsSeparator = ','
	}

	if dec.currencySymbol == 0 {
		dec.currencySymbol = '$'
	}

	d2, err := dec.MakeDecimalFromNumStrDto(nDto)

	if err != nil {
		return fmt.Errorf("SetNumStrDto() Error received from dec.MakeDecimalFromNumStrDto(nDto). nDto.NumStrOut= '%v' Error= %v", nDto.NumStrOut, err)
	}

	dec.CopyIn(d2)

	return nil
}

// Subtract - Subtracts the incoming Decimal from the current
// Decimal and returns the result as Decimal Type.
func (dec *Decimal) Subtract(d2 Decimal) (Decimal, error) {

	if !dec.isValid {
		return Decimal{}, errors.New("This Decimal Value Invalid!")
	}

	if !d2.isValid {
		return Decimal{}, errors.New("Incoming Decimal Invalid!")
	}

	base10 := big.NewInt(10)

	if dec.precision == d2.precision {

		s1Val := big.NewInt(0).Set(dec.signedAllDigitsBigInt)

		s2Val := big.NewInt(0).Set(d2.signedAllDigitsBigInt)

		s3Val := big.NewInt(0).Sub(s1Val, s2Val)

		s3Text := s3Val.String()

		d4, err := dec.NumStrPrecisionToDecimal(s3Text, dec.precision, true)

		if err != nil {
			return Decimal{}, fmt.Errorf("Add() received error from nDto.NumStrPrecisionToDecimal(s3Text, nDto.precision) s3Text='%v' nDto.precision=%v. Error= %v", s3Text, dec.precision, err)
		}

		return d4, nil
	}

	if d2.precision > dec.precision {
		deltaPrecision := big.NewInt(int64(d2.precision - dec.precision))
		deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)

		s1Val := big.NewInt(0).Mul(dec.signedAllDigitsBigInt, deltaPrecisionScale)
		s2Val := big.NewInt(0).Set(d2.signedAllDigitsBigInt)
		s3Val := big.NewInt(0).Sub(s1Val, s2Val)
		s3Text := s3Val.Text(10)

		d4, err := dec.NumStrPrecisionToDecimal(s3Text, d2.precision, true)

		if err != nil {
			return Decimal{}, fmt.Errorf("Add() received error from nDto.NumStrPrecisionToDecimal(s3Text, nDto.precision) s3Text='%v' nDto.precision=%v. Error= %v", s3Text, dec.precision, err)
		}

		return d4, nil

	}

	// must be nDto.precison > d2.precision

	deltaPrecision := big.NewInt(int64(dec.precision - d2.precision))
	deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
	s1Val := big.NewInt(0).Set(dec.signedAllDigitsBigInt)
	s2Val := big.NewInt(0).Mul(d2.signedAllDigitsBigInt, deltaPrecisionScale)

	s3Val := big.NewInt(0).Sub(s1Val, s2Val)

	s3Text := s3Val.Text(10)

	d4, err := dec.NumStrPrecisionToDecimal(s3Text, d2.precision, true)

	if err != nil {
		return Decimal{}, fmt.Errorf("Add() received error from nDto.NumStrPrecisionToDecimal(s3Text, nDto.precision) s3Text='%v' nDto.precision=%v. Error= %v", s3Text, dec.precision, err)
	}

	return d4, nil
}

// SubtractTotal - Subtracts the value of the incoming Decimal type
// from the current Decimal. The updated value is retained in the
// current Decimal type.
func (dec *Decimal) SubtractTotal(d2 Decimal) error {
	d3, err := dec.Subtract(d2)

	if err != nil {
		return err
	}

	dec.CopyIn(d3)

	return nil
}
