package common

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	parseXMLSrcFile    = "main.go"
	parseSMLErrBlockNo = int64(219610000)
)

// ParseXML - Reads and Parses command
// data from an XML file.
func ParseXML(xmlPathFileName string, parent []ErrBaseInfo) (CommandBatch, SpecErr) {
	var isPanic bool

	bi := ErrBaseInfo{}.New(parseXMLSrcFile, "ParseXML", parseSMLErrBlockNo)

	se := SpecErr{}.InitializeBaseInfo(parent, bi)

	var cmds CommandBatch

	xmlFile, err1 := os.Open(xmlPathFileName)

	if err1 != nil {
		isPanic = true
		s := fmt.Sprintf("File Name: %v  - Error opening file: %v", xmlPathFileName, err1.Error())
		return cmds, se.New(s, err1, isPanic, 1005)
	}

	defer xmlFile.Close()

	b, err2 := ioutil.ReadAll(xmlFile)

	if err2 != nil {
		isPanic = true
		s := fmt.Sprintf("Error Reading XML File: %v. Error - %v", xmlPathFileName, err2.Error())
		return cmds, se.New(s, err2, isPanic, 1006)
	}

	xml.Unmarshal(b, &cmds)

	return cmds, se.SignalNoErrors()
}
