package common

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

/*

	The source code repository for numstrutility.go is located at :
					https://github.com/MikeAustin71/numstrutility.git



	Dependencies
  ************

		decimal.go
		numstrdto.go
		intary.go
		nthroot.go

 */

type NumStrUtility struct {
	Nation             string
	CurrencySymbol     rune
	DecimalSeparator   rune
	ThousandsSeparator rune
	StrIn              string
	StrOut             string
	IsFractionalVal    bool
	IntegerStr         string
	FractionStr        string
	Int64Val           int64
	Float64Val         float64
}

func (ns NumStrUtility) DLimInt(num int, delimiter byte) string {
	return ns.DnumStr(strconv.Itoa(num), delimiter)
}

// DLimI64 - Return a delimited number string with
// thousands separator (i.e. 1000000 -> 1,000,000)
func (ns NumStrUtility) DLimI64(num int64, delimiter byte) string {

	return ns.DnumStr(fmt.Sprintf("%v", num), delimiter)
}

func (ns NumStrUtility) DlimDecCurrStr(rawStr string, thousandsSeparator rune, decimal rune, currency rune) string {

	const maxStr = 256
	outStr := [maxStr]rune{}
	inStr := []rune(rawStr)
	lInStr := len(inStr)
	iCnt := 0
	outIdx := maxStr - 1
	outIdx1 := maxStr - 1
	outIdx2 := maxStr - 1
	r1 := [maxStr]rune{}
	r2 := [maxStr]rune{}
	decimalIsFound := false

	for i := lInStr - 1; i >= 0; i-- {
		if inStr[i] == decimal {
			r1[outIdx1] = decimal
			outIdx1--
			decimalIsFound = true
			continue
		}

		if !decimalIsFound {
			r1[outIdx1] = inStr[i]
			outIdx1--
		} else {
			r2[outIdx2] = inStr[i]
			outIdx2--
		}
	}

	var ptr *[maxStr]rune

	if !decimalIsFound {
		ptr = &r1
	} else {
		ptr = &r2
	}

	lIntrPart := len(ptr)

	for i := lIntrPart - 1; i >= 0; i-- {

		if ptr[i] >= '0' && ptr[i] <= '9' {

			iCnt++
			outStr[outIdx] = ptr[i]
			outIdx--

			if iCnt == 3 {
				iCnt = 0
				outStr[outIdx] = thousandsSeparator
				outIdx--
			}

			continue
		}

		// Check and allow for decimal
		// separators and sign designators
		if ptr[i] == '-' ||
			ptr[i] == '+' ||
			(ptr[i] == currency && currency != 0) {

			outStr[outIdx] = ptr[i]
			outIdx--

		}

	}

	if !decimalIsFound {
		return string(outStr[outIdx+1:])
	}

	return string(outStr[outIdx+1:]) + string(r1[outIdx1+1:])

}

// DnumStr - is designed to delimit or format a pure number string with a thousands
// separator (i.e. ','). Example: Input == 1234567890 -> Output == "1,234,567,890".
// NOTE: This method will not handle number strings containing decimal fractions
// and currency characters. For these options see method ns.DlimDecCurrStr(),
// above.
func (ns NumStrUtility) DnumStr(pureNumStr string, thousandsSeparator byte) string {
	const maxStr = 256
	outStr := [maxStr]byte{}
	lInStr := len(pureNumStr)
	iCnt := 0
	outIdx := maxStr - 1

	for i := lInStr - 1; i >= 0; i-- {

		if pureNumStr[i] >= '0' && pureNumStr[i] <= '9' {

			iCnt++
			outStr[outIdx] = pureNumStr[i]
			outIdx--

			if iCnt == 3 && i != 0 {
				iCnt = 0
				outStr[outIdx] = thousandsSeparator
				outIdx--
			}

			continue
		}

		// Check and allow for decimal
		// separators and sign designators
		if pureNumStr[i] == '-' ||
			pureNumStr[i] == '+' {

			outStr[outIdx] = pureNumStr[i]
			outIdx--

		}

	}

	return string(outStr[outIdx+1:])

}

// ConvertNumStrToInt64 - Converts string of numeric digits to an int64 value.
// The str parameter may including a leading sign value ('-' or '+'). If
// any following digits are non-numeric, an error will be
// generated
func (ns *NumStrUtility) ConvertNumStrToInt64(str string) (int64, error) {

	numRunes := ns.ConvertStrToIntNumRunes(strings.TrimLeft(strings.TrimRight(str, " "), " "))
	signVal := 1

	if str[0] == '-' {
		signVal = -1
	}

	return ns.ConvertRunesToInt64(numRunes, signVal)
}

// ConvertStrToIntNumStr - Converts a string of characters to
// a string consisting of a sign character ('-') followed by
// a string of numeric characters.
//
// NOTE: Fractional portions of a number string are ignored and
// not included in the integer number string returned.
func (ns *NumStrUtility) ConvertStrToIntNumStr(str string) string {

	nDto, err := ns.ParseNumString(str)

	if err != nil || !nDto.IsValid {
		return ""
	}

	return string(nDto.NumStrOut)
}

// ConvertInt64ToStr - Converts an int64 to a string of numeric
// characters. If the original number is less than zero, the first
// character of the numeric string is a minus sign ('-').
func (ns NumStrUtility) ConvertInt64ToStr(num int64) (string, error) {

	numStr := ""
	rAry := [248]rune{}

	signVal := int64(1)

	if num < 0 {
		num = num * -1
		signVal = -1
	}

	j := 247

	for num > 9 {
		r := num % 10
		num = num / 10
		rAry[j] = rune(r + 48)
		j--
	}

	rAry[j] = rune(num + 48)

	if signVal == -1 {
		j--
		rAry[j] = '-'
	}

	s := rAry[j:]

	numStr = string(s)

	return numStr, nil
}

func (ns *NumStrUtility) ConvertRunesToInt64(rAry []rune, signVal int) (int64, error) {

	lNumRunes := len(rAry)

	if lNumRunes == 0 {
		return int64(0), errors.New("Incoming rune array is Empty!")
	}

	if rAry[0] == '+' || rAry[0] == '-' {
		rAry = rAry[1:]
		lNumRunes = len(rAry)
	}

	numVal := int64(0)
	for i := 0; i < lNumRunes; i++ {

		if rAry[i] < '0' || rAry[i] > '9' {
			return int64(0), fmt.Errorf("Number string contained non-numeric characters: %v", string(rAry))
		}

		numVal *= 10
		numVal += int64(rAry[i] - 48)

	}

	numVal = numVal * int64(signVal)

	return numVal, nil

}

func (ns *NumStrUtility) ParseNumString(str string) (NumStrDto, error) {
	return NumStrDto{}.NewPtr().ParseNumStr(str)
/*
	nDto.NumStrIn = str

	if len(nDto.NumStrIn) == 0 {
		return nDto, errors.New("Zero length number string!")
	}

	nDto.SignVal = 1
	baseRunes := []rune(nDto.NumStrIn)
	lBaseRunes := len(baseRunes)
	isStartRunes := false
	isEndRunes := false

	for i := 0; i < lBaseRunes && isEndRunes == false; i++ {

		if baseRunes[i] == '-' &&
			isStartRunes == false && isEndRunes == false &&
			i+1 < lBaseRunes &&
			((baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9') ||
				baseRunes[i+1] == '.') {

			nDto.SignVal = -1
			isStartRunes = true
			continue

		} else if isEndRunes == false &&
			baseRunes[i] >= '0' && baseRunes[i] <= '9' {

			nDto.AbsAllNumRunes = append(nDto.AbsAllNumRunes, baseRunes[i])
			isStartRunes = true
			nDto.HasNumericDigits = true

			if nDto.IsFractionalValue {
				nDto.AbsFracRunes = append(nDto.AbsFracRunes, baseRunes[i])
			} else {
				nDto.AbsIntRunes = append(nDto.AbsIntRunes, baseRunes[i])
			}

		} else if (ns.ThousandsSeparator != ' ' && baseRunes[i] == ns.ThousandsSeparator) ||
			(ns.CurrencySymbol != ' ' && baseRunes[i] == ns.CurrencySymbol) ||
			baseRunes[i] == ' ' {

			continue

		} else if isEndRunes == false &&
			i+1 < lBaseRunes &&
			baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9' &&
			baseRunes[i] == '.' {

			nDto.IsFractionalValue = true
			continue

		} else if isStartRunes && !isEndRunes {

			isEndRunes = true

		}

	}

	if len(nDto.AbsAllNumRunes) == 0 {
		nDto.HasNumericDigits = false
		nDto.NumStrOut = "0"
		nDto.AbsIntRunes = append(nDto.AbsIntRunes, '0')
		nDto.SignVal = 1
		nDto.IsFractionalValue = false
		return nDto, nil
	}

	if len(nDto.AbsIntRunes) == 0 {
		nDto.AbsIntRunes = append(nDto.AbsIntRunes, '0')
	}

	if nDto.SignVal < 0 {
		nDto.NumStrOut = "-"
	}

	nDto.NumStrOut += string(nDto.AbsIntRunes)

	if nDto.IsFractionalValue {
		nDto.Precision = uint(len(nDto.AbsFracRunes))
		nDto.NumStrOut += "." + string(nDto.AbsFracRunes)
	}

	nDto.IsValid = true

	return nDto, nil
	*/
}

func (ns *NumStrUtility) ConvertNumStrToDecimal(str string) (Decimal, error) {
	dec := Decimal{}.New()

	err := dec.SetNumStr(str)

	return dec, err

}

// ConvertStrToIntNumRunes - Receives an integer string and returns
// a slice of runes. Note, thousands separator (',') and currency signs ('$')
// will be ignored and excluded from the array of runes returned by this
// method. In order to take advantage of this exclusion, the thousands
// separator and the currency symbol must be previously initialized in
// the NumStrUtility data fields.
func (ns *NumStrUtility) ConvertStrToIntNumRunes(str string) []rune {

	baseRunes := []rune(str)
	var numRunes []rune
	lBaseRunes := len(baseRunes)
	isStartRunes := false
	isEndRunes := false

	for i := 0; i < lBaseRunes; i++ {

		if (baseRunes[i] == '-' || baseRunes[i] == '+') &&
			len(numRunes) == 0 && isEndRunes == false &&
			i < lBaseRunes-2 &&
			baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9' {

			numRunes = append(numRunes, baseRunes[i])
			isStartRunes = true

		} else if isEndRunes == false &&
			baseRunes[i] >= '0' && baseRunes[i] <= '9' {

			numRunes = append(numRunes, baseRunes[i])
			isStartRunes = true

		} else if baseRunes[i] == ns.ThousandsSeparator || baseRunes[i] == ns.CurrencySymbol {

			continue

		} else if isStartRunes && !isEndRunes {

			isEndRunes = true

		}

	}

	return numRunes

}



func (ns *NumStrUtility) ConvertStrToFloat64(str string) (float64, error) {

	numF64 := float64(0.0)

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(str)

	if err != nil {
		return numF64, fmt.Errorf("ConvertStrToFloat64() - Error returned from NumStrDto.ParseNumStr(str). str= '%v' Error= %v",str, err)
	}

	numF64, err = strconv.ParseFloat(nDto.NumStrOut, 64)

	if err != nil {
		return float64(0.0), fmt.Errorf("ConvertStrToFloat64() - Error returned from strconv.ParseFloat(nDto.NumStrOut, 64). nDto.NumStrOut= '%v'. Error = %v", nDto.NumStrOut, err)
	}

	return numF64, nil
}


// ConvertInt64ToIntegerFloat64Value - Receives an int64 value and converts to a
// float64 value.  All of the digits are positioned to the right of the decimal
// place.
func (ns *NumStrUtility) ConvertInt64ToIntegerFloat64Value(i64 int64) (float64, error){

	f64 := float64(i64)

	return f64, nil

}

// ConvertInt64ToFractionalValue - Converts an int64 value to a float64 with
// all digits to the right of the decimal place.
func (ns *NumStrUtility) ConvertInt64ToFractionalValue(i64 int64) (float64, error) {

	ex := 1
	f64 := float64(i64)
	exp := math.Pow10(ex)
	for f64 > exp {
		ex++
		exp = math.Pow10(ex)
	}

	r64 := f64 / math.Pow10(ex)

	return r64, nil
}

// SetPrecision - Receives a string consisting of numeric
// digits. All numeric punctuation  (i.e. '.', ',', '$' is
// automatically removed. The string is converted to a
// floating point number string and the decimal placement
// is made according to the 'precision' parameter.  For example,
// a number string of '123456' with a precision parameter of '3'
// is converted to the number string '123.456'.
//
// Note: A sign character may be placed in the first character
// position before the first digit. Example: '-' minus or
// '+' plus.
func (ns *NumStrUtility) ScaleNumStr(str string, precision uint, roundResult bool) (NumStrDto, error) {



	return NumStrDto{}.NewPtr().SetPrecision(str,precision, roundResult)



}

func (ns *NumStrUtility) SetCountryAndCurrency(country string) error {

	lcStr := strings.ToLower(country)

	if strings.Contains(lcStr, "united states") {
		ns.Nation = "United States"
		ns.CurrencySymbol = NumStrCurrencySymbols[28]
		ns.ThousandsSeparator = ','
		ns.DecimalSeparator = '.'
		return nil
	}

	if strings.Contains(lcStr, "united kingdom") {
		ns.Nation = "United Kingdom"
		ns.CurrencySymbol = NumStrCurrencySymbols[29]
		ns.DecimalSeparator = '.'
		return nil
	}

	if strings.Contains(lcStr, "australia") {
		ns.Nation = "Australia"
		ns.CurrencySymbol = NumStrCurrencySymbols[0]
		return nil
	}

	if strings.Contains(lcStr, "brazil") {
		ns.Nation = "Brazil"
		ns.CurrencySymbol = NumStrCurrencySymbols[1]
		return nil
	}

	if strings.Contains(lcStr, "canada") {
		ns.Nation = "Canada"
		ns.CurrencySymbol = NumStrCurrencySymbols[2]
		ns.ThousandsSeparator = ','
		ns.DecimalSeparator = '.'
		return nil
	}

	if strings.Contains(lcStr, "china") {
		ns.Nation = "China"
		ns.CurrencySymbol = NumStrCurrencySymbols[3]
		return nil
	}

	if strings.Contains(lcStr, "colombia") {
		ns.Nation = "Colombia"
		ns.CurrencySymbol = NumStrCurrencySymbols[4]
		return nil
	}

	if strings.Contains(lcStr, "czech") {
		ns.Nation = "Czechoslovakia"
		ns.CurrencySymbol = NumStrCurrencySymbols[5]
		return nil
	}

	if strings.Contains(lcStr, "egypt") {
		ns.Nation = "Egypt"
		ns.CurrencySymbol = NumStrCurrencySymbols[6]
		return nil
	}

	if strings.Contains(lcStr, "euro") {
		ns.Nation = "Euro"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return nil
	}

	if strings.Contains(lcStr, "germany") {
		ns.Nation = "Germany"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return nil
	}

	if strings.Contains(lcStr, "france") {
		ns.Nation = "France"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return nil
	}

	if strings.Contains(lcStr, "italy") {
		ns.Nation = "Italy"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return nil
	}

	if strings.Contains(lcStr, "spain") {
		ns.Nation = "Spain"
		ns.CurrencySymbol = NumStrCurrencySymbols[7]
		return nil
	}

	if strings.Contains(lcStr, "hungary") {
		ns.Nation = "Hungary"
		ns.CurrencySymbol = NumStrCurrencySymbols[8]
		return nil
	}

	if strings.Contains(lcStr, "iceland") {
		ns.Nation = "Iceland"
		ns.CurrencySymbol = NumStrCurrencySymbols[9]
		return nil
	}

	if strings.Contains(lcStr, "indonesia") {
		ns.Nation = "Indonesia"
		ns.CurrencySymbol = NumStrCurrencySymbols[10]
		return nil
	}

	if strings.Contains(lcStr, "israel") {
		ns.Nation = "Israel"
		ns.CurrencySymbol = NumStrCurrencySymbols[11]
		return nil
	}

	if strings.Contains(lcStr, "japan") {
		ns.Nation = "Japan"
		ns.CurrencySymbol = NumStrCurrencySymbols[12]
		return nil
	}

	if strings.Contains(lcStr, "korea") {
		ns.Nation = "Korea"
		ns.CurrencySymbol = NumStrCurrencySymbols[13]
		return nil
	}

	if strings.Contains(lcStr, "malaysia") {
		ns.Nation = "Malaysia"
		ns.CurrencySymbol = NumStrCurrencySymbols[14]
		return nil
	}

	if strings.Contains(lcStr, "mexico") {
		ns.Nation = "Mexico"
		ns.CurrencySymbol = NumStrCurrencySymbols[15]
		return nil
	}

	if strings.Contains(lcStr, "norway") {
		ns.Nation = "Norway"
		ns.CurrencySymbol = NumStrCurrencySymbols[16]
		return nil
	}

	if strings.Contains(lcStr, "netherlands") {
		ns.Nation = "Netherlands"
		ns.CurrencySymbol = NumStrCurrencySymbols[17]
		return nil
	}

	if strings.Contains(lcStr, "pakistan") {
		ns.Nation = "Pakistan"
		ns.CurrencySymbol = NumStrCurrencySymbols[18]
		return nil
	}

	if strings.Contains(lcStr, "russia") {
		ns.Nation = "Russia"
		ns.CurrencySymbol = NumStrCurrencySymbols[19]
		return nil
	}

	if strings.Contains(lcStr, "saudi") {
		ns.Nation = "Saudi Arabia"
		ns.CurrencySymbol = NumStrCurrencySymbols[20]
		return nil
	}

	if strings.Contains(lcStr, "south africa") {
		ns.Nation = "South Africa"
		ns.CurrencySymbol = NumStrCurrencySymbols[21]
		return nil
	}

	if strings.Contains(lcStr, "sweden") {
		ns.Nation = "Sweden"
		ns.CurrencySymbol = NumStrCurrencySymbols[22]
		return nil
	}

	if strings.Contains(lcStr, "switzerland") {
		ns.Nation = "Switzerland"
		ns.CurrencySymbol = NumStrCurrencySymbols[23]
		return nil
	}

	if strings.Contains(lcStr, "taiwan") {
		ns.Nation = "Taiwan"
		ns.CurrencySymbol = NumStrCurrencySymbols[24]
		return nil
	}

	if strings.Contains(lcStr, "turkey") {
		ns.Nation = "Turkey"
		ns.CurrencySymbol = NumStrCurrencySymbols[25]
		return nil
	}

	if strings.Contains(lcStr, "venezuela") {
		ns.Nation = "Venezuela"
		ns.CurrencySymbol = NumStrCurrencySymbols[26]
		return nil
	}

	if strings.Contains(lcStr, "viet nam") {
		ns.Nation = "Viet Nam"
		ns.CurrencySymbol = NumStrCurrencySymbols[27]
		return nil
	}

	return fmt.Errorf("Failed to initialize country, %v.", country)
}
