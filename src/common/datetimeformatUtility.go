package common

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

type DateTimeWriteFormatsToFileDto struct {
	OutputPathFileName             string
	NumberOfFormatsGenerated       int
	NumberOfFormatMapKeysGenerated int
	FileWriteStartTime             time.Time
	FileWriteEndTime               time.Time
	ElapsedTimeForFileWriteOps     string
}

type DateTimeReadFormatsFromFileDto struct {
	PathFileName                   string
	NumberOfFormatsGenerated       int
	NumberOfFormatMapKeysGenerated int
	FileReadStartTime              time.Time
	FileReadEndTime                time.Time
	ElapsedTimeForFileReadOps      string
}

type DateTimeFormatRecord struct {
	FmtLength int
	FormatStr string
}

type DateTimeFormatGenerator struct {
	DayOfWeek            string
	DayOfWeekSeparator   string
	MthDay               string
	MthDayYear           string
	AfterMthDaySeparator string
	DateTimeSeparator    string
	TimeElement          string
	OffsetSeparator      string
	OffsetElement        string
	TimeZoneSeparator    string
	TimeZoneElement      string
	Year                 string
}

type SearchStrings struct {
	PreTrimSearchStrs [][][]string
	TimeFmtRegEx      [][][]string
}

type ParseDateTimeDto struct {
	IsSuccessful              bool
	FormattedDateTimeStringIn string
	SelectedMapIdx            int
	SelectedFormat            string
	TotalNoOfDictSearches     int
	DateTimeOut               time.Time
	err                       error
}

type DateTimeFormatUtility struct {
	OriginalDateTimeStringIn  string
	FormattedDateTimeStringIn string
	FormatMap                 map[int]map[string]int
	SelectedMapIdx            int
	SelectedFormat            string
	SelectedFormatSource      string
	DictSearches              [][][]int
	TotalNoOfDictSearches     int
	DateTimeOut               time.Time
	NumOfFormatsGenerated     int
	FormatSearchReplaceStrs   SearchStrings
}

// CreateAllFormatsInMemory - Currently this method generates
// approximately 1.5-million permutations of Date Time Formats
// and stores them in memory as a series of "maps of maps" using
// the field: DateTimeFormatUtility.FormatMap. This process currently
// takes about 1.5-seconds on my machine.
//
// 54-map keys are currently generated and used to access the format
// strings. Access is keyed on length of the date time string one is
// attempting to parse.
func (dtf *DateTimeFormatUtility) CreateAllFormatsInMemory() (err error) {

	dtf.FormatMap = make(map[int]map[string]int)
	dtf.NumOfFormatsGenerated = 0
	dtf.assemblePreDefinedFormats()
	dtf.assembleMthDayYearFmts()
	dtf.assembleEdgeCaseFormats()

	dtf.FormatSearchReplaceStrs.PreTrimSearchStrs = dtf.getPreTrimSearchStrings()
	dtf.FormatSearchReplaceStrs.TimeFmtRegEx = dtf.getTimeFmtRegEx()

	return
}

// LoadAllFormatsFromFileIntoMemory - Loads all date time formats from a specified
// text file into memory. The formats are stored in DateTimeFormatUtility.FormatMap.
// This is an alternative means of reading all available date time formats into
// memory so that they may be used to parse time strings. This method assumes that
// the text file containing the format strings was originally created by method
// DateTimeFormatUtility.WriteAllFormatsInMemoryToFile() which employs a specific
// fixed length format which can then be read back into memory.
func (dtf *DateTimeFormatUtility) LoadAllFormatsFromFileIntoMemory(pathFileName string) (DateTimeReadFormatsFromFileDto, error) {

	frDto := DateTimeReadFormatsFromFileDto{}
	frDto.PathFileName = pathFileName
	frDto.FileReadStartTime = time.Now()
	dtf.FormatMap = make(map[int]map[string]int)
	dtf.NumOfFormatsGenerated = 0

	fmtFile, err := os.Open(pathFileName)

	if err != nil {
		return frDto, fmt.Errorf("LoadAllFormatsFromFileIntoMemory- Error Opening File: %v - Error: %v", pathFileName, err.Error())
	}

	defer fmtFile.Close()
	const bufLen int = 2000
	lastBufIdx := 0
	var buffer []byte
	var outRecordBuff []byte
	IsEOF := false
	idx := 0
	isPartialRec := false
	buffer = make([]byte, bufLen)

	// Read File Operation
	// REDO:
	for IsEOF == false {

		bytesRead, err := fmtFile.Read(buffer)

		if err != nil {
			IsEOF = true
		}

		idx = 0

		lastBufIdx = bytesRead - 1

		// Begin Read Record Operation
		for bytesRead > 0 {

			if !isPartialRec {
				outRecordBuff = make([]byte, 0)
			} else {
				isPartialRec = false
			}

			for i := idx; i <= lastBufIdx; i++ {
				// Extract one record from buffer and process
				if buffer[i] == '\n' {
					idx = i + 1
					break
				}

				outRecordBuff = append(outRecordBuff, buffer[i])

				if i == lastBufIdx {
					isPartialRec = true
					idx = 0
				}
			}

			// Break up the record into
			// two fields, int Length and
			// string Format.
			lOutBuff := len(outRecordBuff)

			if isPartialRec || lOutBuff < 7 {
				isPartialRec = true
				break
			}

			lenField := make([]byte, 7)

			for j := 0; j < 7; j++ {

				lenField[j] = outRecordBuff[j]
			}

			s := string(lenField)

			lFmt, err := strconv.Atoi(s)

			if err != nil {
				return frDto, fmt.Errorf(
					"LoadAllFormatsFromFileIntoMemory - Error converting Format Length field from file. Length = %v. Idx= %v. Format Count: %v",
					s, idx, frDto.NumberOfFormatsGenerated)
			}

			fmtFieldLastIdx := 7 + lFmt

			if lOutBuff < fmtFieldLastIdx+1 {
				return frDto, fmt.Errorf(
					"LoadAllFormatsFromFileIntoMemory - Found corrupted Output Buffer. Buffer Length %v, Length Field = %v, Output Buffer= %s Format Count: %v",
					lOutBuff, lFmt, string(outRecordBuff), frDto.NumberOfFormatsGenerated)
			}

			fmtField := make([]byte, lFmt)

			for k := 8; k <= fmtFieldLastIdx; k++ {
				fmtField[k-8] = outRecordBuff[k]
			}

			fmtStr := string(fmtField)

			// Populate DateTimeFormatUtility.FormatMap
			if dtf.FormatMap[lFmt] == nil {
				dtf.FormatMap[lFmt] = make(map[string]int)
				frDto.NumberOfFormatMapKeysGenerated++
			}

			if dtf.FormatMap[lFmt][fmtStr] == 0 {
				dtf.FormatMap[lFmt][fmtStr] = lFmt
				dtf.NumOfFormatsGenerated++
				frDto.NumberOfFormatsGenerated++
			}

			if idx > lastBufIdx {
				break
			}
		}
	}

	frDto.FileReadEndTime = time.Now()
	frDto.NumberOfFormatMapKeysGenerated = len(dtf.FormatMap)
	du := DurationUtility{}
	err = du.SetStartEndTimes(frDto.FileReadStartTime, frDto.FileReadEndTime)

	if err != nil {
		return DateTimeReadFormatsFromFileDto{}, fmt.Errorf("LoadAllFormatsFromFileIntoMemory - Error SetStartEndTimes() - %v", err.Error())
	}

	yrMthDaysTime, err := du.GetYearMthDaysTime()

	if err != nil {
		return DateTimeReadFormatsFromFileDto{}, fmt.Errorf("LoadAllFormatsFromFileIntoMemory - Error GetYearMthDaysTime() - %v", err.Error())
	}

	frDto.ElapsedTimeForFileReadOps = yrMthDaysTime.DisplayStr

	return frDto, nil
}

// WriteAllFormatsInMemoryToFile - Writes all Format Data contained in
// DateTimeFormatUtility.FormatMap field to a specified output file in
// text format. Currently, about 1.4-million formats are generated and
// written to the output file.
//
// IMPORTANT! - Before you call this method, the Format Maps must
// first be created in memory. Call DateTimeFormatUtility.CreateAllFormatsInMemory()
// first, before calling this method.
func (dtf *DateTimeFormatUtility) WriteAllFormatsInMemoryToFile(outputPathFileName string) (DateTimeWriteFormatsToFileDto, error) {

	fwDto := DateTimeWriteFormatsToFileDto{}

	fwDto.FileWriteStartTime = time.Now()
	lFmts := len(dtf.FormatMap)

	if lFmts < 1 {
		return fwDto, errors.New("WriteAllFormatsInMemoryToFile() Error - There are NO Formats in Memory -  FormatMap length == 0. You MUST call CreateAllFormatsInMemory() first!")
	}

	outF, err := os.Create(outputPathFileName)

	if err != nil {
		return fwDto, fmt.Errorf("WriteAllFormatsInMemoryToFile() Error - Failed create output file %v. Error: %v", outputPathFileName, err.Error())
	}

	defer outF.Close()

	var keys []int
	for k := range dtf.FormatMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, k := range keys {

		fwDto.NumberOfFormatMapKeysGenerated++

		for keyFmt := range dtf.FormatMap[k] {
			fwDto.NumberOfFormatsGenerated++
			_, err := outF.WriteString(fmt.Sprintf("%07d %s\n", k, keyFmt))

			if err != nil {
				return DateTimeWriteFormatsToFileDto{}, fmt.Errorf("WriteAllFormatsInMemoryToFile() Error writing Format data to output file %v. Error: %v", outputPathFileName, err.Error())
			}
		}
	}

	outF.Sync()

	fwDto.FileWriteEndTime = time.Now()

	du := DurationUtility{}

	err = du.SetStartEndTimes(fwDto.FileWriteStartTime, fwDto.FileWriteEndTime)

	if err != nil {
		return DateTimeWriteFormatsToFileDto{}, fmt.Errorf("WriteAllFormatsInMemoryToFile() Error Setting Start End Times for Duration Calculation Error: %v", err.Error())
	}

	fwDto.OutputPathFileName = outputPathFileName

	etFileWrite, err := du.GetYearMthDaysTime()

	if err != nil {
		return DateTimeWriteFormatsToFileDto{}, fmt.Errorf("WriteAllFormatsInMemoryToFile() Error du.GetYearMthDaysTime() - %v", err.Error())
	}

	fwDto.ElapsedTimeForFileWriteOps = etFileWrite.DisplayStr

	return fwDto, nil
}

// WriteFormatStatsToFile - This method writes data to a text file.
// The text file is small, currently about 3-kilobytes in size.
// The data output to the text file describes the size of the
// slices contained in dtf.FormatMap.
// IMPORTANT! - Before you call this method, the Format Maps must
// first be created in memory. Call DateTimeFormatUtility.CreateAllFormatsInMemory()
// first, before calling this method.
func (dtf *DateTimeFormatUtility) WriteFormatStatsToFile(outputPathFileName string) (DateTimeWriteFormatsToFileDto, error) {
	outputDto := DateTimeWriteFormatsToFileDto{}
	outputDto.OutputPathFileName = outputPathFileName
	outputDto.FileWriteStartTime = time.Now()

	lFmts := len(dtf.FormatMap)

	if lFmts < 1 {
		return outputDto,
			errors.New("WriteFormatStatsToFile() Error - There are NO Formats in Memory -  FormatMap length == 0. You MUST call CreateAllFormatsInMemory() first!")
	}

	f, err := os.Create(outputPathFileName)

	if err != nil {
		return outputDto, errors.New("Output File Create Error:" + err.Error())
	}

	defer f.Close()

	var keys []int
	for k := range dtf.FormatMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	numOfKeys := 0
	numOfFormats := 0
	mapLen := 0

	_, err = f.WriteString("Length - Number Of Formats\n")

	if err != nil {
		return outputDto, errors.New("Error Writing Header To Output File! Error: " + err.Error())
	}

	for _, k := range keys {

		numOfKeys++
		mapLen = len(dtf.FormatMap[k])
		numOfFormats += mapLen

		_, err = f.WriteString(fmt.Sprintf("%06d%18d\n", k, mapLen))

		if err != nil {
			return outputDto, fmt.Errorf("Error Writing Stats to Output File! mapLen=%v Error: %v", mapLen, err.Error())
		}

	}

	outputDto.FileWriteEndTime = time.Now()
	du := DurationUtility{}
	err = du.SetStartEndTimes(outputDto.FileWriteStartTime, outputDto.FileWriteEndTime)

	if err != nil {
		return outputDto, fmt.Errorf("Error Calculating Duration with SetStartEndTimes() Error: %v", err.Error())
	}

	etFileWrite, err := du.GetYearMthDaysTime()

	if err != nil {
		return outputDto, fmt.Errorf("Error Returning Duration with GetYearMthDaysTime() Error: %v", err.Error())
	}

	outputDto.ElapsedTimeForFileWriteOps = etFileWrite.DisplayStr
	outputDto.NumberOfFormatsGenerated = numOfFormats
	outputDto.NumberOfFormatMapKeysGenerated = numOfKeys

	return outputDto, nil
}

func (dtf *DateTimeFormatUtility) Empty() {
	dtf.OriginalDateTimeStringIn = ""
	dtf.FormattedDateTimeStringIn = ""
	dtf.DateTimeOut = time.Time{}
	dtf.SelectedFormat = ""
	dtf.SelectedFormatSource = ""
	dtf.SelectedMapIdx = -1
	dtf.DictSearches = make([][][]int, 0)
	dtf.TotalNoOfDictSearches = 0

}

// ParseDateTimeString - Parses date time strings passed into the method. If a
// format is passed to the this method as the second parameter, the method will
// use this format first in an attempt to convert the date time string to a time.Time
// structure. If the 'probableFormat' parameter is empty or if it fails to convert
// the time string to a valid time.Time value, this method will run the date time
// string against 1.4-million possible date time string formats in an effort to
// successfully convert the date time string into a valid time.Time value.
func (dtf *DateTimeFormatUtility) ParseDateTimeString(dateTimeStr string, probableFormat string) (time.Time, error) {

	if dateTimeStr == "" {
		return time.Time{}, errors.New("Empty Time String!")
	}

	dtf.Empty()

	xtimeStr := dtf.replaceMultipleStrSequence(dateTimeStr, dtf.FormatSearchReplaceStrs.PreTrimSearchStrs)
	xtimeStr = dtf.replaceDateSuffixStThRd(xtimeStr)
	xtimeStr = dtf.reformatSingleTimeString(xtimeStr, dtf.FormatSearchReplaceStrs.TimeFmtRegEx)
	xtimeStr = dtf.replaceAMPM(xtimeStr)

	ftimeStr, err := dtf.trimEndMultiple(xtimeStr, ' ')

	if err != nil {
		return time.Time{}, err
	}

	if probableFormat != "" {
		t, err := time.Parse(probableFormat, ftimeStr)

		if err == nil {
			dtf.SelectedFormat = probableFormat
			dtf.SelectedFormatSource = "User Provided"
			dtf.SelectedMapIdx = -1
			dtf.DateTimeOut = t
			dtf.OriginalDateTimeStringIn = dateTimeStr
			dtf.FormattedDateTimeStringIn = ftimeStr
			dtf.TotalNoOfDictSearches = 1
			dtf.DictSearches = append(dtf.DictSearches, [][]int{{0, 1}})

			return t, nil
		}

	}

	if len(dtf.FormatMap) == 0 {

		dtf.CreateAllFormatsInMemory()

		if len(dtf.FormatMap) == 0 {
			return time.Time{}, errors.New("Format Map is EMPTY! Load Formats into DateTimeFormatUtility.FormatMap first!")
		}

	}

	lenStr := len(ftimeStr)

	lenSequence := make([][]int, 0)

	lenTests := []int{
		lenStr - 2,
		lenStr - 3,
		lenStr - 1,
		lenStr,
		lenStr + 1,
		lenStr + 2,
		lenStr + 3,
		lenStr - 4,
		lenStr + 4,
		lenStr - 5,
		lenStr + 5,
		lenStr - 6,
		lenStr + 6,
		lenStr - 7,
		lenStr + 7,
		lenStr - 8,
		lenStr + 8,
		lenStr - 9,
		lenStr + 9,
	}

	dtf.TotalNoOfDictSearches = 0
	dtf.OriginalDateTimeStringIn = dateTimeStr
	dtf.FormattedDateTimeStringIn = ftimeStr
	dtf.DateTimeOut = time.Time{}

	threshold := 5

	if lenStr >= 30 {
		threshold = 8
	}

	ary := make([]int, 0)
	for i := 0; i < len(lenTests); i++ {

		if dtf.FormatMap[lenTests[i]] != nil {
			ary = append(ary, lenTests[i])
		}

		if len(ary) == threshold {
			lenSequence = append(lenSequence, ary)
			ary = make([]int, 0)
		}

	}

	if len(lenSequence) > 0 {
		lenSequence = append(lenSequence, ary)
	}

	for j := 0; j < len(lenSequence); j++ {

		if dtf.doParseRun(lenSequence[j], ftimeStr) {
			return dtf.DateTimeOut, nil
		}

	}

	return time.Time{}, errors.New("Failed to locate correct time format!")
}

func (dtf *DateTimeFormatUtility) doParseRun(lenTests []int, ftimeStr string) bool {

	lenLenTests := len(lenTests)
	msg := make(chan ParseDateTimeDto)
	done := uint64(0)

	isSuccessfulParse := false

	for i := 0; i < lenLenTests; i++ {

		go dtf.parseFormatMap(msg, &done, ftimeStr, lenTests[i], dtf.FormatMap[lenTests[i]])

	}

	cnt := 0

	for m := range msg {
		cnt++

		if cnt == lenLenTests {
			close(msg)
			runtime.Gosched()
		}

		dtf.DictSearches =
			append(dtf.DictSearches, [][]int{{m.SelectedMapIdx, m.TotalNoOfDictSearches}})
		dtf.TotalNoOfDictSearches += m.TotalNoOfDictSearches

		if m.IsSuccessful && !isSuccessfulParse {
			isSuccessfulParse = true
			dtf.SelectedFormat = m.SelectedFormat
			dtf.SelectedFormatSource = "Format Map Dictionary"
			dtf.SelectedMapIdx = m.SelectedMapIdx
			dtf.DateTimeOut = m.DateTimeOut

		}

	}

	return isSuccessfulParse
}

func (dtf *DateTimeFormatUtility) parseFormatMap(
	msg chan<- ParseDateTimeDto, done *uint64, timeStr string, idx int, fmtMap map[string]int) {

	var doneTest uint64

	dto := ParseDateTimeDto{}

	dto.FormattedDateTimeStringIn = timeStr
	dto.SelectedMapIdx = idx

	for key := range fmtMap {

		dto.TotalNoOfDictSearches++

		t, err := time.Parse(key, timeStr)

		doneTest = atomic.LoadUint64(done)

		if doneTest > 0 {
			msg <- dto
			return
		}

		if err == nil {
			atomic.AddUint64(done, 1)
			dto.DateTimeOut = t
			dto.SelectedFormat = key
			dto.IsSuccessful = true
			msg <- dto
			runtime.Gosched()
			return
		}

	}

	msg <- dto
	return

}

func (dtf *DateTimeFormatUtility) assembleDayMthYears() error {

	dtf.FormatMap = make(map[int]map[string]int)

	fmtStr := ""

	dayOfWeek, _ := dtf.getDayOfWeekElements()
	dayOfWeekSeparators, _ := dtf.getDayOfWeekSeparator()
	mthDayYearFmts, _ := dtf.getMonthDayYearElements()

	for _, dowk := range dayOfWeek {
		for _, dowkSep := range dayOfWeekSeparators {

			for _, mmddyyy := range mthDayYearFmts {

				if dowk != "" && mmddyyy != "" {
					fmtStr = dowk + dowkSep + mmddyyy
					dtf.assignFormatStrToMap(fmtStr)
				} else if dowk != "" && mmddyyy == "" {
					dtf.assignFormatStrToMap(dowk)
				} else if dowk == "" && mmddyyy != "" {
					dtf.assignFormatStrToMap(mmddyyy)
				}
			}
		}
	}

	return nil
}

func (dtf *DateTimeFormatUtility) assembleMthDayYearFmts() error {

	dayOfWeek, _ := dtf.getDayOfWeekElements()

	dayOfWeekSeparators, _ := dtf.getDayOfWeekSeparator()

	mthDayYearFmts, _ := dtf.getMonthDayYearElements()

	dateTimeSeparators, _ := dtf.getDateTimeSeparators()

	timeFmts, _ := dtf.getTimeElements()

	offsetSeparators, _ := dtf.getTimeOffsetSeparators()

	offsetFmts, _ := dtf.getTimeOffsets()

	tzSeparators, _ := dtf.getTimeZoneSeparators()

	timeZoneFmts, _ := dtf.getTimeZoneElements()

	for _, dowk := range dayOfWeek {
		for _, dowkSep := range dayOfWeekSeparators {
			for _, mmddyyyy := range mthDayYearFmts {
				for _, dtSep := range dateTimeSeparators {
					for _, t := range timeFmts {
						for _, tOffsetSep := range offsetSeparators {
							for _, offFmt := range offsetFmts {
								for _, stdSep := range tzSeparators {
									for _, tzF := range timeZoneFmts {
										fmtGen := DateTimeFormatGenerator{
											DayOfWeek:          dowk,
											DayOfWeekSeparator: dowkSep,
											MthDayYear:         mmddyyyy,
											DateTimeSeparator:  dtSep,
											TimeElement:        t,
											OffsetSeparator:    tOffsetSep,
											OffsetElement:      offFmt,
											TimeZoneSeparator:  stdSep,
											TimeZoneElement:    tzF,
										}

										dtf.analyzeDofWeekMMDDYYYYTimeOffsetTz(fmtGen)
									}
								}

							}
						}
					}
				}

			}
		}

	}

	return nil
}

func (dtf *DateTimeFormatUtility) assembleMthDayTimeOffsetTzYearFmts() error {

	dayOfWeek, _ := dtf.getDayOfWeekElements()

	dayOfWeekSeparators, _ := dtf.getDayOfWeekSeparator()

	mthDayElements, _ := dtf.getMonthDayElements()

	afterMthDaySeparators, _ := dtf.getAfterMthDaySeparators()

	timeFmts, _ := dtf.getTimeElements()

	offsetSeparators, _ := dtf.getTimeOffsetSeparators()

	offsetFmts, _ := dtf.getTimeOffsets()

	tzSeparators, _ := dtf.getTimeZoneSeparators()

	timeZoneFmts, _ := dtf.getTimeZoneElements()

	yearElements, _ := dtf.getYears()

	for _, dowk := range dayOfWeek {
		for _, dowkSep := range dayOfWeekSeparators {
			for _, mthDay := range mthDayElements {
				for _, afterMthDaySeparator := range afterMthDaySeparators {
					for _, t := range timeFmts {
						for _, tOffsetSep := range offsetSeparators {
							for _, offFmt := range offsetFmts {
								for _, stdSep := range tzSeparators {
									for _, tzF := range timeZoneFmts {
										for _, yearEle := range yearElements {

											fmtGen := DateTimeFormatGenerator{
												DayOfWeek:            dowk,
												DayOfWeekSeparator:   dowkSep,
												MthDayYear:           mthDay,
												AfterMthDaySeparator: afterMthDaySeparator,
												TimeElement:          t,
												OffsetSeparator:      tOffsetSep,
												OffsetElement:        offFmt,
												TimeZoneSeparator:    stdSep,
												TimeZoneElement:      tzF,
												Year:                 yearEle,
											}

											dtf.analyzeDofWeekMMDDTimeOffsetTzYYYY(fmtGen)

										}
									}
								}

							}
						}
					}
				}

			}
		}

	}

	return nil
}

func (dtf *DateTimeFormatUtility) analyzeDofWeekMMDDYYYYTimeOffsetTz(dtfGen DateTimeFormatGenerator) {

	fmtStr := ""
	fmtStr2 := ""

	if dtfGen.MthDayYear == "" &&
		dtfGen.TimeElement == "" {
		return
	}

	if dtfGen.DayOfWeek != "" {
		fmtStr += dtfGen.DayOfWeek
	}

	if dtfGen.MthDayYear != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.MthDayYear
		} else {
			fmtStr += dtfGen.DayOfWeekSeparator
			fmtStr += dtfGen.MthDayYear
		}
	}

	if dtfGen.TimeElement != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.TimeElement
		} else {
			fmtStr += dtfGen.DateTimeSeparator
			fmtStr += dtfGen.TimeElement
		}
	}

	fmtStr2 = fmtStr

	if dtfGen.OffsetElement != "" &&
		fmtStr != "" &&
		dtfGen.TimeElement != "" {

		fmtStr += dtfGen.OffsetSeparator
		fmtStr += dtfGen.OffsetElement
	}

	if dtfGen.TimeZoneElement != "" &&
		fmtStr != "" &&
		dtfGen.TimeElement != "" {

		fmtStr += dtfGen.TimeZoneSeparator
		fmtStr += dtfGen.TimeZoneElement
	}

	if fmtStr != "" {
		dtf.assignFormatStrToMap(fmtStr)
	}

	// Calculate variation of format string where
	// Time Zone comes before Offset Element

	if dtfGen.TimeZoneElement != "" &&
		dtfGen.TimeElement != "" &&
		fmtStr2 != "" &&
		dtfGen.OffsetElement != "" {

		fmtStr2 += dtfGen.TimeZoneSeparator
		fmtStr2 += dtfGen.TimeZoneElement
		fmtStr2 += dtfGen.OffsetSeparator
		fmtStr2 += dtfGen.OffsetElement
	}

	if fmtStr2 != "" {
		dtf.assignFormatStrToMap(fmtStr2)
	}

	return

}

func (dtf *DateTimeFormatUtility) assemblePreDefinedFormats() {

	preDefFmts := dtf.getPredefinedFormats()

	for _, pdf := range preDefFmts {

		dtf.assignFormatStrToMap(pdf)

	}

}

func (dtf *DateTimeFormatUtility) assembleEdgeCaseFormats() {
	edgeCases := dtf.getEdgeCases()

	for _, ecf := range edgeCases {
		dtf.assignFormatStrToMap(ecf)
	}
}

func (dtf *DateTimeFormatUtility) analyzeDofWeekMMDDTimeOffsetTzYYYY(dtfGen DateTimeFormatGenerator) {

	fmtStr := ""
	fmtStr2 := ""

	if dtfGen.DayOfWeek != "" {
		fmtStr += dtfGen.DayOfWeek
	}

	if dtfGen.MthDay != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.MthDay
		} else {
			fmtStr += dtfGen.DayOfWeekSeparator
			fmtStr += dtfGen.MthDay
		}
	}

	if dtfGen.TimeElement != "" {
		if fmtStr == "" {
			fmtStr = dtfGen.TimeElement
		} else {
			fmtStr += dtfGen.AfterMthDaySeparator
			fmtStr += dtfGen.TimeElement
		}
	}

	fmtStr2 = fmtStr

	if dtfGen.OffsetElement != "" &&
		fmtStr != "" &&
		dtfGen.TimeElement != "" {
		fmtStr += dtfGen.OffsetSeparator
		fmtStr += dtfGen.OffsetElement

	}

	if dtfGen.TimeZoneElement != "" &&
		fmtStr != "" &&
		dtfGen.TimeElement != "" {
		fmtStr += dtfGen.TimeZoneSeparator
		fmtStr += dtfGen.TimeZoneElement
	}

	if fmtStr != "" {
		dtf.assignFormatStrToMap(fmtStr)
	}

	// Calculate variation of format string where
	// Time Zone comes before Offset Element

	if dtfGen.TimeZoneElement != "" &&
		fmtStr2 != "" &&
		dtfGen.TimeElement != "" {
		fmtStr2 += dtfGen.TimeZoneSeparator
		fmtStr2 += dtfGen.TimeZoneElement
	}

	if dtfGen.OffsetElement != "" &&
		fmtStr2 != "" &&
		dtfGen.TimeElement != "" {
		fmtStr2 += dtfGen.OffsetSeparator
		fmtStr2 += dtfGen.OffsetElement

	}

	if fmtStr2 != "" {
		dtf.assignFormatStrToMap(fmtStr)
	}

	return
}

func (dtf *DateTimeFormatUtility) assignFormatStrToMap(fmtStr string) {

	l := len(fmtStr)

	if l == 0 {
		return
	}

	if dtf.FormatMap[l] == nil {
		dtf.FormatMap[l] = make(map[string]int)
	}

	if dtf.FormatMap[l][fmtStr] != 0 {
		return
	}

	dtf.FormatMap[l][fmtStr] = l
	dtf.NumOfFormatsGenerated++
}

func (dtf *DateTimeFormatUtility) getDayOfWeekElements() ([]string, error) {
	dayOfWeek := make([]string, 0, 10)

	dayOfWeek = append(dayOfWeek, "")
	dayOfWeek = append(dayOfWeek, "Mon")
	dayOfWeek = append(dayOfWeek, "Monday")

	return dayOfWeek, nil
}

func (dtf *DateTimeFormatUtility) getDayOfWeekSeparator() ([]string, error) {
	dayOfWeekSeparator := make([]string, 0, 1024)

	dayOfWeekSeparator = append(dayOfWeekSeparator, " ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, ", ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, " - ")
	dayOfWeekSeparator = append(dayOfWeekSeparator, "-")

	return dayOfWeekSeparator, nil
}

func (dtf *DateTimeFormatUtility) getMonthDayYearElements() ([]string, error) {
	mthDayYr := make([]string, 0, 1024)

	mthDayYr = append(mthDayYr, "2006-1-2")
	mthDayYr = append(mthDayYr, "2006 1 2")
	mthDayYr = append(mthDayYr, "1-2-06")
	mthDayYr = append(mthDayYr, "1-2-2006")

	mthDayYr = append(mthDayYr, "1 2 06")
	mthDayYr = append(mthDayYr, "1 2 2006")

	mthDayYr = append(mthDayYr, "Jan-2-06")
	mthDayYr = append(mthDayYr, "Jan 2 06")
	mthDayYr = append(mthDayYr, "Jan _2 06")
	mthDayYr = append(mthDayYr, "Jan-2-2006")
	mthDayYr = append(mthDayYr, "Jan 2 2006")
	mthDayYr = append(mthDayYr, "Jan _2 2006")

	mthDayYr = append(mthDayYr, "January-2-06")
	mthDayYr = append(mthDayYr, "January 2 06")
	mthDayYr = append(mthDayYr, "January _2 06")
	mthDayYr = append(mthDayYr, "January-2-2006")
	mthDayYr = append(mthDayYr, "January 2 2006")
	mthDayYr = append(mthDayYr, "January _2 2006")

	// European Date Formats DD.MM.YYYY
	mthDayYr = append(mthDayYr, "2.1.06")
	mthDayYr = append(mthDayYr, "2.1.2006")
	mthDayYr = append(mthDayYr, "2.1.'06")

	// Standard Dates with Dot Delimiters
	mthDayYr = append(mthDayYr, "2006.1.2")

	mthDayYr = append(mthDayYr, "2-January-2006")
	mthDayYr = append(mthDayYr, "2-January-06")
	mthDayYr = append(mthDayYr, "2 January 06")
	mthDayYr = append(mthDayYr, "2 January 2006")

	mthDayYr = append(mthDayYr, "2-Jan-2006")
	mthDayYr = append(mthDayYr, "2-Jan-06")
	mthDayYr = append(mthDayYr, "2 Jan 06")
	mthDayYr = append(mthDayYr, "2 Jan 2006")

	mthDayYr = append(mthDayYr, "20060102")
	mthDayYr = append(mthDayYr, "01022006")
	mthDayYr = append(mthDayYr, "")

	return mthDayYr, nil
}

func (dtf *DateTimeFormatUtility) getMonthDayElements() ([]string, error) {
	mthDayElements := make([]string, 0, 124)

	mthDayElements = append(mthDayElements, "Jan 2")
	mthDayElements = append(mthDayElements, "January 2")
	mthDayElements = append(mthDayElements, "Jan _2")
	mthDayElements = append(mthDayElements, "January _2")
	mthDayElements = append(mthDayElements, "1-2")
	mthDayElements = append(mthDayElements, "1-_2")
	mthDayElements = append(mthDayElements, "1 2")
	mthDayElements = append(mthDayElements, "1-2")

	mthDayElements = append(mthDayElements, "01 02")
	mthDayElements = append(mthDayElements, "01 _2")
	mthDayElements = append(mthDayElements, "01-02")

	mthDayElements = append(mthDayElements, "0102")
	// European Format Day Month
	mthDayElements = append(mthDayElements, "02.01")
	mthDayElements = append(mthDayElements, "2.1")
	mthDayElements = append(mthDayElements, "02.1")
	mthDayElements = append(mthDayElements, "2.01")

	return mthDayElements, nil
}

func (dtf *DateTimeFormatUtility) getYears() ([]string, error) {
	yearElements := make([]string, 0, 10)

	yearElements = append(yearElements, "2006")
	yearElements = append(yearElements, "06")
	yearElements = append(yearElements, "'06")

	return yearElements, nil
}

func (dtf *DateTimeFormatUtility) getAfterMthDaySeparators() ([]string, error) {
	mthDayAfterSeparators := make([]string, 0, 10)

	mthDayAfterSeparators = append(mthDayAfterSeparators, " ")
	mthDayAfterSeparators = append(mthDayAfterSeparators, ", ")
	mthDayAfterSeparators = append(mthDayAfterSeparators, ":")
	mthDayAfterSeparators = append(mthDayAfterSeparators, "T")
	mthDayAfterSeparators = append(mthDayAfterSeparators, "")

	return mthDayAfterSeparators, nil

}

func (dtf *DateTimeFormatUtility) getStandardSeparators() ([]string, error) {
	standardSeparators := make([]string, 0, 10)

	standardSeparators = append(standardSeparators, " ")
	standardSeparators = append(standardSeparators, "")

	return standardSeparators, nil
}

func (dtf *DateTimeFormatUtility) getDateTimeSeparators() ([]string, error) {
	dtTimeSeparators := make([]string, 0, 10)

	dtTimeSeparators = append(dtTimeSeparators, " ")
	dtTimeSeparators = append(dtTimeSeparators, ":")
	dtTimeSeparators = append(dtTimeSeparators, "T")
	dtTimeSeparators = append(dtTimeSeparators, "")

	return dtTimeSeparators, nil
}

func (dtf *DateTimeFormatUtility) getTimeElements() ([]string, error) {
	timeElements := make([]string, 0, 512)

	timeElements = append(timeElements, "15:04:05")
	timeElements = append(timeElements, "15:04")
	timeElements = append(timeElements, "15:04:05.000")
	timeElements = append(timeElements, "15:04:05.000000")
	timeElements = append(timeElements, "15:04:05.000000000")

	timeElements = append(timeElements, "03:04:05 pm")
	timeElements = append(timeElements, "03:04 pm")
	timeElements = append(timeElements, "03:04:05.000 pm")
	timeElements = append(timeElements, "03:04:05.000000 pm")
	timeElements = append(timeElements, "03:04:05.000000000 pm")

	timeElements = append(timeElements, "")

	return timeElements, nil
}

func (dtf *DateTimeFormatUtility) getTimeOffsets() ([]string, error) {
	timeOffsetElements := make([]string, 0, 20)

	timeOffsetElements = append(timeOffsetElements, "-0700")
	timeOffsetElements = append(timeOffsetElements, "-07:00")
	timeOffsetElements = append(timeOffsetElements, "-07")
	timeOffsetElements = append(timeOffsetElements, "Z0700")
	timeOffsetElements = append(timeOffsetElements, "Z07:00")
	timeOffsetElements = append(timeOffsetElements, "Z07")
	timeOffsetElements = append(timeOffsetElements, "")

	return timeOffsetElements, nil
}

func (dtf *DateTimeFormatUtility) getTimeOffsetSeparators() ([]string, error) {
	timeOffsetSeparators := make([]string, 0, 20)

	timeOffsetSeparators = append(timeOffsetSeparators, " ")
	timeOffsetSeparators = append(timeOffsetSeparators, "-")
	timeOffsetSeparators = append(timeOffsetSeparators, "")

	return timeOffsetSeparators, nil
}

func (dtf *DateTimeFormatUtility) getTimeZoneElements() ([]string, error) {
	tzElements := make([]string, 0, 20)

	tzElements = append(tzElements, "MST")
	tzElements = append(tzElements, "")

	return tzElements, nil
}

func (dtf *DateTimeFormatUtility) getTimeZoneSeparators() ([]string, error) {
	tzElements := make([]string, 0, 20)

	tzElements = append(tzElements, " ")
	tzElements = append(tzElements, "-")
	tzElements = append(tzElements, "")

	return tzElements, nil
}

func (dtf *DateTimeFormatUtility) getPredefinedFormats() []string {

	preDefinedFormats := make([]string, 0, 20)

	preDefinedFormats = append(preDefinedFormats, time.ANSIC)
	preDefinedFormats = append(preDefinedFormats, time.UnixDate)
	preDefinedFormats = append(preDefinedFormats, time.RubyDate)
	preDefinedFormats = append(preDefinedFormats, time.RFC822)
	preDefinedFormats = append(preDefinedFormats, time.RFC822Z)
	preDefinedFormats = append(preDefinedFormats, time.RFC850)
	preDefinedFormats = append(preDefinedFormats, time.RFC1123)
	preDefinedFormats = append(preDefinedFormats, time.RFC1123Z)
	preDefinedFormats = append(preDefinedFormats, time.RFC3339)
	preDefinedFormats = append(preDefinedFormats, time.RFC3339Nano)
	preDefinedFormats = append(preDefinedFormats, time.Kitchen)
	preDefinedFormats = append(preDefinedFormats, time.Stamp)
	preDefinedFormats = append(preDefinedFormats, time.StampMilli)
	preDefinedFormats = append(preDefinedFormats, time.StampMicro)
	preDefinedFormats = append(preDefinedFormats, time.StampNano)

	return preDefinedFormats
}

func (dtf *DateTimeFormatUtility) getEdgeCases() []string {
	// FmtDateTimeEverything = "Monday January 2, 2006 15:04:05.000000000 -0700 MST"
	edgeCases := make([]string, 0, 20)

	edgeCases = append(edgeCases, "Monday January 2 15:04:05 -0700 MST 2006")

	edgeCases = append(edgeCases, "Mon January 2 15:04:05 -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2 15:04:05 -0700 MST 2006")
	edgeCases = append(edgeCases, "January 2 15:04:05 -0700 MST 2006")

	edgeCases = append(edgeCases, "Monday January 2 15:04 -0700 MST 2006")
	edgeCases = append(edgeCases, "Mon January 2 15:04 -0700 MST 2006")
	edgeCases = append(edgeCases, "Jan 2 15:04 -0700 MST 2006")
	edgeCases = append(edgeCases, "January 2 15:04 -0700 MST 2006")

	edgeCases = append(edgeCases, "January 2 03:04 pm -0700 MST 2006")
	edgeCases = append(edgeCases, "January 2 03:04:05 pm -0700 MST 2006")

	edgeCases = append(edgeCases, "15:04:05 -0700 MST")
	edgeCases = append(edgeCases, "3:04:05 pm -0700 MST")
	edgeCases = append(edgeCases, "15:04 -0700 MST")
	edgeCases = append(edgeCases, "3:04 pm -0700 MST")

	edgeCases = append(edgeCases, "15:04:05 -0700")
	edgeCases = append(edgeCases, "3:04:05 pm -0700")
	edgeCases = append(edgeCases, "15:04 -0700")
	edgeCases = append(edgeCases, "3:04 pm -0700")

	edgeCases = append(edgeCases, "15:04:05")
	edgeCases = append(edgeCases, "3:04:05 pm")
	edgeCases = append(edgeCases, "15:04")
	edgeCases = append(edgeCases, "3:04 pm")

	return edgeCases
}

func (dtf *DateTimeFormatUtility) getPreTrimSearchStrings() [][][]string {
	d := make([][][]string, 0)
	d = append(d, [][]string{{",", " ", "-1"}})
	d = append(d, [][]string{{"/", " ", "-1"}})
	d = append(d, [][]string{{"\\", " ", "-1"}})
	d = append(d, [][]string{{"*", " ", "-1"}})
	d = append(d, [][]string{{"-hrs", ":", "1"}})
	d = append(d, [][]string{{"-mins", ":", "1"}})
	d = append(d, [][]string{{"-secs", "", "1"}})
	d = append(d, [][]string{{"-min", ":", "1"}})
	d = append(d, [][]string{{"-sec", "", "1"}})

	d = append(d, [][]string{{"-Hrs", ":", "1"}})
	d = append(d, [][]string{{"-Mins", ":", "1"}})
	d = append(d, [][]string{{"-Secs", "", "1"}})
	d = append(d, [][]string{{"-Min", ":", "1"}})
	d = append(d, [][]string{{"-Sec", "", "1"}})

	return d
}

func (dtf *DateTimeFormatUtility) getTimeFmtRegEx() [][][]string {
	d := make([][][]string, 0)
	d = append(d, [][]string{{"\\d\\d:\\d\\d:\\d\\d", "%02d:%02d:%02d"}}) // 2:2:2
	d = append(d, [][]string{{"\\d\\d:\\d:\\d\\d", "%02d:%02d:%02d"}})    // 2:1:2
	d = append(d, [][]string{{"\\d\\d:\\d\\d:\\d", "%02d:%02d:%02d"}})    // 2:2:1
	d = append(d, [][]string{{"\\d\\d:\\d:\\d", "%02d:%02d:%02d"}})       // 2:1:1
	d = append(d, [][]string{{"\\d:\\d\\d:\\d\\d", "%02d:%02d:%02d"}})    // 1:2:2
	d = append(d, [][]string{{"\\d:\\d:\\d\\d", "%02d:%02d:%02d"}})       // 1:1:2
	d = append(d, [][]string{{"\\d:\\d\\d:\\d", "%02d:%02d:%02d"}})       // 1:2:1
	d = append(d, [][]string{{"\\d:\\d:\\d", "%02d:%02d:%02d"}})          // 1:1:1

	/*

	   1-	1	1	1
	   2-	1	1	2
	   3-	1	2	2
	   4-	1	2	1
	   5-	2	2	2
	   6-	2	1	1
	   7-	2	2	1
	   8-	2	1	2

	*/

	d = append(d, [][]string{{"\\d\\d:\\d\\d", "%02d:%02d"}}) // 2:2
	d = append(d, [][]string{{"\\d\\d:\\d", "%02d:%02d"}})    // 2:1
	d = append(d, [][]string{{"\\d:\\d\\d", "%02d:%02d"}})    // 1:2
	d = append(d, [][]string{{"\\d:\\d", "%02d:%02d"}})       // 1:1

	/*
	   1- 1:1
	   2- 1:2
	   3- 2:1
	   4- 2:2

	*/

	return d
}

func (dtf *DateTimeFormatUtility) reformatSingleTimeString(targetStr string, regExes [][][]string) string {

	max := len(regExes)

	for i := 0; i < max; i++ {
		re := regexp.MustCompile(regExes[i][0][0])

		idx := re.FindStringIndex(targetStr)

		if idx == nil {
			continue
		}

		s := []byte(targetStr)

		sExtract := string(s[idx[0]:idx[1]])

		timeElements := strings.Split(sExtract, ":")

		lTElements := len(timeElements)

		replaceStr := ""

		for i := 0; i < lTElements; i++ {

			iE, err := strconv.Atoi(timeElements[i])

			if err != nil {
				panic(fmt.Errorf("reformatSingleTimeString() Error converint Time Element %v to ASCII. Error- %v", i, err.Error()))
			}

			if i > 0 {
				replaceStr += ":"
			}

			replaceStr += fmt.Sprintf("%02d", iE)
		}

		return strings.Replace(targetStr, sExtract, replaceStr, 1)

	}

	return targetStr
}

func (dtf *DateTimeFormatUtility) replaceMultipleStrSequence(targetStr string, replaceMap [][][]string) string {

	max := len(replaceMap)

	for i := 0; i < max; i++ {
		if strings.Contains(targetStr, replaceMap[i][0][0]) {
			instances, err := strconv.Atoi(replaceMap[i][0][2])

			if err != nil {
				instances = 1
			}

			targetStr = strings.Replace(targetStr, replaceMap[i][0][0], replaceMap[i][0][1], instances)
		}

	}

	return targetStr
}

func (dtf *DateTimeFormatUtility) replaceAMPM(targetStr string) string {
	d := make([][][]string, 0)

	d = append(d, [][]string{{"\\d{1}\\s{0,4}(?i)a[.]*\\s{0,4}(?i)m[.]*", " am "}})
	d = append(d, [][]string{{"\\d{1}\\s{0,4}(?i)p[.]*\\s{0,4}(?i)m[.]*", " pm "}})

	lD := len(d)

	for i := 0; i < lD; i++ {
		r, err := regexp.Compile(d[i][0][0])

		if err != nil {
			panic(fmt.Errorf("replaceAMPM() Regex failed to Compile. regex== %v. Error: %v", d[i][0][0], err.Error()))
		}

		bTargetStr := []byte(targetStr)

		loc := r.FindIndex(bTargetStr)

		if loc == nil {
			continue
		}

		// Found regex expression

		foundEx := string(bTargetStr[loc[0]+1 : loc[1]])

		return strings.Replace(targetStr, foundEx, d[i][0][1], 1)

	}

	return targetStr
}

func (dtf *DateTimeFormatUtility) replaceDateSuffixStThRd(targetStr string) string {
	// \d{1}\s{0,4}(?i)t\s{0,4}(?i)h
	d := make([][][]string, 0)

	d = append(d, [][]string{{"\\d{1}\\s{0,4}(?i)s\\s{0,4}(?i)t", " "}})
	d = append(d, [][]string{{"\\d{1}\\s{0,4}(?i)n\\s{0,4}(?i)d", " "}})
	d = append(d, [][]string{{"\\d{1}\\s{0,4}(?i)r\\s{0,4}(?i)d", " "}})
	d = append(d, [][]string{{"\\d{1}\\s{0,4}(?i)t\\s{0,4}(?i)h", " "}})

	lD := len(d)

	for i := 0; i < lD; i++ {
		r, err := regexp.Compile(d[i][0][0])

		if err != nil {
			panic(fmt.Errorf("replaceDateSuffixStThRd() Regex failed to Compile. regex== %v. Error: %v", d[i][0][0], err.Error()))
		}

		bTargetStr := []byte(targetStr)

		loc := r.FindIndex(bTargetStr)

		if loc == nil {
			continue
		}

		// Found regex expression

		foundEx := string(bTargetStr[loc[0]+1 : loc[1]])

		return strings.Replace(targetStr, foundEx, d[i][0][1], 1)

	}

	return targetStr
}

// TrimEndMultiple- Performs the following operations on strings:
// 1. Trims Right and Left for all instances of 'trimChar'
// 2. Within the interior of a string, multiple instances
// 		of 'trimChar' are reduce to a single instance.
func (dtf *DateTimeFormatUtility) trimEndMultiple(targetStr string, trimChar rune) (rStr string, err error) {

	if targetStr == "" {
		err = errors.New("trimEndMultiple() - Empty targetStr")
		return
	}

	fStr := []rune(targetStr)
	lenTargetStr := len(fStr)
	outputStr := make([]rune, lenTargetStr)
	lenTargetStr--
	idx := lenTargetStr
	foundFirstChar := false

	for i := lenTargetStr; i >= 0; i-- {

		if !foundFirstChar && fStr[i] == trimChar {
			continue
		}

		if i > 0 && fStr[i] == trimChar && fStr[i-1] == trimChar {
			continue
		}

		if i == 0 && fStr[i] == trimChar {
			continue
		}

		foundFirstChar = true
		outputStr[idx] = fStr[i]
		idx--
	}

	if idx != lenTargetStr {
		idx++
	}

	if outputStr[idx] == trimChar {
		idx++
	}

	result := string(outputStr[idx:])

	return result, nil

}
