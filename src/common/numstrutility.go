package common

import (
	"fmt"
	"strconv"
)

/*
var numStrCurrencies = []rune{
	'\U00000024', // Dollar
	'\U000020ac', // Euro
	'\U000000a3', // Pound
	'\U000000a5', // China Yuan  & Japan Yen
	'\U000020a9', // Korea Won
	'\U000020bd', // Russian Ruble
	'\U000020b9', // Indian Rupee
	'\U0000fdfc', // Saudi Arabia Riyal
	'\U000020a8'} // Pakistan Rupee
*/

type NumStrUtility struct {
	Nation             string
	CurrencySymbol     string
	DecimalSeparator   rune
	ThousandsSeparator rune
	StrIn              string
	StrOut             string
}


func (ns NumStrUtility) DLimInt(num int, delimiter byte ) string {
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
